## rhoas completion powershell

Generate command completion script for Powershell shell

### Synopsis

Install rhoas command completion for the Powershell shell.

1. Create the script file:

   PS> rhoas completion powershell | Out-File -FilePath rhoas.ps1

2. Source this file from your PowerShell profile.

3. Restart your shell for the changes to take effect.


```
rhoas completion powershell
```

### Examples

```
rhoas completion powershell

```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas completion](rhoas_completion.md)	 - Install command completion for your shell (bash, zsh, fish or powershell)

