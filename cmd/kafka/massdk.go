package kafka

import (
	"fmt"
	"os"
	"strings"

	mas "github.com/bf2fc6cc711aee1a0c2a/cli/client/mas"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/config"
)

// TODO refactor into separate config class

func BuildMasClient() *mas.APIClient {

	masCfg := mas.NewConfiguration()
	cfg, err := config.Load()

	if err != nil {
		fmt.Println("Error loading configuration")
		os.Exit(1)
	}

	token := cfg.AccessToken

	if token == "" {
		token = cfg.RefreshToken
	}

	if token == "" {
		fmt.Println("You must be loggen in. To do so use the `rhmas login` command")
		os.Exit(1)
	}

	tokenIsValid, _ := cfg.CheckTokenValidity()

	if !tokenIsValid {
		fmt.Println("Token has expired. Login again using `rhmas login` command")
		os.Exit(1)
	}

	urlSegments := strings.Split(cfg.URL, "://")

	if len(urlSegments) > 1 {
		masCfg.Scheme = urlSegments[0]
		masCfg.Host = urlSegments[1]
	} else {
		masCfg.Host = urlSegments[0]
	}

	masCfg.AddDefaultHeader("Authorization", fmt.Sprintf("Bearer %s", token))
	client := mas.NewAPIClient(masCfg)
	return client
}
