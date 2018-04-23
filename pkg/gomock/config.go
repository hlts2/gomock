package gomock

import (
	"os"

	yaml "gopkg.in/yaml.v2"
)

type Endpoint struct {
	Path         string `yaml:"path"`
	Method       string `yaml:"method"`
	ResponseFile string `yaml:"response_file"`
}

type Config struct {
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

	return nil
}
