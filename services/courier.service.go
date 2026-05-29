package services

import (
	"courier-api/models"
	"courier-api/repositories"
)

// Definisi Interface - Pastikan GetCourierByID ada di sini
type CourierService interface {
	CreateCourier(courier *models.Courier) error
	GetAllCouriers() ([]models.Courier, error)
	GetCourierByID(id uint) (*models.Courier, error) // <--- FUNGSI INI HARUS ADA
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

// Implementasi fungsi GetCourierByID
func (s *courierService) GetCourierByID(id uint) (*models.Courier, error) {
	return s.repo.FindByID(id)
}
