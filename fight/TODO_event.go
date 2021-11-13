package fight

import (
	log "github.com/sirupsen/logrus"
)

// Models several rounds of fights
// initiates the starting pool of dwaynes, either randomly, or from config file

type Event struct {
}

// NewEvent initialises a new Init
func NewEvent() *Event {
	// TODO: create a new Event

	return &Event{}
}

func (e Event) Start() {
	log.Info("Welcome to the Dwayne Fight!")
}
