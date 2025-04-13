#!/bin/bash

set -euo pipefail

export TELEGRAM_BOT_TOKEN=${TELEGRAM_BOT_TOKEN:-"your-telegram-token"}
export TELEGRAM_CHAT_ID=${TELEGRAM_CHAT_ID:-"your-chat-id"}

echo "🚀 Starting Webhook Receiver + Prometheus + Grafana stack..."
docker compose up --build
