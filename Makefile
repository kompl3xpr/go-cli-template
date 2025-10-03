# Makefile

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOFMT=gofmt
GOLINT=golangci-lint
GOBIN?=$$(go env GOPATH)/bin

# Build parameters
EXEC_PATH=./bin/
BINARY=go-cli-template
ifeq ($(OS),Windows_NT)
	BINARY := $(BINARY).exe
endif

# Test parameters
TEST_TIMEOUT=240s
TEST_FLAGS=-race -count=1 -v
TEST_SHORT_FLAGS=-short $(TEST_FLAGS)

# Build targets
.PHONY: all build run clean test test-integration fmt lint install-go-test-coverage check-coverage

# Default target
all: clean build test

# Build the main project
build:
	@echo "Building..."
	$(GOBUILD) -v -o $(EXEC_PATH) ./...

run: build
	@echo "Running..."
	cd $(EXEC_PATH) && "./$(BINARY)"

# Run tests with short flag (excludes long-running tests)
test:
	@echo "Running tests..."
	$(GOTEST) $(TEST_SHORT_FLAGS) -timeout $(TEST_TIMEOUT) ./...

# Run all tests
test-integration:
	@echo "Running all tests..."
	$(GOTEST) $(TEST_FLAGS) -timeout $(TEST_TIMEOUT) ./...

custom-gcl:
	@echo "Checking if golangci-lint is installed..."
	@which $(GOLINT) > /dev/null || (echo "golangci-lint is not installed. Please install it first." && exit 1)
	@echo "Building custom linter plugin..."
	$(GOLINT) custom

# Format the code
fmt:
	@echo "Formatting code..."
	$(GOFMT) -s -w .

# Run lint
lint:
	@echo "Running lint..."
	$(GOLINT) run

# Clean build artifacts
clean:
	@echo "Cleaning..."
	$(GOCLEAN)

install-go-test-coverage:
	go install github.com/vladopajic/go-test-coverage/v2@latest

check-coverage: install-go-test-coverage
	go test ./... -coverprofile=./cover.out -covermode=atomic -coverpkg=./...
	${GOBIN}/go-test-coverage --config=./.testcoverage.yml
