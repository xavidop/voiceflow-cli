# Usage Examples

## Using curl

### 1. Start a test execution
```bash
curl -X POST http://localhost:8080/api/v1/tests/execute \
  -H "Content-Type: application/json" \
  -d '{"suites_path": "/path/to/your/suite.yaml"}'
```

### 2. Check test status
```bash
curl http://localhost:8080/api/v1/tests/status/YOUR_EXECUTION_ID
```

### 3. Health check
```bash
curl http://localhost:8080/health
```

## Using JavaScript/fetch

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

## Using Python requests

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
