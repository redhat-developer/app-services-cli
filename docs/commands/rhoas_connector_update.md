## rhoas connector update

Update a Connectors instance

### Synopsis

Update a Connectors instance.

You can change the details of an existing Connectors instance by editing its configuration file in a text editor. 

Note: To specify an editor, use the EDITOR environment variable before you run the "rhoas connector update" command. For example:
export EDITOR=nvim
export EDITOR=vim
export EDITOR="code -w"

After you edit the configuration file, use the "connector update" command to update an existing Connectors instance by specifying the configuration file.


```
rhoas connector update [flags]
```

### Examples

```
# Update name of the current Connectors instance
rhoas connector update --name=my-connector

# Update Kafka Instance of a Connectors instance by ID
rhoas connector update --kafka-id ce6pg07k09f3rs6us7sg --id ce6tgb1mk0orirpo5i70

```

### Options

```
      --id string             ID of the Connectors instance to be updated (the default is the instance in current context)
      --kafka-id string       ID of of the Kafka instance that you want the Connectors instance to use
      --name string           Override the name of the Connectors instance (the default name is the name specified in the connector configuration file)
      --namespace-id string   ID of the namespace in which you want to deploy the Connectors instance
  -o, --output string         Specify the output format. Choose from: "json", "yaml", "yml"
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas connector](rhoas_connector.md)	 - Connectors commands

