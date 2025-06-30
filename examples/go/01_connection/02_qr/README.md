# 🔗 QR Connection Example - Mobile-Friendly Connections

> **🎯 What you'll learn:** How to create a server that securely connects with mobile apps using QR codes

This example demonstrates **QR CODE-BASED CONNECTIONS** with Self SDK. Perfect for mobile app onboarding and user-friendly visual connection establishment.

## 🟢 Complexity: Beginner

---

## 🚀 Quick Start

```bash
cd examples/go/01_connection/02_qr
go run main.go
```

**What happens:**
1. A QR code appears in your terminal 
2. Scan it with a Self mobile app
3. A secure connection establishes between them

---

## 📱 Perfect for Mobile Integration

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

## 🔑 How It Works

### QR Code Generation
```go
func generateConnectionQR(selfAccount *account.Account) bool {
    // Open inbox for receiving connection requests
    inboxAddress, err := selfAccount.InboxOpen()
    
    // Generate cryptographic key package with expiration
    keyPackage, err := selfAccount.ConnectionNegotiateOutOfBand(
        inboxAddress,
        time.Now().Add(30*time.Minute), // Expires in 30 minutes
    )
    
    // Create discovery request and encode to QR
    content, err := message.NewDiscoveryRequest().
        KeyPackage(keyPackage).
        Expires(time.Now().Add(30 * time.Minute)).
        Finish()
    
    anonymousMsg := event.NewAnonymousMessage(content)
    qrCode, err := anonymousMsg.EncodeToQR(event.QREncodingUnicode)
    
    return true
}
```

### Connection Acceptance
```go
func handleIncomingConnection(acc *account.Account, welcome *event.Welcome) {
    // Connection request received from mobile app
    fmt.Printf("Connection received from: %s\n", welcome.FromAddress().String())
    
    // Accept the connection
    _, err := acc.ConnectionAccept(welcome.ToAddress(), welcome.Welcome())
    
    // Connection established - ready for secure communication!
}
```

**Setup:**
```go
selfAccount := common.SetupAccount(common.AccountConfig{
    Callbacks: account.Callbacks{
        OnWelcome: handleIncomingConnection,  // Handle mobile connections
    },
})
```

---

## 🎬 What You'll See

```
🔗 Connection Example - Server Side
====================================
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

**When someone scans your QR code:**
```
🎉 Connection received from: did:self:mobile_app_address
✅ Connection established successfully!
🚀 Ready to exchange messages and credentials
```

---

## 💡 QR vs Direct Addresses

| **QR Codes** (This Example) | **Direct Addresses** ([../01_direct/](../01_direct/)) |
|------------------------------|---------------------------------------------------|
| 📱 **Mobile-first** connections | 🔄 **Server-to-server** connections |
| 👤 **User-friendly** scanning | 🤖 **Programmatic** workflows |
| 🖼️ **Visual discovery** | 📧 **Address sharing** (email/API) |
| 📷 **Camera-based** | ⚡ **Automated systems** |

---

## 🚀 Next Steps

1. **📨 Add Messaging** → [Chat Example](../../04_chat)
2. **🎫 Issue Credentials** → [Credentials Example](../../02_credentials)  
3. **🔄 Try Direct Approach** → [Direct Connection Example](../01_direct/)

---

## 🧪 Testing with Mobile Apps

1. **Download Self Mobile App**: Get it from your app store
2. **Run This Example**: Generate QR code on your computer  
3. **Scan & Connect**: Use mobile app to scan the QR code
4. **See It Work**: Watch the secure connection establish in real-time

---

**Ready to build mobile-friendly Self connections?** Start here! 🚀
