# Core Advanced Features Example 🔴

> **📖 Learn the concepts first:** [System Architecture](../../../../docs/architecture/system-overview.md)

A comprehensive demonstration of the Self SDK's core advanced capabilities using the core SDK directly. This example showcases production-ready patterns for encrypted storage, event-driven messaging, and advanced account management.

## 🚀 Quick Start

```bash
# Run the advanced features demo
go run main.go
```

The demo automatically showcases core Self SDK features in a working demonstration!

## 📊 Complexity Rating

**7/10** (Advanced) 🔴 - Perfect for learning production-ready Self SDK patterns

- 🔴 **Core SDK integration**: Direct SDK usage with complete control
- 🟠 **Production patterns**: Account configuration, storage management, connection handling
- 🟠 **Real-world features**: Discovery QR codes, message processing, graceful shutdown
- 🟡 **Advanced concepts**: Encrypted storage, event handling, resource management

## 🎯 What This Example Demonstrates

### Core Advanced Features
- ✅ **Core SDK Account Setup** - Direct account creation with advanced configuration
- ✅ **Encrypted Storage** - Secure data persistence using account.ValueStore/ValueLookup
- ✅ **Discovery & Connection** - QR code generation and peer connection handling
- ✅ **Event-Driven Messaging** - Real-time message processing with callbacks
- ✅ **Production Patterns** - Graceful shutdown, error handling, statistics tracking
- ✅ **Storage Management** - JSON serialization and advanced data operations

### Educational Learning Path
1. **Core SDK Setup** - Advanced account configuration with callbacks
2. **Storage Operations** - Encrypted storage with JSON data handling
3. **Discovery System** - QR code generation for peer connections
4. **Message Handling** - Event-driven architecture with real-time processing
5. **Production Patterns** - Proper resource management and monitoring

## 🏃‍♂️ How to Run

### Single Command Demo
```bash
go run main.go
```

The demo runs automatically and demonstrates:
- Advanced account creation with core SDK
- Encrypted storage operations with JSON data
- Discovery QR code generation for peer connections
- Real-time message handling with automatic responses
- Production-ready patterns with graceful shutdown

### What Happens Automatically
1. **Account Creation**: Core SDK account with advanced configuration and callbacks
2. **Storage Demo**: JSON data storage and retrieval with encryption
3. **Discovery QR**: Interactive QR code for peer connections
4. **Message Processing**: Real-time event handling with automatic responses
5. **Monitoring**: Status updates and connection/message statistics
6. **Graceful Shutdown**: Clean exit with Ctrl+C and summary statistics

## 📋 What You'll See

```
🚀 Advanced Self SDK Features Demo (Core SDK)
==============================================
This demo showcases advanced Self SDK capabilities using the core SDK directly.

🔧 Setting up advanced account with core SDK...
✅ Advanced account created successfully
🆔 Account DID: did:self:advanced123...

💾 Advanced Storage Capabilities
================================
🔹 Encrypted Storage Operations
   ✅ Stored user_profile
   ✅ Stored app_settings
   ✅ Stored session_token
   ✅ Stored last_activity
   ✅ Stored connection_count
🔹 Data Retrieval Verification
   📖 Retrieved profile: Advanced User (Developer)
   🎨 Dark mode: true, Notifications: true
   ✅ Storage operations completed successfully

🔍 Discovery & Connection Setup
===============================
📱 QR Code for Discovery:
========================
█▀▀▀▀▀█ ▀▀█▄ ▀▀▀ █▀▀▀▀▀█
█ ███ █ █▄▀▀▀▄█ █ ███ █
█ ▀▀▀ █ ▄▀▀ ▀▀█ █ ▀▀▀ █
▀▀▀▀▀▀▀ █▄█▄▀▄█ ▀▀▀▀▀▀▀
========================

🎯 Scan this QR code with the Self app to connect!
⏰ Expires in 15 minutes (15:30:00)

💡 What happens when someone connects:
   • Secure connection established
   • Welcome message sent automatically
   • Advanced storage updated with connection info
   • Demo showcases real-time message handling

⏳ Waiting for connections and messages...
   • Scan the QR code above to connect
   • Send messages to see advanced features
   • Press Ctrl+C to finish the demo

📊 Status: 0 connections, 0 messages processed (runtime: 30s)

🎉 New connection established!
👤 Peer: did:self:peer456...
🕐 Time: 15:32:15

💾 Connection info stored in advanced storage
📤 Advanced welcome message sent

📨 Message #1 received
👤 From: did:self:peer456...
💬 Content: Hello advanced demo!
🕐 Time: 15:32:45
💾 Message stored in advanced storage
📤 Advanced response sent (Message #1)

📊 Status: 1 connections, 1 messages processed (runtime: 2m15s)

🛑 Shutdown signal received...
🏁 Demo completed by user

📊 Advanced Demo Summary
========================
⏱️  Total runtime: 2m18s
🔗 Connections made: 1
📨 Messages processed: 1
💾 Storage operations: 8+

🎓 What was demonstrated:
   ✅ Core SDK account creation and configuration
   ✅ Advanced encrypted storage with automatic security
   ✅ Discovery QR generation and connection handling
   ✅ Real-time event-driven message processing
   ✅ Production-ready account and storage patterns
   ✅ Graceful shutdown and resource management

🚀 Advanced features unlocked:
   • Direct core SDK usage with complete control
   • Encrypted storage for secure data persistence
   • Event-driven architecture with callbacks
   • Real-time connection and message handling
   • Production patterns for robust applications

📚 Next steps:
   • Explore individual subdirectories for deep dives
   • Build production applications using these patterns
   • Integrate advanced features into your own projects

✅ Advanced features demo completed!
```

