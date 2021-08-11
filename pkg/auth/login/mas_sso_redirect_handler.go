package login

import (
	"context"
	// embed static HTML file
	_ "embed"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/auth/token"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/localize"
	"github.com/redhat-developer/app-services-cli/pkg/logging"
	"golang.org/x/oauth2"
)

//go:embed static/mas-sso-redirect-page.html
var masSSOredirectHTMLPage string

// handler for the MAS-SSO redirect page
type masRedirectPageHandler struct {
	IO            *iostreams.IOStreams
	Config        config.IConfig
	Logger        logging.Logger
	ServerAddr    string
	Port          int
	AuthURL       *url.URL
	AuthOptions   []oauth2.AuthCodeOption
	State         string
	Oauth2Config  *oauth2.Config
	Ctx           context.Context
	TokenVerifier *oidc.IDTokenVerifier
	CancelContext context.CancelFunc
	Localizer     localize.Localizer
}

// nolint:funlen
func (h *masRedirectPageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logger := h.Logger

	callbackURL := fmt.Sprintf("%v%v", h.ServerAddr, r.URL.String())
	logger.Debug("Redirected to callback URL:", callbackURL)
	logger.Debug()

	if r.URL.Query().Get("state") != h.State {
		http.Error(w, "state did not match", http.StatusBadRequest)
		return
	}

	// nolint:govet
	oauth2Token, err := h.Oauth2Config.Exchange(h.Ctx, r.URL.Query().Get("code"), h.AuthOptions...)
	if err != nil {
		http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
		return
	}
	rawIDToken, ok := oauth2Token.Extra("id_token").(string)
	if !ok {
		http.Error(w, "No id_token field in oauth2 token.", http.StatusInternalServerError)
		return
	}
	idToken, err := h.TokenVerifier.Verify(h.Ctx, rawIDToken)
	if err != nil {
		http.Error(w, "Failed to verify ID Token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	resp := struct {
		OAuth2Token   *oauth2.Token
		IDTokenClaims *json.RawMessage // ID Token payload is just JSON.
	}{oauth2Token, new(json.RawMessage)}

	if err = idToken.Claims(&resp.IDTokenClaims); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	accessTkn, _ := token.Parse(resp.OAuth2Token.AccessToken)
	tknClaims, _ := token.MapClaims(accessTkn)
	userName, ok := tknClaims["preferred_username"]
	rawUsername := "unknown"
	if ok {
		rawUsername = fmt.Sprintf("%v", userName)
	}

	pageTitle := h.Localizer.MustLocalize("login.redirectPage.title")
	pageBody := h.Localizer.MustLocalize("login.masRedirectPage.body", localize.NewEntry("Host", h.AuthURL.Host), localize.NewEntry("Username", rawUsername))

	redirectPage := fmt.Sprintf(masSSOredirectHTMLPage, pageTitle, pageTitle, pageBody)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, redirectPage)

	cfg, err := h.Config.Load()
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}
	// save the received tokens to the user's config
	cfg.MasAccessToken = oauth2Token.AccessToken
	cfg.MasRefreshToken = oauth2Token.RefreshToken

	if err = h.Config.Save(cfg); err != nil {
		logger.Error(err)
		os.Exit(1)
	}

	h.CancelContext()
}
