# 🐳 Docker & Local Testing

This project includes Dockerfiles for both the CLI (`argocd-sync`) and the webhook receiver, as well as a `docker-compose.yml` for local development and testing.

---

## 🧱 Docker Overview

| Component        | Image Path                        | Dockerfile                                 |
|------------------|------------------------------------|---------------------------------------------|
| CLI Tool         | `argocd-sync:latest`              | `dockerfiles/trigger/Dockerfile`           |
| Webhook Receiver | `webhook-receiver:latest`         | `dockerfiles/receiver/Dockerfile`          |

---

## ⚙️ Build the CLI Docker Image

```bash
docker build -f dockerfiles/trigger/Dockerfile -t argocd-sync:latest .
```

### Run the CLI Container

```bash
docker run --rm argocd-sync:latest \
  sync \
  --app my-app \
  --server https://argocd.example.com \
  --username admin \
  --password secret \
  --insecure
```

---

## 🌐 Build the Receiver Docker Image

```bash
docker build -f dockerfiles/receiver/Dockerfile -t webhook-receiver:latest .
```

### Run the Webhook Receiver

```bash
docker run --rm -p 8080:8080 \
  -e ENABLE_TELEGRAM=true \
  -e TELEGRAM_BOT_TOKEN=xxxx \
  -e TELEGRAM_CHAT_ID=xxxx \
  webhook-receiver:latest
```

---

## 🧪 Local Testing with Docker Compose

A `docker-compose.yml` is included to spin up:
- Webhook Receiver
- Prometheus (scraping metrics)
- Grafana (dashboard view)

### 🏁 Start it up

```bash
docker-compose up --build
```

### 📬 Test Webhook

```bash
curl -X POST http://localhost:8080/sync \
  -H "Content-Type: application/json" \
  -d '{"app":"demo-app"}'
```

You should see:
- ArgoCD sync log in terminal
- Telegram/Slack notification (if enabled)
- `/metrics` exposed for Prometheus

---

## 🧪 Testing Notification Format Only

```bash
go run tests/test_notify_telegram.go
```

or

```bash
python3 scripts/test_webhook.py http://localhost:8080/sync
```

---

This makes testing and iterating locally easy and fully observable. See the [README](../README.md) for full project context.