package pgrepo

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/ismashin/architecture-sprint-3/smart-home-microservices/internal/devices/config"
	"github.com/ismashin/architecture-sprint-3/smart-home-microservices/internal/devices/domain"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"

	_ "github.com/lib/pq"
)

type PostgresDeviceOutboxRepo struct {
	db        *pgxpool.Pool
	queries   *Queries
	inTx      bool
	converter Converter
}

func NewPostgresDeviceOutboxRepo(cfg config.Config) *PostgresDeviceOutboxRepo {
	db, err := pgxpool.New(context.Background(), cfg.PostgresDNS)
	if err != nil {
		panic(err)
	}

	return &PostgresDeviceOutboxRepo{
		db:      db,
		queries: New(db),
	}
}

func (p *PostgresDeviceOutboxRepo) WithinTransaction(
	ctx context.Context,
	fn func(ctx context.Context, repo domain.DeviceOutboxRepo) error,
) error {
	if p.inTx {
		return fn(ctx, p)
	}

	tx, err := p.db.Begin(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if v := recover(); v != nil {
			if rbErr := tx.Rollback(ctx); rbErr != nil {
				slog.Error("tx err", err, "rb err", rbErr)
			}
			panic(v)
		}
	}()

	wrapped := &PostgresDeviceOutboxRepo{
		db:      p.db,
		queries: p.queries.WithTx(tx),
		inTx:    true,
	}

	err = fn(ctx, wrapped)

	if err != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit(ctx)
}

func (p *PostgresDeviceOutboxRepo) GetDeviceDeletedEvents(ctx context.Context) ([]domain.DeviceEvent, error) {
	events, err := p.queries.GetOutboxMessagesDeviceDeleted(ctx)
	if err != nil {
		return nil, err
	}
	return p.converter.DeviceEventsToDomain(events), nil
}

func (p *PostgresDeviceOutboxRepo) GetDeviceCreatedEvents(ctx context.Context) ([]domain.DeviceEvent, error) {
	events, err := p.queries.GetOutboxMessagesDeviceCreated(ctx)
	if err != nil {
		return nil, err
	}
	return p.converter.DeviceEventsToDomain(events), nil
}

func (p *PostgresDeviceOutboxRepo) GetDeviceUpdatedEvents(ctx context.Context) ([]domain.DeviceEvent, error) {
	events, err := p.queries.GetOutboxMessagesDeviceUpdated(ctx)
	if err != nil {
		return nil, err
	}
	return p.converter.DeviceEventsToDomain(events), nil
}

func (p *PostgresDeviceOutboxRepo) DeleteDeviceDeletedEvents(ctx context.Context) error {
	return p.queries.DeleteOutboxMessagesDeviceDeleted(ctx)
}

func (p *PostgresDeviceOutboxRepo) DeleteDeviceCreatedEvents(ctx context.Context) error {
	return p.queries.DeleteOutboxMessagesDeviceCreated(ctx)
}

func (p *PostgresDeviceOutboxRepo) DeleteDeviceUpdatedEvents(ctx context.Context) error {
	return p.queries.DeleteOutboxMessagesDeviceUpdated(ctx)
}

func (p *PostgresDeviceOutboxRepo) SaveDeviceDeletedEvent(ctx context.Context, event domain.DeviceEvent) error {
	return p.queries.SaveDeviceDeletedToOutbox(ctx, pgtype.Text{
		String: string(event.DeviceID),
		Valid:  true,
	})
}

func (p *PostgresDeviceOutboxRepo) SaveDeviceCreatedEvent(ctx context.Context, event domain.DeviceEvent) error {
	return p.queries.SaveDeviceCreatedToOutbox(ctx, pgtype.Text{
		String: string(event.DeviceID),
		Valid:  true,
	})
}

func (p *PostgresDeviceOutboxRepo) SaveDeviceUpdatedEvent(ctx context.Context, event domain.DeviceEvent) error {
	return p.queries.SaveDeviceUpdatedToOutbox(ctx, pgtype.Text{
		String: string(event.DeviceID),
		Valid:  true,
	})
}

func (p *PostgresDeviceOutboxRepo) GetUserHomes(ctx context.Context, userID domain.ID) ([]*domain.Home, error) {
	homes, err := p.queries.ListUserHomes(ctx, string(userID))
	if err != nil {
		return nil, err
	}

	return p.converter.HomesToDomain(homes), nil
}

func (p *PostgresDeviceOutboxRepo) GetDevicesInHome(
	ctx context.Context,
	userID, homeID domain.ID,
) ([]*domain.Device, error) {
	devices, err := p.queries.ListHomeDevices(ctx, ListHomeDevicesParams{
		HomeID: pgtype.Text{
			String: string(homeID),
			Valid:  true,
		},
		UserID: string(userID),
	})
	if err != nil {
		return nil, err
	}

	return p.converter.DevicesToDomain(devices), nil
}

func (p *PostgresDeviceOutboxRepo) CreateDevice(ctx context.Context, userID domain.ID, device *domain.Device) error {
	return p.queries.CreateDevice(ctx, CreateDeviceParams{
		ID:     string(device.ID),
		Type:   string(device.Type),
		Name:   device.Name,
		Online: device.Online,
		OnOff:  device.On,
		UserID: string(userID),
		HomeID: pgtype.Text{
			String: string(device.HomeID),
			Valid:  true,
		},
	})
}

func (p *PostgresDeviceOutboxRepo) GetDeviceByID(ctx context.Context, userID, id domain.ID) (*domain.Device, error) {
	device, err := p.queries.GetDeviceByID(ctx, GetDeviceByIDParams{
		ID:     string(userID),
		UserID: string(id),
	})
	if err != nil {
		return nil, err
	}

	return p.converter.DeviceToDomain(device), nil
}

func (p *PostgresDeviceOutboxRepo) UpdateDevice(
	ctx context.Context,
	userID, deviceID domain.ID,
	update domain.UpdateDeviceFunc,
) error {
	device, err := p.queries.GetDeviceByIDForUpdate(ctx, GetDeviceByIDForUpdateParams{
		ID:     string(userID),
		UserID: string(deviceID),
	})
	if err != nil {
		return err
	}

	domainDevice := domain.Device{
		ID:     deviceID,
		Type:   domain.DeviceType(device.Type),
		Name:   device.Name,
		Online: device.Online,
		On:     device.OnOff,
		HomeID: domain.ID(device.HomeID.String),
	}

	err = update(&domainDevice)
	if err != nil {
		return err
	}

	return p.queries.UpdateDevice(ctx, UpdateDeviceParams{
		Name:   domainDevice.Name,
		Online: domainDevice.Online,
		OnOff:  domainDevice.On,
		ID:     string(domainDevice.ID),
		UserID: string(userID),
	})
}

func (p *PostgresDeviceOutboxRepo) DeleteDeviceByID(ctx context.Context, userID, id domain.ID) error {
	return p.queries.DeleteDeviceByID(ctx, DeleteDeviceByIDParams{
		ID:     string(id),
		UserID: string(userID),
	})
}
