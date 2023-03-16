#!/bin/bash

THISDIR="$( cd "$(dirname "$0")" ; pwd -P )"
TESTDIR="$THISDIR/../test/e2e"

# Defaults
NPM_TARGET="test"

# Parse script arguments
for i in "$@"
do
case "$i" in
	--all)
	NPM_TARGET="test-all"
	shift # past argument
	;;
	*)
	die "unknown option '$i'"
	;;
esac
done

# Exit as soon as any line fails
set -e

cd "$TESTDIR"

npm install
npm run $NPM_TARGET
