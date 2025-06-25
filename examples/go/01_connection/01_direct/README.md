# 🔗 Direct Connection Example - Address-Based Connections

> **🎯 What you'll learn:** How to establish secure connections using inbox addresses instead of QR codes

This example demonstrates **DIRECT ADDRESS-BASED CONNECTIONS** with Self SDK. Unlike the QR code approach in the parent directory, this method uses shareable inbox addresses for programmatic connections - perfect for server-to-server communication, APIs, and automated systems.

## 🟢 Complexity: Beginner
**Alternative first example** - Shows core connection concepts with a different, more programmatic approach.

---

## 🎓 Learning Objectives

By the end of this example, you'll master **TWO KEY CONCEPTS**:

### 🔑 **Concept 1: Direct Address Creation**
- How to generate shareable inbox addresses for connections
- When to use addresses vs QR codes
- Address-based connection workflows

### 🔑 **Concept 2: Automatic Connection Acceptance** 
- How to handle incoming connection requests via KeyPackage callbacks
- The `OnKeyPackage` callback pattern
- Establishing encrypted communication channels programmatically

**Plus these supporting concepts:**
- Differences between QR and address-based connections
- Server-to-server communication patterns
- API integration approaches

---

## 🚀 Quick Start

### 🔗 Complete Server-Client Test (Recommended!)
```bash
# Terminal 1: Start this server
cd examples/go/01_connection/01_direct
go run main.go
# ✅ Copy the displayed inbox address

# Terminal 2: Connect from the client example
cd ../03_client
go run main.go <paste-inbox-address-here>
# ✅ Watch the connection establish in both terminals!
```

**💡 The `03_client` example is specifically designed to connect TO this server!**

### Run Server Only
```bash
cd examples/go/01_connection/01_direct
go run main.go
```

**What happens:**
1. A Self account starts up on your computer (the "server")
2. An inbox address is generated and displayed
3. Other parties can connect using this address directly
4. All incoming connections are accepted automatically

---

## 📊 Direct Connections vs QR Codes

### This Example: Direct Address Approach
```
🖥️  Your Server
├── Creates Self account
├── Generates inbox address
├── Shares address programmatically  
└── Accepts connections automatically
```

### Parent Directory: QR Code Approach  
```
🖥️  Your Server
├── Creates Self account
├── Generates QR code
├── Displays QR for mobile scanning
└── Accepts connections from mobile apps
```

### When to Use Each Approach

| Use Direct Addresses When... | Use QR Codes When... |
|------------------------------|----------------------|
| 🔄 **Server-to-server** connections | 📱 **Mobile app** connections |
| 🤖 **API integrations** | 👤 **User-facing** applications |
| 📧 **Email/chat sharing** | 🎯 **Quick mobile** onboarding |
| ⚡ **Automated systems** | 🖼️ **Visual discovery** needed |
| 🔗 **Programmatic workflows** | 📷 **Camera-based** interaction |

---

## 🎬 What You'll See

When you run the direct connection example:

```
🔗 Direct Connection Example - Server Side
==========================================
Setting up Self account...
✅ Account ready!

📬 Server Account Information:
   🆔 Account DID: did:self:1234567890abcdef...

🔑 CONCEPT 1: Creating Direct Connection Address
===============================================

📧 DIRECT CONNECTION ADDRESS:
=============================
Address: did:self:inbox:9876543210fedcba...

💡 Other parties can connect using this address directly
   (no QR code scanning required)
📋 How others connect:
   1. Copy the address above
   2. Use it in their Self SDK connection method
   3. Send a connection request to this address

🔑 CONCEPT 2: Accepting All Incoming Connections
==============================================
📧 Share the address above with other parties for direct connection
⏳ Waiting for connections... (Press Ctrl+C to exit)
🤖 All incoming connections will be accepted automatically
```

**When someone connects using your address:**

