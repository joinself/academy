# Setup & Configuration Examples

> **Theory foundation:** [Decentralized Identity Concepts](../concepts/decentralized-identity.md)  
> **What you'll learn:** How to create, load, and manage Self identities with cryptographic ownership

Welcome to your **first hands-on experience** with decentralized identity! These examples transform the concepts you've learned into working code that creates real Self accounts.

---

## What You'll Learn

By completing these setup examples, you'll master:

- **Identity Creation**: Generate cryptographically secure Self accounts from scratch
- **Key Management**: Understand how private keys create ownership and security
- **Persistence**: How Self accounts maintain identity across sessions
- **Network Integration**: Connect your identity to the Self network for communication
- **Account Lifecycle**: Complete workflow from creation to production use

**Time Investment**: 15 minutes to complete all three examples  
**Immediate Result**: Working Self identity ready for messaging and credentials

---

## The Big Picture: From Theory to Practice

In [Decentralized Identity Concepts](../concepts/decentralized-identity.md), you learned **why** traditional identity is broken and **how** decentralized identity fixes it. Now you'll **build** your own decentralized identity and see the theory in action.

### Theory → Practice Connection

```
**Resources:** THEORY                          **Tools:** PRACTICE
Traditional identity fails    →    Create password-free account
DIDs provide ownership        →    Generate your unique DID
Cryptographic keys secure     →    Keys stored safely on device
Self-sovereign control        →    You own your identity completely
```

---

## Complete Setup Journey

Follow this **progressive learning path** to master Self account management:

### **Step 1:** Create Your First Identity

**What it demonstrates:**

- Fresh DID generation from cryptographic keys
- Automatic secure storage initialization  
- Network registration and identity verification
- Foundation for all other Self SDK operations

<div data-github-embed="https://github.com/joinself/academy/blob/main/examples/server/common/account.go#L33-L41"
     data-style="github-dark-dimmed"
     data-show-border="true"
     data-show-line-numbers="true"
     data-show-file-meta="true"
     data-show-full-path="true"
     data-show-copy="true"></div>

**Key Concept**: Unlike username/password systems, this identity is **mathematically yours** - no company can take it away.

**Time**: 2 minutes to complete  
**Success**: You'll see your unique DID and confirm network connection

