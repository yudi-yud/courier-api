package repositories

import (
	"courier-api/config"
	"courier-api/models"
)

type TariffRepository interface {
	Create(tariff *models.Tariff) error
	FindAll() ([]models.Tariff, error)
	FindByRoute(origin, destination, serviceType string) (*models.Tariff, error)
}

type tariffRepository struct{}

func NewTariffRepository() TariffRepository {
	return &tariffRepository{}
}

func (r *tariffRepository) Create(tariff *models.Tariff) error {
	return config.DB.Create(tariff).Error
}

func (r *tariffRepository) FindAll() ([]models.Tariff, error) {
	var tariffs []models.Tariff
	if err := config.DB.Find(&tariffs).Error; err != nil {
		return nil, err
	}
	return tariffs, nil
}

func (r *tariffRepository) FindByRoute(origin, destination, serviceType string) (*models.Tariff, error) {
	var tariff models.Tariff
	err := config.DB.Where("origin_city = ? AND destination_city = ? AND service_type = ?",
		origin, destination, serviceType).First(&tariff).Error
	if err != nil {
		return nil, err
	}
	return &tariff, nil
}
