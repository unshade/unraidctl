package internal

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	Api struct {
		BaseUrl       string `env:"BASE_URL,required"`
		ApiKey        string `env:"KEY,required"`
		SkipTlsVerify bool   `env:"SKIP_TLS_VERIFY" envDefault:"true"`
	} `envPrefix:"API_"`
}

func GetConfig() (*Config, error) {
	_ = godotenv.Load()
	config, err := env.ParseAs[Config]()
	return &config, err
}
