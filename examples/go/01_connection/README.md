# 🔗 Connection Examples - Master Self SDK Connections

> **🎯 What you'll learn:** Complete connection patterns for Self SDK, from server setup to client connections to mobile integration

This directory contains **three complementary examples** that demonstrate the full spectrum of Self SDK connection patterns. Together, they show how to build complete client-server applications with encrypted peer-to-peer communication.

## 🔐 Why Connections Matter - The Foundation

**Before you can send messages, you must establish a connection.** Here's why:

### 🔑 Cryptographic Foundation
- **Shared Encryption Keys**: Each connection establishes unique encryption keys between parties
- **Identity Verification**: Connection process proves both parties are who they claim to be  
- **Secure Channel**: Creates an encrypted tunnel for all future communication
- **No Connection = No Security**: Without this handshake, there's no way to encrypt or verify messages

### 🛡️ The Connection Handshake
```
Party A                    Party B
├── 📧 "I want to connect" ──→ ├── 🔍 Verify identity
├── 🔐 Generate keys       ←── ├── 📧 "Connection accepted"  
├── ✅ Secure channel      ←→ ├── ✅ Secure channel
└── 💬 Ready for messaging ←→ └── 💬 Ready for messaging
```

**Think of it like exchanging phone numbers and setting up encryption - you need both before you can have a private conversation.**

## 🔗 Perfect Pairs: Server + Client Examples

**🎯 Best Learning Experience:** Run `01_direct` and `03_client` together!
- **`01_direct`** creates server addresses → **`03_client`** connects to those addresses
- Watch the complete connection handshake happen across two terminals
- Perfect demonstration of client-server Self SDK patterns

## 🟢 Complexity: Beginner
**Perfect starting point** - All examples are beginner-friendly and teach core Self SDK concepts.

---

## 🎯 Quick Decision Guide

**Choose your learning path based on your goals:**

| I want to... | Start with | Then try | Why this path? |
|--------------|------------|----------|----------------|
| 🔄 **Learn server-to-server** | [01_direct/](01_direct/) → [03_client/](03_client/) | [02_qr/](02_qr/) | Complete direct connection workflow |
| 📱 **Connect mobile apps** | [02_qr/](02_qr/) | [01_direct/](01_direct/) | User-friendly, visual connection |
| 🤖 **Build automated systems** | [01_direct/](01_direct/) → [03_client/](03_client/) | [02_qr/](02_qr/) | Programmatic, no user interaction |
| 👤 **Build user-facing apps** | [02_qr/](02_qr/) | [01_direct/](01_direct/) | QR scanning, mobile-first |
| 📚 **Understand everything** | [01_direct/](01_direct/) → [03_client/](03_client/) → [02_qr/](02_qr/) | - | Complete learning progression |

---

## 📊 The Three Connection Patterns

### 🔗 Direct Server (01_direct/) - "I Accept Connections"
**Creates inbox addresses for others to connect to**

```
🖥️  Direct Server
├── Creates inbox address
├── Shares address programmatically  
├── Handles OnKeyPackage callbacks
└── Accepts incoming connections
```

**Key Features:**
- ✅ **Server-side**: Waits for incoming connections
- ✅ **Address generation**: Creates shareable inbox addresses
- ✅ **OnKeyPackage**: Receives connection requests
- ✅ **Auto-accept**: Handles connections automatically

### 🔗 Direct Client (03_client/) - "I Connect to Addresses"  
**Connects to known inbox addresses created by servers**

```
🖥️  Direct Client
├── Takes inbox address as input
├── Initiates connection request
├── Handles OnWelcome callbacks
└── Completes connection handshake
```

**Key Features:**
- ✅ **Client-side**: Initiates connections to known addresses
- ✅ **Address consumer**: Uses addresses from servers
- ✅ **OnWelcome**: Receives connection acceptance
- ✅ **Connection completion**: Finalizes secure channel

### 📱 QR Code Server (02_qr/) - "I Display QR Codes"
**Generates QR codes for mobile apps to scan**

```
🖥️  QR Server  
├── Generates QR code
├── Displays for mobile scanning
├── Handles OnWelcome callbacks
└── Accepts mobile connections
```

**Key Features:**
- ✅ **Visual discovery**: Scannable QR codes for easy connection
- ✅ **Mobile-first**: Designed for mobile app interactions  
- ✅ **User-friendly**: Visual connection establishment
- ✅ **OnWelcome**: Uses Welcome message callback pattern

---

## 🚀 Quick Start

### 🔗 Try Complete Direct Connection Flow (Recommended!)
```bash
# Terminal 1: Start the server (creates address)
cd 01_direct && go run main.go
# Copy the displayed address

# Terminal 2: Connect as client (uses address)  
cd 03_client && go run main.go <paste-address-here>
# Watch connection establishment in both terminals
```

**💡 The `01_direct` and `03_client` examples are designed to work together!**
- **Server** (`01_direct`) creates inbox addresses and waits for connections
- **Client** (`03_client`) connects TO those addresses and completes the handshake
- **Result**: Complete client-server connection workflow demonstration

### Try QR Code Connections  
```bash
cd 02_qr
go run main.go
# Get a scannable QR code for mobile apps
```

### Run All Three to Compare
```bash
# Terminal 1: Direct server
cd 01_direct && go run main.go

# Terminal 2: Direct client (copy address from Terminal 1)
cd 03_client && go run main.go <address>

# Terminal 3: QR server
cd 02_qr && go run main.go
```

---

## 🎓 Learning Path

### 🌱 New to Self SDK? Start Here!
All three examples teach the same core concepts:
- **Self account creation** and configuration
- **Connection establishment** and handshake
- **Callback patterns** for handling events
- **Encrypted communication** channels

