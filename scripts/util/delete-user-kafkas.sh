# Script to delete a user's Kafka instances
# ./scripts/delete-user-kafkas.sh $username
if [ ! "$(command -v jq)" ]; then
  echo "Error: jq not installed"
  exit 1
fi

if [ -z "$1" ]; then
  echo "Missing username argument"
  exit 1
fi

user_kafkas=$(rhoas kafka list -o json | jq -c --arg username "$1" '.items[] | select(.owner == $username)')

for kafka in $user_kafkas; do
  rhoas kafka delete --id $(echo $kafka | jq -r '.id')
done