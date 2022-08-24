package tui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type IBox interface {
	IView
	SetBorderPadding(top, bottom, left, right int) *tview.Box
	GetRect() (int, int, int, int)
	GetInnerRect() (int, int, int, int)
	SetRect(x, y, width, height int)
	SetDrawFunc(handler func(screen tcell.Screen, x, y, width, height int) (int, int, int, int)) *tview.Box
	GetDrawFunc() func(screen tcell.Screen, x, y, width, height int) (int, int, int, int)
	WrapInputHandler(inputHandler func(*tcell.EventKey, func(p tview.Primitive))) func(*tcell.EventKey, func(p tview.Primitive))
	InputHandler() func(event *tcell.EventKey, setFocus func(p tview.Primitive))
	SetInputCapture(capture func(event *tcell.EventKey) *tcell.EventKey) *tview.Box
	GetInputCapture() func(event *tcell.EventKey) *tcell.EventKey
	WrapMouseHandler(mouseHandler func(tview.MouseAction, *tcell.EventMouse, func(p tview.Primitive)) (bool, tview.Primitive)) func(action tview.MouseAction, event *tcell.EventMouse, setFocus func(p tview.Primitive)) (consumed bool, capture tview.Primitive)
	MouseHandler() func(action tview.MouseAction, event *tcell.EventMouse, setFocus func(p tview.Primitive)) (consumed bool, capture tview.Primitive)
	SetMouseCapture(capture func(action tview.MouseAction, event *tcell.EventMouse) (tview.MouseAction, *tcell.EventMouse)) *tview.Box
	InRect(x, y int) bool
	GetMouseCapture() func(action tview.MouseAction, event *tcell.EventMouse) (tview.MouseAction, *tcell.EventMouse)
	SetBackgroundColor(color tcell.Color) *tview.Box
	SetBorder(show bool) *tview.Box
	SetBorderColor(color tcell.Color) *tview.Box
	SetBorderAttributes(attr tcell.AttrMask) *tview.Box
	GetBorderAttributes() tcell.AttrMask
	GetBorderColor() tcell.Color
	GetBackgroundColor() tcell.Color
	SetTitle(title string) *tview.Box
	GetTitle() string
	SetTitleColor(color tcell.Color) *tview.Box
	SetTitleAlign(align int) *tview.Box
	Draw(screen tcell.Screen)
	DrawForSubclass(screen tcell.Screen, p tview.Primitive)
	SetFocusFunc(callback func()) *tview.Box
	SetBlurFunc(callback func()) *tview.Box
	Focus(delegate func(p tview.Primitive))
	Blur()
	HasFocus() bool
}

type Box struct {
	*tview.Box
}

func newBox() *Box {
	return &Box{
		Box: tview.NewBox(),
	}
}

func NewBox() *Box {
	return newBox()
}
