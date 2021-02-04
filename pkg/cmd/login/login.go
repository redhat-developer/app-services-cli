// Package cluster contains commands for interacting with cluster logic of the service directly instead of through the
// REST API exposed via the serve command.
package login

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/MakeNowJust/heredoc"
	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/config"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/auth/token"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/factory"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/color"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/logging"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/browser"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/connection"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/auth/pkce"

	"github.com/phayes/freeport"

	"golang.org/x/oauth2"

	"github.com/coreos/go-oidc"

	"github.com/spf13/cobra"
)

const (
	devURL          = "http://localhost:8000"
	productionURL   = "https://api.openshift.com"
	stagingURL      = "https://api.stage.openshift.com"
	integrationURL  = "https://api-integration.6943.hive-integration.openshiftapps.com"
	defaultClientID = "rhoas-cli-prod"
)

// HTML template to show on the redirect page
const RedirectURLTemplate = `
<!DOCTYPE html>
<head>
	<link rel="stylesheet" href="https://unpkg.com/@patternfly/patternfly@4.70.2/patternfly.css">
  <title>Welcome to RHOAS</title>
</head>
<body>
	<div class="pf-c-empty-state">
  <div class="pf-c-empty-state__content">
    <i class="fas fa-key pf-c-empty-state__icon" aria-hidden="true"></i>
    <h1 class="pf-c-title pf-m-lg">Welcome to RHOAS</h1>
    <div class="pf-c-empty-state__body">You have successfully logged in to the RHOAS CLI as <span style="font-weight:700">%v</span>. You may now close this tab and return to the command-line.</div>
  </div>
</div>
</body>
`

// When the value of the `--url` option is one of the keys of this map it will be replaced by the
// corresponding value.
var urlAliases = map[string]string{
	"production":  productionURL,
	"prod":        productionURL,
	"prd":         productionURL,
	"staging":     stagingURL,
	"stage":       stagingURL,
	"stg":         stagingURL,
	"integration": integrationURL,
	"int":         integrationURL,
	"dev":         devURL,
	"development": devURL,
}

type Options struct {
	Config config.IConfig
	Logger func() (logging.Logger, error)

	url                   string
	authURL               string
	clientID              string
	scopes                []string
	insecureSkipTLSVerify bool
	printURL              bool
}

// NewLoginCmd gets the command that's log the user in
func NewLoginCmd(f *factory.Factory) *cobra.Command {
	opts := &Options{
		Config: f.Config,
		Logger: f.Logger,
	}

	cmd := &cobra.Command{
		Use:   "login",
		Short: "Log in to RHOAS",
		Long: heredoc.Doc(`
			Log in securely to RHOAS through a web browser.

			This command opens your web browser, where you can enter your credentials.
		`),
		Example: heredoc.Doc(`
			# securely log in to RHOAS by using a web browser
			$ rhoas login

			# print the authentication URL instead of automatically opening the browser
			$ rhoas login --print-sso-url
		`),
		RunE: func(cmd *cobra.Command, _ []string) error {
			return runLogin(opts)
		},
	}

	cmd.Flags().StringVar(&opts.url, "api-gateway", stagingURL, "URL of the API gateway. The value can be the URL or an alias. Valid aliases are: production|staging|integration|development")
	cmd.Flags().BoolVar(&opts.insecureSkipTLSVerify, "insecure", false, "Enables insecure communication with the server. This disables verification of TLS certificates and host names")
	cmd.Flags().StringVar(&opts.clientID, "client-id", defaultClientID, "OpenID client identifier")
	cmd.Flags().StringVar(&opts.authURL, "auth-url", connection.DefaultAuthURL, "The URL of the SSO Authentication server")
	cmd.Flags().BoolVar(&opts.printURL, "print-sso-url", false, "Prints the console login URL, which you can use to log in to RHOAS from a different web browser. This is useful if you need to log in with different credentials than the credentials you used in your default web browser")
	cmd.Flags().StringArrayVar(&opts.scopes, "scope", connection.DefaultScopes, "Override the default OpenID scope. Can be repeated multiple times to specify multiple scopes")

	return cmd
}

