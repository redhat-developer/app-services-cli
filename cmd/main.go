package main

import (
	"flag"
	"fmt"

	"github.com/bf2fc6cc711aee1a0c2a/cli/cmd/auth"
	"github.com/bf2fc6cc711aee1a0c2a/cli/cmd/kafka"
	"github.com/bf2fc6cc711aee1a0c2a/cli/cmd/login"
	"github.com/bf2fc6cc711aee1a0c2a/cli/cmd/tools"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/config"
	"github.com/golang/glog"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:  "rhmas cli",
		Long: "rhmas:  Manage Red Hat Managed Services",
	}
	openHelp = false
)

func main() {
	cobra.OnInitialize(initConfig)

	// This is needed to make `glog` believe that the flags have already been parsed, otherwise
	// every log messages is prefixed by an error message stating the the flags haven't been
	// parsed.
	_ = flag.CommandLine.Parse([]string{})

	//pflag.CommandLine.AddGoFlagSet(flag.CommandLine)

	// Always log to stderr by default
	if err := flag.Set("logtostderr", "true"); err != nil {
		fmt.Printf("Unable to set logtostderr to true")
	}
	rootCmd.PersistentFlags().BoolVarP(&openHelp, "help-browser", "b", false, "help in browser")
	rootCmd.AddCommand(login.NewLoginCommand())
	rootCmd.AddCommand(login.NewLogoutCommand())
	rootCmd.AddCommand(kafka.NewKafkaCommand())
	rootCmd.AddCommand(auth.NewAuthGroupCommand())
	rootCmd.AddCommand(tools.CompletionCmd)

	rootCmd.Version = "0.1.0"

	// Uncomment this to generate docs.
	// tools.DocumentationGenerator(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		glog.Fatalf("error running command: %v", err)
	} else {
		if openHelp {
			fmt.Println("Opening help in browser")
			cmd, err := tools.GetOpenBrowserCommand("http://localhost:3000/docs/commands/rhmas")
			if err != nil {
				glog.Fatal(err)
			} else {
				cmd.Start()
			}
		}
	}
}

func initConfig() {
	cfgFile, err := config.Load()
	if cfgFile != nil {
		return
	}
	if err != nil {
		glog.Fatal(err)
	}
	cfgFile = &config.Config{}
	if err := config.Save(cfgFile); err != nil {
		glog.Fatal(err)
	}
	cfgPath, _ := config.Location()
	glog.Infof("Saved config to %v", cfgPath)
}
