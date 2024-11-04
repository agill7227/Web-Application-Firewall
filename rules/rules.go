package rules

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type Rule interface {
	Check_request(c *fiber.Ctx) bool
}

type UserAgentRule struct {
	BlockedAgents []string
}

type PathRule struct {
	BlockedPaths []string
}

type SqlRule struct{}
type XssRule struct{}

func (r SqlRule) Check_request(c *fiber.Ctx) bool {
	return Define_sql(c)
}

func (r XssRule) Check_request(c *fiber.Ctx) bool {
	return Define_xss("xss_payloads.txt", c)
}

func (r UserAgentRule) Check_request(c *fiber.Ctx) bool {
	userAgent := c.Get("User-Agent")

	for _, agent := range r.BlockedAgents {
		if strings.Contains(userAgent, agent) {
			return true
		}
	}
	return false
}

func (r PathRule) Check_request(c *fiber.Ctx) bool {
	path := c.Path()

	for _, blockedPath := range r.BlockedPaths {
		if strings.Contains(path, blockedPath) {
			return true
		}
	}
	return false
}

func New_user_agent_rule(filepath string) UserAgentRule {
	blockedAgents := []string{}

	file, err := os.Open(filepath)

	if err != nil {
		log.Fatalf("Failed to open the user agent file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		blockedAgents = append(blockedAgents, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Failed to read the user agent file: %v", err)
	}

	return UserAgentRule{BlockedAgents: blockedAgents}

}

func New_path_rule(filepath string) PathRule {
	blockedPaths := []string{}

	file, err := os.Open(filepath)

	if err != nil {
		log.Fatalf("Failed to open the path file: %v", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		blockedPaths = append(blockedPaths, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Failed to read the path file: %v", err)
	}

	return PathRule{BlockedPaths: blockedPaths}
}

func Check_request(c *fiber.Ctx) bool {
	rules := []Rule{
		New_user_agent_rule("blocked_user_agents.txt"),
		New_path_rule("blocked_paths.txt"),
		SqlRule{},
	}

	for _, rule := range rules {
		if rule.Check_request(c) {
			log.Printf("Request blocked by rule: %T", rule)
			return true
		}
	}
	return false

}
