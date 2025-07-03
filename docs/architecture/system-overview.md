# 🌐 System Overview

> **🔧 Hands-on Learning:** After reading this, explore [Advanced Examples](../examples/advanced.md)

## What You'll Learn

Now that you understand the [cryptographic foundations](../concepts/cryptographic-foundations.md), [MLS security](../concepts/message-layer-security.md), and have built [connections](../examples/connections.md), [credentials](../examples/credentials.md), and [chat applications](../examples/chat.md), let's see how everything fits together in the complete Self SDK architecture.

**🎯 Learning Goals:**
- Understand the layered architecture of the Self SDK ecosystem
- See how cryptographic components integrate with application features
- Learn about network topology and infrastructure patterns
- Explore deployment models and integration strategies
- Connect your hands-on examples to the bigger architectural picture

## 🏗️ Complete System Architecture

The Self SDK creates a **decentralized identity and communication ecosystem** built on cryptographic foundations you've already experienced:

```mermaid
graph TB
    subgraph "Application Layer"
        A1[Chat Apps] --> A2[Identity Verification]
        A2 --> A3[Custom Business Logic]
    end
    
    subgraph "Self SDK Layer"
        S1[Account Management] --> S2[Message Handling]
        S2 --> S3[Credential Exchange]
        S3 --> S4[Connection Management]
    end
    
    subgraph "Protocol Layer"
        P1[Message Layer Security] --> P2[Verifiable Credentials]
        P2 --> P3[Connection Protocols]
    end
    
    subgraph "Cryptographic Layer"
        C1[Ed25519 Signatures] --> C2[X25519 Key Exchange]
        C2 --> C3[SHA-2 Hashing]
    end
    
    subgraph "Network Layer"
        N1[Self Network Infrastructure] --> N2[P2P Communication]
        N2 --> N3[Message Routing]
    end
    
    A1 -.-> S1
    S4 -.-> P1
    P3 -.-> C1
    C3 -.-> N1
```

**🔑 Key Insight:** Every layer you've learned builds upon the foundation beneath it. Your chat messages use MLS (Protocol Layer), which uses Ed25519 signatures (Cryptographic Layer), which routes through Self's infrastructure (Network Layer).

## 🎯 Architecture Layers Deep Dive

### 1. Application Layer - Where You Build

This is where your examples live - the applications that users interact with:

**🟢 Basic Examples:**
```bash
examples/server/00_setup/     ← Account creation apps
examples/server/01_connection/ ← Connection establishment apps  
examples/server/03_chat/      ← Messaging applications
```

**🟡 Advanced Examples:**
```bash  
examples/server/02_credentials/ ← Identity verification apps
examples/server/04_advanced/    ← Production-ready apps
```

**Your Application Integration Points:**
- **Account Creation**: Initialize cryptographic identity
- **Connection Management**: Establish secure channels  
- **Message Exchange**: Send/receive encrypted messages
- **Credential Flows**: Issue/verify identity claims
- **Business Logic**: Your custom application features

### 2. Self SDK Layer - Developer Interface

The SDK abstracts complex cryptography into simple APIs you've been using:

#### Core SDK Modules

**Account Management (`account` module):**
```go
// From your examples - this creates complete cryptographic identity
account, err := account.New(&account.Config{
    Callbacks: account.Callbacks{
        OnMessage:    handleMessage,    // Encrypted message handling
        OnWelcome:    handleConnection, // Secure connection establishment  
        OnKeyPackage: handleKeyExchange,// Cryptographic material exchange
    },
})
```

**Message Handling (`message` & `event` modules):**
```go
// From chat examples - automatic encryption/decryption
switch event.ContentTypeOf(msg) {
case message.ContentTypeChat:
    chat, _ := message.DecodeChat(msg.Content()) // Already decrypted!
case message.ContentTypeCredential:
    cred, _ := message.DecodeCredential(msg.Content())
}
```

**Credential Exchange (`credential` module):**
```go
// From credential examples - cryptographic verification
credential, err := credential.NewCredential().
    Claim("email", "user@example.com").
    Evidence(proofDocument).
    Finish()
```

#### SDK Architecture Patterns

