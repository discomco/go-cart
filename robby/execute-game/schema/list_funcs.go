package schema

import (
	"github.com/discomco/go-cart/robby/execute-game/schema/list"
	"github.com/discomco/go-cart/sdk/schema"
)

func NewGameList() *GameList {
	return &GameList{
		ID:    list.DefaultID(),
		Items: make(map[string]*GameItem),
	}
}

func NewGameItem(Id string, name string) *GameItem {
	return &GameItem{
		Id:   Id,
		Name: name,
	}
}

func ListFtor() schema.DocFtor[GameList] {
	return func() *GameList {
		return NewGameList()
	}
}

func (l *GameList) GetItem(Id string) *GameItem {
	it, ok := l.Items[Id]
	if !ok {
		it = NewGameItem(Id, "unknown")
		l.AddItem(it)
	}
	return it
}

func (l *GameList) AddItem(item *GameItem) {
	if l.Items == nil {
		l.Items = make(map[string]*GameItem)
	}
	l.Items[item.Id] = item
}

func (l *GameList) DeleteItem(key string) {
	delete(l.Items, key)
}

func GameListKey() string {
	ID := list.DefaultID()
	return ID.Id()
}
