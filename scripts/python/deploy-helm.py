#!/usr/bin/env python3
import subprocess
import sys

release = sys.argv[1] if len(sys.argv) > 1 else "webhook"
chart = "charts/webhook-receiver"

print(f"📦 Installing Helm chart {chart} as {release}")
subprocess.run(["helm", "upgrade", "--install", release, chart, "--namespace", "default", "--create-namespace"], check=True)
