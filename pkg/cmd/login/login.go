package login

import (
	"context"
	"crypto/tls"
	"errors"
	"net/http"
	"net/url"

	"github.com/redhat-developer/app-services-cli/pkg/core/auth/login"

	"github.com/redhat-developer/app-services-cli/pkg/core/config"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/icon"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/spinner"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection/kcconnection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"

	"github.com/redhat-developer/app-services-cli/internal/build"
	"golang.org/x/oauth2"

	"github.com/spf13/cobra"
)

// When the value of the `--api-gateway` option is one of the keys of this map it will be replaced by the
// corresponding value.
var apiGatewayAliases = map[string]string{
	"production": build.ProductionAPIURL,
	"prod":       build.ProductionAPIURL,
	"staging":    build.StagingAPIURL,
	"stage":      build.StagingAPIURL,
}

// When the value of the `--auth-url` option is one of the keys of this map it will be replaced by the
// corresponding value.
var authURLAliases = map[string]string{
	"production": build.ProductionAuthURL,
	"prod":       build.ProductionAuthURL,
	"staging":    build.ProductionAuthURL,
	"stage":      build.ProductionAuthURL,
}

type options struct {
	Config     config.IConfig
	Logger     logging.Logger
	Connection factory.ConnectionFunc
	IO         *iostreams.IOStreams
	localizer  localize.Localizer
	Context    context.Context

	url                   string
	authURL               string
	deprecatedUrl         string
	clientID              string
	scopes                []string
	insecureSkipTLSVerify bool
	printURL              bool
	offlineToken          string
}

// NewLoginCmd gets the command that's log the user in
func NewLoginCmd(f *factory.Factory) *cobra.Command {
	opts := &options{
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
		IO:         f.IOStreams,
		localizer:  f.Localizer,
		Context:    f.Context,
	}

	cmd := &cobra.Command{
		Use:     "login",
		Short:   opts.localizer.MustLocalize("login.cmd.shortDescription"),
		Long:    opts.localizer.MustLocalize("login.cmd.longDescription", localize.NewEntry("OfflineTokenURL", build.OfflineTokenURL)),
		Example: opts.localizer.MustLocalize("login.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			if opts.offlineToken != "" && opts.clientID == build.DefaultClientID {
				opts.clientID = build.DefaultOfflineTokenClientID
			}

			if opts.IO.IsSSHSession() && opts.offlineToken == "" {
				opts.Logger.Debug(opts.localizer.MustLocalize("login.log.debug.sshLoginDetected", localize.NewEntry("OfflineTokenURL", build.OfflineTokenURL)))
			}

			return runLogin(opts)
		},
	}

	cmd.Flags().StringVar(&opts.url, "api-gateway", build.ProductionAPIURL, opts.localizer.MustLocalize("login.flag.apiGateway"))
	cmd.Flags().BoolVar(&opts.insecureSkipTLSVerify, "insecure", false, opts.localizer.MustLocalize("login.flag.insecure"))
	cmd.Flags().StringVar(&opts.clientID, "client-id", build.DefaultClientID, opts.localizer.MustLocalize("login.flag.clientId"))
	cmd.Flags().StringVar(&opts.authURL, "auth-url", build.ProductionAuthURL, opts.localizer.MustLocalize("login.flag.authUrl"))
	cmd.Flags().StringVar(&opts.deprecatedUrl, "mas-auth-url", "", opts.localizer.MustLocalize("login.flag.masAuthUrl"))
	cmd.Flags().BoolVar(&opts.printURL, "print-sso-url", false, opts.localizer.MustLocalize("login.flag.printSsoUrl"))
	cmd.Flags().StringArrayVar(&opts.scopes, "scope", kcconnection.DefaultScopes, opts.localizer.MustLocalize("login.flag.scope"))
	cmd.Flags().StringVarP(&opts.offlineToken, "token", "t", "", opts.localizer.MustLocalize("login.flag.token", localize.NewEntry("OfflineTokenURL", build.OfflineTokenURL)))

	return cmd
}

// nolint:funlen
func runLogin(opts *options) (err error) {
	gatewayURL, err := getURLFromAlias(opts.url, apiGatewayAliases, opts.localizer)
	if err != nil {
		return err
	}

	authURL, err := getURLFromAlias(opts.authURL, authURLAliases, opts.localizer)
	if err != nil {
		return err
	}
	opts.authURL = authURL.String()

	// log in to SSO
	spinner := spinner.New(opts.IO.ErrOut, opts.localizer)
	spinner.SetLocalizedSuffix("login.log.info.loggingIn")
	spinner.Start()
	if opts.offlineToken == "" {
		tr := createTransport(opts.insecureSkipTLSVerify)
		httpClient := oauth2.NewClient(opts.Context, nil)
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
			RedirectPath: build.SSORedirectPath,
		}

		// Creating a global context with timeout
		ctx, cancel := context.WithTimeout(context.Background(), build.DefaultLoginTimeout)
		defer cancel()

		if err = loginExec.Execute(ctx, ssoCfg, gatewayURL.String()); err != nil {
			spinner.Stop()
			opts.Logger.Info()
			if errors.Is(err, context.DeadlineExceeded) {
				return opts.localizer.MustLocalizeError("login.error.context.deadline.exceeded")
			}

			return err
		}
	}
	spinner.Stop()

	cfg, err := opts.Config.Load()
	if err != nil {
		return err
	}

	if opts.offlineToken != "" {
		cfg.RefreshToken = opts.offlineToken
	}

	cfg.APIUrl = gatewayURL.String()
	cfg.Insecure = opts.insecureSkipTLSVerify
	cfg.ClientID = opts.clientID
	cfg.AuthURL = opts.authURL
	cfg.Scopes = opts.scopes
	// Reset access token on login to avoid reusing previous users valid token
	cfg.AccessToken = ""

	if err = opts.Config.Save(cfg); err != nil {
		return err
	}

	opts.Logger.Info()
	opts.Logger.Info(icon.SuccessPrefix(), opts.localizer.MustLocalize("login.log.info.loginSuccess"))

	return nil
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
		err = localizer.MustLocalizeError("login.error.schemeMissingFromUrl", localize.NewEntry("URL", gatewayURL.String()))
		return nil, err
	}

	return gatewayURL, nil
}
