#!/usr/bin/env bash

# Run pkger to embed static assets into the binary
pkger -o cmd/rhoas

# check if the pkged file has changes
# if so then throw an error as the user should ensure to embed changes into the binary
if [[ $(git diff --stat cmd/rhoas/pkged.go) != '' ]]; then
  echo "./cmd/rhoas/pkged.go has changes"
  echo "Please run `make pkger` to embed the static assets"
  exit 1
fi