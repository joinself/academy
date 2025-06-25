# 🔗 Connection Client Example

> **🎯 What you'll learn:** How to connect TO a known inbox address from the client side

This example demonstrates the **CLIENT SIDE** of direct address-based connections. It connects TO any inbox address from Self SDK sources - servers, mobile apps, or other Self-enabled applications.

## 🟢 Complexity: Beginner
**Perfect for learning** - Shows the client perspective of Self SDK connections

---

## 🎯 Quick Overview

This example shows how to:
- 📞 **Initiate connections** to known inbox addresses
- 🤝 **Handle server responses** to connection requests  
- 🔐 **Establish secure channels** from the client perspective

```bash
# First: Get an inbox address from the server
cd ../01_direct && go run main.go
# Copy the displayed address

# Then: Connect to it from this client
cd ../03_client
go run main.go <inbox-address>
```

---

## 🔄 Connection Workflow

This example completes the **Direct Address Connection** pattern:

```
📊 Complete Direct Connection Flow:

🖥️  Server (01_direct/)          🖥️  Client (03_client/)
├── Creates inbox address   ──►  ├── Takes inbox address
├── Waits for connections        ├── Initiates connection
├── Handles OnKeyPackage    ◄──  ├── Sends connection request
└── Accepts connection           └── Handles OnWelcome response
                                
          🔗 Secure Connection Established! 🔗
```

### Two-Sided Process
1. **Server** creates and shares an inbox address
2. **Client** uses that address to initiate connection
3. **Server** receives KeyPackage and accepts connection
4. **Client** receives Welcome response and completes handshake

---

## 🚀 Quick Start

### 🔗 Complete Client-Server Test (Recommended!)
```bash
# Terminal 1: Start the server (creates address)
cd ../01_direct
go run main.go
# ✅ Copy the displayed inbox address

# Terminal 2: Run this client (connects to address)
cd ../03_client
go run main.go <paste-inbox-address-here>
# ✅ Watch the connection establish in both terminals!

# Example:
# go run main.go did:self:01a2b3c4d5e6f7g8h9i0j1k2l3m4n5o6p7q8r9s0t1u2v3w4x5y6z7
```

**💡 This client is specifically designed to connect TO the `01_direct` server example!**

### Prerequisites
You need an inbox address from any Self SDK source:

**📋 Sources of inbox addresses:**
- **🖥️ Direct server example**: `cd ../01_direct && go run main.go`
- **📱 Self mobile app**: Generate address in app settings
- **🔧 Any Self SDK program**: Using `InboxOpen()` method
- **🌐 Self-enabled service**: API endpoint or web application

```bash
# Example: Get address from direct server
cd ../01_direct
go run main.go
# Copy the displayed inbox address
```

### Run the Client
```bash
# Terminal 2: Connect to the server
cd ../03_client
go run main.go <paste-inbox-address-here>

# Example:
# go run main.go did:self:01a2b3c4d5e6f7g8h9i0j1k2l3m4n5o6p7q8r9s0t1u2v3w4x5y6z7
```

### Expected Output
```
🔗 Connection Client Example
============================
🎯 Target inbox address: did:self:01a2b3c4d5...
Setting up client account...
✅ Client account ready!

🔑 CONCEPT 1: Initiating Connection
==================================
📞 Connecting to inbox: did:self:01a2b3c4d5...
✅ Connection request sent successfully!

🔑 CONCEPT 2: Waiting for Server Response
=======================================
📤 Connection request sent to server
⏳ Waiting for server to accept connection...

🎉 Connection response received from server: did:self:67h8i9j0k1...
✅ Connection established successfully!
🔐 Connected to server: did:self:67h8i9j0k1...
📱 Secure group created: did:self:89l2m3n4o5...
🚀 Ready for secure communication!

🏁 Demo completed - connection established successfully!
```

---

## 🔑 Key Concepts

