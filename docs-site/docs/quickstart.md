# 🚀 Quickstart: ArgoCD Pipeline Trigger

> Get your Webhook Receiver up and running with Helm + Telegram integration in under 2 minutes.

---

## 🧭 Prerequisites

- Kubernetes cluster (Kind, Minikube, EKS, etc.)
- Helm installed (`v3+`)
- Telegram bot created ([via @BotFather](https://t.me/BotFather))

---

## 1. Clone the repository

```bash
git clone https://github.com/giovanni-gava/argocd-pipeline-trigger.git
cd argocd-pipeline-trigger
```

---

## 2. Get your Telegram Chat ID

1. Create your bot with [@BotFather](https://t.me/BotFather)
2. Send a message to the bot
3. Run:

```bash
curl -s "https://api.telegram.org/bot<TELEGRAM_BOT_TOKEN>/getUpdates"
```

4. Copy the value of `chat.id`

---

## 3. Create the secret with your Telegram credentials

```bash
kubectl create secret generic webhook-receiver-secrets \
  --from-literal=telegramBotToken="<TELEGRAM_BOT_TOKEN>" \
  --from-literal=telegramChatId="<TELEGRAM_CHAT_ID>"
```

---

## 4. Install using Helm 🚀

```bash
helm upgrade --install webhook charts/webhook-receiver \
  --set notifications.telegram.enabled=true \
  --set auth.token="my-token" \
  --set hmac.secret="super-hmac"
```

---

## 5. Send a test Webhook

### A. Port-forward the service

```bash
kubectl port-forward svc/webhook-receiver 8080:80
```

### B. Send via HMAC (recommended)

```bash
export HMAC_SECRET="super-hmac"
payload='{"app":"demo-app"}'
signature=$(echo -n "$payload" | openssl dgst -sha256 -hmac "$HMAC_SECRET" | sed 's/^.* //')
```
```bash
curl -X POST http://localhost:8080/sync \
  -H "Content-Type: application/json" \
  -H "X-Signature: sha256=$signature" \
  -d "$payload"
```

### C. Or send with Bearer Token

```bash
curl -X POST http://localhost:8080/sync \
  -H "Authorization: Bearer my-token" \
  -H "Content-Type: application/json" \
  -d '{"app": "demo-app"}'
```

---

## 📬 Expected Result

You should receive a message in Telegram:

```
✅ Sync Triggered
👨🏻‍💻 Giovanni Colognesi - ArgoCD Trigger
🔹 App: `demo-app`
⏱️ Duration: `XYZms`
```

---

## ✅ Done!

Your Webhook Receiver with metrics, logging, and notifications is now up and running.

Explore:
- `/metrics` endpoint for Prometheus
- Full Helm chart in `charts/webhook-receiver`
- CI/CD pipeline integration via `curl`

---

Need help? Reach out on [LinkedIn](https://linkedin.com/in/giovanni-gava) or open an issue on GitHub 💙

