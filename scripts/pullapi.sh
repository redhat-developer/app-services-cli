
echo "Synchronizing managed-services-api"

wget --no-check-certificate https://gitlab.cee.redhat.com/service/managed-services-api/-/raw/master/openapi/kafka-service.yaml
mv kafka-service.yaml ./openapi/kafka-service.yaml

wget https://raw.githubusercontent.com/Apicurio/apicurio-registry/2.0.0.Final/app/src/main/resources-unfiltered/META-INF/resources/api-specifications/registry/v2/openapi.json
mv openapi.json ./openapi/srs-service.json

wget https://github.com/bf2fc6cc711aee1a0c2a/srs-fleet-manager/blob/main/core/src/main/resources/srs-fleet-manager.json
mv srs-fleet-manager.json ./openapi/srs-fleet-manager.json

## Copy api to mock
cp ./openapi/kafka-service.yaml ./mas-mock/kafka-service.yaml

echo "Finished synchronization with managed-services-api"