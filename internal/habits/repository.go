package habits

import (
	"database/sql"
	"log"
)

type Repository interface {
	GetDefaultHabits() *[]Habit
	GetHabitByID(id int64) *Habit
	GetHabitsByUserID(userID int64) *[]Habit

	Save(habit *Habit)
	SaveHabitByUser(userID int64, habitID int64)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetHabitByID(id int64) *Habit {

	row := r.db.QueryRow("SELECT id, title, isDefault FROM habits WHERE id = ?", id)
	var h Habit
	if err := row.Scan(&h.ID, &h.Title); err != nil {
		log.Println("GetHabitByID error:", err)
		return nil
	}
	return &h
}

func (r *repository) GetDefaultHabits() *[]Habit {

	rows, err := r.db.Query(`
        SELECT id, title, isDefault
        FROM habits
        WHERE isDefault = 1
    `)
	if err != nil {
		log.Println("GetDefaultHabits error:", err)
		return nil
	}
	defer rows.Close()

	var habits []Habit
	for rows.Next() {
		var h Habit
		var isDefaultInt int
		if err := rows.Scan(&h.ID, &h.Title, &isDefaultInt); err != nil {
			log.Println("Scan error:", err)
			continue
		}
		h.IsDefault = isDefaultInt == 1
		habits = append(habits, h)
	}
	return &habits
}

func (r *repository) GetHabitsByUserID(userID int64) *[]Habit {
	query := `
        SELECT h.id, h.title, h.isDefault
        FROM habits h
        INNER JOIN user_habits uh ON h.id = uh.habit_id
        WHERE uh.user_id = ?
    `

	rows, err := r.db.Query(query, userID)
	if err != nil {
		log.Println("GetHabitsByUserID query error:", err)
		return nil
	}
	defer rows.Close()

	var habits []Habit
	for rows.Next() {
		var h Habit
		var isDefaultInt int
		if err := rows.Scan(&h.ID, &h.Title, &isDefaultInt); err != nil {
			log.Println("Scan error:", err)
			continue
		}
		h.IsDefault = isDefaultInt == 1
		habits = append(habits, h)
	}

	return &habits
}

func (r *repository) Save(habit *Habit) {
	result, err := r.db.Exec(
		"INSERT INTO habits (title, isDefault) VALUES (?, ?)",
		habit.Title,
		boolToInt(habit.IsDefault),
	)
	if err != nil {
		log.Println("Save error:", err)
		return
	}
	id, _ := result.LastInsertId()
	habit.ID = id
}

func (r *repository) SaveHabitByUser(userId int64, habitId int64) {
	_, err := r.db.Exec(
		"INSERT INTO user_habits (user_id, habit_id) VALUES (?, ?)", userId, habitId,
	)
	if err != nil {
		log.Println("Save error:", err)
		return
	}
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
