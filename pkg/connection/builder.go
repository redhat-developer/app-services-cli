package connection

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"
	"net/url"

	"github.com/redhat-developer/app-services-cli/internal/build"

	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/debug"

	"github.com/Nerzal/gocloak/v7"

	"github.com/redhat-developer/app-services-cli/pkg/auth/token"
	"github.com/redhat-developer/app-services-cli/pkg/logging"
)

// Builder contains the configuration and logic needed to connect to `api.openshift.com`.
// Don't create instances of this type directly, use the NewBulder function instead
type Builder struct {
	trustedCAs        *x509.CertPool
	insecure          bool
	disableKeepAlives bool
	accessToken       string
	refreshToken      string
	masAccessToken    string
	masRefreshToken   string
	clientID          string
	scopes            []string
	apiURL            string
	authURL           string
	masAuthURL        string
	config            config.IConfig
	logger            logging.Logger
	transportWrapper  TransportWrapper
	connectionConfig  *Config
}

// TransportWrapper is a wrapper for a transport of type http.RoundTripper.
// Creating a transport wrapper, enables to preform actions and manipulations on the transport
// request and response.
type TransportWrapper func(http.RoundTripper) http.RoundTripper

// NewBuilder create an builder that knows how to create connections with the default
// configuration.
func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) WithAccessToken(accessToken string) *Builder {
	b.accessToken = accessToken
	return b
}

func (b *Builder) WithRefreshToken(refreshToken string) *Builder {
	b.refreshToken = refreshToken
	return b
}

func (b *Builder) WithMASAccessToken(accessToken string) *Builder {
	b.masAccessToken = accessToken
	return b
}

func (b *Builder) WithMASRefreshToken(refreshToken string) *Builder {
	b.masRefreshToken = refreshToken
	return b
}

func (b *Builder) WithTrustedCAs(value *x509.CertPool) *Builder {
	b.trustedCAs = value
	return b
}

func (b *Builder) WithInsecure(insecure bool) *Builder {
	b.insecure = insecure
	return b
}

func (b *Builder) WithTransportWrapper(transportWrapper TransportWrapper) *Builder {
	b.transportWrapper = transportWrapper
	return b
}

func (b *Builder) WithLogger(logger logging.Logger) *Builder {
	b.logger = logger
	return b
}

func (b *Builder) WithURL(url string) *Builder {
	b.apiURL = url
	return b
}

func (b *Builder) WithAuthURL(authURL string) *Builder {
	b.authURL = authURL
	return b
}

func (b *Builder) WithMASAuthURL(authURL string) *Builder {
	b.masAuthURL = authURL
	return b
}

func (b *Builder) WithClientID(clientID string) *Builder {
	b.clientID = clientID
	return b
}

func (b *Builder) WithScopes(scopes ...string) *Builder {
	b.scopes = append(b.scopes, scopes...)
	return b
}

// DisableKeepAlives disables HTTP keep-alives with the server. This is unrelated to similarly
// named TCP keep-alives.
func (b *Builder) DisableKeepAlives(flag bool) *Builder {
	b.disableKeepAlives = flag
	return b
}

func (b *Builder) WithConfig(cfg config.IConfig) *Builder {
	b.config = cfg
	return b
}

// WithConnectionConfig contains config for the connection instance
func (b *Builder) WithConnectionConfig(cfg *Config) *Builder {
	b.connectionConfig = cfg
	return b
}

// Build uses the configuration stored in the builder to create a new connection. The builder can be
// reused to create multiple connections with the same configuration. It returns a pointer to the
// connection, and an error if something fails when trying to create it.
//
// This operation is potentially lengthy, as it may require network communications. Consider using a
// context and the BuildContext method.
func (b *Builder) Build() (connection *KeycloakConnection, err error) {
	return b.BuildContext(context.Background())
}

