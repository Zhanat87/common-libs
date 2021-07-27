#!/usr/bin/env bash

go mod vendor
make tests
make lint
git add . && git commit -am '0.3.4' && git push
git tag v0.3.4 && git push --tags
