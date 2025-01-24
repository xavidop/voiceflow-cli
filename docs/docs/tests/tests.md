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

  - id: test_4
    user: 
      type: text
      audio: hello
    agent:
      # example with a similarity validation
      validate:
        - type: similarity
          similarityConfig:
            provider: 'openai'
            model: 'gpt-4o'
            temperature: 0.8
            top_k: 5
            top_p: 0.9
            similarityThreshold: 0.5
          values:
            - 'hi'
            - 'Hello'
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

### Equals

The equals validation type is a little bit more complex. It checks if the response returned by the Voiceflow agent is equal to the value specified in the `value` field. To use this type you have to set the `type` field to `equals` and the `value` field to the value you want to check:

```yaml
validate:
  # String validation to check if the response returned by Voiceflow is correct
  - type: equals
    value: Here you have 3 pizzas
```

### Regexp

The regexp validation type is the most complex one. It checks if the response returned by the Voiceflow agent matches the regexp specified in the `value` field. To use this type you have to set the `type` field to `regexp` and the `value` field to the regular expression you want to check:

```yaml
validate:
  # String validation to check if the response returned by Voiceflow is correct
  - type: regexp
    value: '/Here you have \d pizzas/'
```

### TraceType
The traceType validation type checks if the response returned by the Voiceflow agent has the trace type specified in the `value` field. To use this type you have to set the `type` field to `traceType` and the `value` field to the trace type you want to check:

```yaml
validate:
  # String validation to check if the response returned by Voiceflow is correct
  - type: traceType
    value: speak
```

### Similarity
The similarity validation type checks if the response returned by the Voiceflow agent is similar to the values specified in the `values` field. To use this type you have to set the `type` field to `similarity` and the `values` field to the values you want to check:

```yaml
validate:
  # String validation to check if the response returned by Voiceflow is correct
  - type: similarity
    similarityConfig:
      provider: 'openai'
      model: 'gpt-4o'
      temperature: 0.8
      top_k: 5
      top_p: 0.9
      similarityThreshold: 0.5
    values:
      - 'hi'
      - 'Hello'
```

You can also use the `similarityConfig` field to specify the similarity configuration. The `provider` field specifies the similarity provider you want to use. The `model` field specifies the model you want to use. The `temperature` field specifies the temperature you want to use. The `top_k` field specifies the top k you want to use. The `top_p` field specifies the top p you want to use. The `similarityThreshold` field specifies the similarity threshold you want to use.

The only provider available for now is `openai`.

For LLM Providers authentication please check the [Authentication](/overview/authentication) page.


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
