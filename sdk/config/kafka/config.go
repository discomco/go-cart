package kafka

//Config is the configuration for the Kafka bus.
type Config struct {
	BootstrapServers string `mapstructure:"bootstrapServers"`
	GroupId          string `mapstructure:"groupId"`
	AutoOffsetReset  string `mapstructure:"autoOffsetReset"`
	RetentionMs      string `mapstructure:"retentionMs"`
}

func (c *Config) GetBootstrapServers() string {
	return c.BootstrapServers
}

func (c *Config) GetGroupId() string {
	return c.GroupId
}

func (c *Config) GetAutoOffsetReset() string {
	return c.AutoOffsetReset
}

func (c *Config) GetRetentionMs() string {
	return c.RetentionMs
}
