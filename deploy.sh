#!/usr/bin/env bash

make vendor
make tests
make lint
git add . && git commit -am '0.0.6' && git push
git tag v0.0.6 && git push --tags
