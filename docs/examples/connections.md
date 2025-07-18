# Connection Examples

> **Theory foundation:** [Secure Connections Concepts](../concepts/secure-connections.md)  
> **What you'll learn:** How to establish cryptographically secure connections between Self identities

Welcome to **cryptographic connection establishment**! These examples transform the secure connection concepts you've learned into working code that creates encrypted communication channels between Self identities.

---

## What You'll Learn

By completing these connection examples, you'll master:

- **Connection Handshakes**: How two Self identities establish trust cryptographically
- **Key Exchange**: How secure communication keys are generated and shared
- **Connection Patterns**: When to use direct addresses vs QR codes vs client-initiated flows
- **Security Guarantees**: What protections Self connections provide automatically
- **Real-world Integration**: How to build both server and client applications

**Time Investment**: 20 minutes to complete all three patterns  
**Immediate Result**: Working encrypted connections ready for messaging and credentials

---

## The Big Picture: From Theory to Practice

In [Secure Connections Concepts](../concepts/secure-connections.md), you learned **why** connections require cryptographic handshakes and **how** key exchange works. Now you'll **build** secure connections and experience the cryptography in action.

### Theory → Practice Connection

```
THEORY                                PRACTICE
Cryptographic handshake required  →  Implement connection callbacks
Key exchange provides security     →  Generate and share key packages  
Connection patterns vary by use    →  Build direct, QR, and client patterns
Forward secrecy protects history   →  See automatic key rotation
```

---

## Complete Connection Journey

Follow this **progressive learning path** to master Self connection patterns:

### **Step 1:** Master Direct Server Connections
**[Direct Connection Server](../../examples/server/01_connection/01_direct/)**

**What it demonstrates:**

- Creating shareable inbox addresses for receiving connections
- Automatic acceptance of incoming connection requests
- Server-side connection establishment workflow
- Foundation for backend services and APIs

```go
// Server creates address and waits for connections
inboxAddress, err := selfAccount.InboxOpen()
// Displays: did:self:inbox:ABC123... for others to connect to

// Automatically handles incoming connections via callback
OnKeyPackage: handleKeyPackageCallback
// Establishes secure communication channel when connections arrive
```

**Key Concept**: Direct addresses enable programmatic, server-to-server connections without user interaction.

**Time**: 3 minutes to complete  
**Success**: Server displays shareable address and waits for connections

---

### **Step 2:** Build Client Connection Initiation  
**[Client Connection Example](../../examples/server/01_connection/03_client/)**

**What it demonstrates:**

- Connecting TO known inbox addresses from client side
- Handling connection responses from servers
- Client-side connection establishment workflow
- Complete client-server connection testing

```go
// Client connects to known server address
err = clientAccount.ConnectionNegotiate(senderKey, recipientKey, expirationTime)
// Sends connection request to server

// Handles server response via callback
OnWelcome: handleConnectionResponse
// Completes secure channel establishment
```

**Key Concept**: Clients initiate connections to servers by using their published inbox addresses.

**Time**: 3 minutes to complete  
**Success**: Client connects to server, both confirm encrypted channel

---

### **Step 3:** Enable Mobile-Friendly QR Connections
**[QR Code Connections](../../examples/server/01_connection/02_qr/)**

**What it demonstrates:**

- Generating scannable QR codes for mobile discovery
- Mobile app connection workflows
- Visual connection establishment patterns
- User-friendly connection processes

```go
// Generate QR code containing connection information
keyPackage, err := selfAccount.ConnectionNegotiateOutOfBand(inboxAddress, expiration)
qrCode, err := anonymousMsg.EncodeToQR(event.QREncodingUnicode)
// Mobile apps scan QR code to extract connection details

// Handle mobile app connections via callback
OnWelcome: handleIncomingConnection
// Establishes secure channel with mobile app
```

**Key Concept**: QR codes enable user-friendly, visual connection establishment perfect for mobile applications.

