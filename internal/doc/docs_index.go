package doc

import (
	"os"
	"path"
	"path/filepath"
	"text/template"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

const contentTemplate = `:context: rhoas-cli-command-reference
[id="cli-command-reference_{context}"]
= OpenShift Application Services CLI command reference

[role="_abstract"]
You use the ` + "`rhoas`" + ` CLI to manage your application services from the command line.

{{ range .Groups}}
== {{.Description}}
{{ range .Commands}}
include::{{.}}[leveloffset=+2]
{{ end }}
{{ end }}

`

type CommandGroup struct {
	Description string
	Commands    []string
}

func CreateIndexFile(rootCmd *cobra.Command, generationOptions *GeneratorOptions) error {
	var groups []CommandGroup = []CommandGroup{
		{
			Description: "General commands",
			Commands:    []string{},
		},
	}

	children := rootCmd.Commands()
	for _, child := range children {
		if !child.IsAvailableCommand() || child.IsAdditionalHelpTopicCommand() {
			continue
		}

		if child.Annotations[AnnotationName] != "" {
			groupNames := CollectNames(child, generationOptions)
			groups = append(groups, CommandGroup{
				Description: child.Annotations[AnnotationName],
				Commands:    groupNames,
			})
		} else {
			groupNames := CollectNames(child, generationOptions)
			// Add to general command
			groups[0].Commands = append(groups[0].Commands, groupNames...)
		}
	}

	type Vars struct {
		Groups []CommandGroup
	}

	vars := Vars{
		Groups: groups,
	}

	output, err := os.Create(generationOptions.IndexFile)
	if err != nil {
		return errors.WithStack(err)
	}

	err = template.Must(template.New("content").Parse(contentTemplate)).Execute(output, vars)
	if err != nil {
		return errors.WithStack(err)
	}

	err = output.Sync()
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func CollectNames(cmd *cobra.Command, options *GeneratorOptions) []string {
	filename := options.FileNameGenerator(cmd)

	filename, err := filepath.Rel(path.Dir(options.IndexFile), filename)
	if err != nil {
		panic(err)
	}
	names := []string{filename}
	if cmd.HasSubCommands() {
		for _, sub := range cmd.Commands() {
			if sub.IsAvailableCommand() && !sub.IsAdditionalHelpTopicCommand() {
				subNames := CollectNames(sub, options)
				names = append(names, subNames...)
			}
		}
	}

	return names
}
