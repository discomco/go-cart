package tui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type IModal interface {
	IBox
	SetTextColor(color tcell.Color) *tview.Modal
	SetButtonBackgroundColor(color tcell.Color) *tview.Modal
	SetButtonTextColor(color tcell.Color) *tview.Modal
	SetDoneFunc(handler func(buttonIndex int, buttonLabel string)) *tview.Modal
	SetText(text string) *tview.Modal
	AddButtons(labels []string) *tview.Modal
	ClearButtons() *tview.Modal
	SetFocus(index int) *tview.Modal
	Focus(delegate func(p tview.Primitive))
	HasFocus() bool
	Draw(screen tcell.Screen)
	MouseHandler() func(action tview.MouseAction, event *tcell.EventMouse, setFocus func(p tview.Primitive)) (consumed bool, capture tview.Primitive)
}

type Modal struct {
	*tview.Modal
}

func NewModal() *Modal {
	return &Modal{
		Modal: tview.NewModal(),
	}
}
