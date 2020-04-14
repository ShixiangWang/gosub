#!/usr/bin/env bash

version=$1
git tag -a "$version" -m "$version"
git push origin "$version"

if [ "$(uname)" == "Darwin" ];then
# MAC OS
goreleaser --rm-dist
elif [ "$(expr substr $(uname -s) 1 5)" == "Linux" ];then
# GNU/Linux
goreleaser --rm-dist
else
# Windows NT
/d/go/bin/goreleaser.exe --rm-dist
fi
