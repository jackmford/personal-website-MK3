# Makefile for building and running the Go application

APP_NAME := personal-website
GO_FILES := $(shell find . -type f -name '*.go')
TEMPLATE_FILES := $(shell find ui/html -type f -name '*.tmpl')
STATIC_FILES := $(shell find ui/static -type f)
BLOG_FILES := $(wildcard ui/content/blog/*.md)
BUILD_DIR := bin
VERSION := $(shell git describe --tags --always --dirty)

.PHONY: all build run clean test lint fmt tidy

all: build

# Development targets
build: tidy $(GO_FILES) $(TEMPLATE_FILES) $(STATIC_FILES)
	@echo "Building $(APP_NAME) version $(VERSION)..."
	@mkdir -p $(BUILD_DIR)
	go build -ldflags "-X main.version=$(VERSION)" -o $(BUILD_DIR)/$(APP_NAME) ./cmd/web/main.go
	@echo "Build complete. Binary located at $(BUILD_DIR)/$(APP_NAME)."

run: build
	@echo "Running $(APP_NAME)..."
	@$(BUILD_DIR)/$(APP_NAME)

test:
	@echo "Running tests..."
	go test -v ./...

lint:
	@echo "Running linter..."
	golangci-lint run

fmt:
	@echo "Formatting code..."
	go fmt ./...

tidy:
	@echo "Tidying dependencies..."
	go mod tidy

# Cleanup
clean:
	@echo "Cleaning up..."
	@rm -rf $(BUILD_DIR)
	@echo "Clean complete."

# Development server with hot reload
dev: build
	@echo "Starting development server with hot reload..."
	@air

# Docker targets
docker-build:
	@echo "Building Docker image..."
	docker build -t $(APP_NAME):$(VERSION) .

docker-run:
	@echo "Running Docker container..."
	docker run -p 4000:4000 $(APP_NAME):$(VERSION)

# Help target
help:
	@echo "Available targets:"
	@echo "  build      - Build the application"
	@echo "  run        - Run the application"
	@echo "  test       - Run tests"
	@echo "  lint       - Run linter"
	@echo "  fmt        - Format code"
	@echo "  tidy       - Tidy dependencies"
	@echo "  clean      - Clean build artifacts"
	@echo "  dev        - Run development server with hot reload"
	@echo "  docker-build - Build Docker image"
	@echo "  docker-run   - Run Docker container"


