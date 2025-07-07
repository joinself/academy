# Secure Connections Concepts

> **🛠️ Hands-on Learning:** After reading this, try the [Connection Examples](../examples/connections.md)  
> **🎯 What you'll learn:** How cryptographic handshakes establish secure communication channels between Self identities

Welcome to the **heart of Self SDK security**! Here you'll discover how two Self identities establish secure, encrypted communication channels through cryptographic handshakes and key exchange protocols.

---

## What You'll Learn

By understanding secure connections, you'll master:

- **Cryptographic Handshakes**: How two parties prove identity and establish trust
- **Key Exchange**: How communication keys are generated and shared securely
- **Forward Secrecy**: Why past messages stay secure even if current keys are compromised
- **Connection Types**: When to use direct addresses vs QR codes vs client-initiated patterns
- **Security Guarantees**: What protections Self connections provide automatically

**Time Investment**: 10 minutes to understand, lifetime of secure communication  
**Immediate Result**: Deep understanding of why Self connections are cryptographically secure

---

## The Big Picture: Why Connections Matter

In traditional systems, security is an afterthought. In Self SDK, **security is the foundation**. Before any communication happens, parties must establish a **cryptographically secure channel** that guarantees:

### **Four Security Pillars**

```mermaid
graph LR
    A[🔐 CONFIDENTIALITY<br/>Only you can read<br/>your messages]
    B[🔍 AUTHENTICITY<br/>Senders are real<br/>and verified]
    C[🛡️ INTEGRITY<br/>Messages can't<br/>be tampered]
    D[🔄 FORWARD SECRECY<br/>Past messages stay<br/>secure forever]
    
    A --> B --> C --> D
    
    style A fill:#2E3138
    style B fill:#2E3138
    style C fill:#2E3138
    style D fill:#2E3138
```

### **Traditional vs Self Connections**

| Traditional (Broken) | Self SDK (Secure) |
|---------------------|-------------------|
| "HTTPS is enough" | **End-to-end encryption** |
| Company controls keys | **You control cryptographic keys** |
| Anyone can email you | **Connection required before communication** |
| No forward secrecy | **Mathematical forward secrecy guarantee** |

---

## The Cryptographic Handshake: How It Really Works

### **The Three-Act Drama**

Every Self connection follows this **cryptographic protocol**:

```mermaid
graph LR
    A[DISCOVERY<br/>Find each other]
    B[NEGOTIATION<br/>Exchange keys]
    C[ESTABLISHMENT<br/>Secure channel ready]
    
    A --> B --> C
    
    style A fill:#2E3138
    style B fill:#2E3138
    style C fill:#2E3138
```


Let's see this in **actual working code**:

### **Act 1: Discovery**
**How parties find each other**

```go
// Server creates discoverable inbox address
inboxAddress, err := selfAccount.InboxOpen()
// Creates temporary address: did:self:inbox:ABC123...
// Ready to receive connection requests
// Address can be shared via QR code or direct sharing
```

**🔑 Key Concept**: Inbox addresses are **temporary** and **secure** - they expire automatically and can't be guessed or brute-forced.

### **Act 2: Negotiation** 
**How cryptographic material is exchanged**

#### Option A: QR Code Pattern (Mobile-Friendly)

<div data-github-embed="https://github.com/joinself/academy/blob/main/examples/server/01_connection/02_qr/go/main.go#L80-L112"
     data-style="github-dark-dimmed"
     data-show-border="true"
     data-show-line-numbers="true"
     data-show-file-meta="true"
     data-show-full-path="true"
     data-show-copy="true"></div>

#### Option B: Direct Address Pattern (Programmatic)

<div data-github-embed="https://github.com/joinself/academy/blob/main/examples/server/01_connection/03_client/go/main.go#L94-L110"
     data-style="github-dark-dimmed"
     data-show-border="true"
     data-show-line-numbers="true"
     data-show-file-meta="true"
     data-show-full-path="true"
     data-show-copy="true"></div>

> **Key Concept**: Both patterns exchange **cryptographic key packages** that contain the mathematical material needed for secure communication.

### **Act 3: Establishment**
**How the secure channel is created**

<div data-github-embed="https://github.com/joinself/academy/blob/main/examples/server/01_connection/01_direct/go/main.go#L127-L132"
     data-style="github-dark-dimmed"
     data-show-border="true"
     data-show-line-numbers="true"
     data-show-file-meta="true"
     data-show-full-path="true"
     data-show-copy="true"></div>

<div data-github-embed="https://github.com/joinself/academy/blob/main/examples/server/01_connection/02_qr/go/main.go#L141-L146"
     data-style="github-dark-dimmed"
     data-show-border="true"
     data-show-line-numbers="true"
     data-show-file-meta="true"
     data-show-full-path="true"
     data-show-copy="true"></div>


> **Key Concept**: The `groupAddress` represents a **secure communication group** where all messages are end-to-end encrypted.

---

## Cryptographic Deep Dive: The Math Behind The Magic

### **Key Exchange Protocol**

Self SDK uses **Message Layer Security (MLS)** protocol for group key management:

```
Mathematical Guarantees:
├── Perfect Forward Secrecy (PFS)
├── Post-Compromise Security (PCS)  
├── Authentication of all participants
└── End-to-end encryption by default
```

### **Forward Secrecy in Action**

<div data-github-embed="https://github.com/joinself/academy/blob/main/examples/server/01_connection/02_qr/go/main.go#L87-L91"
     data-style="github-dark-dimmed"
     data-show-border="true"
     data-show-line-numbers="true"
     data-show-file-meta="true"
     data-show-full-path="true"
     data-show-copy="true"></div>


