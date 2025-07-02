# 🔗 Connection Client Example

> **🎯 What you'll learn:** How to connect TO a known inbox address from the client side

This example demonstrates the **CLIENT SIDE** of direct address-based connections. It connects TO any inbox address from Self SDK sources - servers, mobile apps, or other Self-enabled applications.

## 🟢 Complexity: Beginner

---

## 🚀 Quick Start

### 🔗 Complete Client-Server Test (Recommended!)
```bash
# Terminal 1: Start the server (creates address)
cd ../01_direct && go run main.go
# Copy the displayed address

# Terminal 2: Run this client (connects to address)
cd ../03_client
go run main.go <paste-inbox-address-here>
# ✅ Watch the connection establish in both terminals!
```

**💡 This client is specifically designed to connect TO the `01_direct` server example!**

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

### 2. Response Handling (`OnWelcome`)
```go
OnWelcome: handleConnectionResponse

func handleConnectionResponse(clientAccount *account.Account, welcome *event.Welcome) {
    // Accept the connection establishment from the server
    groupAddress, err := clientAccount.ConnectionAccept(welcome.ToAddress(), welcome.Welcome())
    // Connection now established!
}
```

---

## 🎬 Expected Output

```
🔗 Connection Client Example
============================
🎯 Target inbox address: did:self:01a2b3c4d5...
Setting up client account...
✅ Client account ready!

📞 Connecting to inbox: did:self:01a2b3c4d5...
✅ Connection request sent successfully!

⏳ Waiting for server to accept connection...

🎉 Connection response received from server: did:self:67h8i9j0k1...
✅ Connection established successfully!
🚀 Ready for secure communication!
```

---

## 📊 Client vs Server Patterns

| Aspect | Client (this example) | Server ([01_direct/](../01_direct/)) |
|--------|----------------------|-----------------------------------|
| **Role** | Initiates connections | Accepts connections |
| **Input** | Inbox address (string) | None (creates own address) |
| **Callback** | `OnWelcome` | `OnKeyPackage` |
| **SDK Method** | `ConnectionNegotiate` | `ConnectionEstablish` |
| **Use Case** | Connect to known servers | Wait for incoming connections |

---

## 💡 Use Cases

**✅ Perfect For:**
- Client applications connecting to backend services
- Microservices (Service A connecting to Service B)
- CLI tools connecting to running services
- Testing server connections
- Mobile integration with Self addresses

---

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

---

## 🚀 Next Steps

1. **📨 Add Messaging** → [Chat Example](../../03_chat)
2. **🎫 Issue Credentials** → [Credentials Example](../../02_credentials)  
3. **🔄 Try Server Side** → [Direct Server Example](../01_direct/)

---

**Ready to build client-server applications with Self SDK?** Start here! 🚀
