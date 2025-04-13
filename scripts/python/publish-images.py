#!/usr/bin/env python3
import subprocess

registry = "giovannigava"

images = {
    "argocd-sync:latest": f"{registry}/argocd-sync:latest",
    "webhook-receiver:latest": f"{registry}/webhook-receiver:latest"
}

for local, remote in images.items():
    print(f"🚀 Pushing {remote}")
    subprocess.run(["docker", "tag", local, remote], check=True)
    subprocess.run(["docker", "push", remote], check=True)
