package tui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type ICheckbox interface {
	IBox
	SetChecked(checked bool) *tview.Checkbox
	IsChecked() bool
	SetLabel(label string) *tview.Checkbox
	GetLabel() string
	SetLabelWidth(width int) *tview.Checkbox
	SetLabelColor(color tcell.Color) *tview.Checkbox
	SetFieldBackgroundColor(color tcell.Color) *tview.Checkbox
	SetFieldTextColor(color tcell.Color) *tview.Checkbox
	SetCheckedString(checked string) *tview.Checkbox
	SetFormAttributes(labelWidth int, labelColor, bgColor, fieldTextColor, fieldBgColor tcell.Color) tview.FormItem
	GetFieldWidth() int
	SetChangedFunc(handler func(checked bool)) *tview.Checkbox
	SetDoneFunc(handler func(key tcell.Key)) *tview.Checkbox
	SetFinishedFunc(handler func(key tcell.Key)) tview.FormItem
	Draw(screen tcell.Screen)
	InputHandler() func(event *tcell.EventKey, setFocus func(p tview.Primitive))
	MouseHandler() func(action tview.MouseAction, event *tcell.EventMouse, setFocus func(p tview.Primitive)) (consumed bool, capture tview.Primitive)
}

type Checkbox struct {
	*tview.Checkbox
}

func NewCheckbox() *Checkbox {
	return &Checkbox{
		Checkbox: tview.NewCheckbox(),
	}
}