// nolint
func runLogin(opts *Options) error {
	logger, err := opts.Logger()
	if err != nil {
		return err
	}

	// If the value of the `--url` is any of the aliases then replace it with the corresponding
	// real URL:
	unparsedGatewayURL, ok := urlAliases[opts.url]
	if !ok {
		unparsedGatewayURL = opts.url
	}

	gatewayURL, err := url.ParseRequestURI(unparsedGatewayURL)
	if err != nil {
		return err
	}
	if gatewayURL.Scheme != "http" && gatewayURL.Scheme != "https" {
		return fmt.Errorf("Scheme missing from URL %v. Please add either 'https' or 'https'.", color.Info(unparsedGatewayURL))
	}

	tr := createTransport(opts.insecureSkipTLSVerify)
	httpClient := &http.Client{Transport: tr}

	parentCtx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ctx := oidc.ClientContext(parentCtx, httpClient)
	provider, err := oidc.NewProvider(ctx, opts.authURL)
	if err != nil {
		return err
	}

	redirectURLPort, err := freeport.GetFreePort()
	if err != nil {
		return err
	}
	redirectURL := url.URL{
		Scheme: "http",
		Host:   fmt.Sprintf("localhost:%v", redirectURLPort),
		Path:   "sso-redhat-callback",
	}
	oauthCfg := oauth2.Config{
		ClientID:    opts.clientID,
		Endpoint:    provider.Endpoint(),
		RedirectURL: redirectURL.String(),
		Scopes:      opts.scopes,
	}

	oidcCfg := &oidc.Config{
		ClientID: oauthCfg.ClientID,
	}

	verifier := provider.Verifier(oidcCfg)

	state, _ := pkce.GenerateVerifier(128)

	// PKCE
	pkceCodeVerifier, err := pkce.GenerateVerifier(128)
	if err != nil {
		return err
	}
	pkceCodeChallenge := pkce.CreateChallenge(pkceCodeVerifier)
	authCodeURL := oauthCfg.AuthCodeURL(state, *pkce.GetAuthCodeURLOptions(pkceCodeChallenge)...)
	logger.Debugf("Created Authorization URL: %v", authCodeURL)

	sm := http.NewServeMux()
	server := http.Server{
		Handler: sm,
		Addr:    redirectURL.Host,
	}

	sm.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, authCodeURL, http.StatusFound)
	})

	sm.HandleFunc("/sso-redhat-callback", func(w http.ResponseWriter, r *http.Request) {
		logger.Debug("Redirected to callback URL")
		if r.URL.Query().Get("state") != state {
			http.Error(w, "state did not match", http.StatusBadRequest)
			return
		}

		oauthExchangeOpts := []oauth2.AuthCodeOption{
			oauth2.SetAuthURLParam("code_verifier", pkceCodeVerifier),
			oauth2.SetAuthURLParam("grant_type", "authorization_code"),
		}

		oauth2Token, err := oauthCfg.Exchange(ctx, r.URL.Query().Get("code"), oauthExchangeOpts...)
		if err != nil {
			http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
			return
		}
		rawIDToken, ok := oauth2Token.Extra("id_token").(string)
		if !ok {
			http.Error(w, "No id_token field in oauth2 token.", http.StatusInternalServerError)
			return
		}
		idToken, err := verifier.Verify(ctx, rawIDToken)
		if err != nil {
			http.Error(w, "Failed to verify ID Token: "+err.Error(), http.StatusInternalServerError)
			return
		}

		resp := struct {
			OAuth2Token   *oauth2.Token
			IDTokenClaims *json.RawMessage // ID Token payload is just JSON.
		}{oauth2Token, new(json.RawMessage)}

		if err = idToken.Claims(&resp.IDTokenClaims); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		cfg, err := opts.Config.Load()
		if err != nil {
			logger.Error(err)
			os.Exit(1)
		}

		cfg.Insecure = opts.insecureSkipTLSVerify
		cfg.ClientID = opts.clientID
		cfg.AuthURL = opts.authURL
		cfg.URL = gatewayURL.String()
		cfg.Scopes = opts.scopes
		cfg.AccessToken = oauth2Token.AccessToken
		cfg.RefreshToken = oauth2Token.RefreshToken

		if err = opts.Config.Save(cfg); err != nil {
			logger.Error(err)
			os.Exit(1)
		}

		accessTkn, _ := token.Parse(resp.OAuth2Token.AccessToken)
		tknClaims, _ := token.MapClaims(accessTkn)
		userName, ok := tknClaims["preferred_username"]
		var rawUsername string = "unknown"
		if ok {
			rawUsername = fmt.Sprintf("%v", userName)
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, RedirectURLTemplate, rawUsername)

		logger.Infof("You are now logged in as %v", color.Info(rawUsername))
		logger.Info("")

		cancel()
	})

	if opts.printURL {
		logger.Info("Open the following URL in your browser to login:\n\n")
		fmt.Println(authCodeURL)
	} else {
		openBrowserExec, _ := browser.GetOpenBrowserCommand(authCodeURL)
		err = openBrowserExec.Run()
		if err != nil {
			return err
		}
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			logger.Errorf("Unable to start server: %v\n", err)
		}
	}()
	<-parentCtx.Done()

	return nil
}

func createTransport(insecure bool) *http.Transport {
	// #nosec 402
	return &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: insecure},
	}
}
