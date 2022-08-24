package tui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type IList interface {
	IBox
	SetCurrentItem(index int) *tview.List
	GetCurrentItem() int
	SetOffset(items, horizontal int) *tview.List
	GetOffset() (int, int)
	RemoveItem(index int) *tview.List
	SetMainTextColor(color tcell.Color) *tview.List
	SetMainTextStyle(style tcell.Style) *tview.List
	SetSecondaryTextColor(color tcell.Color) *tview.List
	SetSecondaryTextStyle(style tcell.Style) *tview.List
	SetShortcutColor(color tcell.Color) *tview.List
	SetShortcutStyle(style tcell.Style) *tview.List
	SetSelectedTextColor(color tcell.Color) *tview.List
	SetSelectedBackgroundColor(color tcell.Color) *tview.List
	SetSelectedStyle(style tcell.Style) *tview.List
	SetSelectedFocusOnly(focusOnly bool) *tview.List
	SetHighlightFullLine(highlight bool) *tview.List
	ShowSecondaryText(show bool) *tview.List
	SetWrapAround(wrapAround bool) *tview.List
	SetChangedFunc(handler func(index int, mainText string, secondaryText string, shortcut rune)) *tview.List
	SetSelectedFunc(handler func(int, string, string, rune)) *tview.List
	SetDoneFunc(handler func()) *tview.List
	AddItem(mainText, secondaryText string, shortcut rune, selected func()) *tview.List
	InsertItem(index int, mainText, secondaryText string, shortcut rune, selected func()) *tview.List
	GetItemCount() int
	GetItemText(index int) (main, secondary string)
	SetItemText(index int, main, secondary string) *tview.List
	FindItems(mainSearch, secondarySearch string, mustContainBoth, ignoreCase bool) (indices []int)
	Clear() *tview.List
	Draw(screen tcell.Screen)
	InputHandler() func(event *tcell.EventKey, setFocus func(p tview.Primitive))
	MouseHandler() func(action tview.MouseAction, event *tcell.EventMouse, setFocus func(p tview.Primitive)) (consumed bool, capture tview.Primitive)
}

type List struct {
	*tview.List
}

func NewList() *List {
	return &List{
		List: tview.NewList(),
	}
}
