package service

import (
	"context"

	dtos "github.com/DevShuxat/restaurant-staff-service/src/application/dtos"
	"github.com/DevShuxat/restaurant-staff-service/src/domain/manager/models"
	"github.com/DevShuxat/restaurant-staff-service/src/infrastructure/jwt"
)
type ManagerApplicationService interface {
	RegisterManager(ctx context.Context, etityID string,req *dtos.RegisterManagerRequest) (*dtos.RegisterManagerResponse, error)
	ChangeManagerPassword(ctx context.Context, managerID string, req *dtos.ChangeManagerPasswordRequest) (*dtos.ChangeManagerResponse, error)
	LoginManager(ctx context.Context, req *dtos.LoginManagerRequest) (*dtos.LoginManagerResponse, error)
	GetManagerProfile(ctx context.Context, managerID string) (*dtos.GetManagerProfileResponse, error)
	UpdateManagerProfile(ctx context.Context, profiel *models.ManagerProfile) (*dtos.UpdateManagerProfileResponse, error)
}

type managerSvcImpl struct {
	managerSvc ManagerService.ManagerService
	tokenSvc jwt.Service
}

func NewManagerApplicationService(
	managerSvc managerSvc.ManagerService,
	tokenSvc jwt.Service,
) ManagerApplicationService {
	return &managerSvcImpl{
		managerSvc: managerSvc,
		tokenSvc: tokenSvc,
	}
}

func (s *managerSvcImpl) RegisterManager(ctx context.Context, entityID string, req *dtos.RegisterManagerRequest) (*dtos.RegisterManagerResponse, error) {
	return nil, nil
}

func (s *managerSvcImpl) ChangeManagerPassword(ctx context.Context, managerID string, req *dtos.ChangeManagerPasswordRequest) (*dtos.ChangeManagerPasswordResponse, error) {
	return nil, nil
}

func (s *managerSvcImpl) LoginManager(ctx context.Context, req *dtos.LoginManagerRequest) (*dtos.LoginManagerResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
} 
profile, err := s.managerSvc.LoginManager(ctx, req.Email, req.Password)
if err != nil {
	return nil, err
}
restaurantID, err := s.managerSvc.GetManagerRestaurant(ctx, profile, ManagerID)
if err != nil {
	return nil, err
}
tokenData, err := s.managerSvc.CreateToken(ctx, profile.ManagerID, tokenData)
if err != nil {
	return nil, err
}
tokenData := map[string]string{
	"restaurant_id": restaurantID,
}

token, err := s.tokenSvc.CreateToken(ctx, profile.ManagerID, tokenData)
if err != nil {
	return nil, err
}
return dtos.NewLoginManagerResponse(token, profile,restaurantID), nil
}

func (s *managerSvcImpl) GetManagerProfile(ctx context.Context, managerID string) (*dtos.GetManagerProfileResponse, error) { 
	return nil, nil
}
func (s *managerSvcImpl) UpdatedManagerProfile(ctx context.Context, profile *models.ManagerProfile) (*dtos.UpdatedManagerProfileResponse, error) {
	return nil, nil
}

