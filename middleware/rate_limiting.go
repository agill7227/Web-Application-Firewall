package middleware

import (
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
)

var requestscounts = make(map[string]int)
var mutex = &sync.Mutex{}

const maxRequests = 10
const windowSeconds = 60

func RateLimting(c *fiber.Ctx) error {
	ip := c.IP()
	mutex.Lock()

	requestscounts[ip]++
	if requestscounts[ip] == 1 {
		go func(ip string) {
			time.Sleep(time.Second * windowSeconds)
			mutex.Lock()
			delete(requestscounts, ip)
			mutex.Unlock()
		}(ip)
	}
	count := requestscounts[ip]
	mutex.Unlock()

	if count > maxRequests {
		return c.Status(fiber.StatusTooManyRequests).SendString("Too many requests")
	}
	return c.Next()

}
