package fight

import "fmt"

// Species represents the rockiness, paperiness or scissoriness of a Dwayne
type Species int

const (
	Rock Species = iota
	Paper
	Scissors
)

func (s Species) String() string {
	switch s {
	case Rock:
		return "Rock"
	case Paper:
		return "Paper"
	case Scissors:
		return "Scissors"
	}

	return fmt.Sprint(int(s))
}
