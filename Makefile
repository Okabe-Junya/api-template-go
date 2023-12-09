DEFAULT_GOAL := build

BIN := bin/server
MAIN := cmd/server/main.go

.PHONY: build
build:
	@echo "Building..."
	@go build -o $(BIN) $(MAIN)

.PHONY: run
run:
	@echo "Running..."
	@go run $(MAIN)

.PHONY: test
test:
	@echo "Testing..."
	@go test -v ./...

.PHONY: clean
clean:
	@echo "Cleaning..."
	@rm -rf $(BIN)
	@mkdir -p bin

.PHONY: lint
lint:
	@echo "Linting..."
	@golangci-lint run ./...

.PHONY: fmt
fmt:
	@echo "Formatting..."
	@go fmt ./...
