package habits

var (
	StepsHabit = HabitDto{
		Name:       "Норма шагов",
		Desc:       "Пройти 10 000 шагов за день",
		Icon:       "👣",
		Color:      "#b2f2bb",
		RepeatType: "daily",
		DaysOfWeek: "",
	}

	MorningExerciseHabit = HabitDto{
		Name:       "Зарядка",
		Desc:       "Сделать утреннюю зарядку",
		Icon:       "🏃",
		Color:      "#c77dff",
		RepeatType: "daily",
		DaysOfWeek: "",
	}

	KbjuHabit = HabitDto{
		Name:       "Соблюдение КБЖУ",
		Desc:       "Соблюдать нормы калорий, белков, жиров и углеводов",
		Icon:       "🥗",
		Color:      "#ffa94d",
		RepeatType: "daily",
		DaysOfWeek: "",
	}

	SleepHabit = HabitDto{
		Name:       "Сон до 23:00",
		Desc:       "Ложиться спать до 23:00",
		Icon:       "🛌",
		Color:      "#63e6be",
		RepeatType: "weekly",
		DaysOfWeek: "wed,fri",
	}

	WorkoutHabit = HabitDto{
		Name:       "Тренировка",
		Desc:       "Провести полноценную тренировку",
		Icon:       "💪",
		Color:      "#fa5252",
		RepeatType: "weekly",
		DaysOfWeek: "mon,wed,fri",
	}
)

var DefaultHabits = []HabitDto{
	StepsHabit,
	MorningExerciseHabit,
	KbjuHabit,
	SleepHabit,
	WorkoutHabit,
}
