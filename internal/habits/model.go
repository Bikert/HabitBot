package habits

import (
	"time"
)

const RepeatTypeDaily = "daily"
const RepeatTypeWeekly = "weekly"

const Monday = "mon"
const Tuesday = "tue"
const Wednesday = "wed"
const Thursday = "thu"
const Friday = "fri"
const Saturday = "sat"
const Sunday = "sun"

type Habit struct {
	ID          int64     `db:"id" json:"id"`
	UserID      int64     `db:"user_id" json:"user_id"`
	GroupID     string    `db:"group_id" json:"group_id"` // UUID в виде строки
	Version     int64     `db:"version" json:"version"`
	Name        string    `db:"name" json:"name"`
	Description string    `db:"description" json:"description,omitempty"`
	Color       string    `db:"color" json:"color,omitempty"`
	Icon        string    `db:"icon" json:"icon,omitempty"`
	IsActive    bool      `db:"is_active" json:"is_active"`
	RepeatType  string    `db:"repeat_type" json:"repeat_type"` // "daily" или "weekly"
	DaysOfWeek  string    `db:"days_of_week" json:"days_of_week"`
	IsDefault   bool      `db:"isDefault" json:"is_default"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
}

type HabitCompletion struct {
	HabitID   int64     `db:"habit_id" json:"habit_id"`
	Date      time.Time `db:"date" json:"date"`
	Completed bool      `db:"completed" json:"completed"`
}

type HabitDto struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	Desc       string `json:"desc"`
	Icon       string `json:"icon"`
	Color      string `json:"color"`
	RepeatType string `json:"repeat_type"`
	DaysOfWeek string `json:"days_of_week"`
}
