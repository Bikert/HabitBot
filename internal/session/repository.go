package session

import (
	_ "HabitMuse/internal/db"
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/jmoiron/sqlx"
	"log"
)

type Repository interface {
	Save(sess Session)
	Get(userID int64) *Session
}

type repository struct {
	DB *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{DB: db}
}

func (r *repository) Get(userID int64) *Session {
	query := "SELECT user_id, next_step, previous_step, data, scenario FROM sessions WHERE user_id = $1"
	row := r.DB.QueryRow(query, userID)
	var sess Session
	var dataJSON []byte
	if err := row.Scan(&sess.UserID, &sess.NextStep, &sess.PreviousStep, &dataJSON, &sess.Scenario); err != nil {
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

func (r *repository) Save(sess Session) {
	dataJSON, err := json.Marshal(sess.Data)
	if err != nil {
		log.Println("Ошибка при сериализации data:", err)
		return
	}

	_, err = r.DB.Exec(`
        INSERT INTO sessions (user_id, next_step, previous_step, data, scenario)
        VALUES ($1, $2, $3, $4, $5)
        ON CONFLICT (user_id)
        DO UPDATE SET next_step = $2, previous_step = $3, data = $4, scenario = $5
    `, sess.UserID, sess.NextStep, sess.PreviousStep, dataJSON, sess.Scenario)

	if err != nil {
		log.Println("Ошибка при сохранении сессии:", err)
	}
	log.Println("Session saved successfully sess = ", sess)
}
