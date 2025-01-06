# Upload URLs to the Knowledge Base

With the `vocieflow-cli` you can upload content from a file to your Voiceflow Knowledge Base with customizable processing options. This is useful when you want to perform a automations around your knowledge base. The `voiceflow-cli` has one command that allows you to update your knowledge base from your terminal:

## Command Usage
```bash
voiceflow document upload-file [flags]
```

### Aliases
- `uf`
- `upload-files`

## Parameters

### Required Flags
- `--file`: Path to local file

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
- `--tags`: Array of tags to associate with document

## Examples

### Basic File Upload
```bash
voiceflow document upload-file \
  --file ./docs/api.pdf
```

### Advanced Upload with Processing
```bash
voiceflow document upload-file \
  --file ./docs/guide.md \
  --max-chunk-size 1000 \
  --markdown-conversion \
  --llm-generated-q \
  --llm-content-summarization \
  --tags guide,user
```

### Upload with Overwrite
```bash
voiceflow document upload-file \
  --file ./docs/updated-policy.pdf \
  --overwrite \
  --tags policy,legal
```

## Supported File Types
- PDF
- TXT
- DOC/DOCX
- MD
- And other text-based formats

## Requirements
- File must be accessible locally
