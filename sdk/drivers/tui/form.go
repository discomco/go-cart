package tui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type IForm interface {
	IBox
	SetItemPadding(padding int) *tview.Form
	SetHorizontal(horizontal bool) *tview.Form
	SetLabelColor(color tcell.Color) *tview.Form
	SetFieldBackgroundColor(color tcell.Color) *tview.Form
	SetFieldTextColor(color tcell.Color) *tview.Form
	SetButtonsAlign(align int) *tview.Form
	SetButtonBackgroundColor(color tcell.Color) *tview.Form
	SetButtonTextColor(color tcell.Color) *tview.Form
	SetFocus(index int) *tview.Form
	AddInputField(label, value string, fieldWidth int, accept func(textToCheck string, lastChar rune) bool, changed func(text string)) *tview.Form
	AddPasswordField(label, value string, fieldWidth int, mask rune, changed func(text string)) *tview.Form
	AddDropDown(label string, options []string, initialOption int, selected func(option string, optionIndex int)) *tview.Form
	AddCheckbox(label string, checked bool, changed func(checked bool)) *tview.Form
	AddButton(label string, selected func()) *tview.Form
	GetButton(index int) *tview.Button
	RemoveButton(index int) *tview.Form
	GetButtonCount() int
	GetButtonIndex(label string) int
	Clear(includeButtons bool) *tview.Form
	ClearButtons() *tview.Form
	AddFormItem(item tview.FormItem) *tview.Form
	GetFormItemCount() int
	GetFormItem(index int) tview.FormItem
	RemoveFormItem(index int) *tview.Form
	GetFormItemByLabel(label string) tview.FormItem
	GetFormItemIndex(label string) int
	GetFocusedItemIndex() (formItem, button int)
	SetCancelFunc(callback func()) *tview.Form
	Draw(screen tcell.Screen)
	Focus(delegate func(p tview.Primitive))
	HasFocus() bool
	MouseHandler() func(action tview.MouseAction, event *tcell.EventMouse, setFocus func(p tview.Primitive)) (consumed bool, capture tview.Primitive)
	InputHandler() func(event *tcell.EventKey, setFocus func(p tview.Primitive))
}

type Form struct {
	*tview.Form
}

func NewForm() *Form {
	return &Form{
		Form: tview.NewForm(),
	}
}
