package internal

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	Api struct {
		BaseUrl       string `env:"BASE_URL,required" json:"base_url"`
		ApiKey        string `env:"KEY,required" json:"api_key"`
		SkipTlsVerify bool   `env:"SKIP_TLS_VERIFY" envDefault:"true" json:"skip_tls_verify"`
	} `envPrefix:"API_" json:"api"`
}

func GetConfig() (*Config, error) {
	_ = godotenv.Load()

	cfg := &Config{}

	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	jsonPath := filepath.Join(home, ".config", "unraidctl", "config.json")
	if b, err := os.ReadFile(jsonPath); err == nil {
		err = json.Unmarshal(b, cfg)
		if err != nil {
			return nil, err
		}

		return cfg, nil
	}

	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