### 1. Connection Initiation (`ConnectionNegotiate`)
```go
// Parse the target inbox address
recipientKey := signing.FromAddress(inboxAddress)

// Get our own address for the sender field
clientInboxAddress, err := clientAccount.InboxOpen()
senderKey := signing.FromAddress(clientInboxAddress.String())

// Initiate the connection with expiration
err = clientAccount.ConnectionNegotiate(senderKey, recipientKey, expirationTime)
```

**What happens:**
- Client parses the server's inbox address
- Client creates its own inbox for receiving responses
- Client sends a connection request with cryptographic key material
- Server receives this as a `KeyPackage` event

### 2. Response Handling (`OnWelcome`)
```go
OnWelcome: handleConnectionResponse

func handleConnectionResponse(clientAccount *account.Account, welcome *event.Welcome) {
    // Accept the connection establishment from the server
    groupAddress, err := clientAccount.ConnectionAccept(welcome.ToAddress(), welcome.Welcome())
    // Connection now established!
}
```

**What happens:**
- Server accepts the connection and responds with a `Welcome` message
- Client receives the Welcome response in the OnWelcome callback
- Client calls `ConnectionAccept` to complete the handshake
- Encrypted communication channel is now established

---

## 🔍 Technical Details

### Callback Pattern
This example uses **OnWelcome** (client-side):
- Receives Welcome messages from servers that accept connections
- Complements the **OnKeyPackage** pattern used by servers
- Handles the "response" side of the connection handshake

### Connection Types
| Aspect | Client (this example) | Server ([01_direct/](../01_direct/)) |
|--------|----------------------|-----------------------------------|
| **Role** | Initiates connections | Accepts connections |
| **Input** | Inbox address (string) | None (creates own address) |
| **Callback** | `OnWelcome` | `OnKeyPackage` |
| **SDK Method** | `ConnectionNegotiate` | `ConnectionEstablish` |
| **Use Case** | Connect to known servers | Wait for incoming connections |

### Address Requirements
- **Valid Format**: Must be a proper Self DID address
- **Active Server**: The target server must be running and listening
- **Network Access**: Both client and server need network connectivity
- **Timing**: Connection requests can expire (default: 5 minutes)

---

## 💡 Common Use Cases

### ✅ Perfect For:
- **Client Applications**: Mobile apps connecting to backend services
- **Microservices**: Service A connecting to Service B
- **API Clients**: REST clients establishing Self connections to APIs
- **Development Tools**: CLI tools connecting to running services
- **Automated Systems**: Scripts that need to connect to specific services
- **Testing**: Verification that server connections work correctly
- **Mobile Integration**: Connecting to Self mobile app addresses
- **Cross-Platform**: Connecting Go services to any Self-enabled application

### 🚫 Not Ideal For:
- **Mobile App Discovery**: Use [QR codes](../02_qr/) for user-friendly mobile connections
- **Unknown Servers**: This requires knowing the exact inbox address
- **One-time Use**: Better for persistent or repeated connections

---

## 🔧 Configuration Options

### Connection Timeout
```go
// Adjust connection timeout (default: 5 minutes)
expirationTime := time.Now().Add(10 * time.Minute) // 10 minutes
err = clientAccount.ConnectionNegotiate(senderKey, recipientKey, expirationTime)
```

### Error Handling
```go
// Add retry logic for failed connections
maxRetries := 3
for i := 0; i < maxRetries; i++ {
    err = clientAccount.ConnectionNegotiate(senderKey, recipientKey, expirationTime)
    if err == nil {
        break
    }
    log.Printf("Connection attempt %d failed: %v", i+1, err)
    time.Sleep(2 * time.Second)
}
```

---

## 🤝 Perfect Partnership with Direct Server

**This client completes the connection workflow started by [01_direct](../01_direct/):**

```bash
# Complete Direct Connection Flow:
# Server (01_direct) creates address → Client (this example) connects → Secure channel established!
```

