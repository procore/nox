package gaia

// Config holds configuration for a connection to elasticsearch
type Config struct {
	User struct {
		Name     string
		Password string
	}
	Net struct {
		Host string
		Port string
		TLS  struct {
			Enable bool
		}
	}
	Pretty bool
	Debug  bool
}

// NewConfig provides a factory for initalizing a config
func NewConfig() *Config {
	c := &Config{}

	c.Net.Host = "localhost"
	c.Net.Port = "9200"
	c.Net.TLS.Enable = false
	c.Pretty = false
	c.Debug = false

	return c
}
