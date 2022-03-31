package generic

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
)

type GenericAPI interface {
	GET(ctx context.Context, path string) (interface{}, *http.Response, error)
	POST(ctx context.Context, path string, body io.Reader) (interface{}, *http.Response, error)
}

// APIConfig defines the available configuration options
// to customize the API client settings
type Config struct {
	// HTTPClient is a custom HTTP client
	HTTPClient *http.Client
	// Debug enables debug-level logging
	Debug bool
	// BaseURL sets a custom API server base URL
	BaseURL string
}

func NewGenericAPIClient(cfg *Config) GenericAPI {
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = http.DefaultClient
	}

	c := APIClient{
		baseURL:    cfg.BaseURL,
		httpClient: cfg.HTTPClient,
	}

	return &c
}

type APIClient struct {
	httpClient *http.Client
	baseURL    string
}

func (c *APIClient) GET(ctx context.Context, path string) (interface{}, *http.Response, error) {
	url := c.baseURL + path

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err, nil, err
	}

	req = req.WithContext(ctx)

	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, resp, err
	}
	if resp.StatusCode > http.StatusBadRequest {
		return nil, resp, errors.New(resp.Status)
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp, err
	}

	return string(b), resp, err
}

func (c *APIClient) POST(ctx context.Context, path string, body io.Reader) (interface{}, *http.Response, error) {
	url := c.baseURL + path
	bodyBinary, err := io.ReadAll(body)

	if err != nil {
		return err, nil, err
	}
	requestBody := strings.NewReader(string(bodyBinary))
	req, err := http.NewRequest("POST", url, requestBody)
	if err != nil {
		return nil, nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, resp, err
	}
	if resp.StatusCode > http.StatusBadRequest {
		return nil, resp, errors.New(resp.Status)
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp, err
	}

	return string(b), resp, err
}
