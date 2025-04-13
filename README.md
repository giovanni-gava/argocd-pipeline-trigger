![ChatGPT Image 13 de abr  de 2025, 00_55_47](https://github.com/user-attachments/assets/d1cbb152-1f89-4f29-942d-8cd232d7b3fe)

# ArgoCD Pipeline Trigger 🔁

> 🚀 Trigger ArgoCD App Sync automatically from CI/CD pipelines with optional Slack/Telegram notifications, full GitOps support, and enterprise-level architecture.


---

## 📚 Overview

**ArgoCD Pipeline Trigger** is a production-ready, extensible, and secure tool built with Go to automate ArgoCD app syncs after pull requests are merged.

It features:

- ✅ CLI tool (`argocd-sync`) written in Go
- ✅ Webhook Receiver service (distroless) for event-driven sync
- ✅ Slack & Telegram notifications with rich formatting
- ✅ Helm chart for easy Kubernetes deployment
- ✅ GitOps-ready architecture and full CI/CD integration

---

## 📐 Architecture

```text
              +-------------+          +-----------------+          +------------------+
              | GitHub/GitLab|  PR → CI|  CLI or Webhook |  Sync → |    ArgoCD App    |
              +-------------+   -----> +-----------------+   -----> +------------------+
                                         |             |                |
                                         |             |-- Notify --> Slack/Telegram
                                  argocd-sync     webhook-receiver
```

---

## 🧰 Project Structure

```
argocd-pipeline-trigger/
├── cmd/                      # CLI and receiver entrypoints
│   ├── root.go
│   ├── sync.go
│   └── webhook/main.go
├── internal/notifier/       # Slack and Telegram integrations
│   ├── notifier.go
│   ├── slack.go
│   ├── telegram.go
│   └── formatter.go
├── dockerfiles/             # Distroless Dockerfiles
│   ├── trigger/Dockerfile
│   └── receiver/Dockerfile
├── scripts/                 # Python/Go automation scripts
├── charts/                  # Helm chart for K8s deployment
│   └── webhook-receiver/
├── tests/                   # Manual validation scripts
├── Makefile                 # Build and automation
├── go.mod / go.sum
└── README.md
```

---

## ⚙️ CLI Tool: `argocd-sync`

### 💻 Usage

```bash
./argocd-sync sync \
  --app my-app \
  --server https://argocd.example.com \
  --username admin \
  --password secret \
  --insecure
```

### ✅ Features
- Go `cobra`-based CLI
- Build with `make build-trigger`
- Docker-ready with `make docker-trigger`

---

## 🌐 Webhook Receiver

### 🚀 Start locally

```bash
ADDR=":8080" go run cmd/webhook/main.go
```

### 🧪 Test
```bash
curl -X POST http://localhost:8080/sync \
  -H "Content-Type: application/json" \
  -d '{"app":"my-app"}'
```

### ✅ Features
- Written in idiomatic Go
- Timeout + recovery
- Secure input validation
- Sends notifications post-sync

---

## 💬 Notifications

### ✅ Slack
1. Create a webhook in Slack workspace
2. Export:
```bash
export ENABLE_SLACK=true
export SLACK_WEBHOOK_URL=https://hooks.slack.com/services/...
```

### ✅ Telegram
1. Create bot via [@BotFather](https://t.me/BotFather)
2. Start chat or group, send a message
3. Get `chat_id` via:
```bash
curl https://api.telegram.org/bot<BOT_TOKEN>/getUpdates
```
4. Export:
```bash
export ENABLE_TELEGRAM=true
export TELEGRAM_BOT_TOKEN=your:token
export TELEGRAM_CHAT_ID=12345678
```

---

## 📦 Helm Chart

### 🧭 Deploy to Kubernetes

```bash
helm upgrade --install webhook charts/webhook-receiver \
  --set notifications.slack.enabled=true \
  --set-file notifications.slack.webhookUrl=./my-slack-secret.txt
```

### 🔐 Secure secrets
Use sealed-secrets or Helm secret values to avoid exposing credentials:
```bash
kubectl create secret generic webhook-receiver-secrets \
  --from-literal=slackWebhookUrl=... \
  --from-literal=telegramBotToken=... \
  --from-literal=telegramChatId=...
```

---

## 🧪 Test Utilities

### ✅ Send test to Telegram
```bash
go run tests/test_notify_telegram.go
```

### 🐍 Python alternative
```bash
python3 scripts/test_webhook.py http://localhost:8080/sync
```

### 🔁 CI Integration
Add in GitHub/GitLab pipelines:
```bash
curl -X POST $WEBHOOK_URL -H "Content-Type: application/json" -d '{"app":"my-app"}'
```

---

## 🔐 Security

- ✅ Uses `distroless` images
- ✅ Build statically with `CGO_ENABLED=0`
- ✅ No tokens hardcoded
- ✅ Timeout in all HTTP clients
- ✅ Secrets loaded via Kubernetes secrets
- ✅ Validated JSON payloads

Run security scan:
```bash
make scan-webhook
```

---

## 🎯 Why This Project?

This project was designed to:
- Replace manual `argocd app sync` after merge
- Bring **GitOps automation** to life in CI/CD
- Showcase clean Go + DevOps engineering
- Offer a plug-and-play solution for teams
- Demonstrate **security, observability, and extensibility**

---

## 👨🏻‍💻 Author

Feito com 💙 por **Giovanni Colognesi**  
DevOps Engineer | Software Engineer | Cloud Architect | AWS | GCP | Python & Golang Builder | Kubernetes | Terraform | CI/CD Evangelist | Infra as Code Expert | Strategic Problem Solver | Creative Mind 
[🔗 LinkedIn](https://linkedin.com/in/giovanni-gava) | [🐙 GitHub](https://github.com/giovanni-gava)

---

## ⭐ Star this repo if it helped you!

Contribuições são bem-vindas — abra uma issue, envie um PR ou compartilhe! 🚀



