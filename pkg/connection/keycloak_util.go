package connection

import (
	"fmt"
	"net/url"
	"strings"
)

// SplitKeycloakRealmURL splits a Keycloak auth URL
// to retrieve the base path and realm separately
func SplitKeycloakRealmURL(u *url.URL) (issuer string, realm string, ok bool) {
	parts := strings.Split(u.Path, "/")
	issuerURL := url.URL{
		Scheme: u.Scheme,
		Host:   u.Host,
	}
	for i, part := range parts {
		if part == "" {
			continue
		}
		// track the path up to the realm part
		// this will be used to form the base URL (issuer)
		issuerURL.Path = fmt.Sprintf("%v/%v", issuerURL.Path, part)
		if part == "realms" {
			realm = parts[i+1]
			ok = true
			break
		}
	}

	if !ok {
		return "", "", false
	}

	return issuerURL.String(), realm, ok
}
