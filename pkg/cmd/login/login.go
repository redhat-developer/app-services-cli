// Package cluster contains commands for interacting with cluster logic of the service directly instead of through the
// REST API exposed via the serve command.
package login

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/auth/login"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/auth/token"

	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/config"
	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/localizer"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/factory"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/iostreams"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/logging"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/connection"

	"github.com/spf13/cobra"
)

const (
	devURL                      = "http://localhost:8000"
	productionURL               = "https://api.openshift.com"
	stagingURL                  = "https://api.stage.openshift.com"
	integrationURL              = "https://api-integration.6943.hive-integration.openshiftapps.com"
	defaultClientID             = "rhoas-cli-prod"
	defaultOfflineTokenClientID = "cloud-services"
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

type Options struct {
	Config     config.IConfig
	Logger     func() (logging.Logger, error)
	Connection factory.ConnectionFunc
	IO         *iostreams.IOStreams

	url                   string
	authURL               string
	masAuthURL            string
	clientID              string
	scopes                []string
	insecureSkipTLSVerify bool
	printURL              bool
	offlineToken          string
	skipMasSSOLogin       bool
}

// NewLoginCmd gets the command that's log the user in
func NewLoginCmd(f *factory.Factory) *cobra.Command {
	opts := &Options{
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
		IO:         f.IOStreams,
	}

	cmd := &cobra.Command{
		Use:     localizer.MustLocalizeFromID("login.cmd.use"),
		Short:   localizer.MustLocalizeFromID("login.cmd.shortDescription"),
		Long:    localizer.MustLocalizeFromID("login.cmd.longDescription"),
		Example: localizer.MustLocalizeFromID("login.cmd.example"),
		RunE: func(cmd *cobra.Command, _ []string) error {
			if opts.offlineToken != "" && opts.clientID == defaultClientID {
				opts.clientID = defaultOfflineTokenClientID
			}
			return runLogin(opts)
		},
	}

	cmd.Flags().StringVar(&opts.url, "api-gateway", stagingURL, localizer.MustLocalizeFromID("login.flag.apiGateway"))
	cmd.Flags().BoolVar(&opts.insecureSkipTLSVerify, "insecure", false, localizer.MustLocalizeFromID("login.flag.insecure"))
	cmd.Flags().StringVar(&opts.clientID, "client-id", defaultClientID, localizer.MustLocalizeFromID("login.flag.clientId"))
	cmd.Flags().StringVar(&opts.authURL, "auth-url", connection.DefaultAuthURL, localizer.MustLocalizeFromID("login.flag.authUrl"))
	cmd.Flags().StringVar(&opts.masAuthURL, "mas-auth-url", connection.DefaultMasAuthURL, localizer.MustLocalizeFromID("login.flag.masAuthUrl"))
	cmd.Flags().BoolVar(&opts.printURL, "print-sso-url", false, localizer.MustLocalizeFromID("login.flag.printSsoUrl"))
	cmd.Flags().StringArrayVar(&opts.scopes, "scope", connection.DefaultScopes, localizer.MustLocalizeFromID("login.flag.scope"))
	cmd.Flags().StringVarP(&opts.offlineToken, "token", "t", "", localizer.MustLocalizeFromID("login.flag.token"))
	cmd.Flags().BoolVar(&opts.skipMasSSOLogin, "skip-mas-login", false, localizer.MustLocalizeFromID("login.flag.skipMasSSOLogin"))

	return cmd
}

// nolint:funlen
func runLogin(opts *Options) (err error) {
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
		return fmt.Errorf(localizer.MustLocalize(&localizer.Config{
			MessageID: "login.error.schemeMissingFromUrl",
			TemplateData: map[string]interface{}{
				"URL": gatewayURL.String(),
			},
		}))
	}

	if opts.offlineToken == "" {
		tr := createTransport(opts.insecureSkipTLSVerify)
		httpClient := &http.Client{Transport: tr}

		loginExec := &login.AuthorizationCodeGrant{
			HTTPClient: httpClient,
			Scopes:     opts.scopes,
			Logger:     logger,
			IO:         opts.IO,
			Config:     opts.Config,
			ClientID:   opts.clientID,
			PrintURL:   opts.printURL,
		}

		ssoCfg := &login.SSOConfig{
			AuthURL:      opts.authURL,
			RedirectPath: "sso-redhat-callback",
		}

		masSsoCfg := &login.SSOConfig{
			AuthURL:      opts.masAuthURL,
			RedirectPath: "mas-sso-callback",
		}

		if err = loginExec.Execute(context.Background(), ssoCfg, masSsoCfg, opts.skipMasSSOLogin); err != nil {
			return err
		}
	}

	if opts.offlineToken != "" {
		if err = loginWithOfflineToken(opts); err != nil {
			return err
		}
	}

	cfg, err := opts.Config.Load()
	if err != nil {
		return err
	}

	cfg.APIUrl = gatewayURL.String()
	cfg.Insecure = opts.insecureSkipTLSVerify
	cfg.ClientID = opts.clientID
	cfg.AuthURL = opts.authURL
	cfg.MasAuthURL = opts.masAuthURL
	cfg.Scopes = opts.scopes

	if err = opts.Config.Save(cfg); err != nil {
		return err
	}

	username, ok := token.GetUsername(cfg.AccessToken)
	if !ok {
		logger.Info("\n", localizer.MustLocalizeFromID("login.log.info.loginSuccessNoUsername"))
	} else {
		logger.Info("\n", localizer.MustLocalize(&localizer.Config{
			MessageID: "login.log.info.loginSuccess",
			TemplateData: map[string]interface{}{
				"Username": username,
			},
		}))
	}

	return nil
}

func loginWithOfflineToken(opts *Options) (err error) {
	cfg, err := opts.Config.Load()
	if err != nil {
		return err
	}
	cfg.Insecure = opts.insecureSkipTLSVerify
	cfg.ClientID = opts.clientID
	cfg.AuthURL = opts.authURL
	cfg.MasAuthURL = opts.masAuthURL
	cfg.Scopes = opts.scopes
	cfg.RefreshToken = opts.offlineToken

	if err = opts.Config.Save(cfg); err != nil {
		return err
	}

	_, err = opts.Connection(connection.DefaultConfigSkipMasAuth)
	return err
}

func createTransport(insecure bool) *http.Transport {
	// #nosec 402
	return &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: insecure},
	}
}
