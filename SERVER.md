# Voiceflow CLI API Server

The Voiceflow CLI now includes an HTTP API server that exposes test execution functionality as REST endpoints with auto-generated OpenAPI/Swagger documentation.

## Features

- **HTTP API**: Execute test suites via REST endpoints
- **Real-time Logging**: Capture and return test execution logs in API responses
- **OpenAPI/Swagger**: Auto-generated API documentation at `/swagger/index.html`
- **Asynchronous Execution**: Non-blocking test execution with status tracking
- **CORS Support**: Enable cross-origin requests for web frontends
- **Health Checks**: Built-in health check endpoints

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

## API Endpoints

### Health Check
```http
GET /health
```

### Execute Test Suite
```http
POST /api/v1/tests/execute
Content-Type: application/json

{
  "suites_path": "/path/to/your/suite.yaml"
}
```

**Response:**
```json
{
  "id": "123e4567-e89b-12d3-a456-426614174000",
  "status": "running",
  "started_at": "2023-01-01T00:00:00Z",
  "logs": ["Test execution started"]
}
```

### Get Test Status
```http
GET /api/v1/tests/status/{execution_id}
```

**Response:**
```json
{
  "id": "123e4567-e89b-12d3-a456-426614174000",
  "status": "completed",
  "started_at": "2023-01-01T00:00:00Z",
  "completed_at": "2023-01-01T00:05:00Z",
  "logs": [
    "Starting test suite execution...",
    "Suite path: /path/to/suite.yaml",
    "Running Test ID: example_test",
    "Test suite execution completed successfully"
  ]
}
```

### System Information
```http
GET /api/v1/system/info
```

**Response:**
```json
{
  "version": "1.0.0",
  "go_version": "go1.20.0",
  "os": "linux",
  "arch": "amd64"
}
```

## OpenAPI/Swagger Documentation

Once the server is running, you can access the interactive API documentation at:

```
http://localhost:8080/swagger/index.html
```

## Usage Examples

### Using curl

#### 1. Start a test execution
```bash
curl -X POST http://localhost:8080/api/v1/tests/execute \
  -H "Content-Type: application/json" \
  -d '{"suites_path": "/path/to/your/suite.yaml"}'
```

#### 2. Check test status
```bash
curl http://localhost:8080/api/v1/tests/status/YOUR_EXECUTION_ID
```

#### 3. Health check
```bash
curl http://localhost:8080/health
```

### Using JavaScript/fetch

```javascript
// Execute a test suite
const response = await fetch('http://localhost:8080/api/v1/tests/execute', {
  method: 'POST',
  headers: {
    'Content-Type': 'application/json',
  },
  body: JSON.stringify({
    suites_path: '/path/to/your/suite.yaml'
  })
});

const execution = await response.json();
console.log('Execution ID:', execution.id);

// Poll for status
const statusResponse = await fetch(`http://localhost:8080/api/v1/tests/status/${execution.id}`);
const status = await statusResponse.json();
console.log('Status:', status.status);
console.log('Logs:', status.logs);
```

### Using Python requests

```python
import requests
import time

# Execute a test suite
response = requests.post('http://localhost:8080/api/v1/tests/execute', json={
    'suites_path': '/path/to/your/suite.yaml'
})
execution = response.json()
print(f"Execution ID: {execution['id']}")

# Poll for completion
while True:
    status_response = requests.get(f"http://localhost:8080/api/v1/tests/status/{execution['id']}")
    status = status_response.json()
    
    print(f"Status: {status['status']}")
    
    if status['status'] in ['completed', 'failed']:
        print("Logs:")
        for log in status['logs']:
            print(f"  {log}")
        break
    
    time.sleep(1)
```

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

## Integration with CI/CD

The API server makes it easy to integrate Voiceflow test execution into CI/CD pipelines:

### GitHub Actions Example

```yaml
name: Voiceflow Tests
on: [push, pull_request]

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      
      - name: Start Voiceflow CLI Server
        run: |
          voiceflow server --port 8080 &
          sleep 5
        env:
          VF_API_KEY: ${{ secrets.VF_API_KEY }}
      
      - name: Run Tests
        run: |
          EXECUTION_ID=$(curl -s -X POST http://localhost:8080/api/v1/tests/execute \
            -H "Content-Type: application/json" \
            -d '{"suites_path": "./tests/suite.yaml"}' | jq -r '.id')
          
          # Poll for completion
          while true; do
            STATUS=$(curl -s http://localhost:8080/api/v1/tests/status/$EXECUTION_ID | jq -r '.status')
            if [ "$STATUS" = "completed" ]; then
              echo "Tests passed!"
              break
            elif [ "$STATUS" = "failed" ]; then
              echo "Tests failed!"
              exit 1
            fi
            sleep 2
          done
```

## Security Considerations

- The server runs on all interfaces (`0.0.0.0`) by default. In production, consider binding to specific interfaces.
- There is no built-in authentication. Consider adding a reverse proxy with authentication if needed.
- Test executions are stored in memory. Consider implementing persistent storage for production use.

## Troubleshooting

### Server Won't Start

1. Check if the port is already in use:
   ```bash
   lsof -i :8080
   ```

2. Try a different port:
   ```bash
   voiceflow server --port 9090
   ```

### API Returns 404

Ensure you're using the correct base path `/api/v1/` for API endpoints.

### Logs Not Appearing

Enable debug mode to see more detailed logging:
```bash
voiceflow server --debug
```
