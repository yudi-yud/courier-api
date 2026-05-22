package models

import (
	"gorm.io/gorm"
)

type Courier struct {
	gorm.Model
	Name         string `json:"name" gorm:"not null"`
	Phone        string `json:"phone" gorm:"not null"`
	VehiclePlate string `json:"vehicle_plate"`
	VehicleType  string `json:"vehicle_type"`
	Status       string `json:"status" gorm:"default:'OFFLINE'"`
	UserID       uint   `json:"user_id"`
	User         User   `json:"user,omitempty"`
}
