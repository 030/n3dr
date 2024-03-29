#!/bin/bash -e
GITHUB_TAG="${GITHUB_TAG:-local}"
SHA512_CMD="${SHA512_CMD:-sha512sum}"
export N3DR_DELIVERABLE="${N3DR_DELIVERABLE:-n3dr}"

echo "GITHUB_TAG: '$GITHUB_TAG' N3DR_DELIVERABLE: '$N3DR_DELIVERABLE'"
cd cmd/n3dr
go build -buildvcs=false -ldflags "-X main.Version=${GITHUB_TAG}" -o "${N3DR_DELIVERABLE}"
$SHA512_CMD "${N3DR_DELIVERABLE}" >"${N3DR_DELIVERABLE}.sha512.txt"
chmod +x "${N3DR_DELIVERABLE}"
cd ../..
