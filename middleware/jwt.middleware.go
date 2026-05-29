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

		if !strings.Contains(authHeader, "Bearer ") {
			return utils.ResponseJSON(c, fiber.StatusUnauthorized, "Invalid token format", nil)
		}

		parts := strings.Split(authHeader, "Bearer ")

		if len(parts) < 2 || parts[1] == "" {
			return utils.ResponseJSON(c, fiber.StatusUnauthorized, "Invalid or expired JWT", nil)
		}

		tokenString := parts[1]

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

// Authorize untuk cek role (misal Admin Only)
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
