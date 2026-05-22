package controllers

import (
	"courier-api/models"
	"courier-api/services"
	"courier-api/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ShipmentController struct {
	service    services.ShipmentService
	pdfService services.PDFService
}

func NewShipmentController() *ShipmentController {
	return &ShipmentController{
		service:    services.NewShipmentService(),
		pdfService: services.NewPDFService(),
	}
}

// Create Shipment
// @Summary Create new shipment
// @Description Membuat data pengiriman baru dengan perhitungan harga otomatis.
// @Tags Shipment
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body models.Shipment true "Shipment Data"
// @Success 201 {object} map[string]interface{} "contoh: {\"status\": 201, \"message\": \"Shipment created\", \"data\": {...}}"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /shipments [post]
func (c *ShipmentController) Create(ctx *fiber.Ctx) error {
	shipment := new(models.Shipment)
	if err := ctx.BodyParser(shipment); err != nil {
		return utils.ResponseJSON(ctx, fiber.StatusBadRequest, "Cannot parse JSON", nil)
	}

	if err := c.service.CreateShipment(shipment); err != nil {
		return utils.ResponseJSON(ctx, fiber.StatusInternalServerError, err.Error(), nil)
	}

	responseData := map[string]interface{}{
		"resi_number":  shipment.ResiNumber,
		"total_weight": shipment.Weight,
		"price":        shipment.Price,
		"service_type": shipment.ServiceType,
		"etd":          shipment.ETD,
		"status":       shipment.Status,
	}

	return utils.ResponseJSON(ctx, fiber.StatusCreated, "Shipment created successfully", responseData)
}

// Track Shipment
// @Summary Track shipment by Resi
// @Description Melacak status pengiriman dan riwayat perjalanan publik (tanpa login).
// @Tags Public
// @Produce json
// @Param resi path string true "Nomor Resi"
// @Success 200 {object} map[string]interface{} "Success"
// @Failure 404 {object} map[string]interface{} "Resi not found"
// @Router /track/{resi} [get]
func (c *ShipmentController) Track(ctx *fiber.Ctx) error {
	resi := ctx.Params("resi")
	shipment, err := c.service.TrackShipment(resi)
	if err != nil {
		return utils.ResponseJSON(ctx, fiber.StatusNotFound, "Resi not found", nil)
	}

	return utils.ResponseJSON(ctx, fiber.StatusOK, "Success", shipment)
}

