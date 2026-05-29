package controllers

import (
	"courier-api/models"
	"courier-api/services"
	"courier-api/utils"

	"github.com/gofiber/fiber/v2"
)

type CourierController struct {
	service services.CourierService
}

func NewCourierController() *CourierController {
	return &CourierController{service: services.NewCourierService()}
}

func (c *CourierController) Create(ctx *fiber.Ctx) error {
	courier := new(models.Courier)
	if err := ctx.BodyParser(courier); err != nil {
		return utils.ResponseJSON(ctx, 400, "Invalid input", nil)
	}

	if err := c.service.CreateCourier(courier); err != nil {
		return utils.ResponseJSON(ctx, 500, err.Error(), nil)
	}

	// PERBAIKAN: Panggil service untuk ambil data lengkap (termasuk User/Role)
	createdCourier, err := c.service.GetCourierByID(courier.ID)
	if err != nil {
		// Jika gagal fetch data lengkap, kembalikan data dasar saja (jarang terjadi)
		return utils.ResponseJSON(ctx, 201, "Courier created but failed to fetch details", courier)
	}

	return utils.ResponseJSON(ctx, 201, "Courier created", createdCourier)
}
func (c *CourierController) GetAll(ctx *fiber.Ctx) error {
	couriers, err := c.service.GetAllCouriers()
	if err != nil {
		return utils.ResponseJSON(ctx, 500, err.Error(), nil)
	}
	return utils.ResponseJSON(ctx, 200, "Success", couriers)
}
