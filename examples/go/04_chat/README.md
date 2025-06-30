# Simple Chat Example 🟢

A straightforward demonstration of the Self SDK's chat messaging capabilities using the core SDK. This example focuses on essential chat functionality with clean, production-ready code patterns.

## 🚀 Quick Start

```bash
go run main.go
```

Then connect from another Self SDK instance (mobile app, another Go client, or the Self app) to start chatting!

## 🎯 What You'll Learn

- **Account Setup**: Create and configure a Self account for messaging
- **Event-Driven Architecture**: Handle incoming messages with callbacks
- **Chat Messaging**: Send and receive chat messages using the core SDK
- **Response Generation**: Build intelligent auto-responses
- **End-to-End Encryption**: Automatic message encryption via the SDK

## 🏗️ How It Works

### Step 1: Account Creation & Configuration
```go
selfAccount := common.SetupAccount(common.AccountConfig{
    Callbacks: account.Callbacks{
        OnMessage: handleMessage,
        OnConnect: func(acc *account.Account) {
            fmt.Println("✅ Connected to Self network")
        },
    },
})
```

**What happens here:**
- Creates a new Self identity with cryptographic keys
- Registers with the Self network
- Sets up event callbacks for real-time message handling
- Configures encrypted storage in `./storage/`

### Step 2: Message Handling
```go
func handleMessage(acc *account.Account, msg *event.Message) {
    timestamp := time.Now().Format("15:04:05")
    
    switch event.ContentTypeOf(msg) {
    case message.ContentTypeChat:
        handleChatMessage(acc, msg, timestamp)
    default:
        fmt.Printf("📨 [%s] Received message from %s\n",
            timestamp, msg.FromAddress().String())
    }
}
```

**Message Processing Flow:**
1. **Message Arrival**: SDK automatically decrypts incoming messages
2. **Type Detection**: Identifies chat messages vs other content types
3. **Routing**: Directs chat messages to specialized handler
4. **Response**: Generates and sends appropriate reply

### Step 3: Smart Response Generation
```go
func generateResponse(messageText, timestamp string) string {
    message := strings.ToLower(strings.TrimSpace(messageText))
    
    switch {
    case strings.Contains(message, "hello") || strings.Contains(message, "hi"):
        return fmt.Sprintf("👋 Hello! Message received at %s", timestamp)
    case strings.Contains(message, "how are you"):
        return "🤖 I'm doing great! I'm a Self SDK chat demo."
    // ... more patterns
    default:
        return fmt.Sprintf("🔄 Echo: %s", messageText)
    }
}
```

**Response Patterns:**
- **Greeting Detection**: Responds to "hello", "hi" with welcome message
- **Status Inquiry**: Handles "how are you" with bot status
- **Help Commands**: Provides usage instructions
- **Time Requests**: Returns current timestamp
- **Echo Fallback**: Repeats unrecognized messages

## 📋 Expected Output

### Initial Setup
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

## 🔍 Code Architecture

### Function Breakdown

| Function | Purpose | Key Concepts |
|----------|---------|--------------|
| `main()` | Demo orchestration and account setup | Account configuration, callbacks |
| `handleMessage()` | Message routing and type detection | Event-driven processing, content types |
| `handleChatMessage()` | Chat-specific message processing | Message decoding, error handling |
| `generateResponse()` | Intelligent response creation | Pattern matching, context awareness |
| `sendResponse()` | Message sending and delivery | Message building, error handling |

### Key SDK Components Used

**Account Management:**
- `account.Account`: Core identity and messaging interface
- `account.Callbacks`: Event-driven message handling
- `common.SetupAccount()`: Simplified account creation

**Message Processing:**
- `event.Message`: Incoming message container
- `message.ContentTypeChat`: Chat message type identifier
- `message.DecodeChat()`: Chat content extraction

**Message Creation:**
- `message.NewChat()`: Chat message builder
- `acc.MessageSend()`: Message delivery interface

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
```go
// Add to generateResponse() function
case strings.Contains(message, "weather"):
    return "🌤️ I can't check weather, but it's always sunny in the Self network!"
case strings.Contains(message, "joke"):
    return "🤪 Why did the developer go broke? Because they used up all their cache!"
```

### Add Message Persistence
```go
// Store conversations in SDK storage
func storeMessage(sender, content string) {
    // Use Self SDK storage APIs to persist chat history
}
```

### Create Custom Commands
```go
// Handle /commands
if strings.HasPrefix(message, "/") {
    return handleCommand(strings.TrimPrefix(message, "/"))
}
```

## 💡 Troubleshooting

### Common Issues

**Build Errors:**
```bash
# Ensure dependencies are up to date
go mod tidy

# Check Go version (requires 1.22+)
go version
```

**Runtime Issues:**
- **Storage permissions**: Ensure write access to `./storage/` directory
- **Network connectivity**: Check internet connection for sandbox environment
- **Account creation fails**: Verify Self C SDK is properly installed

**Connection Problems:**
- **Can't connect from mobile**: Ensure both devices use same environment (sandbox/production)
- **Messages not arriving**: Check that callback functions are properly registered
- **Decryption errors**: Verify proper connection establishment before messaging

### Debug Mode
```bash
# Enable verbose logging
export SELF_LOG_LEVEL=debug
go run main.go
```

## 🚀 Next Steps

After mastering basic chat, explore these progressively complex examples:

| Example | Complexity | Skills Gained |
|---------|------------|---------------|
| **Connection Management** 🟡 | Direct peer connection patterns |
| **Group Chat** 🟡 | Multi-participant messaging, roles |
| **Credential Exchange** 🟠 | Identity verification, VCs |
| **File Sharing** 🟠 | Attachment handling, large data |

## 🛠️ Prerequisites

- **Go**: Version 1.22 or later
- **Self C SDK**: Native dependency for core functionality
- **Storage**: Write permissions for account data persistence
- **Network**: Internet access for Self network connectivity

## ⚡ Performance Notes

- **Memory**: Minimal footprint, suitable for server deployment
- **CPU**: Low overhead, event-driven processing
- **Storage**: Encrypted local storage grows with message history
- **Network**: Efficient binary protocol with automatic compression
