package tools

import (
	"os"
	"fmt"
	"path"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

const fmTemplate = `---
id: %s
title: %s
---

`

/**
* Generates documentation files
 */
func DocumentationGenerator(rootCommand *cobra.Command) {
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
