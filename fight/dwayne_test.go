package fight

import (
	"fmt"
	"testing"
)

func TestDwayneFight(t *testing.T) {
	var tests = []struct {
		a, b *Dwayne
		want Result
	}{
		{NewDwayne(Rock), NewDwayne(Rock), Draw},
		{NewDwayne(Rock), NewDwayne(Paper), Loss},
		{NewDwayne(Rock), NewDwayne(Scissors), Win},
		{NewDwayne(Paper), NewDwayne(Rock), Win},
		{NewDwayne(Paper), NewDwayne(Paper), Draw},
		{NewDwayne(Paper), NewDwayne(Scissors), Loss},
		{NewDwayne(Scissors), NewDwayne(Rock), Loss},
		{NewDwayne(Scissors), NewDwayne(Paper), Win},
		{NewDwayne(Scissors), NewDwayne(Scissors), Draw},
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
