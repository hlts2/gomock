package gomock

type Endpoint struct {
	Path         string `yaml:"path"`
	Method       string `yaml:"method"`
	ResponseFile string `yaml:"response_file"`
}

type Config struct {
	Endpoints []Endpoint `yaml:"endpoints"`
}
