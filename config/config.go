package config

type Config struct {
	HTTPCfg HTTPConfig
	LoggerCfg LoggerConfig
}

type HTTPConfig struct {
	Port string
	Host string
}

type LoggerConfig struct {
	Level string
}

func NewConfig() *Config {
	return &Config{
		HTTPCfg: HTTPConfig{
			Host: "0.0.0.0",
			Port: "8080",
		},
		LoggerCfg: LoggerConfig{
			Level: "debug",
		},
	}
}