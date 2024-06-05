#!/bin/env bash

echo "Last version: $(git tag -l | head -1)"

read -rp "Enter new version: " version

git tag -a "${version}" -m "" &&
git push --tags
