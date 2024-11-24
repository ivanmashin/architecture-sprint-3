package pgrepo

import (
	"context"

	"github.com/ismashin/architecture-sprint-3/smart-home-microservices/internal/telemetry/config"
	"github.com/ismashin/architecture-sprint-3/smart-home-microservices/internal/telemetry/domain"
	"github.com/jackc/pgx/v5/pgxpool"

	_ "github.com/lib/pq"
)

type PostgresTelemetryRepo struct {
	db        *pgxpool.Pool
	queries   *Queries
	converter Converter
}

func NewPostgresDeviceOutboxRepo(cfg config.Config) *PostgresTelemetryRepo {
	db, err := pgxpool.New(context.Background(), cfg.PostgresDNS)
	if err != nil {
		panic(err)
	}

	return &PostgresTelemetryRepo{
		db:      db,
		queries: New(db),
	}
}

func (p *PostgresTelemetryRepo) GetDeviceTelemetry(ctx context.Context, deviceID domain.ID) (*domain.Device, error) {
	telemetryData, err := p.queries.GetDeviceTelemetry(ctx, string(deviceID))
	if err != nil {
		return nil, err
	}
	return p.converter.TelemetryDataToDomain(telemetryData), nil
}

func (p *PostgresTelemetryRepo) SaveCurrentStates(
	ctx context.Context,
	deviceID domain.ID,
	states []domain.State,
) error {
	return p.queries.InsertDeviceState(ctx, p.converter.DomainToInsertDeviceStateParams(deviceID, states))
}

func (p *PostgresTelemetryRepo) DeleteDeviceStates(ctx context.Context, deviceID domain.ID) error {
	return p.queries.DeleteDeviceStates(ctx, string(deviceID))
}
