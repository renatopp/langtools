#!/bin/bash

TAG=$1

if [ -z "$TAG" ]
then
  echo "ERROR: no version provided. Usage: .tools/publish.sh <tag>"
  exit 0
fi

if ! [[ $TAG =~ ^v[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
	echo "ERROR: version does not follow semantic versioning starting with 'v'. Example: v1.0.0"
	exit 0
fi

# Prompt user for description using Vim
touch .publish-description
echo -e "# Please provide a description for release $TAG.\n# Lines starting with '#' like this one will be ignored.\n# Save the message with :wq.\n\n" > .publish-description
vim -c "startinsert | normal G" .publish-description

description=$(cat .publish-description | grep -v "^#")

if [ -z "$description" ]; then
	echo "Description is empty, exiting..."
	exit 0
else
	echo "Variable is not empty"
	echo $description
fi

# Insert the new description
current_date=$(date +'%Y-%m-%d')
description=$(echo -e "<!-- NEWER -->\n\n## $TAG ($current_date)\n$description\n\n")
awk -v search="<\!-- NEWER -->" -v replace="$description" '{ gsub(search, replace); print }' CHANGELOG.md > .publish && mv .publish CHANGELOG.md

rm .publish-description
rm .publish

git add CHANGELOG.md
git commit -m "Update CHANGELOG.md for version $TAG"
git push origin main

git tag $TAG
git push origin $TAG