package habits

import (
	"context"
	"go.uber.org/fx"
	"log"
)

type Service interface {
	CreateNewHabit(Habit) (Habit, error)
	GetHabitsByUser(userId int64) (*[]Habit, error)
	Save(habit *Habit)
	SaveHabitByUser(userId int64, habitID int64)
	SaveDefaultHabitForUser(userId int64)
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
						s.SaveDefaultHabitForUser(newUserID)
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

func (s service) SaveDefaultHabitForUser(userId int64) {
	habits := s.repo.GetDefaultHabits()
	habitIds := make([]int64, len(*habits))
	for i, habit := range *habits {
		habitIds[i] = habit.ID
	}
	s.repo.SaveHabitsByUser(userId, habitIds)
}

func (s service) GetHabitsByUser(userId int64) (*[]Habit, error) {
	return s.repo.GetHabitsByUserID(userId), nil
}

func (s service) Save(habit *Habit) {
	s.repo.Save(habit)
}

func (s service) SaveHabitByUser(userId int64, habitID int64) {
	s.repo.SaveHabitByUser(userId, habitID)
}

func (s service) CreateNewHabit(habit Habit) (Habit, error) {
	//TODO implement me
	panic("implement me")
}
