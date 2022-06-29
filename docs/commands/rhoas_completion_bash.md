## rhoas completion bash

Generate command completion script for Bash shell

### Synopsis

Install rhoas command completion for the Bash shell.

Installing on Linux:

  1. Create the script file:

     $ rhoas completion bash > rhoas_completions

  2. Place the script in a special Bash completions directory:

     $ sudo mv rhoas_completions /etc/bash_completion.d/rhoas

  3. Restart your shell for the changes to take effect.

Installing on macOS:

  1. Create the script file:

     $ rhoas completion bash > rhoas_completions

  2. Place the script in a special Bash completions directory:

     $ sudo mv rhoas_completions /usr/local/etc/bash_completion.d/rhoas

  3. Restart your shell for the changes to take effect.


```
rhoas completion bash
```

### Examples

```
rhoas completion bash

```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas completion](rhoas_completion.md)	 - Install command completion for your shell (bash, zsh, fish or powershell)

