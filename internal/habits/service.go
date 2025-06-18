package habits

type Service interface {
	CreateNewHabit(Habit) (Habit, error)
	GetHabitsByUser(userId int64) (*[]Habit, error)
	Save(habit *Habit)
	SaveHabitByUser(userId int64, habitID int64)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s service) GetHabitsByUser(userId int64) (*[]Habit, error) {
	return s.repo.GetHabitsByUserID(userId), nil
}

func (s service) Save(habit *Habit) {
	s.repo.Save(habit)
}

func (s service) SaveHabitByUser(userId int64, habitID int64) {
	s.repo.SaveHabitByUser(userId, habitID)
}

func (s service) CreateNewHabit(habit Habit) (Habit, error) {
	//TODO implement me
	panic("implement me")
}
