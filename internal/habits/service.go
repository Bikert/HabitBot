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
	habit.CreatedAt = time.Now()
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

func GetUuid() string {
	return uuid.New().String()
}
