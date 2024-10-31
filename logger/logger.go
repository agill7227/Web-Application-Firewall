package logger

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func LogRequest(c *fiber.Ctx, blocked bool) {

	timestamp := time.Now().Format(time.RFC1123)

	ip := c.IP()

	method := c.Method()
	path := c.Path()
	headers := c.GetReqHeaders()
	body := c.Body()

	bodyStr := "No body"

	if len(body) > 0 {
		bodyStr = string(body)

	}

	status := "ALLOWED"
	if blocked {
		status = "BLOCKED"

	}

	fmt.Printf("[%s] IP: %s | Method: %s | Path: %s | Status: %s\n", timestamp, ip, method, path, status)
	fmt.Printf("Headers: %v\n", headers)
	fmt.Printf("Body: %s\n\n", bodyStr)

}
