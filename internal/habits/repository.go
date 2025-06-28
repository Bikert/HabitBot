package habits

import (
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

type Repository interface {
	GetDefaultHabits() *[]Habit
	GetHabitByVersionID(versionId int64) *Habit
	GetActiveHabitByGroupID(groupId string) *Habit
	GetActiveHabitsByUserID(userID int64) *[]Habit
	GetActiveHabitsByUserIDByDate(id int64, date time.Time) *[]Habit
	GetCompletedHabitsByUserIdAndDate(userId int64, date time.Time) []*HabitCompletion
	GetCompletedHabitByHabitIdAndDate(id int64, date time.Time) (*HabitCompletion, error)

	SaveHabit(h *Habit) error
	UpdateHabit(h *Habit) error

	HasCompletions(habitID int64) bool
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

func (r *repository) GetHabitByVersionID(versionId int64) *Habit {
	var h Habit
	err := r.db.Get(&h, `SELECT * FROM habits WHERE version_id = ?`, versionId)
	if err != nil {
		log.Printf("GetHabitByVersionID: failed to get habit with versionId=%d: %v", versionId, err)
		return nil
	}
	return &h
}

func (r *repository) GetActiveHabitByGroupID(groupId string) *Habit {
	var h Habit
	err := r.db.Get(&h, `SELECT * FROM habits WHERE group_id = ? AND is_active = 1`, groupId)
	if err != nil {
		log.Printf("GetHabitByVersionID: failed to get habit with group_id=%d: %v", groupId, err)
		return nil
	}
	return &h
}

func (r *repository) GetActiveHabitsByUserID(userID int64) *[]Habit {
	var habits []Habit
	err := r.db.Select(&habits, `SELECT * FROM habits WHERE user_id = ? AND is_active = 1`, userID)
	if err != nil {
		log.Printf("GetActiveHabitsByUserID: failed to query for user_id=%d: %v", userID, err)
		return &[]Habit{}
	}
	return &habits
}

func (r *repository) GetActiveHabitsByUserIDByDate(userID int64, date time.Time) *[]Habit {
	var habits []Habit
	err := r.db.Select(&habits, `SELECT * FROM habits WHERE user_id = ? 
                       AND date(?) >= date(first_date)
                       AND (last_date IS NULL OR date(?) < date(last_date))
                       `, userID, date, date)
	if err != nil {
		log.Printf("GetActiveHabitsByUserID: failed to query for user_id=%d: %v", userID, err)
		return &[]Habit{}
	}
	return &habits
}

func (r *repository) SaveHabit(h *Habit) error {
	query := `
		INSERT INTO habits (
			user_id, group_id, version, name, description,
			color, icon, repeat_type, days_of_week,
			isDefault, is_active, first_date, last_date
		)
		VALUES (
			:user_id, :group_id, :version, :name, :description,
			:color, :icon, :repeat_type, :days_of_week,
			:isDefault, :is_active, :first_date, :last_date
		)
	`
	res, err := r.db.NamedExec(query, h)
	if err != nil {
		log.Printf("SaveHabit: failed to insert habit: %v", err)
		return err
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		log.Printf("SaveHabit: failed to insert last id: %v", err)
		return err
	}
	h.VersionId = lastID
	return nil
}

func (r *repository) UpdateHabit(h *Habit) error {
	query := `
		UPDATE habits SET
			version = :version,
			name = :name,
			description = :description,
			color = :color,
			icon = :icon,
			repeat_type = :repeat_type,
			days_of_week = :days_of_week,
			isDefault = :isDefault,
			is_active = :is_active,
			first_date = :first_date,
			last_date = :last_date
		WHERE version_id = :version_id
	`
	_, err := r.db.NamedExec(query, h)
	if err != nil {
		log.Printf("UpdateHabit: failed to update habit: %v", err)
		return err
	}
	return nil
}

func (r *repository) GetCompletedHabitsByUserIdAndDate(userID int64, date time.Time) []*HabitCompletion {
	var habits []*HabitCompletion
	err := r.db.Select(&habits, `
		SELECT hc.* FROM habits h
		JOIN habit_completions hc ON h.version_id = hc.habit_version_id
		WHERE h.user_id = ? AND date(hc.date) = date(?)
	`, userID, date)

	if err != nil {
		log.Printf("GetCompletedHabitsByUserIdAndDate: query failed for userID=%d, date=%s: %v",
			userID, date.Format("2006-01-02"), err)
		return []*HabitCompletion{}
	}

	return habits
}

func (r *repository) GetCompletedHabitByHabitIdAndDate(id int64, date time.Time) (*HabitCompletion, error) {
	var completions HabitCompletion
	query := `
        SELECT habit_version_id, date, completed
        FROM habit_completions
        WHERE habit_version_id = ? AND AND date(date) = date(?)
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
        WHERE habit_version_id = $1 AND date = $2
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
            INSERT INTO habit_completions (habit_version_id, date, completed)
            VALUES ($1, $2, TRUE)
            ON CONFLICT (habit_version_id, date) DO NOTHING
        `
		_, err := r.db.Exec(insertQuery, habitID, date)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *repository) DeleteCompletion(habitID int64, date time.Time) error {
	query := `DELETE FROM habit_completions WHERE habit_version_id = $1 AND date = $2`
	_, err := r.db.Exec(query, habitID, date)
	return err
}

func (r *repository) HasCompletions(habitID int64) bool {
	var count int
	err := r.db.Get(&count, `
		SELECT COUNT(*) FROM habit_completions
		WHERE habit_version_id = ?
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
