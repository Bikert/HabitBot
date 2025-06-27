package habits

import (
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

type Repository interface {
	GetDefaultHabits() *[]Habit
	GetHabitByID(id int64) *Habit
	GetHabitsByUserID(userID int64) *[]Habit

	Save(habit *Habit) error
	GetCompletedHabitsByUserIdAndDate(userId int64, date time.Time) *[]Habit
	HasCompletions(habitID int64) bool
	GetCompletedHabitByHabitIdAndDate(id int64, date time.Time) (*HabitCompletion, error)
	SaveOrUpdateCompletion(id int64, date time.Time) error
	DeleteCompletion(id int64, date time.Time) error
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetHabitByID(id int64) *Habit {
	var h Habit
	err := r.db.Get(&h, `SELECT * FROM habits WHERE id = ?`, id)
	if err != nil {
		log.Printf("GetHabitByID: failed to get habit with id=%d: %v", id, err)
		return nil
	}
	return &h
}

func (r *repository) GetHabitsByUserID(userID int64) *[]Habit {
	var habits []Habit
	err := r.db.Select(&habits, `SELECT * FROM habits WHERE user_id = ?`, userID)
	if err != nil {
		log.Printf("GetHabitsByUserID: failed to query for user_id=%d: %v", userID, err)
		return &[]Habit{}
	}
	return &habits
}

func (r *repository) Save(h *Habit) error {
	if h.ID == 0 {
		res, err := r.db.NamedExec(`
			INSERT INTO habits (
				user_id, group_id, version, name, description,
				color, icon, is_active, repeat_type, days_of_week,
				isDefault, created_at
			) VALUES (
				:user_id, :group_id, :version, :name, :description,
				:color, :icon, :is_active, :repeat_type, :days_of_week,
				:isDefault, :created_at
			)
		`, h)
		if err != nil {
			log.Printf("Save: failed to insert habit: %v", err)
			return err
		}
		lastID, err := res.LastInsertId()
		if err == nil {
			h.ID = lastID
		}
	} else {
		_, err := r.db.NamedExec(`
			INSERT INTO habits (
				id, user_id, group_id, version, name, description,
				color, icon, is_active, repeat_type, days_of_week,
				isDefault, created_at
			) VALUES (
				:id, :user_id, :group_id, :version, :name, :description,
				:color, :icon, :is_active, :repeat_type, :days_of_week,
				:isDefault, :created_at
			)
			ON CONFLICT(id) DO UPDATE SET
				user_id = excluded.user_id,
				group_id = excluded.group_id,
				version = excluded.version,
				name = excluded.name,
				description = excluded.description,
				color = excluded.color,
				icon = excluded.icon,
				is_active = excluded.is_active,
				repeat_type = excluded.repeat_type,
				days_of_week = excluded.days_of_week,
				isDefault = excluded.isDefault,
				created_at = excluded.created_at
		`, h)
		if err != nil {
			log.Printf("Save: failed to update habit ID=%d: %v", h.ID, err)
			return err
		}
	}

	return nil
}

func (r *repository) GetCompletedHabitsByUserIdAndDate(userID int64, date time.Time) *[]Habit {
	var habits []Habit
	err := r.db.Select(&habits, `
		SELECT h.* FROM habits h
		JOIN habit_completions hc ON h.id = hc.habit_id
		WHERE h.user_id = ? AND date(hc.date) = date(?)
	`, userID, date.Format("2006-01-02"))

	if err != nil {
		log.Printf("GetCompletedHabitsByUserIdAndDate: query failed for userID=%d, date=%s: %v",
			userID, date.Format("2006-01-02"), err)
		return &[]Habit{}
	}

	return &habits
}

func (r *repository) GetCompletedHabitByHabitIdAndDate(id int64, date time.Time) (*HabitCompletion, error) {
	var completions HabitCompletion
	query := `
        SELECT habit_id, date, completed
        FROM habit_completions
        WHERE habit_id = $1 AND date = $2
    `
	err := r.db.Get(&completions, query, id, date)
	if err != nil {
		return nil, err
	}
	return &completions, nil
}

func (r *repository) SaveOrUpdateCompletion(habitID int64, date time.Time) error {
	updateQuery := `
        UPDATE habit_completions
        SET completed = TRUE
        WHERE habit_id = $1 AND date = $2
    `
	res, err := r.db.Exec(updateQuery, habitID, date)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		insertQuery := `
            INSERT INTO habit_completions (habit_id, date, completed)
            VALUES ($1, $2, TRUE)
            ON CONFLICT (habit_id, date) DO NOTHING
        `
		_, err := r.db.Exec(insertQuery, habitID, date)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *repository) DeleteCompletion(habitID int64, date time.Time) error {
	query := `DELETE FROM habit_completions WHERE habit_id = $1 AND date = $2`
	_, err := r.db.Exec(query, habitID, date)
	return err
}

func (r *repository) HasCompletions(habitID int64) bool {
	var count int
	err := r.db.Get(&count, `
		SELECT COUNT(*) FROM habit_completions
		WHERE habit_id = ?
	`, habitID)

	if err != nil {
		log.Printf("HasCompletions: error checking completions for habit_id=%d: %v", habitID, err)
		return false
	}

	return count > 0
}

func (r *repository) GetDefaultHabits() *[]Habit {
	var habits []Habit
	err := r.db.Select(&habits, `SELECT * FROM habits WHERE isDefault = 1`)
	if err != nil {
		log.Printf("GetDefaultHabits: failed to query: %v", err)
		return &[]Habit{}
	}
	return &habits
}
