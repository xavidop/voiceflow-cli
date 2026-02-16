# Authentication

## Voiceflow API Key

`voiceflow-cli` uses Voiceflow APIs. To interact with your Vocieflow projects you will need a [Voiceflow API Key](https://docs.voiceflow.com/reference/authentication). You can get your API Key in your Voiceflow project > Integration page. You can pass the API Key to the CLI using the `--voiceflow-api-key` flag or by setting the `VF_API_KEY` environment variable. `voiceflow-cli` also works with `.env` files. You can create a `.env` file in the root of your project and add the `VF_API_KEY` variable to it.

The `voiceflow-cli` source code is open source, you can check it out [here](https://github.com/xavidop/voiceflow-cli) to learn more about the actions the tool performs.

## Base URL

The base URL for the Voiceflow API is `https://<api>.<subdomain>.voiceflow.com`. The default value is without subdomain: `https://<api>.voiceflow.com`. If you are using a different Voiceflow environment, you can pass the subdomain using the `--voiceflow-subdomain` flag.

## Open AI PI Key

`voiceflow-cli` uses Open AI APIs. To interact with Open AI you will need an API Key. You can get your API Key in your Open AI account. You can pass the API Key to the CLI using the `--openai-api-key` flag or by setting the `OPENAI_API_KEY` environment variable. `voiceflow-cli` also works with `.env` files. You can create a `.env` file in the root of your project and add the `OPENAI_API_KEY` variable to it.

## OpenAI Base URL (optional)

If you need to target a different OpenAI endpoint (for example EU data residency), you can override the base URL by using:

- CLI flag: `--openai-base-url`
- Environment variable: `OPENAI_BASE_URL`

Example:

```sh
voiceflow test execute evals --openai-base-url https://eu.api.openai.com/v1
```

If not provided, `voiceflow-cli` defaults to `https://api.openai.com/v1`.