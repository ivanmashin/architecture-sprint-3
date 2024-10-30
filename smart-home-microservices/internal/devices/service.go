package devices

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/ismashin/architecture-sprint-3/smart-home-microservices/internal/devices/config"
	"github.com/ismashin/architecture-sprint-3/smart-home-microservices/internal/devices/ports/http"
	"github.com/ismashin/architecture-sprint-3/smart-home-microservices/internal/devices/usecase"
)

func RunService(cfg config.Config) {
	app := usecase.NewApp(cfg)
	server := http.NewHTTPServer(cfg, app)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()

	slog.Info("Devices service started")
	<-quit
	slog.Info("Devices service shutting down")
}