**Time**: 2 minutes to complete  
**Success**: QR code displays, ready for mobile app scanning

---

## What Just Happened? Cryptographic Security in Action

### **Connection Revolution Experienced**
You've just implemented the **connection security revolution**:

- **No passwords required** - connections use cryptographic proofs instead
- **End-to-end encryption** - all communication automatically encrypted
- **Identity verification** - mathematical proof of connection party identity
- **Forward secrecy** - past messages stay secure even if keys compromised

### **Traditional vs Self Connections: You Built Both**

| Traditional Connections | Your Self Connections |
|-------------------------|----------------------|
| "Login with password" | **Cryptographic handshake** |
| "Trust the company" | **Verify cryptographically** |
| "Anyone can contact" | **Connection required first** |
| "Hope it's secure" | **Mathematical guarantee** |

### **Real-World Security Achieved**
Your connection patterns now provide:

- **Mutual authentication** - both parties verify each other cryptographically
- **Automatic encryption** - all messages encrypted without additional code
- **Tamper detection** - any message modification is immediately detected
- **Key rotation** - cryptographic keys refresh automatically for maximum security

---

## Technical Foundation Established

### **Architecture Understanding**
Through these examples, you've built this complete security system:

```
Self Connection Security System
├── Identity Layer (DIDs prove ownership)
├── Discovery Layer (inbox addresses enable finding)
├── Handshake Layer (key exchange establishes trust)
├── Encryption Layer (MLS provides forward secrecy)
└── Transport Layer (secure message delivery)
```

### **Security Model Implemented**
- **Connection Requirement**: No communication possible without prior connection
- **Key Exchange**: Automatic generation and sharing of encryption keys
- **Identity Verification**: Cryptographic proof of sender identity on every message
- **Forward Secrecy**: Compromise of current keys doesn't compromise message history

### **Connection Patterns Mastered**
- **Direct Addresses**: Server-to-server, API integrations, automated systems
- **Client Initiation**: Mobile apps, CLI tools, service connections
- **QR Codes**: Mobile onboarding, user-facing apps, visual discovery

---

## Connection Pattern Selection Guide

### **When to Use Direct Address Connections**
```go
// Perfect for:
inboxAddress, err := selfAccount.InboxOpen()
// Backend services and APIs
// Microservice communication
// Automated systems and scripts
// Development and testing
```

**Benefits**: No user interaction, scalable, integration-friendly
**Security Model**: Address sharing controls connection access

### **When to Use Client-Initiated Connections**
```go
// Perfect for:
err = clientAccount.ConnectionNegotiate(senderKey, recipientKey, expiration)
// Client applications
// Mobile apps connecting to backends
// Service discovery patterns
// Scheduled connection workflows
```

**Benefits**: Programmatic control, automated initiation, scalable
**Security Model**: Knowledge of target address enables connection

### **When to Use QR Code Connections**
```go
// Perfect for:
qrCode, err := anonymousMsg.EncodeToQR(event.QREncodingUnicode)
// Mobile app onboarding
// User-facing applications
// In-person networking
// Conference and event scenarios
```

**Benefits**: User-friendly, mobile-optimized, visual discovery
**Security Model**: Physical access to QR code controls connection

---

## Complete Connection Testing

### **End-to-End Connection Test**
Experience the full client-server connection workflow:

```bash
# Terminal 1: Start direct server
cd examples/server/01_connection/01_direct/go
go run main.go
# Copy the displayed inbox address

# Terminal 2: Connect with client
cd ../../../03_client/go  
go run main.go <paste-address-here>
# Watch secure connection establish in both terminals
```

### **Multi-Pattern Testing**
Compare all three connection approaches:

```bash
# Direct server (programmatic)
cd 01_direct/go && go run main.go

# QR server (mobile-friendly)  
cd ../02_qr/go && go run main.go

# Client connection (to either server)
cd ../03_client/go && go run main.go <server-address>
```

---

## Next Steps: Build On Your Secure Foundation

