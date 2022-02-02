## rhoas service-registry role

Service Registry role management

### Synopsis


Manage Service Registry roles using a set of commands that give users one of following permissions:

* Viewer (provides read access)
* Manager (provides read and write access)
* Admin (enables admin along with read and write access)

Roles can be applied to users (for example, "martin_redhat") and Service Account Client IDs (for example, "srvc-acct-03ddedba-5b49-4aa0-9b68-02e8b8c31add").
These commands are accessible only to users with the organization admin role or owners of the Service Registry instance.


### Examples

```
## Create or update user role
rhoas service-registry role add --role=Admin --username=joedough

## List user and service account roles
rhoas service-registry role list

## Revoke role for user
rhoas service-registry role revoke --username=janedough

```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas service-registry](rhoas_service-registry.md)	 - Service Registry commands
* [rhoas service-registry role add](rhoas_service-registry_role_add.md)	 - Add or update principal role
* [rhoas service-registry role list](rhoas_service-registry_role_list.md)	 - List roles
* [rhoas service-registry role revoke](rhoas_service-registry_role_revoke.md)	 - Revoke role for principal

