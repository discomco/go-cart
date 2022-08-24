package dtos

type IFbk interface {
	IDto
	GetAggregateId() string
	IsSuccess() bool
	GetErrors() []string
	GetAggregateStatus() int
	SetError(s string)
	SetWarning(s string)
	SetInfo(s string)
}

type Fbk2DataFunc func(fbk IFbk) ([]byte, error)

type Fbk struct {
	*Dto
	AggregateStatus int      `json:"aggregate_status"`
	Errors          []string `json:"errors"`
	Warnings        []string `json:"warnings"`
	Infos           []string `json:"infos"`
}

func (f *Fbk) GetAggregateId() string {
	return f.GetId()
}

func (f *Fbk) SetWarning(s string) {
	f.Warnings = append(f.Warnings, s)
}

func (f *Fbk) IsSuccess() bool {
	return len(f.Errors) == 0
}

func (f *Fbk) GetErrors() []string {
	return f.Errors
}

func (f *Fbk) SetError(s string) {
	f.Errors = append(f.Errors, s)
}

func (f *Fbk) SetInfo(s string) {
	f.Infos = append(f.Infos, s)
}

func (f *Fbk) GetAggregateStatus() int {
	return f.AggregateStatus
}

func NewFbk(aggregateId string, status int, err string) *Fbk {
	d := newDto(aggregateId)
	result := &Fbk{
		AggregateStatus: status,
		Errors:          make([]string, 0),
		Warnings:        make([]string, 0),
		Infos:           make([]string, 0),
	}
	result.Dto = d
	if err != "" {
		result.SetError(err)
	}
	return result
}
