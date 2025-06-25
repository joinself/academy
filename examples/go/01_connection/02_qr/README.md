# 🔗 Connection Example - Server Side

> **🎯 What you'll learn:** How to create a server that can securely connect with mobile apps using QR codes

This example demonstrates the **SERVER SIDE** of establishing secure connections with Self SDK. You'll create a Go application that generates QR codes for mobile apps to scan, establishing encrypted communication channels.

## 🟢 Complexity: Beginner
**Perfect first example** - Shows core connection concepts with clear, educational code.

---

## 🎓 Learning Objectives

By the end of this example, you'll master **TWO KEY CONCEPTS**:

### 🔑 **Concept 1: QR Code Generation**
- How to create scannable QR codes for mobile discovery
- The cryptographic process behind secure QR codes
- Managing QR code expiration and security

### 🔑 **Concept 2: Connection Acceptance** 
- How to handle incoming connection requests
- The `OnWelcome` callback pattern
- Establishing encrypted communication channels

**Plus these supporting concepts:**
- Client-Server Architecture in Self SDK
- Real-world integration patterns

---

## 🚀 Quick Start

```bash
cd examples/go/01_connection/02_qr
go run main.go
```

**What happens:**
1. A Self account starts up on your computer (the "server")
2. A QR code appears in your terminal 
3. You scan it with a Self mobile app
4. A secure connection establishes between them

---

## 📱 The Two Sides of Connection

### Server Side (This Example)
```
🖥️  Your Computer
├── Creates Self account
├── Generates QR code  
├── Waits for mobile apps
└── Accepts connections
```

### Mobile Side (Your Phone)
```
📱 Self Mobile App
├── Scans QR code
├── Extracts connection info
├── Sends connection request  
└── Establishes encrypted channel
```

---

## 🎬 What You'll See

When you run the example, you'll see clean, focused output:

```
🔗 Connection Example - Server Side
====================================
Setting up Self account...
✅ Account ready!

📬 Account Address: 00321b095e3dda41452ec7ff57c257fc6fde87a186e2d48f44c2585a5137914781

Generating connection QR code...

██ ▄▄▄▄▄ █▀█ █▄▀▄▀ ▄▄▄▄▄ ██
██ █   █ █▀▀ █ ▀ ▀ █   █ ██
██ █▄▄▄█ █▀█ █▄▀▄▀ █▄▄▄█ ██
██▄▄▄▄▄▄▄█▄▀ ▀▄█▄▄▄▄▄▄▄██
██ ▄▄▄▄▄ █▀█ █▄▀▄▀ ▄▄▄▄▄ ██

⏱️  Expires: 14:35:22
📱 Scan the QR code above with your Self mobile app
⏳ Waiting for connection... (Press Ctrl+C to exit)
```

**📱 To test the connection:**
1. Open the Self mobile app on your phone
2. Scan the QR code displayed in your terminal
3. Watch for connection confirmation

When someone scans your QR code:

```
🎉 Connection received from: did:self:mobile_app_address
✅ Connection established successfully!
🚀 Ready to exchange messages and credentials
```

---

## 🔍 How It Works: The Two Key Concepts

### 🔑 **CONCEPT 1: QR Code Generation**

This is how you create scannable QR codes for mobile apps:

```go
func generateConnectionQR(selfAccount *account.Account) bool {
    // Step 1: Open inbox for receiving connection requests
    inboxAddress, err := selfAccount.InboxOpen()
    
    // Step 2: Generate cryptographic key package
    keyPackage, err := selfAccount.ConnectionNegotiateOutOfBand(
        inboxAddress,
        time.Now().Add(30*time.Minute), // Expires in 30 minutes
    )
    
    // Step 3: Create discovery request
    content, err := message.NewDiscoveryRequest().
        KeyPackage(keyPackage).
        Expires(time.Now().Add(30 * time.Minute)).
        Finish()
    
    // Step 4: Encode to QR code
    anonymousMsg := event.NewAnonymousMessage(content)
    qrCode, err := anonymousMsg.EncodeToQR(event.QREncodingUnicode)
    
    return true
}
```

### 🔑 **CONCEPT 2: Connection Acceptance**

This is how you handle incoming connection requests:

