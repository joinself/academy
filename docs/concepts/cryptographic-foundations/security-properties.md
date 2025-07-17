# Security Properties Achieved

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
// From https://github.com/joinself/academy/tree/main/examples/server/04_advanced_features/01_core_features/go/main.go
cfg := &account.Config{
    // In production, the storage key must be securely generated and managed.
    // The generateStorageKey function is for demonstration purposes only.
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
