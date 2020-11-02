package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/golang/glog"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"gitlab.cee.redhat.com/mas-dx/rhmas/cmd/auth"
	"gitlab.cee.redhat.com/mas-dx/rhmas/cmd/kafka"
	"gitlab.cee.redhat.com/mas-dx/rhmas/cmd/login"
)

func init() {
}

func main() {
	// This is needed to make `glog` believe that the flags have already been parsed, otherwise
	// every log messages is prefixed by an error message stating the the flags haven't been
	// parsed.
	_ = flag.CommandLine.Parse([]string{})

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)

	// Always log to stderr by default
	if err := flag.Set("logtostderr", "true"); err != nil {
		fmt.Printf("Unable to set logtostderr to true")
	}

	rootCmd := &cobra.Command{
		Use:  "rhmas cli",
		Long: "rhmas:  Manage Red Hat Managed Services",
	}

	rootCmd.AddCommand(login.NewLoginCommand())
	rootCmd.AddCommand(kafka.NewKafkaCommand())
	rootCmd.AddCommand(auth.NewAuthGroupCommand())
	rootCmd.AddCommand(completionCmd)

	if err := rootCmd.Execute(); err != nil {
		glog.Fatalf("error running command: %v", err)
	}
}

var completionCmd = &cobra.Command{
	Use:   "completion [bash|zsh|fish|powershell]",
	Short: "Generate completion script",
	Long: `To load completions:

Bash:

$ source <(yourprogram completion bash)

# To load completions for each session, execute once:
Linux:
  $ rhmas completion bash > /etc/bash_completion.d/yourprogram
MacOS:
  $ rhmas completion bash > /usr/local/etc/bash_completion.d/yourprogram

Zsh:

# If shell completion is not already enabled in your environment you will need
# to enable it.  You can execute the following once:

$ echo "autoload -U compinit; compinit" >> ~/.zshrc

# To load completions for each session, execute once:
$ rhmas completion zsh > "${fpath[1]}/_rhmas"

# You will need to start a new shell for this setup to take effect.

Fish:

$ yourprogram completion fish | source

# To load completions for each session, execute once:
$ yourprogram completion fish > ~/.config/fish/completions/yourprogram.fish
`,
	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
	Args:                  cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "bash":
			cmd.Root().GenBashCompletion(os.Stdout)
		case "zsh":
			cmd.Root().GenZshCompletion(os.Stdout)
		case "fish":
			cmd.Root().GenFishCompletion(os.Stdout, true)
		case "powershell":
			cmd.Root().GenPowerShellCompletion(os.Stdout)
		}
	},
}
