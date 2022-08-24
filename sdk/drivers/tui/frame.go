package tui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type IFrame interface {
	IBox
	AddText(text string, header bool, align int, color tcell.Color) *tview.Frame
	Clear() *tview.Frame
	SetBorders(top, bottom, header, footer, left, right int) *tview.Frame
	Draw(screen tcell.Screen)
	Focus(delegate func(p tview.Primitive))
	HasFocus() bool
	MouseHandler() func(action tview.MouseAction, event *tcell.EventMouse, setFocus func(p tview.Primitive)) (consumed bool, capture tview.Primitive)
	InputHandler() func(event *tcell.EventKey, setFocus func(p tview.Primitive))
}

type Frame struct {
	*tview.Frame
}

func NewFrame(primitive tview.Primitive) *Frame {
	return &Frame{
		Frame: tview.NewFrame(primitive),
	}
}