**Event-Driven Design:**
```
Application Events → SDK Callbacks → Cryptographic Operations → Network Actions
```

**Storage Management:**
```bash
./storage/
├── account.db      # Encrypted identity keys (Ed25519 pairs)
├── connections.db  # Secure channel state (MLS groups)  
├── messages.db     # Message history (encrypted at rest)
└── credentials.db  # Verifiable credentials storage
```

### 3. Protocol Layer - Cryptographic Security

This layer implements the security protocols you've learned about:

#### Message Layer Security (MLS)
**What you experienced in [chat examples](../examples/chat.md):**
```go
// Your simple message send...
acc.MessageSend(peer, chatContent)

// ...triggers this MLS protocol flow:
// 1. Generate ephemeral key for this message
// 2. Encrypt content with group key
// 3. Sign with your Ed25519 identity key  
// 4. Add forward secrecy protections
// 5. Route through Self network
```

**MLS Group Management (behind the scenes):**
- **Group Creation**: Each connection becomes an encrypted MLS group
- **Key Rotation**: Automatic key updates for forward secrecy
- **Member Management**: Add/remove participants with security preservation
- **State Synchronization**: Keep all devices in sync with group state

#### Verifiable Credentials Protocol
**What you built in [credential examples](../examples/credentials.md):**
```
Issuer creates credential → Signs with Ed25519 → Holder receives → Verifier checks signature
```

**Protocol Components:**
- **DID-based Identity**: Each account has cryptographic identity (`did:self:...`)
- **JSON-LD Credentials**: W3C-compliant credential format
- **Proof Signatures**: Ed25519 signatures for tamper detection
- **Evidence Attachments**: Cryptographically linked supporting documents

#### Connection Protocols
**What you implemented in [connection examples](../examples/connections.md):**

**Direct Connection Flow:**
```
Client → InboxOpen() → Generate address → Share address → Server receives OnKeyPackage
```

**QR Code Connection Flow:**  
```
Server → Generate QR → Mobile scans → Mobile sends OnWelcome → Server accepts connection
```

### 4. Cryptographic Layer - Mathematical Foundation

The [cryptographic algorithms](../concepts/cryptographic-foundations.md) powering everything:

#### Ed25519 Digital Signatures
**Used throughout your examples:**
- **Account Identity**: Your DID is backed by Ed25519 keypair
- **Message Authentication**: Every message signed automatically
- **Credential Issuance**: Verifiable credentials use Ed25519 proofs
- **Connection Verification**: Prove you control a DID

#### X25519 Key Exchange
**Enables secure communication:**
- **Connection Establishment**: Generate shared secrets for new connections
- **MLS Key Derivation**: Create group encryption keys
- **Forward Secrecy**: Generate ephemeral keys for message encryption

#### SHA-2 Hashing
**Provides integrity and efficiency:**
- **Message Integrity**: Detect tampering attempts
- **Key Derivation**: Generate encryption keys from shared secrets
- **Content Addressing**: Efficiently reference large files and evidence

### 5. Network Layer - Decentralized Infrastructure  

The Self Network provides the infrastructure your applications use:

#### Self Network Topology
```
┌─────────────┐    ┌─────────────┐    ┌─────────────┐
│    Mobile   │    │   Server    │    │     IoT     │
│    Apps     │    │    Apps     │    │   Devices   │
└──────┬──────┘    └──────┬──────┘    └──────┬──────┘
       │                  │                  │
       └──────────────────┼──────────────────┘
                          │
              ┌───────────┴───────────┐
              │   Self Network        │
              │   Infrastructure      │
              │                       │
              │ ┌─────┐ ┌─────┐ ┌───┐ │
              │ │Relay│ │Relay│ │...│ │
              │ └─────┘ └─────┘ └───┘ │
              └───────────────────────┘
```

**Network Services:**
- **Message Routing**: Deliver encrypted messages between accounts
- **Identity Resolution**: Map DID addresses to network locations
- **Connection Brokering**: Help devices find each other for direct connections
- **Offline Storage**: Queue messages when recipients are offline

#### Decentralized Properties
**No Single Point of Failure:**
- Multiple relays across geographic regions
- Automatic failover if relays go offline  
- P2P direct connections when possible
- Cryptographic integrity independent of infrastructure

