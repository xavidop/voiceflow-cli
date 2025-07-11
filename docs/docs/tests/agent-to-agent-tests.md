# Agent-to-Agent Testing

## Overview

Agent-to-Agent testing is a revolutionary approach to conversation testing that uses AI-powered agents to simulate realistic user interactions with your Voiceflow agent. Instead of predefined scripts, these tests use artificial intelligence to conduct natural, goal-oriented conversations.

## How It Works

### The Testing Flow

1. **🚀 Initialization**: An AI agent is configured with a specific goal, persona, and user information
2. **💬 Conversation Start**: The AI agent launches a conversation with your Voiceflow agent
3. **🤖 Dynamic Interaction**: The AI agent responds naturally to your agent's messages, adapting to different conversation paths
4. **📋 Information Requests**: When your agent requests user information, the AI agent provides predefined data or generates realistic responses
5. **🎯 Goal Tracking**: The system continuously evaluates progress toward the specified goal
6. **✅ Completion**: The test succeeds when the goal is achieved or times out after maximum steps

### Key Advantages

- **🎭 Natural Conversations**: AI agents respond like real users, not scripted robots
- **🔄 Multiple Paths**: One test can explore various conversation flows automatically
- **📊 Comprehensive Coverage**: Tests edge cases and unexpected user behaviors
- **⚡ Efficient**: Replaces dozens of traditional tests with one adaptive test
- **🎯 Goal-Focused**: Measures success based on outcomes, not exact responses

## Test Configuration

### Basic Structure

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
```

### Configuration Properties

#### `goal` (Required)
Defines what the AI agent is trying to accomplish. Be specific and measurable.

**Examples:**

- `"Complete a hotel booking for 2 guests for next weekend"`
- `"Report a lost credit card and request a replacement"`
- `"Get technical support for a software installation problem"`
- `"Schedule a doctor's appointment for next month"`

#### `persona` (Required)
Describes the character and context the AI agent should adopt during the conversation.

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

#### `userInformation` (Optional)
Predefined user data that the AI agent can provide when your Voiceflow agent requests personal information.

**Common Information Types:**

- Contact details: `email`, `phone`, `address`
- Account information: `account_number`, `customer_id`, `membership_id`
- Personal details: `name`, `first_name`, `last_name`, `date_of_birth`
- Transaction data: `order_number`, `transaction_id`, `amount`

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

### Creating Realistic Personas

✅ **Good Personas:**

- Include emotional context
- Specify technical skill level
- Mention relevant background

❌ **Avoid:**

- Generic descriptions
- Inconsistent characteristics
- Unrealistic behaviors

### Setting Appropriate Step Limits

- **Too Low**: May timeout before completion
- **Too High**: May hide conversation inefficiencies
- **Just Right**: Allows natural completion with buffer for edge cases

## Authentication Requirements

Agent-to-Agent tests require OpenAI API access for the AI agent functionality. Make sure to:

1. Set up your OpenAI API key in your environment
2. Configure authentication as described in the [Authentication](/overview/authentication) page
3. Ensure sufficient API quota for test execution

## Monitoring and Debugging

### Test Logs

Agent-to-Agent tests provide detailed logs including:

- AI agent's thought process and responses
- Conversation flow and decision points
- Goal achievement evaluation
- User information requests and responses

### Common Issues
- **Goal not achieved**: Review if the goal is realistic and achievable
- **Timeout errors**: Consider increasing `maxSteps` or simplifying the goal
- **Authentication errors**: Verify OpenAI API key configuration
- **Inconsistent behavior**: AI responses may vary; focus on goal achievement rather than exact responses

