package tui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type ITreeView interface {
	IBox
	SetRoot(root *tview.TreeNode) *tview.TreeView
	GetRoot() *tview.TreeNode
	SetCurrentNode(node *tview.TreeNode) *tview.TreeView
	GetCurrentNode() *tview.TreeNode
	SetTopLevel(topLevel int) *tview.TreeView
	SetPrefixes(prefixes []string) *tview.TreeView
	SetAlign(align bool) *tview.TreeView
	SetGraphics(showGraphics bool) *tview.TreeView
	SetGraphicsColor(color tcell.Color) *tview.TreeView
	SetChangedFunc(handler func(node *tview.TreeNode)) *tview.TreeView
	SetSelectedFunc(handler func(node *tview.TreeNode)) *tview.TreeView
	SetDoneFunc(handler func(key tcell.Key)) *tview.TreeView
	GetScrollOffset() int
	GetRowCount() int
	Draw(screen tcell.Screen)
	InputHandler() func(event *tcell.EventKey, setFocus func(p tview.Primitive))
	MouseHandler() func(action tview.MouseAction, event *tcell.EventMouse, setFocus func(p tview.Primitive)) (consumed bool, capture tview.Primitive)
}

type TreeView struct {
	*tview.TreeView
}

func NewTreeView() *TreeView {
	return &TreeView{
		TreeView: tview.NewTreeView(),
	}
}
