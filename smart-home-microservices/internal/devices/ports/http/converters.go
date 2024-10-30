package http

import (
	"github.com/ismashin/architecture-sprint-3/smart-home-microservices/internal/devices/domain"
)

type Converter struct{}

func (c *Converter) HomesFromDomain(homes []*domain.Home) []Home {
	converted := make([]Home, 0, len(homes))
	for _, h := range homes {
		converted = append(converted, c.HomeFromDomain(h))
	}
	return converted
}

func (c *Converter) HomeFromDomain(home *domain.Home) Home {
	return Home{
		Id:   string(home.ID),
		Name: home.Name,
	}
}

func (c *Converter) DevicesFromDomain(devices []*domain.Device) []Device {
	result := make([]Device, 0, len(devices))
	for _, d := range devices {
		result = append(result, c.DeviceFromDomain(d))
	}
	return result
}

func (c *Converter) DeviceFromDomain(device *domain.Device) Device {
	return Device{
		Id:     string(device.ID),
		Name:   device.Name,
		On:     device.On,
		Online: device.Online,
		Type:   DeviceType(device.Type),
	}
}

func (c *Converter) DeviceToDomain(homeID domain.ID, device *Device) *domain.Device {
	return &domain.Device{
		ID:     domain.ID(device.Id),
		Type:   domain.DeviceType(device.Type),
		Name:   device.Name,
		Online: device.Online,
		On:     device.On,
		HomeID: homeID,
	}
}
