const OpenAPIBackend = require("openapi-backend").default;
const express = require("express");
const kafkaHandlers = require("./handlers/kas-fleet-manager");
const srsHandlers = require("./handlers/srs-fleet-manager");
const srsDataHandlers = require("./handlers/srs-data");
const topicHandlers = require("./handlers/kafka-admin");
const path = require("path");
var cors = require('cors');

const api = express();
api.use(express.json());
api.use(cors())

// define api
const kafkaAPI = new OpenAPIBackend({
  definition: path.join(__dirname, "../../openapi/kafka-service.yaml"),
});
const topicAPI = new OpenAPIBackend({
  definition: path.join(__dirname, "../../openapi/strimzi-admin.yaml"),
});

// TODO all API definitions should be done in yaml 
const srsControlApi = new OpenAPIBackend({
  definition: path.join(__dirname, "../../pkg/api/srs/client/api/openapi.yaml"),
});

const srsDataApi = new OpenAPIBackend({
  definition: path.join(__dirname, "../../pkg/api/srsdata/client/api/openapi.yaml"),
});

// register handlers
kafkaAPI.register(kafkaHandlers);
topicAPI.register(topicHandlers);
srsControlApi.register(srsHandlers);
srsDataApi.register(srsDataHandlers);

// register security handler
kafkaAPI.registerSecurityHandler("Bearer", (c, req, res) => {
  return true;
});

srsControlApi.registerSecurityHandler("Bearer", (c, req, res) => {
  return true;
});

// Skipping validation of the schema
// validation fails on this schema definition
// even though it is valid through other validation forms like Swagger.io
topicAPI.validateDefinition = () => {};

kafkaAPI.init();
topicAPI.init();
srsControlApi.init();
srsDataApi.init();

api.use((req, res) => {
  console.debug("URL", req.url);
  if (req.url.startsWith("/api/service-registry/v2")) {
    console.debug("Calling serviceregistry")
    return srsDataApi.handleRequest(req, req, res);
  }
  
  if (req.url.startsWith("/api/serviceregistry_mgmt/v1/registries")) {
    console.debug("Calling serviceregistry_mgmt")
    return srsControlApi.handleRequest(req, req, res);
  }

  if (req.url.startsWith("/api/managed-services-api/v1/")) {
    console.debug("Calling Kafka")
    return kafkaAPI.handleRequest(req, req, res);
  }

  if (req.url.startsWith("/rest")) {
    console.debug("Calling Kafka Admin")
    req.url = req.url.replace("/rest", "");
    return topicAPI.handleRequest(req, req, res);
  }
  res.status(405).status({ err: "Method not allowed" });
});

api.listen(8000, () =>
  console.info("Kafka Service API listening at http://localhost:8000")
);
