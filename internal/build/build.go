package build

import (
	"context"
	"regexp"
	"runtime/debug"

	"github.com/google/go-github/github"
	"github.com/redhat-developer/app-services-cli/pkg/color"
	"github.com/redhat-developer/app-services-cli/pkg/logging"
)

// Version is dynamically set by the toolchain or overridden by the Makefile.
var Version = "dev"
var Language = "en"

// RepositoryOwner is the remote GitHub organization for the releases
var RepositoryOwner = "redhat-developer"

// RepositoryName is the remote GitHub repository for the releases
var RepositoryName = "app-services-cli"

func init() {
	if isDevBuild() {
		if info, ok := debug.ReadBuildInfo(); ok && info.Main.Version != "(devel)" {
			Version = info.Main.Version
		}
	}
}

// CheckForUpdate checks if there is a newer version of the CLI than
// the version currently being used. If so, it logs this information
// to the console.
func CheckForUpdate(ctx context.Context, logger logging.Logger) {
	latest, err := getLatestVersion(ctx)
	if err != nil {
		logger.Debug("Could not check latest version:", err)
		return
	}

	latestVersion := latest.TagName

	// if the user is using a dev or pre-release version, do not check for if an update is available
	if isDevBuild() || isPreRelease(Version) {
		return
	}

	if latestVersion != &Version {
		logger.Info()
		logger.Info(color.Info("A new version of rhoas is available:"), color.CodeSnippet(*latestVersion))
		logger.Info(color.Info(latest.GetHTMLURL()))
		logger.Info()
	}
}

// Get the latest version of the CLI
func getLatestVersion(ctx context.Context) (*github.RepositoryRelease, error) {
	client := github.NewClient(nil)

	latest, _, err := client.Repositories.GetLatestRelease(ctx, RepositoryOwner, RepositoryName)
	if err != nil {
		return nil, err
	}

	return latest, nil
}

// isDevBuild returns true if the current build is "dev" (dev build)
func isDevBuild() bool {
	return Version == "dev"
}

func isPreRelease(tag string) bool {
	match, _ := regexp.MatchString("^[0-9]+\\.[0-9]+\\.[0-9]+$", tag)
	return !match
}
