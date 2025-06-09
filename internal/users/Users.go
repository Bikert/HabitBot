package users

import (
	"database/sql"
	"time"
)

type User struct {
	UserID    int64
	UserName  string
	FirstName string
	LastMane  string
	CreateAt  time.Time
}

type Repository interface {
	SaveOrCreate(id int64, username, firstName, lastName string) *User
	Get(id int64)
}

type Service struct {
	db *sql.DB
}

func InitUserService(db *sql.DB) *Service {
	return &Service{db: db}
}

func (userManager *Service) SaveOrCreate(user User) (User, error) {
	stmt, err := userManager.db.Prepare("INSERT OR IGNORE INTO users(id, username, first_name, last_name, created_at) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return user, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.UserID, user.UserName, user.FirstName, user.LastMane, time.Now())
	return user, err
}

func (userManager *Service) Get(id int64) (User, error) {
	row := userManager.db.QueryRow("SELECT id, username, first_name, last_name FROM users WHERE id = ?", id)
	var user User
	err := row.Scan(&user.UserID, &user.UserName, &user.FirstName, &user.LastMane)
	return user, err
}
