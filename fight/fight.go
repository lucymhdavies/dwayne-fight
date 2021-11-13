package fight

import "fmt"

// Type Fight
// Models fights between individual Dwaynes, outputting winners, losers and drawers

// function Fight(dwayne A, dwayne B) (winner, loser, []drawers)

type FightResult struct {
	winner  *Dwayne
	loser   *Dwayne
	drawers []*Dwayne
}

func (fr FightResult) String() string {
	if len(fr.drawers) > 0 {
		return "Draw"
	}

	return fmt.Sprintf("Win: %v, Loss: %v", fr.winner, fr.loser)
}

func Fight(A, B *Dwayne) FightResult {
	result := A.Fight(B)

	if result == Win {
		return FightResult{
			A,
			B,
			[]*Dwayne{},
		}
	}

	if result == Loss {
		return FightResult{
			B,
			A,
			[]*Dwayne{},
		}
	}

	return FightResult{
		nil,
		nil,
		[]*Dwayne{A, B},
	}
}
