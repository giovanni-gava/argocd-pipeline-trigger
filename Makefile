# Diretórios
BIN_DIR := bin
DOCKERFILES_DIR := dockerfiles
TRIGGER_DOCKERFILE := $(DOCKERFILES_DIR)/trigger/Dockerfile
RECEIVER_DOCKERFILE := $(DOCKERFILES_DIR)/receiver/Dockerfile

# Imagens
TRIGGER_IMAGE := argocd-sync:latest
RECEIVER_IMAGE := webhook-receiver:latest

.PHONY: all build lint test run clean \
        build-trigger docker-trigger run-trigger scan-trigger \
        build-webhook docker-webhook run-webhook scan-webhook

# === Comandos genéricos ===
all: build-trigger build-webhook

clean:
	@echo "🧹 Cleaning binaries..."
	rm -rf $(BIN_DIR)

lint:
	@echo "🔍 Running golint..."
	go install golang.org/x/lint/golint@latest
	golint ./...

test:
	@echo "🧪 Running tests..."
	go test -v ./...

# === CLI Trigger (argocd-sync) ===
build-trigger:
	@echo "🔧 Building CLI argocd-sync..."
	go build -o $(BIN_DIR)/argocd-sync .

docker-trigger:
	@echo "🐳 Building Docker image for CLI (argocd-sync)..."
	docker build -f $(TRIGGER_DOCKERFILE) -t $(TRIGGER_IMAGE) .

run-trigger:
	@echo "🚀 Running CLI (argocd-sync)..."
	./$(BIN_DIR)/argocd-sync sync --app my-app --server https://argocd.example.com --username admin --password secret --insecure

scan-trigger: docker-trigger
	@echo "🔎 Scanning CLI image with Trivy..."
	trivy image --no-progress --severity HIGH,CRITICAL $(TRIGGER_IMAGE)

# === Webhook Receiver ===
build-webhook:
	@echo "🔧 Building webhook receiver binary..."
	go build -o $(BIN_DIR)/webhook-receiver ./cmd/webhook/main.go

docker-webhook:
	@echo "🐳 Building Docker image for webhook receiver..."
	docker build -f $(RECEIVER_DOCKERFILE) -t $(RECEIVER_IMAGE) .

run-webhook:
	@echo "🚀 Running webhook receiver..."
	ADDR=":8080" ./$(BIN_DIR)/webhook-receiver

scan-webhook: docker-webhook
	@echo "🔎 Scanning webhook image with Trivy..."
	trivy image --no-progress --severity HIGH,CRITICAL $(RECEIVER_IMAGE)
