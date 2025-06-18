package users

import (
	"database/sql"
	"errors"
)

type Service interface {
	CreateNewUserFromTG(user UserDTO) (User, error)
	Get(id int64) (User, error)
}

type service struct {
	repo Repository
}

func (s service) Get(id int64) (User, error) {
	return s.repo.Get(id)
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s service) CreateNewUserFromTG(user UserDTO) (User, error) {
	u, err := s.repo.Get(user.UserID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return s.repo.SaveOrCreate(User{
				UserName:  user.UserName,
				FirstName: user.FirstName,
				LastName:  user.LastName,
			})
		}
		return User{}, err
	}
	return u, err
}
