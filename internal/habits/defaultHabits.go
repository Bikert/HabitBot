package habits

var (
	StepsHabit = CreateHabitDto{
		BaseHabitDto: BaseHabitDto{
			Name:       "Норма шагов",
			Desc:       "Пройти 7 000 шагов за день",
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
			Color:      "#ffc09f",
			RepeatType: "daily",
			DaysOfWeek: "",
		},
	}

	KbjuHabit = CreateHabitDto{
		BaseHabitDto: BaseHabitDto{
			Name:       "Соблюдение КБЖУ",
			Desc:       "Соблюдать нормы калорий, белков, жиров и углеводов",
			Icon:       "🥗",
			Color:      "#adf7b6",
			RepeatType: "daily",
			DaysOfWeek: "",
		},
	}

	SleepHabit = CreateHabitDto{
		BaseHabitDto: BaseHabitDto{
			Name:       "Сон до 23:00",
			Desc:       "Ложиться спать до 23:00",
			Icon:       "🛌",
			Color:      "#f3c4fb",
			RepeatType: "daily",
		},
	}

	WorkoutHabit = CreateHabitDto{
		BaseHabitDto: BaseHabitDto{
			Name:       "Тренировка",
			Desc:       "Провести полноценную тренировку",
			Icon:       "💪",
			Color:      "#bde0fe",
			RepeatType: "weekly",
			DaysOfWeek: "tue, thu, sat",
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
