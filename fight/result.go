package fight

import "fmt"

// Result is a Win, Loss or Draw
type Result int

const (
	Win Result = iota
	Draw
	Loss
)

func (r Result) String() string {
	switch r {
	case Win:
		return "Win"
	case Draw:
		return "Draw"
	case Loss:
		return "Loss"
	}

	return fmt.Sprint(int(r))
}
