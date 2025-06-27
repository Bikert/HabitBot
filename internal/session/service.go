package session

import (
	"HabitMuse/internal/channels"
	"HabitMuse/internal/constants"
	"context"
	"go.uber.org/fx"
	"log"
)

type Service interface {
	GetOrCreate(userID int64) Session
	Save(sess Session)
	CreateSessionForNewUser(id int64)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func UserRegistrationListener(lc fx.Lifecycle, s Service, channels channels.InitChannels) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				for {
					select {
					case newUserID := <-channels.AddDefaultHabitsCh:
						s.CreateSessionForNewUser(newUserID)
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

func (s *service) CreateSessionForNewUser(userId int64) {
	var sess = s.repo.Get(userId)
	if sess == nil {
		sess = &Session{
			UserID:       userId,
			NextStep:     "",
			PreviousStep: "",
			Scenario:     constants.ScenarioWelcome,
			Data:         map[string]string{},
		}
	}
	s.repo.Save(*sess)
	log.Println("saved session:", sess)
}

func (s *service) GetOrCreate(userID int64) Session {
	var sess = s.repo.Get(userID)
	if sess == nil {
		sess = &Session{
			UserID:       userID,
			NextStep:     "",
			PreviousStep: "",
			Scenario:     constants.ScenarioWelcome,
			Data:         map[string]string{},
		}
	}
	s.repo.Save(*sess)
	log.Println("saved session:", sess)
	return *sess
}
func (s *service) Save(sess Session) {
	s.repo.Save(sess)
}
