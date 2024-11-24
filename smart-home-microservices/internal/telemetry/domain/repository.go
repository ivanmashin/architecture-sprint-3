package domain

import "context"

type TelemetryRepo interface {
	GetDeviceTelemetry(ctx context.Context, deviceID ID) (*Device, error)
	SaveCurrentStates(ctx context.Context, deviceID ID, states []State) error
	DeleteDeviceStates(ctx context.Context, deviceID ID) error
}
