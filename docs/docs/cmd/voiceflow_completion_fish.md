# voiceflow completion fish

Generate the autocompletion script for fish

## Synopsis

Generate the autocompletion script for the fish shell.

To load completions in your current shell session:

	voiceflow completion fish | source

To load completions for every new session, execute once:

	voiceflow completion fish > ~/.config/fish/completions/voiceflow.fish

You will need to start a new shell for this setup to take effect.


```
voiceflow completion fish [flags]
```

## Options

```
  -h, --help              help for fish
      --no-descriptions   disable completion descriptions
```

## Options inherited from parent commands

```
  -z, --open-api-key string              Open API Key (optional)
      --openai-base-url string           OpenAI API base URL override, e.g. https://eu.api.openai.com/v1 (optional)
  -o, --output-format string             Output Format. Options: text, json. Default: text (optional) (default "text")
      --show-tester-messages             Show tester agent messages in agent-to-agent tests (optional) (default true)
  -u, --skip-update-check                Skip the check for updates check run before every command (optional)
  -v, --verbose                          verbose error output (with stack trace) (optional)
      --voiceflow-analytics-url string   Custom base URL for the Voiceflow analytics API (optional)
  -x, --voiceflow-api-key string         Voiceflow API Key (optional)
      --voiceflow-api-url string         Custom base URL for the Voiceflow API (creator-api), (optional)
      --voiceflow-runtime-url string     Custom base URL for the Voiceflow general-runtime (optional)
  -b, --voiceflow-subdomain string       Voiceflow Base URL (optional). Default: empty
```

## See also

* [voiceflow completion](/cmd/voiceflow_completion/)	 - Generate the autocompletion script for the specified shell

