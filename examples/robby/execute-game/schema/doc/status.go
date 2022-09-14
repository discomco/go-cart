package doc

type Status int

const (
	Unknown            Status = 0 << iota
	Initialized        Status = 1
	CompositionStarted Status = 2
	CompositionEnded   Status = 4
	MapInitialized     Status = 8
	CompositionExpired Status = 16
	GameStarted        Status = 32
	GamePauzed         Status = 64
	GameRunning        Status = 96
)
