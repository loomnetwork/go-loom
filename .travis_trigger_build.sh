#!/bin/bash

set -euxo pipefail

if [ "${TRAVIS_BRANCH}" != "master" ]; then
  echo "Not triggering downstream build since this is not master"
  exit 0
fi

body='{
"request": {
"message": "Triggered by go-loom",
"branch": "master"
}}'

curl -s -X POST \
  -H "Content-Type: application/json" \
  -H "Accept: application/json" \
  -H "Travis-API-Version: 3" \
  -H "Authorization: token ${TRAVIS_TOKEN}" \
  -d "$body" \
  https://api.travis-ci.org/repo/loomnetwork%2Ftiles-chain/requests

