# voiceflow agent export

Export a voiceflow project into a file

```
voiceflow agent export [flags]
```

## Options

```
  -a, --agent-id string      Voiceflow Agent ID (required)
  -h, --help                 help for export
  -d, --output-file string   Output directory to save the VF file. Default is agent.vf (optional) (default "agent.vf")
  -s, --version-id string    Voiceflow Version ID (optional). Default: development (default "development")
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

* [voiceflow agent](/cmd/voiceflow_agent/)	 - Actions on agents

