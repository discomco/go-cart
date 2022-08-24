package tui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"sync"
)

type ITuiApp interface {
	SetInputCapture(capture func(event *tcell.EventKey) *tcell.EventKey) *tview.Application
	GetInputCapture() func(event *tcell.EventKey) *tcell.EventKey
	SetMouseCapture(capture func(event *tcell.EventMouse, action tview.MouseAction) (*tcell.EventMouse, tview.MouseAction)) *tview.Application
	GetMouseCapture() func(event *tcell.EventMouse, action tview.MouseAction) (*tcell.EventMouse, tview.MouseAction)
	SetScreen(screen tcell.Screen) *tview.Application
	EnableMouse(enable bool) *tview.Application
	Run() error
	Stop()
	Suspend(f func()) bool
	Draw() *tview.Application
	ForceDraw() *tview.Application
	Sync() *tview.Application
	SetBeforeDrawFunc(handler func(screen tcell.Screen) bool) *tview.Application
	GetBeforeDrawFunc() func(screen tcell.Screen) bool
	SetAfterDrawFunc(handler func(screen tcell.Screen)) *tview.Application
	GetAfterDrawFunc() func(screen tcell.Screen)
	SetRoot(root tview.Primitive, fullscreen bool) *tview.Application
	ResizeToFullScreen(p tview.Primitive) *tview.Application
	SetFocus(p tview.Primitive) *tview.Application
	GetFocus() tview.Primitive
	QueueUpdate(f func()) *tview.Application
	QueueUpdateDraw(f func()) *tview.Application
	QueueEvent(event tcell.Event) *tview.Application
}

type App struct {
	*tview.Application
}

var (
	singleTui ITuiApp
	cMutex    = &sync.Mutex{}
)

func NewTui() *App {
	a := &App{}
	b := tview.NewApplication()
	a.Application = b
	return a
}

func SingleTui() ITuiApp {
	if singleTui == nil {
		cMutex.Lock()
		defer cMutex.Unlock()
		a := NewTui()
		singleTui = a
	}
	return singleTui
}