**Recommended starting path:**
1. **[01_direct/](01_direct/)** - Understand server-side address creation
2. **[03_client/](03_client/)** - Learn client-side connection initiation  
3. **[02_qr/](02_qr/)** - Explore mobile-friendly QR approach

### 🔄 Want to See the Complete Picture?
**Complete learning progression:**
1. **Direct Server** ([01_direct/](01_direct/)) - Creates and shares addresses
2. **Direct Client** ([03_client/](03_client/)) - Connects to those addresses
3. **QR Server** ([02_qr/](02_qr/)) - Alternative mobile-friendly approach
4. **Compare patterns** - Understand when to use each approach
5. **Move to advanced examples** (chat, credentials, etc.)

---

## 🔍 Technical Differences

| Aspect | Direct Server (01_direct/) | Direct Client (03_client/) | QR Server (02_qr/) |
|--------|--------------------------|-----------------------------|-------------------|
| **Role** | Creates & accepts | Initiates connections | Creates & accepts |
| **Input Required** | None | Inbox address | None |
| **Callback Used** | `OnKeyPackage` | `OnWelcome` | `OnWelcome` |
| **Setup Function** | `InboxOpen()` | `ConnectionNegotiate()` | `ConnectionNegotiateOutOfBand()` |
| **Target Use Case** | Server endpoints | Client applications | Mobile connections |
| **Sharing Method** | Address display | Address consumption | QR code display |
| **User Interaction** | None | Command-line input | QR scanning |
| **Connection Flow** | Waits for requests | Sends requests | Waits for scans |

### Code Comparison

**Direct Server (creates addresses):**
```go
// Generate shareable address
inboxAddress, err := selfAccount.InboxOpen()

// Handle incoming connections via KeyPackage
OnKeyPackage: handleKeyPackageCallback
```

**Direct Client (connects to addresses):**
```go
// Connect to a known address
recipientKey := signing.FromAddress(inboxAddress)
err = clientAccount.ConnectionNegotiate(senderKey, recipientKey, expiration)

// Handle server responses via Welcome
OnWelcome: handleConnectionResponse
```

**QR Server (creates QR codes):**
```go
// Generate QR code with discovery request
keyPackage, err := selfAccount.ConnectionNegotiateOutOfBand(...)
qrCode, err := anonymousMsg.EncodeToQR(...)

// Handle mobile connections via Welcome
OnWelcome: handleIncomingConnection
```

---

## 💡 When to Use Each Pattern

### ✅ Use Direct Server Pattern (01_direct/) For:
- **API Endpoints**: Services that others connect to
- **Backend Services**: Server applications that accept connections
- **Service Discovery**: Publishing connection endpoints
- **Microservice Mesh**: Services that accept connections from other services
- **Development Servers**: Local development with known addresses

### ✅ Use Direct Client Pattern (03_client/) For:
- **Client Applications**: Apps that connect to known servers
- **Microservices**: Services connecting to other services
- **Command-Line Tools**: CLI utilities connecting to servers
- **Automated Systems**: Scripts that connect to specific services
- **API Clients**: REST clients establishing Self connections
- **Testing Tools**: Verification that server connections work

### ✅ Use QR Server Pattern (02_qr/) For:
- **Mobile App Onboarding**: Users connecting via mobile apps
- **Desktop-to-Mobile**: Computer applications connecting to phones
- **Interactive Demos**: Live demonstrations and presentations
- **Consumer Applications**: User-facing connection establishment
- **Event Check-ins**: QR codes for conference/event connections
- **Quick Setup Scenarios**: Fast visual connection establishment

---

## 🚀 Next Steps After Connections

Once you've mastered connections, explore these advanced examples:

### Build on Your Connection Foundation
1. **📨 Messaging** → [Chat Example](../04_chat/)
   - Send messages through established connections
   - Build real-time communication systems

2. **🎫 Credentials** → [Credentials Example](../02_credentials/)
   - Issue and verify digital credentials
   - Build trust and identity systems

3. **👥 Group Communication** → [Group Chat](../04_chat/group_chat/)
   - Multi-party encrypted conversations
   - Group management and coordination

### Real-World Integration Patterns
4. **🔔 Notifications** → [Advanced Features](../08_advanced_features/notifications/)
   - Push notifications for mobile engagement
   - Event-driven communication

5. **💾 Data Storage** → [Advanced Features](../08_advanced_features/)
   - Encrypted data persistence
   - Production-ready patterns

---

## 🧪 Testing All Three Patterns

### Complete Direct Connection Testing
```bash
# Terminal 1: Start server
cd 01_direct && go run main.go
# Copy the displayed address

# Terminal 2: Connect client
cd 03_client && go run main.go <address-from-terminal-1>
# Watch end-to-end connection establishment
```

### QR Code Testing  
```bash
# Terminal 1: Start QR server
cd 02_qr && go run main.go
# Scan QR code with Self mobile app
# Watch mobile-to-server connection
```

### Cross-Pattern Testing
```bash
# Test multiple clients connecting to one server
cd 01_direct && go run main.go  # Get address
cd 03_client && go run main.go <address>  # Client 1
cd 03_client && go run main.go <address>  # Client 2

# Compare callback patterns
# Direct server uses OnKeyPackage
# Both clients use OnWelcome
# QR server uses OnWelcome
```

---

## 📚 Educational Resources

- **Code Examples**: Both subdirectories contain complete, working examples
- **Detailed READMEs**: Each approach has comprehensive documentation
- **Common Library**: Shared utilities in [`../common/`](../common/)
- **Self SDK Docs**: [Official documentation](https://docs.joinself.com)

---

**Ready to establish secure Self SDK connections?** Pick your approach and start building! 🚀

Both paths lead to the same destination: **secure, encrypted, peer-to-peer communication**. The journey just depends on your use case.
