# Security Guarantees Explained

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
<div data-github-embed="https://github.com/joinself/academy/blob/main/examples/server/04_advanced_features/01_core_features/go/main.go#L195-L199"
     data-style="github-dark-dimmed"
     data-show-border="true"
     data-show-line-numbers="true"
     data-show-file-meta="true"
     data-show-full-path="true"
     data-show-copy="true"></div>

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
- DID-based addressing (privacy-friendly identifiers)  
// - Optional proxy routing (traffic analysis protection)
``` 
