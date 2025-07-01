package channels

type InitChannels struct {
	AddDefaultHabitsCh chan int64
}

func NewInitChannels() InitChannels {
	return InitChannels{
		AddDefaultHabitsCh: make(chan int64),
	}
}
