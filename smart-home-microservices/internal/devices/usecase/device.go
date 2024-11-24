package usecase

import (
	"context"

	"github.com/ismashin/architecture-sprint-3/smart-home-microservices/internal/devices/domain"
)

type GetDeviceByID Query[domain.ID, *domain.Device]
type getDeviceByIDUsecase struct {
	repo domain.DeviceRepo
}

func (u *getDeviceByIDUsecase) Handle(ctx context.Context, userID, deviceID domain.ID) (*domain.Device, error) {
	return u.repo.GetDeviceByID(ctx, userID, deviceID)
}

type CreateDevice Command[*domain.Device]
type createDeviceUsecase struct {
	repo        domain.DeviceOutboxRepo
	outboxRelay domain.OutboxRelay
}

func (u *createDeviceUsecase) Handle(ctx context.Context, userID domain.ID, device *domain.Device) error {
	err := u.repo.WithinTransaction(ctx, func(ctx context.Context, repo domain.DeviceOutboxRepo) error {
		err := u.repo.CreateDevice(ctx, userID, device)
		if err != nil {
			return err
		}

		err = u.repo.SaveDeviceCreatedEvent(ctx, domain.DeviceEvent{DeviceID: device.ID})
		if err != nil {
			return err
		}

		u.outboxRelay.TriggerReadDeviceCreated(ctx)
		return nil
	})
	return err
}

type ToggleDevice Command[ToggleCommand]
type toggleDeviceUsecase struct {
	repo        domain.DeviceOutboxRepo
	outboxRelay domain.OutboxRelay
}

func (u *toggleDeviceUsecase) Handle(ctx context.Context, userID domain.ID, cmd ToggleCommand) error {
	err := u.repo.WithinTransaction(ctx, func(ctx context.Context, repo domain.DeviceOutboxRepo) error {
		err := repo.UpdateDevice(ctx, userID, cmd.DeviceID, func(device *domain.Device) error {
			device.Toggle(cmd.On)
			return nil
		})
		if err != nil {
			return err
		}

		err = repo.SaveDeviceUpdatedEvent(ctx, domain.DeviceEvent{DeviceID: cmd.DeviceID})
		if err != nil {
			return err
		}

		u.outboxRelay.TriggerReadDeviceUpdated(ctx)
		return nil
	})
	return err
}

type UpdateDevice Command[UpdateDeviceCommand]
type updateDeviceUsecase struct {
	repo        domain.DeviceOutboxRepo
	outboxRelay domain.OutboxRelay
}

func (u *updateDeviceUsecase) Handle(ctx context.Context, userID domain.ID, cmd UpdateDeviceCommand) error {
	err := u.repo.WithinTransaction(ctx, func(ctx context.Context, repo domain.DeviceOutboxRepo) error {
		err := repo.UpdateDevice(ctx, userID, cmd.DeviceID, func(device *domain.Device) error {
			device.UpdateName(cmd.DeviceName)
			device.UpdateOn(cmd.On)
			return nil
		})
		if err != nil {
			return err
		}

		err = repo.SaveDeviceUpdatedEvent(ctx, domain.DeviceEvent{DeviceID: cmd.DeviceID})
		if err != nil {
			return err
		}

		u.outboxRelay.TriggerReadDeviceUpdated(ctx)
		return nil
	})
	return err
}

type DeleteDevice Command[domain.ID]
type deleteDeviceUsecase struct {
	repo        domain.DeviceOutboxRepo
	outboxRelay domain.OutboxRelay
}

func (u *deleteDeviceUsecase) Handle(ctx context.Context, userID, deviceID domain.ID) error {
	err := u.repo.WithinTransaction(ctx, func(ctx context.Context, repo domain.DeviceOutboxRepo) error {
		err := repo.DeleteDeviceByID(ctx, userID, deviceID)
		if err != nil {
			return err
		}

		err = repo.SaveDeviceDeletedEvent(ctx, domain.DeviceEvent{DeviceID: deviceID})
		if err != nil {
			return err
		}

		u.outboxRelay.TriggerReadDeviceDeleted(ctx)
		return nil
	})
	return err
}
