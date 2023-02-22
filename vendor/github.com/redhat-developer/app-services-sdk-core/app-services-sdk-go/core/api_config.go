package core

import "net/http"

// APIConfig defines the available configuration options
// to customize the API client settings
type APIConfig struct {
	// HTTPClient is a custom HTTP client
	HTTPClient *http.Client
	// Debug enables debug-level logging
	Debug bool
	// BaseURL sets a custom API server base URL
	BaseURL string
	// UserAgent sets a custom user-agent
	UserAgent string
}
