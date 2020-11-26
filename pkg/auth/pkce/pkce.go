package pkce

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"

	"golang.org/x/oauth2"
)

func GenerateVerifier(length int) (string, error) {
	if length > 128 {
		length = 128
	}
	if length < 43 {
		length = 43
	}

	codeVerifier, err := generateRandomString(length)

	return codeVerifier, err
}

func CreateChallenge(verifier string) string {
	sum := sha256.Sum256([]byte(verifier))
	challenge := base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(sum[:])
	return challenge
}

func GetAuthCodeURLOptions(codeChallenge string) *[]oauth2.AuthCodeOption {
	opts := &[]oauth2.AuthCodeOption{
		oauth2.SetAuthURLParam("code_challenge", codeChallenge),
		oauth2.SetAuthURLParam("code_challenge_method", "S256"),
		oauth2.SetAuthURLParam("grant_type", "authorization_code"),
	}

	return opts
}

func generateRandomString(n int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-._~"
	bytes, err := generateRandomBytes(n)
	if err != nil {
		return "", err
	}
	for i, b := range bytes {
		bytes[i] = charset[b%byte(len(charset))]
	}

	return string(bytes), err
}

func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}
