# 🏗️ Architecture

ArgoCD Pipeline Trigger follows a modular and extensible architecture with clear separation of responsibilities, security-first design, and GitOps-native patterns.

---

## 📐 High-Level Diagram

```
+------------------+           +-------------------+           +----------------+
|   GitHub / CI    |  POST --> |  Webhook Receiver  |  Sync --> |   ArgoCD App    |
+------------------+           +-------------------+           +----------------+
                                           |   |                     
                                           |   |--> Notify (Slack/Telegram)
                                           |   |--> Export /metrics (Prometheus)
```

---

## 🔋 Components

| Component        | Description |
|------------------|-------------|
| `argocd-sync` CLI| Local command-line interface to trigger ArgoCD apps |
| Webhook Receiver | Listens to `/sync` POSTs, runs sync, emits metrics & notifs |
| Notifier         | Slack/Telegram notification layer with custom formatter |
| Helm Chart       | Production-ready deployment with observability and security |

---

## 🔐 Security Boundaries

- Sync only happens on verified `POST /sync`
- Tokens are never exposed — injected via Secrets
- Distroless image and non-root runtime
- Metrics and HTTP exposed on separate endpoints

---

## 🔎 Observability

- Prometheus `/metrics` endpoint with detailed counters and histograms
- Grafana dashboard included in Helm chart
- CI/CD traceability via webhook integration

---

## 🔄 Extensibility

The architecture is modular and can be extended with:
- HMAC or mTLS verification (auth)
- Custom notifiers (Discord, MS Teams, etc.)
- App-specific sync rules or retries
- Payload enrichment (commit SHA, PR number)

---

ArgoCD Pipeline Trigger was built to balance simplicity, observability, and flexibility — making it ideal for modern GitOps workflows.