# Voiceflow CLI API Documentation

This directory contains the auto-generated OpenAPI/Swagger documentation for the Voiceflow CLI API server.

## 📁 Files

- **`docs.go`** - Generated Go code with embedded OpenAPI specification
- **`swagger.json`** - OpenAPI 2.0 specification in JSON format
- **`swagger.yaml`** - OpenAPI 2.0 specification in YAML format
- **`docs.go.backup`** - Backup of the previous manual docs.go file (for reference)

## 🔄 Auto-Generation

The documentation is **automatically generated** from the Go code annotations using [swaggo/swag](https://github.com/swaggo/swag).

### Generate Documentation

```bash
# Generate docs using Make
make docs

# Or generate docs directly
./scripts/generate-docs.sh

# Or use swag directly
swag init --generalInfo server/server.go --output server/docs
```

### Watch for Changes

```bash
# Auto-regenerate docs when code changes (requires fswatch)
make watch-docs
```

## 📖 Viewing Documentation

### Swagger UI (Recommended)

Start the API server and visit the Swagger UI:

```bash
# Start development server with Swagger UI
make dev

# Or start server directly
go run . server --debug
```

Then visit: **http://localhost:8080/swagger/index.html**

### Static Files

You can also serve the static files directly:

```bash
# Serve documentation files locally
make serve-docs
```

Then visit:
- JSON: http://localhost:8000/swagger.json
- YAML: http://localhost:8000/swagger.yaml

## 🏷️ API Annotations

The documentation is generated from Go annotations in the source code:

### General API Info

Located in `server/server.go`:

```go
// @title Voiceflow CLI API
// @version 1.0
// @description API server for Voiceflow CLI test execution and management
// @host localhost:8080
// @BasePath /
```

### Endpoint Annotations

Located in `server/handlers/handlers.go`:

```go
// @Summary Execute a test suite
// @Description Execute a Voiceflow test suite and return execution ID
// @Tags tests
// @Accept json
// @Produce json
// @Param request body TestExecutionRequest true "Test execution request"
// @Success 202 {object} TestExecutionResponse
// @Failure 400 {object} ErrorResponse
// @Router /api/v1/tests/execute [post]
```

## 🔧 Customization

### Modify API Information

Edit the annotations in `server/server.go` to change:
- API title and description
- Contact information
- License information
- Host and base path

### Add New Endpoints

1. Add handler function with proper annotations in `server/handlers/handlers.go`
2. Register the route in `server/server.go`
3. Run `make docs` to regenerate documentation

### Modify Response Models

Update the struct definitions in `server/handlers/handlers.go` and add appropriate JSON tags:

```go
type MyResponse struct {
    Field1 string `json:"field1" example:"example value"`
    Field2 int    `json:"field2" example:"42"`
}
```

## 🚀 Integration

The generated documentation is automatically integrated into the server:

1. **Import**: The `docs` package is imported in `server/server.go`
2. **Registration**: Swagger info is registered via `init()` function
3. **Serving**: Swagger UI is served at `/swagger/*` when enabled

## ⚡ Quick Commands

```bash
# Install required tools
make install-tools

# Generate documentation
make docs

# Validate generated OpenAPI spec
make validate-docs

# Start development server with docs
make dev

# Build and run
make build && ./voiceflow-cli server
```

## 📚 Resources

- [Swagger/OpenAPI Specification](https://swagger.io/specification/)
- [swaggo/swag Documentation](https://github.com/swaggo/swag)
- [Gin-Swagger Integration](https://github.com/swaggo/gin-swagger)
- [OpenAPI Examples](https://swagger.io/docs/specification/basic-structure/)

---

> **Note**: This documentation is **auto-generated**. Do not edit `docs.go`, `swagger.json`, or `swagger.yaml` manually. Instead, modify the source code annotations and regenerate using `make docs`.
