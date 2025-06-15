package steps

func GetResponse(key string) string {
	return responses[key]["ru"]
}

var responses = map[string]map[string]string{
	"welcomeMessage": {
		"ru": "%s, Добро пожаловать! Я помогу тебе отслеживать ежедневные привычки \n",
	},

	"askHabitNameMessage": {
		"ru": "📝Давайте создадим новую привычку!\nКак вы хотите её назвать?",
	},
	"newHabitSavedMessage": {
		"ru": "✅ Привычка «%s» сохранена!",
	},
	"mainMenuMessage": {
		"ru": "Главное меню",
	},
}
