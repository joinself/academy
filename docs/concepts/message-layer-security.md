# Message Layer Security

> **Hands-on Learning:** After reading this, try the [Chat Examples](../examples/chat.md)

## What You'll Learn

Message Layer Security (MLS) is the protocol that powers Self SDK's secure communication. It provides end-to-end encryption, forward secrecy, and post-compromise security for both individual and group conversations. This page explains how MLS works and why it's essential for secure messaging.

**Learning Goals:**
- Understand how MLS provides end-to-end encryption
- Learn about forward secrecy and post-compromise security
- Explore group messaging security challenges and solutions
- See how MLS callbacks handle protocol events
- Discover the security guarantees MLS provides

---

## What is Message Layer Security?

**Message Layer Security (MLS)** is a cryptographic protocol designed for secure group messaging. Unlike traditional point-to-point encryption, MLS enables efficient, scalable security for groups of any size while maintaining strong security properties.

### Core Security Properties

**End-to-End Encryption:**
- Messages encrypted before leaving sender's device
- Only intended recipients can decrypt messages
- Network infrastructure cannot read message content

**Forward Secrecy:**
- Compromising current keys doesn't affect past messages
- Keys automatically rotated and old keys deleted
- Past communications remain secure even after key theft

**Post-Compromise Security:**
- System can heal from key compromise
- New secure communication established automatically
- Group continues functioning even after partial compromise

**Efficient Group Operations:**
- Adding/removing members doesn't require N² operations
- Scalable key management for large groups
- Minimal bandwidth overhead for group changes

### Real-World Analogy

Think of MLS like a **secure conference room** that evolves:

```
Traditional Encryption = Simple Door Lock
├── One key fits all situations
├── If key stolen, everything compromised
└── Same security level forever

MLS = Smart Conference Room
├── Keys change automatically every meeting
├── New people get fresh room keys  
├── Old keys stop working immediately
├── Room "heals" if security compromised
└── All handled automatically by your device
```

---

## MLS in the Self SDK

The Self SDK implements MLS transparently - you get enterprise-grade security without managing the complexity.

### Automatic MLS Integration

**Every Self SDK Account Uses MLS:**
```go
// From examples/server/04_advanced_features/01_core_features/go/main.go
cfg := &account.Config{
    Callbacks: account.Callbacks{
        OnConnect:         d.onConnect,         // Account connected to network
        OnDisconnect:      d.onDisconnect,      // Account disconnected from network
        OnAcknowledgement: d.onAcknowledgement, // Message delivery confirmed
        OnError:           d.onError,           // Error occurred during operation
        OnMessage:         d.onMessage,         // Encrypted messages received and decrypted
        OnCommit:          d.onCommit,          // Group state changes confirmed
        OnKeyPackage:      d.onKeyPackage,      // Cryptographic material for establishing connections
        OnProposal:        d.onProposal,        // Proposed group changes
        OnWelcome:         d.onWelcomeMessage,  // New group member joined (including 1:1 "groups")
    },
}
```

**What Each Callback Handles:**
- **OnConnect**: Account successfully connected to Self network infrastructure
- **OnDisconnect**: Account disconnected (with optional error details)
- **OnAcknowledgement**: Message delivery confirmed by recipient
- **OnError**: Error occurred during account operation (with reference and error details)
- **OnMessage**: Encrypted messages received and automatically decrypted by MLS
- **OnCommit**: Group membership or key changes confirmed and applied
- **OnKeyPackage**: Cryptographic material exchanged during connection establishment
- **OnProposal**: Pending changes to group security or membership proposed
- **OnWelcome**: Someone joined a secure group (including 1:1 "groups")

### Message Flow Example

```go
// From https://github.com/joinself/academy/blob/main/examples/server/03_chat/01_basic/go/main.go
func handleMessage(acc *account.Account, msg *event.Message) {
    // MLS automatically decrypted this message before delivery
    switch event.ContentTypeOf(msg) {
    case message.ContentTypeChat:
        handleChatMessage(acc, msg, timestamp)
    }
}

func sendResponse(acc *account.Account, toAddress *signing.PublicKey, responseText string) {
    chatContent, err := message.NewChat().
        Message(responseText).
        Finish()
    
    // MLS automatically encrypts before sending
    err = acc.MessageSend(toAddress, chatContent)
}
```

