#!/usr/bin/env python3
import subprocess

def build_image(tag, dockerfile):
    print(f"🔧 Building image: {tag}")
    subprocess.run(["docker", "build", "-f", dockerfile, "-t", tag, "."], check=True)

build_image("argocd-sync:latest", "dockerfiles/trigger/Dockerfile")
build_image("webhook-receiver:latest", "dockerfiles/receiver/Dockerfile")
