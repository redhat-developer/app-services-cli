package version

import (
	"context"

	"github.com/google/go-github/github"
)

var RepositoryOwner = "redhat-developer"
var RepositoryName = "app-services-cli"

func GetLatest(ctx context.Context) (*string, error) {
	client := github.NewClient(nil)

	latest, _, err := client.Repositories.GetLatestRelease(ctx, RepositoryOwner, RepositoryName)
	if err != nil {
		return nil, err
	}

	return latest.TagName, nil

	// if latest.TagName != &build.Version {
	// 	logger.Info("")
	// 	logger.Info(color.Info("A new version of rhoas is available:"), color.CodeSnippet(*latest.TagName))
	// 	logger.Info(color.Info(latest.GetHTMLURL()))
	// }
}
