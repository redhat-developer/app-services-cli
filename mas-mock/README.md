## Mock for API.openshift.com

Using [openapi-backend](https://github.com/anttiviljami/openapi-backend) on [Express](https://expressjs.com/)

## QuickStart

```
yarn
yarn start # API running at http://localhost:9000
```

## Refresh API

Get recent version of the api

```
yarn refresh
```

> NOTE: This will remove operationId - make sure you add them back when pushing to repository

## Build container

docker build -t mas-mocking-api

 
 ## Calling API

```
yarn api
```
 
