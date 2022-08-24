package tui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type IGrid interface {
	IBox
	SetColumns(columns ...int) *tview.Grid
	SetRows(rows ...int) *tview.Grid
	SetSize(numRows, numColumns, rowSize, columnSize int) *tview.Grid
	SetMinSize(row, column int) *tview.Grid
	SetGap(row, column int) *tview.Grid
	SetBorders(borders bool) *tview.Grid
	SetBordersColor(color tcell.Color) *tview.Grid
	AddItem(p tview.Primitive, row, column, rowSpan, colSpan, minGridHeight, minGridWidth int, focus bool) *tview.Grid
	RemoveItem(p tview.Primitive) *tview.Grid
	Clear() *tview.Grid
	SetOffset(rows, columns int) *tview.Grid
	GetOffset() (rows, columns int)
	Focus(delegate func(p tview.Primitive))
	HasFocus() bool
	InputHandler() func(event *tcell.EventKey, setFocus func(p tview.Primitive))
	Draw(screen tcell.Screen)
	MouseHandler() func(action tview.MouseAction, event *tcell.EventMouse, setFocus func(p tview.Primitive)) (consumed bool, capture tview.Primitive)
}

type Grid struct {
	*tview.Grid
}

func NewGrid() *Grid {
	return &Grid{
		Grid: tview.NewGrid(),
	}
}
