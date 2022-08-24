package jaeger

type Config struct {
	ServiceName string `mapstructure:"serviceName"`
	HostPort    string `mapstructure:"hostPort"`
	Enable      bool   `mapstructure:"enable"`
	LogSpans    bool   `mapstructure:"logSpans"`
}

func (c *Config) GetServiceName() string {
	return c.ServiceName
}

func (c *Config) GetHostPort() string {
	return c.HostPort
}

func (c *Config) IsEnabled() bool {
	return c.Enable
}

func (c *Config) UseLogSpans() bool {
	return c.LogSpans
}
