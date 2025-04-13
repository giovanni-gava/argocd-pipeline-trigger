![ChatGPT Image 13 de abr  de 2025, 00_55_47](https://github.com/user-attachments/assets/d33a7cba-3444-4f69-8eb1-5ca87330ecfa)



# 🚀 ArgoCD Pipeline Trigger


> Automatically trigger ArgoCD sync or refresh after GitHub/GitLab CI pipelines complete, using webhooks, CLI or custom Go tools.

---

## ✨ Overview

**ArgoCD Pipeline Trigger** is a lightweight solution to automate the synchronization of ArgoCD applications after CI/CD pipelines are completed. Ideal for GitOps workflows, this tool supports multiple integration strategies to match your CI and team preferences.

Whether you use GitHub Actions, GitLab CI, or a custom deployment pipeline, this project helps you bridge CI and GitOps with minimal friction.

---

## 📦 Features

- 🔁 Sync ArgoCD applications automatically after PR merges
- 🛠️ Works with **GitHub Actions**, **GitLab CI**, or any CI/CD tool
- 🧩 Supports:
  - Webhook Receiver
  - ArgoCD CLI integration
  - Go-based custom tool
- 🔒 Supports token-based authentication to ArgoCD
- 🧠 Demonstrates real CI/CD + GitOps integration

---

## 🧰 Use Cases

| Trigger              | Method               | Description                         |
|----------------------|----------------------|-------------------------------------|
| PR merged            | Webhook receiver     | CI sends HTTP request to trigger    |
| CI/CD success        | `argocd app sync`    | CLI runs ArgoCD sync from runner    |
| Custom event logic   | Go tool              | Tool connects to ArgoCD API         |

---

## 🖼️ Architecture

```plaintext
          +---------+            +------------------+            +------------------+
          | GitHub  |  PR Merge  | GitHub Actions   |  Webhook   | ArgoCD Sync App  |
          +---------+ ---------->+------------------+----------->+------------------+
                                                              |
                                                              v
                                                    argocd app sync my-app
```

---

## 🔧 Project Structure

```bash
argocd-pipeline-trigger/
├── .github/                # GitHub Actions workflows
├── manifests/              # ArgoCD app and receiver manifests
├── scripts/                # Bash scripts to trigger sync
├── cmd/                    # CLI tool (if using Go)
├── internal/               # Internal logic for CLI/webhook
├── examples/               # Example repo and CI integrations
├── Dockerfile              # For custom webhook receiver image
├── go.mod / go.sum         # If using Go
└── README.md
```

---

## 🚀 Getting Started

### Option 1: Use ArgoCD CLI in your CI

```bash
argocd login $ARGOCD_SERVER --username $USER --password $PASSWORD
argocd app sync my-app
```

### Option 2: Webhook Receiver (sample)

```bash
curl -X POST http://your-receiver/sync \
  -H "Authorization: Bearer $TOKEN" \
  -d '{"app": "my-app", "project": "default"}'
```

### Option 3: Go-based CLI tool (coming soon)

```bash
./argocd-trigger --app my-app --token $ARGOCD_TOKEN
```

---

## 📋 Requirements

- ArgoCD v2.5+
- Token or session access to ArgoCD API
- CI/CD tool with webhook or CLI support

---

## ✅ Roadmap

- [x] Basic sync via CLI
- [x] Bash script with ArgoCD CLI
- [ ] Go-based webhook receiver
- [ ] Custom CLI tool with `cobra`
- [ ] Demo repo with full CI/CD pipeline

---

## 🤝 Contributing

PRs and ideas are welcome!  
Want to share your pipeline? Open an issue or submit an example repo.

---

## 📄 License

[MIT](LICENSE)

---

## 👨‍💻 Author

Made with ❤️ by [Giovanni Gava](https://github.com/giovanni-gava)  
DevOps Engineer | Software Engineer | Cloud Architect | AWS | GCP | Python & Golang Builder | Kubernetes | Terraform | CI/CD Evangelist | Infra as Code Expert | Strategic Problem Solver | Creative Mind 
