
## Architecture

MAS CLI will allow developers to mitigate access to a number of servers:

- cloud.redhat.com (managed service API)
- managed service keycloak (gateway for services like OpenShift Streaming)
- custom kafka protocol for admin tasks like topic creation

To do that effectively our team will build client that could be reused across ecosystem.
To showcase value of client we want to use it in CLI and Operator like in diagram:

## Operator

POC operator will define custom resource that users can use to interact with the MAS API.