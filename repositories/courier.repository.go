package repositories

import (
	"courier-api/config"
	"courier-api/models"
)

type CourierRepository interface {
	Create(courier *models.Courier) error
	FindAll() ([]models.Courier, error)
	FindByID(id uint) (*models.Courier, error)
	FindByUserID(userID uint) (*models.Courier, error)
	Update(courier *models.Courier) error
}

type courierRepository struct{}

func NewCourierRepository() CourierRepository {
	return &courierRepository{}
}

func (r *courierRepository) Create(courier *models.Courier) error {
	return config.DB.Create(courier).Error
}

func (r *courierRepository) FindAll() ([]models.Courier, error) {
	var couriers []models.Courier
	if err := config.DB.Find(&couriers).Error; err != nil {
		return nil, err
	}
	return couriers, nil
}

func (r *courierRepository) FindByID(id uint) (*models.Courier, error) {
	var courier models.Courier
	if err := config.DB.First(&courier, id).Error; err != nil {
		return nil, err
	}
	return &courier, nil
}

func (r *courierRepository) FindByUserID(userID uint) (*models.Courier, error) {
	var courier models.Courier
	if err := config.DB.Where("user_id = ?", userID).First(&courier).Error; err != nil {
		return nil, err
	}
	return &courier, nil
}

func (r *courierRepository) Update(courier *models.Courier) error {
	return config.DB.Save(courier).Error
}
