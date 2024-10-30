package main

import (
	"log/slog"

	"github.com/ismashin/architecture-sprint-3/smart-home-microservices/internal/telemetry"
	"github.com/ismashin/architecture-sprint-3/smart-home-microservices/internal/telemetry/config"
)

func main() {
	cfg := config.DefaultConfig()
	slog.Info("Starting telemetry service...")
	telemetry.RunService(cfg)
}
