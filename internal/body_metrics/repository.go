package body_metrics

import (
	"github.com/jmoiron/sqlx"
	"time"
)

type Repository interface {
	GetByUserIDAndDate(userID int, date time.Time) (*BodyMetric, error)
	GetAllByUserID(userID int64) ([]*BodyMetric, error)
	GetByID(id int64) (*BodyMetric, error)

	Create(model *BodyMetric) (int64, error)
	Update(model *BodyMetric) error
	Delete(id int) error
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(model *BodyMetric) (int64, error) {
	query := `
		INSERT INTO body_metrics
		(user_id, date, weight, biceps_left, biceps_right, chest, waist, belly, hips, thigh_max_left, thigh_max_right, thigh_low_left, thigh_low_right)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	result, err := r.db.Exec(query,
		model.UserID,
		model.Date,
		model.Weight,
		model.BicepsLeft,
		model.BicepsRight,
		model.Chest,
		model.Waist,
		model.Belly,
		model.Hips,
		model.ThighMaxLeft,
		model.ThighMaxRight,
		model.ThighLowLeft,
		model.ThighLowRight,
	)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repository) GetByUserIDAndDate(userID int, date time.Time) (*BodyMetric, error) {
	query := `SELECT * FROM body_metrics WHERE user_id = ? AND date = ?`
	var bm BodyMetric
	err := r.db.Get(&bm, query, userID, date)
	if err != nil {
		return nil, err
	}
	return &bm, nil
}

func (r *repository) Update(model *BodyMetric) error {
	query := `
		UPDATE body_metrics SET
			user_id = ?,
			date = ?,
			weight = ?,
			biceps_left = ?,
			biceps_right = ?,
			chest = ?,
			waist = ?,
			belly = ?,
			hips = ?,
			thigh_max_left = ?,
			thigh_max_right = ?,
			thigh_low_left = ?,
			thigh_low_right = ?
		WHERE id = ?
	`
	_, err := r.db.Exec(query,
		model.UserID,
		model.Date,
		model.Weight,
		model.BicepsLeft,
		model.BicepsRight,
		model.Chest,
		model.Waist,
		model.Belly,
		model.Hips,
		model.ThighMaxLeft,
		model.ThighMaxRight,
		model.ThighLowLeft,
		model.ThighLowRight,
		model.ID,
	)
	return err
}

func (r *repository) Delete(id int) error {
	query := `DELETE FROM body_metrics WHERE id = ?`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *repository) GetByID(id int64) (*BodyMetric, error) {
	query := `
		SELECT id, user_id, date, weight, biceps_left, biceps_right,
		       chest, waist, belly, hips,
		       thigh_max_left, thigh_max_right, thigh_low_left, thigh_low_right
		FROM body_metrics
		WHERE id = ?
	`

	var metric BodyMetric
	err := r.db.Get(&metric, query, id)
	if err != nil {
		return nil, err
	}
	return &metric, nil
}

func (r *repository) GetAllByUserID(userID int64) ([]*BodyMetric, error) {
	query := `
		SELECT id, user_id, date, weight, biceps_left, biceps_right,
		       chest, waist, belly, hips,
		       thigh_max_left, thigh_max_right, thigh_low_left, thigh_low_right
		FROM body_metrics
		WHERE user_id = ?
		ORDER BY date DESC
	`

	var metrics []*BodyMetric
	err := r.db.Select(&metrics, query, userID)
	if err != nil {
		return nil, err
	}
	return metrics, nil
}
