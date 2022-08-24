package tui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type ITreeNode interface {
	Walk(callback func(node, parent *tview.TreeNode) bool) *tview.TreeNode
	SetReference(reference interface{}) *tview.TreeNode
	GetReference() interface{}
	SetChildren(childNodes []*tview.TreeNode) *tview.TreeNode
	GetText() string
	GetChildren() []*tview.TreeNode
	ClearChildren() *tview.TreeNode
	AddChild(node *tview.TreeNode) *tview.TreeNode
	RemoveChild(node *tview.TreeNode) *tview.TreeNode
	SetSelectable(selectable bool) *tview.TreeNode
	SetSelectedFunc(handler func()) *tview.TreeNode
	SetExpanded(expanded bool) *tview.TreeNode
	Expand() *tview.TreeNode
	Collapse() *tview.TreeNode
	ExpandAll() *tview.TreeNode
	CollapseAll() *tview.TreeNode
	IsExpanded() bool
	SetText(text string) *tview.TreeNode
	GetColor() tcell.Color
	SetColor(color tcell.Color) *tview.TreeNode
	SetIndent(indent int) *tview.TreeNode
	GetLevel() int
}

type TreeNode struct {
	*tview.TreeNode
}

func NewTreeNode(text string) *TreeNode {
	return &TreeNode{
		TreeNode: tview.NewTreeNode(text),
	}
}
