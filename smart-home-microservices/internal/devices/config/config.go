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
		PostgresDNS:                "postgres://postgres:postgres@postgresql.default.svc.cluster.local:5432/sh_devices",
		KafkaDNSs:                  []string{"kafka.default.svc.cluster.local"},
		OutboxRelayPollingInterval: 1 * time.Minute,
	}
}
