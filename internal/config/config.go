package config

import "github.com/kelseyhightower/envconfig"

// Config представляет конфигурацию сервиса
type Config struct {
	DatabaseDSN   string `envconfig:"DATABASEDSN"`
	TemplateINDEX string `envconfig:"TEMPLATEINDEX"`
}

// InitConfig возвращает конфиг
func InitConfig() (*Config, error) {
	var cfg Config
	err := envconfig.Process("RASP", &cfg)

	return &cfg, err
}

// MustInitConfig возвращает конфиг или паникует при ошибке
func MustInitConfig() *Config {
	cfg, err := InitConfig()
	if err != nil {
		panic(err)
	}

	return cfg
}
