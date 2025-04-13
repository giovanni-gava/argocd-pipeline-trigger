# 📈 Metrics & Observability

ArgoCD Pipeline Trigger includes native support for Prometheus metrics and Grafana dashboards — allowing teams to monitor sync events, performance, and failure rates.

---

## 🔌 Endpoint

A `/metrics` endpoint is exposed on the receiver:

```
GET /metrics
```

Use `kubectl port-forward` or a Kubernetes ServiceMonitor to scrape it.

---

## 📊 Metrics Available

| Metric                          | Type     | Labels         | Description                            |
|--------------------------------|----------|----------------|----------------------------------------|
| `argocd_sync_total`            | Counter  | `app`, `status`| Total syncs triggered per app/status   |
| `argocd_sync_duration_seconds` | Histogram| `app`          | Duration of sync operation in seconds  |

---

## 🔍 Visualize with Grafana

A dashboard is provided as a ConfigMap and includes:
- Total syncs (stat)
- Sync duration per app (graph)
- Sync rate and errors (table)

See: [`templates/grafana-dashboard.yaml`](https://github.com/giovanni-gava/argocd-pipeline-trigger/blob/main/charts/webhook-receiver/templates/grafana-dashboard.yaml)

---

## 🧭 Prometheus Operator

The Helm chart also includes a `ServiceMonitor`:

```yaml
apiVersion: monitoring.coreos.com/v1
kind: ServiceMonitor
```

Enable it with:
```bash
kubectl apply -f charts/webhook-receiver/templates/servicemonitor.yaml
```

Or template via Helm.

---

## 🧠 Example Queries

```promql
sum(rate(argocd_sync_total[5m])) by (app)
rate(argocd_sync_total{status="failure"}[5m])
avg(argocd_sync_duration_seconds) by (app)
```

These queries can be used to build alerts or dashboards for:
- Deployment frequency (DORA metric)
- Failure rate
- Time to deploy

---

For more, visit the [Grafana page](https://grafana.com/) or use the included dashboard JSON.

