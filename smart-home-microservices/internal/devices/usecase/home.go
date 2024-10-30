package usecase

import (
	"context"

	"github.com/ismashin/architecture-sprint-3/smart-home-microservices/internal/devices/domain"
)

type GetHomes Query[struct{}, []*domain.Home]
type getHomesUsecase struct {
	repo domain.DeviceRepo
}

func (u *getHomesUsecase) Handle(ctx context.Context, userID domain.ID, _ struct{}) ([]*domain.Home, error) {
	return u.repo.GetUserHomes(ctx, userID)
}

type GetDevicesInHome Query[domain.ID, []*domain.Device]
type getDevicesInHomeUsecase struct {
	repo domain.DeviceRepo
}

func (u *getDevicesInHomeUsecase) Handle(ctx context.Context, userID, homeID domain.ID) ([]*domain.Device, error) {
	return u.repo.GetDevicesInHome(ctx, userID, homeID)
}
