package main

import (
	"github.com/lucymhdavies/dwayne-fight/fight"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.DebugLevel)
	fight.NewEvent().Start()
}
