#!/bin/bash

set -e

URL=${1:-http://localhost:8080/sync}
APP_NAME=${2:-my-app}

echo "📡 Sending webhook to $URL..."

curl -X POST "$URL" \
  -H "Content-Type: application/json" \
  -d "{\"app\":\"$APP_NAME\"}"
