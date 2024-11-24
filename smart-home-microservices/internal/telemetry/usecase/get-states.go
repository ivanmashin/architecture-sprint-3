package usecase

import (
	"context"

	"github.com/ismashin/architecture-sprint-3/smart-home-microservices/internal/telemetry/domain"
)

type GetStates Query[domain.ID, []*domain.Device]
type getStatesUsecase struct {
	repo domain.TelemetryRepo
}

func (u *getStatesUsecase) Handle(ctx context.Context, deviceID domain.ID) ([]*domain.Device, error) {
	// TODO: not implemented
	panic("implement me")
}
