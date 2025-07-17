# Real-World Examples

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

---

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
