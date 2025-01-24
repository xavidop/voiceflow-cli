# voiceflow kb query

Query a knowledge base

```
voiceflow kb query [flags]
```

## Options

```
  -c, --chunk-limit int            Chunk limit to use while asking the knowledge base. Default to 2 (optional) (default 2)
  -j, --exclude-operator string    Tags to exclude. Possible values: and/or. Default is empty (optional)
  -y, --exclude-tags stringArray   Tags to exclude. Default is empty (optional)
  -h, --help                       help for query
  -n, --include-all-non-tagged     Filters KB documents to include those that have no KB tags attached. Default to false (optional)
  -g, --include-all-tagged         Filters KB documents to include those that have any KB tags attached. Default to false (optional)
  -i, --include-operator string    Tags to include. Possible values: and/or. Default is empty (optional)
  -t, --include-tags stringArray   Tags to include. Default is empty (optional)
  -m, --model string               Model to use while asking the knowledge base (required)
  -d, --output-file string         Output directory to save the information returned by the CLI. Default is query.json (optional) (default "query.json")
  -q, --question string            Question to ask to the knowledge base (required)
  -s, --synthesis                  Indicates whether to use language models to generate an answer. Default to true (optional) (default true)
  -p, --system-prompt string       System prompt to use while asking the knowledge base. Default is empty (optional)
  -r, --temperature float          Temperature to use while asking the knowledge base. Default to 0.7 (optional) (default 0.7)
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

* [voiceflow kb](/cmd/voiceflow_kb/)	 - Actions on knowledge base

