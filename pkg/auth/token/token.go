package token

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Token contains the current access and refresh tokens from the Authorization server
type Token struct {
	AccessToken  string `json:"access_token,omitempty" doc:"Bearer access token."`
	RefreshToken string `json:"refresh_token,omitempty" doc:"Offline or refresh token."`
}

func (c *Token) IsValid() (tokenIsValid bool, err error) {
	now := time.Now()
	if c.AccessToken != "" {
		var expires bool
		var left time.Duration
		var accessToken *jwt.Token
		accessToken, err = Parse(c.AccessToken)
		if err != nil {
			return
		}
		expires, left, err = GetExpiry(accessToken, now)
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
		var refreshToken *jwt.Token
		refreshToken, err = Parse(c.RefreshToken)
		if err != nil {
			return
		}
		expires, left, err = GetExpiry(refreshToken, now)
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
		err = Errorf(err.Error())
		return
	}
	return token, nil
}

func MapClaims(token *jwt.Token) (jwt.MapClaims, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		err := Errorf("expected map claims but got %T", claims)
		return nil, err
	}

	return claims, nil
}

func GetExpiry(token *jwt.Token, now time.Time) (expires bool,
	left time.Duration, err error) {
	claims, err := MapClaims(token)
	if err != nil {
		return false, 0, err
	}
	var exp float64
	claim, ok := claims["exp"]
	if ok {
		exp, ok = claim.(float64)
		if !ok {
			err = Errorf("expected floating point 'exp' but got %T", claim)
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
