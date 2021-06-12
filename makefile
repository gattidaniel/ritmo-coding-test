all: fmt linter test
	@echo "==> Finished successfully..."

test:
	@echo "==> Running tests..."
	go test ./...

fmt:
	@echo "==> Running format..."
	go fmt ./...

linter:
	@echo "==> Running linter..."
	golangci-lint run ./...
