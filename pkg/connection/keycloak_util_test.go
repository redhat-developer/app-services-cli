package connection

import (
	"net/url"
	"testing"
)

func TestSplitKeycloakRealmURL(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name         string
		args         args
		wantProvider string
		wantRealm    string
		wantOk       bool
	}{
		{
			name:         "should split URL successfully",
			args:         args{"https://identity.test.com/auth/realms/fakerealm"},
			wantProvider: "https://identity.test.com/auth/realms",
			wantRealm:    "fakerealm",
			wantOk:       true,
		},
		{
			name:         "should not split the URL as there is no realm",
			args:         args{"https://identity.test.com/auth/norealmhere"},
			wantProvider: "",
			wantRealm:    "",
			wantOk:       false,
		},
	}
	for _, tt := range tests {
		// nolint:scopelint
		t.Run(tt.name, func(t *testing.T) {
			url, _ := url.Parse(tt.args.url)
			gotProvider, gotRealm, gotOk := SplitKeycloakRealmURL(url)
			if gotProvider != tt.wantProvider {
				t.Errorf("SplitKeycloakRealmURL() gotProvider = %v, want %v", gotProvider, tt.wantProvider)
			}
			if gotRealm != tt.wantRealm {
				t.Errorf("SplitKeycloakRealmURL() gotRealm = %v, want %v", gotRealm, tt.wantRealm)
			}
			if gotOk != tt.wantOk {
				t.Errorf("SplitKeycloakRealmURL() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}
