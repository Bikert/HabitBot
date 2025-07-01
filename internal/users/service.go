package users

import (
	"HabitMuse/internal/channels"
	"database/sql"
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type Service interface {
	GetOrCreateUser(user tgbotapi.User) (User, error)
	Get(id int64) (User, error)
}

type service struct {
	repo               Repository
	addDefaultHabitsCh chan int64
}

func NewService(repo Repository, channels channels.InitChannels) Service {
	return &service{
		repo:               repo,
		addDefaultHabitsCh: channels.AddDefaultHabitsCh,
	}
}

func (s service) GetOrCreateUser(userTg tgbotapi.User) (User, error) {
	log.Println("GetOrCreateUser")
	if userTg.ID == 0 {
		log.Println("GetOrCreateUser: User ID should not be empty")
		return User{}, errors.New("user ID should not be empty")
	}
	user, err := s.repo.Get(userTg.ID)
	if err == nil {
		log.Println("User already exists")
		return user, nil
	}

	if !errors.Is(err, sql.ErrNoRows) {
		log.Println("GetOrCreateUser: failed to get user: %w", err)
		return User{}, errors.New("failed to get user")
	}
	user, err = s.repo.SaveOrCreate(User{
		UserID:    userTg.ID,
		UserName:  userTg.UserName,
		FirstName: userTg.FirstName,
		LastName:  userTg.LastName,
	})
	if err != nil {
		log.Println("GetOrCreateUser: failed to save user: %w", err)
		return User{}, errors.New("failed to save new user")
	}
	log.Println("GetOrCreateUser: saved user:", user)
	s.addDefaultHabitsCh <- user.UserID
	log.Printf("User ID: %d, Name: %s, FirstName: %s, LastName: %s\n Created!", user.UserID, user.UserName, user.FirstName, user.LastName)
	return user, nil
}

func (s service) Get(id int64) (User, error) {
	return s.repo.Get(id)
}
