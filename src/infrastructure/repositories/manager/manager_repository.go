package manager

import (
	"context"

	"github.com/DevShuxat/restaurant-staff-service/src/domain/manager/models"
	repository "github.com/DevShuxat/restaurant-staff-service/src/domain/manager/repositories"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	tableManager                    = "restaurant_staff.managers"
	tableManagerProfile             = "restaurant_staff.manager_profiles"
	tableManagerRestaurantAssigment = "restaurant_staff.managers_restaurant_assigments"
)

type managerSvcImpl struct {
	db *gorm.DB
}

// WithTx implements repository.ManagerRepository.


func NewManagerRepository(db *gorm.DB) repository.ManagerRepository {
	return &managerSvcImpl{
		db: db,
	}
}

func (r *managerSvcImpl) WithTx(ctx context.Context, f func(r repository.ManagerRepository) error) error {
	if err := r.db.Transaction(func(tx *gorm.DB) error {
		r := managerSvcImpl{
			db: tx,
		}
		if err := f(&r); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

// DeleteManagerRestaurantAssignment implements repository.ManagerRepository.
func (r *managerSvcImpl) RemoveManagerRestaurantAssignment(ctx context.Context, assignmentID int) error {
	var assignment *models.ManagerRestaurantAssignment
	result := r.db.WithContext(ctx).Table(tableManagerRestaurantAssigment).Delete(assignment, "id = ?",assignmentID)
	if result.Error != nil {
		return result.Error
	}
	return nil 
}

// GetManager implements repository.ManagerRepository.
func (r *managerSvcImpl) GetManager(ctx context.Context, managerID string) (*models.Manager, error) {
	var manager *models.Manager
	result := r.db.WithContext(ctx).Table(tableManager).First(manager, "id = ?", managerID)
	if result.Error != nil {
		return nil, result.Error
	}
	return manager, nil
}

// GetManagerByEmail implements repository.ManagerRepository.
func (r *managerSvcImpl) GetManagerByEmail(ctx context.Context, email string) (*models.Manager, error) {
	var manager *models.Manager
	result := r.db.WithContext(ctx).Table(tableManager).First(manager, "email = ?", email)
	if result.Error != nil {
		return nil, result.Error
	}
	return manager, nil
}

// GetManagerProfile implements repository.ManagerRepository.
func (r *managerSvcImpl) GetManagerProfile(ctx context.Context, managerID string) (*models.ManagerProfile, error) {
	var profile *models.ManagerProfile
	result := r.db.WithContext(ctx).Table(tableManager).First(profile, "id = ?", managerID)
	if result.Error != nil {
		return nil, result.Error
	}
	return profile, nil
}

// GetManagerRestaurant implements repository.ManagerRepository.
func (r *managerSvcImpl) GetManagerRestaurant(ctx context.Context, managerID string) (string, error) {
 var assignment *models.ManagerRestaurantAssignment
 result := r.db.WithContext(ctx).Table(tableManagerRestaurantAssigment).First(assignment, "manager_id = ?", managerID)
 if result.Error != nil {
	return "", result.Error
 }
 return assignment.RestaurantID, nil
}

// SaveManager implements repository.ManagerRepository.
func (r *managerSvcImpl) SaveManager(ctx context.Context, manager *models.Manager) error {
	result := r.db.WithContext(ctx).Table(tableManager).Create(manager)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// SaveManagerProfile implements repository.ManagerRepository.
func (r *managerSvcImpl) SaveManagerProfile(ctx context.Context, profile *models.ManagerProfile) error {
	result := r.db.WithContext(ctx).Table(tableManagerProfile).Save(profile)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// SaveManagerRestaurantAssignment implements repository.ManagerRepository.
func (r *managerSvcImpl) SaveManagerRestaurantAssignment(ctx context.Context, assignment *models.ManagerRestaurantAssignment) error {
	result := r.db.WithContext(ctx).Table(tableManagerRestaurantAssigment).Clauses(clause.OnConflict{DoNothing: true}).Create(assignment)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// UpdateManager implements repository.ManagerRepository.
func (r *managerSvcImpl) UpdateManager(ctx context.Context, manager *models.Manager) error {
	result := r.db.WithContext(ctx).Table(tableManager).Save(manager)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// UpdateManagerProfile implements repository.ManagerRepository.
func (r *managerSvcImpl) UpdateManagerProfile(ctx context.Context, profile *models.ManagerProfile) error {
	result := r.db.WithContext(ctx).Table(tableManagerProfile).Save(profile)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
