// Package cluster contains commands for interacting with cluster logic of the service directly instead of through the
// REST API exposed via the serve command.
package login

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/redhat-developer/app-services-cli/internal/build"
	"golang.org/x/oauth2"

	"github.com/redhat-developer/app-services-cli/pkg/auth/login"
	"github.com/redhat-developer/app-services-cli/pkg/auth/token"
	"github.com/redhat-developer/app-services-cli/pkg/localize"

	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/debug"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
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
	Logger     logging.Logger
	Connection factory.ConnectionFunc
	IO         *iostreams.IOStreams
	localizer  localize.Localizer

	url                   string
	authURL               string
	masAuthURL            string
	clientID              string
	scopes                []string
	insecureSkipTLSVerify bool
	printURL              bool
	offlineToken          string
}

// NewLoginCmd gets the command that's log the user in
func NewLoginCmd(f *factory.Factory) *cobra.Command {
	opts := &Options{
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
		IO:         f.IOStreams,
		localizer:  f.Localizer,
	}

	cmd := &cobra.Command{
		Use:     opts.localizer.MustLocalize("login.cmd.use"),
		Short:   opts.localizer.MustLocalize("login.cmd.shortDescription"),
		Long:    opts.localizer.MustLocalize("login.cmd.longDescription", localize.NewEntry("OfflineTokenURL", build.OfflineTokenURL)),
		Example: opts.localizer.MustLocalize("login.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			if opts.offlineToken != "" && opts.clientID == build.DefaultClientID {
				opts.clientID = build.DefaultOfflineTokenClientID
			}

			if opts.IO.IsSSHSession() && opts.offlineToken == "" {
				opts.Logger.Info(opts.localizer.MustLocalize("login.log.info.sshLoginDetected", localize.NewEntry("OfflineTokenURL", build.OfflineTokenURL)))
			}

			return runLogin(opts)
		},
	}

	cmd.Flags().StringVar(&opts.url, "api-gateway", build.ProductionAPIURL, opts.localizer.MustLocalize("login.flag.apiGateway"))
	cmd.Flags().BoolVar(&opts.insecureSkipTLSVerify, "insecure", false, opts.localizer.MustLocalize("login.flag.insecure"))
	cmd.Flags().StringVar(&opts.clientID, "client-id", build.DefaultClientID, opts.localizer.MustLocalize("login.flag.clientId"))
	cmd.Flags().StringVar(&opts.authURL, "auth-url", build.ProductionAuthURL, opts.localizer.MustLocalize("login.flag.authUrl"))
	cmd.Flags().StringVar(&opts.masAuthURL, "mas-auth-url", build.ProductionMasAuthURL, opts.localizer.MustLocalize("login.flag.masAuthUrl"))
	cmd.Flags().BoolVar(&opts.printURL, "print-sso-url", false, opts.localizer.MustLocalize("login.flag.printSsoUrl"))
	cmd.Flags().StringArrayVar(&opts.scopes, "scope", connection.DefaultScopes, opts.localizer.MustLocalize("login.flag.scope"))
	cmd.Flags().StringVarP(&opts.offlineToken, "token", "t", "", opts.localizer.MustLocalize("login.flag.token", localize.NewEntry("OfflineTokenURL", build.OfflineTokenURL)))

	return cmd
}

// nolint:funlen
func runLogin(opts *Options) (err error) {
	gatewayURL, err := getURLFromAlias(opts.url, apiGatewayAliases, opts.localizer)
	if err != nil {
		return err
	}

	authURL, err := getURLFromAlias(opts.authURL, authURLAliases, opts.localizer)
	if err != nil {
		return err
	}
	opts.authURL = authURL.String()
	fmt.Println(opts.authURL)

	masAuthURL, err := getURLFromAlias(opts.masAuthURL, masAuthURLAliases, opts.localizer)
	if err != nil {
		return err
	}
	opts.masAuthURL = masAuthURL.String()

	if opts.offlineToken == "" {
		tr := createTransport(opts.insecureSkipTLSVerify)
		httpClient := oauth2.NewClient(context.Background(), nil)
		httpClient.Transport = tr

		loginExec := &login.AuthorizationCodeGrant{
			HTTPClient: httpClient,
			Scopes:     opts.scopes,
			Logger:     opts.Logger,
			IO:         opts.IO,
			Config:     opts.Config,
			ClientID:   opts.clientID,
			PrintURL:   opts.printURL,
			Localizer:  opts.localizer,
		}

		ssoCfg := &login.SSOConfig{
			AuthURL:      authURL,
			RedirectPath: "sso-redhat-callback",
		}

		masSsoCfg := &login.SSOConfig{
			AuthURL:      masAuthURL,
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
	opts.Logger.Info("")

	if !ok {
		opts.Logger.Info(opts.localizer.MustLocalize("login.log.info.loginSuccessNoUsername"))
	} else {
		opts.localizer.MustLocalize("login.log.info.loginSuccess", localize.NewEntry("Username", username))
	}

	// debug mode checks this for a version update also.
	// so we check if is enabled first so as not to print it twice
	if !debug.Enabled() {
		build.CheckForUpdate(context.Background(), opts.Logger, opts.localizer)
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

func getURLFromAlias(urlOrAlias string, urlAliasMap map[string]string, localizer localize.Localizer) (u *url.URL, err error) {
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
		err = errors.New(localizer.MustLocalize("login.error.schemeMissingFromUrl", localize.NewEntry("URL", gatewayURL.String())))
		return nil, err
	}

	return gatewayURL, nil
}
