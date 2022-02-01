## rhoas login

Log in to RHOAS

### Synopsis

Log in securely to RHOAS through a web browser.

This command opens your web browser, where you can enter your credentials.

When using RHOAS in an environment without a web browser, you can log in using an offline-token by using the "--token" flag, which can be obtained at https://console.redhat.com/openshift/token.

Note: Token-based login is not supported by the "rhoas kafka topic" and “rhoas kafka consumer-group" commands.


```
rhoas login [flags]
```

### Examples

```
# Securely log in to RHOAS by using a web browser
$ rhoas login

# Print the authentication URL instead of automatically opening a web browser
$ rhoas login --print-sso-url

# Log in using an offline token
$ rhoas login --token f5cgc...

```

### Options

```
      --api-gateway string   URL of the API gateway (default "https://api.openshift.com")
      --auth-url string      The URL of the SSO Authentication server (default "https://sso.redhat.com/auth/realms/redhat-external")
      --client-id string     OpenID client identifier (default "rhoas-cli-prod")
      --insecure             Allow insecure communication with the server by disabling TLS certificate and host name verification
      --print-sso-url        Print the console login URL, which you can use to log in to RHOAS from a different web browser (this is useful if you need to log in with different credentials than the credentials you used in your default web browser)
      --scope stringArray    Override the default OpenID scope (to specify multiple scopes, use a separate --scope for each scope) (default [openid])
  -t, --token string         Log in using an offline token, which can be obtained at https://console.redhat.com/openshift/token
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas](rhoas.md)	 - RHOAS CLI

