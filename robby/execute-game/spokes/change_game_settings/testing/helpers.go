package testing

import (
	"github.com/discomco/go-cart/robby/execute-game/schema"
	"github.com/discomco/go-cart/robby/execute-game/schema/doc"
	"github.com/discomco/go-cart/robby/execute-game/spokes/change_game_settings/contract"
	"math/rand"
	"sync"
)

func RandomPayload() *contract.Payload {
	dimensions := schema.NewDimensions(rand.Intn(42), rand.Intn(42), rand.Intn(42))
	settings := schema.NewSettings(dimensions, rand.Intn(42))
	return contract.NewPayload(settings)
}

var rMutex = &sync.Mutex{}

func RandomHope() (contract.IHope, error) {
	rMutex.Lock()
	defer rMutex.Unlock()
	aggID, _ := doc.NewGameID()
	pl := RandomPayload()
	return contract.NewHope(aggID.Id(), *pl)
}
