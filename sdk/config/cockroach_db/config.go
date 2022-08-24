package cockroach_db

type Config struct {
	DSN string `mapstructure:"dsn"`
}

func (c *Config) GetDSN() string {
	return c.DSN
}
