package testing

import (
	"github.com/discomco/go-cart/robby/execute-game/schema"
	"github.com/discomco/go-cart/robby/execute-game/schema/doc"
	"github.com/discomco/go-cart/robby/execute-game/spokes/change_game_details/contract"
	"math/rand"
	"sync"
)

var (
	gameNames = []string{
		"John's Bonanza",
		"All quiet on the Southern Front",
		"Resurrections",
		"The Day after Yesterday",
		"Corpses for Sale",
	}

	gameDescriptions = []string{
		"A game where you can win a ton of prizes",
		"A game where not much is happening",
		"Nobody dies in this game",
		"It must be today then",
		"Everybody dies in this game",
	}
)

func RandomPayload() *contract.Payload {
	seed := rand.Intn(5)
	name := gameNames[seed]
	description := gameDescriptions[seed]
	details := schema.NewDetails(name)
	details.Description = description
	return contract.NewPayload(details)
}

var rMutex = &sync.Mutex{}

func RandomHope() (contract.IHope, error) {
	rMutex.Lock()
	defer rMutex.Unlock()
	aggID, _ := doc.NewGameID()
	pl := RandomPayload()
	return contract.NewHope(aggID.Id(), *pl)
}
