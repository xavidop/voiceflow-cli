# Audio input

## Suite file

```yaml
# suite.yaml

name: Example Suite
description: Suite used as an example
projectId: test-cx-346408
locationId: us-central1
agentName: test-agent
tests:
  - id: test_id
    file: ./test.yaml
```

## Test file

```yaml
# test.yaml

name: Example test
description: These are some tests
localeId: en
interactions:
  - id: test_1
    user: 
      type: text
      text: hi
    agent:
      validate:
        - type: contains
          value: hello
```

