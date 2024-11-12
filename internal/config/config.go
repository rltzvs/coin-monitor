package config

type Config struct {
	RedisAddress   string
	UpdateInterval int
}

func LoadConfig() *Config {
	config := &Config{}

	return config
}
