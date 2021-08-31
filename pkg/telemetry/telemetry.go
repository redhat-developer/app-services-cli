package telemetry

import (
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/redhat-developer/app-services-cli/internal/build"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
)

// Telemetry structure
type Telemetry struct {
	telemetryData *TelemetryData
	factory       *factory.Factory
}

// DisableTelemetryEnv is name of environment variable, if set to true it disables telemetry completely
// hiding even the question
const DisableTelemetryEnv = "REDHAT_DISABLE_TELEMETRY"

// Start collecting telemetry data
func CreateTelemetry(f *factory.Factory) *Telemetry {
	if !isTelemetryEnabled() {
		return nil
	}
	data := TelemetryData{}
	data.Properties.Version = build.Version
	data.Properties.TTY = f.IOStreams.CanPrompt()
	data.Properties.OS = runtime.GOOS
	data.Properties.Duration = time.Now().UnixNano()

	return &Telemetry{telemetryData: &data, factory: f}
}

// Finish sending data to telemetry service
func (t Telemetry) Finish(event string, cmdError error) {
	if !isTelemetryEnabled() {
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

// IsTelemetryEnabled returns true if user has consented to telemetry
func isTelemetryEnabled() bool {
	// TODO add config file to store the telemetry enablement
	// TODO ask for permission to collect usage data
	// The env variable gets precedence in this decision.
	// In case a non-bool value was passed to the env var, we ignore it
	disabledTelemetry, _ := strconv.ParseBool(os.Getenv(DisableTelemetryEnv))

	return !disabledTelemetry
}
