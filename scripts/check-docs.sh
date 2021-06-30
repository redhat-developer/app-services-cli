#!/usr/bin/env bash

if [[ $(git diff --stat docs/commands | wc -w | tr -d ' ') != 0 ]]; then
  echo "./docs/commands has changes"
  exit 1
fi