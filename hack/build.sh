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

if [[ ! -d ./.git ]]; then
    echo "error: the build_deploy.sh script must be executed from the project root"
    exit 1
fi

midstream_target="gitlab.cee.redhat.com/mk-ci-cd/cli.git"
git_origin="https://app:${APP_SRE_BOT_PUSH_TOKEN}@${midstream_target}"

function exec_go() {
    ./hack/tools_exec.sh go $@
}

function exec_git() {
    ./hack/tools_exec.sh git $@
}

branch_name="update-vendor"
git checkout -B $branch_name

step "Display env vars"
exec_go env

step "Set GOMODCACHE"
export GOMODCACHE="${PWD}/gotmp"
mkdir -p "${GOMODCACHE}"
echo $GOMODCACHE
chown 1000 /var/lib/jenkins/workspace/mk-ci-cd-cli-gl-build-main/gotmp

step "Clean up code and dependencies"
exec_go mod tidy

step "Vendorize components"
exec_go mod vendor

step "Add vendor files to staging area"
git add vendor -f

step "Commit and push vendor files if there are changes to push"

num_changes=$(git diff --cached --numstat | wc -l)
if [[ $(($num_changes)) == 0 ]]; then
    exit 0
fi
