package rules

import (
	"bufio"
	"log"
	"os"
	"regexp"

	"github.com/gofiber/fiber/v2"
)

func Define_xss(filepath string, c *fiber.Ctx) bool {

	file, err := os.Open(filepath)

	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}

	defer file.Close()

	var xssPayloads []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		xssPayloads = append(xssPayloads, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Failed to scan file: %s", err)
	}

	body := string(c.Body())

	for _, payload := range xssPayloads {
		match, err := regexp.MatchString(payload, body)
		if err != nil {
			log.Fatalf("Failed to match string: %s", err)
		}
		if match {
			return true
		}
	}
	return false

}
