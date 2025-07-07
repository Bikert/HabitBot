package habits

import (
	"HabitMuse/internal/channels"
	"HabitMuse/internal/users"
	"context"
	"errors"
	"github.com/google/uuid"
	"go.uber.org/fx"
	"log"
	"sort"
	"strings"
	"time"
)

type Service interface {
	GetHabitsByUser(userId int64) ([]*HabitDto, error)
	GetHabitByGroupID(groupId string) *HabitDto
	GetHabitByVersionId(versionId int64) *HabitDto

	CreateHabit(dto *CreateHabitDto, user *users.User) (*HabitDto, error)
	UpdateHabit(groupId string, dto *UpdateHabitDto, user *users.User) (*HabitDto, error)
	CreateBaseHabitsForNewUser(userId int64)
	GetCompletionHabitsForUserByDate(user *users.User, date time.Time) ([]*HabitCompletionDto, error)
	ToggleHabitCompletion(user *users.User, id int64, date time.Time, completed bool) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func UserRegistrationListener(lc fx.Lifecycle, s Service, channels channels.InitChannels) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				for {
					select {
					case newUserID := <-channels.AddDefaultHabitsCh:
						s.CreateBaseHabitsForNewUser(newUserID)
					case <-ctx.Done():
						log.Println("user registration listener stopped")
						return
					}
				}
			}()
			return nil
		},
		OnStop: func(context.Context) error {
			return nil
		},
	})
}

