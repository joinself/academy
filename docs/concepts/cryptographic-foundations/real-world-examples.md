# Real-World Examples

### Example 1: Account Creation Cryptography

```go
// From https://github.com/joinself/academy/tree/main/examples/server/00_setup/01_new_account/go/main.go
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
// From https://github.com/joinself/academy/tree/main/examples/server/01_connection/01_direct/go/main.go
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
// From https://github.com/joinself/academy/tree/main/examples/server/02_credentials/01_issuing_credentials/01_basic/go/main.go
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
