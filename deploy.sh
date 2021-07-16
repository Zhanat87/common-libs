#!/usr/bin/env bash

go mod vendor
make tests
make lint
git add . && git commit -am '0.1.0' && git push
git tag v0.1.0 && git push --tags
