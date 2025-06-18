package session

type Service interface {
	GetOrCreate(userID int64) *Session
	Save(sess *Session)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetOrCreate(userID int64) *Session {
	var sess = s.repo.Get(userID)
	if sess == nil {
		sess = &Session{
			UserID:       userID,
			NextStep:     "new_user",
			PreviousStep: "",
			Data:         map[string]string{},
		}
	}
	s.repo.Save(sess)
	return sess
}
func (s *service) Save(sess *Session) {
	s.repo.Save(sess)
}
