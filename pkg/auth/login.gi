package auth

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	sdk "github.com/openshift-online/ocm-sdk-go"
)

// Aliases
const (
	devURL         = "http://localhost:8000"
	productionURL  = "https://api.openshift.com"
	stagingURL     = "https://api.stage.openshift.com"
	integrationURL = "https://api-integration.6943.hive-integration.openshiftapps.com"
)

// Login credentials
const (
	// #nosec G101
	DefaultTokenURL     = "https://sso.redhat.com/auth/realms/redhat-external/protocol/openid-connect/token"
	DefaultClientID     = "cloud-services"
	DefaultClientSecret = ""
	DefaultURL          = "https://api.openshift.com"
	DefaultAgent        = "rhmas/" + 
)

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


/**
* Perform loging command
*/
func login(cfg: *Config) {
	// If the value of the `--url` is any of the aliases then replace it with the corresponding
	// real URL:
	gatewayURL, ok := urlAliases[args.url]
	if !ok {
		gatewayURL = args.url
	}
	

	cfg.URL = gatewayURL

	if len(args.token) > 0 {
		var parsedToken *jwt.Token
		parser := new(jwt.Parser)
		parsedToken, _, err = parser.ParseUnverified(args.token, jwt.MapClaims{})
		if err != nil {
			fmt.Fprintf(os.Stderr, "Can't parse token '%s': %v\n", args.token, err)
			os.Exit(1)
		}
		tokenType, err := tokenType(parsedToken)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Can't extract type from 'typ' claim of token '%s': %v\n", args.token, err)
			os.Exit(1)
		}

		cfg.TokenURL = args.tokenURL
		switch tokenType {
		case "Bearer":
			cfg.AccessToken = args.token
		case "Refresh", "Offline":
			cfg.RefreshToken = args.token
		case "":
			fmt.Fprintf(os.Stderr, "Don't know how to handle empty type in token '%s'\n", args.token)
			os.Exit(1)
		default:
			fmt.Fprintf(os.Stderr, "Don't know how to handle token type '%s' in token '%s'\n", tokenType, args.token)
			os.Exit(1)
		}

		// Create a connection and get the token to verify that the crendentials are correct:
		connection, err := cfg.Connection()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Can't create connection: %v", err)
			os.Exit(1)
		}
		accessToken, refreshToken, err := connection.Tokens()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Can't get token: %v\n", err)
			os.Exit(1)
		}

		// Save the configuration, but clear the user name and password before unless we have
		// explicitly been asked to store them persistently:
		cfg.AccessToken = accessToken
		cfg.RefreshToken = refreshToken
		err = config.Save(cfg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Can't save config file: %v\n", err)
			os.Exit(1)
		}
}else{
	fmt.Fprintf(os.Stderr, "Missing token value")
	os.Exit(1)
	
}
fmt.Fprintln(os.Stderr, "Successfully logged in using token")
	return
}


// Connection creates a connection using this configuration.
func (c *Config) Connection() (connection *sdk.Connection, err error) {
	if err != nil {
		return
	}

	builder := sdk.NewConnectionBuilder()
	if c.TokenURL != "" {
		builder.TokenURL(c.TokenURL)
	}

	// TODO read these from CLI
	builder.Client(sdk.DefaultClientID, sdk.DefaultClientSecret)
	builder.Scopes(sdk.DefaultScopes...)
	builder.URL(c.URL)

	tokens := make([]string, 0, 2)
	if c.AccessToken != "" {
		tokens = append(tokens, c.AccessToken)
	}
	if c.RefreshToken != "" {
		tokens = append(tokens, c.RefreshToken)
	}
	if len(tokens) > 0 {
		builder.Tokens(tokens...)
	}
	// disable TLS certification verification for now.
	builder.Insecure(true)

	// Create the connection:
	connection, err = builder.Build()
	if err != nil {
		return
	}

	return
}

// CheckTokenValidity checks if the configuration contains either credentials or tokens that haven't expired, so
// that it can be used to perform authenticated requests.
func (c *Config) CheckTokenValidity() (tokenIsValid bool, err error) {
	now := time.Now()
	if c.AccessToken != "" {
		var expires bool
		var left time.Duration
		var accessToken *jwt.Token
		accessToken, err = parseToken(c.AccessToken)
		if err != nil {
			return
		}
		expires, left, err = sdk.GetTokenExpiry(accessToken, now)
		if err != nil {
			return
		}
		if !expires || left > 5*time.Second {
			tokenIsValid = true
			return
		}
	}
	if c.RefreshToken != "" {
		var expires bool
		var left time.Duration
		var refreshToken *jwt.Token
		refreshToken, err = parseToken(c.RefreshToken)
		if err != nil {
			return
		}
		expires, left, err = sdk.GetTokenExpiry(refreshToken, now)
		if err != nil {
			return
		}
		if !expires || left > 10*time.Second {
			tokenIsValid = true
			return
		}
	}
	return
}

func parseToken(textToken string) (token *jwt.Token, err error) {
	parser := new(jwt.Parser)
	token, _, err = parser.ParseUnverified(textToken, jwt.MapClaims{})
	if err != nil {
		err = fmt.Errorf("can't parse token: %v", err)
		return
	}
	return token, nil
}

// tokenType extracts the value of the `typ` claim. It returns the value as a string, or the empty
// string if there is no such claim.
func tokenType(token *jwt.Token) (typ string, err error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		err = fmt.Errorf("expected map claims but got %T", claims)
		return
	}
	claim, ok := claims["typ"]
	if !ok {
		return
	}
	value, ok := claim.(string)
	if !ok {
		err = fmt.Errorf("expected string 'typ' but got %T", claim)
		return
	}
	typ = value
	return
}
