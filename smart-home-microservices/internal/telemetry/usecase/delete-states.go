package usecase

import (
	"context"

	"github.com/ismashin/architecture-sprint-3/smart-home-microservices/internal/telemetry/domain"
)

type DeleteStates Command[domain.ID]
type deleteStatesUsecase struct {
	repo domain.TelemetryRepo
}

func (u *deleteStatesUsecase) Handle(ctx context.Context, deviceID domain.ID) error {
	// TODO: not implemented
	panic("implement me")
}