**Privacy Protection:**
- End-to-end encryption (relays can't read messages)
- Minimal metadata exposure (only routing information)
- Optional proxy routing for traffic analysis protection

## 🔄 System Data Flow Patterns

Let's trace how data flows through the complete system using examples you've built:

### Chat Message Flow

**From your [chat example](../examples/chat.md):**
```go
// 1. Application Layer: You call MessageSend
acc.MessageSend(peer, chatContent)

// 2. SDK Layer: Message gets processed
// - Content encoded as chat type
// - Recipient address resolved

// 3. Protocol Layer: MLS encryption applied  
// - Message signed with Ed25519
// - Encrypted with MLS group key
// - Forward secrecy headers added

// 4. Cryptographic Layer: Security operations
// - SHA-256 message integrity hash
// - Key derivation for encryption  
// - Digital signature generation

// 5. Network Layer: Message routing
// - Encrypted payload sent to Self network
// - Routed to recipient's device
// - Delivered via OnMessage callback
```

### Credential Verification Flow

**From your [credential examples](../examples/credentials.md):**
```go
// 1. Issuer creates credential
credential := credential.NewCredential().
    Issuer(issuerDID).
    Claim("email", "user@example.com").
    Finish()

// 2. Cryptographic signing
// - Ed25519 signature over credential content
// - Proof attached to credential

// 3. Secure transmission
// - Credential sent via MLS-encrypted message
// - Received by holder's OnMessage callback

// 4. Verification process
// - Verifier receives credential presentation
// - Ed25519 signature verified against issuer DID
// - Claims validated for authenticity
```

### Connection Establishment Flow

**From your [connection examples](../examples/connections.md):**
```go
// 1. Server opens inbox
inboxAddr, _ := acc.InboxOpen()

// 2. Client initiates connection
acc.ConnectionNegotiate(serverInboxAddr)

// 3. Cryptographic handshake
// - X25519 key exchange for shared secret
// - Ed25519 signatures for mutual authentication  
// - MLS group creation for future messages

// 4. Connection established
// - OnWelcome/OnKeyPackage callbacks fire
// - Secure channel ready for messaging
```

## 🏭 Deployment Architecture Patterns

### Single Service Deployment

**Example: [Chat Bot](../examples/chat.md)**
```
┌─────────────────────┐
│   Chat Bot Service  │ ← Your Go/Java application
│                     │
│ ┌─────────────────┐ │
│ │   Self SDK      │ │ ← Embedded SDK
│ │                 │ │
│ │ ┌─────────────┐ │ │
│ │ │   Storage   │ │ │ ← Local encrypted storage
│ │ └─────────────┘ │ │
│ └─────────────────┘ │
└─────────────────────┘
           │
    ┌──────┴──────┐
    │ Self Network │ ← Decentralized infrastructure
    └─────────────┘
```

**Use Cases:**
- Customer support bots
- IoT device controllers  
- Automated service endpoints
- Development and testing

### Multi-Service Architecture

**Example: [Credential Issuance System](../examples/credentials.md)**
```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│ Issuer Service  │    │ Holder Service  │    │Verifier Service │
│                 │    │                 │    │                 │
│ ┌─────────────┐ │    │ ┌─────────────┐ │    │ ┌─────────────┐ │
│ │  Self SDK   │ │    │ │  Self SDK   │ │    │ │  Self SDK   │ │
│ └─────────────┘ │    │ └─────────────┘ │    │ └─────────────┘ │
└─────────────────┘    └─────────────────┘    └─────────────────┘
         │                       │                       │
         └───────────────────────┼───────────────────────┘
                                 │
                      ┌──────────┴──────────┐
                      │    Self Network     │
                      └─────────────────────┘
```

**Use Cases:**
- Enterprise identity systems
- Educational credential platforms  
- Healthcare verification networks
- Financial service authentication

### Hybrid Cloud Deployment

**Example: [Advanced Features](../examples/advanced.md)**
```
┌─────────────────────────────────────────────────────────────┐
│                    Enterprise Cloud                        │
│                                                             │
│ ┌─────────────────┐  ┌─────────────────┐  ┌───────────────┐ │
│ │   Web Portal    │  │  API Gateway    │  │   Analytics   │ │
│ └─────────────────┘  └─────────────────┘  └───────────────┘ │
│                                                             │
│ ┌─────────────────────────────────────────────────────────┐ │
│ │              Self SDK Service Layer                     │ │
│ │                                                         │ │
│ │ ┌─────────────┐ ┌─────────────┐ ┌─────────────────────┐ │ │
│ │ │ User Mgmt   │ │ Credentials │ │    Messaging        │ │ │
│ │ └─────────────┘ └─────────────┘ └─────────────────────┘ │ │
│ └─────────────────────────────────────────────────────────┘ │
└─────────────────────────────────────────────────────────────┘
                                 │
                      ┌──────────┴──────────┐
                      │    Self Network     │
                      └─────────────────────┘
```

**Use Cases:**
- Large-scale enterprise deployments
- Multi-tenant SaaS platforms
- Government and healthcare systems  
- Global identity networks

## 🎯 Integration Patterns

### SDK Integration Models

#### Embedded SDK Pattern
**What you've been using in examples:**
```go
// SDK embedded directly in your application
func main() {
    account := common.SetupAccount(common.AccountConfig{
        Callbacks: account.Callbacks{
            OnMessage: handleMessage,
        },
    })
    // Your application logic here
}
```

**Benefits:**
- Simple deployment (single binary)
- Direct API access
- Maximum performance
- Full control over SDK lifecycle

#### Service-Oriented Pattern
**For microservice architectures:**
```
Application Service → REST API → Self SDK Service → Self Network
```

**Benefits:**
- Language-agnostic integration
- Centralized identity management
- Scalable service deployment
- Clear separation of concerns

#### Event-Driven Pattern
**For reactive architectures:**
```
Application → Message Queue → Self SDK Handler → Event Bus → Consumers
```

**Benefits:**
- Asynchronous processing
- Event sourcing capabilities
- High throughput handling
- Resilient failure handling

### Storage Integration

#### Local Storage (Examples Pattern)
```bash
./storage/              # Used in all your examples
├── account.db         # Identity and keys
├── connections.db     # Secure channels  
├── messages.db        # Message history
└── credentials.db     # Verifiable credentials
```

#### Database Integration
```go
// Custom storage backend
cfg := &account.Config{
    StoragePath: "postgresql://user:pass@host/db",
    // SDK handles encryption automatically
}
```

#### Cloud Storage Integration
```go
// Encrypted cloud storage
cfg := &account.Config{
    StoragePath: "s3://bucket/path",
    StorageKey:  customerManagedKey,
}
```

## 🚀 Performance & Scalability

### Performance Characteristics

**Message Throughput:**
- **Direct Connections**: ~10,000 messages/second per connection
- **Group Messages**: ~1,000 messages/second per group (100 members)
- **Credential Verification**: ~5,000 verifications/second per service

**Latency Profiles:**
- **Connection Establishment**: 200-500ms (includes cryptographic handshake)
- **Message Delivery**: 50-200ms (depending on geographic distance)
- **Credential Verification**: 1-5ms (local cryptographic operations)

**Resource Usage:**
- **Memory**: ~10-50MB per SDK instance (varies with connection count)
- **CPU**: Low baseline, spikes during cryptographic operations
- **Storage**: ~1-10MB per 10,000 messages (compressed and encrypted)

### Scaling Patterns

#### Horizontal Scaling
```
Load Balancer → Multiple SDK Services → Self Network
```

**Each service handles:**
- Independent account identities
- Separate connection pools
- Local message processing
- Coordinated via application logic

#### Vertical Scaling
```
More CPU/Memory → Single SDK Service → Higher throughput
```

**Optimization points:**
- Connection pooling and reuse
- Message batching for efficiency
- Cryptographic operation caching
- Storage performance tuning

#### Geographic Distribution
```
Region A ← Self Network → Region B
   ↑                        ↑  
SDK Services         SDK Services
```

**Benefits:**
- Reduced latency for regional users
- Regulatory compliance (data residency)
- Disaster recovery capabilities
- Independent regional operation

## 🔐 Security Architecture

### Defense in Depth

The Self SDK implements multiple security layers:

**1. Application Security (Your Code):**
- Input validation and sanitization
- Business logic access controls
- Audit logging and monitoring
- Secure configuration management

**2. SDK Security (Automatic):**
- End-to-end message encryption  
- Cryptographic identity verification
- Secure storage encryption
- Protocol-level protections

**3. Network Security (Self Infrastructure):**
- TLS transport encryption
- Relay authentication
- DDoS protection and filtering
- Network monitoring and anomaly detection

**4. Cryptographic Security (Mathematical):**
- Ed25519 signature security
- X25519 key exchange security  
- SHA-2 hash function security
- Forward secrecy guarantees

### Threat Model Coverage

**✅ Protected Against:**
- **Man-in-the-Middle**: Cryptographic authentication prevents impersonation
- **Eavesdropping**: End-to-end encryption protects message content
- **Message Tampering**: Digital signatures detect modifications  
- **Replay Attacks**: Nonce-based freshness prevents message replay
- **Key Compromise**: Forward secrecy limits damage from key exposure

**⚠️ Implementation Considerations:**
- **Traffic Analysis**: Message timing and size patterns may be observable
- **Endpoint Security**: Device compromise affects stored keys  
- **Implementation Bugs**: Software vulnerabilities in application code
- **Social Engineering**: Human factors outside cryptographic protection

## 🎓 What Just Happened? Complete System Understanding

### ✅ **Architectural Mastery Achieved**

You now understand how the Self SDK creates a **complete secure ecosystem**:

- **🏗️ Layered Architecture**: From cryptographic primitives to application features
- **🔐 Security Integration**: How Ed25519, MLS, and protocols work together  
- **🌐 Network Topology**: Decentralized infrastructure with no single points of failure
- **⚙️ Integration Patterns**: How to deploy and scale Self SDK applications
- **📊 Performance Profile**: Throughput, latency, and resource characteristics

### ✅ **Your Examples in Context**

Every example you built connects to the larger architecture:

| Your Example | Architecture Layer | Security Features | Network Impact |
|--------------|-------------------|-------------------|----------------|
| **[Setup](../examples/setup.md)** | SDK + Crypto Layer | Ed25519 identity generation | DID registration |
| **[Connections](../examples/connections.md)** | Protocol + Network | X25519 key exchange, MLS setup | P2P channel establishment |  
| **[Credentials](../examples/credentials.md)** | Application + Protocol | Ed25519 signatures, W3C compliance | Verifiable claim exchange |
| **[Chat](../examples/chat.md)** | Application + MLS | End-to-end encryption, forward secrecy | Real-time message routing |

### ✅ **Production Readiness Understanding**

You now know how to think about:
- **Deployment Models**: Single service → Multi-service → Cloud hybrid
- **Integration Patterns**: Embedded → Service-oriented → Event-driven
- **Scaling Strategies**: Horizontal → Vertical → Geographic distribution
- **Security Posture**: Multiple layers of cryptographic and operational protection

## 📚 Next Steps

With complete system understanding, you're ready for advanced topics:

1. **🛡️ [Security Model](security-model.md)** - Detailed threat analysis and security guarantees
2. **🔧 [Integration Patterns](integration-patterns.md)** - Production deployment and integration strategies  
3. **🚀 [Advanced Examples](../examples/advanced.md)** - Complex real-world implementation patterns
4. **📱 Mobile Integration** - Cross-platform identity and messaging

## 💡 Key Takeaways

- **Architecture = Simplicity**: Complex cryptography hidden behind simple APIs
- **Security = Automatic**: Mathematical guarantees built into every operation
- **Scale = Distributed**: No central points of failure or bottlenecks
- **Integration = Flexible**: Multiple patterns for different architectural needs
- **Future = Extensible**: Foundation ready for new protocols and capabilities

The Self SDK architecture shows how **academic cryptography becomes practical infrastructure** - you've experienced the power of decentralized identity and secure communication through hands-on examples, and now understand how it all fits together in a production-ready system.

Ready to architect the future of secure applications? 🏗️
