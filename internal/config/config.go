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
	
}

func NewConfig() *Config {
	return nil
}