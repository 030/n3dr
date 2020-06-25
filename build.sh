#!/bin/bash -e
echo "TRAVIS_TAG: '$TRAVIS_TAG' DELIVERABLE: '$DELIVERABLE'"
go build -ldflags "-X n3dr/cmd.Version=${TRAVIS_TAG}" -o "${DELIVERABLE}"
$SHA512_CMD "${TOOL}" > "${DELIVERABLE}.sha512.txt"
chmod +x "${DELIVERABLE}"
