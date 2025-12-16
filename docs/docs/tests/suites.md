# Suites

## Reference

A suite is a yaml file with the following structure:

```yaml
# suite.yaml

# Name of the suite.
name: Example Suite
# Brief description of the suite.
description: Suite used as an example
# Environment name of your Voiceflow agent. It could be development, or production.
environmentName: development
# Optional: Create a new user session for each test (default: false)
# When enabled, each test will run with a fresh user session instead of sharing one session across all tests
newSessionPerTest: false
# You can have multiple tests defined in separated files
tests:
  # ID of the test.
  - id: test_id
    # File where the test specification is located
    file: ./test.yaml
```

It has the same structure as the NLU Profiler suite.

### Session Management

By default, all tests within a suite share the same user session (user ID). This means that:

- Variables set in one test persist to the next test
- The conversation context carries over between tests
- Tests are executed sequentially with the same user state

If you want each test to start with a fresh session, set `newSessionPerTest: true`. This will:

- Generate a new user ID for each test
- Clear all conversation context between tests
- Ensure tests are completely isolated from each other

## JSON Schema

`voiceflow-cli` also has a [jsonschema](http://json-schema.org/draft/2020-12/json-schema-validation.html) file, which you can use to have better
editor support:

```sh
https://voiceflow.xavidop.me/static/conversationsuite.json
```

You can also specify it in your `yml` config files by adding a
comment like the following:
```yaml
# yaml-language-server: $schema=https://voiceflow.xavidop.me/static/conversationsuite.json
```