```
🎉 Connection request received from: did:self:connecting_party...
✅ Successfully established encrypted connection!
📱 Connected to: did:self:connecting_party...
🔐 Secure group created: did:self:group:secure_channel...
🚀 Connection is now ready for secure messaging!

🏁 Demo completed - connection established successfully!
```

---

## 🔍 How It Works: The Two Key Concepts

### 🔑 **CONCEPT 1: Direct Address Creation**

This is how you create shareable addresses for direct connections:

```go
func displayConnectionAddress(selfAccount *account.Account) bool {
    // Step 1: Open inbox for receiving connection requests
    // This creates a temporary "mailbox" where others can send connection requests
    inboxAddress, err := selfAccount.InboxOpen()
    if err != nil {
        log.Printf("❌ Failed to open inbox: %v", err)
        return false
    }

    // Step 2: Extract the inbox address (returned directly from InboxOpen)
    // The address is a Self DID that others can use to connect directly

    // Step 3: Display the inbox address for direct connection
    fmt.Printf("Address: %s\n", inboxAddress)
    
    return true
}
```

### 🔑 **CONCEPT 2: Automatic Connection Acceptance**

This is how you handle incoming connection requests:

```go
func handleKeyPackageCallback(selfAccount *account.Account, kpg *event.KeyPackage) {
    // Step 1: Connection request received
    fmt.Printf("🎉 Connection request received from: %s\n", kpg.FromAddress().String())
    
    // Step 2: Establish encrypted connection using their key package
    groupAddress, err := selfAccount.ConnectionEstablish(
        kpg.ToAddress(),    // Our address (where they sent the request)
        kpg.KeyPackage(),   // Their cryptographic key package for encryption
    )

    // Step 3: Connection established - ready for secure communication!
    fmt.Printf("✅ Successfully established encrypted connection!\n")
}
```

**Setup:** Configure your account to use the connection handler:
```go
selfAccount := common.SetupAccount(common.AccountConfig{
    Callbacks: account.Callbacks{
        OnKeyPackage: handleKeyPackageCallback,  // Use your connection handler
    },
})
```

---

## 🎓 What Just Happened?

When you run this example, here's the **magic** that happens behind the scenes:

### During Setup:
1. **Server Creates Identity**: Your computer generates a unique Self identity (DID)
2. **Inbox Opens**: Creates a shareable "mailbox" address for receiving connections  
3. **Address Generated**: Produces a DID address that others can use programmatically

### During Connection:
4. **Address Sharing**: Other parties get your inbox address (via API, email, etc.)
5. **Connection Request**: They send a secure connection request to your address
6. **Automatic Accept**: Your server automatically accepts the connection via KeyPackage callback
7. **Handshake**: Both sides exchange cryptographic keys automatically
8. **Channel Established**: End-to-end encrypted communication is ready!

**The beautiful part:** No QR codes needed - perfect for automated, programmatic connections.

---

## 🔄 Comparing with QR Code Approach

### Code Comparison

| **Direct Address** (This Example) | **QR Code** (Parent Directory) |
|-----------------------------------|--------------------------------|
| `inboxAddress, err := selfAccount.InboxOpen()` | `keyPackage, err := selfAccount.ConnectionNegotiateOutOfBand(...)` |
| `OnKeyPackage: handleKeyPackageCallback` | `OnWelcome: handleIncomingConnection` |
| Share address via API/email | Display QR code for scanning |
| Perfect for automation | Perfect for mobile users |

### Workflow Comparison

**Direct Address Workflow:**
1. 🔧 Server generates inbox address
2. 📤 Address shared programmatically (API, email, etc.)
3. 🔗 Other party connects using address
4. ✅ Connection established via KeyPackage

**QR Code Workflow:**
1. 🔧 Server generates QR code with discovery request
2. 📱 QR code displayed for mobile scanning
3. 📷 Mobile app scans QR code
4. ✅ Connection established via Welcome message

---

## 🚀 Next Steps

### Build on This Foundation

