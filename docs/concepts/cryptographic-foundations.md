# 🔐 Cryptographic Foundations

> **🔧 Hands-on Learning:** These concepts underpin all [Academy Examples](../examples/overview.md)

## What You'll Learn

The Self SDK builds on proven cryptographic primitives to provide secure, decentralized identity and messaging. This page explains the mathematical foundations that make everything work - from your first account creation to complex credential exchanges.

**🎯 Learning Goals:**
- Understand the cryptographic algorithms powering Self SDK
- Learn about Ed25519 signatures and X25519 key exchange  
- Discover how hash functions provide integrity and security
- Explore key management and storage encryption patterns
- See how these primitives combine for end-to-end security

---

## 🔑 Core Cryptographic Algorithms

The Self SDK uses state-of-the-art cryptographic algorithms chosen for security, performance, and interoperability.

### Ed25519 Digital Signatures

**Ed25519** is the foundation of Self SDK's digital signature system, providing authentication and non-repudiation.

**Key Properties:**
- **Algorithm**: Edwards-curve Digital Signature Algorithm (EdDSA)
- **Curve**: Curve25519 (edwards25519 form)
- **Key Size**: 32 bytes (private and public keys)
- **Signature Size**: 64 bytes
- **Security Level**: 128-bit (equivalent to 3072-bit RSA)
- **Hash Function**: SHA-512 for signature generation

**Why Ed25519?**
```
✅ High Performance: Faster than RSA and ECDSA
✅ Small Keys & Signatures: Efficient network usage
✅ Deterministic: No need for secure random per signature
✅ Side-Channel Resistant: Protection against timing attacks
✅ Simple Implementation: Fewer edge cases than ECDSA
✅ Collision Resistant: Hash collisions don't break signatures
```

**Real Example from Self SDK:**
```go
// From examples/server/04_advanced_features/01_core_features/go/main.go
import "github.com/joinself/self-go-sdk/keypair/signing"

// Ed25519 signatures are used throughout the SDK:
func (d *AdvancedDemo) sendAdvancedWelcomeMessage(peer *signing.PublicKey) {
    // peer represents an Ed25519 public key for message verification
    chatContent, err := message.NewChat().Message(welcomeText).Finish()
    if err := d.account.MessageSend(peer, chatContent); err != nil {
        // Message is automatically signed with Ed25519
    }
}
```

### X25519 Key Exchange

**X25519** enables secure key exchange for establishing encrypted communication channels.

**Key Properties:**
- **Algorithm**: Elliptic Curve Diffie-Hellman (ECDH)
- **Curve**: Curve25519 (Montgomery form)
- **Key Size**: 32 bytes
- **Security Level**: 128-bit
- **Purpose**: Establish shared secrets for encryption

**Connection Establishment Flow:**
```
Party A                           Party B
├── Generate X25519 keypair   ←→  ├── Generate X25519 keypair  
├── Share public key          ←→  ├── Share public key
├── Compute shared secret     ←→  ├── Compute shared secret
└── Derive encryption keys    ←→  └── Derive encryption keys
```

**Real Example from Self SDK:**
```go
// From examples/server/01_connection/01_direct/go/main.go
kpg.KeyPackage(), // Their cryptographic key package for encryption
```

### SHA-256 and SHA-512 Hash Functions

**Hash functions** provide data integrity, digital fingerprinting, and key derivation.

**SHA-256 Usage:**
- **Key Derivation**: Storage key generation from seeds
- **Integrity Checking**: Data verification and tamper detection  
- **Deterministic Generation**: Fallback for key generation

**SHA-512 Usage:**
- **Ed25519 Signatures**: Internal hash function for signature generation
- **Key Expansion**: Expanding seed material for key generation

**Real Example from Self SDK:**
```go
// From examples/server/04_advanced_features/01_core_features/go/main.go
import (
    "crypto/rand"
    "crypto/sha256"
)

// generateStorageKey creates a cryptographically secure 32-byte key
func generateStorageKey(seed string) []byte {
    key := make([]byte, 32)
    if _, err := rand.Read(key); err != nil {
        // Fallback to deterministic key generation if crypto/rand fails
        h := sha256.Sum256([]byte(fmt.Sprintf("self-sdk-%s-%d", seed, time.Now().UnixNano())))
        return h[:]
    }
    return key
}
```

---

## 🛡️ Security Properties Achieved

### Confidentiality (Encryption)

**Automatic End-to-End Encryption:**
- All messages between Self SDK accounts are automatically encrypted
- No plaintext data transmitted over networks
- Forward secrecy through ephemeral key exchange

**Storage Encryption:**
- Account data encrypted with 32-byte AES keys
- Keys derived from secure random generation or deterministic fallback
- Isolated storage per account

