
echo "Synchronizing managed-services-api"

wget --no-check-certificate https://gitlab.cee.redhat.com/service/managed-services-api/-/raw/master/openapi/kafka-service.yaml
mv kafka-service.yaml ./openapi/kafka-service.yaml

## Copy api to mock
cp ./openapi/kafka-service.yaml ./mas-mock/kafka-service.yaml

echo "Finished synchronization with managed-services-api"