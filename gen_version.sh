#!/bin/bash

git describe > version_tag_and_commit.txt
git log -1 --format="%at" | xargs -I '{}' date -d '@{}' '+%Y-%m-%d' > version_date.txt
printf "%s (%s)" "$(cat version_tag_and_commit.txt)" "$(cat version_date.txt)" > version.txt
