# Export an agent

With the `vocieflow-cli` you can query your agents' knowledge base. This is useful when you want to perform a query on your knowledge base. The `voiceflow-cli` has one command that allows you to query your knowledge base from your terminal:

## Parameters

### Required Parameters

- `--model, -m`: AI model to process knowledge base queries
    * Required: Yes
    * Example: `--model gpt-4`

### Optional Parameters

#### Model Configuration
- `--temperature, -r`
    - Range: 0.0 to 1.0
    - Default: 0.7
    - Purpose: Controls response randomness

- `--chunk-limit, -c`
    - Default: 2
    - Purpose: Maximum chunks to process

- `--synthesis, -s`
    - Default: true
    - Purpose: Enable/disable AI answer generation

- `--system-prompt, -p`
    - Default: empty
    - Purpose: Custom system instructions

#### Tag Filtering
- `--include-tags, -t`
    - Default: []
    - Purpose: Tags to include in search

- `--include-operator, -i`
    - Values: "and"/"or"
    - Purpose: Logic operator for included tags

- `--exclude-tags, -y`
    - Default: []
    - Purpose: Tags to exclude from search

- `--exclude-operator, -j`
    - Values: "and"/"or"
    - Purpose: Logic operator for excluded tags

- `--include-all-tagged, -g`
    - Default: false
    - Purpose: Include all documents with tags

- `--include-all-non-tagged, -n`
    - Default: false
    - Purpose: Include all documents without tags

#### Output
- `--output-file, -d`
    - Default: "query.json"
    - Purpose: Results output location

## Examples

### Basic Query

```bash
vf kb query --quesiton "How does feature X work?" --model gpt-4
```