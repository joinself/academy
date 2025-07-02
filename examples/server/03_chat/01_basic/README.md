# Basic Chat Example 🟢

A straightforward demonstration of the Self SDK's chat messaging capabilities using the core SDK. This example focuses on essential chat functionality with clean, production-ready code patterns.

## 🎯 What You'll Learn

- **Account Setup**: Create and configure a Self account for messaging
- **Event-Driven Architecture**: Handle incoming messages with callbacks
- **Chat Messaging**: Send and receive chat messages using the core SDK
- **Response Generation**: Build intelligent auto-responses
- **End-to-End Encryption**: Automatic message encryption via the SDK

## 🚀 Quick Start

### Go
```bash
cd go
go run main.go
```

### Java
```bash
cd java
# Coming soon - Java implementation
```

Then connect from another Self SDK instance (mobile app, another SDK client, or the Self app) to start chatting!

### Expected Output

```
💬 Simple Chat Demo
===================

🔧 Setting up Self account...
✅ Self account created successfully
✅ Connected to Self network

📬 Connection Information:
   🆔 Inbox Address: did:self:abc123...
   📱 To connect: Use this address in another Self SDK instance
   🔐 All messages will be automatically encrypted

✅ Chat demo ready! Press Ctrl+C to exit.
```

### During Conversation

```
📨 [15:04:07] did:self:peer456: "Hello there!"
📤 [15:04:07] Sent: "👋 Hello! Message received at 15:04:07"

📨 [15:04:15] did:self:peer456: "How are you?"
📤 [15:04:15] Sent: "🤖 I'm doing great! I'm a Self SDK chat demo."

📨 [15:04:23] did:self:peer456: "What can you do?"
📤 [15:04:23] Sent: "🔄 Echo: What can you do?"
```

## 🏗️ How It Works

### Step 1: Account Creation & Configuration

**Conceptual Setup:**
```
Create Self Account:
  - Generate unique identity and cryptographic keys
  - Register with Self network
  - Configure event callbacks for message handling
  - Initialize encrypted storage
```

**What happens here:**
- Creates a new Self identity with cryptographic keys
- Registers with the Self network
- Sets up event callbacks for real-time message handling
- Configures encrypted storage for account data

### Step 2: Message Handling

**Message Processing Architecture:**
```
Incoming Message Flow:
  1. SDK receives encrypted message
  2. Automatically decrypts using account keys
  3. Determines message content type
  4. Routes to appropriate handler
  5. Triggers user-defined callback
```

**Message Processing Flow:**
1. **Message Arrival**: SDK automatically decrypts incoming messages
2. **Type Detection**: Identifies chat messages vs other content types
3. **Routing**: Directs chat messages to specialized handler
4. **Response**: Generates and sends appropriate reply

### Step 3: Smart Response Generation

**Response Logic:**
```
Message Analysis:
  - Extract text content from message
  - Apply pattern matching rules
  - Generate contextual response
  - Send reply back to sender

Pattern Examples:
  - Greetings: "hello", "hi" → Welcome message
  - Status: "how are you" → Bot status
  - Help: "help" → Usage instructions
  - Time: "time" → Current timestamp
  - Default: Echo the original message
```

**Response Patterns:**
- **Greeting Detection**: Responds to "hello", "hi" with welcome message
- **Status Inquiry**: Handles "how are you" with bot status
- **Help Commands**: Provides usage instructions
- **Time Requests**: Returns current timestamp
- **Echo Fallback**: Repeats unrecognized messages

## 🔍 Code Architecture

### Core Components

| Component | Purpose | Key Concepts |
|-----------|---------|--------------|
| **Account Setup** | Demo orchestration and account creation | Account configuration, callbacks |
| **Message Router** | Message routing and type detection | Event-driven processing, content types |
| **Chat Handler** | Chat-specific message processing | Message decoding, error handling |
| **Response Engine** | Intelligent response creation | Pattern matching, context awareness |
| **Message Sender** | Message building and delivery | Message composition, error handling |

