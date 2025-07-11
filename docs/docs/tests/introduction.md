# Conversation Profiler

## What is this?

<p align="center">
  <img alt="Flow" src="/images/flow.png" style="height:512px;width:512px" />
</p>

Use the Conversation Profiler to test user utterances and improve your agent's interaction model.

The Conversation Profiler supports **two distinct testing approaches** to validate your agent's conversation flow:

### 🔧 Traditional Interaction-Based Testing
Test the conversation flow with **predefined interactions** where you send specific user utterances to your agent and validate exact responses. This approach is ideal for:

- **Regression testing** to ensure specific responses remain consistent
- **Validation of exact conversation flows** with predetermined inputs and outputs
- **Quality assurance** for specific features or conversation paths

#### Reference

It is important to know which [suites](/tests/suites) and [tests](/tests/interaction-tests) you can build. Because of that, you can find the entire reference on the [Reference](/tests/suites) page. Suites and test are defined as `yaml` files.

### 🤖 Agent-to-Agent Testing
Simulate **realistic conversations** using AI-powered agents that interact naturally with your Voiceflow agent to achieve specific goals. This approach is ideal for:

- **End-to-end conversation testing** with natural, adaptive interactions
- **User behavior simulation** where the AI agent responds dynamically like real users
- **Goal-oriented testing** to ensure your agent can handle varied conversation paths

Both testing approaches can be run in your CI/CD pipelines and include additional features beyond the Voiceflow console's Test Agent feature. Every suite is executed in the same Voiceflow user's session.

All of the commands that are available in `voiceflow-cli` to execute the Conversation profiler are located within the [`voiceflow test` subcommand](/cmd/voiceflow_test).

#### Reference

It is important to know which [suites](/tests/suites) and [tests](/tests/agent-to-agent-tests) you can build. Because of that, you can find the entire reference on the [Reference](/tests/suites) page. Suites and test are defined as `yaml` files.

## Execution

The `voiceflow-cli` has a command that allows you to run these suites from your terminal or from your CI pipelines.

To execute a suite, you can run the `voiceflow test execute [suitesPath]` command. For the usage, please refer to this [page](/cmd/voiceflow_test).

## Examples

You can find some useful examples on our [GitHub repo](https://github.com/xavidop/voiceflow-cli/tree/main/examples) and the [Examples](/tests/examples) page.


## Execution Example

Here is a simple example of the `voiceflow test execute` command:

```sh
voiceflow test execute examples/test/
```

The above command will give you output similar to the following:

```sh
$ voiceflow test execute examples/test/

Dec 31 10:54:01.664 [INFO] Suite: Example Conversation Profiler Suite
Description: Suite used as an example
Environment: development
Dec 31 10:54:01.664 [INFO] Running Tests:
Dec 31 10:54:01.664 [INFO] Running Test ID: Example test
Dec 31 10:54:01.664 [INFO] Interaction ID: test_1_1
Dec 31 10:54:01.664 [INFO]      Interaction Request Type: launch
Dec 31 10:54:02.693 [INFO]      Interaction Response Type: text
Dec 31 10:54:02.693 [INFO]      Interaction Response Message: Hey there! 🌟 Welcome to the Isla Experience! I’m like a warm cup of cocoa on a chilly day—sweet, comforting, and maybe a little too hot if you’re not careful! How’s your day going?
Dec 31 10:54:02.693 [INFO] All validations passed for Interaction ID: test_1_1
Dec 31 10:54:02.693 [INFO] Interaction ID: test_1_2
Dec 31 10:54:02.693 [INFO]      Interaction Request Type: text
Dec 31 10:54:02.693 [INFO]      Interaction Request Payload: I am doing well
Dec 31 10:54:03.889 [INFO]      Interaction Response Type: text
Dec 31 10:54:03.889 [INFO]      Interaction Response Message: Awesome! Glad to hear it! Are you riding the wave of good vibes, or did you just find a hidden stash of chocolate? 🍫 Either way, I’m here for it! What’s been the highlight of your day so far?
Dec 31 10:54:03.889 [INFO] All validations passed for Interaction ID: test_1_2
Dec 31 10:54:03.889 [INFO] Interaction ID: test_1_3
Dec 31 10:54:03.889 [INFO]      Interaction Request Type: text
Dec 31 10:54:03.889 [INFO]      Interaction Request Payload: I have been working very hard
Dec 31 10:54:06.090 [INFO]      Interaction Response Type: text
Dec 31 10:54:06.091 [INFO]      Interaction Response Message: Ah, the classic “I’m working hard” routine! It’s like a superhero origin story, but instead of gaining superpowers, you just gain a lot of coffee stains and a questionable relationship with your chair. What kind of work are you diving into?
Dec 31 10:54:06.091 [INFO] All validations passed for Interaction ID: test_1_3
```

!!! info "Are you running this command in a CI/CD pipeline?"
    If this is the case, we recommend that you set the `--output-format` parameter to `json`.
