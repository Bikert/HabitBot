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

type BaseHabitDto struct {
	Name       string     `json:"name" binding:"required"`
	Desc       string     `json:"desc"`
	Icon       string     `json:"icon" binding:"required"`
	Color      string     `json:"color" binding:"required"`
	RepeatType string     `json:"repeatType" binding:"required"`
	DaysOfWeek string     `json:"daysOfWeek"`
	FirstDate  *time.Time `json:"firstDate" binding:"required"` //"2025-06-28T00:00:00Z"
}

type HabitDto struct {
	BaseHabitDto
	GroupId   string `json:"id" binding:"required"`
	VersionId int64  `json:"versionId" binding:"required"`
}

type CreateHabitDto struct {
	BaseHabitDto
}

type UpdateHabitDto struct {
	BaseHabitDto
}

type HabitCompletionDto struct {
	Habit        HabitDto `json:"habit"        binding:"required"`
	Completed    bool     `json:"completed"    binding:"required"`
	CompletedDay string   `json:"completedDay" binding:"required" format:"2006-01-02" `
}

type CompletionRequest struct {
	Completed bool `json:"completed"`
}
