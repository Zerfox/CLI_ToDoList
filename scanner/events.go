package scanner

import "time"

type Event struct {
	Description error
	UserInput   string
	DateAt      time.Time
}

func NewEvent(description error, userInpute string) Event {
	return Event{
		Description: description,
		UserInput:   userInpute,
		DateAt:      time.Now(),
	}
}
