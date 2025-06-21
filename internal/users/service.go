package users

import (
	"HabitMuse/internal/http"
	"database/sql"
	"errors"
	"log"
)

type Service interface {
	RegistrationUser(user UserTGDTO) (User, error)
	Get(id int64) (User, error)
}

type service struct {
	repo             Repository
	userRegisteredCh chan int64
}

func NewService(repo Repository, userRegisteredCh chan int64) Service {
	return &service{
		repo:             repo,
		userRegisteredCh: userRegisteredCh,
	}
}

func (s service) RegistrationUser(userTg UserTGDTO) (User, error) {
	log.Println("RegistrationUser")
	if userTg.UserID == 0 {
		return User{}, http.ErrBadRequest("RegistrationUser: User ID should not be empty", nil)
	}
	log.Println("RegistrationUser UserID = ", userTg.UserID)
	user, err := s.repo.Get(userTg.UserID)
	if err == nil {
		log.Println("User already exists")
		return user, nil
	}

	if !errors.Is(err, sql.ErrNoRows) {
		log.Println("RegistrationUser: failed to get user: %w", err)
		return User{}, http.ErrInternal("RegistrationUser: failed to get user: %w", err)
	}
	log.Println("RegistrationUser SaveOrCreate started ")
	user, err = s.repo.SaveOrCreate(User{
		UserID:    userTg.UserID,
		UserName:  userTg.UserName,
		FirstName: userTg.FirstName,
		LastName:  userTg.LastName,
	})
	log.Println("RegistrationUser SaveOrCreate finished")

	if err != nil {
		log.Println("RegistrationUser: failed to save user: %w", err)
		return User{}, http.ErrInternal("failed to save user", err)
	}
	log.Printf("Befor userRegisteredCh")
	s.userRegisteredCh <- user.UserID
	log.Printf("User ID: %d, Name: %s, FirstName: %s, LastName: %s\n Created!", user.UserID, user.UserName, user.FirstName, user.LastName)
	return user, nil
}

func (s service) Get(id int64) (User, error) {
	return s.repo.Get(id)
}
