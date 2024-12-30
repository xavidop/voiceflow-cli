# Conversation profiler

## What is this?

<p align="center">
  <img alt="Flow" src="/images/flow.png" style="height:512px;width:512px" />
</p>

Use the Conversation Profiler to test user utterances and improve your agent's interaction model.

With the Conversation Profiler, you can test the conversation flow of your agent. This means that you can send user's utterances to your agent and check if the agent is responding as expected (interactions). This feature is also useful when you have generative AI models in your agent, and you want to test if the agent is responding as expected. It is similar to the Voiceflow console's Test Agent feature, but with the Conversation Profiler, you can run your tests in your CI/CD pipelines and also it has additional features. Every suite is executed in the same Voiceflow user's session.

All of the commands that are available in `voiceflow-cli` to execute the Conversation profiler are located within the [`voiceflow test` subcommand](/cmd/voiceflow_test).

## Reference

It is important to know which [suites](/tests/suites) and [tests](/tests/tests) you can build. Because of that, you can find the entire reference on the [Reference](/tests/suites) page. Suites and test are defined as `yaml` files.

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

[INFO] Running suite: Example Conversation Profiler Suite
[INFO][test-file:test_1][interaction:test_1_1][input:prompt] User> Hello! (auto-generated from prompt: "give me a one line hello")
[INFO][test-file:test_1][interaction:test_1_1][input:prompt] Agent> Hi! How are you doing?
[INFO][test-file:test_1][interaction:test_1_1][validation:equals] Validation with value "hi" 
[INFO][test-file:test_1][interaction:test_1_1][validation:equals] Validation configuration: {CaseSensitive:false}
[INFO][test-file:test_1][interaction:test_1_2][input:text] User> hi
[INFO][test-file:test_1][interaction:test_1_2][input:text] Agent> Hi! How are you doing?
[INFO][test-file:test_1][interaction:test_1_2][validation:contains] Validation with value "hi" 
[INFO][test-file:test_1][interaction:test_1_2][validation:contains] Validation configuration: {CaseSensitive:false}
[INFO][test-file:test_1][interaction:test_1_3][input:text] User> hello
[INFO][test-file:test_1][interaction:test_1_3][input:text] Agent> Hi! How are you doing?
[INFO][test-file:test_1][interaction:test_1_3][validation:equals] Validation with value "hello" 
[INFO][test-file:test_1][interaction:test_1_3][validation:equals] Validation configuration: {CaseSensitive:false}
[INFO][test-file:test_1][interaction:test_1_4][input:text] User> hi
[INFO][test-file:test_1][interaction:test_1_4][input:text] Agent> Hi! How are you doing?
[INFO][test-file:test_1][interaction:test_1_4][validation:similarity] Validation with value "hello" 
[INFO][test-file:test_1][interaction:test_1_4][validation:similarity] Validation configuration: {CaseSensitive:false InsertCost:0 DeleteCost:0 ReplaceCost:0}
```

!!! info "Are you running this command in a CI/CD pipeline?"
    If this is the case, we recommend that you set the `--output-format` parameter to `json`.
