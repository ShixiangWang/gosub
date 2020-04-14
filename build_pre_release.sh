#!/usr/bin/env bash

if [ "$(uname)" == "Darwin" ];then
# MAC OS
goreleaser --snapshot --skip-publish --rm-dist
elif [ "$(expr substr $(uname -s) 1 5)" == "Linux" ];then
# GNU/Linux
goreleaser --snapshot --skip-publish --rm-dist
else
# Windows NT
/d/go/bin/goreleaser.exe --snapshot --skip-publish --rm-dist
fi


