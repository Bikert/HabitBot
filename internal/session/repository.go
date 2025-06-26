package session

import (
	_ "HabitMuse/internal/db"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/jmoiron/sqlx"
	"log"
)

type Repository interface {
	Save(sess *Session)
	Get(userID int64) *Session
}

type repository struct {
	DB *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{DB: db}
}

func (r *repository) Get(userID int64) *Session {
	query := "SELECT next_step, previous_step, data FROM sessions WHERE user_id = $1"
	// TODO вынести контекст наверх
	ctx := context.Background()
	row := r.DB.QueryRowContext(ctx, query, userID)
	var sess Session
	var dataJSON []byte
	if err := row.Scan(&sess.NextStep, &sess.PreviousStep, &dataJSON); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		log.Fatal(err)
	}

	var data map[string]string
	if err := json.Unmarshal(dataJSON, &data); err != nil {
		log.Println("Ошибка чтения JSON из data:", err)
		data = map[string]string{}
	}
	sess.Data = data

	return &sess
}

func (r *repository) Save(sess *Session) {
	// TODO вынести контекст наверх
	ctx := context.Background()

	dataJSON, err := json.Marshal(sess.Data)
	if err != nil {
		log.Println("Ошибка при сериализации data:", err)
		return
	}

	_, err = r.DB.ExecContext(ctx, `
        INSERT INTO sessions (user_id, next_step, previous_step, data)
        VALUES ($1, $2, $3, $4)
        ON CONFLICT (user_id)
        DO UPDATE SET next_step = $2, previous_step = $3, data = $4
    `, sess.UserID, sess.NextStep, sess.PreviousStep, dataJSON)

	if err != nil {
		log.Println("Ошибка при сохранении сессии:", err)
	}
}
