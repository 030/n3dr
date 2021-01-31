#!/bin/bash -e
TRAVIS_TAG="${TRAVIS_TAG:-local}"
SHA512_CMD="${SHA512_CMD:-sha512sum}"
export DELIVERABLE="${DELIVERABLE:-n3dr}"

echo "TRAVIS_TAG: '$TRAVIS_TAG' DELIVERABLE: '$DELIVERABLE'"
go build -ldflags "-X github.com/030/n3dr/cmd.Version=${TRAVIS_TAG}" -o "${DELIVERABLE}"
$SHA512_CMD "${DELIVERABLE}" > "${DELIVERABLE}.sha512.txt"
chmod +x "${DELIVERABLE}"
