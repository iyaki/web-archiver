#!/bin/env sh

git tag -a "${1:?version as first argument is required}" &&
git push --tags
