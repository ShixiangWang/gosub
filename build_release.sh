#!/usr/bin/env bash

version=$1
git tag -a "$version" -m "$version"
git push origin "$version"
/d/go/bin/goreleaser.exe
