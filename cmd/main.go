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

	rootCmd.AddCommand(login.NewLoginCommand())
	rootCmd.AddCommand(login.NewLogoutCommand())
	rootCmd.AddCommand(kafka.NewKafkaCommand())
	rootCmd.AddCommand(auth.NewAuthGroupCommand())
	rootCmd.AddCommand(tools.CompletionCmd)

	if err := rootCmd.Execute(); err != nil {
		glog.Fatalf("error running command: %v", err)
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
