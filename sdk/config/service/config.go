package service

import "fmt"

type Config struct {
	NameSpace string `mapstructure:"nameSpace"`
	SubSystem string `mapstructure:"subSystem"`
}

func (c *Config) GetNamespace() string {
	return c.NameSpace
}

func (c *Config) GetSubSystem() string {
	return c.SubSystem
}

func (c *Config) GetServiceName() string {
	return fmt.Sprintf("%s.%s", c.NameSpace, c.SubSystem)
}
