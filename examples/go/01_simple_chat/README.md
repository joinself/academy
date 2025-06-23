# Simple Chat Example (Using Core SDK)

A straightforward demonstration of the underlying Self SDK's core chat capabilities, designed for educational purposes. This example uses the core Self SDK directly and focuses on essential chat functionality.

## 🚀 Quick Start

```bash
# Start the chat server
go run main.go
```

Then connect from another Self SDK instance to start chatting!

## 📊 Complexity Rating

**4/10** (Simple) - Perfect for understanding basic Self SDK concepts

- 🟢 **Clear flow**: Simple account setup and message handling
- 🟢 **Core concepts**: Account, messages, and callbacks
- 🟡 **Event-driven**: Callback-based message processing
- 🟡 **SDK usage**: Direct use of underlying SDK components

## 🎯 What This Example Demonstrates

### Core SDK Features
- ✅ **Account Management** - Simple Self account creation and configuration
- ✅ **Event System** - Callback-based message handling
- ✅ **Chat Messaging** - Send and receive chat messages directly
- ✅ **End-to-end encryption** - Automatic message encryption via the SDK
- ✅ **Inbox Management** - Open inbox addresses for receiving messages

### Educational Learning Path
1. **Account Setup** - Initialize Self SDK account with essential callbacks
2. **Inbox Management** - Open inbox addresses for message receiving
3. **Event Handling** - Process incoming chat messages
4. **Response Generation** - Build and send automatic responses

## 🏃‍♂️ How to Run

### Step 1: Start the Chat Server
```bash
go run main.go
```

You'll see:
- Account initialization with core SDK
- Inbox address for connections
- Event handler setup

### Step 2: Connect and Chat
- Use the displayed inbox address to connect from another Self SDK instance
- Send messages to see automatic responses
- All communication is end-to-end encrypted

## 📋 What You'll See

```
💬 Simple Chat Demo (Core SDK)
===============================
This demo shows basic chat messaging using the underlying Self SDK.

🔧 Setting up Self account...
✅ Self account created successfully
🔗 Connected to Self network

📬 Connection Information:
   🆔 Inbox Address: did:self:example123...
   📱 To connect: Use this address in another Self SDK instance
   🔐 All messages will be automatically encrypted

✅ Chat demo ready!

🎓 What's running:
   • Self account with event-driven message handling
   • Automatic chat responses to incoming messages
   • End-to-end encrypted communication

💡 To test: Connect from another Self SDK instance and send messages!
Press Ctrl+C to exit.

📨 [15:04:07] Chat message from did:self:peer456:
   💬 "Hello there!"
📤 [15:04:07] Sending response...
✅ [15:04:07] Response sent: "👋 Hello! Message received at 15:04:07 via Self SDK"
```

## 🔍 Key Code Sections

| Function | Lines | Purpose |
|----------|-------|---------|
| `main()` | 35-55 | Main demo orchestration |
| `setupAccount()` | 60-85 | Account creation and configuration |
| `showConnectionInfo()` | 90-100 | Display connection details |
| `handleMessage()` | 105-115 | Event-driven message routing |
| `handleChatMessage()` | 120-135 | Chat message processing and response |
| `generateResponse()` | 140-155 | Smart response generation |
| `sendResponse()` | 160-175 | Message sending |

## 🎓 Educational Notes

### Core SDK Concepts
- **Account**: The fundamental Self SDK entity for identity and messaging
- **Events**: Callback-based system for handling incoming messages
- **Inbox**: Address where others can send you messages
- **Content Types**: Different message types (focused on chat in this example)

### Key Patterns
- **No global variables**: Clean function parameter passing
- **Event-driven**: Callback-based instead of polling
- **Error handling**: Proper error checking throughout
- **Simple configuration**: Minimal setup for basic functionality

### Real-time Features
- **Event callbacks** for instant message processing
- **Automatic encryption** for all messages via core SDK
- **Smart responses** based on message content

## 🔧 Customization Ideas

Try modifying the code to:
- Add new response patterns in `generateResponse()`
- Implement message persistence using SDK storage APIs
- Add logging of all conversations
- Create custom commands (like `/time`, `/help`)
- Add message formatting with emojis or markdown

## 🚀 Next Steps

After understanding this core SDK example, explore:

| Example | Complexity | Description |
|---------|------------|-------------|
| Connection Management | 5/10 | Establish connections between peers |
| Credential Exchange | 7/10 | Issue and verify credentials using core APIs |
| Group Chat | 6/10 | Multi-participant encrypted conversations |
| File Sharing | 8/10 | Send files and attachments |

## 🛠️ Prerequisites

- Go 1.22 or later
- Self C SDK (native dependency)
- Basic understanding of Go functions and error handling

## 💡 Troubleshooting

**Build Issues:**
- Ensure Self C SDK is properly installed
- Run `go mod tidy` to resolve dependencies
- Check Go version with `go version`

**Runtime Issues:**
- Check storage directory permissions (creates `./storage/` folder)
- Ensure storage key is properly configured
- Verify network connectivity for sandbox environment

**Connection Issues:**
- Share the inbox address with another Self SDK instance
- Ensure both instances are using the same environment (sandbox)
- Check that callback functions are properly registered
