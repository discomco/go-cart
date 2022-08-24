package http

type HttpConfig struct {
	Port                string   `mapstructure:"port" validate:"required"`
	Development         bool     `mapstructure:"development"`
	BasePath            string   `mapstructure:"basePath" validate:"required"`
	AppPath             string   `mapstructure:"appPath" validate:"required"`
	DebugErrorsResponse bool     `mapstructure:"debugErrorsResponse"`
	IgnoreLogUrls       []string `mapstructure:"ignoreLogUrls"`
}

func (h *HttpConfig) GetPort() string {
	return h.Port
}

func (h *HttpConfig) IsDevelopment() bool {
	return h.Development
}

func (h *HttpConfig) GetBasePath() string {
	return h.BasePath
}

func (h *HttpConfig) GetAppPath() string {
	return h.AppPath
}

func (h *HttpConfig) WithDebugErrorsResponse() bool {
	return h.DebugErrorsResponse
}

func (h *HttpConfig) GetIgnoreLogUrls() []string {
	return h.IgnoreLogUrls
}
