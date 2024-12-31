# Tests

## Reference

A test is a yaml file with the following structure:

```yaml
# test.yaml

# Name of the test.
name: Example test
# Brief description of the test.
description: These are some tests

# A interactions is the test itself: given an input, you will validate the agent response returned by Voiceflow
# You can have multiple interactions defined
interactions:
  # The ID of the interactions
  - id: test_1
    user:
      # the input type
      # it could be text, audio or prompt
      type: text
      # The input itself in text format. For type: audio, you have to specify the audio file.
      text: I want 3 pizzas
    agent:
      validate:
        # String validation to check if the response returned by Voiceflow is correct
        - type: contains
          value: pizza

  - id: test_2
    user: 
      type: text
      text: hi
    agent:
      # example with a traceType validation
      validate:
        - type: traceType
          value: speak

  - id: test_3
    user: 
      type: text
      audio: hello
    agent:
      # example with a regexp validation
      validate:
        - type: regexp
          value: '/my-regex/'
```

## Input types

### Text input

The input text is the simplest one. You just have to specify the text you want to send to Voiceflow. Make sure that the text is in the language you specified in the `localeId` field. to use this type you have to set the `type` field to `text` and the `text` field to the text you want to send.



## Validation types

### Contains

The contains validation type is the simplest one. It just checks if the response returned by the Voiceflow agent contains the value specified in the `value` field. To use this type you have to set the `type` field to `contains` and the `value` field to the value you want to check:

```yaml
validate:
  # String validation to check if the response returned by Voiceflow is correct
  - type: contains
    value: pizza
```

The `contains` validation has its own options:

```yaml
validate:
  # String validation to check if the response returned by Voiceflow is correct
  - type: contains
    value: pizza
```

If you set the `casesensitive` field to `true`, the validation will be case sensitive. By default, it is set to `false`.

### Equals

The equals validation type is a little bit more complex. It checks if the response returned by the Voiceflow agent is equal to the value specified in the `value` field. To use this type you have to set the `type` field to `equals` and the `value` field to the value you want to check:

```yaml
validate:
  # String validation to check if the response returned by Voiceflow is correct
  - type: equals
    value: Here you have 3 pizzas
```

The `equals` validation has its own options:

```yaml
validate:
  # String validation to check if the response returned by Voiceflow is correct
  - type: equals
    value: Here you have 3 pizzas
```

If you set the `casesensitive` field to `true`, the validation will be case sensitive. By default, it is set to `false`.

### Regexp

The regexp validation type is the most complex one. It checks if the response returned by the Voiceflow agent matches the regexp specified in the `value` field. To use this type you have to set the `type` field to `regexp` and the `value` field to the regular expression you want to check:

```yaml
validate:
  # String validation to check if the response returned by Voiceflow is correct
  - type: regexp
    value: '/Here you have \d pizzas/'
```

The `regexp` validation has its own options:

```yaml
validate:
  # String validation to check if the response returned by Voiceflow is correct
  - type: regexp
    value: '/Here you have \d pizzas/'
```

If you set the `findinsubmatches` field to `true`, the validation will check if the regexp matches any of the submatches. By default, it is set to `false`.

## JSON Schema

`voiceflow-cli` also has a [jsonschema](http://json-schema.org/draft/2020-12/json-schema-validation.html) file, which you can use to have better
editor support:

```sh
https://voiceflow.xavidop.me/static/conversationtest.json
```

You can also specify it in your `yml` config files by adding a
comment like the following:
```yaml
# yaml-language-server: $schema=https://voiceflow.xavidop.me/static/conversationtest.json
```
