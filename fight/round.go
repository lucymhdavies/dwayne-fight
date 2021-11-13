package fight

import (
	"fmt"
	"math/rand"
	"time"
)

type Round struct {
	participants []*Dwayne
	shuffle      bool
}

type RoundResult struct {
	winners []*Dwayne
	losers  []*Dwayne
}

func (rr RoundResult) String() string {
	return fmt.Sprintf("Win: %v, Loss: %v", rr.winners, rr.losers)
}

func NewRound(ps []*Dwayne) *Round {
	rand.Seed(time.Now().UnixNano())

	return &Round{
		participants: ps,
		shuffle:      true,
	}
}

func (r Round) Fights() RoundResult {
	fmt.Println("====================")
	fmt.Printf("Round: %v\n", r.participants)
	fmt.Println("====================")

	winners := []*Dwayne{}
	losers := []*Dwayne{}

	// Keep track of participants of last set of fights
	//previousParticipants := []*Dwayne{}

	// While we still have participants...
	for len(r.participants) > 0 {
		// If there is only 1 dwayne, he has nobody to fight, so loses
		if len(r.participants) == 1 {
			fmt.Printf("%v has nobody to fight\n", r.participants[0])
			losers = append(losers, r.participants[0])
			r.participants = []*Dwayne{}
			continue
		}

		// Participants who drew in this set of fights go on to fight again
		// in this round
		drawers := []*Dwayne{}

		// If our participants are all the same species...
		//   nobody can win, so add them all to losers then end round
		firstSpecies := r.participants[0].species
		allSameSpecies := true
		for i := 0; i < len(r.participants); i++ {
			if r.participants[i].species != firstSpecies {
				allSameSpecies = false
				break
			}
		}
		if allSameSpecies {
			losers = append(losers, r.participants...)
			fmt.Printf("All remaining participants %v same species\n",
				r.participants)
			r.participants = []*Dwayne{}
			continue
		}

		if r.shuffle {
			rand.Shuffle(
				len(r.participants),
				func(i, j int) {
					r.participants[i], r.participants[j] =
						r.participants[j], r.participants[i]
				},
			)
		} else {
			// TODO: if our remaining participants are the same as our previous
			// set of participants
			//   we're not shuffling, so we're guaranteed an infinite loop if we
			//   keep going, so mark all remaining participants as losers
		}

		// If we have an odd number of dwaynes, the last one has nobody to fight
		// so we need to add him to the next set of fights
		if len(r.participants)%2 != 0 {
			drawers = append(drawers,
				r.participants[len(r.participants)-1])
			r.participants = r.participants[:len(r.participants)-1]
		}

		// First two Dwaynes fight
		fightResult := Fight(r.participants[0], r.participants[1])
		// and we remove them from participants
		r.participants = r.participants[2:]

		// Add the winner & loser, if they exist, to our winners & losers
		if fightResult.winner != nil {
			winners = append(winners, fightResult.winner)
			losers = append(losers, fightResult.loser)
		}

		// If we have drawers instead, add those to the next set of fights
		drawers = append(drawers, fightResult.drawers...)

		// Now that we have finished all fights, set the fighters for the next
		// set of fights
		r.participants = append(r.participants, drawers...)
	}

	results := RoundResult{
		winners,
		losers,
	}
	fmt.Printf("Results: %v\n", results)

	return results
}
