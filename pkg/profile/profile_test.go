package profile

import (
	"testing"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
)

func TestEnableDevPreviewConfig(t *testing.T) {
	testVal := ""
	config, err := EnableDevPreview(factory.New("dev", nil), &testVal)
	if config != nil {
		t.Errorf("TestEnableDevPreviewConfig config = %v, want %v", config, nil)
	}
	if err != nil {
		t.Errorf("TestEnableDevPreviewConfig error = %v, want %v", err, nil)
	}

	testVal = "giberish"
	config, err = EnableDevPreview(factory.New("dev", nil), &testVal)
	if config.DevPreviewEnabled == true {
		t.Errorf("TestEnableDevPreviewConfig() config.DevPreviewEnabled = %v, want %v", config.DevPreviewEnabled, false)
	}
	if err != nil {
		t.Errorf("TestEnableDevPreviewConfig error = %v, want %v", err, nil)
	}

	testVal = "true"
	config, err = EnableDevPreview(factory.New("dev", nil), &testVal)
	if config.DevPreviewEnabled != true {
		t.Errorf("TestEnableDevPreviewConfig config.DevPreviewEnabled = %v, want %v", config.DevPreviewEnabled, true)
	}
	if err != nil {
		t.Errorf("TestEnableDevPreviewConfig error = %v, want %v", err, nil)
	}

	testVal = "yes"
	config, err = EnableDevPreview(factory.New("dev", nil), &testVal)
	if config.DevPreviewEnabled != true {
		t.Errorf("TestEnableDevPreviewConfig config.DevPreviewEnabled = %v, want %v", config.DevPreviewEnabled, true)
	}
	if err != nil {
		t.Errorf("TestEnableDevPreviewConfig error = %v, want %v", err, nil)
	}
}

func TestDevPreviewEnabled(t *testing.T) {
	f := factory.New("dev", nil)
	testVal := "false"
	config, _ := EnableDevPreview(f, &testVal)
	t.Log("Ess", config)
	enabled := DevPreviewEnabled(f)

	if enabled {
		t.Errorf("TestEnableDevPreviewConfig enabled = %v, want %v", enabled, false)
	}

	testVal = "true"
	_, _ = EnableDevPreview(f, &testVal)

	enabled = DevPreviewEnabled(f)
	if !enabled {
		t.Errorf("TestEnableDevPreviewConfig enabled = %v, want %v", enabled, true)
	}
}
