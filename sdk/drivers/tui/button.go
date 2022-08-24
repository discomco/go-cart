package tui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type IButton interface {
	IBox
	SetLabel(label string) *tview.Button
	GetLabel() string
	SetLabelColor(color tcell.Color) *tview.Button
	SetLabelColorActivated(color tcell.Color) *tview.Button
	SetBackgroundColorActivated(color tcell.Color) *tview.Button
	SetSelectedFunc(handler func()) *tview.Button
	SetExitFunc(handler func(key tcell.Key)) *tview.Button
	Draw(screen tcell.Screen)
	InputHandler() func(event *tcell.EventKey, setFocus func(p tview.Primitive))
	MouseHandler() func(action tview.MouseAction, event *tcell.EventMouse, setFocus func(p tview.Primitive)) (consumed bool, capture tview.Primitive)
}

type Button struct {
	*tview.Button
}

func NewButton(label string) *Button {
	return &Button{
		Button: tview.NewButton(label),
	}
}
