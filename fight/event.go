package fight

import (
	log "github.com/sirupsen/logrus"
)

// Models several rounds of fights
// initiates the starting pool of dwaynes, either randomly, or from config file

type Event struct {
	participants []*Dwayne
}

// NewEvent initialises a new Init
func NewEvent() *Event {
	// TODO: set random participants, or read from config file if one exists

	return &Event{
		participants: []*Dwayne{
			NewNamedDwayne(Rock, "R1"),
			NewNamedDwayne(Rock, "R2"),
			NewNamedDwayne(Rock, "R3"),
			NewNamedDwayne(Rock, "R4"),
			NewNamedDwayne(Rock, "R5"),
			NewNamedDwayne(Rock, "R6"),
			NewNamedDwayne(Rock, "R7"),
			NewNamedDwayne(Rock, "R8"),
			NewNamedDwayne(Rock, "R9"),
			NewNamedDwayne(Rock, "R10"),
			NewNamedDwayne(Rock, "R11"),
			NewNamedDwayne(Rock, "R12"),
			NewNamedDwayne(Rock, "R13"),
			NewNamedDwayne(Rock, "R14"),
			NewNamedDwayne(Rock, "R15"),
			NewNamedDwayne(Rock, "R16"),
			NewNamedDwayne(Paper, "P1"),
			NewNamedDwayne(Paper, "P2"),
			NewNamedDwayne(Paper, "P3"),
			NewNamedDwayne(Paper, "P4"),
			NewNamedDwayne(Scissors, "S1"),
			NewNamedDwayne(Scissors, "S2"),
			NewNamedDwayne(Scissors, "S3"),
			NewNamedDwayne(Scissors, "S4"),
			NewNamedDwayne(Scissors, "S5"),
			NewNamedDwayne(Scissors, "S6"),
			NewNamedDwayne(Scissors, "S7"),
			NewNamedDwayne(Scissors, "S8"),
			NewNamedDwayne(Scissors, "S9"),
			NewNamedDwayne(Scissors, "S10"),
			NewNamedDwayne(Scissors, "S11"),
			NewNamedDwayne(Scissors, "S12"),
		},
	}
}

func (e Event) Start() {
	log.Info("Welcome to the Dwayne Fight!")
	log.Infof("Today's participants: %v", e.participants)

	roundNum := 0
	weHaveAWinner := false

	for !weHaveAWinner {
		roundNum++
		if roundNum > 50 {
			log.Fatalf("Too Many Rounds!")
		}

		r := NewNumberedRound(e.participants, roundNum)
		results := r.Fights()

		if len(results.winners) == 0 {
			weHaveAWinner = true
			log.Infof("Nobody Won!")
		} else {
			e.participants = results.winners
		}

		if len(results.winners) == 1 {
			weHaveAWinner = true
			log.Infof("Event Winner: %v", results.winners[0])
		} else {
			e.participants = results.winners
		}
	}

}
