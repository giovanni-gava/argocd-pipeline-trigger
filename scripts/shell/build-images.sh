#!/bin/bash

set -e

echo "🔧 Building CLI image..."
docker build -f dockerfiles/trigger/Dockerfile -t argocd-sync:latest .

echo "🔧 Building Webhook Receiver image..."
docker build -f dockerfiles/receiver/Dockerfile -t webhook-receiver:latest .
