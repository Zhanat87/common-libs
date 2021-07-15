#!/usr/bin/env bash

go mod vendor
make tests
make lint
git add . && git commit -am '0.0.7' && git push
git tag v0.0.7 && git push --tags