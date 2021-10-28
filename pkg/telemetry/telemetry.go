package telemetry

import (
	"os"
	"runtime"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/redhat-developer/app-services-cli/internal/build"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
)

// Telemetry structure
type Telemetry struct {
	telemetryData *TelemetryData
	factory       *factory.Factory
	enabled       bool
}

// DisableTelemetryEnv is name of environment variable, if set to true it disables telemetry completely
// hiding even the question
const DisableTelemetryEnv = "RED_HAT_DISABLE_TELEMETRY"

// Start collecting telemetry data
func CreateTelemetry(f *factory.Factory) (*Telemetry, error) {
	t := &Telemetry{factory: f}
	err := t.Init()
	if err != nil {
		return nil, err
	}
	if t.enabled {
		data := TelemetryData{}
		data.Properties.Version = build.Version
		data.Properties.TTY = f.IOStreams.CanPrompt()
		data.Properties.OS = runtime.GOOS
		data.Properties.Duration = time.Now().UnixNano()
		t.telemetryData = &data
	}
	return t, nil
}

func (t *Telemetry) Init() error {
	// if we are in
	if !t.factory.IOStreams.CanPrompt() {
		t.enabled = false
		return nil
	}

	// The env variable with any value
	if os.Getenv(DisableTelemetryEnv) != "" {
		t.enabled = false
		return nil
	}

	cfg, err := t.factory.Config.Load()
	if err != nil {
		return err
	}

	// Check if user had seen telemetry consent
	if cfg.Telemetry == "" {
		var consentTelemetry bool
		prompt := &survey.Confirm{Message: t.factory.Localizer.MustLocalize("common.telemetry.consent"), Default: false}
		err = survey.AskOne(prompt, &consentTelemetry, nil)
		if err != nil {
			return err
		}
		t.enabled = consentTelemetry
		if consentTelemetry {
			cfg.Telemetry = "enabled"
		} else {
			cfg.Telemetry = "disabled"
		}

		err = t.factory.Config.Save(cfg)
		if err != nil {
			return err
		}
		return nil
	}
	t.enabled = cfg.Telemetry != "disabled"

	return nil
}

// Finish sending data to telemetry service
func (t *Telemetry) Finish(event string, cmdError error) {
	if !t.enabled {
		return
	}
	t.telemetryData.Event = event
	// convert to milliseconds
	t.telemetryData.Properties.Duration = (time.Now().UnixNano() - t.telemetryData.Properties.Duration) / 1000000
	t.telemetryData.Properties.Success = cmdError == nil
	t.telemetryData.Properties.Error = SanitizeError(cmdError)

	telemetryClient, err := NewClient()
	if err != nil {
		t.factory.Logger.Info("Cannot create a telemetryClient %q", err)
	}
	defer telemetryClient.Close()

	err = telemetryClient.Upload(t.telemetryData)
	if err != nil {
		t.factory.Logger.Info("Cannot send data to telemetry: %q", err)
	}

	telemetryClient.Close()

}
