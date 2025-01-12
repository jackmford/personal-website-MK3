# Makefile for building and running the Go application

APP_NAME := personal-website 
GO_FILES := $(shell find . -type f -name '*.go')
BUILD_DIR := bin

.PHONY: all build run clean

all: build

build: $(GO_FILES)
	@echo "Building $(APP_NAME)..."
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(APP_NAME) ./cmd/web/main.go
	@echo "Build complete. Binary located at $(BUILD_DIR)/$(APP_NAME)."

run: build
	@echo "Running $(APP_NAME)..."
	@$(BUILD_DIR)/$(APP_NAME)

clean:
	@echo "Cleaning up..."
	@rm -rf $(BUILD_DIR)
	@echo "Clean complete."


