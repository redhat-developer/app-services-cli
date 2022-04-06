package main

import (
	"context"
	"fmt"
	"os"

	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/plugins/request/pkg/cmd/plugin"

	"github.com/redhat-developer/app-services-cli/pkg/core/localize/goi18n"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory/defaultfactory"

	"github.com/redhat-developer/app-services-cli/internal/build"
)

func main() {
	// TODO - wrap into SDK
	localizer, err := goi18n.New(nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	buildVersion := build.Version
	cmdFactory := defaultfactory.New(localizer)

	rootCmd := plugin.NewPluginCommand(cmdFactory, buildVersion)

	err = rootCmd.Execute()

	// TODO replicate all handling into SDK
	if err == nil {
		if flagutil.DebugEnabled() {
			build.CheckForUpdate(cmdFactory.Context, build.Version, cmdFactory.Logger, localizer)
		}
		return
	}

	cmdFactory.Logger.Errorf("%v\n", err)
	build.CheckForUpdate(context.Background(), build.Version, cmdFactory.Logger, localizer)
	os.Exit(1)
}
