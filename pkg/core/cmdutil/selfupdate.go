package cmdutil

import (
	"time"

	"github.com/blang/semver"
	"github.com/redhat-developer/app-services-cli/internal/build"
	"github.com/redhat-developer/app-services-cli/pkg/core/config"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
)

func DoSelfUpdate(logger logging.Logger) (bool, error) {
	version := build.Version
	if build.IsDevBuild() {
		// TODO uncomment later
		//return false, nil
		version = "0.0.0"
	}

	v := semver.MustParse(version)
	latest, err := selfupdate.UpdateSelf(v, build.RepositoryOwner+"/"+build.RepositoryName)
	if err != nil {
		return false, err
	}
	if latest.Version.Equals(v) {
		// latest version is the same as current version. It means current binary is up to date.
		logger.Debug("Current binary is the latest version", version)
	} else {
		logger.Info("Successfully updated RHOAS CLI to latest version", latest.Version)
		logger.Info("Release notes:\n", latest.ReleaseNotes)
		return true, err
	}

	return false, nil
}

func DoSelfUpdateOnceADay(logger logging.Logger, loader config.IConfig) (bool, error) {
	cfg, err := loader.Load()
	if err != nil {
		return false, err
	}
	logger.Debug("Last updated cli", cfg.LastUpdated)

	if cfg.LastUpdated < time.Now().AddDate(0, 0, -1).UnixMilli() {
		logger.Debug("Updating CLI")

		updated, err := DoSelfUpdate(logger)
		if err != nil {
			return false, err
		}
		cfg.LastUpdated = time.Now().UnixMilli()
		err = loader.Save(cfg)
		if err != nil {
			return false, err
		}
		return updated, nil
	}
	return false, nil
}
