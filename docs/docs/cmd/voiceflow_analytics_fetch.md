# voiceflow analytics fetch

Fetch all project analytics. They could write into a file

```
voiceflow analytics fetch [flags]
```

## Options

```
  -a, --agent-id string         Voiceflow Agent ID (required)
  -t, --analytics stringArray   Analytics to fetch. Default is interactions,sessions,top_intents,top_slots,understood_messages,unique_users,token_usage (optional) (default [interactions,sessions,top_intents,top_slots,understood_messages,unique_users,token_usage])
  -e, --end-time string         Start time in ISO-8601 format to fetch the analytics. Default is current day ago (optional)
  -h, --help                    help for fetch
  -l, --limit int               Limit of analytics to fetch. Default is 100 (optional) (default 100)
  -d, --output-file string      Output directory to save the analytics. Default is analytics.json (optional) (default "analytics.json")
  -s, --start-time string       Start time in ISO-8601 format to fetch the analytics. Default is current day but a month ago (optional)
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

* [voiceflow analytics](/cmd/voiceflow_analytics/)	 - Actions on analytics

