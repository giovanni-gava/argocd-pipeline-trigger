#!/bin/bash

set -e

REGISTRY=giovannigava  # ou seu user no Docker Hub

echo "🚀 Tagging and pushing CLI image..."
docker tag argocd-sync:latest $REGISTRY/argocd-sync:latest
docker push $REGISTRY/argocd-sync:latest

echo "🚀 Tagging and pushing Webhook Receiver..."
docker tag webhook-receiver:latest $REGISTRY/webhook-receiver:latest
docker push $REGISTRY/webhook-receiver:latest
