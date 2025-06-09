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

	row := m.DB.QueryRowContext(ctx, `SELECT scenario, step, data FROM sessions WHERE user_id = $1`, userID)

	var scenario, step string
	var dataJSON []byte

	err := row.Scan(&scenario, &step, &dataJSON)
	if err == sql.ErrNoRows {
		sess := &Session{
			UserID:   userID,
			Scenario: "welcome",
			Step:     "new_user",
			Data:     map[string]string{},
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
		UserID:   userID,
		Scenario: scenario,
		Step:     step,
		Data:     data,
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
        INSERT INTO sessions (user_id, scenario, step, data)
        VALUES ($1, $2, $3, $4)
        ON CONFLICT (user_id)
        DO UPDATE SET scenario = $2, step = $3, data = $4
    `, sess.UserID, sess.Scenario, sess.Step, dataJSON)

	if err != nil {
		log.Println("Ошибка при сохранении сессии:", err)
	}
}
