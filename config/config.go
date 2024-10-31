package config

// "log"
// "os"
// "gopkg.in/yaml.v2"

type DetectionConfig struct {
	BlockedUserAgentsFile string `yaml:"blocked_user_agents_file"`
	BlockedPathsFile      string `yaml:"blocked_paths_file"`
	XssPayloadsFile       string `yaml:"xss_payloads_file"`
}
