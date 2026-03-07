# voiceflow transcript fetch

Fetch one transcripts from a project

```
voiceflow transcript fetch [flags]
```

## Options

```
  -a, --agent-id string           Voiceflow Agent ID (required)
  -h, --help                      help for fetch
  -d, --output-directory string   Output directory to save the transcripts. Default is ./output (optional) (default "./output")
  -t, --transcript-id string      Voiceflow Transcript ID (required)
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

* [voiceflow transcript](/cmd/voiceflow_transcript/)	 - Actions on transcripts