1. **📨 Add Messaging** → [Chat Example](../../04_chat)
   - Send text messages through the established connection
   - Build real-time communication on top of direct connections

2. **🎫 Issue Credentials** → [Credentials Example](../../02_credentials)  
   - Create and send verifiable credentials via direct connections
   - Build trust relationships programmatically

3. **🔄 Try QR Approach** → [QR Connection Example](../)
   - Compare with the mobile-friendly QR code approach
   - Understand when to use each method

### 🤝 Works Perfectly With Client Example

**This server pairs with [03_client](../03_client/) for complete testing:**

```bash
# Server (this example) creates address → Client connects to address → Secure communication!
```

| This Server Does | Client Example Does |
|------------------|-------------------|
| ✅ Creates inbox address | ✅ Takes inbox address as input |
| ✅ Waits for connections | ✅ Initiates connection request |
| ✅ Uses OnKeyPackage callback | ✅ Uses OnWelcome callback |
| ✅ Accepts connection with ConnectionEstablish | ✅ Completes connection with ConnectionAccept |

**🎯 Perfect for learning the complete client-server connection flow!**

### Real-World Applications

**💡 What can you build with direct connections?**

- **API Services**: Backend services that other servers connect to
- **Microservices**: Self-enabled communication between services
- **Integration Platforms**: Connect existing systems via Self addresses
- **Developer Tools**: Command-line tools that establish Self connections
- **Automated Workflows**: Scripts that connect and exchange data automatically

---

## 🔧 Code Structure

| Function | Purpose | Educational Focus |
|----------|---------|-------------------|
| `main()` | Program flow and concept demonstration | Shows two key concepts clearly |
| `displayConnectionAddress()` | Address generation with detailed steps | Explains inbox creation and sharing |
| `handleKeyPackageCallback()` | Connection acceptance with explanations | Shows KeyPackage-based connection flow |

**Total: ~140 lines** - Clean, educational, and well-documented code.

---

## 🆚 When to Use This vs QR Codes

### ✅ Use Direct Addresses For:
- **Server-to-Server Communication**: APIs and microservices
- **Command-Line Tools**: Developer utilities and scripts
- **Email/Chat Integration**: Share addresses in messages
- **Automated Systems**: Programmatic connection establishment
- **Backend Services**: Non-interactive connection scenarios

### ✅ Use QR Codes For:
- **Mobile App Onboarding**: User-friendly connection establishment
- **Desktop-to-Mobile**: Computer apps connecting to phones
- **Interactive Demos**: Visual connection establishment
- **User-Facing Applications**: Consumer-oriented connection flows
- **Quick Setup**: Fast visual connection establishment

---

## 💡 Key Concepts Learned

### Direct Address Model
- **Server** (this example): Generates shareable addresses, accepts connections programmatically
- **Client** (other systems): Use addresses to connect directly via APIs
- **Result**: Secure, automated communication channels

### Callback Differences
- **OnKeyPackage**: Used for direct address connections (this example)
- **OnWelcome**: Used for QR code connections (parent directory)
- **Both**: Establish the same secure, encrypted communication

### Integration Benefits
- **No User Interaction**: Perfect for automated systems
- **API-Friendly**: Addresses can be shared via REST APIs
- **Scalable**: Handle multiple connections programmatically
- **Server-Oriented**: Designed for backend service integration

---

## 📚 Files in This Example

- **`main.go`**: The complete direct connection example with educational comments
- **`go.mod`**: Go module with Self SDK dependency  
- **`README.md`**: This comprehensive tutorial (you're reading it!)

## 🧪 Testing Direct Connections

1. **Run This Example**: Get your inbox address
2. **Use Another Self SDK Instance**: Connect using the displayed address
3. **Try Both Approaches**: Compare with the QR code example in parent directory
4. **Build Integration**: Use addresses in your own applications

---

**Ready to build programmatic, server-oriented Self connections?** This is your starting point! 🚀

For mobile-friendly QR connections, check out the [parent directory example](../README.md). 
