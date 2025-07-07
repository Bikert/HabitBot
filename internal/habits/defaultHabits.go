package habits

var (
	StepsHabit = CreateHabitDto{
		BaseHabitDto: BaseHabitDto{
			Name:       "Норма шагов",
			Desc:       "Пройти 10 000 шагов за день",
			Icon:       "👣",
			Color:      "#b2f2bb",
			RepeatType: "daily",
			DaysOfWeek: "",
		},
	}

	MorningExerciseHabit = CreateHabitDto{
		BaseHabitDto: BaseHabitDto{
			Name:       "Зарядка",
			Desc:       "Сделать утреннюю зарядку",
			Icon:       "🏃",
			Color:      "#c77dff",
			RepeatType: "daily",
			DaysOfWeek: "",
		},
	}

	KbjuHabit = CreateHabitDto{
		BaseHabitDto: BaseHabitDto{
			Name:       "Соблюдение КБЖУ",
			Desc:       "Соблюдать нормы калорий, белков, жиров и углеводов",
			Icon:       "🥗",
			Color:      "#ffa94d",
			RepeatType: "daily",
			DaysOfWeek: "",
		},
	}

	SleepHabit = CreateHabitDto{
		BaseHabitDto: BaseHabitDto{
			Name:       "Сон до 23:00",
			Desc:       "Ложиться спать до 23:00",
			Icon:       "🛌",
			Color:      "#63e6be",
			RepeatType: "weekly",
			DaysOfWeek: "wed,fri",
		},
	}

	WorkoutHabit = CreateHabitDto{
		BaseHabitDto: BaseHabitDto{
			Name:       "Тренировка",
			Desc:       "Провести полноценную тренировку",
			Icon:       "💪",
			Color:      "#fa5252",
			RepeatType: "weekly",
			DaysOfWeek: "mon,wed,fri",
		},
	}
)

var DefaultHabits = []CreateHabitDto{
	StepsHabit,
	MorningExerciseHabit,
	KbjuHabit,
	SleepHabit,
	WorkoutHabit,
}
