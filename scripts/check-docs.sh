#!/usr/bin/env bash

set -uo pipefail
commands=$(git diff --stat docs/commands | wc -l)
if [[ $(($commands)) != 0 ]]; then
  echo "./docs/commands directory has " $(($commands)) " changes"
  git diff --stat docs/commands
  exit 1
fi