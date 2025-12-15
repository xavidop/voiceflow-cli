# voiceflow document upload-url

Upload a URL to a knowledge base

```
voiceflow document upload-url [flags]
```

## Options

```
  -h, --help                        help for upload-url
  -g, --llm-based-chunking          LLM to determine the optimal chunking of the document content based on semantic similarity and retrieval effectiveness. Default is false (optional)
  -s, --llm-content-summarization   LLM to summarize and rewrite the content, removing unnecessary information and focusing on important parts to optimize for retrieval. Default is false (optional)
  -q, --llm-generated-q             If an LLM to generate a question based on the document context and specific chunk, and prepend it to the chunk. Default is false (optional)
  -p, --llm-prepend-context         LLM to generate a context summary based on the document and chunk context, and prepend it to each chunk. Default is false (optional)
  -k, --markdown-conversion         Enable HTML to markdown conversion. Default is false (optional)
  -m, --max-chunk-size int          Determines how granularly each document is broken up. Default is 1000 (optional) (default 1000)
  -n, --name string                 Name of the document that will be uploaded to the knowledge base (required)
  -w, --overwrite                   Overwrite the document if it already exists in the knowledge base. Default is false (optional)
  -t, --tags stringArray            An array of tag labels to attach to a KB document that can be used to filter document eligibility in query retrieval. Default is empty (optional)
  -l, --url string                  URL to upload to the knowledge base (required)
```

## Options inherited from parent commands

```
  -z, --open-api-key string          Open API Key (optional)
  -o, --output-format string         Output Format. Options: text, json. Default: text (optional) (default "text")
      --show-tester-messages         Show tester agent messages in agent-to-agent tests (optional) (default true)
  -u, --skip-update-check            Skip the check for updates check run before every command (optional)
  -v, --verbose                      verbose error output (with stack trace) (optional)
  -x, --voiceflow-api-key string     Voiceflow API Key (optional)
  -b, --voiceflow-subdomain string   Voiceflow Base URL (optional). Default: empty
```

## See also

* [voiceflow document](/cmd/voiceflow_document/)	 - Actions on documents

