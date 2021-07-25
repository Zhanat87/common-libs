#!/usr/bin/env bash

go mod vendor
make tests
make lint
git add . && git commit -am '0.2.6' && git push
git tag v0.2.6 && git push --tags
