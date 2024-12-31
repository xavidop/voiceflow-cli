# Authentication

`voiceflow-cli` uses Voiceflow APIs. To interact with your Vocieflow projects you will need a [Voiceflow API Key](https://docs.voiceflow.com/reference/authentication). You can get your API Key in your Voiceflow project > Integration page. You can pass the API Key to the CLI using the `--voiceflow-api-key` flag or by setting the `VF_API_KEY` environment variable.

The `voiceflow-cli` source code is open source, you can check it out [here](https://github.com/xavidop/voiceflow-cli) to learn more about the actions the tool performs.

## Base URL

The base URL for the Voiceflow API is `https://general-runtime.voiceflow.com`. This is the default value for the `--voiceflow-base-url` flag. If you are using a different Voiceflow environment, you can pass the base URL using the `--voiceflow-base-url` flag.