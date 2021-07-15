#!/usr/bin/env bash

make tests
make lint
git add . && git commit -am '0.0.3' && git push
git tag v0.0.3 && git push --tags
