# Traditional Interaction-Based Tests

## Overview
Test the conversation flow with **predefined interactions** where you send specific user utterances to your agent and validate exact responses.

## Reference

A traditional test is a YAML file with the following structure:

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

  - id: test_5
    user: 
      type: text
      text: 'myVariableValue1'
    agent:
      # example with a variable validation
      validate:
        - type: variable
          value: 'myVariableValue1'
          variableConfig:
            name: 'variableName1'
```

## Input types

### Text input

The input text is the simplest one. You just have to specify the text you want to send to Voiceflow. Make sure that the text is in the language you specified in the `localeId` field. to use this type you have to set the `type` field to `text` and the `text` field to the text you want to send.

```yaml
user:
  type: text
  text: I want 3 pizzas
```

### Launch input

The launch input is used to start a new conversation session. This is typically the first interaction in a test. To use this type you have to set the `type` field to `launch`.

```yaml
user:
  type: launch
```

### Event input

The event input allows you to send custom events to your Voiceflow agent. Events can be used to trigger specific flows or actions in your agent. To use this type you have to set the `type` field to `event` and the `event` field to the event name you want to send.

```yaml
user:
  type: event
  event: user_logged_in
```

Example with validation:

```yaml
- id: event_example
  user:
    type: event
    event: user_logged_in
  agent:
    validate:
      - type: contains
        value: "Welcome back!"
```

### Intent input

The intent input allows you to directly send an intent to your Voiceflow agent, bypassing NLU processing. This is useful when you have your own NLU matching or want to test specific intent handling. To use this type you have to set the `type` field to `intent` and provide an `intent` object with the intent name and optional entities.

```yaml
user:
  type: intent
  intent:
    name: order_pizza
    entities:
      - name: pizza_type
        value: pepperoni
      - name: size
        value: large
```

The `intent` object accepts the following properties:
- `name`: (Required) The name of the intent to trigger
- `entities`: (Optional) An array of entity objects with `name` and `value` fields

Examples:

```yaml
# Intent with entities
- id: intent_with_entities
  user:
    type: intent
    intent:
      name: order_pizza
      entities:
        - name: pizza_type
          value: pepperoni
        - name: size
          value: large
  agent:
    validate:
      - type: contains
        value: "one large pepperoni pizza"

# Intent without entities
- id: intent_no_entities
  user:
    type: intent
    intent:
      name: get_help
  agent:
    validate:
      - type: contains
        value: "How can I help you?"
```

### Button input

The button input allows you to simulate clicking a button that was presented in a previous choice/button response from your Voiceflow agent. This is useful for testing conversational flows that include button interactions. To use this type you have to set the `type` field to `button` and the `value` field to the button label you want to click.

```yaml
user:
  type: button
  value: Yes, continue
```

The button interaction automatically:
- Finds the matching button from the previous `choice` trace by its label
- Sends the complete button request (including path type and payload) back to Voiceflow
- Handles the button click as if a user clicked it in a real conversation

**Important**: A button interaction must follow an interaction that returned a `choice` trace type with buttons. The `value` must match the `label` field in one of the button's payload.

Example workflow:

```yaml
# First interaction receives a choice with buttons
- id: show_options
  user:
    type: text
    text: "Show me options"
  agent:
    validate:
      - type: traceType
        value: choice

# Second interaction clicks one of the buttons
- id: select_option
  user:
    type: button
    value: "Yes, continue"  # Must match button label
  agent:
    validate:
      - type: contains
        value: "Great! Continuing..."
```

Complete example:

```yaml
name: Button Interaction Example
description: Test demonstrating button click simulation

interactions:
  - id: launch_conversation
    user:
      type: launch
    agent:
      validate:
        - type: contains
          value: "Welcome"
        - type: traceType
          value: choice
  
  - id: click_yes_button
    user:
      type: button
      value: "Yes"  # Clicks the button with label "Yes"
    agent:
      validate:
        - type: contains
          value: "You selected yes"
```



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


### Variable
The variable validation type checks if a variable in the Voiceflow agent has the expected value. To use this type you have to set the `type` field to `variable`, the `value` field to the expected value, and provide a `variableConfig` object with the variable details:

```yaml
validate:
  # Variable validation to check if a variable has the expected value
  - type: variable
    value: 'myVariableValue1'
    variableConfig:
      name: 'variableName1'
```

The `variableConfig` object accepts the following properties:
- `name`: (Required) The name of the variable to validate
- `jsonPath`: (Optional) A JSONPath expression to extract nested values from JSON/object variables

Examples:

```yaml
validate:
  # Simple variable validation
  - type: variable
    value: 'myVariableValue1'
    variableConfig:
      name: 'variableName1'
  
  # Multiple variable validations
  - type: variable
    value: 'myVariableValue2'
    variableConfig:
      name: 'variableName2'
  
  # Variable validation with JSONPath if the variable is a JSON/object
  - type: variable
    value: 'myVariableValue3'
    variableConfig:
      name: 'variableName3'
      jsonPath: '$.hello'
```

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