## 🔍 Key Code Sections

| Function | Lines | Purpose |
|----------|-------|---------|
| `main()` | 50-65 | Demo orchestration and setup |
| `setupAdvancedAccount()` | 85-115 | Core SDK account creation with callbacks |
| `demonstrateAdvancedStorage()` | 120-170 | Encrypted storage operations with JSON |
| `generateDiscoveryQR()` | 175-220 | QR code generation for peer connections |
| `onWelcomeMessage()` | 225-270 | Connection acceptance and welcome handling |
| `onMessage()` | 285-320 | Real-time message processing and storage |
| `sendAdvancedWelcomeMessage()` | 325-350 | Advanced welcome message creation |
| `sendAdvancedResponse()` | 365-395 | Intelligent auto-response system |
| `waitForInteractions()` | 410-440 | Interactive demo with status monitoring |
| `displaySummary()` | 445-485 | Final statistics and educational summary |

## 🎓 Educational Notes

### Core SDK Integration
- **Direct Usage**: Complete control with direct account interaction
- **Callback Architecture**: Event-driven design with OnWelcome, OnMessage, OnKeyPackage
- **Account Configuration**: Advanced setup with storage encryption and environment
- **Resource Management**: Proper account.Close() and graceful shutdown

### Advanced Storage Patterns
- **JSON Serialization**: Convert complex data to JSON for storage
- **Encrypted Persistence**: All data automatically encrypted using storage key
- **Value Operations**: Direct use of account.ValueStore() and account.ValueLookup()
- **Connection Tracking**: Store connection metadata for advanced features

### Discovery & Connection
- **Key Package Generation**: ConnectionNegotiateOutOfBand for QR connections
- **Discovery Requests**: Message builders with expiration and key packages
- **QR Code Creation**: Anonymous messages with sandbox flags for testing
- **Connection Acceptance**: Automatic acceptance with ConnectionAccept()

### Real-Time Message Processing
- **Event Callbacks**: OnMessage callback for all incoming messages
- **Message Filtering**: Only processes chat messages to avoid system message flooding
- **Message Decoding**: Proper handling of chat vs non-chat messages
- **Auto-Response**: Intelligent responses showing demo statistics
- **Message Storage**: Persistent message history with metadata

### Production Benefits
- **Error Handling**: Comprehensive error checking and graceful degradation
- **Monitoring**: Real-time status updates and connection tracking
- **Statistics**: Detailed summary with runtime and operation counts
- **Resource Cleanup**: Proper shutdown with signal handling

## 🔧 Customization Ideas

Try modifying the code to:
- Add different message types beyond chat messages
- Implement custom storage namespaces for organization
- Create advanced connection filtering and management
- Add file attachment handling for richer messaging
- Implement connection persistence across restarts
- Add custom analytics and monitoring
- Create advanced auto-response patterns

## 🚀 Next Steps

After understanding this example, you're ready for:

| Next Level | Complexity | Description |
|------------|------------|-------------|
| **Production Apps** | 8-9/10 | Build real applications using these core patterns |
| **Custom Features** | 8/10 | Add custom functionality using core SDK |
| **Advanced Integration** | 9/10 | Integrate with existing systems and databases |

## 🛠️ Prerequisites

- Go 1.19 or later
- Self SDK v0.59.0+ (handled by go.mod)
- Understanding of previous examples (simple_chat, credential_issuance, discovery)
- Basic knowledge of JSON and encryption concepts

## 💡 Troubleshooting

**Account Creation Issues:**
- Ensure storage path is writable (`./advanced_demo_storage`)
- Check that storage key generation is working
- Verify environment is set to sandbox for testing

**Storage Problems:**
- Confirm JSON marshaling/unmarshaling is working
- Check that ValueStore/ValueLookup methods exist
- Verify storage key is 32 bytes and properly generated

**Connection Issues:**
- Ensure QR code is properly formatted and scannable
- Check that key package generation succeeds
- Verify sandbox environment flags are set correctly

**Message Handling Problems:**
- Confirm callback signatures match SDK requirements
- Check that message decoding is handling different types
- Verify MessageSend is using correct content format

**Build Issues:**
- Run `go mod tidy` to ensure dependencies
- Check Go version with `go version` (need 1.19+)
- Verify you're in the correct directory with go.mod

## 🎯 Core SDK Capabilities

This example demonstrates the full power of the Self SDK with direct core integration:

| Feature | Implementation | Benefits |
|---------|---------------|----------|
| **Account Setup** | `account.New()` with config | Complete control over configuration |
| **Storage** | `account.ValueStore/ValueLookup()` | Direct encrypted storage access |
| **Messages** | `account.MessageSend()` with builders | Full message customization |
| **Events** | Direct account callbacks | Real-time event processing |
| **Complexity** | **Full SDK control** | Maximum flexibility |
| **Flexibility** | **Complete customization** | Production-ready patterns |

## 🏗️ Architecture Patterns Demonstrated

### Core SDK Architecture
```
account.Account
├── Callbacks (OnWelcome, OnMessage, OnKeyPackage)
├── Storage (ValueStore, ValueLookup)
├── Messaging (MessageSend, ConnectionAccept)
└── Discovery (ConnectionNegotiateOutOfBand)
```

### Event-Driven Flow
```
QR Scan → OnWelcome → ConnectionAccept → Welcome Message
Message → OnMessage → Process → Store → Auto-Response
```

This example provides the foundation for building production applications with direct core SDK usage! 🚀
