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

	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/config"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/auth/token"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/factory"

	"github.com/MakeNowJust/heredoc"

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

type LoginOptions struct {
	Config config.IConfig

	url                   string
	authURL               string
	clientID              string
	insecureSkipTLSVerify bool
	printURL              bool
}

// NewLoginCmd gets the command that's log the user in
func NewLoginCmd(f *factory.Factory) *cobra.Command {
	opts := &LoginOptions{
		Config: f.Config,
	}

	cmd := &cobra.Command{
		Use:   "login",
		Short: "Login to Managed Application Services",
		Long:  "Login to Managed Application Services in order to manage your services",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return runLogin(opts)
		},
	}

	cmd.Flags().StringVar(&opts.url, "url", stagingURL, "URL of the API gateway. The value can be the complete URL or an alias. The valid aliases are 'production', 'staging', 'integration', 'development' and their shorthands.")
	cmd.Flags().BoolVar(&opts.insecureSkipTLSVerify, "insecure", false, "Enables insecure communication with the server. This disables verification of TLS certificates and host names.")
	cmd.Flags().StringVar(&opts.clientID, "client-id", defaultClientID, "OpenID client identifier.")
	cmd.Flags().StringVar(&opts.authURL, "auth-url", connection.DefaultAuthURL, "SSO Authentication server")
	cmd.Flags().BoolVar(&opts.printURL, "print-sso-url", false, "Prints the login URL to the console so you can control which browser to open it in. Useful if you need to log in with a user that is different to the one logged in on your default web browser.")

	return cmd
}

// nolint
func runLogin(opts *LoginOptions) error {
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
		return fmt.Errorf("Scheme missing from URL '%v'. Please add either 'https' or 'https'.", unparsedGatewayURL)
	}

	tr := createTransport(opts.insecureSkipTLSVerify)
	httpClient := &http.Client{Transport: tr}

	parentCtx, cancel := context.WithCancel(context.Background())
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

		cfg, err := opts.Config.Load()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not load config: %v\n", err)
			os.Exit(1)
		}

		cfg.Insecure = opts.insecureSkipTLSVerify
		cfg.ClientID = opts.clientID
		cfg.AuthURL = opts.authURL
		cfg.URL = gatewayURL.String()
		cfg.Scopes = oauthCfg.Scopes
		cfg.AccessToken = oauth2Token.AccessToken
		cfg.RefreshToken = oauth2Token.RefreshToken

		if err = opts.Config.Save(cfg); err != nil {
			fmt.Fprintf(os.Stderr, "Could not save config: %v\n", err)
			os.Exit(1)
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintln(w, PostLoginPage)

		accessTkn, _ := token.Parse(resp.OAuth2Token.AccessToken)
		tknClaims, _ := token.MapClaims(accessTkn)
		userName, ok := tknClaims["preferred_username"]
		if !ok {
			fmt.Fprintln(os.Stderr, "\nYou are now logged in")
		} else {
			fmt.Fprintf(os.Stderr, "\nYou are now logged in as %v\n", userName)
		}
		cancel()
		return
	})

	if opts.printURL {
		fmt.Println(heredoc.Docf(`
		Login URL: 
		
		%v`, authCodeURL))
	} else {
		openBrowserExec, _ := browser.GetOpenBrowserCommand(authCodeURL)
		err = openBrowserExec.Run()
		if err != nil {
			return err
		}
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			fmt.Fprintf(os.Stderr, "Error starting server: %v\n", err)
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
