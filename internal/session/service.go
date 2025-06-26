package session

import (
	"HabitMuse/internal/constants"
	"log"
)

type Service interface {
	GetOrCreate(userID int64) Session
	Save(sess Session)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetOrCreate(userID int64) Session {
	var sess = s.repo.Get(userID)
	if sess == nil {
		sess = &Session{
			UserID:       userID,
			NextStep:     "",
			PreviousStep: "",
			Scenario:     constants.Registration.Title,
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
