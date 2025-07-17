# 🎯 Advanced Features

### Asynchronous Messaging

**Challenge**: Recipients might be offline when message sent
**MLS Solution**: Pre-generated key packages enable offline delivery

<div data-github-embed="https://github.com/joinself/academy/blob/main/examples/server/04_advanced_features/01_core_features/go/main.go#L195-L199"
     data-style="github-dark-dimmed"
     data-show-border="true"
     data-show-line-numbers="true"
     data-show-file-meta="true"
     data-show-full-path="true"
     data-show-copy="true"></div>


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
