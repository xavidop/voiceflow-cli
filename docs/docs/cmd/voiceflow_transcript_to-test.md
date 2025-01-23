# voiceflow transcript to-test

Transforms a transcript into a test

```
voiceflow transcript to-test [flags]
```

## Options

```
  -a, --agent-id string           Voiceflow Agent ID (required)
  -h, --help                      help for to-test
  -d, --output-file string        Output file to save the test. Default is test.yaml (optional) (default "test.yaml")
  -e, --test-description string   Test description (optional) (default "Test")
  -n, --test-name string          Test name (optional) (default "Test")
  -t, --transcript-id string      Voiceflow Transcript ID (required)
```

## Options inherited from parent commands

```
  -o, --output-format string         Output Format. Options: text, json. Default: text (optional) (default "text")
  -u, --skip-update-check            Skip the check for updates check run before every command (optional)
  -v, --verbose                      verbose error output (with stack trace) (optional)
  -x, --voiceflow-api-key string     Voiceflow API Key (optional)
  -b, --voiceflow-subdomain string   Voiceflow Base URL (optional). Default: empty
```

## See also

* [voiceflow transcript](/cmd/voiceflow_transcript/)	 - Actions on transcripts

