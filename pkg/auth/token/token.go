package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/localizer"
	"github.com/dgrijalva/jwt-go"
)

// Token contains the current access and refresh tokens from the Authorization server
type Token struct {
	AccessToken  string `json:"access_token,omitempty" doc:"Bearer access token."`
	RefreshToken string `json:"refresh_token,omitempty" doc:"Offline or refresh token."`
}

func init() {
	localizer.LoadMessageFiles("auth/token")
}

func (c *Token) IsValid() (tokenIsValid bool, err error) {
	now := time.Now()
	if c.AccessToken != "" {
		var expires bool
		var left time.Duration
		expires, left, err = GetExpiry(c.AccessToken, now)
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
		expires, left, err = GetExpiry(c.RefreshToken, now)
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
