# API Server Introduction

The Voiceflow CLI now includes an HTTP API server that exposes test execution functionality as REST endpoints with auto-generated OpenAPI/Swagger documentation.

## Features

- **HTTP API**: Execute test suites via REST endpoints
- **Real-time Logging**: Capture and return test execution logs in API responses
- **OpenAPI/Swagger**: Auto-generated API documentation at `/swagger/index.html`
- **Asynchronous Execution**: Non-blocking test execution with status tracking
- **CORS Support**: Enable cross-origin requests for web frontends
- **Health Checks**: Built-in health check endpoints

## OpenAPI Specifications

The server provides OpenAPI specifications in both YAML and JSON formats:

- **YAML Format**: Available at [`/static/swagger.yaml`](/static/swagger.yaml)
- **JSON Format**: Available at [`/static/swagger.json`](/static/swagger.json)

These specifications can be used to generate client libraries, import into API testing tools, or integrate with other OpenAPI-compatible tooling.

## Starting the Server

### Basic Usage

```bash
# Start server on default port (8080)
voiceflow server

# Start server on custom port
voiceflow server --port 9090

# Start server with debug mode
voiceflow server --debug

# Start server with custom host
voiceflow server --host 127.0.0.1 --port 8080
```

### Command Line Options

| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--port` | `-p` | `8080` | Port to run the server on |
| `--host` | `-H` | `0.0.0.0` | Host to bind the server to |
| `--debug` | `-d` | `false` | Enable debug mode |
| `--cors` | | `true` | Enable CORS middleware |
| `--swagger` | | `true` | Enable Swagger documentation endpoint |

## Configuration

### Environment Variables

The server respects all existing Voiceflow CLI environment variables:

- `VF_API_KEY`: Voiceflow API Key
- `OPENAI_API_KEY`: OpenAI API Key (for similarity validations)

### CORS Configuration

CORS is enabled by default. To disable CORS:

```bash
voiceflow server --cors=false
```

### Debug Mode

Enable debug mode for detailed logging:

```bash
voiceflow server --debug
```
