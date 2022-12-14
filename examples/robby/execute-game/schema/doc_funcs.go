package schema

import (
	sdk_model "github.com/discomco/go-cart/sdk/schema"
)

func (r *GameDoc) GetStatus() int {
	return int(r.Status)
}

func DocFtor() sdk_model.DocFtor[GameDoc] {
	return func() *GameDoc {
		return newGameDoc()
	}
}

func NewGameDoc() *GameDoc {
	return newGameDoc()
}

func newGameDoc() *GameDoc {
	return &GameDoc{
		ID:      nil,
		Details: NewDetails("New GameDoc"),
	}
}

func NewDetails(name string) *Details {
	return &Details{
		Name: name,
	}
}

func NewDimensions(x, y, z int) *Dimensions {
	return &Dimensions{
		X: x,
		Y: y,
		Z: z,
	}
}

func NewSettings(mapSize *Dimensions, nbrOfPlayers int) *Settings {
	return &Settings{
		MapSize:      mapSize,
		NbrOfPlayers: nbrOfPlayers,
	}
}
