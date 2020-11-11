package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra/doc"

	"github.com/bf2fc6cc711aee1a0c2a/cli/cmd/rhmas/auth"
	"github.com/bf2fc6cc711aee1a0c2a/cli/cmd/rhmas/completion"
	"github.com/bf2fc6cc711aee1a0c2a/cli/cmd/rhmas/docs"
	"github.com/bf2fc6cc711aee1a0c2a/cli/cmd/rhmas/kafka"
	"github.com/bf2fc6cc711aee1a0c2a/cli/cmd/rhmas/login"
	"github.com/bf2fc6cc711aee1a0c2a/cli/cmd/rhmas/logout"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/config"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:  "rhmas cli",
		Long: "rhmas:  Manage Red Hat Managed Services",
	}
	generateDocs = os.Getenv("GENERATE_DOCS") == "true"
)

func main() {
	cobra.OnInitialize(initConfig)

	rootCmd.AddCommand(login.NewLoginCommand())
	rootCmd.AddCommand(logout.NewLogoutCommand())
	rootCmd.AddCommand(kafka.NewKafkaCommand())
	rootCmd.AddCommand(auth.NewAuthGroupCommand())
	rootCmd.AddCommand(docs.NewDocsCommand())
	rootCmd.AddCommand(completion.CompletionCmd)

	rootCmd.Version = "0.1.0"

	if generateDocs {
		generateDocumentation(rootCmd)
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error running command: %v\n", err)
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

const fmTemplate = `---
id: %s
title: %s
---
`

/**
* Generates documentation files
 */
func generateDocumentation(rootCommand *cobra.Command) {
	fmt.Print("Generating docs. Config to put into sitebars\n\n")
	filePrepender := func(filename string) string {
		name := filepath.Base(filename)
		base := strings.TrimSuffix(name, path.Ext(name))
		fmt.Printf("\"commands/%s\",", base)
		finalName := strings.Replace(base, "_", " ", -1)
		return fmt.Sprintf(fmTemplate, base, finalName)
	}
	fmt.Print("\n")
	linkHandler := func(s string) string { return s }

	err := doc.GenMarkdownTreeCustom(rootCommand, "./docs/commands", filePrepender, linkHandler)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
