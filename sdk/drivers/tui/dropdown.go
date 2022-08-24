package tui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type IDropDown interface {
	IBox
	SetCurrentOption(index int) *tview.DropDown
	GetCurrentOption() (int, string)
	SetTextOptions(prefix, suffix, currentPrefix, currentSuffix, noSelection string) *tview.DropDown
	SetLabel(label string) *tview.DropDown
	GetLabel() string
	SetLabelWidth(width int) *tview.DropDown
	SetLabelColor(color tcell.Color) *tview.DropDown
	SetFieldBackgroundColor(color tcell.Color) *tview.DropDown
	SetFieldTextColor(color tcell.Color) *tview.DropDown
	SetPrefixTextColor(color tcell.Color) *tview.DropDown
	SetListStyles(unselected, selected tcell.Style) *tview.DropDown
	SetFormAttributes(labelWidth int, labelColor, bgColor, fieldTextColor, fieldBgColor tcell.Color) tview.FormItem
	SetFieldWidth(width int) *tview.DropDown
	GetFieldWidth() int
	AddOption(text string, selected func()) *tview.DropDown
	SetOptions(texts []string, selected func(text string, index int)) *tview.DropDown
	GetOptionCount() int
	RemoveOption(index int) *tview.DropDown
	SetSelectedFunc(handler func(text string, index int)) *tview.DropDown
	SetDoneFunc(handler func(key tcell.Key)) *tview.DropDown
	SetFinishedFunc(handler func(key tcell.Key)) tview.FormItem
	Draw(screen tcell.Screen)
	InputHandler() func(event *tcell.EventKey, setFocus func(p tview.Primitive))
	Focus(delegate func(p tview.Primitive))
	HasFocus() bool
	MouseHandler() func(action tview.MouseAction, event *tcell.EventMouse, setFocus func(p tview.Primitive)) (consumed bool, capture tview.Primitive)
}

type DropDown struct {
	*tview.DropDown
}

func NewDropDown() *DropDown {
	return &DropDown{
		DropDown: tview.NewDropDown(),
	}
}
