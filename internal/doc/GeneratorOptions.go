package doc

import "github.com/spf13/cobra"

// GeneratorOptions all modifiers for the generator
type GeneratorOptions struct {
	// Directory to write the documentation files
	Dir string

	// FileNameGenerator - provides custom file name for each documentation file
	FileNameGenerator func(cmd *cobra.Command) string

	// FilePrepender - Prepend content to the generated file (add header)
	FilePrepender func(string) string

	// LinkHandler - function to handle links that lets you to transform them to different format
	LinkHandler func(string) string

	// GenerateIndex - generate index file
	GenerateIndex bool

	// IndexLocation - name and folder of the assembly file (typically ./README.adoc)
	IndexLocation string
}
