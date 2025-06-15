package constants

const MainMenu = "main_menu"

type RegistrationSteps struct {
	Title            string
	VerifyProfile    string
	FirstNameReceive string
	LastNameReceive  string
}

type HabitCreationSteps struct {
	Title              string
	AskTitle           string
	ReceiveNameAndSave string
}

var Registration = RegistrationSteps{
	Title:            "registration",
	VerifyProfile:    "verify_profile",
	FirstNameReceive: "receive_first_name",
	LastNameReceive:  "receive_last_name",
}

var HabitCreation = HabitCreationSteps{
	Title:              "habit_creation",
	AskTitle:           "ask_title",
	ReceiveNameAndSave: "receive_name_and_save",
}
