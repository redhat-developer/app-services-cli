if [ -z "$1" ]; then
  echo "You cannot delete all service accounts. Pass an argument to filter on the name"
  exit 1
fi

echo "Deleting service accounts which include '$1' in the name"

serviceaccounts=$(rhoas serviceaccount list -o json | jq -r -c --arg filter "$1" '.items[] | select(.name | contains($filter))')

for sa in ${serviceaccounts}; do
  name=$(echo $sa | jq -r '.name')
  echo "Deleting '$name'..."
  rhoas serviceaccount delete -f --id $(echo $sa | jq -r '.id')
done