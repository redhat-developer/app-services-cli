package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/root"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/config"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/version"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

var (
	generateDocs = os.Getenv("GENERATE_DOCS") == "true"
)

func main() {
	cobra.OnInitialize(initConfig)
	rootCmd := root.NewCmdRoot()
	rootCmd.Version = version.CLI_VERSION

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error running command: %v\n", err)
	}

	if generateDocs {
		generateDocumentation(rootCmd)
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
	fmt.Fprint(os.Stderr, "Generating docs. Config to put into sidebars.json: \n\n")
	filePrepender := func(filename string) string {
		name := filepath.Base(filename)
		base := strings.TrimSuffix(name, path.Ext(name))
		fmt.Printf("\"commands/%s\",", base)
		finalName := strings.ReplaceAll(base, "_", " ")
		return fmt.Sprintf(fmTemplate, base, finalName)
	}

	linkHandler := func(s string) string { return s }

	err := doc.GenMarkdownTreeCustom(rootCommand, "./docs/commands", filePrepender, linkHandler)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
