package models

import (
	"time"

	"gorm.io/gorm"
)

type Shipment struct {
	gorm.Model
	ResiNumber string `gorm:"uniqueIndex;not null" json:"resi_number"`

	SenderName    string `json:"sender_name"`
	SenderPhone   string `json:"sender_phone"`
	SenderAddress string `json:"sender_address"`
	SenderCity    string `json:"sender_city"`

	ReceiverName    string `json:"receiver_name"`
	ReceiverPhone   string `json:"receiver_phone"`
	ReceiverAddress string `json:"receiver_address"`
	ReceiverCity    string `json:"receiver_city"`

	Weight        float64 `json:"weight"`
	ServiceType   string  `json:"service_type"`
	PaymentStatus string  `json:"payment_status" gorm:"default:'PENDING'"`
	Price         float64 `json:"price"`
	ETD           string  `json:"etd"`

	Status      string `gorm:"default:'PENDING'" json:"status"`
	PODImageURL string `json:"pod_image_url,omitempty"`

	CourierID *uint    `json:"courier_id"`
	Courier   *Courier `json:"courier,omitempty"`

	Items []ShipmentItem `json:"items,omitempty"`

	Histories []TrackingHistory `json:"histories,omitempty"`
}

type ShipmentItem struct {
	gorm.Model
	ShipmentID uint    `json:"shipment_id"`
	ItemName   string  `json:"item_name"`
	Quantity   int     `json:"quantity"`
	Weight     float64 `json:"weight"`
}

type TrackingHistory struct {
	gorm.Model
	ShipmentID uint      `json:"shipment_id"`
	Status     string    `json:"status"`
	Location   string    `json:"location"`
	Note       string    `json:"note"`
	Timestamp  time.Time `json:"timestamp"`
}
