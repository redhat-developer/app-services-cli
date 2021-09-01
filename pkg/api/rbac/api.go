package rbac

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
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

type RbacAPI struct {
	PrincipalAPI func() PrincipalAPI
}

type PrincipalAPI interface {
	GetPrincipals(ctx context.Context) (*PrincipalList, *http.Response, error)
}

// APIConfig defines the available configuration options
// to customize the API client settings
type Config struct {
	// HTTPClient is a custom HTTP client
	HTTPClient *http.Client
	// Debug enables debug-level logging
	Debug bool
	// BaseURL sets a custom API server base URL
	BaseURL     *url.URL
	AccessToken string
}

// NewPrincipalAPIClient returns a new v1 API client
// using a custom config
func NewPrincipalAPIClient(cfg *Config) PrincipalAPI {
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = http.DefaultClient
	}

	c := APIClient{
		baseURL:     cfg.BaseURL,
		AccessToken: cfg.AccessToken,
		httpClient:  cfg.HTTPClient,
	}

	return &c
}

type APIClient struct {
	httpClient  *http.Client
	baseURL     *url.URL
	AccessToken string
}

func (c *APIClient) GetPrincipals(ctx context.Context) (*PrincipalList, *http.Response, error) {
	rel := url.URL{Path: "/api/rbac/v1/principals/"}
	u := c.baseURL.ResolveReference(&rel)
	req, err := http.NewRequest("GET", u.String(), nil)

	req = req.WithContext(ctx)

	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("Accept", "application/json")
	if c.AccessToken != "" {
		req.Header.Set("Authorization", "Bearer "+c.AccessToken)
	}

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
