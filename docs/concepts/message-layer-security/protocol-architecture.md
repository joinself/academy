# **Network:** Protocol Architecture

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
| Group Efficiency | **Success:** Efficient tree-based | ❌ Pairwise only |
| Forward Secrecy | **Success:** Group forward secrecy | **Success:** Pairwise forward secrecy |
| Post-Compromise Security | **Success:** Automatic healing | ⚠️ Manual detection required |
| Standardization | **Success:** IETF RFC 9420 | ❌ Proprietary specification |
| Scalability | **Success:** Large groups | ❌ Small groups only |

**MLS vs. Traditional TLS:**

| Feature | MLS | TLS |
|---------|-----|-----|
| End-to-End | **Success:** Client-to-client | ❌ Client-to-server |
| Forward Secrecy | **Success:** Automatic rotation | ⚠️ Depends on configuration |
| Group Messaging | **Success:** Native support | ❌ Not supported |
| Server Access | **Success:** Zero server access | ❌ Server can decrypt | 
