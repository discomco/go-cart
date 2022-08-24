package probes

type Config struct {
	ReadinessPath        string `mapstructure:"readinessPath"`
	LivenessPath         string `mapstructure:"livenessPath"`
	Port                 string `mapstructure:"port"`
	Pprof                string `mapstructure:"pprof"`
	PrometheusPath       string `mapstructure:"prometheusPath"`
	PrometheusPort       string `mapstructure:"prometheusPort"`
	CheckIntervalSeconds int    `mapstructure:"checkIntervalSeconds"`
}

func (c *Config) GetReadinessPath() string {
	return c.ReadinessPath
}

func (c *Config) GetLivenessPath() string {
	return c.LivenessPath

}

func (c *Config) GetPort() string {
	return c.Port
}

func (c *Config) GetPProf() string {
	return c.Pprof
}

func (c *Config) GetPrometheusPath() string {
	return c.PrometheusPath
}

func (c *Config) GetCheckIntervalSeconds() int {
	return c.CheckIntervalSeconds
}
