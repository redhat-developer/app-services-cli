#!/usr/bin/env bash
#
# Copyright RedHat.
# License: MIT License see the file LICENSE

# Inspired by https://disconnected.systems/blog/another-bash-strict-mode/
set -eEu -o pipefail
trap 's=$?; echo "[ERROR] [$(date +"%T")] on $0:$LINENO"; exit $s' ERR

function log() {
    echo "[$1] [$(date +"%T")] - ${2}"
}

function step() {
    log "STEP" "$1"
}

CONTAINER_ENGINE=${CONTAINER_ENGINE:-"docker"}
MK_TOOLS_IMAGE=${MK_TOOLS_IMAGE:-"quay.io/app-sre/mk-ci-tools:latest"}

${CONTAINER_ENGINE} run \
    -v $(pwd):/workspace -u $UID \
    -w /workspace -e GOMODCACHE=/var/lib/jenkins/workspace/mk-ci-cd-cli-gl-build-main/gotmp \
    ${MK_TOOLS_IMAGE} \
    $@
