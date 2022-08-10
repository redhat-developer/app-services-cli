package login

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/redhat-developer/app-services-cli/internal/build"
	"github.com/redhat-developer/app-services-cli/pkg/core/auth/pkce"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/browser"
	"github.com/redhat-developer/app-services-cli/pkg/core/config"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/phayes/freeport"
	"golang.org/x/oauth2"
)

type AuthorizationCodeGrant struct {
	HTTPClient *http.Client
	Config     config.IConfig
	Logger     logging.Logger
	IO         *iostreams.IOStreams
	Localizer  localize.Localizer
	ClientID   string
	Scopes     []string
	PrintURL   bool
}

type SSOConfig struct {
	AuthURL      *url.URL
	RedirectPath string
}

// Execute runs an Authorization Code flow login
// https://tools.ietf.org/html/rfc6749#section-4.1
func (a *AuthorizationCodeGrant) Execute(ctx context.Context,
	ssoCfg *SSOConfig, apiUrl string) error {
	if err := a.loginSSO(ctx, ssoCfg); err != nil {
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
	provider, err := oidc.NewProvider(ctx, cfg.AuthURL.String())
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
	a.Logger.Debug("Opening Authorization URL:", authCodeURL)
	a.Logger.Debug()

	// create a localhost server to handle redirects and exchange tokens securely
	sm := http.NewServeMux()
	server := http.Server{
		Handler:           sm,
		Addr:              redirectURL.Host,
		ReadHeaderTimeout: build.DefaultLoginTimeout,
	}

	sm.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, authCodeURL, http.StatusFound)
	})

	authURL, err := url.Parse(cfg.AuthURL.String())
	if err != nil {
		return err
	}

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
		AuthURL:       authURL,
		Localizer:     a.Localizer,
		AuthOptions: []oauth2.AuthCodeOption{
			oauth2.SetAuthURLParam("code_verifier", pkceCodeVerifier),
			oauth2.SetAuthURLParam("grant_type", "authorization_code"),
		},
	})

	a.openBrowser(authCodeURL, redirectURL)

	// start the local server
	a.startServer(clientCtx, &server)

	return nil
}

func (a *AuthorizationCodeGrant) openBrowser(authCodeURL string, redirectURL *url.URL) {
	if a.PrintURL {
		a.Logger.Info(a.Localizer.MustLocalize("login.log.info.openSSOUrl"), "\n")
		fmt.Fprintln(a.IO.Out, authCodeURL)
		a.Logger.Info("")
	} else {
		err := browser.Open(redirectURL.Scheme + "://" + redirectURL.Host)
		if err != nil {
			a.printAuthURLFallback(authCodeURL, redirectURL, err)
			return
		}
	}
}

// starts the local HTTP webserver to handle redirect from the Auth server
func (a *AuthorizationCodeGrant) startServer(ctx context.Context, server *http.Server) {
	go func() {
		log.Fatal(server.ListenAndServe())
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

// when there is an error trying to automatically open the browser on the login page
// fallback to printing the URL to the user terminal instead.
func (a *AuthorizationCodeGrant) printAuthURLFallback(authCodeURL string, redirectURL *url.URL, err error) {
	a.PrintURL = true
	a.Logger.Debug("Error opening browser:", err, "\nPrinting Auth URL to console instead")
	a.openBrowser(authCodeURL, redirectURL)
}
