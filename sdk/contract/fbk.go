package contract

type IFbk interface {
	IDto
	GetBehaviorId() string
	IsSuccess() bool
	GetErrors() []string
	GetStatus() int
	SetError(s string)
	SetWarning(s string)
	SetInfo(s string)
	SetStatus(s int) int
}

type FFbk2Data func(fbk IFbk) ([]byte, error)

type Fbk struct {
	*Dto
	Status   int      `json:"status"`
	Errors   []string `json:"errors"`
	Warnings []string `json:"warnings"`
	Infos    []string `json:"infos"`
}

func (f *Fbk) SetStatus(s int) int {
	f.Status = s
	return f.Status
}

func (f *Fbk) GetBehaviorId() string {
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

func (f *Fbk) GetStatus() int {
	return f.Status
}

func NewFbk(behId string, status int, err string) IFbk {
	d := newDto(behId)
	result := &Fbk{
		Status:   status,
		Errors:   make([]string, 0),
		Warnings: make([]string, 0),
		Infos:    make([]string, 0),
	}
	result.Dto = d
	if err != "" {
		result.SetError(err)
	}
	return result
}
