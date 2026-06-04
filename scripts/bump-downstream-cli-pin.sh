#!/usr/bin/env bash
# Bump pinned cli release in a downstream MCOS repository checkout.
#
# Usage: bump-downstream-cli-pin.sh <repo-name> <checkout-root> <version>
#   repo-name: plugin-cli | addons
#   version: GitHub release tag (optional mc_ prefix is stripped)

set -euo pipefail

REPO="${1:?repo name required}"
ROOT="${2:?checkout root required}"
VER="${3:?version required}"
VER="${VER#mc_}"

log() { echo "[bump-downstream-cli-pin] $*"; }

case "${REPO}" in
  plugin-cli)
    f="${ROOT}/Dockerfile"
    [[ -f "${f}" ]] || {
      log "missing ${f}"
      exit 1
    }
    sed -i -E "s/^ARG CLI_VERSION=.*/ARG CLI_VERSION=${VER}/" "${f}"
    log "updated Dockerfile"
    ;;
  addons)
    f="${ROOT}/ssh/build.yaml"
    [[ -f "${f}" ]] || {
      log "missing ${f}"
      exit 1
    }
    sed -i -E "s/^  CLI_VERSION: .*/  CLI_VERSION: ${VER}/" "${f}"
    log "updated ssh/build.yaml"
    ;;
  *)
    log "unknown repo: ${REPO}"
    exit 1
    ;;
esac
