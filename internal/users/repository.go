package users

import (
	"github.com/jmoiron/sqlx"
	"time"
)

type Repository interface {
	SaveOrCreate(user User) (User, error)
	Get(id int64) (User, error)
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{db: db}
}

func (repo *repository) SaveOrCreate(user User) (User, error) {
	stmt, err := repo.db.Prepare("INSERT OR IGNORE INTO users(id, username, first_name, last_name, created_at) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return user, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.UserID, user.UserName, user.FirstName, user.LastName, time.Now())
	return user, err
}

func (repo *repository) Get(id int64) (User, error) {
	row := repo.db.QueryRow("SELECT id, username, first_name, last_name FROM users WHERE id = ?", id)
	var user User
	err := row.Scan(&user.UserID, &user.UserName, &user.FirstName, &user.LastName)
	return user, err
}
