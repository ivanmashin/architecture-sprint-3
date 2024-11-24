package usecase

import (
	"context"

	"github.com/ismashin/architecture-sprint-3/smart-home-microservices/internal/devices/adapters/outbox-relay"
	"github.com/ismashin/architecture-sprint-3/smart-home-microservices/internal/devices/adapters/pgrepo"
	"github.com/ismashin/architecture-sprint-3/smart-home-microservices/internal/devices/config"
	"github.com/ismashin/architecture-sprint-3/smart-home-microservices/internal/devices/domain"
)

// App описывает все варианты использования приложения
type App struct {
	// GetHomes возвращает список домов
	GetHomes GetHomes
	// GetDevicesInHome возвращает список устройств в доме
	GetDevicesInHome GetDevicesInHome
	// GetDeviceByID возвращает устройство пользователя
	GetDeviceByID GetDeviceByID

	// CreateDevice создает новое устройство
	CreateDevice CreateDevice
	// ToggleDevice переключает состояние устройства (вкл/выкл)
	ToggleDevice ToggleDevice
	// UpdateDevice обновляет устройство
	UpdateDevice UpdateDevice
	// DeleteDevice удаляет устройство
	DeleteDevice DeleteDevice
}

type ToggleCommand struct {
	DeviceID domain.ID
	On       bool
}

type UpdateDeviceCommand struct {
	DeviceID   domain.ID
	DeviceName *string
	On         *bool
}

type Command[I any] interface {
	Handle(ctx context.Context, userID domain.ID, in I) error
}

type Query[I, O any] interface {
	Handle(ctx context.Context, userID domain.ID, in I) (O, error)
}

func NewApp(cfg config.Config) *App {
	repository := pgrepo.NewPostgresDeviceOutboxRepo(cfg)
	outboxRelay := outbox.NewKafkaPollingRelay(cfg, repository)
	return &App{
		GetHomes:         &getHomesUsecase{repo: repository},
		GetDevicesInHome: &getDevicesInHomeUsecase{repo: repository},
		GetDeviceByID:    &getDeviceByIDUsecase{repo: repository},
		CreateDevice:     &createDeviceUsecase{repo: repository, outboxRelay: outboxRelay},
		ToggleDevice:     &toggleDeviceUsecase{repo: repository, outboxRelay: outboxRelay},
		UpdateDevice:     &updateDeviceUsecase{repo: repository, outboxRelay: outboxRelay},
		DeleteDevice:     &deleteDeviceUsecase{repo: repository, outboxRelay: outboxRelay},
	}
}