```go
func handleIncomingConnection(acc *account.Account, welcome *event.Welcome) {
    // Step 1: Connection request received from mobile app
    fmt.Printf("Connection received from: %s\n", welcome.FromAddress().String())
    
    // Step 2: Accept the connection (this is the critical call!)
    _, err := acc.ConnectionAccept(welcome.ToAddress(), welcome.Welcome())
    
    // Step 3: Connection established - ready for secure communication!
}
```

**Setup:** Configure your account to use the connection handler:
```go
selfAccount := common.SetupAccount(common.AccountConfig{
    Callbacks: account.Callbacks{
        OnWelcome: handleIncomingConnection,  // Use your connection handler
    },
})
```

---

## 🎓 What Just Happened?

When you run this example, here's the **magic** that happens behind the scenes:

### During Setup:
1. **Server Creates Identity**: Your computer generates a unique Self identity (DID)
2. **Inbox Opens**: Creates a temporary "mailbox" address for receiving connections  
3. **Key Package Generated**: Creates cryptographic material for secure communication
4. **QR Code Encoded**: Packages connection info into a scannable QR code

### During Connection:
5. **Mobile Scans**: A Self mobile app extracts the connection information from the QR code
6. **Connection Request**: The mobile app sends a secure connection request to your server
7. **Automatic Accept**: Your server automatically accepts the connection
8. **Handshake**: Both sides exchange cryptographic keys automatically
9. **Channel Established**: End-to-end encrypted communication is ready!

**The beautiful part:** All the complex cryptography happens automatically. You just focus on building your application logic.

### Why the Simple Output?

The example now shows minimal output to keep things clean and focused. The detailed step-by-step progress that was shown before is documented here in the README instead. This makes the actual running experience smoother while keeping all the educational content easily accessible.

---

## 🚀 Next Steps

### Build on This Foundation

1. **📨 Add Messaging** → [Chat Example](../04_chat)
   - Send text messages through the connection
   - Build real-time communication

2. **🎫 Issue Credentials** → [Credentials Example](../02_credentials)  
   - Create and send verifiable credentials
   - Build trust relationships

3. **👥 Group Connections** → [Group Chat Example](../04_chat/group_chat)
   - Connect multiple mobile apps
   - Manage group communications

### Real-World Applications

**💡 What can you build with this?**

- **Desktop Apps**: Apps that mobile users connect to via QR codes
- **Web Services**: Websites that authenticate users through mobile apps  
- **IoT Devices**: Hardware that mobile apps can securely configure
- **Developer Tools**: Build your own Self-powered applications

---

## 🔧 Code Structure

| Function | Purpose | Educational Focus |
|----------|---------|-------------------|
| `main()` | Program flow and user guidance | Shows step-by-step progression |
| `generateConnectionQR()` | QR code creation with detailed logging | Explains each cryptographic step |
| `handleIncomingConnection()` | Connection acceptance with explanations | Shows what happens during handshake |

**Total: ~180 lines** - Educational and well-commented code.

---

## 🧪 Testing with Real Mobile Apps

1. **Download Self Mobile App**: Get it from your app store
2. **Run This Example**: Generate QR code on your computer  
3. **Scan & Connect**: Use mobile app to scan the QR code
4. **See It Work**: Watch the secure connection establish in real-time

---

## 💡 Key Concepts Learned

### Client-Server Model
- **Server** (this example): Generates QR codes, accepts connections
- **Client** (mobile app): Scans QR codes, initiates connections  
- **Result**: Secure, encrypted communication channel

### Security Features
- **End-to-End Encryption**: Only you and connected mobile apps can read messages
- **Mutual Authentication**: Both sides verify each other's identity
- **No Central Server**: Direct peer-to-peer communication
- **Time-Limited QR Codes**: Connections expire for security

### Self SDK Benefits  
- **Automatic Cryptography**: Complex security handled transparently
- **Cross-Platform**: Works with iOS, Android, and other platforms
- **Standard Protocols**: Uses W3C standards for DIDs and VCs
- **Developer Friendly**: Simple APIs for complex functionality

---

## 📚 Files in This Example

- **`main.go`**: The complete, educational connection example  
- **`go.mod`**: Go module with Self SDK dependency
- **`README.md`**: This comprehensive tutorial

Ready to build secure, connected applications? Start here! 🚀
