#!/bin/bash

set -e

RELEASE_NAME=${1:-webhook}
CHART_PATH=charts/webhook-receiver

echo "🚀 Installing/Upgrading Helm chart..."
helm upgrade --install "$RELEASE_NAME" "$CHART_PATH" --namespace default --create-namespace
