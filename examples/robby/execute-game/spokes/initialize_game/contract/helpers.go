package contract

import (
	"github.com/discomco/go-cart/robby/execute-game/schema/doc"
	"math/rand"
	"sync"
)

var (
	gameNames = []string{
		"John's Bonanza",
		"All quiet on the Southern Front",
		"Resurrection",
		"The Day after Yesterday",
		"Corpses for Sale",
	}
)

func RandomPayload() *Payload {
	ID, _ := doc.NewGameID()
	seed := rand.Intn(5)
	name := gameNames[seed]
	x := rand.Intn(42) + 3
	y := rand.Intn(42) + 3
	z := rand.Intn(42) + 3
	nbrOfPlayers := rand.Intn(12) + 2
	return NewPayload(ID.Id(), name, x, y, z, nbrOfPlayers)
}

var rMutex = &sync.Mutex{}

func RandomHope() (IHope, error) {
	rMutex.Lock()
	defer rMutex.Unlock()
	aggID, _ := doc.NewGameID()
	pl := RandomPayload()
	return NewHope(aggID.Id(), *pl)
}
