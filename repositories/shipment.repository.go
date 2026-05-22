package repositories

import (
	"courier-api/config"
	"courier-api/models"
)

type ShipmentRepository interface {
	Create(shipment *models.Shipment) error
	FindByResi(resi string) (*models.Shipment, error)

	FindAll(page, limit int, search string) ([]models.Shipment, int64, error)

	Update(shipment *models.Shipment) error
	AddHistory(history *models.TrackingHistory) error

	CountByStatus(status string) (int64, error)
}

type shipmentRepository struct{}

func NewShipmentRepository() ShipmentRepository {
	return &shipmentRepository{}
}

func (r *shipmentRepository) Create(shipment *models.Shipment) error {
	return config.DB.Create(shipment).Error
}

func (r *shipmentRepository) FindByResi(resi string) (*models.Shipment, error) {
	var shipment models.Shipment
	if err := config.DB.Preload("Courier").Preload("Histories").Where("resi_number = ?", resi).First(&shipment).Error; err != nil {
		return nil, err
	}
	return &shipment, nil
}

func (r *shipmentRepository) FindAll(page, limit int, search string) ([]models.Shipment, int64, error) {
	var shipments []models.Shipment
	var total int64

	offset := (page - 1) * limit
	query := config.DB.Model(&models.Shipment{})

	if search != "" {
		query = query.Where("resi_number LIKE ? OR receiver_name LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	query.Count(&total)

	result := query.Preload("Courier").Offset(offset).Limit(limit).Find(&shipments)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	return shipments, total, nil
}

func (r *shipmentRepository) Update(shipment *models.Shipment) error {
	return config.DB.Save(shipment).Error
}

func (r *shipmentRepository) AddHistory(history *models.TrackingHistory) error {
	return config.DB.Create(history).Error
}

func (r *shipmentRepository) CountByStatus(status string) (int64, error) {
	var count int64
	if err := config.DB.Model(&models.Shipment{}).Where("status = ?", status).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
