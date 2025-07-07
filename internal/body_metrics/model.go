package body_metrics

import "time"

type BodyMetric struct {
	ID            int64     `db:"id"`
	UserID        int64     `db:"user_id"`
	Date          time.Time `db:"date"`
	Weight        *float64  `db:"weight"`
	BicepsLeft    *float64  `db:"biceps_left"`
	BicepsRight   *float64  `db:"biceps_right"`
	Chest         *float64  `db:"chest"`
	Waist         *float64  `db:"waist"`
	Belly         *float64  `db:"belly"`
	Hips          *float64  `db:"hips"`
	ThighMaxLeft  *float64  `db:"thigh_max_left"`
	ThighMaxRight *float64  `db:"thigh_max_right"`
	ThighLowLeft  *float64  `db:"thigh_low_left"`
	ThighLowRight *float64  `db:"thigh_low_right"`
}

type BodyMetricDTO struct {
	Date          time.Time `json:"date" binding:"required"`
	Weight        *float64  `json:"weight,omitempty"`
	BicepsLeft    *float64  `json:"biceps_left,omitempty"`
	BicepsRight   *float64  `json:"biceps_right,omitempty"`
	Chest         *float64  `json:"chest,omitempty"`
	Waist         *float64  `json:"waist,omitempty"`
	Belly         *float64  `json:"belly,omitempty"`
	Hips          *float64  `json:"hips,omitempty"`
	ThighMaxLeft  *float64  `json:"thigh_max_left,omitempty"`
	ThighMaxRight *float64  `json:"thigh_max_right,omitempty"`
	ThighLowLeft  *float64  `json:"thigh_low_left,omitempty"`
	ThighLowRight *float64  `json:"thigh_low_right,omitempty"`
}
