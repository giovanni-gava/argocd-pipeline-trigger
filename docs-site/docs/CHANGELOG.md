# 📦 Changelog

## [v1.1.0] - 2025-04-18

🔐 **HMAC Signature Authentication**, improved Helm integration, and production-grade security enhancements.

### 🚀 Features
- ✅ Added **HMAC-SHA256 signature verification** for the webhook receiver
- ✅ Compatible with existing Bearer Token authentication (can use both)
- ✅ New environment variable: `HMAC_SECRET`
- ✅ Updated Helm `values.yaml` with `hmac.secret` and `secretRef`
- ✅ Example `curl` and `openssl` usage for CI/CD integrations

### 📚 Documentation
- ➕ `security.md`: authentication methods clarified and expanded
- ➕ `receiver.md`: how to test, use, and configure HMAC
- ➕ `quickstart.md`: includes HMAC test via curl

### 🔐 Security
- ✅ Dual-authentication support (Bearer + HMAC)
- ✅ Runtime-safe body reuse using `io.NopCloser`
- ✅ Skips auth if unset (safe for dev, protected in prod)

### 🧠 Internal Improvements
- 📦 Prometheus metrics preserved with per-app labeling
- 📁 Architecture remains clean and performant with no extra dependencies

---

## [v1.0.0] - 2025-04-16

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

