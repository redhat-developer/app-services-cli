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

# branch_name="update-vendor"
# git checkout -B $branch_name

# step "Vendorize components"
# exec_go mod vendor

# step "Add vendor files to staging area"
# git add vendor -f

# step "Commit and push vendor files if there are changes to push"

# num_changes=$(git diff --cached --numstat | wc -l)
# if [[ $(($num_changes)) == 0 ]]; then
#     exit 0
# fi

# step "Commit vendor updates and create a merge request"
# git commit -m "chore: update vendor folder"
# exec_git push -u ${git_origin} ${branch_name} -o merge_request.create \
#     -o ci.skip \
#     -o merge_request.target="main" \
#     -o merge_request.merge_when_pipeline_succeeds
#     -o merge_request.remove_source_branch