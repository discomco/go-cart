package tui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type ITable interface {
	IBox
	SetContent(content tview.TableContent) *tview.Table
	Clear() *tview.Table
	SetBorders(show bool) *tview.Table
	SetBordersColor(color tcell.Color) *tview.Table
	SetSelectedStyle(style tcell.Style) *tview.Table
	SetSeparator(separator rune) *tview.Table
	SetFixed(rows, columns int) *tview.Table
	SetSelectable(rows, columns bool) *tview.Table
	GetSelectable() (rows, columns bool)
	GetSelection() (row, column int)
	Select(row, column int) *tview.Table
	SetOffset(row, column int) *tview.Table
	GetOffset() (row, column int)
	SetEvaluateAllRows(all bool) *tview.Table
	SetSelectedFunc(handler func(row, column int)) *tview.Table
	SetSelectionChangedFunc(handler func(row, column int)) *tview.Table
	SetDoneFunc(handler func(key tcell.Key)) *tview.Table
	SetCell(row, column int, cell *tview.TableCell) *tview.Table
	SetCellSimple(row, column int, text string) *tview.Table
	GetCell(row, column int) *tview.TableCell
	RemoveRow(row int) *tview.Table
	RemoveColumn(column int) *tview.Table
	InsertRow(row int) *tview.Table
	InsertColumn(column int) *tview.Table
	GetRowCount() int
	GetColumnCount() int
	ScrollToBeginning() *tview.Table
	ScrollToEnd() *tview.Table
	SetWrapSelection(vertical, horizontal bool) *tview.Table
	Draw(screen tcell.Screen)
	InputHandler() func(event *tcell.EventKey, setFocus func(p tview.Primitive))
	MouseHandler() func(action tview.MouseAction, event *tcell.EventMouse, setFocus func(p tview.Primitive)) (consumed bool, capture tview.Primitive)
}

type Table struct {
	*tview.Table
}

func NewTable() *Table {
	return &Table{
		Table: tview.NewTable(),
	}
}
