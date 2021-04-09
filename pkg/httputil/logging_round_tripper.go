// Package httputil contains functions that act as middleware for api interactions
package httputil

import (
	"net/http"
	"net/http/httputil"

	"github.com/redhat-developer/app-services-cli/pkg/logging"
)

// LoggingRoundTripper implements http.RoundTripper. When set as Transport of http.Client, it executes HTTP requests with logging.
type LoggingRoundTripper struct {
	Proxied http.RoundTripper
	Logger  logging.Logger
}

// RoundTrip logs the http request and response in debug mode
func (c LoggingRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	resp, err := c.Proxied.RoundTrip(r)
	if err != nil {
		return nil, err
	}

	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		return nil, err
	}

	c.Logger.Debug(string(requestDump))

	responseDump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		return nil, err
	}

	c.Logger.Debug(string(responseDump))

	return resp, nil
}
