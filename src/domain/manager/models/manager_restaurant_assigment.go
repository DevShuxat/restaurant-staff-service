package models
import "time"

type ManagerRestaurantAssignment struct {
	ID           int       `json:"id" gorm:"primaryKey"`
	ManagerID    string    `json:"manager_id" gorm:"index:idx_restaurant_by_manager"`
	RestaurantID string    `json:"restaurant_id" gorm:"index:idx_manager_by_restaurant"`
	CreatedAt    time.Time `json:"created_at"`
}
