# 📦 Changelog

## [v1.0.0] - 2025-04-18

🎉 **Initial production-ready release**

### 🚀 Features
- CLI `argocd-sync` written in Go with Cobra
- Webhook Receiver service with Prometheus `/metrics`
- Slack and Telegram notifications with rich formatting
- Full Helm Chart for Kubernetes deployment
- Dockerfiles (distroless, secure, non-root)
- Grafana dashboard via ConfigMap
- CI/CD integration examples (GitHub, GitLab, Bitbucket, Jenkins, CircleCI)
- MkDocs documentation site with Material theme

### 🔐 Security
- Auth via `Bearer Token` in `/sync` endpoint
- Secrets handled via ENV or Kubernetes Secrets
- JSON request validation for safety
- Non-root distroless images

### 📈 Observability
- Prometheus metrics: `sync_total`, `sync_duration`
- Grafana dashboard for sync monitoring
- Custom labels: `app`, `status`, `duration`

### 🧪 Tooling
- `Makefile` with build, lint, scan, docker
- `docker-compose.yml` for local Prometheus + Receiver
- Test scripts in Go and Python

---

Looking for older releases? Stay tuned for versioned tags and [GitHub Releases](https://github.com/giovanni-gava/argocd-pipeline-trigger/releases).

Want to contribute? Check out [CONTRIBUTING.md](./CONTRIBUTING.md).
