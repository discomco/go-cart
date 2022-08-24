package tui

type ITextViewWriter interface {
	Close() error
	Clear()
	Write(p []byte) (n int, err error)
	HasFocus() bool
}
