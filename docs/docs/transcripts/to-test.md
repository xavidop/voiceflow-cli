# Transform Transcripts into Tests

## Overview
Convert a Voiceflow transcript into a reusable [test case](/tests/tests).

## Command
```bash
voiceflow transcript to-test [flags]
```

## Parameters

### Required Flags
- `--agent-id`: Voiceflow Agent ID
- `--transcript-id`: ID of the transcript to convert

### Optional Flags
- `--output-file`: Path to save the generated test (default: test.yaml)
- `--test-name`: Name for the generated test
- `--test-description`: Description for the generated test

## Examples

### Basic Usage
```bash
voiceflow transcript to-test \
  --agent-id your-agent-id \
  --transcript-id transcript-123
```

### Full Example with Options
```bash
voiceflow transcript to-test \
  --agent-id your-agent-id \
  --transcript-id transcript-123 \
  --output-file my-test.yaml \
  --test-name "Payment Flow Test" \
  --test-description "Validates the payment processing dialogue"
```

## Output
The command generates a YAML file containing:

  - Test metadata (name, description)
  - User interactions
  - Expected agent responses
  - Validation rules