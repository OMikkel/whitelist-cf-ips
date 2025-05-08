package config

import (
	"log"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	HETZNER_API_TOKEN string `env:"HETZNER_API_TOKEN,required"`
	HETZNER_FIREWALL_ID string `env:"HETZNER_FIREWALL_ID,required"`
}

func GetConfig() *Config {
	cfg := Config{}

	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("Failed to parse environment variables: %+v", err)
	}

	return &cfg
}