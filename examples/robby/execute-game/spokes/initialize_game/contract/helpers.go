package contract

import (
	"fmt"
	"github.com/discomco/go-cart/examples/robby/execute-game/schema/doc"
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
	names      = []string{"John", "Paul", "George", "Ringo", "Brian", "Mick", "Charlie", "Linda", "Yoko", "Pattie", "Barbara"}
	games      = []string{"Chess", "Checkers", "Risk", "Monopoly", "Bridge", "Poker", "Whist", "Roulette", "Blackjack", "Go"}
	colors     = []string{"Red", "Orange", "Yellow", "Green", "Blue", "Indigo", "Violet"}
	adjectives = []string{"awesome", "dreaded", "sweet", "bitter", "sour", "salty", "overwhelming", "boring"}
	happenings = []string{"game", "marathon", "bonanza", "freak-out", "festival", "sit-in", "morning", "day", "night", "evening"}
)

func randomGameName() string {
	iN1 := rand.Intn(len(names))
	iN2 := rand.Intn(len(names))
	iG1 := rand.Intn(len(games))
	iG2 := rand.Intn(len(games))
	iC1 := rand.Intn(len(colors))
	iA1 := rand.Intn(len(adjectives))
	iH1 := rand.Intn(len(happenings))
	return fmt.Sprintf(
		"%+v and %+v's %+v %+v & %+v,%+v %+v",
		names[iN1],
		names[iN2],
		colors[iC1],
		games[iG1],
		games[iG2],
		adjectives[iA1],
		happenings[iH1],
	)
}

func RandomPayload() *Payload {
	ID, _ := doc.NewGameID()
	name := randomGameName()
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
