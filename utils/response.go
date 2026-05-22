package utils

import "github.com/gofiber/fiber/v2"

type ResponseFormat struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResponseJSON(c *fiber.Ctx, status int, message string, data interface{}) error {
	return c.Status(status).JSON(ResponseFormat{
		Status:  status,
		Message: message,
		Data:    data,
	})
}