```go
// From examples/server/04_advanced_features/01_core_features/go/main.go
cfg := &account.Config{
    StorageKey:  generateStorageKey("advanced_demo"), // 32-byte AES key
    StoragePath: "./advanced_demo_storage",           // Encrypted storage
    Environment: account.TargetSandbox,
}
```

### Integrity (Digital Signatures)

**Message Authentication:**
- Every message signed with Ed25519 sender's private key
- Recipients verify signatures with sender's public key
- Tampering detection through signature verification

**Credential Integrity:**
- Verifiable credentials cryptographically signed by issuers
- Claims cannot be modified without detection
- Signature verification proves credential authenticity

```go
// From examples/server/02_credentials/01_issuing_credentials/05_comprehensive/go/main.go
// centralized authorities while maintaining cryptographic integrity and privacy.
```

### Non-Repudiation (Digital Signatures)

**Proof of Origin:**
- Ed25519 signatures prove message sender identity
- Private key holder uniquely able to create valid signatures
- Mathematical proof of authorship

**Credential Provenance:**
- Issuer signatures prove credential authenticity
- Non-repudiable proof of credential issuance
- Verifiable audit trail

### Authenticity (Identity Verification)

**DID-Based Identity:**
- Each account has unique Decentralized Identifier (DID)
- DID cryptographically derived from Ed25519 public key
- Self-sovereign identity without central authority

```go
// From examples/server/02_credentials/01_issuing_credentials/05_comprehensive/go/main.go
// DIDs (Decentralized Identifiers) are cryptographically verifiable identities
```

---

## 🔧 Key Management Patterns

### Key Generation

**Secure Random Generation:**
```go
// Primary method - cryptographically secure random
func generateStorageKey(seed string) []byte {
    key := make([]byte, 32)
    if _, err := rand.Read(key); err != nil {
        // Fallback to deterministic generation
        h := sha256.Sum256([]byte(fmt.Sprintf("self-sdk-%s-%d", seed, time.Now().UnixNano())))
        return h[:]
    }
    return key
}
```

**Key Properties:**
- **32-byte keys**: Standard for AES-256 encryption
- **Cryptographic quality**: Using `crypto/rand` for true randomness
- **Deterministic fallback**: SHA-256 based generation if random fails
- **Unique per account**: Isolated storage and key spaces

### Key Storage

**Encrypted Account Storage:**
- Private keys never leave encrypted storage
- Storage isolation per account prevents key mixing
- Cross-platform compatible storage format

**Storage Paths Pattern:**
```go
// From multiple examples - each account gets isolated storage
issuerCfg := &account.Config{
    StorageKey:  generateStorageKey("demo_issuer"),
    StoragePath: "./demo_issuer_storage",  // Isolated storage
}

holderCfg := &account.Config{
    StorageKey:  generateStorageKey("demo_holder"), 
    StoragePath: "./demo_holder_storage",  // Separate storage
}
```

### Key Distribution

**Public Key Sharing:**
- Ed25519 public keys shared through DIDs
- Public keys distributed in connection establishment
- No secret material shared over networks

**Connection Key Exchange:**
```go
// From examples/server/01_connection/02_qr/go/main.go
// Step 2: Generate cryptographic key package for secure communication
kpg, err := acc.KeyPackageGenerate()
```

---

## 🌐 Cryptographic Protocols

### Message Layer Security (MLS)

The Self SDK implements **Message Layer Security** for group messaging:

**Key Features:**
- **Forward Secrecy**: Compromised keys don't affect past messages
- **Post-Compromise Security**: Automatic healing from key compromise
- **Efficient Group Operations**: Scalable key management for groups
- **End-to-End Encryption**: No server can decrypt messages

**Integration with Examples:**
```go
// From examples/server/04_advanced_features/01_core_features/go/main.go
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
// From examples/server/01_connection/01_direct/go/main.go
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
// From examples/server/02_credentials/01_issuing_credentials/01_basic/go/main.go
emailCred, err := credential.NewCredential().
    ClaimObject(credential.ClaimTypeEmailAddress, map[string]interface{}{
        "emailAddress": "alice@example.com",
    }).
    Subject(holderAddress). // Who the credential is about
    Finish()                // Sign with issuer's Ed25519 key
```

---

## 🚀 Performance Characteristics

### Ed25519 Performance

**Speed Advantages:**
- **Signing**: ~40,000 signatures/second on modern CPUs
- **Verification**: ~15,000 verifications/second  
- **Key Generation**: ~10,000 keypairs/second
- **Memory**: Minimal RAM requirements (~1KB per operation)

**Comparison with RSA:**
| Operation | Ed25519 | RSA-2048 | Speedup |
|-----------|---------|----------|---------|
| Sign      | 68μs    | 1100μs   | 16x     |
| Verify    | 201μs   | 33μs     | 0.16x*  |
| Key Gen   | 51μs    | 226ms    | 4400x   |

