package fight

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFight(t *testing.T) {
	dwayneTheRock := NewDwayne(Rock)
	dwayneThePaper := NewDwayne(Paper)
	dwayneTheScissors := NewDwayne(Scissors)

	var tests = []struct {
		a, b *Dwayne
		want FightResult
	}{
		{dwayneTheRock, dwayneTheRock,
			FightResult{
				winner:  nil,
				loser:   nil,
				drawers: []*Dwayne{dwayneTheRock, dwayneTheRock},
			},
		},
		{dwayneTheRock, dwayneThePaper,
			FightResult{
				winner:  dwayneThePaper,
				loser:   dwayneTheRock,
				drawers: []*Dwayne{},
			},
		},
		{dwayneTheRock, dwayneTheScissors,
			FightResult{
				winner:  dwayneTheRock,
				loser:   dwayneTheScissors,
				drawers: []*Dwayne{},
			},
		},

		{dwayneThePaper, dwayneTheRock,
			FightResult{
				winner:  dwayneThePaper,
				loser:   dwayneTheRock,
				drawers: []*Dwayne{},
			},
		},
		{dwayneThePaper, dwayneThePaper,
			FightResult{
				winner:  nil,
				loser:   nil,
				drawers: []*Dwayne{dwayneThePaper, dwayneThePaper},
			},
		},
		{dwayneThePaper, dwayneTheScissors,
			FightResult{
				winner:  dwayneTheScissors,
				loser:   dwayneThePaper,
				drawers: []*Dwayne{},
			},
		},

		{dwayneTheScissors, dwayneTheRock,
			FightResult{
				winner:  dwayneTheRock,
				loser:   dwayneTheScissors,
				drawers: []*Dwayne{},
			},
		},
		{dwayneTheScissors, dwayneThePaper,
			FightResult{
				winner:  dwayneTheScissors,
				loser:   dwayneThePaper,
				drawers: []*Dwayne{},
			},
		},
		{dwayneTheScissors, dwayneTheScissors,
			FightResult{
				winner:  nil,
				loser:   nil,
				drawers: []*Dwayne{dwayneTheScissors, dwayneTheScissors},
			},
		},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v vs %v", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := Fight(tt.b, tt.a)
			if !reflect.DeepEqual(ans, tt.want) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
