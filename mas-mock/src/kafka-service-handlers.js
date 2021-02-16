module.exports = {
  createServiceAccount: async (c, req, res) => {
    const clientId = Number.MAX_SAFE_INTEGER - new Date().getTime();
    const clientSecret = Number.MAX_SAFE_INTEGER - new Date().getTime();
    res.status(200).json({
      name: req.body.name,
      description: req.body.description,
      clientID: clientId.toString(),
      clientSecret: clientSecret.toString(),
    });
  },
  createKafka: async (c, req, res) => {
    res.status(202).json({
      id: "1iSY6RQ3JKI8Q0OTmjQFd3ocFRg",
      kind: "kafka",
      href: "/api/managed-services-api/v1/kafkas/1iSY6RQ3JKI8Q0OTmjQFd3ocFRg",
      status: "complete",
      cloud_provider: "aws",
      multi_az: false,
      region: "us-east-1",
      owner: "api_kafka_service",
      name: "serviceapi",
      bootstrapServerHost:
        "serviceapi-1isy6rq3jki8q0otmjqfd3ocfrg.apps.ms-bttg0jn170hp.x5u8.s1.devshift.org",
      created_at: "2020-10-05T12:51:24.053142Z",
      updated_at: "2020-10-05T12:56:36.362208Z",
    });
  },

  deleteKafkaById: async (c, req, res) => {
    res.status(204).json({
      id: "1iSY6RQ3JKI8Q0OTmjQFd3ocFRg",
      kind: "kafka",
      href: "/api/managed-services-api/v1/kafkas/1iSY6RQ3JKI8Q0OTmjQFd3ocFRg",
      status: "complete",
      cloud_provider: "aws",
      multi_az: false,
      region: "us-east-1",
      owner: "api_kafka_service",
      name: "serviceapi",
      bootstrapServerHost:
        "serviceapi-1isy6rq3jki8q0otmjqfd3ocfrg.apps.ms-bttg0jn170hp.x5u8.s1.devshift.org",
      created_at: "2020-10-05T12:51:24.053142Z",
      updated_at: "2020-10-05T12:56:36.362208Z",
    });
  },

  getKafkaById: async (c, req, res) => {
    res.status(200).json({
      id: "1iSY6RQ3JKI8Q0OTmjQFd3ocFRg",
      kind: "kafka",
      href: "/api/managed-services-api/v1/kafkas/1iSY6RQ3JKI8Q0OTmjQFd3ocFRg",
      status: "complete",
      cloud_provider: "aws",
      multi_az: false,
      region: "us-east-1",
      owner: "api_kafka_service",
      name: "serviceapi",
      bootstrapServerHost:
        "serviceapi-1isy6rq3jki8q0otmjqfd3ocfrg.apps.ms-bttg0jn170hp.x5u8.s1.devshift.org",
      created_at: "2020-10-05T12:51:24.053142Z",
      updated_at: "2020-10-05T12:56:36.362208Z",
    });
  },

  listKafkas: async (c, req, res) => {
    const { status, mock } = c.api.mockResponseForOperation(c.operation.operationId);
    return res.status(status).json(mock);
  },

  listCloudProviders: async (_c, _req, res) => {
    return res.status(200).json({
      "kind": "CloudProviderList",
      "page": 1,
      "size": 7,
      "total": 7,
      "items": [
        {
          "kind": "CloudProvider",
          "id": "aws",
          "display_name": "Amazon Web Services",
          "name": "aws",
          "enabled": true
        },
        {
          "kind": "CloudProvider",
          "id": "azure",
          "display_name": "Microsoft Azure",
          "name": "azure",
          "enabled": false
        },
      ]
    })
  },

  listCloudProviderRegions: async (_c, _req, res) => {
    return res.status(200).json(
      {
        "kind": "CloudRegionList",
        "page": 1,
        "size": 17,
        "total": 17,
        "items": [
          {
            "kind": "CloudRegion",
            "id": "eu-west-2",
            "display_name": "EU, London",
            "enabled": true
          }
        ]
      })
  },

  // Handling auth
  notFound: async (c, req, res) => res.status(404).json({ err: "not found" }),
  unauthorizedHandler: async (c, req, res) =>
    res.status(401).json({ err: "unauthorized" }),
};
