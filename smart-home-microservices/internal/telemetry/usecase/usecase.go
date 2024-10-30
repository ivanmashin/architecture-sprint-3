package usecase

import (
	"context"

	"github.com/ismashin/architecture-sprint-3/smart-home-microservices/internal/telemetry/adapters/pgrepo"
	"github.com/ismashin/architecture-sprint-3/smart-home-microservices/internal/telemetry/config"
)

// App описывает все варианты использования приложения
type App struct {
	// GetStates возвращает состояния устройств
	GetStates GetStates
	// SetStates устанавливает состояния устройств
	SetStates SetStates
	// DeleteStates удаляет состояния устройств
	DeleteStates DeleteStates
}

type Command[I any] interface {
	Handle(ctx context.Context, in I) error
}

type Query[I, O any] interface {
	Handle(ctx context.Context, in I) (O, error)
}

func NewApp(cfg config.Config) *App {
	repo := pgrepo.NewPostgresDeviceOutboxRepo(cfg)
	return &App{
		GetStates: &getStatesUsecase{repo: repo},
		SetStates: &setStatesUsecase{repo: repo},
	}
}
