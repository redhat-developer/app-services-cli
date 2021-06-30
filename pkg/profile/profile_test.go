package profile

import (
	"testing"

	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/internal/mockutil"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
)

func TestEnableDevPreviewConfig(t *testing.T) {
	testVal := ""
	factoryObj := factory.New("dev", nil)
	factoryObj.Config = mockutil.NewConfigMock(&config.Config{})
	config, err := EnableDevPreview(factoryObj, testVal)
	if config != nil {
		t.Errorf("TestEnableDevPreviewConfig config = %v, want %v", config, nil)
	}
	if err != nil {
		t.Errorf("TestEnableDevPreviewConfig error = %v, want %v", err, nil)
	}

	testVal = "giberish"
	config, err = EnableDevPreview(factoryObj, testVal)
	if config.DevPreviewEnabled == true {
		t.Errorf("TestEnableDevPreviewConfig() config.DevPreviewEnabled = %v, want %v", config.DevPreviewEnabled, false)
	}
	if err != nil {
		t.Errorf("TestEnableDevPreviewConfig error = %v, want %v", err, nil)
	}

	testVal = "true"
	config, err = EnableDevPreview(factoryObj, testVal)
	if config.DevPreviewEnabled != true {
		t.Errorf("TestEnableDevPreviewConfig config.DevPreviewEnabled = %v, want %v", config.DevPreviewEnabled, true)
	}
	if err != nil {
		t.Errorf("TestEnableDevPreviewConfig error = %v, want %v", err, nil)
	}

	testVal = "yes"
	config, err = EnableDevPreview(factoryObj, testVal)
	if config.DevPreviewEnabled != true {
		t.Errorf("TestEnableDevPreviewConfig config.DevPreviewEnabled = %v, want %v", config.DevPreviewEnabled, true)
	}
	if err != nil {
		t.Errorf("TestEnableDevPreviewConfig error = %v, want %v", err, nil)
	}
}

func TestDevPreviewEnabled(t *testing.T) {
	factoryObj := factory.New("dev", nil)
	factoryObj.Config = mockutil.NewConfigMock(&config.Config{})
	testVal := "false"
	config, _ := EnableDevPreview(factoryObj, testVal)
	t.Log("Ess", config)
	enabled := DevPreviewEnabled(factoryObj)

	if enabled {
		t.Errorf("TestEnableDevPreviewConfig enabled = %v, want %v", enabled, false)
	}

	testVal = "true"
	_, _ = EnableDevPreview(factoryObj, testVal)

	enabled = DevPreviewEnabled(factoryObj)
	if !enabled {
		t.Errorf("TestEnableDevPreviewConfig enabled = %v, want %v", enabled, true)
	}
}
