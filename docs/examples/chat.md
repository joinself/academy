# **Chat:** Chat & Messaging Examples

> ****Documentation:** Learn the concepts first:** [Message Layer Security](../concepts/message-layer-security.md) | [Cryptographic Foundations](../concepts/cryptographic-foundations.md)

## What You'll Learn

Self SDK makes secure messaging surprisingly simple. Behind the scenes, every chat message benefits from [Ed25519 signatures](../concepts/cryptographic-foundations.md#digital-signatures), [MLS end-to-end encryption](../concepts/message-layer-security.md#encryption-properties), and [forward secrecy](../concepts/message-layer-security.md#forward-secrecy) - but you just focus on building great chat experiences.

****What you will learn:** Learning Goals:**
- Master event-driven message handling patterns
- Build intelligent chat bots with response systems  
- Understand message types and content routing
- Implement real-time secure communication
- Connect cryptographic theory to practical messaging

## **Architecture:** Progressive Learning Path

The chat examples follow our **Theory Meets Practice** approach, building from basic concepts to production-ready systems:

| **What you will learn:** Example | **Learning:** Complexity | **Resources:** Focus |  
|------------|---------------|----------|
| **🟢 [Basic Chat](#basic-chat-messaging)** | Beginner | Core messaging patterns, response generation |
| **🟡 Group Chat** | Intermediate | Multi-participant MLS groups *(Coming Soon)* |
| **🟠 Rich Media** | Advanced | File attachments, media sharing *(Coming Soon)* |
| **🔴 Production Features** | Expert | Typing indicators, read receipts *(Coming Soon)* |

## 🟢 Basic Chat Messaging

**Path:** [`examples/server/03_chat/01_basic/`](../../examples/server/03_chat/01_basic/)

Your first step into secure messaging. This example demonstrates core chat functionality with intelligent auto-responses, showing how the SDK handles encryption automatically while you focus on message logic.

### Quick Start

```bash
cd examples/server/03_chat/01_basic/go
go run main.go
```

**What happens:**
1. Creates a Self account with message callbacks
2. Displays your inbox address for connections
3. Listens for incoming chat messages
4. Responds intelligently based on message content

### Expected Output

```bash
**Chat:** Simple Chat Demo
===================

**Troubleshooting:** Setting up Self account...
**Success:** Connected to Self network

**Inbox:** Connection Information:
   🆔 Inbox Address: did:self:abc123...
   **Security:** All messages automatically encrypted

**Success:** Chat demo ready! Press Ctrl+C to exit.

# After someone sends "Hello!"
📨 [15:04:07] did:self:peer456: "Hello!"
📤 [15:04:07] Sent: "👋 Hello! Message received at 15:04:07"
```

### How It Works

#### 1. Event-Driven Message Handling

The foundation of all chat applications is responding to incoming messages. The SDK uses callbacks to notify your application:

```bash
67:82:examples/server/03_chat/01_basic/go/main.go
// handleMessage processes all incoming messages
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

****Key Concepts:** Key Concepts:**
- **Message Routing**: `event.ContentTypeOf()` determines message type (chat, credentials, etc.)
- **Type Safety**: Each content type has specific decoding methods
- **Event Loop**: Your application responds to messages as they arrive
- **Automatic Decryption**: The SDK handles all cryptographic operations

#### 2. Chat Message Processing

Once you know a message is a chat type, extract and process the content:

```bash
84:98:examples/server/03_chat/01_basic/go/main.go
// handleChatMessage processes incoming chat messages and sends responses
func handleChatMessage(acc *account.Account, msg *event.Message, timestamp string) {
	chat, err := message.DecodeChat(msg.Content())
	if err != nil {
		fmt.Printf("❌ [%s] Failed to decode chat message: %v\n", timestamp, err)
		return
	}

	fmt.Printf("📨 [%s] %s: \"%s\"\n", timestamp, msg.FromAddress().String(), chat.Message())

	response := generateResponse(chat.Message(), timestamp)
	sendResponse(acc, msg.FromAddress(), response, timestamp)
}
```

****Key Concepts:** Key Concepts:**
- **Content Decoding**: `message.DecodeChat()` extracts text from encrypted message
- **Sender Identification**: `msg.FromAddress()` provides cryptographic identity
- **Error Handling**: Always check for decoding failures
- **Response Pipeline**: Process → Generate → Send pattern

#### 3. Intelligent Response Generation  

Build smart bots that understand context and respond appropriately:

```bash
100:117:examples/server/03_chat/01_basic/go/main.go
// generateResponse creates appropriate responses based on the message content
func generateResponse(messageText, timestamp string) string {
	message := strings.ToLower(strings.TrimSpace(messageText))

	switch {
	case strings.Contains(message, "hello") || strings.Contains(message, "hi"):
		return fmt.Sprintf("👋 Hello! Message received at %s", timestamp)
	case strings.Contains(message, "how are you"):
		return "🤖 I'm doing great! I'm a Self SDK chat demo."
	case strings.Contains(message, "help"):
		return "💡 I'm a simple chat bot. Try saying 'hello', 'how are you', or send any message for an echo!"
	case strings.Contains(message, "time"):
		return fmt.Sprintf("🕐 Current time is %s", timestamp)
	default:
		return fmt.Sprintf("**Process:** Echo: %s", messageText)
	}
}
```

****Key Concepts:** Key Concepts:**
- **Pattern Matching**: Use string analysis for message understanding
- **Context Awareness**: Include timestamps and dynamic data
- **Fallback Responses**: Always have a default response
- **Extensible Design**: Easy to add new response patterns

#### 4. Secure Message Sending

Create and send responses using the SDK's message builders:

```bash
119:136:examples/server/03_chat/01_basic/go/main.go
// sendResponse sends a chat response to the peer
func sendResponse(acc *account.Account, toAddress *signing.PublicKey, responseText, timestamp string) {
	chatContent, err := message.NewChat().
		Message(responseText).
		Finish()
	if err != nil {
		fmt.Printf("❌ [%s] Failed to build response: %v\n", timestamp, err)
		return
	}

	err = acc.MessageSend(toAddress, chatContent)
	if err != nil {
		fmt.Printf("❌ [%s] Failed to send response: %v\n", timestamp, err)
	} else {
		fmt.Printf("📤 [%s] Sent: \"%s\"\n", timestamp, responseText)
	}
}
```

****Key Concepts:** Key Concepts:**
- **Message Builders**: `message.NewChat()` creates properly formatted content
- **Fluent API**: Chain methods for clean message construction
- **Automatic Encryption**: `MessageSend()` handles all cryptographic operations
- **Error Handling**: Always check for build and send failures

### **Security:** Security Features in Action

When you run this example, several security features work automatically:

#### End-to-End Encryption
```
Your Message → Ed25519 Signature → MLS Encryption → Network → MLS Decryption → Peer
```

- **Signing**: Every message signed with your Ed25519 private key
- **Encryption**: MLS protocol encrypts content with rotating keys  
- **Authentication**: Receiver cryptographically verifies your identity
- **Forward Secrecy**: Past messages remain secure even if current keys are compromised

#### Identity Verification
```
Message Source: did:self:abc123...
├── Cryptographic Identity (not just username)
├── Unforgeable (backed by Ed25519 signatures)  
└── Verifiable (anyone can validate the signature)
```

### **Learning:** What Just Happened?

Let's trace what happens when someone sends you "Hello!":

1. **Message Arrival**: Their Self SDK encrypts and signs the message
2. **Network Transport**: Message travels over Self's decentralized infrastructure  
3. **Automatic Decryption**: Your SDK decrypts using your private keys
4. **Callback Trigger**: `OnMessage` callback fires with decrypted content
5. **Type Detection**: `ContentTypeOf()` identifies this as a chat message
6. **Content Extraction**: `DecodeChat()` gets the text "Hello!"
7. **Response Generation**: Pattern matching produces a greeting response
8. **Response Encryption**: Your response gets signed and encrypted
9. **Delivery**: Response reaches the sender's device automatically

**The magic:** You wrote simple string processing code, but every message benefits from advanced cryptography, forward secrecy, and verifiable identity.

## **What you will learn:** Real-World Applications

### Customer Support Bot

Transform the basic pattern into an intelligent customer service system:

**Enhanced Response Patterns:**
```go
case strings.Contains(message, "refund"):
    return "💰 I can help with refunds! Please provide your order number."
case strings.Contains(message, "order") && containsOrderNumber(message):
    orderNum := extractOrderNumber(message)
    return fmt.Sprintf("📦 Looking up order %s...", orderNum)
case strings.Contains(message, "hours"):
    return "🕐 We're open Monday-Friday 9AM-5PM EST"
```

**Why Self SDK for Support:**
- **Verified Identity**: Know exactly who you're talking to
- **Message Integrity**: Guarantee messages aren't tampered with
- **Audit Trail**: Cryptographic proof of all conversations
- **Privacy**: End-to-end encryption protects sensitive data

### Team Notification Bot

Build an intelligent alert system for team communications:

**Notification Patterns:**
```go  
case strings.Contains(message, "deploy"):
    return "**Quick Start:** Deployment initiated! I'll notify the team of progress."
case strings.Contains(message, "alert"):
    severity := extractSeverity(message)
    return fmt.Sprintf("🚨 %s alert logged. Escalating to on-call team.", severity)
```

**Why Self SDK for Teams:**
- **Group Security**: MLS provides efficient group encryption
- **No Central Server**: Messages route through decentralized infrastructure  
- **Cross-Platform**: Works on mobile, web, server, IoT devices
- **Developer Friendly**: Simple APIs for complex security

### IoT Command Interface

Create secure device control with chat commands:

**Device Control Patterns:**
```go
case strings.HasPrefix(message, "/lights"):
    return handleLightingCommand(extractCommand(message))
case strings.HasPrefix(message, "/temp"):
    return handleTemperatureCommand(extractCommand(message))  
case strings.HasPrefix(message, "/status"):
    return getDeviceStatus()
```

**Why Self SDK for IoT:**
- **Device Identity**: Each device has cryptographic identity
- **Secure Commands**: No plaintext commands over networks
- **Access Control**: Only authorized users can send commands
- **Message Authentication**: Verify commands come from legitimate sources

## **Troubleshooting:** Customization Patterns

### Adding Message Types

Extend the routing system to handle different content types:

```go
switch event.ContentTypeOf(msg) {
case message.ContentTypeChat:
    handleChatMessage(acc, msg, timestamp)
case message.ContentTypeCredentialRequest:
    handleCredentialRequest(acc, msg, timestamp)  
case message.ContentTypeCredentialShare:
    handleCredentialShare(acc, msg, timestamp)
default:
    handleUnknownMessage(acc, msg, timestamp)
}
```

### Message Persistence

Store conversations for history and analytics:

```go
func storeMessage(sender, content, timestamp string) {
    messageRecord := map[string]interface{}{
        "sender":    sender,
        "content":   content, 
        "timestamp": timestamp,
        "type":      "incoming",
    }
    
    // Use SDK storage for encrypted persistence
    jsonData, _ := json.Marshal(messageRecord)
    acc.ValueStore(fmt.Sprintf("message_%s", timestamp), jsonData)
}
```

### Command Processing

Build sophisticated command interfaces:

```go
func processCommand(message string) string {
    if !strings.HasPrefix(message, "/") {
        return ""  // Not a command
    }
    
    parts := strings.Fields(message)
    command := parts[0][1:]  // Remove "/"
    args := parts[1:]
    
    switch command {
    case "help":
        return listAvailableCommands()
    case "status":
        return getSystemStatus()
    case "user":
        return handleUserCommand(args)
    default:
        return fmt.Sprintf("Unknown command: %s", command)
    }
}
```

## **Quick Start:** Advanced Chat Features (Coming Soon)

The next examples in the progression will demonstrate:

### 🟡 Group Chat
- **MLS Group Management**: Create and manage encrypted group conversations
- **Member Lifecycle**: Add/remove participants with forward secrecy
- **Group Identity**: Cryptographic group identity and verification
- **Efficient Encryption**: MLS scales to hundreds of participants

### 🟠 Rich Media Chat  
- **File Attachments**: Secure file sharing with integrity verification
- **Media Streaming**: Large file support with resumable transfers
- **Content Types**: Images, documents, audio, video support
- **Storage Integration**: Encrypted content storage and retrieval

### 🔴 Production Features
- **Typing Indicators**: Real-time activity notifications
- **Read Receipts**: Message delivery and read confirmations
- **Message History**: Persistent conversation storage
- **Offline Sync**: Message queuing and delivery when back online

## 🔗 Integration with Other Examples

Chat builds on and connects with other Academy examples:

### After Setup Examples
- **[New Account](setup.md#new-account-creation)**: Create the identity for chat
- **[Existing Account](setup.md#loading-existing-accounts)**: Restore chat identity from storage

### Before Chat Examples  
- **[Connection Examples](connections.md)**: Establish secure channels for messaging
- **[Credential Examples](credentials.md)**: Add identity verification to conversations

### Advanced Integration
- **[Advanced Features](advanced.md)**: Production-ready chat systems
- **[Notification Examples](advanced.md#notifications)**: Real-time chat alerts

## **Troubleshooting:** Need Help?

**Having chat issues?** Check our comprehensive **[Troubleshooting Guide](troubleshooting.md)** for solutions to messaging problems, including:

- **Chat:** **[Chat & Messaging Issues](troubleshooting.md#chat--messaging-issues)** - Message send failures, decode errors, and no messages received
- **Connections:** **[Connection Issues](troubleshooting.md#connection-issues)** - Connection problems that affect messaging
- **Network:** **[Network Issues](troubleshooting.md#network--connectivity-issues)** - Connectivity problems affecting chat operations

The troubleshooting guide includes detailed solutions, common causes, and debugging tips for all Self SDK examples.

## **Resources:** Next Steps

After mastering chat examples:

1. **🟡 [Message Layer Security Concepts](../concepts/message-layer-security.md)** - Understand the theory behind chat security
2. **🟠 [Credential Examples](credentials.md)** - Add identity verification to conversations  
3. **🔴 [Advanced Examples](advanced.md)** - Build production chat systems
4. **📱 Mobile Integration** - Connect with mobile chat apps

## 💡 Key Takeaways

- **Security by Default**: Every message automatically encrypted and signed
- **Event-Driven**: Build reactive applications with message callbacks
- **Identity-First**: Every participant has verifiable cryptographic identity  
- **Developer Friendly**: Simple APIs for complex security operations
- **Theory Meets Practice**: Academic cryptography made practically usable

The chat examples show how Self SDK makes advanced cryptography accessible - you focus on building great messaging experiences while the SDK handles Ed25519 signatures, MLS encryption, and forward secrecy automatically.

Ready to build the future of secure communication? **Chat:**
