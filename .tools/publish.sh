#!/bin/bash

TAG=$1

if [ -z "$TAG" ]
then
  echo "ERROR: no version provided. Usage: make publish version=<tag>"
  exit 0
fi

if ! [[ $TAG =~ ^v[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
	echo "ERROR: version does not follow semantic versioning starting with 'v'. Example: v1.0.0"
	exit 0
fi

git tag $TAG
git push origin $TAG