package completion

import (
	"fmt"
	"os"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/factory"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/iostreams"
	"github.com/spf13/cobra"
)

type Options struct {
	IO *iostreams.IOStreams
}

func NewCompletionCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		IO: f.IOStreams,
	}

	cmd := &cobra.Command{
		Use:   "completion [bash|zsh|fish|powershell]",
		Short: "Generate completion script",
		Long: `To load completions:
	
	Bash:
	
	$ source <(rhoas completion bash)
	
	# To load completions for each session, execute once:
	Linux:
		$ rhoas completion bash > /etc/bash_completion.d/rhoas
	MacOS:
		$ rhoas completion bash > /usr/local/etc/bash_completion.d/rhoas
	
	Zsh:
	
	# If shell completion is not already enabled in your environment you will need
	# to enable it.  You can execute the following once:
	
	$ echo "autoload -U compinit; compinit" >> ~/.zshrc
	
	# To load completions for each session, execute once:
	$ rhoas completion zsh > "${fpath[1]}/_rhoas"
	
	# You will need to start a new shell for this setup to take effect.
	
	Fish:
	
	$ rhoas completion fish | source
	
	# To load completions for each session, execute once:
	$ rhoas completion fish > ~/.config/fish/completions/rhoas.fish
	`,
		DisableFlagsInUseLine: true,
		ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
		Args:                  cobra.ExactValidArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			var err error
			switch args[0] {
			case "bash":
				err = cmd.Root().GenBashCompletion(opts.IO.Out)
			case "zsh":
				err = cmd.Root().GenZshCompletion(opts.IO.Out)
			case "fish":
				err = cmd.Root().GenFishCompletion(opts.IO.Out, true)
			case "powershell":
				err = cmd.Root().GenPowerShellCompletion(opts.IO.Out)
			}

			if err != nil {
				fmt.Fprintln(opts.IO.ErrOut, err)
				os.Exit(1)
			}
		},
	}

	return cmd
}
