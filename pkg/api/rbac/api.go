package rbac

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"
)

type PrincipalList struct {
	Meta struct {
		Count int `json:"count"`
	} `json:"meta"`
	Links struct {
		First    string `json:"first"`
		Previous string `json:"previous"`
		Next     string `json:"next"`
		Last     string `json:"last"`
	} `json:"links"`
	Data []Principal `json:"data"`
}

type Principal struct {
	Username   string `json:"username"`
	Email      string `json:"email"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	IsActive   bool   `json:"is_active"`
	IsOrgAdmin bool   `json:"is_org_admin"`
}

// RbacAPI defines a collection of APIs grouped under the RBAC API
type RbacAPI struct {
	PrincipalAPI func() PrincipalAPI
}

// PrincipalAPI is the API definition for the RBAC Principal API
type PrincipalAPI interface {
	GetPrincipals(ctx context.Context, opts ...QueryParam) (*PrincipalList, *http.Response, error)
}

// APIConfig defines the available configuration options
// to customize the API client settings
type Config struct {
	// HTTPClient is a custom HTTP client
	HTTPClient *http.Client
	// Debug enables debug-level logging
	Debug bool
	// BaseURL sets a custom API server base URL
	BaseURL *url.URL
}

// NewPrincipalAPIClient returns a new v1 API client
// using a custom config
func NewPrincipalAPIClient(cfg *Config) PrincipalAPI {
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
	baseURL    *url.URL
}

// GetPrincipals returns the list of user's in the current users organization/tenant
func (c *APIClient) GetPrincipals(ctx context.Context, opts ...QueryParam) (*PrincipalList, *http.Response, error) {
	u := resolveURI(c.baseURL, "/api/rbac/v1/principals/", opts...)

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, nil, err
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

	var principalList PrincipalList
	err = json.NewDecoder(resp.Body).Decode(&principalList)
	return &principalList, resp, err
}

// QueryParam is a function defining the query param options return signature
type QueryParam func() (key string, value string)

// WithIntQueryParam accepts an integer query parameter and formats it as a string
func WithIntQueryParam(key string, value int) QueryParam {
	val := strconv.Itoa(value)

	return WithQueryParam(key, val)
}

// WithIntParam accepts a boolean query parameter and formats it as a string
func WithBoolQueryParam(key string, value bool) QueryParam {
	val := strconv.FormatBool(value)

	return WithQueryParam(key, val)
}

// WithQueryParam accepts a string query parameter
func WithQueryParam(key string, value string) QueryParam {
	return func() (string, string) {
		return key, value
	}
}

// resolveURI builds and returns a URI with query parameters
func resolveURI(baseURL *url.URL, path string, opts ...QueryParam) *url.URL {
	rel := url.URL{Path: path}
	u := baseURL.ResolveReference(&rel)
	q := u.Query()
	for _, opt := range opts {
		if opt == nil {
			continue
		}
		key, val := opt()
		q.Set(key, val)
	}
	u.RawQuery = q.Encode()

	return u
}
