## rhoas connector update

Update a Connectors instance

### Synopsis

Update a Connectors instance

Allows you to change the details of a already existing connectors instance.
By changing its configuration in a text editor. To change which editor use
edit the EDITOR environment variable. Below are some comman editors you may
use.

export EDITOR=nvim
export EDITOR=vim
export EDITOR="code -w"


```
rhoas connector update [flags]
```

### Examples

```
# Update a Connectors instance
rhoas connector update --id=my-connector --file=myconnector.json

# Update a Connectors instance from stdin
cat myconnector.json | rhoas connector update

```

### Options

```
      --kafka-id string       ID of of the namespace you want the connector to be updated to
      --name string           Override name of the connector (by default name in the connector spec would be used)
      --namespace-id string   ID of of the kafka you want the connector to be updated to
  -o, --output string         Specify the output format. Choose from: "json", "yaml", "yml"
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas connector](rhoas_connector.md)	 - Connectors instance commands

