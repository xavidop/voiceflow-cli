---
sidebar_position: 2
---

# Replay Command

The `replay` command allows you to replay previously recorded conversations with your Voiceflow project. This is useful for testing changes to your project with consistent inputs, demonstrating flows, or debugging issues.

## Usage

```bash
voiceflow dialog replay -f RECORD_FILE [options]
```

## Options

| Option | Shorthand | Description |
|--------|-----------|-------------|
| `--record-file` | `-f` | Path to the recorded conversation file (required) |
| `--environment` | `-e` | Environment to use (default: "development") |
| `--user-id` | `-u` | User ID for the conversation (optional) |

## Examples

### Replay a recorded conversation

```bash
voiceflow dialog replay -f my-conversation.json
```

This will replay all interactions from the recording file against your Voiceflow project.

### Replay with a specific user ID

```bash
voiceflow dialog replay -f my-conversation.json --user-id user123
```

This allows you to maintain consistent user state or test with specific user profiles.

### Replay in production environment

```bash
voiceflow dialog replay -f my-conversation.json -e production
```

Replays the conversation using your production environment settings.

## How Replay Works

The `replay` command:

1. Reads the recorded conversation file specified with `-f`
2. Processes each interaction in sequence, automatically sending user inputs to the Voiceflow API
3. Displays the responses from your Voiceflow project for each interaction
4. Adds brief pauses between interactions to simulate natural conversation timing

## Creating Recording Files

To create a file for replay, use the `dialog start` command with the `--record-file` option:

```bash
voiceflow dialog start --record-file my-conversation.json
```

During the conversation, every interaction will be saved to the specified file. When you finish the conversation (by typing `exit` or pressing `Ctrl+C`), the complete recording will be available for replay.

## Troubleshooting

If the replay produces different results than the original conversation:

1. Check if your Voiceflow project has been modified since the recording
2. Verify you're using the same environment that was used during recording
3. Consider using a consistent user ID if your project relies on user-specific state
4. Ensure any external APIs or services your project depends on are available
