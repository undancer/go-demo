#!/usr/bin/env bash

echo build

export TARGET="build/app"
export SOURCE="github.com/undancer/go-demo/cmd"

go build -o "${TARGET}" "${SOURCE}"
