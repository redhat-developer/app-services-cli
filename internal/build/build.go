package build

import (
	"context"
	"runtime/debug"
	"strings"
	"time"

	"github.com/google/go-github/v39/github"
	"github.com/redhat-developer/app-services-cli/pkg/color"
	"github.com/redhat-developer/app-services-cli/pkg/localize"
	"github.com/redhat-developer/app-services-cli/pkg/logging"
)

type buildSource string

const (
	githubBuildSource buildSource = "github"
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

	// TermsReviewSpecURL Url used to download terms and conditions specification
	TermsReviewSpecURL = "https://console.redhat.com/apps/application-services/terms-conditions-spec.json"

	// DefaultPageSize is the default number of items per page when using list commands
	DefaultPageSize = "10"

	// DefaultPageNumber is the default page number when using list commands
	DefaultPageNumber = "1"

	// SSORedirectPath is the default SSO redirect path
	SSORedirectPath = "sso-redhat-callback"

	// MASSSORedirectPath is the default MAS-SSO redirect path
	MASSSORedirectPath = "mas-sso-callback"

	// BuildSource is a unique key which indicates the infrastructure on which the binary was built
	BuildSource = "local"
)

// Auth Build variables
var (
	ProductionAPIURL            = "https://api.openshift.com"
	StagingAPIURL               = "https://api.stage.openshift.com"
	ConsoleURL                  = "https://console.redhat.com"
	DefaultClientID             = "rhoas-cli-prod"
	DefaultUserAgentPrefix      = "rhoas-cli_"
	DefaultOfflineTokenClientID = "cloud-services"
	DefaultLoginTimeout         = 60 * time.Second
	OfflineTokenURL             = ConsoleURL + "/openshift/token"
	ProductionAuthURL           = "https://sso.redhat.com/auth/realms/redhat-external"
	ProductionMasAuthURL        = "https://identity.api.openshift.com/auth/realms/rhoas"
	StagingMasAuthURL           = "https://identity.api.stage.openshift.com/auth/realms/rhoas"
)

func init() {
	if IsDevBuild() {
		if info, ok := debug.ReadBuildInfo(); ok && info.Main.Version != "(devel)" {
			Version = info.Main.Version
		}
	}
}

// CheckForUpdate checks if there is a newer version of the CLI than
// the version currently being used. If so, it logs this information
// to the console.
func CheckForUpdate(ctx context.Context, version string, logger logging.Logger, localizer localize.Localizer) {
	if BuildSource != string(githubBuildSource) {
		return
	}
	// prefix version with a v to correspond with Git tag
	if !strings.HasPrefix(version, "v") {
		version = "v" + version
	}

	releases, err := getReleases(ctx)
	if err != nil {
		return
	}

	var latestRelease *github.RepositoryRelease
	releaseTagIndexMap := map[string]int{}
	for i, release := range releases {
		// assign the latest non-pre release as the latest public release
		if latestRelease == nil && !release.GetPrerelease() {
			latestRelease = release
		}

		// create an tag:index map of the releases
		// the first index (0) is the latest release
		releaseTagIndexMap[release.GetTagName()] = i
		if release.GetTagName() == version {
			break
		}
	}

	currentVersionIndex, ok := releaseTagIndexMap[version]
	if !ok {
		// the currently used version does not exist as a public release
		// assume it to be an unpublished or dev release
		return
	}

	latestVersionIndex := releaseTagIndexMap[latestRelease.GetTagName()]

	// if the index of the current version is greater than the latest release
	// this means it is older, and therefore, an update is available.
	if currentVersionIndex > latestVersionIndex {
		logger.Info()
		logger.Info(color.Info(localizer.MustLocalize("common.log.info.updateAvailable")), color.CodeSnippet(latestRelease.GetTagName()))
		logger.Info(color.Info(latestRelease.GetHTMLURL()))
		logger.Info()
	}
}

func getReleases(ctx context.Context) ([]*github.RepositoryRelease, error) {
	client := github.NewClient(nil)

	releases, _, err := client.Repositories.ListReleases(ctx, RepositoryOwner, RepositoryName, nil)
	if err != nil {
		return nil, err
	}

	return releases, nil
}

// IsDevBuild returns true if the current build is "dev" (dev build)
func IsDevBuild() bool {
	return Version == "dev"
}
