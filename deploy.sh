#!/usr/bin/env bash

go mod vendor
make tests
make lint
git add . && git commit -am '0.4.2' && git push
git tag v0.4.2 && git push --tags
