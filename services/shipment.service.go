package services

import (
	"courier-api/models"
	"courier-api/repositories"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"time"
)

type ShipmentService interface {
	CreateShipment(shipment *models.Shipment) error
	TrackShipment(resi string) (*models.Shipment, error)
	UpdateStatus(resi, status, location, note string, courierID uint) error
	UploadPOD(resi string, filePath string, userID uint) error
	GetAllShipments(page, limit int, search string) (map[string]interface{}, error)
	AssignCourier(resi string, courierID uint) error
	GetDashboardStats() (map[string]interface{}, error)
	GetMyTasks(userID uint) ([]models.Shipment, error)
}

type shipmentService struct {
	repo        repositories.ShipmentRepository
	courierRepo repositories.CourierRepository
	tariffRepo  repositories.TariffRepository
}

func NewShipmentService() ShipmentService {
	return &shipmentService{
		repo:        repositories.NewShipmentRepository(),
		courierRepo: repositories.NewCourierRepository(),
		tariffRepo:  repositories.NewTariffRepository(),
	}
}

func (s *shipmentService) CreateShipment(shipment *models.Shipment) error {
	rand.Seed(time.Now().UnixNano())
	shipment.ResiNumber = fmt.Sprintf("TRK%d", rand.Intn(1000000)+1000000)
	shipment.Status = "PENDING"

	tariff, err := s.tariffRepo.FindByRoute(shipment.SenderCity, shipment.ReceiverCity, shipment.ServiceType)
	if err != nil {
		return errors.New("tariff not found for this route and service type")
	}

	var totalWeight float64
	for _, item := range shipment.Items {
		totalWeight += item.Weight * float64(item.Quantity)
	}
	if totalWeight == 0 {
		totalWeight = shipment.Weight
	}

	shipment.Weight = totalWeight
	shipment.Price = totalWeight * tariff.PricePerKg
	shipment.ETD = tariff.ETD

	err = s.repo.Create(shipment)
	if err != nil {
		return err
	}

	history := &models.TrackingHistory{
		ShipmentID: shipment.ID,
		Status:     "PENDING",
		Location:   fmt.Sprintf("Warehouse %s", shipment.SenderCity),
		Note:       fmt.Sprintf("Shipment created. ETD: %s", shipment.ETD),
		Timestamp:  time.Now(),
	}
	return s.repo.AddHistory(history)
}

func (s *shipmentService) TrackShipment(resi string) (*models.Shipment, error) {
	shipment, err := s.repo.FindByResi(resi)
	if err != nil {
		return nil, errors.New("shipment not found")
	}
	return shipment, nil
}

func (s *shipmentService) GetAllShipments(page, limit int, search string) (map[string]interface{}, error) {
	shipments, total, err := s.repo.FindAll(page, limit, search)
	if err != nil {
		return nil, err
	}

	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	return map[string]interface{}{
		"shipments": shipments,
		"total":     total,
		"page":      page,
		"last_page": totalPages,
	}, nil
}

func (s *shipmentService) AssignCourier(resi string, courierID uint) error {
	_, err := s.courierRepo.FindByID(courierID)
	if err != nil {
		return errors.New("courier not found")
	}

	shipment, err := s.repo.FindByResi(resi)
	if err != nil {
		return err
	}

	shipment.CourierID = &courierID

	if err := s.repo.Update(shipment); err != nil {
		return err
	}

	history := &models.TrackingHistory{
		ShipmentID: shipment.ID,
		Status:     "COURIER_ASSIGNED",
		Location:   "Sorting Center",
		Note:       fmt.Sprintf("Assigned to Courier ID %d", courierID),
		Timestamp:  time.Now(),
	}
	return s.repo.AddHistory(history)
}

func (s *shipmentService) UpdateStatus(resi, status, location, note string, courierID uint) error {
	shipment, err := s.repo.FindByResi(resi)
	if err != nil {
		return err
	}

	shipment.Status = status

	if err := s.repo.Update(shipment); err != nil {
		return err
	}

	history := &models.TrackingHistory{
		ShipmentID: shipment.ID,
		Status:     status,
		Location:   location,
		Note:       note,
		Timestamp:  time.Now(),
	}
	return s.repo.AddHistory(history)
}

func (s *shipmentService) UploadPOD(resi string, filePath string, userID uint) error {
	shipment, err := s.repo.FindByResi(resi)
	if err != nil {
		return err
	}

	courier, err := s.courierRepo.FindByUserID(userID)
	if err != nil {
		return errors.New("you are not a registered courier")
	}

	if shipment.CourierID == nil || *shipment.CourierID != courier.ID {
		return errors.New("unauthorized: you are not assigned to this shipment")
	}

	shipment.PODImageURL = filePath
	shipment.Status = "DELIVERED"

	if err := s.repo.Update(shipment); err != nil {
		return err
	}

	history := &models.TrackingHistory{
		ShipmentID: shipment.ID,
		Status:     "DELIVERED",
		Location:   shipment.ReceiverAddress,
		Note:       "Proof of Delivery uploaded",
		Timestamp:  time.Now(),
	}
	return s.repo.AddHistory(history)
}

func (s *shipmentService) GetDashboardStats() (map[string]interface{}, error) {
	pending, _ := s.repo.CountByStatus("PENDING")
	shipped, _ := s.repo.CountByStatus("SHIPPED")
	delivered, _ := s.repo.CountByStatus("DELIVERED")

	return map[string]interface{}{
		"pending_count":   pending,
		"shipped_count":   shipped,
		"delivered_count": delivered,
	}, nil
}
func (s *shipmentService) GetMyTasks(userID uint) ([]models.Shipment, error) {
	courier, err := s.courierRepo.FindByUserID(userID)
	if err != nil {
		return nil, errors.New("courier profile not found for this user")
	}

	shipments, err := s.repo.FindByCourierID(courier.ID)
	if err != nil {
		return nil, err
	}

	return shipments, nil
}
