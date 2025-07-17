# Decentralized Identifiers (DIDs) Explained

### What is a DID?

A **DID** (Decentralized Identifier) is like a permanent, globally unique username that no one can take away from you.

```
Traditional: @john_doe_123 (controlled by Twitter/X)
DID: did:self:1:ABC123XYZ... (controlled by John)
```

### DID Structure Breakdown

```
did:self:1:ABC123XYZ789DEF456GHI...
│   │    │  │
│   │    │  └─ Unique identifier (derived from public key)
│   │    └─ Version number
│   └─ Self network identifier
└─ DID scheme
```

### Key Properties

#### **Cryptographically Secure**
- Generated from cryptographic key pairs
- Impossible to forge or duplicate
- Mathematically verifiable

#### **Globally Unique**
- No central registry needed
- No naming conflicts possible
- Works across all applications

#### **User-Controlled**
- Only the user has the private key
- No authority can revoke or change it
- Portable between applications 
