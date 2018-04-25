package gomock

import (
	"os"

	"regexp"

	yaml "gopkg.in/yaml.v2"
)

type Endpoint struct {
	Request  Request  `yaml:request`
	Response Response `yaml:"response"`
}

type Request struct {
	Path       string `yaml:"path"`
	Method     string `yaml:"method"`
	RegexRoute *regexp.Regexp
}

type Response struct {
	Code    string            `yaml:"code"`
	Body    string            `yaml:"body"`
	Headers map[string]string `yaml:"headers"`
}

type Config struct {
	Port      string     `yaml:port`
	Endpoints []Endpoint `yaml:"endpoints"`
}

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

	for _, endpoint := range config.Endpoints {
		rgx, err := newRegexRoute(endpoint.Request.Path)
		if err != nil {
			return err
		}
		endpoint.Request.RegexRoute = rgx
	}

	return nil
}