### Key SDK Features Used

**Account Management:**
- **Account Interface**: Core identity and messaging functionality
- **Event Callbacks**: Event-driven message handling system
- **Setup Utilities**: Simplified account creation and configuration

**Message Processing:**
- **Message Events**: Incoming message container and metadata
- **Content Types**: Chat message type identification
- **Content Decoders**: Chat content extraction and parsing

**Message Creation:**
- **Message Builders**: Chat message composition utilities
- **Send Interface**: Message delivery and status tracking

## 🎓 What Just Happened

When you run this example:

1. **Identity Creation**: The SDK generates a unique DID and cryptographic keypair
2. **Network Registration**: Your account registers with the Self network
3. **Inbox Setup**: An inbox address is created for receiving messages
4. **Event Loop**: The program enters a listening state for incoming messages
5. **Message Processing**: Each incoming message triggers the callback chain
6. **Auto-Response**: The bot analyzes messages and sends intelligent replies

**Security Features:**
- All messages are end-to-end encrypted automatically
- Your private keys never leave your device
- Communication uses Self's decentralized infrastructure

## 🔧 Customization Ideas

### Extend Response Patterns

**Weather Responses:**
```
Pattern: "weather" → "🌤️ I can't check weather, but it's always sunny in the Self network!"
```

**Humor Responses:**
```
Pattern: "joke" → "🤪 Why did the developer go broke? Because they used up all their cache!"
```

### Add Message Persistence

**Store Conversations:**
```
For each message:
  - Extract sender ID and content
  - Store in SDK storage with timestamp
  - Enable conversation history retrieval
```

### Create Custom Commands

**Command Processing:**
```
If message starts with "/":
  - Extract command name
  - Parse command parameters
  - Execute command logic
  - Return command result
```

## 🎯 Real-World Applications

### Customer Support Bot
- **Use Case**: Automated customer service responses
- **Features**: Pattern-based responses, escalation triggers
- **Response Types**: Greetings, FAQ answers, contact information

### Team Notification Bot
- **Use Case**: Send automated updates to team members
- **Features**: Event-triggered messages, status broadcasts
- **Message Types**: Alerts, reminders, status updates

### Integration Hub
- **Use Case**: Bridge between Self SDK and external systems
- **Features**: API integration, data transformation
- **Capabilities**: Webhook handling, service notifications

## 🔐 Security Considerations

### Message Encryption
- **End-to-End**: All messages encrypted between sender and receiver
- **Key Management**: SDK handles cryptographic key lifecycle
- **Forward Secrecy**: Past messages remain secure even if keys are compromised

### Identity Verification
- **DID-Based**: Each participant has cryptographically verifiable identity
- **Anti-Spoofing**: Messages cannot be forged or impersonated
- **Network Security**: Self network provides secure message routing

## 🚀 Next Steps

After mastering basic chat:

1. **🔗 Connection Examples**: Learn how to establish connections → `../../01_connection`
2. **🎫 Credentials Integration**: Combine chat with credential verification → `../../02_credentials`
3. **⚙️ Advanced Features**: Explore rich messaging features → `../../04_advanced_features`
4. **📱 Mobile Integration**: Connect with mobile apps → `../../../mobile`

## 🔧 Troubleshooting

### No Messages Received
```
⏳ Waiting for messages... (no activity)
```
**Solution**: 
- Verify another client has the correct inbox address
- Check network connectivity
- Ensure account is properly initialized

### Message Send Failed
```
❌ Failed to send message
```
**Solution**: 
- Verify connection to recipient exists
- Check network connectivity
- Ensure recipient address is valid

### Account Setup Failed
```
❌ Failed to create Self account
```
**Solution**: 
- Check internet connection
- Verify storage permissions
- Ensure SDK dependencies are available

---

**Ready to build intelligent chat experiences?** This foundation enables powerful messaging workflows! 💬
