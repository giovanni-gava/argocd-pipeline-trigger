#!/usr/bin/env python3
import requests
import sys
import json

url = sys.argv[1] if len(sys.argv) > 1 else "http://localhost:8080/sync"
app = sys.argv[2] if len(sys.argv) > 2 else "my-app"

print(f"📡 Sending webhook to {url} for app: {app}")
payload = {"app": app}

response = requests.post(url, json=payload)
print(f"✅ Status: {response.status_code}")
print(response.text)
