# Upload URLs to the Knowledge Base

With the `vocieflow-cli` you can upload content from a URL to your Voiceflow Knowledge Base with customizable processing options. This is useful when you want to perform a automations around your knowledge base. The `voiceflow-cli` has one command that allows you to update your knowledge base from your terminal:

## Command Usage
```bash
vf document upload-url [flags]
```

### Aliases
- `ur`
- `upload-urls`

## Parameters

### Required Flags
- `--url`: URL to upload content from
- `--name`: Name for the uploaded document

### Processing Options
- `--max-chunk-size`: Maximum size of content chunks
- `--markdown-conversion`: Convert content to markdown format
- `--overwrite`: Overwrite existing document if present

### LLM Processing Options
- `--llm-generated-q`: Enable LLM-generated questions
- `--llm-prepend-context`: Prepend context using LLM
- `--llm-based-chunking`: Use LLM for content chunking
- `--llm-content-summarization`: Enable content summarization

### Metadata
- `--tags`: Array of tags to associate with the document

## Examples

### Basic Upload
```bash
vf document upload-url --url https://docs.example.com/api --name "API Documentation"
```

### Advanced Upload with LLM Processing

```bash
vf document upload-url \
  --url https://docs.example.com/api \
  --name "API Documentation" \
  --max-chunk-size 1000 \
  --markdown-conversion \
  --llm-generated-q \
  --llm-content-summarization \
  --tags api,documentation
```

### Upload with Overwrite

```bash
vf document upload-url \
  --url https://docs.example.com/api \
  --name "API Documentation" \
  --overwrite \
  --tags api,v2
```

