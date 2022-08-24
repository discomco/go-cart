package tui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type ITextView interface {
	IBox
	SetScrollable(scrollable bool) *tview.TextView
	SetWrap(wrap bool) *tview.TextView
	SetWordWrap(wrapOnWords bool) *tview.TextView
	SetMaxLines(maxLines int) *tview.TextView
	SetTextAlign(align int) *tview.TextView
	SetTextColor(color tcell.Color) *tview.TextView
	SetText(text string) *tview.TextView
	GetText(stripAllTags bool) string
	GetOriginalLineCount() int
	SetDynamicColors(dynamic bool) *tview.TextView
	SetRegions(regions bool) *tview.TextView
	SetChangedFunc(handler func()) *tview.TextView
	SetDoneFunc(handler func(key tcell.Key)) *tview.TextView
	SetHighlightedFunc(handler func(added, removed, remaining []string)) *tview.TextView
	ScrollTo(row, column int) *tview.TextView
	ScrollToBeginning() *tview.TextView
	ScrollToEnd() *tview.TextView
	GetScrollOffset() (row, column int)
	Clear() *tview.TextView
	Highlight(regionIDs ...string) *tview.TextView
	GetHighlights() (regionIDs []string)
	SetToggleHighlights(toggle bool) *tview.TextView
	ScrollToHighlight() *tview.TextView
	GetRegionText(regionID string) string
	Focus(delegate func(p tview.Primitive))
	HasFocus() bool
	Write(p []byte) (n int, err error)
	BatchWriter() tview.TextViewWriter
	Draw(screen tcell.Screen)
	InputHandler() func(event *tcell.EventKey, setFocus func(p tview.Primitive))
	MouseHandler() func(action tview.MouseAction, event *tcell.EventMouse, setFocus func(p tview.Primitive)) (consumed bool, capture tview.Primitive)
}

type TextView struct {
	*tview.TextView
}

func NewTextView() *TextView {
	return &TextView{
		TextView: tview.NewTextView(),
	}
}
