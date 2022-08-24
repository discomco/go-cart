package eventstore_db

// Config represents the configuration structure for EventStoreDB.
type Config struct {
	ConnectionString string `mapstructure:"connectionString"`
	User             string `mapstructure:"user"`
	Pwd              string `mapstructure:"password"`
}

// GetConnectionString returns the connection string for EventStoreDB. (i.e. esdb://localhost:2113?tls=false).
func (c *Config) GetConnectionString() string {
	return c.ConnectionString
}

// GetUser returns the user for EventStoreDB.
func (c *Config) GetUser() string {
	return c.User
}

// GetPwd returns the password for EventStoreDB.
func (c *Config) GetPwd() string {
	return c.Pwd
}

type ProjectionConfig struct {
	PoolSize    int    `mapstructure:"poolSize" validate:"required,gte=0"`
	EventPrefix string `mapstructure:"eventPrefix" validate:"required"`
	Name        string `mapstructure:"name" validate:"required"`
	Group       string `mapstructure:"group" validate:"required"`
}

func (c *ProjectionConfig) GetName() string {
	return c.Name
}

func (c *ProjectionConfig) GetGroup() string {
	return c.Group
}

func (c *ProjectionConfig) GetPoolSize() int {
	return c.PoolSize
}

func (c *ProjectionConfig) GetEventPrefix() string {
	return c.EventPrefix
}
