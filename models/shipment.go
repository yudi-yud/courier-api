package models

import (
	"time"

	"gorm.io/gorm"
)

type Shipment struct {
	gorm.Model
	ResiNumber string `gorm:"uniqueIndex;not null" json:"resi_number"`

	// Pengirim & Penerima
	SenderName    string `json:"sender_name"`
	SenderPhone   string `json:"sender_phone"`
	SenderAddress string `json:"sender_address"`
	SenderCity    string `json:"sender_city"`

	ReceiverName    string `json:"receiver_name"`
	ReceiverPhone   string `json:"receiver_phone"`
	ReceiverAddress string `json:"receiver_address"`
	ReceiverCity    string `json:"receiver_city"`

	// Detail Pengiriman
	Weight        float64 `json:"weight"`
	ServiceType   string  `json:"service_type"`
	PaymentStatus string  `json:"payment_status" gorm:"default:'PENDING'"`
	Price         float64 `json:"price"`
	ETD           string  `json:"etd"` // Estimasi waktu (dari Tarif)

	Status      string `gorm:"default:'PENDING'" json:"status"`
	PODImageURL string `json:"pod_image_url,omitempty"`

	// Relasi Ke Kurir (menggunakan pointer agar bisa null)
	CourierID *uint    `json:"courier_id"`
	Courier   *Courier `json:"courier,omitempty"`

	// Relasi Baru: Satu Shipment punya banyak Item
	Items []ShipmentItem `json:"items,omitempty"`

	// Relasi History
	Histories []TrackingHistory `json:"histories,omitempty"`
}

// DEFINISI STRUCT BARU: ShipmentItem
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
