
SHELL = bash
BIN_DIR = bin
BUILD_TIME=$(shell sh -c 'date +%FT%T%z')
VERSION := $(shell sh -c 'git describe --always --tags')
BRANCH := $(shell sh -c 'git rev-parse --abbrev-ref HEAD')
COMMIT := $(shell sh -c 'git rev-parse --short HEAD')
GIT_SHORT_HASH := $(shell sh -c 'git rev-parse --short HEAD')
LDFLAGS=-ldflags "-s -w -X main.version=$(VERSION) -X main.commit=$(COMMIT) -X main.branch=$(BRANCH) -X main.buildDate=$(BUILD_TIME)"
BUILD_TAGS=-tags aws_lambda
TIMESTAMP := $(shell date +'%s')
ROOT_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

LINT_TOOL=$(shell go env GOPATH)/bin/golangci-lint
LINT_VERSION=v1.44.0


$(LINT_TOOL):
	@echo "Installing golangci linter version $(LINT_VERSION)..."
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(shell go env GOPATH)/bin $(LINT_VERSION)

.PHONY: go-lint
go-lint:
	@echo "Running lint..." && $(LINT_TOOL) --version && $(LINT_TOOL) run --allow-parallel-runners --config=.golangci.yaml ./... && echo "Congratulations! Lint check passed."

.PHONY: check-headers
check-headers:
	@echo "Running license header check..."
	@$(ROOT_DIR)/check-headers.sh

.PHONY: lint
lint: go-lint check-headers

.PHONY: dependencies deps
dependencies: deps

deps:
	@go env -w GOPRIVATE=github.com/aharal/*
	@go mod download

.PHONY: build
build: 
	@echo "Building Application..."
	env GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o $(BIN_DIR)/server main.go
