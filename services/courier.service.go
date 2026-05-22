package services

import (
	"courier-api/models"
	"courier-api/repositories"
)

type CourierService interface {
	CreateCourier(courier *models.Courier) error
	GetAllCouriers() ([]models.Courier, error)
}

type courierService struct {
	repo repositories.CourierRepository
}

func NewCourierService() CourierService {
	return &courierService{repo: repositories.NewCourierRepository()}
}

func (s *courierService) CreateCourier(courier *models.Courier) error {
	return s.repo.Create(courier)
}

func (s *courierService) GetAllCouriers() ([]models.Courier, error) {
	return s.repo.FindAll()
}
