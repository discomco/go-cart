package tui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type IFlex interface {
	IBox
	SetDirection(direction int) *tview.Flex
	SetFullScreen(fullScreen bool) *tview.Flex
	AddItem(item tview.Primitive, fixedSize, proportion int, focus bool) *tview.Flex
	RemoveItem(p tview.Primitive) *tview.Flex
	GetItemCount() int
	GetItem(index int) tview.Primitive
	Clear() *tview.Flex
	ResizeItem(p tview.Primitive, fixedSize, proportion int) *tview.Flex
	Draw(screen tcell.Screen)
	Focus(delegate func(p tview.Primitive))
	HasFocus() bool
	MouseHandler() func(action tview.MouseAction, event *tcell.EventMouse, setFocus func(p tview.Primitive)) (consumed bool, capture tview.Primitive)
	InputHandler() func(event *tcell.EventKey, setFocus func(p tview.Primitive))
}

type Flex struct {
	*tview.Flex
}

func NewFlex() *Flex {
	return &Flex{
		Flex: tview.NewFlex(),
	}
}
