package domain

import "context"

type UpdateDeviceFunc func(device *Device) error

type DeviceRepo interface {
	GetUserHomes(ctx context.Context, userID ID) ([]*Home, error)
	GetDevicesInHome(ctx context.Context, userID, homeID ID) ([]*Device, error)
	CreateDevice(ctx context.Context, userID ID, device *Device) error
	GetDeviceByID(ctx context.Context, userID, id ID) (*Device, error)
	UpdateDevice(ctx context.Context, userID, deviceID ID, update UpdateDeviceFunc) error
	DeleteDeviceByID(ctx context.Context, userID, id ID) error
}
