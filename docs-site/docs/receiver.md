# 🌐 Webhook Receiver

The Webhook Receiver is a lightweight, stateless HTTP service built in Go that listens for `POST` requests from CI/CD tools and triggers an ArgoCD sync operation.

---

## 🔧 How it works

1. CI merges a PR (e.g., GitHub/GitLab)
2. CI sends a `POST /sync` to the Webhook Receiver
3. Receiver runs `argocd app sync <app>` locally
4. Duration and status are measured and exposed as Prometheus metrics
5. A notification is sent to Slack/Telegram (optional)

---

## 🔐 Security Considerations

| Area         | Strategy |
|--------------|----------|
| Transport    | Use TLS (via ingress/controller) |
| Auth         | Can be extended with token header (planned) |
| Secrets      | Managed via Kubernetes Secret only |
| Timeout      | Sync process has hard timeout (default: 10s) |
| Rate limit   | Recommend external NGINX / Gateway limiters |
| Input        | Strict JSON decoding and validation |

---

## 📨 Payload format

```json
{
  "app": "my-argocd-app"
}
```

If the `app` field is missing or malformed, the receiver responds with `400 Bad Request`.

---

## ✅ Sample Response

```json
HTTP/1.1 200 OK

ok
```

On error:
```json
HTTP/1.1 500 Internal Server Error

sync failed: app not found
```

---

## 🔁 Why not use ArgoCD webhooks directly?

This project decouples sync automation from ArgoCD itself, allowing teams to:
- Add metrics, alerts, dashboards
- Control sync timing and conditions
- Integrate notifications
- Keep ArgoCD declarative while syncing on-demand