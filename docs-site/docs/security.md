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
| HMAC Secret         | Kubernetes Secret or Env Var |

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

## 🔐 Authentication Methods

The `/sync` endpoint supports multiple methods of authentication. You can use one or both:

### 1. Bearer Token Authentication

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

### 2. HMAC Signature Authentication (recommended for CI/CD)

Set a shared secret in your receiver:
```bash
export HMAC_SECRET="super-secure-hmac-key"
```

Then compute the HMAC hash in your CI or test script:
```bash
payload='{"app":"your-app"}'
signature=$(echo -n "$payload" | openssl dgst -sha256 -hmac "$HMAC_SECRET" | sed 's/^.* //')

curl -X POST http://receiver-url/sync \
  -H "Content-Type: application/json" \
  -H "X-Signature: sha256=$signature" \
  -d "$payload"
```

If `HMAC_SECRET` is **not set**, the receiver will skip verification.

---

## 🚨 Rate Limiting (recommended)

Use a reverse proxy or ingress controller to add rate limiting:

```yaml
nginx.ingress.kubernetes.io/limit-connections: "5"
nginx.ingress.kubernetes.io/limit-rpm: "30"
```

---

## 🧠 Suggested Enhancements

| Feature            | Status         |
|--------------------|----------------|
| Bearer token auth  | ✅ Implemented |
| HMAC validation    | ✅ Implemented |
| mTLS enforcement   | ❓ Advanced     |

---

Security is an ongoing effort. Always review deployments, scan images with [Trivy](https://github.com/aquasecurity/trivy), and follow best practices from CNCF and OWASP.