func (s *service) CreateBaseHabitsForNewUser(userId int64) {
	for _, habitDto := range DefaultHabits {
		now := time.Now()
		habitDto.FirstDate = &now
		habit := buildNewModelFromCreateDTO(&habitDto)
		habit.UserID = userId
		habit.CreatedAt = now
		habit.GroupId = GetUuid()
		err := s.repo.SaveHabit(habit)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (s *service) GetHabitsByUser(userId int64) ([]*HabitDto, error) {
	habits := s.repo.GetActiveHabitsByUserID(userId)
	log.Println("ActiveHabitsByUserID:", habits)
	habitsDto := make([]*HabitDto, len(*habits))
	for i, habit := range *habits {
		habitsDto[i] = buildHabitDtoByModel(habit)
	}
	return habitsDto, nil
}

func (s *service) GetHabitByVersionId(id int64) *HabitDto {
	h := s.repo.GetHabitByVersionID(id)
	if h == nil {
		return nil
	}
	return buildHabitDtoByModel(*h)
}

func (s *service) GetHabitByGroupID(groupId string) *HabitDto {
	h := s.repo.GetActiveHabitByGroupID(groupId)
	if h == nil {
		return nil
	}
	return buildHabitDtoByModel(*h)
}

func (s *service) GetCompletionHabitsForUserByDate(user *users.User, date time.Time) ([]*HabitCompletionDto, error) {
	habits := s.repo.GetActiveHabitsByUserIDByDate(user.UserID, date)
	log.Println("ActiveHabitsByUserID:", habits)
	compareHabits := s.repo.GetCompletedHabitsByUserIdAndDate(user.UserID, date)
	habitsMap := make(map[int64]*HabitCompletionDto)

	for _, h := range *habits {
		if h.RepeatType == RepeatTypeDaily {
			habitsMap[h.VersionId] = buildHabitCompletionDtoByModel(h)
			continue
		}
		if h.RepeatType == RepeatTypeWeekly {
			dayOfWeek := date.Weekday()
			shortDay := strings.ToLower(dayOfWeek.String()[:3])

			allowedDays := make(map[string]struct{})
			for _, day := range strings.Split(h.DaysOfWeek, ",") {
				allowedDays[strings.ToLower(strings.TrimSpace(day))] = struct{}{}
			}
			_, ok := allowedDays[shortDay]
			if ok {
				habitsMap[h.VersionId] = buildHabitCompletionDtoByModel(h)
				continue
			}
		}
	}
	for _, ch := range compareHabits {
		if habitsMap[ch.HabitID] != nil {
			habitsMap[ch.HabitID].CompletedDay = ch.Date.Format("2006-01-02")
			habitsMap[ch.HabitID].Completed = ch.Completed
			continue
		}
		h := s.repo.GetHabitByVersionID(ch.HabitID)
		habitsMap[ch.HabitID] = buildHabitCompletionDtoByModel(*h)
		habitsMap[ch.HabitID].CompletedDay = ch.Date.Format("2006-01-02")
		habitsMap[ch.HabitID].Completed = ch.Completed
	}

	return GetValues(habitsMap), nil
}

func (s *service) ToggleHabitCompletion(user *users.User, habitId int64, date time.Time, completed bool) error {
	habit := s.repo.GetHabitByVersionID(habitId)
	if habit == nil {
		return errors.New("Habit not found")
	}
	if user.UserID != habit.UserID {
		return errors.New("habit is not owned by this user")
	}
	if completed {
		return s.repo.SaveOrUpdateCompletion(habitId, date)
	} else {
		return s.repo.DeleteCompletion(habitId, date)
	}
}

func (s *service) CreateHabit(dto *CreateHabitDto, user *users.User) (*HabitDto, error) {
	log.Println("Habit dto = ", dto)
	err := validateHabit(dto.BaseHabitDto)
	if err != nil {
		return nil, err
	}
	habit := buildNewModelFromCreateDTO(dto)
	habit.UserID = user.UserID
	habit.GroupId = GetUuid()
	log.Println("Habit = ", habit)
	err = s.repo.SaveHabit(habit)
	if err != nil {
		return nil, err
	}
	return buildHabitDtoByModel(*habit), nil
}

func (s *service) UpdateHabit(habitGroupId string, dto *UpdateHabitDto, user *users.User) (*HabitDto, error) {
	habit := s.repo.GetActiveHabitByGroupID(habitGroupId)
	if habit == nil {
		return nil, errors.New("Habit not found")
	}

	// ничего не изменилось — не сохраняем
	if !isHabitChanged(dto, habit) {
		return &HabitDto{
			BaseHabitDto: dto.BaseHabitDto,
			GroupId:      dto.GroupId,
			VersionId:    habit.VersionId,
		}, nil
	}

	if user.UserID != habit.UserID {
		return nil, errors.New("habit is not owned by this user")
	}

	err := validateHabit(dto.BaseHabitDto)
	if err != nil {
		return nil, err
	}

	newHabit := *habit
	updateModelFromDTO(&dto.BaseHabitDto, &newHabit)

	habit.IsActive = false
	habit.LastDate = dto.FirstDate
	err = s.repo.SaveHabit(&newHabit)
	if err != nil {
		return nil, err
	}
	err = s.repo.UpdateHabit(habit)
	if err != nil {
		return nil, err
	}
	return buildHabitDtoByModel(newHabit), nil
}

func buildNewModelFromCreateDTO(dto *CreateHabitDto) *Habit {
	var habit Habit
	habit.Version = 1
	habit.Name = dto.Name
	habit.Description = dto.Desc
	habit.Color = dto.Color
	habit.Icon = dto.Icon
	habit.IsActive = true
	habit.RepeatType = dto.RepeatType
	habit.DaysOfWeek = dto.DaysOfWeek
	habit.IsDefault = false
	habit.FirstDate = *dto.FirstDate
	habit.CreatedAt = time.Now()
	return &habit
}

func updateModelFromDTO(dto *BaseHabitDto, habit *Habit) {
	habit.Version = habit.Version + 1
	habit.Name = dto.Name
	habit.Description = dto.Desc
	habit.Color = dto.Color
	habit.Icon = dto.Icon
	habit.IsActive = true
	habit.RepeatType = dto.RepeatType
	habit.DaysOfWeek = dto.DaysOfWeek
	habit.IsDefault = false
	habit.FirstDate = *dto.FirstDate
	habit.CreatedAt = time.Now()
}

func buildHabitCompletionDtoByModel(model Habit) *HabitCompletionDto {
	habit := HabitDto{
		GroupId:   model.GroupId,
		VersionId: model.VersionId,
		BaseHabitDto: BaseHabitDto{
			Name:       model.Name,
			Desc:       model.Description,
			Color:      model.Color,
			Icon:       model.Icon,
			RepeatType: model.RepeatType,
			DaysOfWeek: model.DaysOfWeek,
			FirstDate:  &model.FirstDate,
		},
	}

	return &HabitCompletionDto{
		Habit:        habit,
		Completed:    false,
		CompletedDay: "",
	}
}

func buildHabitDtoByModel(model Habit) *HabitDto {
	return &HabitDto{
		GroupId:   model.GroupId,
		VersionId: model.VersionId,
		BaseHabitDto: BaseHabitDto{
			Name:       model.Name,
			Desc:       model.Description,
			Color:      model.Color,
			Icon:       model.Icon,
			RepeatType: model.RepeatType,
			DaysOfWeek: model.DaysOfWeek,
			FirstDate:  &model.FirstDate,
		},
	}
}

func validateHabit(h BaseHabitDto) error {
	if strings.TrimSpace(h.Name) == "" {
		log.Println("Name is empty")
		return errors.New("name is required")
	}

	if h.RepeatType != RepeatTypeDaily && h.RepeatType != RepeatTypeWeekly {

		log.Println("Repeat type is invalid = ", h.RepeatType)
		return errors.New("repeat_type must be 'daily' or 'weekly'")
	}

	if h.RepeatType == RepeatTypeWeekly {
		if len(h.DaysOfWeek) == 0 {
			log.Println("DaysOfWeek is empty")
			return errors.New("days_of_week is required when repeat_type is 'weekly'")
		}
		if !isValidDaysOfWeek(h.DaysOfWeek) {
			log.Println("DaysOfWeek is invalid")
			return errors.New("days_of_week contains invalid day(s)")
		}
	}
	if h.FirstDate == nil {
		log.Println("FirstDate is empty")
		return errors.New("firstDate is required")
	}
	if h.Color == "" {
		log.Println("Color is empty")
		return errors.New("color is required")
	}
	if h.Icon == "" {
		log.Println("Icon is empty")
		return errors.New("icon is required")
	}
	return nil
}

func isValidDaysOfWeek(days string) bool {
	validDays := map[string]bool{
		Monday: true, Tuesday: true, Wednesday: true, Thursday: true,
		Friday: true, Saturday: true, Sunday: true,
	}

	parts := strings.Split(days, ",")
	for _, day := range parts {
		day = strings.ToLower(strings.TrimSpace(day))
		if !validDays[day] {
			return false
		}
	}

	return true
}

func isHabitChanged(dto *UpdateHabitDto, habit *Habit) bool {
	return dto.Name != habit.Name ||
		dto.Desc != habit.Description ||
		dto.Color != habit.Color ||
		dto.Icon != habit.Icon ||
		dto.RepeatType != habit.RepeatType ||
		dto.DaysOfWeek != habit.DaysOfWeek
}

func GetUuid() string {
	return uuid.New().String()
}

func GetValues(m map[int64]*HabitCompletionDto) []*HabitCompletionDto {
	values := make([]*HabitCompletionDto, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}

	sort.Slice(values, func(i, j int) bool {
		return values[i].Habit.Name < values[j].Habit.Name
	})
	return values
}
