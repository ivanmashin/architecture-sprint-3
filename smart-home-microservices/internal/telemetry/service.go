package telemetry

import (
	"context"
	"log/slog"
	"os/signal"
	"syscall"

	"github.com/ismashin/architecture-sprint-3/smart-home-microservices/internal/telemetry/config"
	"github.com/ismashin/architecture-sprint-3/smart-home-microservices/internal/telemetry/ports/kafka"
	"github.com/ismashin/architecture-sprint-3/smart-home-microservices/internal/telemetry/usecase"
)

func RunService(cfg config.Config) {
	app := usecase.NewApp(cfg)
	consumer := kafka.NewConsumer(cfg, app)

	runtimeCtx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	go func() {
		err := consumer.Consume(runtimeCtx)
		if err != nil {
			panic(err)
		}
	}()

	slog.Info("Devices service started")
	<-runtimeCtx.Done()
	slog.Info("Devices service shutting down")
}
