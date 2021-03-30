// Package httputil contains functions that act as middleware for api interactions
package httputil

import (
	"net/http"
	"net/http/httputil"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/logging"
)

type LoggingRoundTripper struct {
	Proxied http.RoundTripper
	Logger  logging.Logger
}

// RoundTrip logs the http response in debug mode
func (c LoggingRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	resp, err := c.Proxied.RoundTrip(r)
	if err != nil {
		return nil, err
	}

	responseDump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		return nil, err
	}

	c.Logger.Debug(string(responseDump))

	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		return nil, err
	}

	c.Logger.Debug(string(requestDump))

	return resp, nil
}
