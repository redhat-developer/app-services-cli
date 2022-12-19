#!/usr/bin/env bash

set -uo pipefail

go mod vendor

commands=$(git diff --stat vendor | wc -l)

echo $((commands))
if [[ $(($commands)) != 0 ]]; then  
  echo echo "inconsistent vendoring detected"
  exit 1
fi