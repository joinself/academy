# Message Layer Security

> **Hands-on Learning:** After reading this, try the [Chat Examples](../examples/chat.md)

## What You'll Learn

Message Layer Security (MLS) is the protocol that powers Self SDK's secure communication. It provides end-to-end encryption, forward secrecy, and post-compromise security for both individual and group conversations. This page explains how MLS works and why it's essential for secure messaging.

**Learning Goals:**

- Understand how MLS provides end-to-end encryption
- Learn about forward secrecy and post-compromise security
- Explore group messaging security challenges and solutions
- See how MLS callbacks handle protocol events
- Discover the security guarantees MLS provides

---

## What is Message Layer Security?

**Message Layer Security (MLS)** is a cryptographic protocol designed for secure group messaging. Unlike traditional point-to-point encryption, MLS enables efficient, scalable security for groups of any size while maintaining strong security properties.

### Core Security Properties

**End-to-End Encryption:**

- Messages encrypted before leaving sender's device
- Only intended recipients can decrypt messages
- Network infrastructure cannot read message content

**Forward Secrecy:**

- Compromising current keys doesn't affect past messages
- Keys automatically rotated and old keys deleted
- Past communications remain secure even after key theft

**Post-Compromise Security:**

- System can heal from key compromise
- New secure communication established automatically
- Group continues functioning even after partial compromise

**Efficient Group Operations:**

- Adding/removing members doesn't require N² operations
- Scalable key management for large groups
- Minimal bandwidth overhead for group changes

### Real-World Analogy

Think of MLS like a **secure conference room** that evolves:

```
Traditional Encryption = Simple Door Lock
├── One key fits all situations
├── If key stolen, everything compromised
└── Same security level forever

MLS = Smart Conference Room
├── Keys change automatically every meeting
├── New people get fresh room keys  
├── Old keys stop working immediately
├── Room "heals" if security compromised
└── All handled automatically by your device
```

---

## Resources

**Continue MLS Path:**

- **[Key Management](message-layer-security/key-management.md)** - Mathematical primitives used in MLS
- **[Group Messaging](message-layer-security/group-messaging.md)** - Mathematical primitives used in MLS
- **[Security Guarantees](message-layer-security/security-guarantees.md)** - Mathematical primitives used in MLS
- **[Protocol Architecture](message-layer-security/protocol-architecture.md)** - Mathematical primitives used in MLS
- **[Real-World Examples](message-layer-security/real-world-examples.md)** - Mathematical primitives used in MLS
- **[Advanced Features](message-layer-security/advanced-features.md)** - Mathematical primitives used in MLS

**Learn related concepts:**

- **[Cryptographic Foundations](cryptographic-foundations.md)** - Mathematical primitives used in MLS
- **[Secure Connections](secure-connections.md)** - Connection establishment details
- **[Decentralized Identity](decentralized-identity.md)** - Identity layer above MLS

**Practice with Examples:**

- **[Chat Examples](../examples/chat.md)** - Experience MLS encryption in action
- **[Connection Examples](../examples/connections.md)** - See group formation process
- **[Advanced Features](../examples/advanced.md)** - Explore MLS callbacks and events

**Deep Dive Resources:**

- **[RFC 9420](https://datatracker.ietf.org/doc/rfc9420/)**: MLS specification
- **[IETF MLS Working Group](https://datatracker.ietf.org/wg/mls/about/)**: Latest protocol developments
