package session

type Session struct {
	UserID       int64
	NextStep     string
	PreviousStep string
	Scenario     string
	Data         map[string]string
}
