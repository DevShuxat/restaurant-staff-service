package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/DevShuxat/restaurant-staff-service/src/infrastructure/rand"
	"github.com/DevShuxat/restaurant-staff-service/src/domain/manager/models"
	"github.com/DevShuxat/restaurant-staff-service/src/domain/manager/repositories"
	

	"github.com/DevShuxat/logistic-staff-service/src/infrastructure/crypto"
	"go.uber.org/zap"
)

type ManagerService interface {
	RegisterManager(ctx context.Context, entityID, restaurantID, email, password string) (string, error)
	ChangeManagerPassword(ctx context.Context, managerID, currentPassword, newPassword string) error
	LoginManager(ctx context.Context, email, password string) (*models.ManagerProfile, error)
	GetManagerProfile(ctx context.Context, managerID string) (*models.ManagerProfile, error)
	UpdateManagerProfile(ctx context.Context, profile *models.ManagerProfile) error
	AssignManagerToRestaurant(ctx context.Context, managerID, restaurantID string) (*models.ManagerRestaurantAssignment, error)
	RemoveManagerAssignmentFromRestaurant(ctx context.Context, assignmentID int) error
	GetManagerRestaurant(ctx context.Context, managerID string) (string, error)
}

type managerSvcImpl struct {
	managerRepo repositories.ManagerRepository,
	logger      *zap.Logger,
}

func NewManagerService(
	managerSvcImpl *repositories.ManagerRepository,
	logger *zap.Logger,
) ManagerService {
	return &managerSvcImpl{
		manageSvcImpl: managerRepo,
		logger:        logger,
	}
}

func (s *managerSvcImpl) RegisterManager(ctx context.Context, entityID, restaurantID, phoneNumber, email, password string) (string, error) {
	manager, err := s.managerRepo.GetManagerByEmail(ctx, email)
	if err == nil {
		return manager.ID, fmt.Errorf("manager with this email already exists: %s", email)
	}
	var (
		managerID   = rand.UUID()
		managerName = fmt.Sprintf("manager-%s", rand.NumericString(5))
		salt        = crypto.GenerateSalt()
		saltedPass  = crypto.Combine(salt, phoneNumber)
		passHash    = crypto.HashPassword(saltedPass)
		now         = time.Now().UTC()
	)
	manager = &models.Manager{
		ID:           managerID,
		Email:        email,
		PasswordHash: passHash,
		PasswordSalt: salt,
		CreatedAt:    now,
		UpdatedAt:    now,
	}
	profile = &models.ManagerProfile{
		ManagerID:   managerID,
		Name:        managerName,
		ImageUrl:    "",
		PhoneNumber: "",
		Email:       email,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	err = s.managerRepo.WithTx(ctx, func(r repository.ManagerRepository) error {
		if err := r.SaveManager(ctx, manager); err != nil {
			return err
		}

		if err := r.SaveManagerProfile(ctx, &profile); err != nil {
			return err
		}

		if err := r.SaveManagerRestaurantAssignment(ctx, &assignment); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return "", err
	}
	return managerID, nil
}

func (s *managerSvcImpl) ChangeManagerPassword(ctx context.Context, managerID, currentPassword, newPassword string) error {
	manager, err := s.managerRepo.GetManager(ctx, managerID)
	if err != nil {
		return err
	}

	if !crypto.PasswordMatch(currentPassword, manager.PasswordSalt, manager.PasswordHash) {
		return errors.New("password mismatch")
	}

	var (
		salt       = crypto.GenerateSalt()
		saltedPass = crypto.Combine(salt, newPassword)
		passHash   = crypto.HashPassword(saltedPass)
	)
	manager.UpdatePassword(passHash, saltedPass)
	if err := s.managerRepo.UpdateManager(ctx, manager); err != nil {
		return err
	}
	return nil
}

func (s *managerSvcImpl) LoginManager(ctx context.Context, email, password string) (*models.ManagerProfile, error) {
	manager, err := s.managerRepo.GetManagerByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	
	if !crypto.PasswordMatch(password, manager.PasswordSalt, manager.PasswordHash) {
		return nil, errors.New("password does not match")
	}

		profile, err := s.managerRepo.GetManagerProfile(ctx, manager.ID)
		if err != nil {
			return nil, err
		}
		return profile, nil
}

func (s *managerSvcImpl) GetManagerProfile(ctx context.Context, managerID string) (*models.ManagerProfile, error) {
	profile, err := s.managerRepo.GetManagerProfile(ctx, managerID)
	if err != nil {
		return nil, err
	}
	return profile, nil
}

func (s *managerSvcImpl) UpdateManagerProfile(ctx context.Context, profile *models.ManagerProfile) error {
	profile.Updated()
	if err := s.managerRepo.UpdateManagerProfile(ctx, &profile); err != nil {
		return err
	}
	return nil
}

func (s *managerSvcImpl) AssignManagerToRestaurant(ctx context.Context, managerID, restaurantID string) *models.ManagerRestaurantAssignment {
	assigment := models.ManagerRestaurantAssignment{
		ManagerID:    managerID,
		RestaurantID: restaurantID,
		CreatedAt:    time.Now(),
	}
	if err := s.managerRepo.SaveManagerRestaurantAssignment(ctx, &assigment); err != nil {
		return nil, err
	}
	return &assigment
}

// func (s *managerSvcImpl) RemoveManagerAssignmentFromRestaurant(ctx context.Context, assignmentID int) error {
// 	if err := s.managerRepo.DeleteManagerRestaurantAssignment(ctx, &assignmentID); err != nil {
// 		return err
// 	}
// 	return nil
// }

func (s *managerSvcImpl) RemoveManagerAssignmentFromRestaurant(ctx context.Context, assignmentID int) error {
	if err := s.managerRepo.RemoveManagerAssignmentFromRestaurant(ctx, assignmentID); err !=nil {
		return err
	}
	return nil
}

func (s *managerSvcImpl) GetManagerRestaurant(ctx context.Context, managerID string) (string, error) {
	restaurantID, err := s.managerRepo.GetManagerRestaurant(ctx, managerID)
	if err != nil {
		return "", err
	}
	return restaurantID, nil
}