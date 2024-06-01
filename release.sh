#!/bin/env sh

git tag -a "${1:?version is as first argument is required}" &&
git push --tags
