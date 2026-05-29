// Buat file baru: constants/status.go
package constants

const (
	// Shipment Status
	StatusPending   = "PENDING"
	StatusAssigned  = "COURIER_ASSIGNED"
	StatusInTransit = "IN_TRANSIT"
	StatusDelivered = "DELIVERED"
	StatusCancelled = "CANCELLED"

	// Courier Status
	CourierOffline = "OFFLINE"
	CourierOnline  = "ONLINE"
	CourierBusy    = "BUSY"
)