**What Just Happened:**
1. **Incoming**: MLS decrypted message using current group keys
2. **Processing**: Your app handles the plaintext message content
3. **Outgoing**: MLS encrypts response with fresh cryptographic material
4. **Delivery**: Recipient's MLS automatically decrypts upon receipt

---

## Key Management Deep Dive

MLS's security comes from sophisticated key management that happens automatically.

### Key Lifecycle

**Key Generation:**
```go
// From https://github.com/joinself/academy/blob/main/examples/server/04_advanced_features/01_core_features/go/main.go
// Generate key package for connection (valid for 15 minutes)
keyPackage, err := d.account.ConnectionNegotiateOutOfBand(
    inboxAddress,
    time.Now().Add(15*time.Minute), // Keys expire automatically!
)
```

**Key Rotation:**
- **Automatic**: Keys rotate on schedule and group changes
- **Triggered**: Member joins/leaves, security events
- **Transparent**: Applications never see raw cryptographic material

**Key Deletion:**
- **Forward Secrecy**: Old keys deleted after rotation
- **Memory Safety**: Keys zeroed out when no longer needed
- **Storage Cleanup**: Historical keys removed from device storage

### Key Package Structure

```
Key Package (exchanged during connection setup):
├── Identity Key (Ed25519)
│   └── Long-term identity verification
├── Encryption Key (X25519)  
│   └── Message encryption material
├── Expiration Time
│   └── Automatic key material expiry
├── Signature
│   └── Proof of authenticity
└── Supported Features
    └── Protocol capabilities
```

**Real Example from Self SDK:**
```go
// From https://github.com/joinself/academy/blob/main/examples/server/04_advanced_features/01_core_features/go/main.go
func (d *AdvancedDemo) onKeyPackage(acc *account.Account, keyPackageMsg *event.KeyPackage) {
    fmt.Printf("🔑 Key package received from %s\n", keyPackageMsg.FromAddress().String())
    // SDK automatically processes the cryptographic material
}
```

---

## 👥 Group Messaging Security

MLS shines in group scenarios where traditional encryption becomes complex and inefficient.

### The Group Challenge

**Without MLS (Traditional Approach):**
```
Group of 4 people = 6 pairwise connections
├── A ↔ B, A ↔ C, A ↔ D
├── B ↔ C, B ↔ D  
└── C ↔ D

Problems:
❌ Message sent 3 times (inefficient)
❌ Keys managed separately (complex)
❌ Adding member requires 4 new connections
❌ No forward secrecy across group
```

**With MLS (Efficient Approach):**
```
Group of 4 people = 1 secure group
├── Shared group encryption key
├── Individual identity verification
├── Efficient key derivation tree
└── Automatic member management

Benefits:
✅ Message sent once (efficient)
✅ One key derivation for all (simple)
✅ Adding member requires one operation
✅ Forward secrecy for entire group
```

### Group Operations in Self SDK

**Creating Secure Groups:**
```go
// Groups are created automatically when connections are established
// Every connection in Self SDK is actually a "group" (even 1:1 conversations)

groupAddress, err := selfAccount.ConnectionEstablish(
    kpg.ToAddress(),  // Group coordination address
    kpg.KeyPackage(), // Cryptographic material
)
// Result: groupAddress represents a secure MLS group
```

**Group State Management:**
- **Membership**: Who can send/receive messages
- **Keys**: Current encryption material for the group
- **History**: Secure audit trail of group changes
- **Epochs**: Security periods with distinct keys

### Real Group Example

```go
// From https://github.com/joinself/academy/blob/main/examples/server/04_advanced_features/01_core_features/go/main.go
func (d *AdvancedDemo) onWelcomeMessage(acc *account.Account, welcomeMsg *event.Welcome) {
    // New member joined our secure group
    peerAddress := welcomeMsg.FromAddress()
    d.connections[peerAddress.String()] = true

    // Accept them into the group
    _, err := acc.ConnectionAccept(welcomeMsg.ToAddress(), welcomeMsg.Welcome())
    
    // Send welcome message (automatically encrypted for group)
    d.sendAdvancedWelcomeMessage(peerAddress)
}
```

**What Just Happened:**
1. **Welcome Event**: Someone wants to join secure communication
2. **Group Update**: MLS adds them to the secure group  
3. **Key Derivation**: New group keys derived for all members
4. **Message Encryption**: Welcome message encrypted with new group key
5. **Forward Secrecy**: Previous keys deleted, old messages stay secure

