package cmdutil

import (
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/blang/semver"
	"github.com/redhat-developer/app-services-cli/internal/build"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
)

// DoSelfUpdate checks for updates and prompts the user to update if there is a newer version available
func DoSelfUpdate(f *factory.Factory) (bool, error) {
	version := build.Version
	// TODO temp
	if build.IsDevBuild() {
		version = "0.0.0"
	}

	v := semver.MustParse(version)
	versionToUpdate, found, err := selfupdate.DefaultUpdater().DetectLatest(version)

	if found && versionToUpdate.Version.Equals(v) {
		// latest version is the same as current version. It means current binary is up to date.
		f.Logger.Debug("Current binary is the latest version", version)
		return false, nil
	}

	promptConfirmName := &survey.Confirm{
		Message: f.Localizer.MustLocalize("common.selfupdate.confirm"),
	}

	var confirmUpdate bool
	err = survey.AskOne(promptConfirmName, &confirmUpdate)
	if err != nil {
		return false, err
	}

	if !confirmUpdate {
		return false, nil
	}

	latest, err := selfupdate.UpdateSelf(v, build.RepositoryOwner+"/"+build.RepositoryName)
	if err != nil {
		return false, err
	}

	f.Logger.Info(f.Localizer.MustLocalize("common.selfupdate.success", localize.NewEntry("Version", latest.Version)))
	return true, err

}

func DoSelfUpdateOnceADay(f *factory.Factory) (bool, error) {
	if !f.IOStreams.CanPrompt() {
		// Do not prompt if we are not in interactive mode
		return false, nil
	}

	if build.IsDevBuild() {
		return false, nil
	}

	cfg, err := f.Config.Load()

	logger := f.Logger

	if err != nil {
		return false, err
	}
	logger.Debug("Checking for updates. Last check was done:", cfg.LastUpdated)

	if cfg.LastUpdated < time.Now().AddDate(0, 0, -1).UnixMilli() {
		updated, err := DoSelfUpdate(f)
		if err != nil {
			return false, err
		}
		cfg.LastUpdated = time.Now().UnixMilli()
		err = f.Config.Save(cfg)
		if err != nil {
			return false, err
		}
		return updated, nil
	}
	return false, nil
}
