package config

import "github.com/kelseyhightower/envconfig"

// Config представляет конфигурацию сервиса
type Config struct {
	Environment   string `envconfig:"ENVIRONMENT" required:"true"`
	MigrationsDIR  string `envconfig:"MIGRATIONS_DIR"`
	APPPort       int64  `envconfig:"APP_PORT" required:"true"`
	DatabaseDSN   string `envconfig:"DATABASE_DSN" required:"true"`
	TemplateINDEX string `envconfig:"TEMPLATE_INDEX" required:"true"`
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
