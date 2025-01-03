# Fetching transcripts

With the `vocieflow-cli` you can fetch the transcripts of your project. This is useful when you want to analyze the conversation flow of your agent. The `voiceflow-cli` has 2 commands that allow you to fetch the transcripts from your terminal:

## Fetching all transcripts

To fetch all transcripts, you need to know the `agent-id` of the agent you want to fetch the transcripts from. You can find that information in the Voiceflow Agent section under your Agent Settings on [voiceflow.com](https://voiceflow.com).

```sh
voiceflow transcript fetch-all --agent-id <your-agent-id>
```
### Time Range Filters

- Start Time
    * Flag: `--start-time, -s`
    * Format: ISO-8601
    * Default: Current date minus one month
    * Example: `--start-time 2024-01-01T00:00:00Z`

- End Time
    * Flag: `--end-time, -e`
    * Format: ISO-8601
    * Default: Current date
    * Example: `--end-time 2024-02-01T00:00:00Z`

### Content Filters

- Tag Filter
    * Flag: `--tag, -g`
    * Description: Filter transcripts by specific tag
    * Default: Empty (no filter)
    * Example: `--tag production`

- Range Filter
    * Flag: `--range, -r`
    * Description: Filter transcripts by date range
    * Default: Empty (no filter)
    * Example: `--range Yesterday`

### Example Usage

```bash
voiceflow transcript fetch-all \
  --agent-id abc123 \
  --start-time 2024-01-01T00:00:00Z \
  --end-time 2024-02-01T00:00:00Z \
  --tag production \
  --range Yesterday \
  --output-directory ./my-transcripts
```

## Fetching a specific transcript

To fetch a specific transcript, you need to know the `transcript-id` of the transcript you want to fetch. You can find the `transcript-id` in the Voiceflow Transcript section.

```sh
voiceflow transcript fetch --agent-id <your-agent-id> --transcript-id <your-transcript-id>
```
