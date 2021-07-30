#!/usr/bin/env bash

go mod vendor
make tests
make lint
git add . && git commit -am '0.4.1' && git push
git tag v0.4.1 && git push --tags
