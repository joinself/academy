# Core Cryptographic Algorithms

At Self we use state-of-the-art cryptographic algorithms chosen for security, performance, and interoperability.

### Ed25519 Digital Signatures

**Ed25519** is the foundation of Self's digital signature system, providing authentication and non-repudiation.

**Key Properties:**

- **Algorithm**: Edwards-curve Digital Signature Algorithm (EdDSA)
- **Curve**: Curve25519 (edwards25519 form)
- **Key Size**: 32 bytes (private and public keys)
- **Signature Size**: 64 bytes
- **Security Level**: 128-bit (equivalent to 3072-bit RSA)
- **Hash Function**: SHA-512 for signature generation

**Why Ed25519?**
```
- High Performance: Faster than RSA and ECDSA
- Small Keys & Signatures: Efficient network usage
- Deterministic: No need for secure random per signature
- Side-Channel Resistant: Protection against timing attacks
- Simple Implementation: Fewer edge cases than ECDSA
- Collision Resistant: Hash collisions don't break signatures
```

**Real Example from Self SDK:**
<div data-github-embed="https://github.com/joinself/academy/blob/main/examples/server/04_advanced_features/01_core_features/go/main.go#L242-L250"
     data-style="github-dark-dimmed"
     data-show-border="true"
     data-show-line-numbers="true"
     data-show-file-meta="true"
     data-show-full-path="true"
     data-show-copy="true"></div>

Peer represents an Ed25519 public key for message verification, which can be used to sign any messages.

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
// From https://github.com/joinself/academy/tree/main/examples/server/01_connection/01_direct/go/main.go
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
<div data-github-embed="https://github.com/joinself/academy/blob/main/examples/server/04_advanced_features/01_core_features/go/main.go#L87-L96"
     data-style="github-dark-dimmed"
     data-show-border="true"
     data-show-line-numbers="true"
     data-show-file-meta="true"
     data-show-full-path="true"
     data-show-copy="true"></div>

> This is not usually created by the SDKs, but instead you have to provide a SHA key to encrypt the SDK local storage.

---

## Deep Dive Resources

- **[RFC 8032](https://www.rfc-editor.org/rfc/rfc8032)**: EdDSA specification
- **[RFC 7748](https://www.rfc-editor.org/rfc/rfc7748)**: X25519 specification
- **[NIST SP 800-186](https://csrc.nist.gov/publications/detail/sp/800-186/final)**: ECC cryptography recommendations
- **[Self SDK Documentation](https://joinself.com/docs/sdk/self-sdk/introduction)**: Implementation-specific details
