package fight

// Models an entire round of Dwayne fights
// including:
// - shuffling the list of Dwaynes
//     - handles odd-number dwayne pools (this Dwayne had nobody to fight, so he is classed as a Draw)
// - initiating a Fight
// - Keeping track of round winners, losers, and drawers
// - once all Dwaynes have fought
//     - if all remaining Drawers are of the same species, mark them as losers (they could not win a fight)
//     - brings in drawers, and re-shuffles
