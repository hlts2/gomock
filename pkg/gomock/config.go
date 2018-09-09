package gomock

import (
	"os"
	"regexp"

	yaml "gopkg.in/yaml.v2"
)

// Endpoint represents endpoint info of API mock
type Endpoint struct {
	Request  Request  `yaml:"request"`
	Response Response `yaml:"response"`
}

// Request represents request info of API mock
type Request struct {
	Path       string `yaml:"path"`
	Method     string `yaml:"method"`
	RegexRoute *regexp.Regexp
}

// Response represents response info of API mock
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

// GetMachingEndpointIndex returns the element number of the endpoint matching the given method name and path
func (endpoints Endpoints) GetMachingEndpointIndex(method, path string) int {
	for i, endpoint := range endpoints {
		ok := endpoint.Request.RegexRoute.MatchString(method + path)
		if ok {
			return i
		}
	}
	return -1
}

// LoadConfig load configuration file of given file path
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

	for i := 0; i < len(config.Endpoints); i++ {
		request := &config.Endpoints[i].Request

		regex, err := newRegexRoute(request.Method + request.Path)
		if err != nil {
			return err
		}

		request.RegexRoute = regex
	}

	return nil
}
