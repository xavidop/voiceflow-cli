# Server Mode (BETA)

The Voiceflow CLI can run in server mode, providing a local HTTP server with REST API endpoints.

## Starting the Server

There are three ways to start and configure the server:

1. Using command-line flags:
```bash
# Long form with subdomain
voiceflow --server --port 3000 --voiceflow-subdomain staging

# Short form
voiceflow -s -p 3000
```

2. Using environment variables in your `.env` file:
```bash
SERVER=true
PORT=3000 # Optional: Change to any port number
VOICEFLOW_SUBDOMAIN=staging # Optional: Change to your Voiceflow subdomain (private cloud only)
VOICEFLOW_API_KEY=your_api_key
TRUSTED_PROXIES=10.0.0.1,10.0.0.2
```

3. Using both (command-line flags take precedence):
```bash
# Environment variables in .env
SERVER=false
PORT=4000

# Command line overrides
voiceflow --server
```

## Configuration Options

### Server Settings
- `--server, -s`: Enable server mode
- `--port, -p`: Port to run the server on (default: 8080)

### Environment Variables
- `SERVER`: Set to "true" to enable server mode
- `PORT`: Port number for the server
- `VOICEFLOW_SUBDOMAIN`: Voiceflow environment subdomain (private cloud only)
- `VOICEFLOW_API_KEY`: Your Voiceflow API key
- `TRUSTED_PROXIES`: Comma-separated list of trusted proxy IPs

### Authentication
The server requires a Voiceflow API key for authentication. You can provide it in one of two ways:

1. Set it in the `Authorization` header:
```
Authorization: YOUR_API_KEY
```

2. Set it in your `.env` file:
```
VOICEFLOW_API_KEY=YOUR_API_KEY
```

### Proxy Configuration
By default, the server only trusts localhost proxies (`127.0.0.1` and `::1`). If you're running behind a proxy (like Nginx, load balancer, etc.), configure trusted proxy IPs using the `TRUSTED_PROXIES` environment variable:

```bash
TRUSTED_PROXIES=10.0.0.1,10.0.0.2
```

## API Endpoints

For detailed information about available endpoints, see the [API Reference](../api/endpoints.md).

## Agent Endpoints

### Export Agent
`GET /api/agent/export`

Export an agent's configuration as JSON.

**Parameters:**

  - `agent-id` (required): The ID of the agent to export
  - `version-id` (optional): The version ID to export (defaults to "development")

**Response:** JSON configuration of the agent

## Analytics Endpoints

### Fetch Analytics
`GET /api/analytics/fetch`

Fetch analytics data for an agent.

**Parameters:**

- `agent-id` (required): The ID of the agent
- `start-time` (optional): Start time for analytics data
- `end-time` (optional): End time for analytics data
- `limit` (optional): Maximum number of results (default: 100)
- `analytics` (optional): Array of analytics types to fetch. If not specified, fetches all types:
    `interactions`, `sessions`, `top_intents`, `top_slots`, `understood_messages`, `unique_users`, `token_usage`

## Document Endpoints

### Fetch Documents
`GET /api/document/fetch`

Fetch documents from a knowledge base.

**Parameters:**

- `page` (optional): Page number (default: 1)
- `limit` (optional): Results per page (default: 10)
- `document-type` (optional): Type of documents to fetch
- `include-tags` (optional): Array of tags to include
- `exclude-tags` (optional): Array of tags to exclude
- `include-all-non-tagged` (optional): Include all non-tagged documents (default: false)
- `include-all-tagged` (optional): Include all tagged documents (default: false)

### Upload Document URL
`POST /api/document/upload-url`

Upload a document from a URL.

**Request Body:**
```json
{
  "url": "string (required)",
  "name": "string (required)",
  "overwrite": "boolean",
  "max_chunk_size": "number",
  "markdown_conversion": "boolean",
  "llm_generated_q": "boolean",
  "llm_prepend_context": "boolean",
  "llm_based_chunking": "boolean",
  "llm_content_summarization": "boolean",
  "tags": "string[]"
}
```

### Upload Document File
`POST /api/document/upload-file`

Upload a document file.

**Form Data:**

- `file` (required): The file to upload
- `overwrite` (optional): Whether to overwrite existing document
- `max_chunk_size` (optional): Maximum chunk size
- `markdown_conversion` (optional): Convert to markdown
- `llm_generated_q` (optional): Generate questions using LLM
- `llm_prepend_context` (optional): Prepend context using LLM
- `llm_based_chunking` (optional): Use LLM for chunking
- `llm_content_summarization` (optional): Summarize content using LLM
- `tags` (optional): Array of tags

## Knowledge Base Endpoints

### Query Knowledge Base
`POST /api/kb/query`

Query a knowledge base.

**Request Body:**
```json
{
  "question": "string (required)",
  "model": "string",
  "temperature": "number",
  "chunk_limit": "number",
  "synthesis": "boolean",
  "system_prompt": "string",
  "include_tags": "string[]",
  "include_operator": "string",
  "exclude_tags": "string[]",
  "exclude_operator": "string",
  "include_all_tagged": "boolean",
  "include_all_non_tagged": "boolean"
}
```

## Transcript Endpoints

### Fetch Transcript
`GET /api/transcript/fetch`

Fetch a conversation transcript.

**Parameters:**

- `agent-id` (required): The ID of the agent
- `transcript-id` (required): The ID of the transcript
- `output-directory` (required): Directory to save the transcript
