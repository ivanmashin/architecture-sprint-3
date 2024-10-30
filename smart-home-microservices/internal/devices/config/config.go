package config

import "time"

type Config struct {
	Secret                     string
	HTTPServerAddress          string
	PostgresDNS                string
	KafkaDNSs                  []string
	OutboxRelayPollingInterval time.Duration
}

func DefaultConfig() Config {
	return Config{
		Secret:                     "secret",
		HTTPServerAddress:          ":8080",
		PostgresDNS:                "postgres://postgres:postgres@localhost:5432/sh_devices",
		KafkaDNSs:                  []string{"localhost:9092"},
		OutboxRelayPollingInterval: 1 * time.Minute,
	}
}
