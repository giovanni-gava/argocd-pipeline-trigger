#!/usr/bin/env python3
import subprocess
import sys

cluster = sys.argv[1] if len(sys.argv) > 1 else "argocd-trigger"

print(f"🧱 Creating Kind cluster: {cluster}")
subprocess.run(["kind", "create", "cluster", "--name", cluster, "--wait", "60s"], check=True)

print("🧪 Verifying cluster...")
subprocess.run(["kubectl", "cluster-info", "--context", f"kind-{cluster}"], check=True)

print("🚀 Deploying webhook receiver via Helm...")
subprocess.run(["helm", "upgrade", "--install", "webhook", "charts/webhook-receiver"], check=True)
