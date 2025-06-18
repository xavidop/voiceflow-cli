.PHONY: help build test clean docs serve-docs dev install-tools

# Default target
help: ## Show this help message
	@echo "Available targets:"
	@awk 'BEGIN {FS = ":.*##"; OFS = ": "} /^[a-zA-Z_-]+:.*##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

# Build the application
build: ## Build the voiceflow-cli binary
	@echo "🔨 Building voiceflow-cli..."
	go build -o voiceflow-cli .

# Run tests
test: ## Run all tests
	@echo "🧪 Running tests..."
	go test -v ./...

# Clean build artifacts
clean: ## Clean build artifacts and generated files
	@echo "🧹 Cleaning up..."
	rm -f voiceflow-cli
	rm -rf dist/

# Install development tools
install-tools: ## Install required development tools
	@echo "🔧 Installing development tools..."
	go install github.com/swaggo/swag/cmd/swag@latest
	go install github.com/go-swagger/go-swagger/cmd/swagger@latest
	@echo "✅ Development tools installed"

# Generate API documentation
docs: ## Generate OpenAPI/Swagger documentation
	@echo "📚 Generating API documentation..."
	./scripts/cmd_docs.sh

# Serve documentation locally (requires Python)
serve-docs: docs ## Generate and serve documentation locally
	@echo "🌐 Starting documentation server..."
	@if command -v python3 >/dev/null 2>&1; then \
		cd server/docs && python3 -m http.server 8000 & \
		echo "📖 Documentation server started at http://localhost:8000"; \
		echo "   • swagger.json: http://localhost:8000/swagger.json"; \
		echo "   • swagger.yaml: http://localhost:8000/swagger.yaml"; \
		echo "Press Ctrl+C to stop the server"; \
		wait; \
	else \
		echo "❌ Python3 not found. Please install Python3 to serve docs locally."; \
		echo "   Alternatively, run 'make dev' to start the API server with Swagger UI"; \
	fi

# Development server with live reload
dev: docs ## Build and run development server with Swagger UI
	@echo "🚀 Starting development server..."
	go run . server --debug --port 8080
	@echo "🎉 Server started at http://localhost:8080"
	@echo "📖 Swagger UI available at http://localhost:8080/swagger/index.html"

# Build for multiple platforms
build-all: ## Build for multiple platforms
	@echo "🏗️  Building for multiple platforms..."
	GOOS=linux GOARCH=amd64 go build -o dist/voiceflow-cli-linux-amd64 .
	GOOS=darwin GOARCH=amd64 go build -o dist/voiceflow-cli-darwin-amd64 .
	GOOS=darwin GOARCH=arm64 go build -o dist/voiceflow-cli-darwin-arm64 .
	GOOS=windows GOARCH=amd64 go build -o dist/voiceflow-cli-windows-amd64.exe .
	@echo "✅ Multi-platform builds completed in dist/"

# Run linting
lint: ## Run linting
	@echo "🔍 Running linters..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run; \
	else \
		echo "⚠️  golangci-lint not found. Install it with: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"; \
	fi

# Format code
fmt: ## Format Go code
	@echo "✨ Formatting code..."
	go fmt ./...
	goimports -w .

# Watch for changes and regenerate docs
watch-docs: ## Watch for changes and regenerate docs automatically
	@echo "👀 Watching for changes to regenerate docs..."
	@if command -v fswatch >/dev/null 2>&1; then \
		fswatch -o server/handlers/ server/server.go | xargs -n1 -I{} make docs; \
	elif command -v inotifywait >/dev/null 2>&1; then \
		while inotifywait -r -e modify server/handlers/ server/server.go; do make docs; done; \
	else \
		echo "❌ Neither fswatch nor inotifywait found."; \
		echo "   Install fswatch (macOS: brew install fswatch) or inotify-tools (Linux)"; \
	fi

# Validate OpenAPI spec
validate-docs: docs ## Validate generated OpenAPI specification
	@echo "✅ Validating OpenAPI specification..."
	@if command -v swagger >/dev/null 2>&1; then \
		swagger validate server/docs/swagger.yaml; \
	else \
		echo "⚠️  swagger CLI not found. Install it with: go install github.com/go-swagger/go-swagger/cmd/swagger@latest"; \
		echo "   Using basic validation..."; \
		@if [ -f server/docs/swagger.json ]; then \
			echo "📄 swagger.json exists and is valid JSON"; \
		else \
			echo "❌ swagger.json not found or invalid"; \
		fi \
	fi
