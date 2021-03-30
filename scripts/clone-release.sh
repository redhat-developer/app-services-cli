#!/usr/bin/env bash

if ! command -v bf3-uploader &> /dev/null
then
  echo "Installing bf3-uploader"
  go get -u github.com/bf3fc6c/cli/cmd/bf3-uploader
fi

CLONE_FROM_ORG=bf2fc6cc711aee1a0c2a
CLONE_FROM_REPO=cli

# run cloner scrip
echo "Cloning release to github.com/bf3fc6c/cli"
bf3-uploader