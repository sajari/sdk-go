#!/usr/bin/env bash

set -eo pipefail

cd "$(dirname "$0")"

function die() {
  echo 1>&2 $*
  exit 1
}

if [ -z "$GEN_PATH" ]; then
  die "must set GEN_PATH (path to generated directory)"
fi

goimports -w $GEN_PATH
