package habits

type Habit struct {
	ID        int64
	Title     string
	IsDefault bool
}

type HabitDTO struct {
	id           int64
	title        string
	isDefault    bool
	color        string
	description  string
	repeatType   string // daily / weekly
	selectedDays []string
}

type UserHabit struct {
	UserID  int
	HabitID int
}
