# Dialog Commands

The `dialog` commands allow you to interact with your Voiceflow project through a conversational interface. You can start new conversations, record them for later use, replay previous conversations, and create tests from your interactions.

## Available Commands

| Command | Description |
|---------|-------------|
| [start](./start.md) | Start a new conversation with your Voiceflow project |
| [replay](./replay.md) | Replay a previously recorded conversation |

## Common Options

All dialog commands support these common options:

| Option | Description |
|--------|-------------|
| `--environment`, `-e` | Voiceflow environment to use (default: "development") |
| `--user-id`, `-u` | User ID for the conversation (optional, will generate a random ID if not provided) |

## Basic Usage

```bash
# Start a new conversation in the development environment
voiceflow dialog start

# Replay a recorded conversation
voiceflow dialog replay -f conversation.json

# Start a conversation in the production environment
voiceflow dialog start -e production
```

For detailed information about each command, refer to their specific documentation pages.
