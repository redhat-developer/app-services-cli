package login

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/config"
	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/localizer"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/auth/pkce"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/browser"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/iostreams"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/logging"
	"github.com/coreos/go-oidc"
	"github.com/phayes/freeport"
	"golang.org/x/oauth2"
)

type AuthorizationCodeGrant struct {
	HTTPClient *http.Client
	Config     config.IConfig
	Logger     logging.Logger
	IO         *iostreams.IOStreams
	ClientID   string
	Scopes     []string
	PrintURL   bool
}

type SSOConfig struct {
	AuthURL      string
	RedirectPath string
}

// Execute runs an Authorization Code flow login
// enabling the user to log in to SSO and MAS-SSO in succession
// https://tools.ietf.org/html/rfc6749#section-4.1
func (a *AuthorizationCodeGrant) Execute(ctx context.Context, ssoCfg *SSOConfig, masSSOCfg *SSOConfig) (err error) {
	// log in to SSO
	if err = a.loginSSO(ctx, ssoCfg); err != nil {
		return err
	}

	// log in to MAS-SSO
	if err = a.loginMAS(ctx, masSSOCfg); err != nil {
		return err
	}

	return nil
}

// log the user in to the main authorization server
// this can be configured with the `--auth-url` flag
func (a *AuthorizationCodeGrant) loginSSO(ctx context.Context, cfg *SSOConfig) error {
	a.Logger.Debug("Logging into", cfg.AuthURL, "\n")
	clientCtx, cancel := createClientContext(ctx, a.HTTPClient)
	defer cancel()
	provider, err := oidc.NewProvider(ctx, cfg.AuthURL)
	if err != nil {
		return err
	}

	redirectURL, redirectURLPort, err := createRedirectURL(cfg.RedirectPath)
	if err != nil {
		return err
	}

	oauthConfig := &oauth2.Config{
		ClientID:    a.ClientID,
		Endpoint:    provider.Endpoint(),
		RedirectURL: redirectURL.String(),
		Scopes:      a.Scopes,
	}

	oidcConfig := &oidc.Config{
		ClientID: a.ClientID,
	}

	verifier := provider.Verifier(oidcConfig)
	state, _ := pkce.GenerateVerifier(128)

	// PKCE
	pkceCodeVerifier, err := pkce.GenerateVerifier(128)
	if err != nil {
		return err
	}
	pkceCodeChallenge := pkce.CreateChallenge(pkceCodeVerifier)
	authCodeURL := oauthConfig.AuthCodeURL(state, *pkce.GetAuthCodeURLOptions(pkceCodeChallenge)...)
	a.Logger.Debug(localizer.MustLocalize(&localizer.Config{
		MessageID: "login.log.debug.createdAuthorizationUrl",
		TemplateData: map[string]interface{}{
			"URL": authCodeURL,
		},
	}), "\n")

	// create a localhost server to handle redirects and exchange tokens securely
	sm := http.NewServeMux()
	server := http.Server{
		Handler: sm,
		Addr:    redirectURL.Host,
	}

	sm.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, authCodeURL, http.StatusFound)
	})

	// HTTP handler for the redirect URL
	sm.Handle("/"+redirectURL.Path, &redirectPageHandler{
		CancelContext: cancel,
		Ctx:           clientCtx,
		Port:          redirectURLPort,
		Config:        a.Config,
		Logger:        a.Logger,
		IO:            a.IO,
		ServerAddr:    server.Addr,
		Oauth2Config:  oauthConfig,
		State:         state,
		TokenVerifier: verifier,
		AuthOptions: []oauth2.AuthCodeOption{
			oauth2.SetAuthURLParam("code_verifier", pkceCodeVerifier),
			oauth2.SetAuthURLParam("grant_type", "authorization_code"),
		},
	})

	if err = a.openBrowser(authCodeURL, redirectURL); err != nil {
		return err
	}

	// start the local server
	a.startServer(clientCtx, &server)

	return nil
}

