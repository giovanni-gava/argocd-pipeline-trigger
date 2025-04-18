# 🔐 Security Best Practices

ArgoCD Pipeline Trigger is designed with safety and production-readiness in mind. Below are the key practices and features included to support secure deployment and usage.

---

## ✅ Transport Security

- All traffic should be secured via **TLS termination**
- Use Ingress controllers or service mesh (Istio, Linkerd) for HTTPS
- Do **not expose** the webhook directly without protection

---

## 🔑 Secret Management

| What                | How it's secured |
|---------------------|------------------|
| Slack Webhook URL   | Kubernetes Secret |
| Telegram Bot Token  | Kubernetes Secret |
| Telegram Chat ID    | Kubernetes Secret |
| Receiver Auth Token | Kubernetes Secret or Env Var |

Use sealed-secrets or SOPS for GitOps-friendly secret management.

---

## 🧱 Runtime Protections

- Webhook requests use **timeouts (10s)** to prevent long-running calls
- Docker images are based on **distroless** to minimize attack surface
- User runs as non-root (`runAsUser: 65532`)

---

## 🧪 Input Validation

- Payloads are parsed as **strict JSON**
- Requests missing `app` or malformed will return `400 Bad Request`

---

## 🔐 Authentication (Bearer Token)

For additional protection, the `/sync` endpoint supports **optional Bearer Token authentication**.

### Enable Authentication:
Set an environment variable in the receiver:
```bash
export AUTH_TOKEN="your-super-secret-token"
```

Then send requests using:
```bash
curl -X POST http://receiver-url/sync \
  -H "Authorization: Bearer your-super-secret-token" \
  -H "Content-Type: application/json" \
  -d '{"app":"your-app"}'
```

If `AUTH_TOKEN` is **not set**, the endpoint is **open** (useful for internal testing).

For production, it's recommended to use this token via:
- ENV in Kubernetes Deployment
- `valueFrom` linked to a Kubernetes Secret

---

## 🚨 Rate Limiting (recommended)

Use a reverse proxy or ingress controller to add rate limiting:

```yaml
nginx.ingress.kubernetes.io/limit-connections: "5"
nginx.ingress.kubernetes.io/limit-rpm: "30"
```

---

## 🧠 Suggested Enhancements (optional)

| Feature            | Status |
|--------------------|--------|
| Bearer token auth  | ✅ implemented |
| HMAC validation    | 🔜 planned |
| mTLS enforcement   | ❓ advanced |

---

Security is an ongoing effort. Always review deployments, scan images with [Trivy](https://github.com/aquasecurity/trivy), and follow best practices from CNCF and OWASP.

