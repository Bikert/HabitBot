package session

import (
	"HabitMuse/internal/bot/constants"
	"log"
)

type Service interface {
	GetOrCreateSessionForUser(userID int64) *Session
	Save(sess Session)
	CreateSessionForNewUser(id int64) *Session
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateSessionForNewUser(userId int64) *Session {
	log.Println("CreateSessionForNewUser")
	sess := &Session{
		UserID:       userId,
		NextStep:     "",
		PreviousStep: "",
		Scenario:     constants.ScenarioWelcome,
		Data:         map[string]string{},
	}

	s.repo.Save(*sess)
	log.Println("saved session for NEW user:", sess)
	return sess
}

func (s *service) GetOrCreateSessionForUser(userID int64) *Session {
	var sess = s.repo.Get(userID)
	if sess == nil {
		return s.CreateSessionForNewUser(userID)
	}
	return sess
}
func (s *service) Save(sess Session) {
	s.repo.Save(sess)
}
