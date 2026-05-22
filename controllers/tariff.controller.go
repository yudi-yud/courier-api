package controllers

import (
	"courier-api/models"
	"courier-api/repositories"
	"courier-api/utils"

	"github.com/gofiber/fiber/v2"
)

type TariffController struct {
	repo repositories.TariffRepository
}

func NewTariffController() *TariffController {
	return &TariffController{repo: repositories.NewTariffRepository()}
}

func (c *TariffController) Create(ctx *fiber.Ctx) error {
	tariff := new(models.Tariff)
	if err := ctx.BodyParser(tariff); err != nil {
		return utils.ResponseJSON(ctx, 400, "Invalid input", nil)
	}

	if err := c.repo.Create(tariff); err != nil {
		return utils.ResponseJSON(ctx, 500, "Failed to create tariff", nil)
	}

	return utils.ResponseJSON(ctx, 201, "Tariff created", tariff)
}

func (c *TariffController) GetAll(ctx *fiber.Ctx) error {
	tariffs, err := c.repo.FindAll()
	if err != nil {
		return utils.ResponseJSON(ctx, 500, "Failed to fetch tariffs", nil)
	}
	return utils.ResponseJSON(ctx, 200, "Success", tariffs)
}