With secure connections established, you're ready for advanced communication:

### **Level 1: Encrypted Messaging**
- **[Chat Examples](chat.md)** - Send encrypted messages over secure connections
- **Message Types** - Text, files, structured data transmission

### **Level 2: Verifiable Credentials**  
- **[Credential Examples](credentials.md)** - Exchange cryptographic proofs securely
- **Trust Networks** - Build webs of verifiable claims

### **Level 3: Advanced Features**
- **[Advanced Examples](advanced.md)** - Production patterns and optimization
- **Group Communication** - Multi-party secure messaging

### **Architecture Deep Dive**
- **[Message Layer Security](../concepts/message-layer-security.md)** - Deep dive into MLS protocol

---

## Production Deployment

**Ready for production connections?** Check our comprehensive **[Production Deployment Guide](production.md)** for everything you need to deploy secure connection systems:

- **[Moving to Production](production.md#moving-to-production)** - Connection security patterns, key management, and monitoring strategies
- **[Security Hardening](production.md#security-hardening)** - Connection filtering, rate limiting, access control, and audit trails
- **[Performance Optimization](production.md#performance-optimization)** - Connection pooling, load balancing, and resource management
- **[Monitoring & Observability](production.md#monitoring--observability)** - Connection monitoring, security event logging, and performance metrics
- **[Scalability Patterns](production.md#scalability-patterns)** - Distributed connection handling and high-availability architectures

The production guide includes connection-specific security patterns, performance optimization techniques, and enterprise deployment strategies.

---

## Success Checklist

Confirm you've mastered secure connections:

**Connection Establishment**

- [ ] Can create inbox addresses for receiving connections
- [ ] Understand key package generation and exchange
- [ ] Can implement all three connection patterns
- [ ] Know how to handle connection callbacks

**Security Understanding**  

- [ ] Know why connections are required before communication
- [ ] Understand cryptographic handshake process
- [ ] Can explain forward secrecy benefits
- [ ] Recognize different security models

**Pattern Selection**

- [ ] Can choose appropriate connection pattern for use case
- [ ] Know when to use direct vs QR vs client-initiated patterns
- [ ] Understand scalability and usability tradeoffs
- [ ] Can implement production security best practices

**Technical Implementation**

- [ ] Master connection callback patterns (OnKeyPackage, OnWelcome)
- [ ] Understand connection lifecycle and error handling
- [ ] Can build complete client-server applications
- [ ] Ready for message exchange over secure connections

---

## Need Help?

**Having connection issues?** Check our comprehensive **[Troubleshooting Guide](troubleshooting.md)** for solutions to connection problems, including:

- **[Connection Issues](troubleshooting.md#connection-issues)** - Connection establishment, handshake failures, and QR code problems
- **[Network Issues](troubleshooting.md#network--connectivity-issues)** - Connectivity and Self network problems
- **[Setup Issues](troubleshooting.md#setup--account-issues)** - Account initialization and configuration problems

The troubleshooting guide includes detailed solutions, common causes, and debugging tips for all Self SDK examples.

---

## Resources & Next Steps

**Need more connection resources?** Check our comprehensive **[Resources & Community Guide](resources.md)** for everything you need to build secure connections:

- **[Cryptographic Foundations](resources.md#related-concepts)** - Mathematical basis of connection security and MLS protocol details
- **[Developer Tools](resources.md#developer-tools)** - Connection testing utilities, debug logging, and network monitoring
- **[Community Support](resources.md#community-support)** - Get help with connection implementation and report connection bugs

The resources guide includes complete documentation for connection patterns, security standards, and community guidelines.

---

**Congratulations!** You've mastered cryptographic connection establishment and experienced state-of-the-art security firsthand. Your Self connections now provide military-grade encryption, identity verification, and forward secrecy automatically.

**Ready to send encrypted messages?** Continue with [Chat Examples](chat.md) to build secure messaging over your encrypted connections!
