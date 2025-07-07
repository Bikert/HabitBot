package body_metrics

import (
	"fmt"
	"reflect"
)

type MetricInfo struct {
	Name  string
	Emoji string
	Unit  string
}

var MetricsMeta = map[string]MetricInfo{
	"Weight":        {"Вес", "⚖️", "кг"},
	"BicepsLeft":    {"Бицепс (лев)", "💪", "см"},
	"BicepsRight":   {"Бицепс (прав)", "💪", "см"},
	"Chest":         {"Грудь", "🏋️", "см"},
	"Waist":         {"Талия", "📏", "см"},
	"Belly":         {"Живот", "🧍", "см"},
	"Hips":          {"Бёдра", "🦵", "см"},
	"ThighMaxLeft":  {"Бедро верх (лев)", "🦿", "см"},
	"ThighMaxRight": {"Бедро верх (прав)", "🦿", "см"},
	"ThighLowLeft":  {"Бедро низ (лев)", "🦿", "см"},
	"ThighLowRight": {"Бедро низ (прав)", "🦿", "см"},
}

type BodyProgress struct {
	Weight        *MetricProgress
	BicepsLeft    *MetricProgress
	BicepsRight   *MetricProgress
	Chest         *MetricProgress
	Waist         *MetricProgress
	Belly         *MetricProgress
	Hips          *MetricProgress
	ThighMaxLeft  *MetricProgress
	ThighMaxRight *MetricProgress
	ThighLowLeft  *MetricProgress
	ThighLowRight *MetricProgress
}

type MetricProgress struct {
	Current              *float64 // Последнее измерение
	ChangeFromFirst      *float64 // Разница между первым и текущим
	ChangeFromSecondLast *float64 // Разница между предыдущим и текущим
}

func (s *service) GenerationReport(userId int64) (string, error) {
	metrics, err := s.repo.GetAllByUserID(userId)
	if err != nil {
		return "", err
	}
	if len(metrics) == 0 {
		return "", fmt.Errorf("Для получения отчёта необходимо сначала добавить хотя бы одну запись с параметрами тела. Пожалуйста, заполните метрики и повторите попытку.")
	}
	firstMetric := metrics[0]
	if len(metrics) == 1 {
		str := "🚀 Ты сделал первый шаг — и это уже победа!\n Параметры сохранены:\n"
		bodyProgress := makeBodyProgress(firstMetric, nil, nil)
		str += getMeasurementsString(bodyProgress)
		str += "🏁 Вперёд к цели! Добавь новый замер через несколько дней, чтобы увидеть прогресс ✨"
		return str, nil
	}

	lastMetric := metrics[len(metrics)-1]
	if len(metrics) == 2 {
		str := "📊 Появился первый прогресс! 🎉\n "
		bodyProgress := makeBodyProgress(lastMetric, firstMetric, nil)
		str += getMeasurementsString(bodyProgress)
		str += "🏁 Добавляй новые данные регулярно, чтобы видеть свои успехи ✨"
		return str, nil
	}

	secondLastMetric := metrics[len(metrics)-2]
	str := "📈 У тебя накопилась хорошая статистика! 🎉\n"
	bodyProgress := makeBodyProgress(lastMetric, firstMetric, secondLastMetric)
	str += getMeasurementsString(bodyProgress)
	str += "📊 Продолжай добавлять замеры, чтобы отслеживать динамику и вдохновляться прогрессом.\n🏁 Вперёд к новым вершинам! ✨"
	return str, nil
}

func makeBodyProgress(last, first, secondLast *BodyMetric) *BodyProgress {
	if last == nil {
		return nil
	}

	f := func(m *BodyMetric) *BodyMetric {
		if m == nil {
			return &BodyMetric{}
		}
		return m
	}

	first = f(first)
	secondLast = f(secondLast)

	return &BodyProgress{
		Weight:        makeMetricProgress(last.Weight, first.Weight, secondLast.Weight),
		BicepsLeft:    makeMetricProgress(last.BicepsLeft, first.BicepsLeft, secondLast.BicepsLeft),
		BicepsRight:   makeMetricProgress(last.BicepsRight, first.BicepsRight, secondLast.BicepsRight),
		Chest:         makeMetricProgress(last.Chest, first.Chest, secondLast.Chest),
		Waist:         makeMetricProgress(last.Waist, first.Waist, secondLast.Waist),
		Belly:         makeMetricProgress(last.Belly, first.Belly, secondLast.Belly),
		Hips:          makeMetricProgress(last.Hips, first.Hips, secondLast.Hips),
		ThighMaxLeft:  makeMetricProgress(last.ThighMaxLeft, first.ThighMaxLeft, secondLast.ThighMaxLeft),
		ThighMaxRight: makeMetricProgress(last.ThighMaxRight, first.ThighMaxRight, secondLast.ThighMaxRight),
		ThighLowLeft:  makeMetricProgress(last.ThighLowLeft, first.ThighLowLeft, secondLast.ThighLowLeft),
		ThighLowRight: makeMetricProgress(last.ThighLowRight, first.ThighLowRight, secondLast.ThighLowRight),
	}

}

func makeMetricProgress(last, first, secondLast *float64) *MetricProgress {
	if last == nil {
		return nil
	}
	var changeFromFirst *float64
	if first != nil {
		diff := *last - *first
		changeFromFirst = &diff
	}

	var changeFromSecondLast *float64
	if secondLast != nil {
		diff := *last - *secondLast
		changeFromSecondLast = &diff
	}

	return &MetricProgress{
		Current:              last,
		ChangeFromFirst:      changeFromFirst,
		ChangeFromSecondLast: changeFromSecondLast,
	}
}

func getMeasurementsString(progress *BodyProgress) string {
	result := ""
	t := reflect.TypeOf(*progress)
	valueRef := reflect.ValueOf(*progress)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		valueField := valueRef.Field(i)
		if valueField.Kind() != reflect.Ptr || valueField.IsNil() {
			continue
		}

		val := valueField.Elem().Interface().(MetricProgress)
		info, ok := MetricsMeta[field.Name]
		if !ok {
			info = MetricInfo{
				Name:  field.Name,
				Unit:  "nil",
				Emoji: "❓",
			}
		}
		result += formatMetric(info, val)
	}
	return result
}

func formatMetric(info MetricInfo, progress MetricProgress) string {
	if progress.Current == nil {
		return ""
	}
	currentVal := fmt.Sprintf("%.1f", *progress.Current)
	if progress.ChangeFromFirst == nil {
		return fmt.Sprintf("%s %s: %s %s\n", info.Emoji, info.Name, currentVal, info.Unit)
	}

	f := func(f float64) string {
		if *progress.ChangeFromFirst > 0 {
			return fmt.Sprintf("+%.1f", f)
		}
		return fmt.Sprintf("%.1f", f)
	}
	changeFromFirstVal := f(*progress.ChangeFromFirst)
	if progress.ChangeFromSecondLast == nil {
		return fmt.Sprintf("%s %s: %s %s / %s %s\n", info.Emoji, info.Name, currentVal, info.Unit, changeFromFirstVal, info.Unit)
	}
	changeFromSecondLastVal := f(*progress.ChangeFromSecondLast)
	return fmt.Sprintf("%s %s: %s %s / %s %s, за все время: %s %s\n", info.Emoji, info.Name, currentVal, info.Unit, changeFromSecondLastVal, info.Unit, changeFromFirstVal, info.Unit)
}
