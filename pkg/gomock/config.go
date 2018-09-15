package gomock

import (
	"os"

	yaml "gopkg.in/yaml.v2"
)

// Endpoint represents endpoint info for API mock
type Endpoint struct {
	Request  Request  `yaml:"request"`
	Response Response `yaml:"response"`
}

// Request represents request info for API mock
type Request struct {
	Path   string `yaml:"path"`
	Method string `yaml:"method"`
}

// Response represents response info for API mock
type Response struct {
	Code    int               `yaml:"code"`
	Body    string            `yaml:"body"`
	Headers map[string]string `yaml:"headers"`
}

// Config is core gomock config struct
type Config struct {
	Port      string    `yaml:"port"`
	Endpoints Endpoints `yaml:"endpoints"`
}

// Endpoints is Endpoint slice
type Endpoints []Endpoint

// LoadConfig loads configuration file of given the path
func LoadConfig(path string, config *Config) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	err = yaml.NewDecoder(f).Decode(config)
	if err != nil {
		return err
	}

	return nil
}
