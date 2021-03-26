package httputil

import (
	"fmt"
	"net/http"
)

type LoggingRoundTripper struct {
	Proxied http.RoundTripper
}

func (c LoggingRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	resp, err := c.Proxied.RoundTrip(r)
	if err != nil {
		return nil, err
	}

	fmt.Println(*resp)

	return resp, nil
}