// GetAllShipments godoc
// @Summary Get all shipments with pagination
// @Description Mengambil semua data pengiriman dengan fitur pencarian dan paginasi.
// @Tags Shipment
// @Produce json
// @Security BearerAuth
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Limit per page" default(10)
// @Param search query string false "Search by Resi or Receiver Name"
// @Success 200 {object} map[string]interface{} "Success"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /shipments [get]
func (c *ShipmentController) GetAll(ctx *fiber.Ctx) error {
	page, _ := strconv.Atoi(ctx.Query("page", "1"))
	limit, _ := strconv.Atoi(ctx.Query("limit", "10"))
	search := ctx.Query("search", "")

	result, err := c.service.GetAllShipments(page, limit, search)
	if err != nil {
		return utils.ResponseJSON(ctx, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return utils.ResponseJSON(ctx, fiber.StatusOK, "Success", result)
}

// AssignCourier godoc
// @Summary Assign Courier to Shipment
// @Description Menugaskan kurir tertentu ke sebuah pengiriman (Admin Only).
// @Tags Shipment
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param resi path string true "Nomor Resi"
// @Param request body map[string]uint true "Courier ID"
// @Success 200 {object} map[string]interface{} "Success"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /shipments/{resi}/assign [post]
func (c *ShipmentController) AssignCourier(ctx *fiber.Ctx) error {
	resi := ctx.Params("resi")

	type Input struct {
		CourierID uint `json:"courier_id"`
	}
	input := new(Input)
	if err := ctx.BodyParser(input); err != nil {
		return utils.ResponseJSON(ctx, fiber.StatusBadRequest, "Invalid input", nil)
	}

	if err := c.service.AssignCourier(resi, input.CourierID); err != nil {
		return utils.ResponseJSON(ctx, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return utils.ResponseJSON(ctx, fiber.StatusOK, "Courier assigned successfully", nil)
}

// UpdateStatus godoc
// @Summary Update Shipment Status
// @Description Update status pengiriman (oleh Kurir/Admin).
// @Tags Shipment
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param resi path string true "Nomor Resi"
// @Param request body map[string]string true "Status Update Data"
// @Success 200 {object} map[string]interface{} "Success"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Router /shipments/{resi}/status [patch]
func (c *ShipmentController) UpdateStatus(ctx *fiber.Ctx) error {
	resi := ctx.Params("resi")

	type UpdateInput struct {
		Status   string `json:"status"`
		Location string `json:"location"`
		Note     string `json:"note"`
	}
	input := new(UpdateInput)
	if err := ctx.BodyParser(input); err != nil {
		return utils.ResponseJSON(ctx, fiber.StatusBadRequest, "Invalid input", nil)
	}

	userID := ctx.Locals("userID").(uint)

	if err := c.service.UpdateStatus(resi, input.Status, input.Location, input.Note, userID); err != nil {
		return utils.ResponseJSON(ctx, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return utils.ResponseJSON(ctx, fiber.StatusOK, "Status updated", nil)
}

// UploadPOD godoc
// @Summary Upload Proof of Delivery
// @Description Upload bukti foto penerimaan paket (Format: multipart/form-data).
// @Tags Shipment
// @Accept multipart/form-data
// @Produce json
// @Security BearerAuth
// @Param resi path string true "Nomor Resi"
// @Param pod_image formData file true "POD Image File"
// @Success 200 {object} map[string]interface{} "Success"
// @Failure 400 {object} map[string]interface{} "Image required"
// @Failure 500 {object} map[string]interface{} "Server Error"
// @Router /shipments/{resi}/pod [post]
func (c *ShipmentController) UploadPOD(ctx *fiber.Ctx) error {
	resi := ctx.Params("resi")

	file, err := ctx.FormFile("pod_image")
	if err != nil {
		return utils.ResponseJSON(ctx, fiber.StatusBadRequest, "Image is required", nil)
	}

	filePath := "./uploads/" + file.Filename
	if err := ctx.SaveFile(file, filePath); err != nil {
		return utils.ResponseJSON(ctx, fiber.StatusInternalServerError, "Failed to save image", nil)
	}

	userID := ctx.Locals("userID").(uint)

	if err := c.service.UploadPOD(resi, filePath, userID); err != nil {
		return utils.ResponseJSON(ctx, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return utils.ResponseJSON(ctx, fiber.StatusOK, "POD Uploaded", fiber.Map{"image_url": filePath})
}

// GenerateAirwayBillPDF godoc
// @Summary Generate Airway Bill PDF
// @Description Generate dan download label cetak PDF untuk pengiriman.
// @Tags Shipment
// @Produce application/pdf
// @Security BearerAuth
// @Param resi path string true "Nomor Resi"
// @Success 200 {file} file "PDF File"
// @Failure 404 {object} map[string]interface{} "Not Found"
// @Router /shipments/{resi}/pdf [get]
func (c *ShipmentController) GenerateAirwayBillPDF(ctx *fiber.Ctx) error {
	resi := ctx.Params("resi")

	shipment, err := c.service.TrackShipment(resi)
	if err != nil {
		return utils.ResponseJSON(ctx, fiber.StatusNotFound, "Resi not found", nil)
	}

	pdfBytes, err := c.pdfService.GenerateAirwayBill(shipment)
	if err != nil {
		return utils.ResponseJSON(ctx, fiber.StatusInternalServerError, "Failed to generate PDF", nil)
	}

	ctx.Set("Content-Type", "application/pdf")
	ctx.Set("Content-Disposition", "inline; filename="+resi+".pdf")
	return ctx.Send(pdfBytes)
}

// GetDashboardStats godoc
// @Summary Get Dashboard Statistics
// @Description Mengambil ringkasan statistik pengiriman.
// @Tags Shipment
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "Success"
// @Router /shipments/stats [get]
func (c *ShipmentController) GetStats(ctx *fiber.Ctx) error {
	stats, err := c.service.GetDashboardStats()
	if err != nil {
		return utils.ResponseJSON(ctx, fiber.StatusInternalServerError, err.Error(), nil)
	}
	return utils.ResponseJSON(ctx, fiber.StatusOK, "Success", stats)
}
