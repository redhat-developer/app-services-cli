## rhoas completion zsh

Generate command completion script for Zsh shell

### Synopsis

Install rhoas command completion  for the Zsh shell.

1. Install the completion script:

   $ rhoas completion zsh > "${fpath[1]}/_rhoas"

2. Unless already installed, enable shell completions for Zsh:

   $ echo "autoload -U compinit; compinit" >> ~/.zshrc

3. Restart your shell for the changes to take effect.


```
rhoas completion zsh
```

### Examples

```
rhoas completion zsh

```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas completion](rhoas_completion.md)	 - Install command completion for your shell (bash, zsh, fish or powershell)

