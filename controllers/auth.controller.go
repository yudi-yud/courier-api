package controllers

import (
	"courier-api/services"
	"courier-api/utils"

	"github.com/gofiber/fiber/v2"
)

// Login godoc
// @Summary Login user
// @Description Login untuk mendapatkan token JWT
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body map[string]string true "Login Data"
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /login [post]

type AuthController struct {
	service services.AuthService
}

func NewAuthController() *AuthController {
	return &AuthController{service: services.NewAuthService()}
}

func (c *AuthController) Login(ctx *fiber.Ctx) error {
	type LoginInput struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	input := new(LoginInput)
	if err := ctx.BodyParser(input); err != nil {
		return utils.ResponseJSON(ctx, fiber.StatusBadRequest, "Invalid input", nil)
	}

	token, err := c.service.Login(input.Username, input.Password)
	if err != nil {
		return utils.ResponseJSON(ctx, fiber.StatusUnauthorized, err.Error(), nil)
	}

	return utils.ResponseJSON(ctx, fiber.StatusOK, "Success", fiber.Map{"token": token})
}

func (c *AuthController) Register(ctx *fiber.Ctx) error {
	type RegisterInput struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}

	input := new(RegisterInput)
	if err := ctx.BodyParser(input); err != nil {
		return utils.ResponseJSON(ctx, fiber.StatusBadRequest, "Invalid input", nil)
	}

	// Hanya admin yang bisa register user baru
	if err := c.service.Register(input.Username, input.Password, input.Role); err != nil {
		return utils.ResponseJSON(ctx, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return utils.ResponseJSON(ctx, fiber.StatusCreated, "User created", nil)
}
