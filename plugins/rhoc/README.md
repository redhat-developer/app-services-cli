## RHOAS CLI example plugin

### Plugin specific information

Plugin contains single request command that allows developers to execute any call against control plane.
Command is hidden by default and embedded into rhoas cli:

- At runtime: Command is included into RHOAS CLI in the build time
- Dynamically: When new version of the plugin is released 

### Running plugin as standalone cli

Run following commands to compile plugin locally.

```
make install
rhrequest 
```

## How it works

// TODO add SDK information

// TODO add info how to host plugin in separate repository

// TODO add information how to publish new version of plugin
