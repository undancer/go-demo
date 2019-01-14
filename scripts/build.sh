#!/usr/bin/env bash

GOOS="${GOOS:-$(go env GOHOSTOS)}"
GOARCH="${GOARCH:-$(go env GOHOSTARCH)}"

export TARGET="build/app-$GOOS-$GOARCH"
export SOURCE="github.com/undancer/go-demo/cmd"

go build -o "${TARGET}" "${SOURCE}"
