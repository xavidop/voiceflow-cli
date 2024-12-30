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
# You can have multiple tests defined in separated files
tests:
  # ID of the test.
  - id: test_id
    # File where the test specification is located
    file: ./test.yaml
```

It has the same structure as the NLU Profiler suite.

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