**View complete implementation:** [New Account Creation](https://github.com/joinself/academy/blob/main/examples/server/00_setup/01_new_account/)

---

### **Step 2:** Master Account Persistence

**What it demonstrates:**

- How Self accounts persist across application restarts
- Identity preservation and verification
- Network reconnection with existing credentials
- Real-world application startup patterns

**Key Insight**: The code to recover an existing account is **exactly the same** as creating a new one - as long as you use the same storage path and storage key, the Self SDK automatically detects and loads your existing identity.

<div data-github-embed="https://github.com/joinself/academy/blob/main/examples/server/common/account.go#L33-L41"
     data-style="github-dark-dimmed"
     data-show-border="true"
     data-show-line-numbers="true"
     data-show-file-meta="true"
     data-show-full-path="true"
     data-show-copy="true"></div>

**How it works:**
- **Same function call**: `common.SetupAccount()` works for both creation and recovery
- **Automatic detection**: SDK checks if storage exists and loads existing account
- **Storage key matters**: Same key = same account, different key = different account
- **Storage path matters**: Same path = same account, different path = different account

**Key Concept**: Your identity is **permanently yours** - it persists across devices, applications, and time, and the same simple code recovers it.

**Time**: 2 minutes to complete  
**Success**: Same DID loads perfectly, proving identity ownership

**View complete implementation:** [Existing Account Loading](https://github.com/joinself/academy/blob/main/examples/server/00_setup/02_existing_account/)

---

### **Step 3:** Access Your Identity Network

**What it demonstrates:**
- How to get your shareable inbox address
- The difference between DIDs and inbox addresses
- Network communication setup
- Foundation for receiving messages and credentials

<div data-github-embed="https://github.com/joinself/academy/blob/main/examples/server/00_setup/03_inbox_access/go/main.go#L18-28"
     data-style="github-dark-dimmed"
     data-show-border="true"
     data-show-line-numbers="true"
     data-show-file-meta="true"
     data-show-full-path="true"
     data-show-copy="true"></div>

**Key Concept**: Your inbox address is like your **email for the decentralized web** - others need it to reach you.

**Time**: 1 minute to complete  
**Success**: Working inbox address ready for communication

**View complete implementation:** [Inbox Access & Addressing](https://github.com/joinself/academy/blob/main/examples/server/00_setup/03_inbox_access/)

---

## What Just Happened? Theory in Action

### **Identity Revolution Experienced**
You've just experienced the **identity revolution** firsthand:

- **No passwords** - Your identity is secured by cryptography, not memorized secrets
- **True ownership** - Your private key means YOU control your identity  
- **Universal identity** - One identity works across all Self-enabled applications
- **Maximum security** - Cryptographic proofs instead of vulnerable databases

### **Decentralized vs Traditional: You Built Both**

| Traditional Identity | Your Self Identity |
|---------------------|-------------------|
| Company controls | **You control** |
| Password vulnerable | **Cryptographically secure** |
| App-specific | **Universal across apps** |
| Can be revoked | **Permanently yours** |

### **Real-World Impact Achieved**
Your new Self identity can now:

- **Replace passwords** in Self-enabled applications
- **Receive messages** without revealing personal information
- **Store credentials** that others can verify cryptographically
- **Connect securely** with other Self users worldwide

---

## Technical Foundation Established

### **Architecture Understanding**
Through these examples, you've built this complete system:

```
Your Self Identity System
├── Cryptographic Keys (your ownership proof)
├── Unique DID (your permanent identifier)  
├── Encrypted Storage (your secure data vault)
├── Network Connection (your communication channel)
└── Inbox Address (your public communication endpoint)
```

### **Security Model Implemented**
- **Private Key**: Never leaves your device, proves identity ownership
- **Public Key**: Shared with network, enables others to verify your signatures
- **DID**: Permanent identifier derived from your public key
- **Encrypted Storage**: All account data protected on your device

### **Network Integration Achieved**  
- **Registration**: Your DID is registered with Self network for discovery
- **Communication**: Inbox address enables peer-to-peer messaging
- **Verification**: Network can verify your identity cryptographically
- **Interoperability**: Works with all Self-enabled applications

---

## Next Steps: Build On Your Foundation

With your Self identity established, you're ready for advanced patterns:

### **Level 1: Communication** 🟡
- **[Connection Examples](connections.md)** - Connect securely with other Self users
- **[Chat Examples](chat.md)** - Send encrypted messages using your identity

### **Level 2: Credentials** 🟡  
- **[Credential Examples](credentials.md)** - Issue and verify digital credentials
- **Verifiable Claims** - Make cryptographically provable statements

### **Level 3: Advanced Features** 🟠
- **[Advanced Examples](advanced.md)** - Production patterns and optimization
- **Multi-device** - Use your identity across multiple devices

---

## **Production:** Production Deployment

**Ready for production?** Check our comprehensive **[Production Deployment Guide](production.md)** for everything you need to deploy Self SDK applications in production:

- **Quick Start:** **[Moving to Production](production.md#moving-to-production)** - Environment migration, configuration management, and data migration strategies
- **Security:** **[Security Hardening](production.md#security-hardening)** - Key management, network security, and application security best practices  
- **Performance:** **[Performance Optimization](production.md#performance-optimization)** - Connection pooling, caching strategies, and resource optimization
- **Overview:** **[Monitoring & Observability](production.md#monitoring--observability)** - Logging, health checks, and alerting systems
- 📈 **[Scalability Patterns](production.md#scalability-patterns)** - Load balancing, data partitioning, and distributed architectures
- **Process:** **[Deployment & Operations](production.md#deployment--operations)** - Blue-green deployments, configuration management, and maintenance

The production guide includes code examples, checklists, and templates for enterprise-grade Self SDK deployments.

---

## Success Checklist

Confirm you've mastered the setup fundamentals:

**Identity Creation**

- [ ] Can create new Self accounts from scratch
- [ ] Understand how DIDs are generated from keys
- [ ] Know where account data is stored securely

**Account Management**  

- [ ] Can load existing accounts reliably
- [ ] Understand identity persistence across sessions
- [ ] Can verify account integrity after loading

**Network Integration**

- [ ] Can access inbox addresses for communication
- [ ] Understand how others can reach your identity
- [ ] Ready to receive connections and messages

**Conceptual Understanding**

- [ ] Know why decentralized identity is revolutionary  
- [ ] Understand cryptographic ownership model
- [ ] Can explain benefits to other developers

---

## **Troubleshooting:** Need Help?

**Having issues?** Check our comprehensive **[Troubleshooting Guide](troubleshooting.md)** for solutions to common setup and account problems, including:

- **Architecture:** **[Setup & Account Issues](troubleshooting.md#setup--account-issues)** - Storage conflicts, account creation, and loading problems
- **Network:** **[Network Issues](troubleshooting.md#network--connectivity-issues)** - Connectivity and Self network problems  
- **Storage:** **[Storage & Permission Issues](troubleshooting.md#storage--permission-issues)** - Directory permissions and storage corruption

The troubleshooting guide includes detailed solutions, common causes, and debugging tips for all Self SDK examples.

---

## **Resources:** Resources & Next Steps

**Ready to dive deeper?** Check our comprehensive **[Resources & Community Guide](resources.md)** for everything you need to build with Self SDK:

- **Documentation:** **[Related Concepts](resources.md#related-concepts)** - Deep dives into cryptographic foundations and system architecture
- **Tools:** **[Developer Tools](resources.md#developer-tools)** - SDK documentation, testing utilities, and debugging tools
- **Community:** **[Community Support](resources.md#community-support)** - Forums, Discord, and GitHub for getting help and contributing

The resources guide includes complete SDK documentation, standards references, community guidelines, and integration resources for all platforms.

---

**Congratulations!** You've built your first decentralized identity and experienced the future of digital identity firsthand. Your Self account is now ready for secure communication, credential exchange, and passwordless authentication across the Self ecosystem.

**Ready to connect with others?** Continue with [Connection Examples](connections.md) to establish secure relationships with other Self users!
