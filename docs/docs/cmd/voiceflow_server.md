# voiceflow server

Start the Voiceflow CLI API server

## Synopsis

Start the Voiceflow CLI API server to expose test execution endpoints.

The server provides HTTP endpoints for:
- Executing test suites
- Checking test execution status
- Retrieving system information

The server includes auto-generated OpenAPI/Swagger documentation available at /swagger/index.html

```
voiceflow server [flags]
```

## Examples

```
  # Start server on default port (8080)
  voiceflow server

  # Start server on custom port
  voiceflow server --port 9090

  # Start server with debug mode
  voiceflow server --debug

  # Start server with custom host
  voiceflow server --host 127.0.0.1 --port 8080
```

## Options

```
      --cors          Enable CORS middleware (default true)
  -d, --debug         Enable debug mode
  -h, --help          help for server
  -H, --host string   Host to bind the server to (default "0.0.0.0")
  -p, --port string   Port to run the server on (default "8080")
      --swagger       Enable Swagger documentation endpoint (default true)
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

* [voiceflow](/cmd/voiceflow/)	 - Voiceflow CLI

