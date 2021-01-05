package connection

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"github.com/Nerzal/gocloak/v7"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/auth/token"
)

// Builder contains the configuration and logic needed to connect to `api.openshift.com`.
// Don't create instances of this type directly, use the NewBulder function instead
type Builder struct {
	trustedCAs       *x509.CertPool
	insecure         bool
	accessToken      string
	refreshToken     string
	clientID         string
	scopes           []string
	apiURL           string
	authURL          string
	transportWrapper TransportWrapper
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

func (b *Builder) WithURL(url string) *Builder {
	b.apiURL = url
	return b
}

func (b *Builder) WithAuthURL(authURL string) *Builder {
	b.authURL = authURL
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

// Build uses the configuration stored in the builder to create a new connection. The builder can be
// reused to create multiple connections with the same configuration. It returns a pointer to the
// connection, and an error if something fails when trying to create it.
//
// This operation is potentially lengthy, as it may require network communications. Consider using a
// context and the BuildContext method.
func (b *Builder) Build() (connection *Connection, err error) {
	return b.BuildContext(context.Background())
}

// BuildContext uses the configuration stored in the builder to create a new connection. The builder
// can be reused to create multiple connections with the same configuration. It returns a pointer to
// the connection, and an error if something fails when trying to create it.
func (b *Builder) BuildContext(ctx context.Context) (connection *Connection, err error) {
	if b.clientID == "" {
		return nil, AuthErrorf("Missing client ID")
	}

	if b.accessToken == "" && b.refreshToken == "" {
		return nil, notLoggedInError()
	}

	tkn := token.Token{
		AccessToken:  b.accessToken,
		RefreshToken: b.refreshToken,
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
		rawAPIURL = DefaultURL
	}
	apiURL, err := url.Parse(rawAPIURL)
	if err != nil {
		err = AuthErrorf("unable to parse API URL '%s': %w", rawAPIURL, err)
		return
	}

	authURL, err := url.Parse(b.authURL)
	if err != nil {
		err = AuthErrorf("unable to parse Auth URL '%s': %w", DefaultAuthURL, err)
		return
	}

	// Create the cookie jar:
	jar, err := b.createCookieJar()
	if err != nil {
		return
	}

	// Create the transport:
	transport := b.createTransport()
	if err != nil {
		return
	}

	client := &http.Client{
		Jar:       jar,
		Transport: transport,
	}

	baseAuthURL := fmt.Sprintf("%v://%v", authURL.Scheme, authURL.Host)

	authClient := gocloak.NewClient(baseAuthURL)
	restyClient := *authClient.RestyClient()
	// #nosec 402
	restyClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: b.insecure})
	authClient.SetRestyClient(&restyClient)

	connection = &Connection{
		insecure:   b.insecure,
		trustedCAs: b.trustedCAs,
		clientID:   b.clientID,
		scopes:     scopes,
		apiURL:     apiURL,
		client:     client,
		authClient: authClient,
		Token:      &tkn,
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
		Proxy: http.ProxyFromEnvironment,
	}

	// Wrap the transport with the round trippers provided by the user:
	if b.transportWrapper != nil {
		transport = b.transportWrapper(transport)
	}

	return
}

func (b *Builder) createCookieJar() (jar http.CookieJar, err error) {
	jar, err = cookiejar.New(nil)
	return
}
