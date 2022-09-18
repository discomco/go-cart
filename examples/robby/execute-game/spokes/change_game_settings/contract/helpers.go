package contract

import (
	"github.com/discomco/go-cart/examples/robby/execute-game/schema"
	"github.com/discomco/go-cart/examples/robby/execute-game/schema/doc"
	"math/rand"
	"sync"
)

func RandomPayload() *Payload {
	dimensions := schema.NewDimensions(rand.Intn(42), rand.Intn(42), rand.Intn(42))
	settings := schema.NewSettings(dimensions, rand.Intn(42))
	return NewPayload(settings)
}

var rMutex = &sync.Mutex{}

func RandomHope() (IHope, error) {
	rMutex.Lock()
	defer rMutex.Unlock()
	aggID, _ := doc.NewGameID()
	pl := RandomPayload()
	return NewHope(aggID.Id(), *pl)
}
