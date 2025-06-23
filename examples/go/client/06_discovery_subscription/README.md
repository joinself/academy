# Simple Discovery Demo

This example demonstrates **peer discovery and connection** using the core Self SDK directly. It shows a clean, focused workflow: generate one QR code, wait for a peer to connect, send them a welcome message, and complete the demo.

> **📢 Simplified Focus**: This example uses the underlying Self SDK directly (v0.59.0+) to demonstrate the essential discovery workflow without complexity. It's perfect for understanding the core concepts before building more advanced applications.

## 🚀 Quick Start

| 🎯 Goal | 🏃‍♂️ Command | ⏱️ Time |
|---------|-------------|---------|
| **Run discovery demo** | `go run main.go` | 2-3 min |
| **Test with another client** | Scan QR code with Self SDK | Real-time connection |
| **Learn core concepts** | Run + read guide | 5-10 min |

## 📚 What You'll Experience

### 🎯 Simple Discovery Workflow

- **📱 Single QR Code Generation** - Creates one discovery endpoint (30-minute timeout)
- **⏳ Connection Waiting** - Waits for one peer to scan and connect  
- **🎉 Automatic Connection** - Accepts the connection and sends welcome message
- **✅ Clean Completion** - Demo finishes after successful message delivery

### 🔄 Step-by-Step Process

```
1. Account Setup
   └── Create account with Welcome event handler
   └── Display account DID for reference

2. QR Code Generation  
   └── Generate key package for connection (30-minute validity)
   └── Create discovery request with embedded key package
   └── Encode as QR code and display

3. Connection Handling
   └── Wait for peer to scan QR code
   └── Automatically accept incoming connection
   └── Send welcome message to connected peer

4. Demo Completion
   └── Show summary of successful connection
   └── Exit cleanly after message delivery
```

### 📱 Demo Output Example

```
🔍 Simple Discovery Demo
=========================
This demo shows basic discovery workflow:
• Generate one QR code for connection
• Wait for a peer to scan and connect
• Send a welcome message
• Complete the demo

🔧 Setting up discovery account...
🔗 Connected to Self network
✅ Account created successfully
🆔 Your DID: 0056301398177fde1b53fa84aae917eba906f651559076c6665020f79221b87bcc

📱 Generating discovery QR code...
--- Discovery QR Code ---
Valid for: 30 minutes

[QR CODE DISPLAYED]

✅ QR code generated successfully!

⏳ Waiting for peer connection...
📱 Scan the QR code above with a Self client to connect
🔄 Press Ctrl+C to cancel

🤝 Connection request from: 001234...
🎉 Connected to peer: 001234...
📤 Sending welcome message...
✅ Welcome message sent successfully!
✅ Demo completed successfully!

📋 Demo Summary
================
👤 Connected peer: 001234...
💬 Welcome message sent

🎓 What was demonstrated:
   • QR code generation for discovery
   • Automatic connection acceptance
   • Direct peer messaging
   • Clean demo completion
```

## 🎓 Learning Outcomes

After running this simple demo, you'll understand:

### Discovery Fundamentals
- ✅ **QR code generation** using `ConnectionNegotiateOutOfBand()` and `EncodeToQR()`
- ✅ **Connection handling** with `OnWelcome` callback and `ConnectionAccept()`
- ✅ **Message sending** using `MessageSend()` with chat content
- ✅ **Clean demo patterns** with proper setup, execution, and completion

### Core SDK Concepts
- ✅ **Account configuration** with event callbacks for connection handling
- ✅ **Discovery requests** containing key packages for secure connection establishment
- ✅ **Event-driven architecture** using callbacks instead of polling or blocking
- ✅ **Message creation** with `message.NewChat()` and content encoding

### Essential Patterns
- ✅ **Single-purpose demos** that focus on one core concept
- ✅ **Graceful completion** with timeout handling and clean exit
- ✅ **Error handling** with appropriate logging and user feedback
- ✅ **Resource management** with proper cleanup and storage handling

## 🔧 Technical Implementation

### Core SDK Usage

