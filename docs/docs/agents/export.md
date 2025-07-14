# Export an agent

With the `voiceflow-cli` you can export your agent information. This is useful when you want to get the `.vf` file of your project. The `voiceflow-cli` has one command that allows you to export an agent from your terminal:

To export an agent, you need to know the `agent-id` and the `version-id` of the agent you want to export from. You can find that information in the Voiceflow Agent section under your Agent Settings on [voiceflow.com](https://voiceflow.com).

```sh
voiceflow agent export --agent-id <your-agent-id> --version-id <your-version-id> --output-file <path-to-save>
```