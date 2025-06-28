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
	VersionId   int64      `db:"version_id" `
	UserID      int64      `db:"user_id" `
	GroupId     string     `db:"group_id" `
	Version     int64      `db:"version" `
	Name        string     `db:"name" `
	Description string     `db:"description" `
	Color       string     `db:"color"`
	Icon        string     `db:"icon" `
	IsActive    bool       `db:"is_active" `
	RepeatType  string     `db:"repeat_type" `
	DaysOfWeek  string     `db:"days_of_week" `
	IsDefault   bool       `db:"isDefault" `
	FirstDate   time.Time  `db:"first_date" `
	LastDate    *time.Time `db:"last_date" `
	CreatedAt   time.Time  `db:"created_at" `
}

type HabitCompletion struct {
	HabitID   int64     `db:"habit_version_id"`
	Date      time.Time `db:"date"`
	Completed bool      `db:"completed"`
}

type HabitDto struct {
	GroupId    string     `json:"id"`
	VersionId  int64      `json:"versionId"`
	Name       string     `json:"name"`
	Desc       string     `json:"desc"`
	Icon       string     `json:"icon"`
	Color      string     `json:"color"`
	RepeatType string     `json:"repeatType"`
	DaysOfWeek string     `json:"daysOfWeek"`
	FirstDate  *time.Time `json:"firstDate"` //"2025-06-28T00:00:00Z"
}

type HabitCompletionDto struct {
	Habit        HabitDto `json:"habit"`
	Completed    bool     `json:"completed"`
	CompletedDay string   `json:"completedDay" format:"2006-01-02"`
}

type CompletionRequest struct {
	Completed bool `json:"completed"`
}
