package fight

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

// NewStaticRound returns a Round without the participants shuffled
// this allows a predictable outcome, which is more useful for tests
func NewStaticRound(ps []*Dwayne) *Round {
	return &Round{
		participants: ps,
		shuffle:      false,
	}
}

func TestRoundFight(t *testing.T) {
	dwayneTheRock := NewNamedDwayne(Rock, "R1")
	dwayneThePaper := NewNamedDwayne(Paper, "P1")
	dwayneTheScissors := NewNamedDwayne(Scissors, "S1")
	dwayneTheRock2 := NewNamedDwayne(Rock, "R2")
	dwayneThePaper2 := NewNamedDwayne(Paper, "P2")
	dwayneTheScissors2 := NewNamedDwayne(Scissors, "S2")

	var tests = []struct {
		ps   []*Dwayne
		want RoundResult
	}{
		// A round with all the Dwaynes fighting another Dwayne of the same
		// species, results in a total loss
		// (they couldn't win in a fight)
		{
			[]*Dwayne{
				dwayneTheRock, dwayneTheRock2,
			},
			RoundResult{
				winners: []*Dwayne{},
				losers: []*Dwayne{
					dwayneTheRock, dwayneTheRock2,
				},
			},
		},

		// A round with all Dwaynes fighting another Dwayne of a different
		// species, results in some winners and some losers
		{
			[]*Dwayne{
				dwayneTheRock, dwayneTheScissors2,
				dwayneThePaper, dwayneTheRock2,
				dwayneTheScissors, dwayneThePaper2,
			},
			RoundResult{
				winners: []*Dwayne{
					dwayneTheRock, dwayneThePaper, dwayneTheScissors,
				},
				losers: []*Dwayne{
					dwayneTheScissors2, dwayneTheRock2, dwayneThePaper2,
				},
			},
		},

		// A round where we initially have two drawers, and one non-participant
		//
		// Participants: [Rock, Rock2, Paper]
		// Fights:
		//   Rock vs Rock2 = Draw
		//   Paper - nobody to fight
		// Results: all drawers
		//
		// Participants: [Paper, Rock, Rock2]
		// Fights:
		//   Paper vs Rock = Paper Win, Rock Loss
		//   Rock2 - nobody to fight
		//
		// Participants: [Rock2]
		// Winners: [Paper]
		// Losers: [Rock]
		// Fights: none
		//
		// Winners: [Paper]
		// Losers: [Rock, Rock2]
		{
			[]*Dwayne{
				dwayneTheRock, dwayneTheRock2, dwayneThePaper,
			},
			RoundResult{
				winners: []*Dwayne{dwayneThePaper},
				losers:  []*Dwayne{dwayneTheRock, dwayneTheRock2},
			},
		},

		// Another simple round with some winners, some losers, and some drawers
		{
			[]*Dwayne{
				dwayneTheRock, dwayneThePaper,
				dwayneThePaper2, dwayneTheRock2,
				dwayneTheScissors, dwayneTheScissors2,
			},
			RoundResult{
				winners: []*Dwayne{dwayneThePaper, dwayneThePaper2},
				losers: []*Dwayne{
					dwayneTheRock, dwayneTheRock2,
					dwayneTheScissors, dwayneTheScissors2,
				},
			},
		},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.ps)
		t.Run(testname, func(t *testing.T) {
			timeout := time.After(1 * time.Millisecond)
			done := make(chan bool)
			go func() {
				round := NewStaticRound(tt.ps)

				ans := round.Fights()
				if !reflect.DeepEqual(ans, tt.want) {
					t.Errorf("got %v, want %v", ans, tt.want)
				}
				done <- true
			}()

			select {
			case <-timeout:
				t.Fatal("Test didn't finish in time")
			case <-done:
			}
		})
	}
}
