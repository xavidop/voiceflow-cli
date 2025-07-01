# Audio input

## Suite file

```yaml
# suite.yaml

name: Example Suite
description: Suite used as an example
environmentName: production
tests:
  - id: test_id
    file: ./test.yaml
```

## Test file

```yaml
# test.yaml

name: Example test
description: These are some tests
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

