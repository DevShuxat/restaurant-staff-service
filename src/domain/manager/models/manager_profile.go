package models

import "time"

type ManagerProfile struct {
	ManagerID   string    `json:"manager_id" gorm:"primaryKey"`
	EntityID    string    `json:"entity_id" gorm:"index:idx_manager_by_entity"`
	Name        string    `json:"name"`
	ImageUrl    string    `json:"image_url"`
	PhoneNumber string    `json:"phone_number"`
	Email       string    `json:"email"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
