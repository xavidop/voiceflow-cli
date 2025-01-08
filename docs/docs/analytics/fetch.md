# Fetch Analytics

With the `vocieflow-cli` you can fetch the analytics of your project. This is useful when you want to get the get the analytics and import them to another system. The `voiceflow-cli` has one command that allows you to export your voiceflow agent Analytics from your terminal:

To export the analytics, you need to know the `agent-id` of the agent you want to export from. You can find that information in the Voiceflow Agent section under your Agent Settings on [voiceflow.com](https://voiceflow.com).

```sh
voiceflow analytics fetch --agent-id <your-agent-id> --output-file <path-to-save>
```

### Filters

The Voiceflow analytics command has a few filters that you can use to narrow down the data you want to export. The filters are:

#### Time Range

- Start Time

    * Flag: `--start-time, -s`
    * Format: ISO-8601
    * Default: Current day minus one month
    * Example: `--start-time 2025-01-01T00:00:00.000Z`

- End Time

    * Flag: `--end-time, -s`
    * Format: ISO-8601
    * Default: Current day
    * Example: `--end-time 2025-01-02T00:00:00.000Z`


#### Limit

- Flag: `--limit, -l`
- Description: Maximum number of records to fetch
- Default: 100
- Example: `--limit 500`

#### Output File

- Flag: `--output-file, -d`
- Description: Path where analytics will be saved
- Default: `analytics.json`
- Example: `--output-file my-analytics.json`

#### Analytics Types

- Flag: `--analytics, -t`
- Description: Types of analytics to fetch
- Default: All types listed below
- Multiple values allowed: Yes

| Analytics Type | Description |
|---------------|-------------|
| `interactions` | User interaction data |
| `sessions` | Session-level analytics |
| `top_intents` | Most triggered intents |
| `top_slots` | Most used slots |
| `understood_messages` | Successfully parsed messages |
| `unique_users` | Distinct user counts |
| `token_usage` | API token consumption |

## Example Usage

```bash
voiceflow analytics fetch \
  --agent-id abc123 \
  --start-time 2025-01-01T00:00:00.000Z \
  --end-time 2025-01-02T00:00:00.000Z \
  --limit 500 \
  --analytics interactions,sessions \
  --output-file jan-2024-analytics.json
```
