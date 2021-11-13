package fight

import "fmt"

// Dwayne is our fighter
type Dwayne struct {
	species Species
	// TODO: fight history
	// TODO: name
}

func (d Dwayne) String() string {
	return fmt.Sprintf("Dwayne \"The %v\" Johnson", d.species)
}

// NewDwayne returns a pointer to a new Dwayne
func NewDwayne(s Species) *Dwayne {
	return &Dwayne{
		species: s,
	}
}

func (d Dwayne) Fight(opponent *Dwayne) Result {
	// Dwaynes of the same species trivially draw
	if d.species == opponent.species {
		return Draw
	}

	// Handle other cases
	switch d.species {
	case Rock:
		switch opponent.species {
		case Paper:
			return Loss
		case Scissors:
			return Win
		}
	case Paper:
		switch opponent.species {
		case Rock:
			return Win
		case Scissors:
			return Loss
		}
	case Scissors:
		switch opponent.species {
		case Rock:
			return Loss
		case Paper:
			return Win
		}
	}

	// Shouldn't happen, but if we have a mutant Dwayne species, it could happen
	return 42
}
