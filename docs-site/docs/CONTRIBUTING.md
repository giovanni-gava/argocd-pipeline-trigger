# 🤝 Contributing to ArgoCD Pipeline Trigger

Welcome! 🎉 First of all, thank you for considering contributing to this project. Whether you’re fixing bugs, adding new features, or improving documentation, we’re grateful for your help.

---

## 🧱 Project Philosophy

ArgoCD Pipeline Trigger aims to be:
- 💡 **Simple**: Easy to use and extend

- 🔐 **Secure**: Safe by default

- 📊 **Observable**: Prometheus and alert-ready

- 🔁 **GitOps-native**: Ready for real CI/CD pipelines

- 🧰 **DevOps-first**: Designed for real-world infra

---

## 🛠 How to Contribute

### 1. Fork and Clone

```bash
git clone https://github.com/YOUR_USERNAME/argocd-pipeline-trigger.git
cd argocd-pipeline-trigger
git checkout -b my-feature
```

### 2. Run Locally

```bash
make build-webhook
make run-webhook
```

Or with Docker:

```bash
make docker-webhook
```

### 3. Add Your Changes

Follow existing coding patterns, use idiomatic Go, and include comments when necessary.

### 4. Test Your Feature

```bash
go test ./...
```

Check with:
```bash
make lint
```

### 5. Open Pull Request

Please include:
- A clear title and description
- Screenshots or logs if applicable
- Reference to any issues (e.g. `Fixes #12`)

---

## ✅ Code Guidelines

- Follow [Effective Go](https://golang.org/doc/effective_go.html)
- Format with `go fmt`
- Use `cobra` for CLI
- Tests with `testing` or `testify`
- Metrics via `prometheus/client_golang`
- Secrets via env/Helm only

---

## 💬 Communication

Feel free to open [issues](https://github.com/giovanni-gava/argocd-pipeline-trigger/issues), or start a discussion via Pull Request. Feature ideas are welcome!

---

## 📄 License

By contributing, you agree that your contributions will be licensed under the [MIT License](../LICENSE).

Thank you for making ArgoCD Pipeline Trigger better! 🙏

— Giovanni Colognesi

