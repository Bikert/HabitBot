package habits

import (
	"HabitMuse/internal/users"
	"context"
	"errors"
	"github.com/google/uuid"
	"go.uber.org/fx"
	"log"
	"strings"
	"time"
)

type Service interface {
	CreateHabit(dto *HabitDto, user *users.User) (*Habit, error)
	GetHabitsByUser(userId int64) (*[]Habit, error)
	CreateBaseHabitsForNewUser(userId int64)
	GetHabitById(id int64) *Habit
	GetHabitsForUserByDate(user *users.User, date time.Time) ([]*HabitDto, error)
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

func UserRegistrationListener(lc fx.Lifecycle, s Service, userRegisteredCh chan int64) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				for {
					select {
					case newUserID := <-userRegisteredCh:
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
		habit := buildNewModelFromDTO(&habitDto)
		habit.UserID = userId
		habit.CreatedAt = time.Now()
		err := s.repo.Save(habit)
		if err != nil {
			log.Fatal(err)
		}

	}
}

func (s *service) GetHabitsForUserByDate(user *users.User, date time.Time) ([]*HabitDto, error) {
	habit := s.repo.GetHabitsByUserID(user.UserID)
	compareHabits := s.repo.GetCompletedHabitsByUserIdAndDate(user.UserID, date)
	habitsMap := make(map[int64]*HabitDto)

	for _, h := range *habit {
		if h.RepeatType == RepeatTypeDaily {
			habitsMap[h.ID] = buildDtoByModel(h)
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
				habitsMap[h.ID] = buildDtoByModel(h)
				continue
			}
		}
	}
	for _, ch := range *compareHabits {
		if habitsMap[ch.ID] != nil {
			habitsMap[ch.ID].CompletedDay = date.Format("2006-01-02")
			habitsMap[ch.ID].Completed = true
			continue
		}
	}

	return GetValues(habitsMap), nil
}

func (s *service) ToggleHabitCompletion(user *users.User, habitId int64, date time.Time, completed bool) error {
	habit := s.repo.GetHabitByID(habitId)
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

func (s *service) CreateHabit(dto *HabitDto, user *users.User) (*Habit, error) {
	err := validateHabit(dto)
	if err != nil {
		return nil, err
	}

	habit := s.repo.GetHabitByID(dto.Id)
	if habit == nil || s.repo.HasCompletions(habit.ID) {
		habit = buildNewModelFromDTO(dto)
		habit.UserID = user.UserID
		habit.CreatedAt = time.Now()

		err = s.repo.Save(habit)
		if err != nil {
			return nil, err
		}
		return habit, nil
	}
	if !isHabitChanged(dto, habit) {
		// ничего не изменилось — не сохраняем
		return habit, nil
	}

	newHabit := *habit
	updateModelFromDTO(dto, &newHabit)

	habit.IsActive = false
	err = s.repo.Save(&newHabit)
	if err != nil {
		return nil, err
	}
	err = s.repo.Save(habit)
	if err != nil {
		return nil, err
	}
	return &newHabit, nil
}

func buildNewModelFromDTO(dto *HabitDto) *Habit {
	var habit Habit
	habit.GroupID = GetUuid()
	habit.Version = 1

	habit.Name = dto.Name
	habit.Description = dto.Desc
	habit.Color = dto.Color
	habit.Icon = dto.Icon
	habit.IsActive = true
	habit.RepeatType = dto.RepeatType
	habit.DaysOfWeek = dto.DaysOfWeek
	habit.IsDefault = false
	habit.CreatedAt = time.Now()
	return &habit
}

func updateModelFromDTO(dto *HabitDto, habit *Habit) {
	habit.Version = habit.Version + 1
	habit.Name = dto.Name
	habit.Description = dto.Desc
	habit.Color = dto.Color
	habit.Icon = dto.Icon
	habit.IsActive = true
	habit.RepeatType = dto.RepeatType
	habit.DaysOfWeek = dto.DaysOfWeek
	habit.IsDefault = false
	habit.ID = 0
	habit.CreatedAt = time.Now()
}

func buildDtoByModel(model Habit) *HabitDto {
	return &HabitDto{
		Id:           model.ID,
		Name:         model.Name,
		Desc:         model.Description,
		Color:        model.Color,
		Icon:         model.Icon,
		RepeatType:   model.RepeatType,
		DaysOfWeek:   model.DaysOfWeek,
		Completed:    false,
		CompletedDay: "",
	}
}

func (s *service) GetHabitById(id int64) *Habit {
	return s.repo.GetHabitByID(id)
}

func (s *service) GetHabitsByUser(userId int64) (*[]Habit, error) {
	return s.repo.GetHabitsByUserID(userId), nil
}

func validateHabit(h *HabitDto) error {
	if strings.TrimSpace(h.Name) == "" {
		return errors.New("name is required")
	}

	if h.RepeatType != RepeatTypeDaily && h.RepeatType != RepeatTypeWeekly {
		return errors.New("repeat_type must be 'daily' or 'weekly'")
	}

	if h.RepeatType == RepeatTypeWeekly {
		if len(h.DaysOfWeek) == 0 {
			return errors.New("days_of_week is required when repeat_type is 'weekly'")
		}
		if !isValidDaysOfWeek(h.DaysOfWeek) {
			return errors.New("days_of_week contains invalid day(s)")
		}

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

func isHabitChanged(dto *HabitDto, habit *Habit) bool {
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

func GetValues(m map[int64]*HabitDto) []*HabitDto {
	values := make([]*HabitDto, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}
