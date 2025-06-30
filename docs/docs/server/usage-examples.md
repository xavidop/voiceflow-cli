# Usage Examples

## Using curl

### 1. Start a test execution
```bash
curl -X POST http://localhost:8080/api/v1/tests/execute \
  -H "Content-Type: application/json" \
  -d '{"api_key": "your_api_key (optional)","voiceflow_subdomain": "your_custom_subdomain (optional)","suite": {"name": "Example Suite","description": "Suite used as an example","environment_name": "production","tests": [{"id": "test_1","test": {"name": "Example test","description": "These are some tests","interactions": [{"id": "test_1_1","user": {"type": "text","text": "hi"},"agent": {"validate": [{"type": "contains","value": "hello"}]}}]}}]}}'
```

### 1b. Start a test execution with custom subdomain
```bash
curl -X POST http://localhost:8080/api/v1/tests/execute \
  -H "Content-Type: application/json" \
  -d '{
    "api_key": "VF.DM.YOUR_API_KEY",
    "voiceflow_subdomain": "staging-env",
    "suite": {
      "name": "Staging Environment Test",
      "description": "Testing against staging environment",
      "environment_name": "production",
      "tests": [
        {
          "id": "staging_test_1",
          "test": {
            "name": "Basic staging test",
            "description": "Test against staging subdomain",
            "interactions": [
              {
                "id": "staging_interaction_1",
                "user": {
                  "type": "launch"
                },
                "agent": {
                  "validate": [
                    {
                      "type": "contains",
                      "value": "welcome"
                    }
                  ]
                }
              }
            ]
          }
        }
      ]
    }
  }'
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
// Execute a test suite with custom subdomain
const response = await fetch('http://localhost:8080/api/v1/tests/execute', {
  method: 'POST',
  headers: {
    'Content-Type': 'application/json',
  },
  body: JSON.stringify({
    api_key: "VF.DM.YOUR_API_KEY",
    voiceflow_subdomain: "staging-env", // Optional: use custom subdomain
    suite: {
      name: "Example Suite",
      description: "Suite used as an example",
      environment_name: "production",
      tests: [
        {
          id: "test_1",
          test: {
            name: "Example test",
            description: "These are some tests",
            interactions: [
              {
                id: "test_1_1",
                user: {
                  type: "text",
                  text: "hi"
                },
                agent: {
                  validate: [
                    {
                      type: "contains",
                      value: "hello"
                    }
                  ]
                }
              }
            ]
          }
        }
      ]
    }
  })
});

const execution = await response.json();
console.log('Execution ID:', execution.id);

// Poll for status
const statusResponse = await fetch(`http://localhost:8080/api/v1/tests/status/${execution.id}`);
const status = await statusResponse.json();
console.log('Status:', status.status);
console.log('Logs:', status.logs);

// Example with multiple environments
const environments = [
  { name: "production", subdomain: "" }, // Use global subdomain
  { name: "staging", subdomain: "staging-env" },
  { name: "development", subdomain: "dev-env" }
];

for (const env of environments) {
  const envResponse = await fetch('http://localhost:8080/api/v1/tests/execute', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({
      api_key: "VF.DM.YOUR_API_KEY",
      voiceflow_subdomain: env.subdomain,
      suite: {
        name: `${env.name} Test Suite`,
        description: `Testing against ${env.name} environment`,
        environment_name: "production",
        tests: [/* your tests here */]
      }
    })
  });
  
  const envExecution = await envResponse.json();
  console.log(`${env.name} execution started:`, envExecution.id);
}
```

## Using Python requests

```python
import requests
import time

# Execute a test suite with custom subdomain
response = requests.post('http://localhost:8080/api/v1/tests/execute', json={
    'api_key': 'VF.DM.YOUR_API_KEY',
    'voiceflow_subdomain': 'staging-env',  # Optional: use custom subdomain
    'suite': {
        'name': 'Example Suite',
        'description': 'Suite used as an example',
        'environment_name': 'production',
        'tests': [
            {
                'id': 'test_1',
                'test': {
                    'name': 'Example test',
                    'description': 'These are some tests',
                    'interactions': [
                        {
                            'id': 'test_1_1',
                            'user': {
                                'type': 'text',
                                'text': 'hi'
                            },
                            'agent': {
                                'validate': [
                                    {
                                        'type': 'contains',
                                        'value': 'hello'
                                    }
                                ]
                            }
                        }
                    ]
                }
            }
        ]
    }
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

# Example: Testing multiple environments
environments = [
    {"name": "production", "subdomain": ""},  # Use global subdomain
    {"name": "staging", "subdomain": "staging-env"},
    {"name": "development", "subdomain": "dev-env"}
]

execution_ids = []

for env in environments:
    print(f"\nStarting test for {env['name']} environment...")
    
    response = requests.post('http://localhost:8080/api/v1/tests/execute', json={
        'api_key': 'VF.DM.YOUR_API_KEY',
        'voiceflow_subdomain': env['subdomain'],
        'suite': {
            'name': f"{env['name']} Test Suite",
            'description': f"Testing against {env['name']} environment",
            'environment_name': 'production',
            'tests': [
                # Your test definitions here
            ]
        }
    })
    
    execution = response.json()
    execution_ids.append((env['name'], execution['id']))
    print(f"{env['name']} execution started: {execution['id']}")

# Monitor all executions
print("\nMonitoring all executions...")
for env_name, execution_id in execution_ids:
    print(f"\nChecking {env_name} ({execution_id})...")
    status_response = requests.get(f"http://localhost:8080/api/v1/tests/status/{execution_id}")
    status = status_response.json()
    print(f"Status: {status['status']}")
```
