package login

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/coreos/go-oidc"
	"github.com/markbates/pkger"
	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/internal/localizer"
	"github.com/redhat-developer/app-services-cli/pkg/auth/token"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/logging"
	"golang.org/x/oauth2"
)

// handler for the SSO redirect page
type redirectPageHandler struct {
	IO            *iostreams.IOStreams
	Config        config.IConfig
	Logger        logging.Logger
	ServerAddr    string
	Port          int
	AuthOptions   []oauth2.AuthCodeOption
	State         string
	Oauth2Config  *oauth2.Config
	Ctx           context.Context
	TokenVerifier *oidc.IDTokenVerifier
	CancelContext context.CancelFunc
}

// nolint:funlen
func (h *redirectPageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f, _ := pkger.Open("/static/login/sso-redirect-page.html")

	b := bytes.NewBufferString("")
	if _, err := io.Copy(b, f); err != nil {
		fmt.Fprintln(h.IO.ErrOut, err)
		f.Close()
		os.Exit(1)
	}

	out, _ := ioutil.ReadAll(b)

	h.Logger.Debug(localizer.MustLocalize(&localizer.Config{
		MessageID: "login.log.debug.redirectedToCallbackUrl",
		TemplateData: map[string]interface{}{
			"URL": fmt.Sprintf("%v%v", h.ServerAddr, r.URL.String()),
		},
	}), "\n")

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

	cfg, err := h.Config.Load()
	if err != nil {
		h.Logger.Error(err)
		os.Exit(1)
	}

	username, ok := token.GetUsername(oauth2Token.AccessToken)
	if !ok {
		username = "unknown"
	}

	pageTitle := localizer.MustLocalizeFromID("login.redirectPage.title")
	pageBody := localizer.MustLocalize(&localizer.Config{
		MessageID: "login.redirectPage.body",
		TemplateData: map[string]interface{}{
			"Username": username,
		},
	})

	redirectPage := fmt.Sprintf((string(out)), pageTitle, pageTitle, pageBody)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, redirectPage)

	// save the received tokens to the user's config
	cfg.AccessToken = oauth2Token.AccessToken
	cfg.RefreshToken = oauth2Token.RefreshToken

	if err = h.Config.Save(cfg); err != nil {
		h.Logger.Error(err)
		os.Exit(1)
	}

	h.CancelContext()
}
