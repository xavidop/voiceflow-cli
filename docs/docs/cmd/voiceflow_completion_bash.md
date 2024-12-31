# voiceflow completion bash

Generate the autocompletion script for bash

## Synopsis

Generate the autocompletion script for the bash shell.

This script depends on the 'bash-completion' package.
If it is not installed already, you can install it via your OS's package manager.

To load completions in your current shell session:

	source <(voiceflow completion bash)

To load completions for every new session, execute once:

### Linux:

	voiceflow completion bash > /etc/bash_completion.d/voiceflow

### macOS:

	voiceflow completion bash > $(brew --prefix)/etc/bash_completion.d/voiceflow

You will need to start a new shell for this setup to take effect.


```
voiceflow completion bash
```

## Options

```
  -h, --help              help for bash
      --no-descriptions   disable completion descriptions
```

## Options inherited from parent commands

```
  -o, --output-format string        Output Format. Options: text, json. Default: text (optional) (default "text")
  -u, --skip-update-check           Skip the check for updates check run before every command (optional)
  -v, --verbose                     verbose error output (with stack trace) (optional)
  -x, --voiceflow-api-key string    Voiceflow API Key (optional)
  -b, --voiceflow-base-url string   Voiceflow Base URL (optional). Default: https://general-runtime.voiceflow.com (default "https://general-runtime.voiceflow.com")
```

## See also

* [voiceflow completion](/cmd/voiceflow_completion/)	 - Generate the autocompletion script for the specified shell

