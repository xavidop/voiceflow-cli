# voiceflow completion zsh

Generate the autocompletion script for zsh

## Synopsis

Generate the autocompletion script for the zsh shell.

If shell completion is not already enabled in your environment you will need
to enable it.  You can execute the following once:

	echo "autoload -U compinit; compinit" >> ~/.zshrc

To load completions in your current shell session:

	source <(voiceflow completion zsh)

To load completions for every new session, execute once:

### Linux:

	voiceflow completion zsh > "${fpath[1]}/_voiceflow"

### macOS:

	voiceflow completion zsh > $(brew --prefix)/share/zsh/site-functions/_voiceflow

You will need to start a new shell for this setup to take effect.


```
voiceflow completion zsh [flags]
```

## Options

```
  -h, --help              help for zsh
      --no-descriptions   disable completion descriptions
```

## Options inherited from parent commands

```
  -o, --output-format string       Output Format. Options: text, json. Default: text (optional) (default "text")
  -u, --skip-update-check          Skip the check for updates check run before every command (optional)
  -v, --verbose                    verbose error output (with stack trace) (optional)
  -c, --voiceflow-api-key string   Voiceflow API Key (optional)
```

## See also

* [voiceflow completion](/cmd/voiceflow_completion/)	 - Generate the autocompletion script for the specified shell

