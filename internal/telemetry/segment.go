package telemetry

import (
	"net"
	"os"
	"path/filepath"

	"gopkg.in/segmentio/analytics-go.v3"
)

// writekey will be the API key used to send data to the correct source on Segment. Default is the dev key
var writeKey = "DBNhpBzdrdof9K7g8OJcvfDfAJnBs1Sb"

// TelemetryProperties contains all of the properties that are sent as the telemetry data
type TelemetryProperties struct {
	Duration int64  `json:"duration"`
	Error    string `json:"error"`
	Success  bool   `json:"success"`
	TTY      bool   `json:"tty"`
	Version  string `json:"version"`
	OS       string `json:"os"`
}

// TelemetryData contains all of the data that is sent to Segment for telemetry
type TelemetryData struct {
	Event      string              `json:"event"`
	Properties TelemetryProperties `json:"properties"`
}

// TelemetryClient client used to send data to Segment
type TelemetryClient struct {
	// SegmentClient helps interact with the segment API
	SegmentClient analytics.Client
	// TelemetryFilePath points to the file containing anonymousID used for tracking odo commands executed by the user
	TelemetryFilePath string
}

// NewClient returns a Client created with the default args
func NewClient() (*TelemetryClient, error) {
	homeDir, _ := os.UserHomeDir()
	return newCustomClient(filepath.Join(homeDir, ".redhat", "anonymousId"),
		analytics.DefaultEndpoint,
	)
}

// newCustomClient returns a Client created with custom args
func newCustomClient(telemetryFilePath string, segmentEndpoint string) (*TelemetryClient, error) {
	// DefaultContext has IP set to 0.0.0.0 so that it does not track user's IP, which it does in case no IP is set
	client, err := analytics.NewWithConfig(writeKey, analytics.Config{
		Endpoint: segmentEndpoint,
		Verbose:  false,
		DefaultContext: &analytics.Context{
			IP: net.IPv4(0, 0, 0, 0),
		},
	})
	if err != nil {
		return nil, err
	}
	return &TelemetryClient{
		SegmentClient:     client,
		TelemetryFilePath: telemetryFilePath,
	}, nil
}

// Close client connection and send the data
func (c *TelemetryClient) Close() error {
	return c.SegmentClient.Close()
}

// Upload prepares the data to be sent to segment and send it once the client connection closes
func (c *TelemetryClient) Upload(data *TelemetryData) error {
	// obtain the user ID
	userId, uerr := GetUserIdentity(c, c.TelemetryFilePath)
	if uerr != nil {
		return uerr
	}

	// add information to the data
	properties := analytics.NewProperties()

	properties = properties.Set("version", data.Properties.Version).
		Set("success", data.Properties.Success).
		Set("duration(ms)", data.Properties.Duration).
		Set("tty", data.Properties.TTY).
		Set("os", data.Properties.OS).
		Set("error", data.Properties.Error)

	// queue the data that has telemetry information
	return c.SegmentClient.Enqueue(analytics.Track{
		UserId:     userId,
		Event:      data.Event,
		Properties: properties,
	})
}
