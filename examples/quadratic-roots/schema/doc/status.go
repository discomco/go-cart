package doc

type Status int

const (
	Unknown                 Status = 0 << iota
	Initialized             Status = 1
	DiscriminatorCalculated Status = 2
	RootsCalculated         Status = 4
)
