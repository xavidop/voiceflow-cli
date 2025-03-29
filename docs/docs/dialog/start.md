# Start Command

The `start` command initiates an interactive conversation with your Voiceflow project. This allows you to test your project's dialog flow by sending text inputs and receiving responses.

## Usage

```bash
voiceflow dialog start [options]
```

## Options

| Option | Shorthand | Description |
|--------|-----------|-------------|
| `--environment` | `-e` | Environment to use (default: "development") |
| `--user-id` | `-u` | User ID for the conversation (optional) |
| `--record-file` | `-f` | File to save the conversation recording (optional) |
| `--save-test` | `-t` | Save the conversation as a test file (optional) |

## Examples

### Start a basic conversation

```bash
voiceflow dialog start
```

This starts a conversation with your Voiceflow project in the development environment. You can type messages and see the responses from your project.

### Start with a specific user ID

```bash
voiceflow dialog start --user-id user123
```

Using a consistent user ID allows the conversation to maintain state across multiple sessions.

### Record a conversation

```bash
voiceflow dialog start --record-file my-conversation.json
```

This will save the entire conversation to a file that can be replayed later using the `replay` command.

### Start a conversation and save it as a test

```bash
voiceflow dialog start --save-test
```

This records the conversation and automatically saves it as a YAML test file that can be used with the `voiceflow test` commands.

### Conversation in production environment

```bash
voiceflow dialog start -e production
```

Starts the conversation using your production environment settings.

## Interactive Commands

During an active conversation session, you can use these special commands:

| Command | Action |
|---------|--------|
| `exit` or `quit` | End the conversation and exit |
| `Ctrl+C` | Interrupt the conversation (will save recordings if enabled) |

## Recording Format

When you use the `--record-file` option, the conversation is saved in JSON format with the following structure:

```json
{
  "name": "Recording_YYYYMMDD_HHMMSS",
  "interactions": [
    {
      "id": "launch",
      "user": {
        "type": "launch"
      },
      "agent": [
        {
          "type": "text",
          "value": "Hello! How can I help you today?"
        }
      ]
    },
    {
      "id": "interaction_1",
      "user": {
        "type": "text",
        "text": "What's the weather like?"
      },
      "agent": [
        {
          "type": "text",
          "value": "I don't have access to weather information."
        }
      ]
    }
  ]
}
```

This recording can be used with the `voiceflow dialog replay` command to repeat the conversation.
