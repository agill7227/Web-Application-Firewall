package middleware

import (
	"github.com/agill7227/Web-Application-Firewall/logger"
	"github.com/agill7227/Web-Application-Firewall/rules"

	"github.com/gofiber/fiber/v2"
)

func WAFMiddleware(c *fiber.Ctx) error {

	if err := RateLimting(c); err != nil {
		return err
	}

	blocked := rules.Check_request(c)

	logger.LogRequest(c, blocked)

	if blocked {
		return c.Status(fiber.StatusForbidden).SendString("Forbidden")
	}
	if blocked {
		return c.Status(fiber.StatusForbidden).SendString("Forbidden")

	}

	return c.Next()

}
