package tui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type IInputField interface {
	IBox
	SetText(text string) *tview.InputField
	GetText() string
	SetLabel(label string) *tview.InputField
	GetLabel() string
	SetLabelWidth(width int) *tview.InputField
	SetPlaceholder(text string) *tview.InputField
	SetLabelColor(color tcell.Color) *tview.InputField
	SetLabelStyle(style tcell.Style) *tview.InputField
	GetLabelStyle() tcell.Style
	SetFieldBackgroundColor(color tcell.Color) *tview.InputField
	SetFieldTextColor(color tcell.Color) *tview.InputField
	SetFieldStyle(style tcell.Style) *tview.InputField
	GetFieldStyle() tcell.Style
	SetPlaceholderTextColor(color tcell.Color) *tview.InputField
	SetPlaceholderStyle(style tcell.Style) *tview.InputField
	GetPlaceholderStyle() tcell.Style
	SetAutocompleteStyles(background tcell.Color, main, selected tcell.Style) *tview.InputField
	SetFormAttributes(labelWidth int, labelColor, bgColor, fieldTextColor, fieldBgColor tcell.Color) tview.FormItem
	SetFieldWidth(width int) *tview.InputField
	GetFieldWidth() int
	SetMaskCharacter(mask rune) *tview.InputField
	SetAutocompleteFunc(callback func(currentText string) (entries []string)) *tview.InputField
	Autocomplete() *tview.InputField
	SetAcceptanceFunc(handler func(textToCheck string, lastChar rune) bool) *tview.InputField
	SetChangedFunc(handler func(text string)) *tview.InputField
	SetDoneFunc(handler func(key tcell.Key)) *tview.InputField
	SetFinishedFunc(handler func(key tcell.Key)) tview.FormItem
	Draw(screen tcell.Screen)
	InputHandler() func(event *tcell.EventKey, setFocus func(p tview.Primitive))
	MouseHandler() func(action tview.MouseAction, event *tcell.EventMouse, setFocus func(p tview.Primitive)) (consumed bool, capture tview.Primitive)
}

type InputField struct {
	*tview.InputField
}

func NewInputField() *InputField {
	return &InputField{
		InputField: tview.NewInputField(),
	}
}
