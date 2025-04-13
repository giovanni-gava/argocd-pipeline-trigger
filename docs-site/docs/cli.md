# ⚙️ CLI: argocd-sync

The CLI tool `argocd-sync` allows you to trigger ArgoCD application syncs directly from pipelines, scripts, or terminals.

---

## 📦 Installation

### 🧪 Build locally

```bash
make build-trigger
```

Or build manually:

```bash
go build -o bin/argocd-sync .
```

---

## 🚀 Usage

```bash
./bin/argocd-sync sync \
  --app my-app \
  --server https://argocd.example.com \
  --username admin \
  --password secret \
  --insecure
```

### 📖 Flags

| Flag        | Description                       |
|-------------|-----------------------------------|
| `--app`     | ArgoCD Application name           |
| `--server`  | ArgoCD API Server URL             |
| `--username`| ArgoCD login username             |
| `--password`| ArgoCD password or token          |
| `--insecure`| Skip TLS verification (optional)  |

---

## 🔄 Use in CI/CD

GitHub Actions:
```yaml
- name: Sync app with ArgoCD
  run: ./bin/argocd-sync sync --app my-app --server https://argocd.example.com --username $ARGO_USER --password $ARGO_PASS
```

GitLab CI:
```yaml
sync:
  script:
    - ./bin/argocd-sync sync --app $CI_APP --server $ARGO_URL --username $ARGO_USER --password $ARGO_PASS
```

---

## 📁 Source

See the implementation in [`cmd/sync.go`](https://github.com/giovanni-gava/argocd-pipeline-trigger/blob/main/cmd/sync.go).