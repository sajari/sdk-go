#!/usr/bin/env bash

set -eo pipefail

cd "$(dirname "$0")"

function die() {
    echo 1>&2 $*
    exit 1
}

if [ -z "$GEN_PATH" ]; then
    die "GEN_PATH must be set, e.g. /path/to/sajari/sdk-go/internal/openapi"
fi

# There's a bug in the generator which turns an object into the wrong Go map.
# It's probably to do with { "additionalProperties": { "type": "object" } } but
# either way we need to fix it in post.
sed -i 's/map\[string\]map\[string\]interface{}/map\[string\]interface{}/g' "$GEN_PATH"/model_send_event_request.go

goimports -w "$GEN_PATH"
