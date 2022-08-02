## rhoas service-registry setting

Configure settings for a Service Registry instance

### Synopsis

Service Registry instance owners, instance administrators, and organization administrators can 
configure settings for a Service Registry instance. 

The available settings include the following options: 

* registry.auth.authenticated-read-access.enabled - Specifies whether Service Registry grants at least 
  read-only access to requests from any authenticated user in the same organization, regardless of their 
  user role. Defaults to false.  
* registry.auth.basic-auth-client-credentials.enabled - Specifies whether Service Registry users can 
  authenticate using HTTP basic authentication, in addition to OAuth. Defaults to true.
* registry.auth.owner-only-authorization - Specifies whether only the user who creates an artifact can 
  modify that artifact. Defaults to false. 
* registry.ccompat.legacy-id-mode.enabled - Specifies whether the Confluent Schema Registry compatibility 
  API uses globalId instead of contentId as an artifact identifier. Defaults to false.


### Examples

```
## List all settings for the current Service Registry instance
$ rhoas service-registry setting list

## Set the value of setting
$ rhoas service-registry setting set --name registry.ccompat.legacy-id-mode.enabled --value true

```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas service-registry](rhoas_service-registry.md)	 - Service Registry commands
* [rhoas service-registry setting get](rhoas_service-registry_setting_get.md)	 - Get value of the setting for a Service Registry instance
* [rhoas service-registry setting list](rhoas_service-registry_setting_list.md)	 - List settings for a Service Registry instance
* [rhoas service-registry setting set](rhoas_service-registry_setting_set.md)	 - Set value of the setting for a Service Registry instance

