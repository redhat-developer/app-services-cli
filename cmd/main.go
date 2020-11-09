package main

import (
	"os"
	"fmt"

	"github.com/bf2fc6cc711aee1a0c2a/cli/cmd/auth"
	"github.com/bf2fc6cc711aee1a0c2a/cli/cmd/kafka"
	"github.com/bf2fc6cc711aee1a0c2a/cli/cmd/login"
	"github.com/bf2fc6cc711aee1a0c2a/cli/cmd/tools"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/config"
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
		fmt.Fprintf(os.Stderr, "error running command: %v\n", err)
	} else {
		if openHelp {
			fmt.Fprintln(os.Stderr, "Opening help in browser")
			cmd, err := tools.GetOpenBrowserCommand("http://localhost:3000/docs/commands/rhmas")
			if err != nil {
				fmt.Fprint(os.Stderr, err)
				os.Exit(1)
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
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	cfgFile = &config.Config{}
	if err := config.Save(cfgFile); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	cfgPath, _ := config.Location()
	fmt.Fprintf(os.Stderr, "Saved config to %v\n", cfgPath)
}
