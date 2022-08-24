package tui

import "github.com/rivo/tview"

type ITableContentReadOnly interface {
	SetCell(row, column int, cell *tview.TableCell)
	RemoveRow(row int)
	RemoveColumn(column int)
	InsertRow(row int)
	InsertColumn(column int)
	Clear()
}
