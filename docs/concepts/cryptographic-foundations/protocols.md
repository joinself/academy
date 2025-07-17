# Cryptographic Protocols

### Message Layer Security (MLS)

The Self SDK implements **Message Layer Security** for group messaging:

**Key Features:**

- **Forward Secrecy**: Compromised keys don't affect past messages
- **Post-Compromise Security**: Automatic healing from key compromise
- **Efficient Group Operations**: Scalable key management for groups
- **End-to-End Encryption**: No server can decrypt messages

**Integration with Examples:**
```go
// From https://github.com/joinself/academy/tree/main/examples/server/04_advanced_features/01_core_features/go/main.go
Callbacks: account.Callbacks{
    OnWelcome:    d.onWelcomeMessage,    // MLS welcome handling
    OnKeyPackage: d.onKeyPackage,        // Key package management
    OnMessage:    d.onMessage,           // Encrypted message handling
}
```

### Connection Establishment Protocol

**Three-Phase Handshake:**

1. **Discovery Phase**: Share identity and capabilities
2. **Key Exchange Phase**: Establish shared encryption keys using X25519
3. **Secure Channel Phase**: Begin encrypted communication

```go
// From https://github.com/joinself/academy/tree/main/examples/server/01_connection/01_direct/go/main.go
connection, err := senderAccount.ConnectionNegotiate(
    recipientKey,  // Ed25519 public key for verification
    kpg.KeyPackage(), // X25519 key package for encryption
)
```

### Credential Signing Protocol

**Verifiable Credential Creation:**

1. **Claim Assembly**: Structure credential claims and metadata
2. **Signature Generation**: Ed25519 signature over credential content
3. **Credential Packaging**: Combine claims + signature + metadata

```go
// From https://github.com/joinself/academy/tree/main/examples/server/02_credentials/01_issuing_credentials/01_basic/go/main.go
emailCred, err := credential.NewCredential().
    ClaimObject(credential.ClaimTypeEmailAddress, map[string]interface{}{
        "emailAddress": "alice@example.com",
    }).
    Subject(holderAddress). // Who the credential is about
    Finish()                // Sign with issuer's Ed25519 key
``` 