// log in to MAS-SSO
func (a *AuthorizationCodeGrant) loginMAS(ctx context.Context, cfg *SSOConfig) error {
	a.Logger.Debug("Logging into", cfg.AuthURL, "\n")

	clientCtx, cancel := createClientContext(ctx, a.HTTPClient)
	defer cancel()
	provider, err := oidc.NewProvider(ctx, cfg.AuthURL)
	if err != nil {
		return err
	}

	redirectURL, redirectURLPort, err := createRedirectURL(cfg.RedirectPath)
	if err != nil {
		return err
	}

	oauthConfig := &oauth2.Config{
		ClientID:    a.ClientID,
		Endpoint:    provider.Endpoint(),
		RedirectURL: redirectURL.String(),
		Scopes:      a.Scopes,
	}

	oidcConfig := &oidc.Config{
		ClientID: a.ClientID,
	}

	// Configure PKCE challenge and verifier
	// https://tools.ietf.org/html/rfc7636
	verifier := provider.Verifier(oidcConfig)
	state, _ := pkce.GenerateVerifier(128)
	pkceCodeVerifier, err := pkce.GenerateVerifier(128)
	if err != nil {
		return err
	}
	pkceCodeChallenge := pkce.CreateChallenge(pkceCodeVerifier)

	authCodeURL := oauthConfig.AuthCodeURL(state, *pkce.GetAuthCodeURLOptions(pkceCodeChallenge)...)
	a.Logger.Debug(localizer.MustLocalize(&localizer.Config{
		MessageID: "login.log.debug.createdAuthorizationUrl",
		TemplateData: map[string]interface{}{
			"URL": authCodeURL,
		},
	}), "\n")

	sm := http.NewServeMux()
	server := http.Server{
		Handler: sm,
		Addr:    redirectURL.Host,
	}

	sm.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, authCodeURL, http.StatusFound)
	})

	// HTTP handler for the redirect page
	sm.Handle("/"+redirectURL.Path, &masRedirectPageHandler{
		CancelContext: cancel,
		Ctx:           clientCtx,
		Port:          redirectURLPort,
		Config:        a.Config,
		Logger:        a.Logger,
		IO:            a.IO,
		ServerAddr:    server.Addr,
		Oauth2Config:  oauthConfig,
		State:         state,
		TokenVerifier: verifier,
		AuthOptions: []oauth2.AuthCodeOption{
			oauth2.SetAuthURLParam("code_verifier", pkceCodeVerifier),
			oauth2.SetAuthURLParam("grant_type", "authorization_code"),
		},
	})

	if err = a.openBrowser(authCodeURL, redirectURL); err != nil {
		return err
	}

	a.startServer(clientCtx, &server)

	return nil
}

func (a *AuthorizationCodeGrant) openBrowser(authCodeURL string, redirectURL *url.URL) error {
	if a.PrintURL {
		a.Logger.Info(localizer.MustLocalizeFromID("login.log.info.openSSOUrl"), "\n")
		fmt.Fprintln(a.IO.Out, authCodeURL)
		a.Logger.Info("")
	} else {
		openBrowserExec, err := browser.GetOpenBrowserCommand(redirectURL.Scheme + "://" + redirectURL.Host)
		if err != nil {
			return err
		}
		if err = openBrowserExec.Run(); err != nil {
			return err
		}
	}

	return nil
}

// starts the local HTTP webserver to handle redirect from the Auth server
func (a *AuthorizationCodeGrant) startServer(ctx context.Context, server *http.Server) {
	go func() {
		if err := server.ListenAndServe(); err == nil {
			a.Logger.Error(localizer.MustLocalize(&localizer.Config{
				MessageID: "login.log.error.unableToStartServer",
				TemplateData: map[string]interface{}{
					"Error": err,
				},
			}))
		}
	}()
	<-ctx.Done()
}

// create an OIDC client context which is cancellable
func createClientContext(ctx context.Context, httpClient *http.Client) (context.Context, context.CancelFunc) {
	cancelCtx, cancel := context.WithCancel(ctx)
	clientCtx := oidc.ClientContext(cancelCtx, httpClient)

	return clientCtx, cancel
}

// creates a redirect URL with a random port which is available
// on the user's system
func createRedirectURL(path string) (*url.URL, int, error) {
	port, err := freeport.GetFreePort()
	if err != nil {
		return nil, 0, err
	}

	redirectURL := &url.URL{
		Scheme: "http",
		Host:   fmt.Sprintf("localhost:%v", port),
		Path:   path,
	}

	return redirectURL, port, nil
}
