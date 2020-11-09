// Package cluster contains commands for interacting with cluster logic of the service directly instead of through the
// REST API exposed via the serve command.
package login

import (
	"fmt"
	"time"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/golang/glog"

	"github.com/bf2fc6cc711aee1a0c2a/cli/cmd/tools"
	sdk "github.com/openshift-online/ocm-sdk-go"
	"github.com/spf13/cobra"
)

var args struct {
	tokenURL string
	token    string
	url      string
}

const (
	devURL         = "http://localhost:8000"
	productionURL  = "https://api.openshift.com"
	stagingURL     = "https://api.stage.openshift.com"
	integrationURL = "https://api-integration.6943.hive-integration.openshiftapps.com"
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

// NewLoginCommand gets the command that's log the user in
func NewLoginCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "login",
		Short: "Login to Managed Application Services",
		Long:  "Login to Managed Application Services in order to manage your services",
		Run:   runLogin,
	}

	cmd.Flags().StringVar(&args.token, "token", "", "access token that can be used for login")
	cmd.Flags().StringVar(&args.tokenURL, "token-url", sdk.DefaultTokenURL, "OpenID token URL")
	cmd.Flags().StringVar(&args.url, "url", "staging", "URL of the API gateway. The value can be the complete URL or an alias. The valid aliases are 'production', 'staging', 'integration', 'development' and their shorthands.")

	return cmd
}

func runLogin(cmd *cobra.Command, _ []string) {
	cfg, err := config.Load()
	if err != nil {
		fmt.Errorf("Can't load config file: %v", err)
		return
	}
	if cfg == nil {
		cfg = new(config.Config)
	}

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
			fmt.Errorf("Can't parse token '%s': %v", args.token, err)
			return
		}
		tokenType, err := tokenType(parsedToken)
		if err != nil {
			fmt.Errorf("Can't extract type from 'typ' claim of token '%s': %v", args.token, err)
			return
		}

		cfg.TokenURL = args.tokenURL
		switch tokenType {
		case "Bearer":
			cfg.AccessToken = args.token
		case "Refresh", "Offline":
			cfg.RefreshToken = args.token
		case "":
			fmt.Errorf("Don't know how to handle empty type in token '%s'", args.token)
			return
		default:
			fmt.Errorf("Don't know how to handle token type '%s' in token '%s'", tokenType, args.token)
			return
		}

		// Create a connection and get the token to verify that the crendentials are correct:
		connection, err := cfg.Connection()
		if err != nil {
			fmt.Errorf("Can't create connection: %v", err)
			return
		}
		accessToken, refreshToken, err := connection.Tokens()
		if err != nil {
			fmt.Errorf("Can't get token: %v", err)
			return
		}

		// Save the configuration, but clear the user name and password before unless we have
		// explicitly been asked to store them persistently:
		cfg.AccessToken = accessToken
		cfg.RefreshToken = refreshToken
		err = config.Save(cfg)
		if err != nil {
			fmt.Errorf("Can't save config file: %v", err)
			return
		}

		fmt.Println("Successfully logged in using token")
	} else {
		glog.Infof("Redirecting to login page")
		cmd, err := tools.GetOpenBrowserCommand("https://sso.redhat.com/auth/realms/redhat-external/protocol/openid-connect/auth?client_id=cloud-services&redirect_uri=https%3A%2F%2Fcloud.redhat.com%2F&state=d8b10b88-8699-4c9b-80fd-665c39343e53&response_mode=fragment&response_type=code&scope=openid&nonce=7ba8050f-5f7b-4a1c-80dd-0392c922f5f8")
		if err != nil {
			glog.Fatal(err)
		} else {
			cmd.Start()
			time.Sleep(30 * time.Second)
		}
	}
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
