# Connection Patterns: Choose Your Security Model

Self SDK provides **three connection patterns** for different use cases:

### **Pattern 1: Direct Address Connections** 
**[Try it: Direct Connection Example](https://github.com/joinself/academy/tree/main/examples/server/01_connection/01_direct/)**

```go
// Perfect for server-to-server communication
inboxAddress, err := selfAccount.InboxOpen()
// Share address via API, email, configuration, etc.
// Other services connect using address directly
```

**Use when:**

- Building APIs and backend services
- Microservice communication  
- Automated systems and scripts
- Development and testing environments

**Security model:** Address sharing controls who can initiate connections

### 📱 **Pattern 2: QR Code Connections**
**[Try it: QR Connection Example](https://github.com/joinself/academy/tree/main/examples/server/01_connection/02_qr/)**

```go
// Mobile-friendly visual discovery
qrCode, err := anonymousMsg.EncodeToQR(event.QREncodingUnicode)
// Users scan QR code with mobile apps
// Automatic connection establishment
```

**Use when:**

- Mobile app onboarding
- User-facing applications
- In-person connection establishment
- Conference/event networking

**Security model:** Physical access to QR code controls who can connect

### **Pattern 3: Client-Initiated Connections**
**[Try it: Client Connection Example](https://github.com/joinself/academy/tree/main/examples/server/01_connection/03_client/)**

```go
// Programmatic connection initiation
err = clientAccount.ConnectionNegotiate(senderKey, recipientKey, expirationTime)
// Client actively connects to known service
// Server handles via OnKeyPackage callback
```

**Use when:**

- Client applications connecting to services
- Mobile apps connecting to backend APIs
- Scheduled/automated connection workflows
- Service discovery patterns

**Security model:** Knowledge of target address controls connection initiation 
