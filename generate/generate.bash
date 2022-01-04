#!/usr/bin/env bash

set -eo pipefail

cd "$(dirname "$0")"

function die() {
    echo 1>&2 "$@"
    exit 1
}

if [ -z "$GEN_PATH" ]; then
    die "GEN_PATH must be set, e.g. /path/to/sajari/sdk-go/internal/openapi"
fi

docker-entrypoint.sh generate \
    -i /openapi.json \
    -g go \
    --git-user-id sajari \
    --git-repo-id sdk-go \
    --additional-properties enumClassPrefix=true \
    -o "$GEN_PATH"
