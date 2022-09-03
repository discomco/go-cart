package avatar

type Status int

const (
	Unknown     Status = 0 << iota //
	Initialized Status = 1
	Thinking    Status = 2
	Attacking   Status = 4
	Defending   Status = 8
	Alive       Status = 16
	Dead        Status = 32
	Moving      Status = 64
	Idle        Status = 128
)
