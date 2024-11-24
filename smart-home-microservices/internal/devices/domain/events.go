package domain

import (
	"context"

	"github.com/ismashin/architecture-sprint-3/smart-home-microservices/pkg/transactions"
)

type DeviceEvent struct {
	DeviceID ID
}

type DeviceOutboxRepo interface {
	transactions.Transactor[DeviceOutboxRepo]
	EventsOutbox
	DeviceRepo
}

type EventsOutbox interface {
	SaveDeviceDeletedEvent(ctx context.Context, event DeviceEvent) error
	SaveDeviceCreatedEvent(ctx context.Context, event DeviceEvent) error
	SaveDeviceUpdatedEvent(ctx context.Context, event DeviceEvent) error

	GetDeviceDeletedEvents(ctx context.Context) ([]DeviceEvent, error)
	GetDeviceCreatedEvents(ctx context.Context) ([]DeviceEvent, error)
	GetDeviceUpdatedEvents(ctx context.Context) ([]DeviceEvent, error)

	DeleteDeviceDeletedEvents(ctx context.Context) error
	DeleteDeviceCreatedEvents(ctx context.Context) error
	DeleteDeviceUpdatedEvents(ctx context.Context) error
}

type OutboxRelay interface {
	TriggerReadDeviceDeleted(ctx context.Context)
	TriggerReadDeviceCreated(ctx context.Context)
	TriggerReadDeviceUpdated(ctx context.Context)
}
