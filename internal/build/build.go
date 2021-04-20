package build

import (
	"context"
	"regexp"
	"runtime/debug"

	"github.com/google/go-github/github"
	"github.com/redhat-developer/app-services-cli/pkg/color"
	"github.com/redhat-developer/app-services-cli/pkg/logging"
)

// Define public variables here which you wish to be configurable at build time
var (
	// Version is dynamically set by the toolchain or overridden by the Makefile.
	Version = "dev"

	// Language used, can be overridden by Makefile or CI
	Language = "en"

	// RepositoryOwner is the remote GitHub organization for the releases
	RepositoryOwner = "redhat-developer"

	// RepositoryName is the remote GitHub repository for the releases
	RepositoryName = "app-services-cli"

	// TermsReviewEventCode is the event code used when checking the terms review
	TermsReviewEventCode = "onlineService"

	// TermsReviewSiteCode is the site code used when checking the terms review
	TermsReviewSiteCode = "ocm"
)

// Auth Build variables
var (
	ProductionAPIURL            = "https://api.openshift.com"
	StagingAPIURL               = "https://api.stage.openshift.com"
	DefaultClientID             = "rhoas-cli-prod"
	DefaultOfflineTokenClientID = "cloud-services"
	ProductionAuthURL           = "https://sso.redhat.com/auth/realms/redhat-external"
	ProductionMasAuthURL        = "https://identity.api.openshift.com/auth/realms/rhoas"
	StagingMasAuthURL           = "https://identity.api.stage.openshift.com/auth/realms/rhoas"
)

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

// check if the tag is a pre-release tag
// true it if contains anything other than MAJOR.MINOR.PATCH
func isPreRelease(tag string) bool {
	match, _ := regexp.MatchString("^[0-9]+\\.[0-9]+\\.[0-9]+$", tag)
	return !match
}
