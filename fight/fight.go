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
	fmt.Printf("Fight: %v vs %v", A, B)
	result := A.Fight(B)

	if result == Win {
		fmt.Printf(" - Winner: %v\n", A)
		return FightResult{
			A,
			B,
			[]*Dwayne{},
		}
	}

	if result == Loss {
		fmt.Printf(" - Winner: %v\n", B)
		return FightResult{
			B,
			A,
			[]*Dwayne{},
		}
	}

	fmt.Printf(" - Draw\n")
	return FightResult{
		nil,
		nil,
		[]*Dwayne{A, B},
	}
}
