package session

import (
	_ "HabitMuse/internal/db"
	"context"
	"database/sql"
	"encoding/json"
	"log"
)

type SessionService struct {
	DB *sql.DB
}

func InitSessionService(db *sql.DB) *SessionService {
	return &SessionService{DB: db}
}

func (m *SessionService) GetOrCreate(userID int64) *Session {
	ctx := context.Background()

	row := m.DB.QueryRowContext(ctx, `SELECT next_step, previous_step, data FROM sessions WHERE user_id = $1`, userID)

	var next_step, previous_step string
	var dataJSON []byte

	err := row.Scan(&next_step, &previous_step, &dataJSON)
	if err == sql.ErrNoRows {
		sess := &Session{
			UserID:       userID,
			NextStep:     "new_user",
			PreviousStep: "",
			Data:         map[string]string{},
		}
		m.Save(sess)
		return sess
	} else if err != nil {
		log.Println("DB Get error:", err)
		return &Session{UserID: userID, Data: map[string]string{}}
	}

	var data map[string]string
	if err := json.Unmarshal(dataJSON, &data); err != nil {
		log.Println("Ошибка чтения JSON из data:", err)
		data = map[string]string{}
	}

	return &Session{
		UserID:       userID,
		NextStep:     next_step,
		PreviousStep: previous_step,
		Data:         data,
	}
}

func (m *SessionService) Save(sess *Session) {
	ctx := context.Background()

	dataJSON, err := json.Marshal(sess.Data)
	if err != nil {
		log.Println("Ошибка при сериализации data:", err)
		return
	}

	_, err = m.DB.ExecContext(ctx, `
        INSERT INTO sessions (user_id, next_step, previous_step, data)
        VALUES ($1, $2, $3, $4)
        ON CONFLICT (user_id)
        DO UPDATE SET next_step = $2, previous_step = $3, data = $4
    `, sess.UserID, sess.NextStep, sess.PreviousStep, dataJSON)

	if err != nil {
		log.Println("Ошибка при сохранении сессии:", err)
	}
}
