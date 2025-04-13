# 🧭 Helm Deployment

This project includes a production-ready [Helm chart](https://helm.sh/) to deploy the webhook receiver and all its capabilities (metrics, secrets, notifications).

---

## 📁 Chart Location

```bash
charts/webhook-receiver/
```

---

## 🚀 Install

```bash
helm upgrade --install webhook charts/webhook-receiver \
  --namespace default --create-namespace
```

---

## ⚙️ Values Overview

| Key | Description |
|-----|-------------|
| `notifications.slack.enabled` | Enable Slack notifier |
| `notifications.slack.webhookUrl` | Slack Webhook URL (via Secret) |
| `notifications.telegram.enabled` | Enable Telegram notifier |
| `notifications.telegram.botToken` | Bot Token (via Secret) |
| `notifications.telegram.chatId` | Chat ID (via Secret) |
| `metrics.enabled` | Enable Prometheus exporter |
| `metrics.path` | Path to metrics endpoint |
| `metrics.port` | Port for metrics exposure |

---

## 🔐 Secrets Handling

We recommend creating Kubernetes secrets externally:

```bash
kubectl create secret generic webhook-receiver-secrets \
  --from-literal=telegramBotToken="<token>" \
  --from-literal=telegramChatId="<chat_id>"
```

---

## 📡 ServiceMonitor (Prometheus Operator)

Deploy manually or use `helm template`:
```bash
kubectl apply -f charts/webhook-receiver/templates/servicemonitor.yaml
```

---

## 🎨 Grafana Dashboard

A dashboard is included via ConfigMap:
```bash
kubectl apply -f charts/webhook-receiver/templates/grafana-dashboard.yaml
```

This allows dashboards to be versioned as code and loaded via sidecar.

---

## 🧪 Test Receiver

```bash
kubectl port-forward svc/webhook-receiver 8080:80
```

Then send a sync request:
```bash
curl -X POST http://localhost:8080/sync -H "Content-Type: application/json" -d '{"app":"test-app"}'
```

---

The Helm chart is maintained under `charts/` and accepts PRs for new features, values, and templates.