```go
// Account setup with Welcome callback
cfg := &account.Config{
    StorageKey:  generateStorageKey("simple_discovery"),
    StoragePath: "./simple_discovery_storage", 
    Environment: account.TargetSandbox,
    Callbacks: account.Callbacks{
        OnWelcome: d.onWelcome, // Handle connection requests
    },
}
```

### QR Code Generation

```go
// Generate key package for connections
keyPackage, err := account.ConnectionNegotiateOutOfBand(
    inboxAddress,
    time.Now().Add(30*time.Minute), // 30-minute validity
)

// Create discovery request
discoveryContent, err := message.NewDiscoveryRequest().
    KeyPackage(keyPackage).
    Expires(time.Now().Add(30*time.Minute)).
    Finish()

// Encode as QR code
anonymousMsg := event.NewAnonymousMessage(discoveryContent)
qrCodeData, err := anonymousMsg.EncodeToQR(event.QREncodingUnicode)
```

### Connection Handling

```go
func (d *SimpleDiscoveryDemo) onWelcome(acc *account.Account, wlc *event.Welcome) {
    // Accept the connection
    _, err := acc.ConnectionAccept(wlc.ToAddress(), wlc.Welcome())
    
    // Send welcome message
    d.sendWelcomeMessage(wlc.FromAddress())
    
    // Signal completion
    d.done <- true
}
```

### Message Sending

```go
// Build and send chat message
chatContent, err := message.NewChat().
    Message("🎉 Welcome! Connection established!").
    Finish()

err = account.MessageSend(peerAddress, chatContent)
```

## 🚀 Getting Started

### Prerequisites

1. **Go 1.19 or later**
2. **Self SDK dependencies** (handled automatically by go.mod)
3. **Self client for testing** (mobile app, another SDK instance, web client)

### Running the Demo

```bash
# Run the simple discovery demo
go run main.go

# The program will:
# 1. Set up an account with connection handling
# 2. Generate and display one QR code (valid for 30 minutes)
# 3. Wait for a peer to scan the QR code and connect
# 4. Send a welcome message to the connected peer
# 5. Show a summary and exit cleanly
```

### 📱 Testing the Demo

1. **Start the demo** - Run `go run main.go` to see the QR code
2. **Scan the QR code** - Use any Self SDK client to scan and connect
3. **Watch the connection** - See real-time connection acceptance and message sending
4. **See the summary** - Demo completes with connection details

### Expected Behavior

- **QR code displays** immediately after account setup (valid for 30 minutes)
- **Connection event** appears when peer scans QR code  
- **Welcome message** is sent automatically to the connected peer
- **Demo completion** with summary showing successful connection and message delivery
- **Clean exit** after successful message sending

## 🎯 Use Cases & Applications

### Educational Applications

1. **Discovery Concept Learning**
   - Understand QR-based peer discovery
   - Learn connection establishment flow
   - See message sending in action
   - Foundation for more complex applications

2. **SDK Integration Testing**
   - Test Self SDK setup and configuration
   - Verify connection handling works correctly
   - Confirm message sending functionality
   - Debug connection issues

### Production Foundations

1. **Simple Connection Services**
   - One-time device pairing
   - Basic peer introduction
   - Quick message delivery
   - Temporary connection establishment

2. **Building Blocks for Complex Apps**
   - Foundation for multi-peer discovery
   - Base pattern for subscription services
   - Starting point for credential exchange
   - Core component for larger workflows

## 🔄 Simple vs. Advanced Discovery

### Simple Discovery Demo (This Example)

| Feature | Description | Focus |
|---------|-------------|-------|
| **Scope** | One QR code, one connection, one message | Essential concepts |
| **Complexity** | Minimal - easy to understand and follow | Learning foundation |
| **Duration** | Completes after first successful connection | Quick demonstration |
| **Use Case** | Learning discovery basics, simple pairing | Educational, basic apps |

### Advanced Discovery Applications

| Feature | Description | Focus |
|---------|-------------|-------|
| **Scope** | Multiple QR codes, multiple connections, ongoing messaging | Production patterns |
| **Complexity** | Full - connection management, message routing, error handling | Real-world applications |
| **Duration** | Runs continuously with connection pooling | Production services |
| **Use Case** | Event management, customer service, IoT platforms | Enterprise applications |

