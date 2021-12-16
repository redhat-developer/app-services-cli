package doc

// TODO remove
import (
	"os"
	"text/template"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

const contentTemplate = `:context: rhoas-cli-command-reference
[id="cli-command-reference_{context}"]
= CLI command reference (rhoas)

[role="_abstract"]
You use the ` + "`rhoas`" + ` CLI to manage your application services from the command line.

[#service-account-commands]
== Service account commands
{{ range .Commands}}
include::{{.}}[leveloffset=+1]
{{ end }}
`

func CreateIndexFile(rootCmd *cobra.Command, location string) error {

	commandFileNames := []string{"todo.adoc"}

	type Vars struct {
		Commands []string
	}

	vars := Vars{
		Commands: commandFileNames,
	}

	output, err := os.Create(location)
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
