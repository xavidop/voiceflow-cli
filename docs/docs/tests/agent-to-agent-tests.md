# Agent-to-Agent Testing

## Overview

Agent-to-Agent testing is a revolutionary approach to conversation testing that uses AI-powered agents to simulate realistic user interactions with your Voiceflow agent. Instead of predefined scripts, these tests use artificial intelligence to conduct natural, goal-oriented conversations.

There are two types of agent-to-agent testing available:

1. **OpenAI-Powered Testing**: Uses OpenAI models (GPT-4, GPT-3.5, etc.) to simulate user behavior
2. **Voiceflow Agent Testing**: Uses another Voiceflow agent as the tester to interact with your target agent

## How It Works

### OpenAI-Powered Testing Flow

<p align="center">
  <img alt="OpenAI Agent To agent" src="/images/openai-agent-to-agent.png" />
</p>

1. **🚀 Initialization**: An AI agent is configured with a specific goal, persona, and user information
2. **💬 Conversation Start**: The AI agent launches a conversation with your Voiceflow agent
3. **🤖 Dynamic Interaction**: The AI agent responds naturally to your agent's messages, adapting to different conversation paths
4. **📋 Information Requests**: When your agent requests user information, the AI agent provides predefined data or generates realistic responses
5. **🎯 Goal Tracking**: The system continuously evaluates progress toward the specified goal
6. **✅ Completion**: The test succeeds when the goal is achieved or times out after maximum steps

### Voiceflow Agent-to-Agent Testing Flow

<p align="center">
  <img alt="Voiceflow Agent To agent" src="/images/voiceflow-agent-to-agent.png" />
</p>

1. **🚀 Initialization**: Two Voiceflow agents are configured - one as the tester and one as the target
2. **💬 Conversation Start**: Both agents are launched simultaneously
3. **🤖 Agent Interaction**: The tester agent conducts a conversation with your target agent
4. **🎯 Goal Tracking**: OpenAI evaluates whether the specified goal is achieved based on the conversation
5. **✅ Completion**: The test succeeds when the goal is achieved or times out after maximum steps

### Key Advantages

- **🎭 Natural Conversations**: AI agents respond like real users, not scripted robots
- **🔄 Multiple Paths**: One test can explore various conversation flows automatically
- **📊 Comprehensive Coverage**: Tests edge cases and unexpected user behaviors
- **⚡ Efficient**: Replaces dozens of traditional tests with one adaptive test
- **🎯 Goal-Focused**: Measures success based on outcomes, not exact responses
- **🤖 Dual Testing Modes**: Choose between OpenAI-powered testing or Voiceflow agent testing based on your needs

## Test Configuration

### OpenAI-Powered Testing Structure

```yaml
name: Customer Support Agent Test
description: Test agent's ability to resolve customer issues

agent:
  goal: "Get help with a billing issue and update my account information"
  persona: "A confused customer who received an unexpected charge on their account"
  maxSteps: 20
  userInformation:
    - name: 'email'
      value: 'john.doe@example.com'
    - name: 'account_number'
      value: 'ACC-123456'
    - name: 'phone'
      value: '555-0123'
  openAIConfig:
    model: gpt-4o
    temperature: 0.7
```

### Voiceflow Agent-to-Agent Testing Structure

```yaml
name: Customer Support Agent Test
description: Test using a Voiceflow agent as the tester

agent:
  goal: "Get help with a billing issue and update account information"
  maxSteps: 15
  # OpenAI config is still used for goal evaluation
  openAIConfig:
    model: gpt-4o
    temperature: 0.3
  voiceflowAgentTesterConfig:
    environmentName: "production"  # Environment of the tester agent
    apiKey: "VF.DM.your-tester-agent-api-key"
  # Note: userInformation is not used with Voiceflow agent testing
  # The tester agent should be pre-configured with any needed information
```

### Configuration Properties

#### `goal` (Required)
Defines what the AI agent is trying to accomplish. Be specific and measurable.

**Examples:**

- `"Complete a hotel booking for 2 guests for next weekend"`
- `"Report a lost credit card and request a replacement"`
- `"Get technical support for a software installation problem"`
- `"Schedule a doctor's appointment for next month"`

