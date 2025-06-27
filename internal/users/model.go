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
	UserID    int64  `json:"userId"`
	UserName  string `json:"userName"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	CreateAt  time.Time
}

type UserTGDTO struct {
	UserID    int64  `json:"id"`
	UserName  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
