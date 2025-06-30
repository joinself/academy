# 🔗 Direct Connection Example - Address-Based Connections

> **🎯 What you'll learn:** How to establish secure connections using inbox addresses for programmatic server-to-server communication

This example demonstrates **DIRECT ADDRESS-BASED CONNECTIONS** with Self SDK. Perfect for server-to-server communication, APIs, and automated systems.

## 🟢 Complexity: Beginner

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

---

## 📊 Direct vs QR Approach

| **Direct Addresses** (This Example) | **QR Codes** ([../02_qr/](../02_qr/)) |
|-------------------------------------|--------------------------------------|
| 🔄 **Server-to-server** connections | 📱 **Mobile app** connections |
| 🤖 **API integrations** | 👤 **User-facing** applications |
| ⚡ **Automated systems** | 🖼️ **Visual discovery** |
| 🔗 **Programmatic workflows** | 📷 **Camera-based** interaction |

---

## 🔑 How It Works

### Address Creation
```go
func displayConnectionAddress(selfAccount *account.Account) bool {
    // Create a shareable "mailbox" for receiving connection requests
    inboxAddress, err := selfAccount.InboxOpen()
    if err != nil {
        return false
    }
    
    // Display the address for others to use
    fmt.Printf("Address: %s\n", inboxAddress)
    return true
}
```

### Connection Acceptance
```go
func handleKeyPackageCallback(selfAccount *account.Account, kpg *event.KeyPackage) {
    // Automatically accept incoming connection requests
    groupAddress, err := selfAccount.ConnectionEstablish(
        kpg.ToAddress(),    // Our address
        kpg.KeyPackage(),   // Their key package
    )
    
    fmt.Printf("✅ Connection established: %s\n", groupAddress)
}
```

**Setup:**
```go
selfAccount := common.SetupAccount(common.AccountConfig{
    Callbacks: account.Callbacks{
        OnKeyPackage: handleKeyPackageCallback,  // Handle incoming connections
    },
})
```

---

## 🎬 What You'll See

```
🔗 Direct Connection Example - Server Side
==========================================
📧 DIRECT CONNECTION ADDRESS:
Address: did:self:inbox:9876543210fedcba...

📧 Share the address above with other parties for direct connection
⏳ Waiting for connections... (Press Ctrl+C to exit)
```

**When someone connects:**
```
🎉 Connection request received from: did:self:connecting_party...
✅ Successfully established encrypted connection!
🚀 Connection is now ready for secure messaging!
```

---

## 🤝 Works with Client Example

**This server pairs with [03_client](../03_client/) for complete testing:**

| This Server Does | Client Example Does |
|------------------|-------------------|
| ✅ Creates inbox address with `InboxOpen()` | ✅ Takes address as input |
| ✅ Uses `OnKeyPackage` callback | ✅ Uses `OnWelcome` callback |
| ✅ Calls `ConnectionEstablish()` | ✅ Calls `ConnectionAccept()` |

---

## 💡 Use Cases

**✅ Perfect for:**
- API endpoints and backend services
- Microservice communication
- Command-line tools and automation
- Email/chat address sharing
- Development servers

---

## 🚀 Next Steps

1. **📨 Add Messaging** → [Chat Example](../../04_chat)
2. **🎫 Issue Credentials** → [Credentials Example](../../02_credentials)  
3. **🔄 Try QR Approach** → [QR Connection Example](../02_qr/)

---

**Ready to build programmatic Self connections?** This is your starting point! 🚀
