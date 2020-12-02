// Package ms is the Managed Services API client
package builders

import (
	"fmt"
	"net/url"
	"os"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/managedservices"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/config"
)

// TODO refactor into separate config class

func BuildClient() *managedservices.APIClient {
	masCfg := managedservices.NewConfiguration()
	cfg, err := config.Load()

	if err != nil {
		fmt.Fprintln(os.Stderr, "Can't load config file: %w", err)
		os.Exit(1)
	}
	if cfg == nil {
		fmt.Fprintln(os.Stderr, "Not logged in, run the login command")
		os.Exit(1)
	}

	armed, err := cfg.Armed()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Can't check if tokens have expired: %w", err)
		os.Exit(1)
	}
	if !armed {
		fmt.Fprintln(os.Stderr, "Tokens have expired, run the login command")
		os.Exit(1)
	}

	apiURL, err := url.Parse(cfg.URL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to parse url '%v': %v", config.AuthURL, err)
		os.Exit(1)
	}

	if apiURL.Scheme == "" {
		apiURL.Scheme = "http"
	}

	masCfg.Scheme = apiURL.Scheme
	masCfg.Host = apiURL.Host

	// Refresh tokens
	if err = cfg.TokenRefresh(); err != nil {
		fmt.Fprintf(os.Stderr, "Unable to refresh access token: %v", err)
		os.Exit(1)
	}
	err = config.Save(cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not save config file: %v", err)
	}

	masCfg.AddDefaultHeader("Authorization", fmt.Sprintf("Bearer %s", cfg.AccessToken))

	return managedservices.NewAPIClient(masCfg)
}