// BuildContext uses the configuration stored in the builder to create a new connection. The builder
// can be reused to create multiple connections with the same configuration. It returns a pointer to
// the connection, and an error if something fails when trying to create it.
// nolint:funlen
func (b *Builder) BuildContext(ctx context.Context) (connection *KeycloakConnection, err error) {
	if b.connectionConfig.RequireAuth && b.accessToken == "" && b.refreshToken == "" {
		return nil, &AuthError{notLoggedInError()}
	}

	if b.connectionConfig.RequireMASAuth && b.masAccessToken == "" && b.masRefreshToken == "" {
		return nil, &MasAuthError{notLoggedInMASError()}
	}

	if b.clientID == "" {
		return nil, AuthErrorf("missing client ID")
	}

	if b.config == nil {
		return nil, fmt.Errorf("Missing IConfig")
	}

	if b.logger == nil {
		loggerBuilder := logging.NewStdLoggerBuilder()
		debugEnabled := debug.Enabled()
		loggerBuilder = loggerBuilder.Debug(debugEnabled)

		b.logger, err = loggerBuilder.Build()
		if err != nil {
			return nil, err
		}
	}

	tkn := token.Token{
		AccessToken:  b.accessToken,
		RefreshToken: b.refreshToken,
		Logger:       b.logger,
	}

	masTk := token.Token{
		AccessToken:  b.masAccessToken,
		RefreshToken: b.masRefreshToken,
		Logger:       b.logger,
	}

	tokenIsValid, err := tkn.IsValid()
	if err != nil {
		return nil, err
	}
	if !tokenIsValid {
		return nil, sessionExpiredError()
	}

	scopes := b.scopes
	if len(scopes) == 0 {
		scopes = DefaultScopes
	} else {
		scopes = make([]string, len(b.scopes))
		for i := range b.scopes {
			scopes[i] = b.scopes[i]
		}
	}

	// Set the default URL, if needed:
	rawAPIURL := b.apiURL
	if rawAPIURL == "" {
		rawAPIURL = build.ProductionAPIURL
	}
	apiURL, err := url.Parse(rawAPIURL)
	if err != nil {
		err = AuthErrorf("unable to parse API URL '%s': %w", rawAPIURL, err)
		return
	}

	authURL, err := url.Parse(b.authURL)
	if err != nil {
		err = AuthErrorf("unable to parse Auth URL '%s': %w", b.authURL, err)
		return
	}

	masAuthURL, err := url.Parse(b.masAuthURL)
	if err != nil {
		err = AuthErrorf("unable to parse Auth URL '%s': %w", b.masAuthURL, err)
		return
	}

	// Create the transport:
	transport := b.createTransport()
	if err != nil {
		return
	}

	client := &http.Client{
		Transport: transport,
	}

	baseAuthURL := fmt.Sprintf("%v://%v", authURL.Scheme, authURL.Host)

	_, kcRealm, ok := SplitKeycloakRealmURL(authURL)
	if !ok {
		return nil, fmt.Errorf("unable to get realm name from Auth URL: '%s'", b.authURL)
	}

	keycloak := gocloak.NewClient(baseAuthURL)
	restyClient := *keycloak.RestyClient()
	// #nosec 402
	restyClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: b.insecure})
	keycloak.SetRestyClient(&restyClient)

	baseMasAuthURL := fmt.Sprintf("%v://%v", masAuthURL.Scheme, masAuthURL.Host)
	masKc := gocloak.NewClient(baseMasAuthURL)
	masRestyClient := *keycloak.RestyClient()

	_, masKcRealm, ok := SplitKeycloakRealmURL(masAuthURL)
	if !ok {
		return nil, fmt.Errorf("unable to get realm name from Auth URL: '%s'", b.masAuthURL)
	}

	// #nosec 402
	restyClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: b.insecure})
	masKc.SetRestyClient(&masRestyClient)

	connection = &KeycloakConnection{
		insecure:          b.insecure,
		trustedCAs:        b.trustedCAs,
		clientID:          b.clientID,
		scopes:            scopes,
		apiURL:            apiURL,
		defaultHTTPClient: client,
		keycloakClient:    keycloak,
		masKeycloakClient: masKc,
		Token:             &tkn,
		MASToken:          &masTk,
		defaultRealm:      kcRealm,
		masRealm:          masKcRealm,
		logger:            b.logger,
		Config:            b.config,
		connectionConfig:  b.connectionConfig,
	}

	return connection, nil
}

func (b *Builder) createTransport() (transport http.RoundTripper) {
	// Create the raw transport:
	// #nosec 402
	transport = &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: b.insecure,
			RootCAs:            b.trustedCAs,
		},
		Proxy:             http.ProxyFromEnvironment,
		DisableKeepAlives: b.disableKeepAlives,
	}

	// Wrap the transport with the round trippers provided by the user:
	if b.transportWrapper != nil {
		transport = b.transportWrapper(transport)
	}

	return
}
