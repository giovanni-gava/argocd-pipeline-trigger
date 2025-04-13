![ChatGPT Image 13 de abr  de 2025, 00_55_47](https://github.com/user-attachments/assets/d1cbb152-1f89-4f29-942d-8cd232d7b3fe)


# 🚀 ArgoCD Pipeline Trigger

> Automatically trigger ArgoCD sync after GitHub/GitLab CI pipelines using webhooks, CLI or custom Go tools.

![Go](https://img.shields.io/badge/built%20with-Go-00ADD8?logo=go&logoColor=white)
![ArgoCD](https://img.shields.io/badge/argocd-integrated-brightgreen?logo=argo)
![CI/CD](https://img.shields.io/badge/ci/cd-github--actions-blue?logo=githubactions)
![Security](https://img.shields.io/badge/scanned%20with-Trivy-critical?logo=aqua)
![License](https://img.shields.io/github/license/giovanni-gava/argocd-pipeline-trigger)

---

## ✨ Overview

**ArgoCD Pipeline Trigger** is a lightweight solution to automate the synchronization of ArgoCD applications after CI/CD pipelines complete.  
Ideal for GitOps workflows, this tool supports multiple integration strategies: Webhooks, CLI, or a custom Go-based tool.

Whether you're using GitHub Actions, GitLab CI, or a custom platform, this project helps you bridge **CI pipelines ↔ GitOps delivery** in a secure, portable, and extensible way.

---

## 📦 Features

- 🔁 Automatically sync ArgoCD apps after PR merges
- 🛠️ Works with **GitHub Actions**, **GitLab CI**, or any CI/CD platform
- 🧩 Supports:
  - Webhook Receiver (coming soon)
  - ArgoCD CLI integration ✅
  - Go-based custom CLI tool ✅
- 🔒 Secure, token-based authentication to ArgoCD
- ⚡ Easy to deploy and integrate into existing pipelines

---

## 🧰 Use Cases

| Trigger            | Method               | Description                           |
|--------------------|----------------------|---------------------------------------|
| PR merged          | Webhook receiver     | CI sends HTTP request to trigger sync |
| CI/CD completed    | `argocd app sync`    | Run ArgoCD CLI directly               |
| Custom event logic | Go CLI tool          | Executes sync securely via API/exec   |

---

## 🖼️ Architecture

```plaintext
          +---------+            +------------------+            +------------------+
          | GitHub  |  PR Merge  | GitHub Actions   |  Webhook   | ArgoCD Controller |
          +---------+ ---------->+------------------+----------->+------------------+
                                                              |
                                                              v
                                                    argocd app sync my-app


🔧 Project Structure

argocd-pipeline-trigger/
├── cmd/                    # CLI command logic (cobra)
├── internal/               # Core logic to connect to ArgoCD
├── .github/workflows/      # GitHub Actions pipeline
├── Dockerfile              # Secure distroless image
├── Makefile                # Build, lint, test, scan, docker
├── go.mod / go.sum         # Module definitions
├── bin/                    # Compiled binary output
└── README.md

🚀 Getting Started
✅ Option 1: Use CLI directly

./bin/argocd-sync sync \
  --app my-app \
  --server https://argocd.example.com \
  --username admin \
  --password secret \
  --insecure

🐳 Option 2: Use Docker image

docker build -t argocd-sync .
docker run --rm argocd-sync sync \
  --app my-app \
  --server https://argocd.example.com \
  --username admin \
  --password secret \
  --insecure

🛠 Option 3: Run via GitHub Actions

- name: 🔑 Login and Sync ArgoCD
  run: |
    argocd login $ARGOCD_SERVER \
      --username $ARGOCD_USERNAME \
      --password $ARGOCD_PASSWORD \
      --insecure
    argocd app sync my-app

🧪 Development
🔧 Build

make build


🧼 Lint & Test

make lint
make test


🔐 Security Scan

make scan

Uses Trivy to detect HIGH/CRITICAL vulnerabilities in the Docker image.


✅ Requirements
Go 1.21+

ArgoCD CLI installed (inside Docker or system)

ArgoCD v2.5+

Valid user/token for authentication

📋 Roadmap
 CLI tool with Cobra

 GitHub Actions pipeline integration

 Secure Docker image (distroless)

 Trivy scan included in Makefile

 Webhook receiver in Go

 Helm chart for deploying CLI/receiver

 Optional Slack/Discord notifier


🤝 Contributing
Pull requests, ideas, and discussions are welcome!
Share your integration or fork with #argocd-pipeline-trigger ❤️

👨‍💻 Author
Made with 💙 by Giovanni Gava
DevOps Engineer | Software Engineer (Go) | Cloud Architect | GitOps Evangelist

![ChatGPT Image 13 de abr  de 2025, 01_22_58](https://github.com/user-attachments/assets/0ed8c488-51de-49cf-8006-00cf9a52367d)


