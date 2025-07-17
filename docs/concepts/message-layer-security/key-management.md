# Key Management Deep Dive

MLS's security comes from sophisticated key management that happens automatically.

### Key Lifecycle

**Key Generation:**
<div data-github-embed="https://github.com/joinself/academy/blob/main/examples/server/04_advanced_features/01_core_features/go/main.go#L195-L199"
     data-style="github-dark-dimmed"
     data-show-border="true"
     data-show-line-numbers="true"
     data-show-file-meta="true"
     data-show-full-path="true"
     data-show-copy="true"></div>

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
<div data-github-embed="https://github.com/joinself/academy/blob/main/examples/server/04_advanced_features/01_core_features/go/main.go#L285-L288"
     data-style="github-dark-dimmed"
     data-show-border="true"
     data-show-line-numbers="true"
     data-show-file-meta="true"
     data-show-full-path="true"
     data-show-copy="true"></div> 
