# 🌐 Webhook Receiver

The webhook receiver is a lightweight Go service that listens for incoming CI/CD events (e.g., PR merges) and triggers an `argocd app sync` operation.

---

## 🚀 How it works

- Exposes a single POST endpoint: `/sync`
- Receives JSON payload with `{"app":"your-app"}`
- Executes `argocd app sync <app>` under the hood
- Sends notification via Slack or Telegram (if enabled)
- Exports Prometheus metrics (`/metrics`)
- Now supports: **Bearer Token** + **HMAC-SHA256 Signature** auth

---

## 📦 Environment Variables

| Name           | Required | Description |
|----------------|----------|-------------|
| `ADDR`         | No       | Port to bind (default `:8080`) |
| `AUTH_TOKEN`   | Optional | Bearer token for authorization |
| `HMAC_SECRET`  | Optional | HMAC secret for signature validation |
| `ENABLE_SLACK` | Optional | Enable Slack notifications |
| `ENABLE_TELEGRAM` | Optional | Enable Telegram notifications |

---

## 🔐 Authentication Options

### 1. Bearer Token

Set `AUTH_TOKEN` on the receiver:
```bash
export AUTH_TOKEN="my-token"
```

Send requests with:
```bash
curl -X POST http://localhost:8080/sync \
  -H "Authorization: Bearer my-token" \
  -H "Content-Type: application/json" \
  -d '{"app":"demo-app"}'
```

### 2. HMAC Signature (recommended)

Set `HMAC_SECRET` on the receiver:
```bash
export HMAC_SECRET="super-hmac"
```

Generate signature:
```bash
payload='{"app":"demo-app"}'
signature=$(echo -n "$payload" | openssl dgst -sha256 -hmac "$HMAC_SECRET" | sed 's/^.* //')
```

Send request:
```bash
curl -X POST http://localhost:8080/sync \
  -H "Content-Type: application/json" \
  -H "X-Signature: sha256=$signature" \
  -d "$payload"
```

---

## 📈 Metrics

Prometheus endpoint available at `/metrics`.

| Metric Name            | Type     | Description |
|------------------------|----------|-------------|
| `sync_total`           | Counter  | Total syncs triggered |
| `sync_duration_seconds`| Histogram | Sync durations |

---

## 🧪 Test locally

```bash
go run cmd/webhook/main.go
```

```bash
curl -X POST http://localhost:8080/sync \
  -H "Authorization: Bearer my-token" \
  -H "Content-Type: application/json" \
  -d '{"app":"demo-app"}'
```

---

## ✅ Production Tips

- Use `HMAC_SECRET` or `AUTH_TOKEN` (or both!)
- Use a reverse proxy for TLS and rate limiting
- Mount secrets via Kubernetes or Vault
- Scrape `/metrics` with Prometheus

---

For full deployment, see the [Helm Chart](helm.md) and [Security Guide](security.md).