#### `persona` (OpenAI-Powered Testing Only)
Describes the character and context the AI agent should adopt during the conversation. **This property is only used with OpenAI-powered testing and is ignored when using Voiceflow agent testing.**

**Examples:**

- `"An elderly customer who is not tech-savvy and needs extra help"`
- `"A busy professional who wants quick, efficient service"`
- `"A frustrated customer whose previous issue wasn't resolved"`
- `"A new user who doesn't understand the product yet"`

#### `maxSteps` (Required)
Maximum number of conversation turns before the test is considered failed. Consider your conversation complexity when setting this value.

**Recommendations:**

- Simple tasks: 5-10 steps
- Medium complexity: 10-20 steps
- Complex scenarios: 20-30 steps

#### `userInformation` (OpenAI-Powered Testing Testing Only)
Predefined user data that the AI agent can provide when your Voiceflow agent requests personal information. **This property is only used with OpenAI-powered testing.**

For Voiceflow agent testing, any required user information should be pre-configured within the tester agent itself.

**Common Information Types:**

- Contact details: `email`, `phone`, `address`
- Account information: `account_number`, `customer_id`, `membership_id`
- Personal details: `name`, `first_name`, `last_name`, `date_of_birth`
- Transaction data: `order_number`, `transaction_id`, `amount`

#### `openAIConfig` (Optional)
Configures the OpenAI model and parameters used for the AI agent in this specific test. This configuration overrides any suite-level OpenAI settings.

**For OpenAI Testing**: Used to power the AI agent that conducts the conversation.
**For Voiceflow Agent Testing**: Used only for goal evaluation to determine if the test objective has been achieved.

**Properties:**

- `model`: The OpenAI model to use (default: `gpt-4o`)
- `temperature`: Controls response randomness from 0.0 (deterministic) to 1.0 (creative) (default: `0.7`)

#### `voiceflowAgentTesterConfig` (Voiceflow Agent-to-Agent Testing Only)
Configures a Voiceflow agent to act as the tester instead of using OpenAI. When this configuration is present, the system will use agent-to-agent testing with two Voiceflow agents.

**Properties:**

- `environmentName`: The environment name of the tester Voiceflow agent (e.g., "production", "development")
- `apiKey`: The API key for the tester Voiceflow agent (format: `VF.DM.xxxxx.xxxxx`)

**Important Notes:**

- When using Voiceflow agent testing, the `persona` and `userInformation` properties are ignored
- The tester agent should be pre-configured with appropriate conversation logic and any required user data
- OpenAI is still used for goal evaluation even when using Voiceflow agent testing

**Example:**

```yaml
# OpenAI-powered testing configuration
agent:
  goal: "Get technical support for a complex software issue"
  persona: "A software developer who needs detailed technical assistance"
  maxSteps: 15
  openAIConfig:
    model: gpt-4o
    temperature: 0.3  # Lower temperature for more focused technical responses
```

```yaml
# Voiceflow agent-to-agent testing configuration
agent:
  goal: "Complete a hotel booking for this weekend"
  maxSteps: 12
  openAIConfig:
    model: gpt-4o
    temperature: 0.3  # Used only for goal evaluation
  voiceflowAgentTesterConfig:
    environmentName: "production"
    apiKey: "VF.DM.your-tester-agent-key"
```

## Choosing Between Testing Methods

### OpenAI-Powered Testing

**Best for:**

- ✅ Flexible persona and behavior simulation
- ✅ Dynamic user information generation
- ✅ Complex reasoning and decision-making scenarios
- ✅ Testing edge cases and unexpected user behaviors
- ✅ Rapid prototyping and testing different user types

**Requirements:**

- OpenAI API key and sufficient quota
- Persona and user information configuration

### Voiceflow Agent Testing

**Best for:**

- ✅ Consistent, reproducible test behavior
- ✅ Testing specific conversation flows designed in Voiceflow
- ✅ Using existing Voiceflow agents as test users
- ✅ Avoiding OpenAI API costs for conversation simulation
- ✅ Pre-configured user personas built in Voiceflow

**Requirements:**

