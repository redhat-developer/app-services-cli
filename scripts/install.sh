#!/usr/bin/env bash

{ # this ensures the entire script is downloaded #

  BINARY_NAME="rhoas"
  SRC_ORG="redhat-developer"
  SRC_REPO="app-services-cli"
  OS_TYPE="linux"
  OS_LONG_BIT="64"
  RELEASE_TAG="${1:-latest}"
  API_BASE_URL="https://api.github.com"
  API_RELEASES_BASE_URL="${API_BASE_URL}/repos/${SRC_ORG}/${SRC_REPO}/releases"
  DOWNLOAD_DIR="/tmp"
  BINARY_DEST="$HOME/.local/bin"

  API_RELEASE_URL="$API_RELEASES_BASE_URL/latest"
  if [ "$RELEASE_TAG" != "latest" ]; then
    API_RELEASE_URL="$API_RELEASES_BASE_URL/tags/$RELEASE_TAG"
  fi

  has_in_path() {
    type "$1" >/dev/null 2>&1
  }

  source_file() {
    # shellcheck disable=SC1090
    source "$1" 2>/dev/null
  }

  # update config based on OS type
  if [[ "$OSTYPE" == "darwin"* ]]; then
    OS_TYPE="macOS"
    BINARY_DEST="$HOME/bin"
  fi

  if ! has_in_path "curl"; then
    echo "curl is required to download this binary"
    exit 1
  fi

  if [ ! -d "$BINARY_DEST" ]; then
    mkdir -p "$BINARY_DEST"
  fi

  DOWNLOAD_TAG=$(curl -s "${API_RELEASE_URL}" |
    grep "tag_name.*" |
    cut -d '"' -f 4)

  if [ -z "$DOWNLOAD_TAG" ]; then
    API_RELEASE_URL="$API_RELEASES_BASE_URL/tags/v$RELEASE_TAG"

    DOWNLOAD_TAG=$(curl -s "${API_RELEASE_URL}" |
    grep "tag_name.*" |
    cut -d '"' -f 4)

    if [ -z "$DOWNLOAD_TAG" ]; then
      echo "Release tag $RELEASE_TAG not found"
      exit 1
    fi

    DOWNLOAD_TAG="$RELEASE_TAG"
  fi

  if [[ $DOWNLOAD_TAG == v* ]]; then
    DOWNLOAD_TAG="${DOWNLOAD_TAG:1}"
  fi


  ASSET_NAME="${BINARY_NAME}_${DOWNLOAD_TAG}_${OS_TYPE}_amd${OS_LONG_BIT}"
  ASSET_NAME_COMPRESSED="${ASSET_NAME}.tar.gz"

  DOWNLOAD_URL=$(curl -s "${API_RELEASE_URL}" |
    grep "browser.download_url.*${ASSET_NAME_COMPRESSED}" |
    cut -d '"' -f 4)

  cd "$DOWNLOAD_DIR" || exit

  echo "Downloading $BINARY_NAME v${DOWNLOAD_TAG}"
  if curl -sL "$DOWNLOAD_URL" --output "${ASSET_NAME_COMPRESSED}"; then
    echo "$BINARY_NAME v${DOWNLOAD_TAG} downloaded"
  else
    echo "Error downloading $BINARY_NAME ${DOWNLOAD_TAG}"
  fi

  # unpack and place the binary in the users PATH
  tar xf "$ASSET_NAME_COMPRESSED"
  if [ "$?" -ne 0 ]; then
    exit 1
  fi

  cp "${ASSET_NAME}/bin/${BINARY_NAME}" "${BINARY_DEST}/${BINARY_NAME}"
  if [ "$?" -ne 0 ]; then
    exit 1
  fi

  echo "rhoas has been installed succesfully to $BINARY_DEST"
  echo "Please ensure that $BINARY_DEST is on your PATH"
} # this ensures the entire script is downloaded #
