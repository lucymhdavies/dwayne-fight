package fight

import (
	"fmt"

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
			NewNamedDwayne(Paper, "P1"),
			NewNamedDwayne(Paper, "P2"),
			NewNamedDwayne(Paper, "P3"),
			NewNamedDwayne(Paper, "P4"),
			NewNamedDwayne(Scissors, "S1"),
			NewNamedDwayne(Scissors, "S2"),
			NewNamedDwayne(Scissors, "S3"),
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
			// Prevent infinite loop from terrible code
			log.Fatalf("Too Many Rounds!")
		}

		r := NewNumberedRound(e.participants, roundNum)
		results := r.Fights()

		if len(results.winners) == 0 {
			weHaveAWinner = true
			log.Infof("Nobody Won!")
			fmt.Printf("%v, %v,\n", "Nobody", roundNum)
		} else {
			e.participants = results.winners
		}

		if len(results.winners) == 1 {
			weHaveAWinner = true
			log.Infof("Event Winner: %v", results.winners[0])
			fmt.Printf("%v, %v,\n", results.winners[0].species, roundNum)
		} else {
			e.participants = results.winners
		}
	}

}
