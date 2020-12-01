// Package cluster contains commands for interacting with cluster logic of the service directly instead of through the
// REST API exposed via the serve command.
package login

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/browser"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/auth/pkce"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/config"

	"github.com/phayes/freeport"

	"golang.org/x/oauth2"

	"github.com/coreos/go-oidc"

	"github.com/spf13/cobra"
)

const (
	devURL            = "http://localhost:8000"
	productionURL     = "https://api.openshift.com"
	stagingURL        = "https://api.stage.openshift.com"
	integrationURL    = "https://api-integration.6943.hive-integration.openshiftapps.com"
	productionAuthURL = "https://sso.qa.redhat.com/auth/realms/redhat-external"
	defaultClientID   = "rhoas-cli"
)

const PostLoginPage = `
<link rel="preconnect" href="https://fonts.gstatic.com">
<link href="https://fonts.googleapis.com/css2?family=Red+Hat+Display&display=swap" rel="stylesheet">
<style>
.content {
	font-family: 'Red Hat Display', sans-serif;
	margin: auto;
  width: 50%;
	padding: 10px;
	margin-top: 350px;
	text-align: center;
}
</style> 

<div class="content">
<h1>Logged in to RHOAS. Return to your terminal to begin.</h1>
</div>
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

var args struct {
	url                   string
	authURL               string
	clientID              string
	insecureSkipTLSVerify bool
}

// NewLoginCmd gets the command that's log the user in
func NewLoginCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "login",
		Short: "Login to Managed Application Services",
		Long:  "Login to Managed Application Services in order to manage your services",
		RunE:  runLogin,
	}

	cmd.Flags().StringVar(&args.url, "url", stagingURL, "URL of the API gateway. The value can be the complete URL or an alias. The valid aliases are 'production', 'staging', 'integration', 'development' and their shorthands.")
	cmd.Flags().StringVar(&args.authURL, "auth-url", productionAuthURL, "URL of the authorization server.")
	cmd.Flags().BoolVar(&args.insecureSkipTLSVerify, "insecure", false, "Enables insecure communication with the server. This disables verification of TLS certificates and host names.")
	cmd.Flags().StringVar(&args.clientID, "client-id", defaultClientID, "OpenID client identifier.")

	return cmd
}

// nolint
func runLogin(cmd *cobra.Command, _ []string) error {
	cfg, _ := config.Load()
	cfg.SetInsecure(args.insecureSkipTLSVerify)

	// If the value of the `--url` is any of the aliases then replace it with the corresponding
	// real URL:
	gatewayURL, ok := urlAliases[args.url]
	if !ok {
		gatewayURL = args.url
	}

	var authURL string
	if args.authURL != "" {
		authURL = args.authURL
	}

	httpClient := cfg.CreateHTTPClient()

	parentCtx, cancel := context.WithCancel(context.Background())
	ctx := oidc.ClientContext(parentCtx, httpClient)
	provider, err := oidc.NewProvider(ctx, authURL)
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
		ClientID:    args.clientID,
		Endpoint:    provider.Endpoint(),
		RedirectURL: redirectURL.String(),
		Scopes:      []string{oidc.ScopeOpenID},
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

	sm := http.NewServeMux()
	server := http.Server{
		Handler: sm,
		Addr:    redirectURL.Host,
	}

	sm.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, authCodeURL, http.StatusFound)
	})

	sm.HandleFunc("/sso-redhat-callback", func(w http.ResponseWriter, r *http.Request) {
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

		cfg.SetClientID(args.clientID)
		cfg.SetAuthURL(authURL)
		cfg.SetURL(gatewayURL)
		cfg.SetScopes(oauthCfg.Scopes)
		cfg.SetInsecure(args.insecureSkipTLSVerify)
		cfg.SetAccessToken(oauth2Token.AccessToken)
		cfg.SetRefreshToken(oauth2Token.RefreshToken)

		if err = config.Save(cfg); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			return
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, PostLoginPage)
		fmt.Fprintln(os.Stderr, "Successfully logged in to RHOAS")
		cancel()
	})

	openBrowserExec, _ := browser.GetOpenBrowserCommand(authCodeURL)
	_ = openBrowserExec.Run()

	go func() {
		if err := server.ListenAndServe(); err != nil {
			fmt.Fprintf(os.Stderr, "Error starting server: %v", err)
		}
	}()
	<-parentCtx.Done()

	return nil
}
