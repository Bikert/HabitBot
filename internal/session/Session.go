package session

type Session struct {
	UserID   int64
	Step     string
	Scenario string
	Data     map[string]string
}

type Repository interface {
	GetOrCreate(userID int64) *Session
	Save(sess *Session)
}
