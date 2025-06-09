package habits

type Habit struct {
	ID        int64
	Title     string
	IsDefault bool
}

type UserHabit struct {
	UserID  int
	HabitID int
}

type Repository interface {
	GetDefaultHabits() *[]Habit
	GetHabitByID(id int64) *Habit
	GetHabitsByUserID(userID int64) *[]Habit

	Save(habit *Habit)
	SetToUser(userID int64, habitID int64)
}
