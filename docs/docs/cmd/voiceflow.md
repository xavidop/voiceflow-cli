# voiceflow

Voiceflow CLI

## Synopsis

Welcome to voiceflow-cli!

This utility provides you with an easy way to interact
with your Voiceflow agents.

You can find the documentation at https://github.com/xavidop/voiceflow-cli.

Please file all bug reports on GitHub at https://github.com/xavidop/voiceflow-cli/issues.

```
voiceflow [flags]
```

## Options

```
  -h, --help                         help for voiceflow
  -z, --open-api-key string          Open API Key (optional)
  -o, --output-format string         Output Format. Options: text, json. Default: text (optional) (default "text")
      --show-tester-messages         Show tester agent messages in agent-to-agent tests (optional) (default true)
  -u, --skip-update-check            Skip the check for updates check run before every command (optional)
  -v, --verbose                      verbose error output (with stack trace) (optional)
  -x, --voiceflow-api-key string     Voiceflow API Key (optional)
  -b, --voiceflow-subdomain string   Voiceflow Base URL (optional). Default: empty
```

## See also

* [voiceflow agent](/cmd/voiceflow_agent/)	 - Actions on agents
* [voiceflow analytics](/cmd/voiceflow_analytics/)	 - Actions on analytics
* [voiceflow completion](/cmd/voiceflow_completion/)	 - Generate the autocompletion script for the specified shell
* [voiceflow dialog](/cmd/voiceflow_dialog/)	 - Start a dialog with the Voiceflow project
* [voiceflow document](/cmd/voiceflow_document/)	 - Actions on documents
* [voiceflow jsonschema](/cmd/voiceflow_jsonschema/)	 - outputs voiceflow's JSON schema
* [voiceflow kb](/cmd/voiceflow_kb/)	 - Actions on knowledge base
* [voiceflow server](/cmd/voiceflow_server/)	 - Start the Voiceflow CLI API server
* [voiceflow test](/cmd/voiceflow_test/)	 - Actions on conversation testing
* [voiceflow transcript](/cmd/voiceflow_transcript/)	 - Actions on transcripts
* [voiceflow version](/cmd/voiceflow_version/)	 - Get voiceflow-cli version

