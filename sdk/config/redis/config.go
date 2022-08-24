package redis

type Config struct {
	Url string `json:"url,omitempty"`
}

func (c *Config) GetUrl() string {
	return c.Url
}
