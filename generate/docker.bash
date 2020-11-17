#!/usr/bin/env bash

set -eo pipefail

cd "$(dirname "$0")"

function die() {
  echo 1>&2 $*
  exit 1
}

GEN_PATH="$(pwd)/../internal/openapi"

OPENAPI_URL=${OPENAPI_URL:-https://api-gateway.sajari.com/v4/openapi.json}

if [ -z "$OPENAPI_URL" ]; then
    die "OPENAPI_URL must be set (e.g. https://api-gateway.sajari.com/v4/openapi.json)"
fi

rm -rf $GEN_PATH
mkdir -p $GEN_PATH
cp .openapi-generator-ignore $GEN_PATH/

wget -O openapi.json $OPENAPI_URL

docker build -f Dockerfile.generate -t sdk-go-generate .
docker run --rm -it \
    -v $(pwd)/openapi.json:/openapi.json \
    -v "$GEN_PATH":/gen \
    -v $(pwd)/generate.bash:/generate.bash \
    -e GEN_PATH=/gen \
    sdk-go-generate \
    ./generate.bash

docker build -f Dockerfile.post-generate -t sdk-go-post-generate .

docker run --rm -it \
    -v "$GEN_PATH":/app/gen \
    -v $(pwd)/post-generate.bash:/app/post-generate.bash \
    -e GEN_PATH=/app/gen \
    sdk-go-post-generate \
    ./post-generate.bash
