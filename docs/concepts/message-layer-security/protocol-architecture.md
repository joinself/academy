# 🌐 Protocol Architecture

### MLS Protocol Stack

```
┌─────────────────────────────────────┐
│           Application Layer         │ ← Your chat, credentials, etc.
├─────────────────────────────────────┤
│      Self SDK Message Layer         │ ← Message routing, addressing
├─────────────────────────────────────┤  
│     Message Layer Security (MLS)    │ ← End-to-end encryption
├─────────────────────────────────────┤
│        Transport Security           │ ← TLS, connection security
├─────────────────────────────────────┤
│           Network Layer             │ ← Internet routing
└─────────────────────────────────────┘
```

### MLS vs. Other Protocols

**MLS vs. Signal Protocol:**

| Feature | MLS | Signal Protocol |
|---------|-----|-----------------|
| Group Efficiency | ✅ Efficient tree-based | ❌ Pairwise only |
| Forward Secrecy | ✅ Group forward secrecy | ✅ Pairwise forward secrecy |
| Post-Compromise Security | ✅ Automatic healing | ⚠️ Manual detection required |
| Standardization | ✅ IETF RFC 9420 | ❌ Proprietary specification |
| Scalability | ✅ Large groups | ❌ Small groups only |

**MLS vs. Traditional TLS:**

| Feature | MLS | TLS |
|---------|-----|-----|
| End-to-End | ✅ Client-to-client | ❌ Client-to-server |
| Forward Secrecy | ✅ Automatic rotation | ⚠️ Depends on configuration |
| Group Messaging | ✅ Native support | ❌ Not supported |
| Server Access | ✅ Zero server access | ❌ Server can decrypt | 
