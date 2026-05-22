package middleware

import (
	"courier-api/utils"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTProtected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return utils.ResponseJSON(c, fiber.StatusUnauthorized, "Missing or malformed JWT", nil)
		}

		tokenString := strings.Split(authHeader, "Bearer ")[1]
		if tokenString == "" {
			return utils.ResponseJSON(c, fiber.StatusUnauthorized, "Missing or malformed JWT", nil)
		}

		token, err := jwt.ParseWithClaims(tokenString, &utils.JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			return utils.ResponseJSON(c, fiber.StatusUnauthorized, "Invalid or expired JWT", nil)
		}

		claims := token.Claims.(*utils.JwtClaims)
		c.Locals("userID", claims.UserID)
		c.Locals("role", claims.Role)

		return c.Next()
	}
}

// Middleware untuk otorisasi role (Admin saja)
func Authorize(roles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userRole := c.Locals("role").(string)
		for _, role := range roles {
			if userRole == role {
				return c.Next()
			}
		}
		return utils.ResponseJSON(c, fiber.StatusForbidden, "You do not have permission", nil)
	}
}
