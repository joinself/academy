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

---

## 🎯 Choose Your Path

| I want to... | Start with | Why this path? |
|--------------|------------|----------------|
| 🔄 **Learn server-to-server** | [01_direct/](01_direct/) → [03_client/](03_client/) | Complete direct connection workflow |
| 📱 **Connect mobile apps** | [02_qr/](02_qr/) | User-friendly, visual connection |
| 🤖 **Build automated systems** | [01_direct/](01_direct/) → [03_client/](03_client/) | Programmatic, no user interaction |
| 👤 **Build user-facing apps** | [02_qr/](02_qr/) | QR scanning, mobile-first |

---

## 📊 The Three Connection Patterns

### 🔗 Direct Server (01_direct/) - "I Accept Connections"
Creates inbox addresses for others to connect to

**Key Features:**
- ✅ **Server-side**: Waits for incoming connections via `OnKeyPackage` callbacks
- ✅ **Address generation**: Creates shareable inbox addresses with `InboxOpen()`
- ✅ **Use cases**: API endpoints, backend services, microservice mesh

### 🔗 Direct Client (03_client/) - "I Connect to Addresses"  
Connects to known inbox addresses created by servers

**Key Features:**
- ✅ **Client-side**: Initiates connections via `ConnectionNegotiate()` 
- ✅ **Address consumer**: Uses addresses from servers via `OnWelcome` callbacks
- ✅ **Use cases**: Client applications, CLI tools, API clients

### 📱 QR Code Server (02_qr/) - "I Display QR Codes"
Generates QR codes for mobile apps to scan

**Key Features:**
- ✅ **Visual discovery**: Scannable QR codes via `ConnectionNegotiateOutOfBand()`
- ✅ **Mobile-first**: Designed for mobile app interactions via `OnWelcome` callbacks  
- ✅ **Use cases**: Mobile onboarding, interactive demos, consumer applications

---

## 🚀 Quick Start & Testing

### 🔗 Complete Direct Connection (Recommended!)
```bash
# Terminal 1: Start the server (creates address)
cd 01_direct && go run main.go
# Copy the displayed address

# Terminal 2: Connect as client (uses address)  
cd 03_client && go run main.go <paste-address-here>
# Watch connection establishment in both terminals
```

### 📱 QR Code Connections  
```bash
cd 02_qr && go run main.go
# Get a scannable QR code for mobile apps
```

### 🧪 Compare All Three
```bash
# Terminal 1: Direct server
cd 01_direct && go run main.go

# Terminal 2: Direct client 
cd 03_client && go run main.go <address-from-terminal-1>

# Terminal 3: QR server
cd 02_qr && go run main.go
```

---

## 🎓 Learning Path

**New to Self SDK?** All examples teach core concepts:
- Self account creation and configuration
- Connection establishment and handshake
- Callback patterns for handling events
- Encrypted communication channels

**Recommended progression:**
1. **[01_direct/](01_direct/)** - Server-side address creation
2. **[03_client/](03_client/)** - Client-side connection initiation  
3. **[02_qr/](02_qr/)** - Mobile-friendly QR approach

---

## 🚀 Next Steps After Connections

Once you've mastered connections, explore these advanced examples:

1. **📨 Messaging** → [Chat Example](../04_chat/)
2. **🎫 Credentials** → [Credentials Example](../02_credentials/)
3. **👥 Group Communication** → [Group Chat](../04_chat/group_chat/)
4. **🔔 Notifications** → [Advanced Features](../08_advanced_features/notifications/)

---

## 📚 Resources

- **Detailed READMEs**: Each subdirectory has comprehensive documentation
- **Common Library**: Shared utilities in [`../common/`](../common/)
- **Self SDK Docs**: [Official documentation](https://docs.joinself.com)

---

**Ready to establish secure Self SDK connections?** Pick your approach and start building! 🚀
