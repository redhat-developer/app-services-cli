
echo "Downloading Connectors schemas into current folder"

for id in $(rhoas connector type list --limit=200 | jq -rc '.id'); do
  echo "fetching schema '$id'..."
  rhoas connector type describe --type $id | jq .schema > ${id}.json
done