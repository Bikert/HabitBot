package session

type Session struct {
	UserID       int64
	NextStep     string
	PreviousStep string
	Data         map[string]string
}

type Repository interface {
	GetOrCreate(userID int64) *Session
	Save(sess *Session)
}