**What this means**:
- **Each connection uses different keys** - compromising one doesn't affect others
- **Keys expire automatically** - old keys become useless over time  
- **Key rotation happens automatically** - SDK handles key refresh behind the scenes
- **Past messages stay secure** - even if current keys are stolen

### **Identity Verification Process**

```go
// Every connection cryptographically verifies identity
fmt.Printf("Connection received from: %s\n", kpg.FromAddress().String())
// FromAddress is cryptographically verified
// Cannot be spoofed or impersonated
// Backed by their private key signature
```

**Security guarantee**: When you see a `FromAddress`, you can **mathematically prove** that message came from the holder of that DID's private key.

---

## Connection Patterns: Choose Your Security Model

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

---

## What Just Happened? Cryptographic Security Achieved

### **Security Revolution Experienced**

You now understand how Self SDK provides **revolutionary security**:

- **Military-grade encryption** - Every connection uses state-of-the-art cryptography
- **Identity verification** - Mathematical proof of who you're talking to
- **Forward secrecy** - Past messages stay secure even if keys are compromised  
- **Automatic protection** - All security happens behind the scenes

### **Connection Security vs Traditional**

| Traditional Connections | Self SDK Connections |
|------------------------|---------------------|
| "Login with password" | **Cryptographic handshake** |
| "Trust the server" | **Mutual authentication** |
| "Anyone can email" | **Connection required first** |
| "Hope it's secure" | **Mathematical security guarantee** |

### **Real-World Security Impact**

Your connections now provide:

- **No password vulnerabilities** - connections use cryptographic proofs
- **End-to-end encryption** - even Self can't read your messages
- **Identity authenticity** - impossible to impersonate connection parties
- **Perfect forward secrecy** - compromised keys don't compromise history

---

## Technical Architecture: How It All Fits Together

### **Complete Connection Architecture**

Through the examples, you've built this **cryptographically secure system**:

```
Self Connection Architecture
├── Identity Layer (DIDs + Cryptographic Keys)
├── Connection Layer (Handshakes + Key Exchange)  
├── Encryption Layer (MLS + Forward Secrecy)
├── Transport Layer (Network + Routing)
└── Application Layer (Messages + Credentials)
```

### **Network Security Model**

- **Discovery**: Inbox addresses enable secure discovery
- **Handshake**: MLS protocol ensures secure key exchange
- **Encryption**: All communication is end-to-end encrypted
- **Verification**: Every message is cryptographically authenticated
- **Rotation**: Keys rotate automatically for forward secrecy

---

## Next Steps: Build On Your Secure Foundation

With secure connections mastered, you're ready for advanced communication:

### **Level 1: Messaging** 🟡
- **[Chat Examples](../examples/chat.md)** - Send encrypted messages over secure connections
- **Message Types** - Text, files, structured data over encrypted channels

### **Level 2: Credentials** 🟡
- **[Credential Examples](../examples/credentials.md)** - Exchange verifiable credentials securely
- **Verification** - Cryptographic proof of claims over secure channels

### **Level 3: Advanced Security** 🟠
- **[Message Layer Security](message-layer-security.md)** - Deep dive into MLS protocol
- **[Cryptographic Foundations](cryptographic-foundations.md)** - Mathematical primitives

### **Architecture Deep Dive** 🔴
- **[Security Model](../architecture/security-model.md)** - Complete threat analysis
- **[System Overview](../architecture/system-overview.md)** - How security fits the bigger picture

---

## Common Security Questions & Troubleshooting

### **"How secure are Self connections really?"**
**Answer**: Self connections use **Message Layer Security (MLS)**, the same protocol being standardized by the IETF for next-generation secure messaging. This provides:
- Perfect Forward Secrecy (PFS)
- Post-Compromise Security (PCS)  
- Cryptographic authentication
- End-to-end encryption

### **"Why do connections expire?"**
**Answer**: Expiration provides **forward secrecy** - if keys are compromised later, past communications remain secure. It also prevents indefinite accumulation of stale connection attempts.

### **"Can connections be intercepted?"**
**Answer**: Self connections use **end-to-end encryption** with **identity verification**. Even if network traffic is intercepted, attackers cannot:
- Decrypt the messages (no keys)
- Impersonate parties (no private keys)
- Modify messages (cryptographic integrity)

### **"What happens if connection establishment fails?"**
```bash
❌ Failed to establish connection
```
**Solution**: 
- Check network connectivity
- Verify address format and validity
- Ensure both parties are online
- Check key package expiration times

---

## Additional Security Resources

### **Cryptographic Concepts**
- **[Cryptographic Foundations](cryptographic-foundations.md)** - Mathematical primitives explained
- **[Message Layer Security](message-layer-security.md)** - Deep dive into MLS protocol

### **Implementation Guides**  
- **[Connection Examples](../examples/connections.md)** - Hands-on connection patterns
- **[Security Best Practices](../architecture/security-model.md)** - Production security guidance

### **Standards & Specifications**
- **[IETF MLS Working Group](https://datatracker.ietf.org/wg/mls/)** - Message Layer Security standard
- **[RFC 9420](https://tools.ietf.org/rfc/rfc9420.txt)** - The MLS Protocol specification
- **[Double Ratchet Algorithm](https://signal.org/docs/specifications/doubleratchet/)** - Forward secrecy foundations

---

**Congratulations!** You now understand the **cryptographic foundations** that make Self SDK communications secure. Every connection you create provides military-grade security with the simplicity of a single function call.

**Ready to send secure messages?** Continue with [Message Layer Security](message-layer-security.md) to understand how messages are protected over your secure connections!
