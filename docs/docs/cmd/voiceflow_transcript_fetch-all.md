# voiceflow transcript fetch-all

Fetch all transcripts from a project

```
voiceflow transcript fetch-all [flags]
```

## Options

```
  -a, --agent-id string           Voiceflow Agent ID (required)
  -e, --end-time string           Start time in ISO-8601 format to fetch the analytics. Default is current day ago (optional)
  -h, --help                      help for fetch-all
  -d, --output-directory string   Output directory to save the transcripts. Default is ./output (optional) (default "./output")
  -r, --range string              Range to filter the transcripts. Default is empty (optional)
  -s, --start-time string         Start time in ISO-8601 format to fetch the analytics. Default is current day but a month ago (optional)
  -g, --tag string                Tag to filter the transcripts. Default is empty (optional)
```

## Options inherited from parent commands

```
  -z, --open-api-key string          Open API Key (optional)
  -o, --output-format string         Output Format. Options: text, json. Default: text (optional) (default "text")
  -u, --skip-update-check            Skip the check for updates check run before every command (optional)
  -v, --verbose                      verbose error output (with stack trace) (optional)
  -x, --voiceflow-api-key string     Voiceflow API Key (optional)
  -b, --voiceflow-subdomain string   Voiceflow Base URL (optional). Default: empty
```

## See also

* [voiceflow transcript](/cmd/voiceflow_transcript/)	 - Actions on transcripts

