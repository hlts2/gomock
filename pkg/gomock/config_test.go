package gomock

import (
	"path/filepath"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	var conf Config

	err := LoadConfig(filepath.Join("test_data", "config.yml"), &conf)
	if err != nil {
		t.Errorf("LoadConfig is error: %v", err)
	}

	if len(conf.Port) == 0 {
		t.Errorf("LoadConfig conf is empty")
	}
}