*Note: Ed25519 verification is slower than RSA but still very fast for most applications

### Network Efficiency

**Bandwidth Optimization:**
- **Public Keys**: 32 bytes vs 256+ bytes for RSA
- **Signatures**: 64 bytes vs 256+ bytes for RSA  
- **Total Overhead**: ~50% reduction in cryptographic data size

---

## 🔮 Post-Quantum Considerations

### Current Quantum Threat

**Timeline Assessment:**
- **Short-term (5-10 years)**: Current algorithms remain secure
- **Medium-term (10-20 years)**: Quantum computers may threaten ECC
- **Long-term (20+ years)**: Large-scale quantum computers likely

**Algorithm Vulnerabilities:**
- **Ed25519**: Vulnerable to Shor's algorithm on quantum computers
- **X25519**: Also vulnerable to quantum ECDLP attacks
- **SHA-256/SHA-512**: Quantum resistant (halved security but still strong)

### Migration Strategy

**Hybrid Approach:**
The Self SDK is designed for algorithm agility, enabling future migration:

```go
// Current pattern allows for future algorithm changes
type account.Config struct {
    StorageKey  []byte           // Could use post-quantum keys
    // Future: QuantumSafeMode bool
    // Future: SignatureAlgorithm string
}
```

**Recommended Post-Quantum Algorithms:**
- **Signatures**: Dilithium, Falcon, or SPHINCS+
- **Key Exchange**: Kyber or NTRU
- **Hash Functions**: SHA-256/SHA-512 (already quantum resistant)

**Migration Considerations:**
- **Larger Key Sizes**: Post-quantum keys are typically much larger
- **Performance Impact**: Some post-quantum algorithms are slower
- **Backwards Compatibility**: Need to support legacy and new algorithms

---

## 📚 Real-World Examples

### Example 1: Account Creation Cryptography

```go
// From examples/server/00_setup/01_new_account/go/main.go
// When you create a new Self account:

1. Generate Ed25519 keypair for signatures
2. Create X25519 keys for message encryption  
3. Derive DID from Ed25519 public key
4. Generate storage encryption key
5. Register cryptographic identity on network
```

**What Just Happened:**
- Your device generated unique cryptographic keys
- Ed25519 keypair enables message signing and verification
- X25519 keys enable encrypted communication setup
- Storage key protects your private keys at rest
- DID provides your verifiable digital identity

### Example 2: Secure Connection Establishment

```go
// From examples/server/01_connection/01_direct/go/main.go
// When two accounts connect:

1. Exchange Ed25519 public keys (identity verification)
2. Generate X25519 ephemeral keys (forward secrecy)
3. Perform ECDH key agreement (shared secret)
4. Derive message encryption keys from shared secret
5. Begin encrypted communication
```

**What Just Happened:**
- Both parties verified each other's cryptographic identity
- Ephemeral keys provide forward secrecy
- No eavesdropper can decrypt your messages
- Connection is authenticated and encrypted

### Example 3: Credential Issuance Cryptography

```go
// From examples/server/02_credentials/01_issuing_credentials/01_basic/go/main.go
// When issuing a verifiable credential:

1. Issuer structures credential claims
2. Generate Ed25519 signature over claims + metadata
3. Package signature with credential content  
4. Send to holder's encrypted channel
5. Holder verifies issuer's signature
```

**What Just Happened:**
- Credential is cryptographically bound to issuer's identity
- Claims cannot be modified without signature breaking
- Anyone can verify credential authenticity using issuer's public key
- Credential provides non-repudiable proof of claims

---

## 🎯 Next Steps

Now that you understand the cryptographic foundations, explore how they enable Self SDK's capabilities:

**Continue Learning:**
- **[Message Layer Security](message-layer-security.md)** - Group messaging encryption protocols
- **[Decentralized Identity](decentralized-identity.md)** - How cryptography enables DIDs
- **[Secure Connections](secure-connections.md)** - Connection establishment details
- **[Verifiable Credentials](verifiable-credentials.md)** - Digital signature applications

**Practice with Examples:**
- **[Setup Examples](../examples/setup.md)** - See key generation in action
- **[Connection Examples](../examples/connections.md)** - Observe key exchange protocols  
- **[Credential Examples](../examples/credentials.md)** - Experience digital signatures
- **[Chat Examples](../examples/chat.md)** - Use end-to-end encryption

**Deep Dive Resources:**
- **RFC 8032**: EdDSA specification
- **RFC 7748**: X25519 specification  
- **NIST SP 800-186**: ECC cryptography recommendations
- **Self SDK Documentation**: Implementation-specific details
