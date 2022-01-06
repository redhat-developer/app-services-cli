## rhoas service-registry role

Service Registry role management

### Synopsis


Service Registry role management
Set of commands that give users one of following permissions:

- READ_ONLY (read artifacts)
- DEVELOPER (write access to all resources)
- ADMIN (Export/Import artifacts, Manage Roles)

Roles can be applied to users (e.g martin_redhat) and Service Account Client IDs (e.g srvc-acct-03ddedba-5b49-4aa0-9b68-02e8b8c31add).
These commands are only accessible to users with the organization admin role or owners of the Service Registry instance.


### Examples

```
## Create or update user role
rhoas service-registry role add --role=ADMIN --username=joedough

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

