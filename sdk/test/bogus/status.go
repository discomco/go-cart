package bogus

type Status int

const (
	Unknown     Status = 0 << iota
	Initialized Status = 1
	Started     Status = 2
	Stopped     Status = 4
)

func (s Status) HasFlag(flag Status) bool {
	return s|flag == s
}

func (s Status) NotHasFlag(flag Status) bool {
	return !s.HasFlag(flag)
}

func (s Status) Unset(flag Status) Status {
	return s &^ flag
}

func (s Status) Set(flag Status) Status {
	return s | flag
}
