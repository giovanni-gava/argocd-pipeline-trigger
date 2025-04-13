# 🔔 Integrations: Slack, Telegram, CI/CD

ArgoCD Pipeline Trigger offers seamless integration with modern toolchains including CI/CD, Slack, and Telegram to improve visibility and automation.

---

## 💬 Slack Notifications

### 🔧 Setup

1. Create an [Incoming Webhook](https://api.slack.com/messaging/webhooks)
2. Store it in a Kubernetes secret:

```bash
kubectl create secret generic webhook-receiver-secrets \
  --from-literal=slackWebhookUrl="https://hooks.slack.com/services/..."
```

3. Enable in `values.yaml`:
```yaml
notifications:
  slack:
    enabled: true
```

### ✅ Message Example
```
✅ Sync Triggered
👨🏻‍💻 Giovanni Colognesi - ArgoCD Trigger
🔹 App: `my-app`
⏱️ Duration: `0.82s`
```

---

## 📲 Telegram Notifications

### 🔧 Setup
1. Talk to [@BotFather](https://t.me/BotFather)
2. Create a bot and get the token
3. Start a conversation or group with the bot
4. Find your chat ID:
```bash
curl https://api.telegram.org/bot<token>/getUpdates
```

5. Create a Kubernetes secret:
```bash
kubectl create secret generic webhook-receiver-secrets \
  --from-literal=telegramBotToken="123:ABC" \
  --from-literal=telegramChatId="12345678"
```

6. Enable in `values.yaml`:
```yaml
notifications:
  telegram:
    enabled: true
```

---

## 🔁 CI/CD Pipelines

Use `curl` to trigger the webhook after deploys:

### GitHub Actions
```yaml
- name: Trigger ArgoCD sync
  run: |
    curl -X POST https://your-receiver/sync \
      -H "Content-Type: application/json" \
      -d '{"app":"my-app"}'
```

### GitLab CI
```yaml
trigger:
  script:
    - curl -X POST https://receiver/sync -H "Content-Type: application/json" -d '{"app":"my-app"}'
```

---

## 🧪 Testing Notifications

Use the Go or Python scripts provided under `scripts/`:

```bash
go run scripts/test_notify_telegram.go
python3 scripts/test_webhook.py
```

These allow you to validate formatting and delivery without deploying to K8s.

