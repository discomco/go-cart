package list

import "github.com/discomco/go-cart/robby/execute-game/schema/doc"

type GameList struct {
	items map[string]*GameItem `json:"items"`
}

type GameItem struct {
	Names           string     `json:"names"`
	NumberOfPlayers int        `json:"numberOfPlayers"`
	Status          doc.Status `json:"status"`
}
