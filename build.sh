#!/bin/bash -e
TOOL="${1:-./n3dr}"
TRAVIS_TAG="${TRAVIS_TAG:-local}"
SHA512_CMD="${SHA512_CMD:-sha512sum}"
DELIVERABLE="${DELIVERABLE:-n3dr}"

echo "TRAVIS_TAG: '$TRAVIS_TAG' DELIVERABLE: '$DELIVERABLE'"
go build -ldflags "-X n3dr/cmd.Version=${TRAVIS_TAG}" -o "${DELIVERABLE}"
$SHA512_CMD "${TOOL}" > "${DELIVERABLE}.sha512.txt"
chmod +x "${DELIVERABLE}"
