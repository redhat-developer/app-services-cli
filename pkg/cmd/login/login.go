// Package cluster contains commands for interacting with cluster logic of the service directly instead of through the
// REST API exposed via the serve command.
package login

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"

	"github.com/redhat-developer/app-services-cli/internal/build"

	"github.com/redhat-developer/app-services-cli/pkg/auth/login"
	"github.com/redhat-developer/app-services-cli/pkg/auth/token"

	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/internal/localizer"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/debug"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/httputil"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/logging"

	"github.com/redhat-developer/app-services-cli/pkg/connection"

	"github.com/spf13/cobra"
)

// When the value of the `--api-gateway` option is one of the keys of this map it will be replaced by the
// corresponding value.
var apiGatewayAliases = map[string]string{
	"production": build.ProductionAPIURL,
	"prod":       build.ProductionAPIURL,
	"prd":        build.ProductionAPIURL,
	"staging":    build.StagingAPIURL,
	"stage":      build.StagingAPIURL,
	"stg":        build.StagingAPIURL,
}

// When the value of the `--auth-url` option is one of the keys of this map it will be replaced by the
// corresponding value.
var authURLAliases = map[string]string{
	"production": build.ProductionAuthURL,
	"prod":       build.ProductionAuthURL,
	"prd":        build.ProductionAuthURL,
	"staging":    build.ProductionAuthURL,
	"stage":      build.ProductionAuthURL,
	"stg":        build.ProductionAuthURL,
}

// When the value of the `--mas-auth-url` option is one of the keys of this map it will be replaced by the
// corresponding value.
var masAuthURLAliases = map[string]string{
	"production": build.ProductionMasAuthURL,
	"prod":       build.ProductionMasAuthURL,
	"prd":        build.ProductionMasAuthURL,
	"staging":    build.StagingMasAuthURL,
	"stage":      build.StagingMasAuthURL,
	"stg":        build.StagingMasAuthURL,
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
	remoteSession         bool
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
		Use:   localizer.MustLocalizeFromID("login.cmd.use"),
		Short: localizer.MustLocalizeFromID("login.cmd.shortDescription"),
		Long: localizer.MustLocalize(&localizer.Config{
			MessageID: "login.cmd.longDescription",
			TemplateData: map[string]interface{}{
				"OfflineTokenURL": build.OfflineTokenURL,
			},
		}),
		Example: localizer.MustLocalizeFromID("login.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			if opts.offlineToken != "" && opts.clientID == build.DefaultClientID {
				opts.clientID = build.DefaultOfflineTokenClientID
			}

			if 

			return runLogin(opts)
		},
	}

	cmd.Flags().StringVar(&opts.url, "api-gateway", build.ProductionAPIURL, localizer.MustLocalizeFromID("login.flag.apiGateway"))
	cmd.Flags().BoolVar(&opts.insecureSkipTLSVerify, "insecure", false, localizer.MustLocalizeFromID("login.flag.insecure"))
	cmd.Flags().StringVar(&opts.clientID, "client-id", build.DefaultClientID, localizer.MustLocalizeFromID("login.flag.clientId"))
	cmd.Flags().StringVar(&opts.authURL, "auth-url", build.ProductionAuthURL, localizer.MustLocalizeFromID("login.flag.authUrl"))
	cmd.Flags().StringVar(&opts.masAuthURL, "mas-auth-url", build.ProductionMasAuthURL, localizer.MustLocalizeFromID("login.flag.masAuthUrl"))
	cmd.Flags().BoolVar(&opts.printURL, "print-sso-url", false, localizer.MustLocalizeFromID("login.flag.printSsoUrl"))
	cmd.Flags().StringArrayVar(&opts.scopes, "scope", connection.DefaultScopes, localizer.MustLocalizeFromID("login.flag.scope"))
	cmd.Flags().StringVarP(&opts.offlineToken, "token", "t", "", localizer.MustLocalize(&localizer.Config{
		MessageID: "login.flag.token",
		TemplateData: map[string]interface{}{
			"OfflineTokenURL": build.OfflineTokenURL,
		},
	}))

	return cmd
}

// nolint:funlen
func runLogin(opts *Options) (err error) {
	isRemoteLogin := opts.IO.IsSSHSession()

	logger, err := opts.Logger()
	if err != nil {
		return err
	}

	gatewayURL, err := getURLFromAlias(opts.url, apiGatewayAliases)
	if err != nil {
		return err
	}

	authURL, err := getURLFromAlias(opts.authURL, authURLAliases)
	if err != nil {
		return err
	}
	opts.authURL = authURL.String()

	masAuthURL, err := getURLFromAlias(opts.masAuthURL, masAuthURLAliases)
	if err != nil {
		return err
	}
	opts.masAuthURL = masAuthURL.String()

	if opts.offlineToken == "" {
		tr := createTransport(opts.insecureSkipTLSVerify)
		httpClient := &http.Client{
			Transport: httputil.LoggingRoundTripper{
				Proxied: tr,
				Logger:  logger,
			},
		}

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

		if err = loginExec.Execute(context.Background(), ssoCfg, masSsoCfg); err != nil {
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
	logger.Info("")
	if !ok {
		logger.Info(localizer.MustLocalizeFromID("login.log.info.loginSuccessNoUsername"))
	} else {
		logger.Info(localizer.MustLocalize(&localizer.Config{
			MessageID: "login.log.info.loginSuccess",
			TemplateData: map[string]interface{}{
				"Username": username,
			},
		}))
	}

	// debug mode checks this for a version update also.
	// so we check if is enabled first so as not to print it twice
	if !debug.Enabled() {
		build.CheckForUpdate(context.Background(), logger)
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
	// remove MAS-SSO tokens, as this does not support token login
	cfg.MasAccessToken = ""
	cfg.MasRefreshToken = ""

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

func getURLFromAlias(urlOrAlias string, urlAliasMap map[string]string) (*url.URL, error) {
	// If the URL value is any of the aliases then replace it with the corresponding
	// real URL:
	unparsedGatewayURL, ok := urlAliasMap[urlOrAlias]
	if !ok {
		unparsedGatewayURL = urlOrAlias
	}

	gatewayURL, err := url.ParseRequestURI(unparsedGatewayURL)
	if err != nil {
		return nil, err
	}
	if gatewayURL.Scheme != "http" && gatewayURL.Scheme != "https" {
		return nil, fmt.Errorf(localizer.MustLocalize(&localizer.Config{
			MessageID: "login.error.schemeMissingFromUrl",
			TemplateData: map[string]interface{}{
				"URL": gatewayURL.String(),
			},
		}))
	}

	return gatewayURL, nil
}
