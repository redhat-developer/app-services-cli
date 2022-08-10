package telemetry

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/xtgo/uuid"
	"gopkg.in/segmentio/analytics-go.v3"
)

// GetUserIdentity returns the anonymous ID if it exists, else creates a new one and sends the data to Segment
func GetUserIdentity(client *TelemetryClient, telemetryFilePath string) (string, error) {
	var id []byte

	// Get-or-Create the '$HOME/.redhat' directory
	if err := os.MkdirAll(filepath.Dir(telemetryFilePath), os.ModePerm); err != nil {
		return "", err
	}

	// Get-or-Create the anonymousID file that contains a UUID
	if _, err := os.Stat(telemetryFilePath); !os.IsNotExist(err) {
		id, err = os.ReadFile(telemetryFilePath)
		if err != nil {
			return "", err
		}
	}

	// check if the id is a valid uuid, if not, nil is returned
	uuidValue := string(id)
	if _, err := uuid.Parse(uuidValue); err != nil {
		id = []byte(uuid.NewRandom().String())
		if err := os.WriteFile(telemetryFilePath, id, 0600); err != nil {
			return "", err
		}
		// Since a new ID was created, send the Identify message data that helps identify the user on segment
		if err1 := client.SegmentClient.Enqueue(analytics.Identify{
			UserId: strings.TrimSpace(string(id)),
			Traits: addConfigTraits(),
		}); err1 != nil {
			return "", err1
		}

	}
	return string(id), nil
}

// addConfigTraits adds information about the system
func addConfigTraits() analytics.Traits {
	traits := analytics.NewTraits().Set("os", runtime.GOOS)
	return traits
}
