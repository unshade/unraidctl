package internal

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	Api struct {
		BaseUrl       string `env:"BASE_URL"`
		ApiKey        string `env:"KEY"`
		SkipTlsVerify bool   `env:"SKIP_TLS_VERIFY"`
	} `envPrefix:"API_"`
}

func GetConfig() (*Config, error) {
	_ = godotenv.Load()
	config, err := env.ParseAs[Config]()
	return &config, err
}