| Server Example Does ([01_direct](../01_direct/)) | This Client Does |
|------------------|-------------------|
| ✅ Creates inbox address with `InboxOpen()` | ✅ Takes inbox address as command-line argument |
| ✅ Waits for connections | ✅ Initiates connection with `ConnectionNegotiate()` |
| ✅ Uses `OnKeyPackage` callback | ✅ Uses `OnWelcome` callback |
| ✅ Calls `ConnectionEstablish()` | ✅ Calls `ConnectionAccept()` to complete |

**🎯 Together, they demonstrate the complete client-server Self SDK connection pattern!**

## 🧪 Testing

### End-to-End Test with Server Example
```bash
# Complete connection test between server and client

# Terminal 1: Start server
cd ../01_direct && go run main.go
# Server displays: "Address: did:self:01a2b3c4d5..."

# Terminal 2: Connect client  
cd ../03_client && go run main.go did:self:01a2b3c4d5...
# Watch connection establishment in both terminals
```

### Multiple Clients
```bash
# Test multiple clients connecting to one server

# Terminal 1: Server (accepts multiple connections)
cd ../01_direct && go run main.go

# Terminal 2: Client 1
cd ../03_client && go run main.go <address>

# Terminal 3: Client 2  
cd ../03_client && go run main.go <same-address>
```

### Testing with Different Address Sources
```bash
# Connect to direct server example
cd ../01_direct && go run main.go  # Get address
cd ../03_client && go run main.go <server-address>

# Connect to QR server (if you have the underlying address)
cd ../02_qr && go run main.go  # Generate QR, extract address
cd ../03_client && go run main.go <qr-server-address>

# Connect to Self mobile app
# 1. Open Self app → Settings → Generate connection address
# 2. Copy the address from mobile app
cd ../03_client && go run main.go <mobile-app-address>

# Connect to any Self-enabled service
cd ../03_client && go run main.go <service-inbox-address>
```

---

## 🚀 Next Steps

### Immediate Next Steps
1. **Try the Server Example**: Run [01_direct/](../01_direct/) to understand the server perspective
2. **Compare Approaches**: Try [02_qr/](../02_qr/) to see QR-based connections
3. **Add Messaging**: Extend this example to send messages after connection

### Advanced Applications
4. **Build Message Exchange**: Add message handling after connection establishment
5. **Connection Management**: Track and manage multiple server connections
6. **Error Recovery**: Implement robust retry and failover logic
7. **Production Patterns**: Add logging, monitoring, and health checks

### Integration Patterns
8. **REST API Client**: Use this pattern to add Self connections to HTTP clients
9. **CLI Tools**: Build command-line utilities that connect to Self services
10. **Service Mesh**: Connect microservices using Self instead of traditional networking

---

## 📚 Related Examples

### Connection Learning Path
1. **[01_direct/](../01_direct/)** - Server side (creates addresses, accepts connections)
2. **[03_client/](../03_client/)** - Client side (this example, connects to addresses)  
3. **[02_qr/](../02_qr/)** - QR codes (alternative mobile-friendly approach)

### Next Examples to Explore
- **[Chat Example](../../04_chat/)** - Send messages through established connections
- **[Credentials](../../02_credentials/)** - Exchange digital credentials
- **[Advanced Features](../../08_advanced_features/)** - Production-ready patterns

---

## 🎯 Summary

This client example demonstrates the **second half** of direct address connections. Combined with the [Direct Server Example](../01_direct/), you can establish secure, encrypted, peer-to-peer connections between any two Self SDK instances.

**Key Takeaways:**
- ✅ Clients initiate connections using `ConnectionNegotiate`
- ✅ Servers respond via `OnWelcome` callbacks  
- ✅ Both sides use `ConnectionAccept` to complete the handshake
- ✅ The result is a secure, encrypted communication channel

**Ready to build client-server applications with Self SDK?** Start here and work your way up to more complex patterns! 🚀 
