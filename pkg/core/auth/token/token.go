package token

import (
	"fmt"
	"time"

	"github.com/redhat-developer/app-services-cli/pkg/core/logging"

	"github.com/golang-jwt/jwt/v4"
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
		t.Logger.Debug("Error while checking token expiry:", err)
		return false
	}

	if !expires || left > 5*time.Minute {
		t.Logger.Debug("Token is still valid. Expires in", left)
		return false
	}

	return true
}

func Parse(textToken string) (token *jwt.Token, err error) {
	parser := new(jwt.Parser)
	token, _, err = parser.ParseUnverified(textToken, jwt.MapClaims{})
	if err != nil {
		err = fmt.Errorf("%v: %w", "unable to parse token", err)
		return
	}
	return token, nil
}

func MapClaims(token *jwt.Token) (claims jwt.MapClaims, err error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		err = fmt.Errorf("expected map claims but got \"%v\"", claims)
	}

	return claims, err
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
			err = fmt.Errorf("expected floating point \"exp\" but got \"%v\"", claim)
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

// GetUsername extracts the username claim value from the JWT
func GetUsername(tokenStr string) (username string, ok bool) {
	if tokenStr == "" {
		return "", false
	}
	accessTkn, err := Parse(tokenStr)
	if err != nil {
		return "", false
	}
	tknClaims, _ := MapClaims(accessTkn)
	u, ok := tknClaims["preferred_username"]
	if ok {
		username = fmt.Sprintf("%v", u)
		return
	}
	u, ok = tknClaims["username"]
	if ok {
		username = fmt.Sprintf("%v", u)
		return
	}

	return username, ok
}

// IsOrgAdmin returns the value of the `is_org_admin` claim
func IsOrgAdmin(tokenStr string) bool {
	accessTkn, _ := Parse(tokenStr)
	tknClaims, _ := MapClaims(accessTkn)
	isAdminClaim, ok := tknClaims["is_org_admin"]
	if !ok {
		return false
	}
	orgAdmin, _ := isAdminClaim.(bool)
	return orgAdmin
}
