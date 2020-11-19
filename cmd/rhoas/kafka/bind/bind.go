package bind

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var dryrun bool

// NewGetCommand gets a new command for getting kafkas.
func NewCredentialsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bind",
		Short: "Bind currently selected kafa cluster credentials to your cluster",
		Long: `Bind command will use current kubernetes or openshift context (namespace/project you have selected)
		Bind command will retrieve credentials for your kafka and mount them as secret into your project.
		Additionally we going to provide extra metadata for the service binding operator:

		https://github.com/redhat-developer/service-binding-operator

		If your cluster has service binding operator installed, you would be able to bind your application with credentials directly from the console etc.

		To use command your cluster needs to have rhoas-operator installed:

		<Link to operator here>
		
		`,
		Run: runBind,
	}

	cmd.Flags().BoolVarP(&dryrun, "dryrun", "d", false, "Provide yaml file containing changes without applying them to the cluster. Developers can use `oc apply -f kafka.yml` to apply it manually")
	return cmd
}

func runBind(cmd *cobra.Command, _ []string) {
	if dryrun {
		fmt.Fprintf(os.Stderr, "Generating CR files")
	}
	fmt.Fprintf(os.Stderr, "Successfully bound kafka credentials to your cluster")
}
