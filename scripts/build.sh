#!/bin/bash -e
GITHUB_TAG="${GITHUB_TAG:-local}"
SHA512_CMD="${SHA512_CMD:-sha512sum}"
export DELIVERABLE="${DELIVERABLE:-n3dr}"

echo "GITHUB_TAG: '$GITHUB_TAG' DELIVERABLE: '$DELIVERABLE'"
go build -ldflags "-X github.com/030/n3dr/cmd.Version=${GITHUB_TAG}" -o "${DELIVERABLE}"
$SHA512_CMD "${DELIVERABLE}" > "${DELIVERABLE}.sha512.txt"
chmod +x "${DELIVERABLE}"
