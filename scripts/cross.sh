#!/usr/bin/env bash

BUILDDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

declare -A map=(
  ["linux"]="386 amd64 arm"
  ["darwin"]="386 amd64"
  ["freebsd"]="386 amd64"
  ["windows"]="386 amd64"
)

for key in ${!map[@]}
do
  for value in ${map[$key]}
  do
    export GOOS=${key}
    export GOARCH=${value}
    ${BUILDDIR}/build.sh
  done
done