---

## Security Guarantees Explained

### Forward Secrecy in Practice

**Timeline-Based Security:**
```
Time: 10:00 AM - Message 1 sent (Key A)
Time: 10:15 AM - Key rotation (Key A deleted, Key B active)  
Time: 10:30 AM - Message 2 sent (Key B)
Time: 10:45 AM - Attacker compromises device

Result:
Message 2 can be decrypted (Key B compromised)
Message 1 CANNOT be decrypted (Key A was deleted)
```

**Real Implementation:**
```go
// From https://github.com/joinself/academy/blob/main/examples/server/04_advanced_features/01_core_features/go/main.go
keyPackage, err := d.account.ConnectionNegotiateOutOfBand(
    inboxAddress,
    time.Now().Add(15*time.Minute), // Keys automatically expire
)
```

### Post-Compromise Security

**Healing Process:**
```
1. Compromise Detected/Suspected
   └── Device compromised or key leaked
   
2. Healing Initiated  
   └── New key material generated
   
3. Key Distribution
   └── Fresh keys distributed to group
   
4. Security Restored
   └── Future communication secure again
```

**Automatic Healing:**
- MLS continuously rotates keys
- Even undetected compromises get "healed"
- No manual intervention required
- Security improves over time

### Metadata Protection

**What MLS Protects:**
- ✅ **Message Content**: Encrypted end-to-end
- ✅ **Message Authentication**: Sender verification
- ✅ **Message Integrity**: Tamper detection
- ✅ **Key History**: No access to past keys

**What MLS Doesn't Protect:**
- ❌ **Traffic Analysis**: Message timing and size
- ❌ **Participant Identity**: Who's in the conversation
- ❌ **Communication Metadata**: When messages sent

**Additional Protection in Self SDK:**
```go
// Self SDK adds layers beyond MLS:
// - Decentralized infrastructure (harder to monitor)
// - DID-based addressing (privacy-friendly identifiers)  
// - Optional proxy routing (traffic analysis protection)
```

---

## 🌐 Protocol Architecture

### MLS Protocol Stack

```
┌─────────────────────────────────────┐
│           Application Layer         │ ← Your chat, credentials, etc.
├─────────────────────────────────────┤
│      Self SDK Message Layer         │ ← Message routing, addressing
├─────────────────────────────────────┤  
│     Message Layer Security (MLS)    │ ← End-to-end encryption
├─────────────────────────────────────┤
│        Transport Security           │ ← TLS, connection security
├─────────────────────────────────────┤
│           Network Layer             │ ← Internet routing
└─────────────────────────────────────┘
```

### MLS vs. Other Protocols

**MLS vs. Signal Protocol:**
| Feature | MLS | Signal Protocol |
|---------|-----|-----------------|
| Group Efficiency | ✅ Efficient tree-based | ❌ Pairwise only |
| Forward Secrecy | ✅ Group forward secrecy | ✅ Pairwise forward secrecy |
| Post-Compromise Security | ✅ Automatic healing | ⚠️ Manual detection required |
| Standardization | ✅ IETF RFC 9420 | ❌ Proprietary specification |
| Scalability | ✅ Large groups | ❌ Small groups only |

**MLS vs. Traditional TLS:**
| Feature | MLS | TLS |
|---------|-----|-----|
| End-to-End | ✅ Client-to-client | ❌ Client-to-server |
| Forward Secrecy | ✅ Automatic rotation | ⚠️ Depends on configuration |
| Group Messaging | ✅ Native support | ❌ Not supported |
| Server Access | ✅ Zero server access | ❌ Server can decrypt |

---

## Real-World Examples

### Example 1: Simple Chat Security

```go
// From https://github.com/joinself/academy/blob/main/examples/server/03_chat/01_basic/go/main.go
// When you send a chat message:

1. Your app calls acc.MessageSend(peer, content)
2. MLS encrypts message with current group key
3. Message sent through Self network infrastructure  
4. Recipient's MLS automatically decrypts message
5. OnMessage callback delivers plaintext to their app
```

**Security Properties Applied:**
- **Confidentiality**: Only recipient can read message
- **Authenticity**: Recipient knows message came from you
- **Integrity**: Any tampering detected and rejected
- **Forward Secrecy**: Future key compromise won't affect this message

### Example 2: Group Connection Process