- A separate Voiceflow agent configured as the tester
- API key for the tester agent
- OpenAI API key still needed for goal evaluation

## OpenAI Model Configuration

### Model Recommendations

- `gpt-4o`: Best for complex reasoning and nuanced conversations
- `gpt-4o-mini`: Good balance of performance and cost for most use cases
- `gpt-3.5-turbo`: Budget-friendly option for simpler interactions

### Temperature Guidelines

- `0.0-0.3`: Highly focused, deterministic responses (ideal for technical support)
- `0.4-0.7`: Balanced responses with some creativity (good for general conversations)
- `0.8-1.0`: More creative and varied responses (useful for casual interactions)

## Suite-Level OpenAI Configuration

You can also configure OpenAI settings at the suite level, which applies to all agent tests unless overridden at the test level:

```yaml
name: Customer Service Test Suite
description: Comprehensive customer service scenarios
environmentName: production

# Suite-level OpenAI configuration
openAIConfig:
  model: gpt-4o-mini
  temperature: 0.5

tests:
  - id: billing_support
    file: ./tests/billing_test.yaml
  - id: technical_support
    file: ./tests/technical_test.yaml  # Can override with test-level config
  - id: voiceflow_agent_test
    file: ./tests/voiceflow_agent_test.yaml  # Uses suite config for goal evaluation
```

## Best Practices

### Writing Effective Goals

✅ **Good Goals:**

- Specific and measurable
- Achievable within the conversation scope
- Focused on user outcomes

❌ **Avoid:**

- Vague objectives
- Impossible tasks
- Testing internal system functions

### Creating Realistic Personas (OpenAI Testing)

✅ **Good Personas:**

- Include emotional context
- Specify technical skill level
- Mention relevant background

❌ **Avoid:**

- Generic descriptions
- Inconsistent characteristics
- Unrealistic behaviors

### Configuring Voiceflow Tester Agents

✅ **Best Practices:**

- Design the tester agent with clear conversation flows
- Include appropriate user information within the agent
- Test the tester agent independently before using in tests
- Use meaningful environment names and secure API keys

❌ **Avoid:**

- Using production agents directly as testers
- Hardcoding sensitive information in tester agents
- Creating overly complex tester agent flows

### Setting Appropriate Step Limits

- **Too Low**: May timeout before completion
- **Too High**: May hide conversation inefficiencies
- **Just Right**: Allows natural completion with buffer for edge cases

## Authentication Requirements

### OpenAI Testing Requirements

OpenAI-powered agent tests require OpenAI API access for the AI agent functionality. Make sure to:

1. Set up your OpenAI API key in your environment
2. Configure authentication as described in the [Authentication](/overview/authentication) page
3. Ensure sufficient API quota for test execution

### Voiceflow Agent Testing Requirements

Voiceflow agent-to-agent tests require:

1. **Target Agent**: API key for the agent being tested
2. **Tester Agent**: API key for the agent acting as the tester (specified in `voiceflowAgentTesterConfig`)
3. **OpenAI API**: Still required for goal evaluation functionality
4. **Environment Access**: Ensure both agents are accessible in their respective environments

## Monitoring and Debugging

### Test Logs

Both testing methods provide detailed logs including:

**OpenAI Testing Logs:**

- AI agent's thought process and responses
- Conversation flow and decision points
- Goal achievement evaluation
- User information requests and responses

**Voiceflow Agent Testing Logs:**

- Interaction between tester and target agents
- Message exchange flow
- Goal achievement evaluation
- Agent response details

### Common Issues

**General Issues:**

- **Goal not achieved**: Review if the goal is realistic and achievable
- **Timeout errors**: Consider increasing `maxSteps` or simplifying the goal
- **Authentication errors**: Verify API key configuration

**OpenAI Testing Specific:**

- **Inconsistent behavior**: AI responses may vary; focus on goal achievement rather than exact responses
- **OpenAI API errors**: Check API key and quota limits

**Voiceflow Agent Testing Specific:**

- **Tester agent errors**: Verify the tester agent is properly configured and accessible
- **API key issues**: Ensure both target and tester agent API keys are valid
- **Environment mismatches**: Verify environment names are correct for both agents