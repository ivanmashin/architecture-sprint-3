package main

import (
	"log/slog"

	"github.com/ismashin/architecture-sprint-3/smart-home-microservices/internal/devices"
	"github.com/ismashin/architecture-sprint-3/smart-home-microservices/internal/devices/config"
)

func main() {
	cfg := config.DefaultConfig()
	slog.Info("Starting devices service...")
	devices.RunService(cfg)
}
