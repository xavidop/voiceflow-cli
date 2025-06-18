# API Endpoints

## Health Check
```http
GET /health
```

Returns the health status of the server.

## Execute Test Suite
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

Executes a test suite asynchronously and returns an execution ID for tracking.

## Get Test Status
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

Retrieves the current status and logs of a test execution.

## System Information
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

Returns system information about the running server instance.

## OpenAPI/Swagger Documentation

Once the server is running, you can access the interactive API documentation at:

```
http://localhost:8080/swagger/index.html
```
