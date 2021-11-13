package fight

import (
	"fmt"
	"testing"
)

func TestDwayneFight(t *testing.T) {

	dwayneTheRock := NewNamedDwayne(Rock, "R")
	dwayneThePaper := NewNamedDwayne(Paper, "P")
	dwayneTheScissors := NewNamedDwayne(Scissors, "S")

	var tests = []struct {
		a, b *Dwayne
		want Result
	}{
		{dwayneTheRock, dwayneTheRock, Draw},
		{dwayneTheRock, dwayneThePaper, Loss},
		{dwayneTheRock, dwayneTheScissors, Win},
		{dwayneThePaper, dwayneTheRock, Win},
		{dwayneThePaper, dwayneThePaper, Draw},
		{dwayneThePaper, dwayneTheScissors, Loss},
		{dwayneTheScissors, dwayneTheRock, Loss},
		{dwayneTheScissors, dwayneThePaper, Win},
		{dwayneTheScissors, dwayneTheScissors, Draw},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v vs %v", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := tt.a.Fight(tt.b)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
