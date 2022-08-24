package mongo_db

type Config struct {
	Uri           string `mapstructure:"uri"`
	User          string `mapstructure:"user"`
	Password      string `mapstructure:"password"`
	AuthMechanism string `mapstructure:"auth_mechanism"`
}

func (c *Config) GetUri() string {
	return c.Uri
}

func (c *Config) GetUser() string {
	return c.User
}

func (c *Config) GetPassword() string {
	return c.Password
}

func (c *Config) GetAuthMechanism() string {
	return c.AuthMechanism
}
