package pgrepo

import (
	"github.com/ismashin/architecture-sprint-3/smart-home-microservices/internal/devices/domain"
	"github.com/jackc/pgx/v5/pgtype"
)

type Converter struct{}

func (c Converter) DeviceToDomain(device Device) *domain.Device {
	return &domain.Device{
		ID:     domain.ID(device.ID),
		Type:   domain.DeviceType(device.Type),
		Name:   device.Name,
		On:     device.OnOff,
		Online: device.Online,
		HomeID: domain.ID(device.HomeID.String),
	}
}

func (c Converter) DevicesToDomain(devices []Device) []*domain.Device {
	domainDevices := make([]*domain.Device, 0, len(devices))
	for _, device := range devices {
		domainDevices = append(domainDevices, c.DeviceToDomain(device))
	}
	return domainDevices
}

func (c Converter) HomeToDomain(home Home) *domain.Home {
	return &domain.Home{
		ID:   domain.ID(home.ID),
		Name: home.Name,
	}
}

func (c Converter) HomesToDomain(homes []Home) []*domain.Home {
	domainHomes := make([]*domain.Home, 0, len(homes))
	for _, home := range homes {
		domainHomes = append(domainHomes, c.HomeToDomain(home))
	}
	return domainHomes
}

func (c Converter) DeviceEventsToDomain(events []pgtype.Text) []domain.DeviceEvent {
	domainEvents := make([]domain.DeviceEvent, 0, len(events))
	for _, event := range events {
		domainEvents = append(domainEvents, domain.DeviceEvent{
			DeviceID: domain.ID(event.String),
		})
	}
	return domainEvents
}
