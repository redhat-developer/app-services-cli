
echo "Synchronizing managed-services-api"

wget --no-check-certificate https://gitlab.cee.redhat.com/service/managed-services-api/-/raw/master/openapi/managed-services-api.yaml
mv managed-services-api.yaml ./openapi/managed-services-api.yaml

## Copy api to mock
cp ./openapi/managed-services-api.yaml ./mas-mock/managed-services-api.yaml

echo "Finished synchronization with managed-services-api"