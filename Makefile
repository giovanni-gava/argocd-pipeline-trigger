APP_NAME := argocd-sync
BIN_DIR := bin
IMAGE_NAME := $(APP_NAME):latest

.PHONY: all build lint test run clean scan docker-build docker-run

all: build

build:
	@echo "🔧 Building $(APP_NAME)..."
	go build -o $(BIN_DIR)/$(APP_NAME) .

lint:
	@echo "🔍 Running golint..."
	go install golang.org/x/lint/golint@latest
	golint ./...

test:
	@echo "🧪 Running tests..."
	go test -v ./...

run:
	@echo "🚀 Running CLI..."
	./$(BIN_DIR)/$(APP_NAME) sync --app my-app --server https://argocd.example.com --username admin --password secret --insecure

clean:
	@echo "🧹 Cleaning..."
	rm -rf $(BIN_DIR)

docker-build:
	@echo "🐳 Building Docker image..."
	docker build -t $(IMAGE_NAME) .

docker-run:
	@echo "🐳 Running Docker container..."
	docker run --rm $(IMAGE_NAME) sync --app my-app --server https://argocd.example.com --username admin --password secret --insecure

scan: docker-build
	@echo "🔎 Scanning Docker image with Trivy..."
	trivy image --no-progress --severity HIGH,CRITICAL $(IMAGE_NAME)
