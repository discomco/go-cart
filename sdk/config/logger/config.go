package logger

type Config struct {
	LogLevel string `mapstructure:"level"`
	DevMode  bool   `mapstructure:"devMode"`
	Encoder  string `mapstructure:"encoder"`
}

func (c *Config) GetLogLevel() string {
	return c.LogLevel
}

func (c *Config) IsDevelopment() bool {
	return c.DevMode
}

func (c *Config) GetEncoder() string {
	return c.Encoder
}
