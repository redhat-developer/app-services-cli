package kcconnection

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"
	"net/url"

	"github.com/redhat-developer/app-services-cli/pkg/core/auth/token"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"

	"github.com/redhat-developer/app-services-cli/pkg/core/config"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"

	"github.com/redhat-developer/app-services-cli/internal/build"

	"github.com/Nerzal/gocloak/v7"
)

// ConnectionBuilder contains the configuration and logic needed to connect to `api.openshift.com`.
// Don't create instances of this type directly, use the NewConnectionBuilder function instead
type ConnectionBuilder struct {
	trustedCAs        *x509.CertPool
	insecure          bool
	disableKeepAlives bool
	accessToken       string
	refreshToken      string
	clientID          string
	scopes            []string
	apiURL            string
	authURL           string
	consoleURL        string
	config            config.IConfig
	logger            logging.Logger
	transportWrapper  TransportWrapper
	connectionConfig  *connection.Config
}

// TransportWrapper is a wrapper for a transport of type http.RoundTripper.
// Creating a transport wrapper, enables to preform actions and manipulations on the transport
// request and response.
type TransportWrapper func(http.RoundTripper) http.RoundTripper

// NewConnectionBuilder create an builder that knows how to create connections with the default
// configuration.
func NewConnectionBuilder() *ConnectionBuilder {
	return &ConnectionBuilder{}
}

func (b *ConnectionBuilder) WithURL(url string) *ConnectionBuilder {
	b.apiURL = url
	return b
}

func (b *ConnectionBuilder) WithAccessToken(accessToken string) *ConnectionBuilder {
	b.accessToken = accessToken
	return b
}

func (b *ConnectionBuilder) WithRefreshToken(refreshToken string) *ConnectionBuilder {
	b.refreshToken = refreshToken
	return b
}

func (b *ConnectionBuilder) WithTrustedCAs(value *x509.CertPool) *ConnectionBuilder {
	b.trustedCAs = value
	return b
}

func (b *ConnectionBuilder) WithInsecure(insecure bool) *ConnectionBuilder {
	b.insecure = insecure
	return b
}

func (b *ConnectionBuilder) WithTransportWrapper(transportWrapper TransportWrapper) *ConnectionBuilder {
	b.transportWrapper = transportWrapper
	return b
}

func (b *ConnectionBuilder) WithLogger(logger logging.Logger) *ConnectionBuilder {
	b.logger = logger
	return b
}

func (b *ConnectionBuilder) WithConsoleURL(url string) *ConnectionBuilder {
	b.consoleURL = url
	return b
}

func (b *ConnectionBuilder) WithAuthURL(authURL string) *ConnectionBuilder {
	b.authURL = authURL
	return b
}

func (b *ConnectionBuilder) WithClientID(clientID string) *ConnectionBuilder {
	b.clientID = clientID
	return b
}

func (b *ConnectionBuilder) WithScopes(scopes ...string) *ConnectionBuilder {
	b.scopes = append(b.scopes, scopes...)
	return b
}

// DisableKeepAlives disables HTTP keep-alives with the server. This is unrelated to similarly
// named TCP keep-alives.
func (b *ConnectionBuilder) DisableKeepAlives(flag bool) *ConnectionBuilder {
	b.disableKeepAlives = flag
	return b
}

func (b *ConnectionBuilder) WithConfig(cfg config.IConfig) *ConnectionBuilder {
	b.config = cfg
	return b
}

// WithConnectionConfig contains config for the connection instance
func (b *ConnectionBuilder) WithConnectionConfig(cfg *connection.Config) *ConnectionBuilder {
	b.connectionConfig = cfg
	return b
}

// Build uses the configuration stored in the builder to create a new connection. The builder can be
// reused to create multiple connections with the same configuration. It returns a pointer to the
// connection, and an error if something fails when trying to create it.
//
// This operation is potentially lengthy, as it may require network communications. Consider using a
// context and the BuildContext method.
func (b *ConnectionBuilder) Build() (connection *Connection, err error) {
	return b.BuildContext(context.Background())
}

// BuildContext uses the configuration stored in the builder to create a new connection. The builder
// can be reused to create multiple connections with the same configuration. It returns a pointer to
// the connection, and an error if something fails when trying to create it.
// nolint:funlen
func (b *ConnectionBuilder) BuildContext(ctx context.Context) (connection *Connection, err error) {
	if b.connectionConfig.RequireAuth && b.accessToken == "" && b.refreshToken == "" {
		return nil, &AuthError{notLoggedInError()}
	}

	if b.clientID == "" {
		return nil, AuthErrorf("missing client ID")
	}

	if b.config == nil {
		return nil, fmt.Errorf("missing IConfig")
	}

	if b.logger == nil {
		loggerBuilder := logging.NewStdLoggerBuilder()
		debugEnabled := flagutil.DebugEnabled()
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
		err = fmt.Errorf("unable to parse API URL '%s': %w", rawAPIURL, err)
		return
	}

	authURL, err := url.Parse(b.authURL)
	if err != nil {
		err = AuthErrorf("unable to parse Auth URL '%s': %w", b.authURL, err)
		return
	}

	consoleURL, err := url.Parse(b.consoleURL)
	if err != nil {
		err = fmt.Errorf("unable to parse Console URL '%s': %w", b.consoleURL, err)
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

	connection = &Connection{
		insecure:          b.insecure,
		trustedCAs:        b.trustedCAs,
		clientID:          b.clientID,
		consoleURL:        consoleURL,
		scopes:            scopes,
		apiURL:            apiURL,
		defaultHTTPClient: client,
		keycloakClient:    keycloak,
		Token:             &tkn,
		defaultRealm:      kcRealm,
		logger:            b.logger,
		Config:            b.config,
		connectionConfig:  b.connectionConfig,
	}

	return connection, nil
}

func (b *ConnectionBuilder) createTransport() (transport http.RoundTripper) {
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
