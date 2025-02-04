package middleware

import (
	"github.com/gofiber/fiber/v2"
)

// Setup middleware function
func Setup() fiber.Handler {

	return func(c *fiber.Ctx) error {
		// JWT Middelware
		return AuthAndRefreshMiddleware(c)
	}
}
