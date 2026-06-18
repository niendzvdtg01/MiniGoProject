package config

type Config struct {
	ServeAddress string
}

func NewConfig() *Config {
	return &Config{
		ServeAddress: ":8085",
	}
}