## 🔗 Next Steps & Learning Path

### Immediate Next Steps

After mastering this simple discovery demo:

1. **🔄 Connection Management** - Learn to handle multiple simultaneous connections
2. **💬 Message Exchange** - Add bidirectional messaging and conversation handling
3. **📋 Credential Exchange** - Use discovery for credential presentation workflows
4. **🏭 Production Patterns** - Add persistence, monitoring, and error recovery

### 🎓 Educational Progression

1. **This Simple Demo** - Basic discovery with one connection ← **You are here**
2. **Chat Examples** - Bidirectional messaging between discovered peers
3. **Credential Examples** - Using discovery for credential exchange workflows
4. **Group Examples** - Multi-peer discovery and group formation
5. **Advanced Examples** - Production-ready discovery services

### Code Evolution

```go
// Current Level: Simple Discovery
demo := NewSimpleDiscoveryDemo()
demo.Run() // Connect to one peer, send message, exit

// Next Level: Connection Management  
connectionManager := NewConnectionManager()
connectionManager.HandleMultiplePeers()
connectionManager.MaintainConnections()

// Advanced Level: Discovery Service
discoveryService := NewDiscoveryService()
discoveryService.HandleMultipleQRCodes()
discoveryService.RouteMessagesToHandlers()
discoveryService.PersistConnections()
```

## 💡 Why Start Simple?

### Focus on Core Concepts

This simplified approach helps you understand:

- **Essential discovery workflow** without distracting complexity
- **Core SDK patterns** that apply to all Self applications  
- **Connection fundamentals** before advanced connection management
- **Message basics** before complex routing and handling

### Building Solid Foundations

Starting with one connection and one message:

- **Reduces cognitive load** - focus on understanding rather than managing complexity
- **Establishes patterns** - learn the right way to use callbacks and event handling
- **Enables experimentation** - easy to modify and test different approaches
- **Provides confidence** - see immediate results and build understanding incrementally

## 📖 Technical Notes

### SDK Version & Dependencies
```go
// go.mod
module discovery_subscription
go 1.22
require github.com/joinself/self-go-sdk v0.59.0
```

### Storage & Environment
- Creates `simple_discovery_storage` directory for account data
- Configured for sandbox environment (safe for testing)
- Automatically cleans up on demo completion

### Timeouts & Limits
- **QR code validity**: 30 minutes (configurable)
- **Connection timeout**: Waits indefinitely until connection or manual cancellation
- **Demo completion**: Automatic exit after successful message delivery

## 🛠️ Troubleshooting

### Common Issues

1. **Storage Errors** - Delete `simple_discovery_storage` directory and retry
2. **QR Code Scanning Issues** - Ensure QR code is not expired and properly displayed
3. **Connection Failures** - Check that both peers are using compatible SDK versions
4. **Message Sending Errors** - Verify connection was established before message sending

### Debug Tips

```bash
# Clean storage and retry
rm -rf simple_discovery_storage && go run main.go

# Monitor demo execution
# Watch for connection events and message delivery confirmation

# Test with different clients
# Try scanning from mobile apps, web clients, or other SDK instances
```

### Expected Timing

- **Account setup**: 1-2 seconds
- **QR code generation**: 1-2 seconds  
- **Connection establishment**: 2-5 seconds after scanning
- **Message delivery**: 1-2 seconds after connection
- **Demo completion**: Immediate after message delivery

## 🤝 Contributing

If you have suggestions for improving this simple discovery demo or ideas for additional educational examples, please contribute back to the Self SDK project.

## 📖 Additional Resources

- [Self SDK Documentation](https://docs.joinself.com)
- [Discovery Protocol Specification](https://docs.joinself.com/discovery)
- [QR Code Standards](https://www.iso.org/standard/62021.html)
- [Event-Driven Programming Patterns](https://martinfowler.com/articles/201701-event-driven.html)

---

**Simple and effective! 🎉**

This simplified discovery demo provides a clean, focused introduction to peer discovery with the Self SDK. Master these essential concepts first, then progress to more advanced connection management and messaging patterns as your applications require them.