```go
// From https://github.com/joinself/academy/blob/main/examples/server/04_advanced_features/01_core_features/go/main.go
// When someone scans your QR code:

1. QR contains discovery request with key package
2. Their device sends connection request to your account
3. OnKeyPackage callback receives their cryptographic material
4. MLS negotiates shared group keys
5. OnWelcome callback confirms secure connection established
6. Future messages encrypted with shared group key
```

**Security Properties Applied:**
- **Mutual Authentication**: Both parties verify identities
- **Key Exchange**: Secure establishment of shared secrets
- **Group Formation**: Secure "group" created (even for 1:1 chat)
- **Future Secrecy**: Connection ready for secure messaging

### Example 3: Group Member Addition

```go
// When adding a new member to group conversation:

1. Existing member generates invitation key package
2. New member receives invitation (via QR, deep link, etc.)
3. MLS performs group key derivation for new member set
4. All existing members receive OnCommit callback (group changed)
5. New member receives OnWelcome callback (joined group)
6. All future messages use new group key (includes new member)
7. Past messages remain unreadable to new member (forward secrecy)
```

**Security Properties Applied:**
- **Group Forward Secrecy**: New member can't read past messages
- **Efficient Key Update**: O(log n) complexity instead of O(n²)
- **Automatic Healing**: Any compromise before addition is healed
- **Consistent Security**: Same security level regardless of group size

---

## 🎯 Advanced Features

### Asynchronous Messaging

**Challenge**: Recipients might be offline when message sent
**MLS Solution**: Pre-generated key packages enable offline delivery

```go
// From https://github.com/joinself/academy/blob/main/examples/server/04_advanced_features/01_core_features/go/main.go
keyPackage, err := d.account.ConnectionNegotiateOutOfBand(
    inboxAddress,
    time.Now().Add(15*time.Minute), // Valid even if offline
)
```

**How It Works:**
1. **Key Pre-Generation**: Recipients create key packages in advance
2. **Key Publication**: Key packages stored on Self network
3. **Offline Encryption**: Senders use published keys to encrypt
4. **Delayed Decryption**: Recipients decrypt when they come online
5. **Security Maintained**: Full MLS security even for offline recipients

### Credential Integration

**MLS + Verifiable Credentials = Enhanced Security:**

```go
// From https://github.com/joinself/academy/blob/main/examples/server/02_credentials/01_issuing_credentials/01_basic/go/main.go
// Credentials can be shared through MLS-protected channels

1. Issue credential using normal credential API
2. Share credential through encrypted MLS message
3. Recipient receives credential via OnMessage callback
4. Credential integrity protected by both:
   - Ed25519 signature (credential authenticity)
   - MLS encryption (transmission confidentiality)
```

**Benefits:**
- **Double Protection**: Credential signatures + MLS encryption
- **Private Sharing**: No eavesdropping on credential exchange
- **Authentic Delivery**: Guaranteed sender identity
- **Selective Disclosure**: Share different credentials to different groups

### Cross-Device Synchronization

**Multi-Device Security:**
```go
// Each device gets its own identity but can join same groups
// MLS treats each device as separate group member
// Provides security isolation between devices
```

**Benefits:**
- **Device Independence**: Compromise of one device doesn't affect others
- **Selective Sync**: Choose which conversations sync to which devices
- **Recovery Capability**: Add new device without compromising security
- **Granular Control**: Different security policies per device

---

### Federation and Interoperability

**MLS Standardization Benefits:**
- Other messaging systems can implement MLS
- Potential for cross-platform secure messaging
- Self SDK ready for federated communication

---

## 🎯 Next Steps

Now that you understand MLS security concepts, explore them in practice:

**Continue Learning:**
- **[Cryptographic Foundations](cryptographic-foundations.md)** - Mathematical primitives used in MLS
- **[Secure Connections](secure-connections.md)** - Connection establishment details
- **[Decentralized Identity](decentralized-identity.md)** - Identity layer above MLS

**Practice with Examples:**
- **[Chat Examples](../examples/chat.md)** - Experience MLS encryption in action
- **[Connection Examples](../examples/connections.md)** - See group formation process
- **[Advanced Features](../examples/advanced.md)** - Explore MLS callbacks and events

**Deep Dive Resources:**
- **RFC 9420**: MLS specification
- **IETF MLS Working Group**: Latest protocol developments
- **Self SDK Documentation**: Implementation-specific details
