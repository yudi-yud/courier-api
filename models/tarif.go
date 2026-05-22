package models

import "gorm.io/gorm"

type Tariff struct {
	gorm.Model
	OriginCity      string  `json:"origin_city" gorm:"not null"`
	DestinationCity string  `json:"destination_city" gorm:"not null"`
	ServiceType     string  `json:"service_type" gorm:"not null"`
	PricePerKg      float64 `json:"price_per_kg" gorm:"not null"`
	ETD             string  `json:"etd"`
}
