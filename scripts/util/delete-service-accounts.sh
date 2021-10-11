#!/usr/bin/env bash

set -uo pipefail

echo "Deleting all service accounts"

serviceaccount_uuids=$(rhoas service-account list -o json | jq --arg iamwho $(rhoas whoami) -rc '.items[] | select( .owner == $iamwho)' | jq -rc '.id')

for id in ${serviceaccount_uuids}; do
  echo "Deleting service account '$id'..."
  rhoas service-account delete -y --id $id
done