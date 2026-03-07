# voiceflow dialog start

Start a dialog with the Voiceflow project

```
voiceflow dialog start [flags]
```

## Options

```
  -e, --environment string   Environment to use (optional). Default to development (default "development")
  -h, --help                 help for start
  -f, --record-file string   Record file to use (optional)
  -t, --save-as-test         Save conversation as a test (optional)
  -r, --user-id string       User ID for the dialog (optional)
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

* [voiceflow dialog](/cmd/voiceflow_dialog/)	 - Start a dialog with the Voiceflow project

