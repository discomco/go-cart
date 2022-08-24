package tui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type ITableCell interface {
	SetText(text string) *tview.TableCell
	SetAlign(align int) *tview.TableCell
	SetMaxWidth(maxWidth int) *tview.TableCell
	SetExpansion(expansion int) *tview.TableCell
	SetTextColor(color tcell.Color) *tview.TableCell
	SetBackgroundColor(color tcell.Color) *tview.TableCell
	SetTransparency(transparent bool) *tview.TableCell
	SetAttributes(attr tcell.AttrMask) *tview.TableCell
	SetStyle(style tcell.Style) *tview.TableCell
	SetSelectable(selectable bool) *tview.TableCell
	SetReference(reference interface{}) *tview.TableCell
	GetReference() interface{}
	GetLastPosition() (x, y, width int)
	SetClickedFunc(clicked func() bool) *tview.TableCell
}

type TableCell struct {
	*tview.TableCell
}

func NewTableCell(text string) *TableCell {
	return &TableCell{
		TableCell: tview.NewTableCell(text),
	}
}
