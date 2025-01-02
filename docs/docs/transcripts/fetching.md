# Fetching transcripts

With the `vocieflow-cli` you can fetch the transcripts of your project. This is useful when you want to analyze the conversation flow of your agent. The `voiceflow-cli` has 2 commands that allow you to fetch the transcripts from your terminal:

## Fetching all transcripts

To fetch all transcripts, you need to know the `agent-id` of the agent you want to fetch the transcripts from. You can find the `agent-id` in the Voiceflow Agent section.

```sh
voiceflow transcript fetch-all --agent-id <your-agent-id>
```

## Fetching a specific transcript

To fetch a specific transcript, you need to know the `transcript-id` of the transcript you want to fetch. You can find the `transcript-id` in the Voiceflow Transcript section.

```sh
voiceflow transcript fetch --agent-id <your-agent-id> --transcript-id <your-transcript-id>
```
