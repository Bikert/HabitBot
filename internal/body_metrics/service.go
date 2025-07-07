package body_metrics

import (
	"fmt"
	"log"
)

type Service interface {
	GetMetricsList(userId int64) ([]*BodyMetricDTO, error)
	Create(userId int64, dto *BodyMetricDTO) (*BodyMetricDTO, error)
	Update(userId int64, metricId int64, dto *BodyMetricDTO) error
	GetReport(userId int64) (string, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetMetricsList(userId int64) ([]*BodyMetricDTO, error) {
	metrics, err := s.repo.GetAllByUserID(userId)
	if err != nil {
		return nil, err
	}

	result := make([]*BodyMetricDTO, 0, len(metrics))
	for _, m := range metrics {
		dto := mapToDTO(m)
		result = append(result, dto)
	}
	return result, nil
}

func (s *service) Create(userId int64, dto *BodyMetricDTO) (*BodyMetricDTO, error) {
	m := mapToModel(dto)
	m.UserID = userId
	id, err := s.repo.Create(m)
	if err != nil {
		return nil, err
	}
	log.Printf("BodyMetric id = %d Created = %+v\n", id, dto)
	return dto, nil
}

func (s *service) Update(userId int64, metricId int64, dto *BodyMetricDTO) error {
	existing, err := s.repo.GetByID(metricId)
	if err != nil {
		return err
	}
	if existing.UserID != userId {
		return fmt.Errorf("user ID mismatch")
	}
	m := mapToModel(dto)
	m.UserID = userId
	m.ID = metricId
	return s.repo.Update(m)
}

func (s *service) GetReport(userId int64) (string, error) {
	metrics, err := s.repo.GetAllByUserID(userId)
	if err != nil {
		return "", err
	}
	if len(metrics) == 0 {
		return "", fmt.Errorf("Для получения отчёта необходимо сначала добавить хотя бы одну запись с параметрами тела. Пожалуйста, заполните метрики и повторите попытку.")
	}
	lastMetric, firstMetric := metrics[0], metrics[0]
	for _, m := range metrics {
		if m.Date.After(lastMetric.Date) {
			lastMetric = m
		}
		if m.Date.Before(firstMetric.Date) {
			firstMetric = m
		}
	}
	str := fmt.Sprintf("Дата вашего последнего измерения: %s, \n"+
		"Дата вашего первого измерения: %s, \n", lastMetric.Date, firstMetric.Date)
	return str, nil
}

func mapToDTO(m *BodyMetric) *BodyMetricDTO {
	return &BodyMetricDTO{
		Date:          m.Date,
		Weight:        &m.Weight,
		BicepsLeft:    &m.BicepsLeft,
		BicepsRight:   &m.BicepsRight,
		Chest:         &m.Chest,
		Waist:         &m.Waist,
		Belly:         &m.Belly,
		Hips:          &m.Hips,
		ThighMaxLeft:  &m.ThighMaxLeft,
		ThighMaxRight: &m.ThighMaxRight,
		ThighLowLeft:  &m.ThighLowLeft,
		ThighLowRight: &m.ThighLowRight,
	}
}

func mapToModel(dto *BodyMetricDTO) *BodyMetric {
	return &BodyMetric{
		Date:          dto.Date,
		Weight:        *dto.Weight,
		BicepsLeft:    *dto.BicepsLeft,
		BicepsRight:   *dto.BicepsRight,
		Chest:         *dto.Chest,
		Waist:         *dto.Waist,
		Belly:         *dto.Belly,
		Hips:          *dto.Hips,
		ThighMaxLeft:  *dto.ThighMaxLeft,
		ThighMaxRight: *dto.ThighMaxRight,
		ThighLowLeft:  *dto.ThighLowLeft,
		ThighLowRight: *dto.ThighLowRight,
	}
}
