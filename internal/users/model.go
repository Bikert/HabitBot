package users

import "time"

type User struct {
	UserID    int64
	UserName  string
	FirstName string
	LastName  string
	CreateAt  time.Time
}

type UserDTO struct {
	UserID    int64
	UserName  string
	FirstName string
	LastName  string
	CreateAt  time.Time
}
