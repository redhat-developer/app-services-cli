package profile

import (
	"testing"

	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/internal/mockutil"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/localize/goi18n"
)

func TestEnableDevPreviewConfig(t *testing.T) {
	localizer, _ := goi18n.New(nil)
	testVal := true
	factoryObj := factory.New("dev", localizer)

	factoryObj.Config = mockutil.NewConfigMock(&config.Config{})
	config, err := EnableDevPreview(factoryObj, testVal)
	if config.DevPreviewEnabled == false {
		t.Errorf("TestEnableDevPreviewConfig config = %v, want %v", config.DevPreviewEnabled, true)
	}
	if err != nil {
		t.Errorf("TestEnableDevPreviewConfig error = %v, want %v", err, nil)
	}

	testVal = false
	config, err = EnableDevPreview(factoryObj, testVal)
	if config.DevPreviewEnabled == true {
		t.Errorf("TestEnableDevPreviewConfig() config.DevPreviewEnabled = %v, want %v", config.DevPreviewEnabled, false)
	}
	if err != nil {
		t.Errorf("TestEnableDevPreviewConfig error = %v, want %v", err, nil)
	}

}

func TestDevPreviewEnabled(t *testing.T) {
	localizer, _ := goi18n.New(nil)
	factoryObj := factory.New("dev", localizer)
	factoryObj.Config = mockutil.NewConfigMock(&config.Config{})
	testVal := false
	_, err := EnableDevPreview(factoryObj, testVal)
	if err != nil {
		t.Errorf("TestEnableDevPreviewConfig error = %v, want %v", err, nil)
	}
	enabled := DevPreviewEnabled(factoryObj)

	if enabled {
		t.Errorf("TestEnableDevPreviewConfig enabled = %v, want %v", enabled, false)
	}

	testVal = true
	_, _ = EnableDevPreview(factoryObj, testVal)

	enabled = DevPreviewEnabled(factoryObj)
	if !enabled {
		t.Errorf("TestEnableDevPreviewConfig enabled = %v, want %v", enabled, true)
	}
}
