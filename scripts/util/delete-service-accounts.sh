echo "Deleting all service accounts"

serviceaccount_uuids=$(rhoas serviceaccount list -o json | jq -rc '.items[].id')

for id in ${serviceaccount_uuids}; do
  echo "Deleting service account '$id'..."
  rhoas serviceaccount delete -y --id $id
done