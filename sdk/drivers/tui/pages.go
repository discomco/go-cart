package tui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type IPages interface {
	IBox
	SetChangedFunc(handler func()) *tview.Pages
	GetPageCount() int
	AddPage(name string, item tview.Primitive, resize, visible bool) *tview.Pages
	AddAndSwitchToPage(name string, item tview.Primitive, resize bool) *tview.Pages
	RemovePage(name string) *tview.Pages
	HasPage(name string) bool
	ShowPage(name string) *tview.Pages
	HidePage(name string) *tview.Pages
	SwitchToPage(name string) *tview.Pages
	SendToFront(name string) *tview.Pages
	SendToBack(name string) *tview.Pages
	GetFrontPage() (name string, item tview.Primitive)
	HasFocus() bool
	Focus(delegate func(p tview.Primitive))
	Draw(screen tcell.Screen)
	MouseHandler() func(action tview.MouseAction, event *tcell.EventMouse, setFocus func(p tview.Primitive)) (consumed bool, capture tview.Primitive)
	InputHandler() func(event *tcell.EventKey, setFocus func(p tview.Primitive))
}

// Slide is a function which returns the slide's main primitive and its title.
// It receives a "nextSlide" function which can be called to advance the
// presentation to the next slide.
type Slide func(nextSlide func()) (title string, content tview.Primitive)

type Pages struct {
	*tview.Pages
}

func newPages() *Pages {
	return &Pages{
		Pages: tview.NewPages(),
	}
}

func NewPages() *Pages {
	return newPages()
}
