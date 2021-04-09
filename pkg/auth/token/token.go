package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/redhat-developer/app-services-cli/internal/localizer"
	"github.com/redhat-developer/app-services-cli/pkg/logging"
)

// Token contains the current access and refresh tokens from the Authorization server
type Token struct {
	Logger       logging.Logger
	AccessToken  string `json:"access_token,omitempty" doc:"Bearer access token."`
	RefreshToken string `json:"refresh_token,omitempty" doc:"Offline or refresh token."`
}

func (t *Token) IsValid() (tokenIsValid bool, err error) {
	now := time.Now()
	if t.AccessToken != "" {
		var expires bool
		var left time.Duration
		expires, left, err = GetExpiry(t.AccessToken, now)
		if err != nil {
			return
		}
		if !expires || left > 5*time.Second {
			tokenIsValid = true
			return
		}
	}
	if t.RefreshToken != "" {
		var expires bool
		var left time.Duration
		expires, left, err = GetExpiry(t.RefreshToken, now)
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

// NeedsRefresh checks if the access token is missing,
// expired or nearing expiry and should be refreshed
func (t *Token) NeedsRefresh() bool {
	if t.AccessToken == "" && t.RefreshToken != "" {
		return true
	}

	now := time.Now()
	expires, left, err := GetExpiry(t.AccessToken, now)
	if err != nil {
		t.Logger.Debug(localizer.MustLocalize(&localizer.Config{
			MessageID: "connection.tokenNeedsRefresh.log.debug.expiryCheckError",
			TemplateData: map[string]interface{}{
				"Reason": err.Error(),
			},
		}))
		return false
	}

	if !expires || left > 5*time.Minute {
		t.Logger.Debug(localizer.MustLocalize(&localizer.Config{
			MessageID: "connection.tokenNeedsRefresh.log.debug.tokenIsStillValid",
			TemplateData: map[string]interface{}{
				"TimeLeft": left,
			},
		}))
		return false
	}

	return true
}

func Parse(textToken string) (token *jwt.Token, err error) {
	parser := new(jwt.Parser)
	token, _, err = parser.ParseUnverified(textToken, jwt.MapClaims{})
	if err != nil {
		err = fmt.Errorf("%v: %w", localizer.MustLocalizeFromID("auth.token.parse.error.parseError"), err)
		return
	}
	return token, nil
}

func MapClaims(token *jwt.Token) (jwt.MapClaims, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		err := errors.New(localizer.MustLocalize(&localizer.Config{
			MessageID: "auth.token.mapClaims.error.claimsError",
			TemplateData: map[string]interface{}{
				"Claims": claims,
			},
		}))
		return nil, err
	}

	return claims, nil
}

func GetExpiry(tokenStr string, now time.Time) (expires bool,
	left time.Duration, err error) {

	token, err := Parse(tokenStr)
	if err != nil {
		return false, 0, err
	}

	claims, err := MapClaims(token)
	if err != nil {
		return false, 0, err
	}
	var exp float64
	claim, ok := claims["exp"]
	if ok {
		exp, ok = claim.(float64)
		if !ok {
			err = errors.New(localizer.MustLocalize(&localizer.Config{
				MessageID: "auth.token.getExpiry.error.expectedExpiryClaimError",
				TemplateData: map[string]interface{}{
					"Claim": claim,
				},
			}))
			return
		}
	}
	if exp == 0 {
		expires = false
		left = 0
	} else {
		expires = true
		left = time.Unix(int64(exp), 0).Sub(now)
	}

	return
}

func GetUsername(tokenStr string) (username string, ok bool) {
	accessTkn, _ := Parse(tokenStr)
	tknClaims, _ := MapClaims(accessTkn)
	userName, ok := tknClaims["preferred_username"]
	if ok {
		username = fmt.Sprintf("%v", userName)
	}

	return username, ok
}
