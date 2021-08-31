#!/usr/bin/env bash

set -uox pipefail
commands=$(git diff --stat docs/commands | wc -l)
if [[ $(($commands)) != 0 ]]; then
  echo "./docs/commands has changes"
  exit 1
fi