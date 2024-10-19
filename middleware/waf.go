package middleware

import (
	"WAF/logger"

	"github.com/gofiber/fiber/v2"
)

func WAFMiddleware(c *fiber.Ctx) error {

	blocked := false

	logger.LogRequest(c, blocked)

	if blocked {
		return c.Status(fiber.StatusForbidden).SendString("Forbidden")
	}
	// if isRequestBlocked(c) {
	// 	return c.Status(fiber.StatusForbidden).SendString("Forbidden")

	// }

	return c.Next()

}
