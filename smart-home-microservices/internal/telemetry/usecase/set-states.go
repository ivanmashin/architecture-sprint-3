package usecase

import (
	"context"

	"github.com/ismashin/architecture-sprint-3/smart-home-microservices/internal/telemetry/domain"
)

type SetStates Command[domain.Device]
type setStatesUsecase struct {
	repo domain.TelemetryRepo
}

func (u *setStatesUsecase) Handle(ctx context.Context, device domain.Device) error {
	currentStates, ok := device.CurrentStates()
	if !ok {
		return nil
	}
	return u.repo.SaveCurrentStates(ctx, device.ID, currentStates)
}
