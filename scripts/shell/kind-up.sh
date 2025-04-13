#!/bin/bash

set -e

CLUSTER_NAME=${1:-argocd-trigger}
KUBECONFIG_PATH=$HOME/.kube/kind-config-$CLUSTER_NAME

echo "🧱 Creating Kind cluster: $CLUSTER_NAME..."
kind create cluster --name "$CLUSTER_NAME" --wait 60s

echo "🧪 Testing kubectl..."
kubectl cluster-info --context kind-"$CLUSTER_NAME"

echo "⚙️ Installing Helm chart..."
helm upgrade --install webhook charts/webhook-receiver
