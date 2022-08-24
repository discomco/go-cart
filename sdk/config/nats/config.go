package nats

type Config struct {
	Url  string `mapstructure:"url"`
	User string `mapstructure:"user"`
	Pwd  string `mapstructure:"password"`
}

func (c *Config) GetUrl() string {
	return c.Url
}

func (c *Config) GetUser() string {
	return c.User
}

func (c *Config) GetPwd() string {
	return c.Pwd
}
