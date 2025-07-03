# 🛠️ Setup & Configuration Examples 🟢

> **📖 Theory foundation:** [Decentralized Identity Concepts](../concepts/decentralized-identity.md)  
> **🎯 What you'll learn:** How to create, load, and manage Self identities with cryptographic ownership

Welcome to your **first hands-on experience** with decentralized identity! These examples transform the concepts you've learned into working code that creates real Self accounts.

---

## 🎯 What You'll Learn

By completing these setup examples, you'll master:

- **🆔 Identity Creation**: Generate cryptographically secure Self accounts from scratch
- **🔐 Key Management**: Understand how private keys create ownership and security
- **💾 Persistence**: How Self accounts maintain identity across sessions
- **📬 Network Integration**: Connect your identity to the Self network for communication
- **🔄 Account Lifecycle**: Complete workflow from creation to production use

**⏱️ Time Investment**: 15 minutes to complete all three examples  
**💡 Immediate Result**: Working Self identity ready for messaging and credentials

---

## 🌟 The Big Picture: From Theory to Practice

In [Decentralized Identity Concepts](../concepts/decentralized-identity.md), you learned **why** traditional identity is broken and **how** decentralized identity fixes it. Now you'll **build** your own decentralized identity and see the theory in action.

### 🔄 Theory → Practice Connection

```
📚 THEORY                          🛠️ PRACTICE
Traditional identity fails    →    Create password-free account
DIDs provide ownership       →    Generate your unique DID
Cryptographic keys secure   →    Keys stored safely on device
Self-sovereign control      →    You own your identity completely
```

---

## 📋 Complete Setup Journey 🟢

Follow this **progressive learning path** to master Self account management:

### **Step 1:** 🆕 Create Your First Identity
**[New Account Creation](../../examples/server/00_setup/01_new_account/)**

🎯 **What it demonstrates:**
- Fresh DID generation from cryptographic keys
- Automatic secure storage initialization  
- Network registration and identity verification
- Foundation for all other Self SDK operations

```go
// This simple code creates a permanent digital identity
account := common.SetupAccount(config)
// ✅ Generates unique DID: did:self:ABC123...
// ✅ Creates cryptographic key pair
// ✅ Registers with Self network
// ✅ Ready for secure communication
```

**🔑 Key Concept**: Unlike username/password systems, this identity is **mathematically yours** - no company can take it away.

**⏱️ Time**: 2 minutes to complete  
**✅ Success**: You'll see your unique DID and confirm network connection

---

### **Step 2:** 📂 Master Account Persistence
**[Existing Account Loading](../../examples/server/00_setup/02_existing_account/)**

🎯 **What it demonstrates:**
- How Self accounts persist across application restarts
- Identity preservation and verification
- Network reconnection with existing credentials
- Real-world application startup patterns

```go
// Load your previously created identity
selfAccount := loadExistingAccount()
// ✅ Restores your exact same DID
// ✅ Reloads cryptographic keys securely
// ✅ Reconnects to Self network
// ✅ Maintains all connections and data
```

**🔑 Key Concept**: Your identity is **permanently yours** - it persists across devices, applications, and time.

**⏱️ Time**: 2 minutes to complete  
**✅ Success**: Same DID loads perfectly, proving identity ownership

---

### **Step 3:** 📧 Access Your Identity Network
**[Inbox Access & Addressing](../../examples/server/00_setup/03_inbox_access/)**

🎯 **What it demonstrates:**
- How to get your shareable inbox address
- The difference between DIDs and inbox addresses
- Network communication setup
- Foundation for receiving messages and credentials

```go
// Get your public address for communication
inboxAddress, err := selfAccount.InboxOpen()
fmt.Printf("📬 Share this address: %s\n", inboxAddress.String())
// ✅ Others can now send you messages
// ✅ Ready for credential exchange
// ✅ Connected to decentralized network
```

**🔑 Key Concept**: Your inbox address is like your **email for the decentralized web** - others need it to reach you.

**⏱️ Time**: 1 minute to complete  
**✅ Success**: Working inbox address ready for communication

---

## 🎓 What Just Happened? Theory in Action

### ✅ **Identity Revolution Experienced**
You've just experienced the **identity revolution** firsthand:

- **🚫 No passwords** - Your identity is secured by cryptography, not memorized secrets
- **🔐 True ownership** - Your private key means YOU control your identity  
- **🌐 Universal identity** - One identity works across all Self-enabled applications
- **🛡️ Maximum security** - Cryptographic proofs instead of vulnerable databases

### ✅ **Decentralized vs Traditional: You Built Both**

| Traditional Identity | Your Self Identity |
|---------------------|-------------------|
| 🏢 Company controls | 👤 **You control** |
| 🔓 Password vulnerable | 🔐 **Cryptographically secure** |
| 📱 App-specific | 🌐 **Universal across apps** |
| ❌ Can be revoked | ✅ **Permanently yours** |

### ✅ **Real-World Impact Achieved**
Your new Self identity can now:

- **Replace passwords** in Self-enabled applications
- **Receive messages** without revealing personal information
- **Store credentials** that others can verify cryptographically
- **Connect securely** with other Self users worldwide

---

## 🔧 Technical Foundation Established

### 🏗️ **Architecture Understanding**
Through these examples, you've built this complete system:

```
Your Self Identity System
├── 🔑 Cryptographic Keys (your ownership proof)
├── 🆔 Unique DID (your permanent identifier)  
├── 💾 Encrypted Storage (your secure data vault)
├── 📡 Network Connection (your communication channel)
└── 📬 Inbox Address (your public communication endpoint)
```

### 🔐 **Security Model Implemented**
- **Private Key**: Never leaves your device, proves identity ownership
- **Public Key**: Shared with network, enables others to verify your signatures
- **DID**: Permanent identifier derived from your public key
- **Encrypted Storage**: All account data protected on your device

### 📡 **Network Integration Achieved**  
- **Registration**: Your DID is registered with Self network for discovery
- **Communication**: Inbox address enables peer-to-peer messaging
- **Verification**: Network can verify your identity cryptographically
- **Interoperability**: Works with all Self-enabled applications

---

## 🚀 Next Steps: Build On Your Foundation

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

### **Architecture Deep Dive** 🔴
- **[System Overview](../architecture/system-overview.md)** - Complete technical architecture
- **[Security Model](../architecture/security-model.md)** - Threat analysis and mitigations

---

## 💡 Production Considerations

### 🚀 **Moving to Production**
Your setup examples taught you development patterns. For production:

- **🔒 Key Management**: Implement proper key backup and recovery
- **📊 Monitoring**: Add logging and health checks for network connectivity  
- **🔄 Error Handling**: Robust error handling for network and storage issues
- **📱 Multi-Device**: Sync identity across user devices securely

### ⚡ **Performance Optimization**
- **Connection Pooling**: Reuse network connections efficiently
- **Storage Optimization**: Index frequently accessed account data
- **Background Sync**: Keep account state synchronized asynchronously
- **Caching**: Cache inbox addresses and network status

### 🛡️ **Security Hardening**
- **Key Rotation**: Implement periodic key rotation strategies
- **Storage Encryption**: Additional encryption layers for sensitive data
- **Network Security**: TLS configuration and certificate pinning
- **Access Control**: Proper permissions for storage directories

---

## 🎯 Success Checklist

Confirm you've mastered the setup fundamentals:

**✅ Identity Creation**
- [ ] Can create new Self accounts from scratch
- [ ] Understand how DIDs are generated from keys
- [ ] Know where account data is stored securely

**✅ Account Management**  
- [ ] Can load existing accounts reliably
- [ ] Understand identity persistence across sessions
- [ ] Can verify account integrity after loading

**✅ Network Integration**
- [ ] Can access inbox addresses for communication
- [ ] Understand how others can reach your identity
- [ ] Ready to receive connections and messages

**✅ Conceptual Understanding**
- [ ] Know why decentralized identity is revolutionary  
- [ ] Understand cryptographic ownership model
- [ ] Can explain benefits to other developers

---

## 🔧 Troubleshooting & Common Issues

### 📁 **Storage Issues**
```bash
⚠️ Account already exists in storage directory
```
**Solution**: Delete `./storage/` for fresh start, or use existing account example

### 🌐 **Network Issues**  
```bash
❌ Failed to connect to Self network
```
**Solution**: Check internet connection and firewall settings

### 🔑 **Permission Issues**
```bash
❌ Failed to create storage directory  
```
**Solution**: Ensure write permissions in working directory

### 💾 **Storage Corruption**
```bash
❌ Failed to load account: invalid storage format
```
**Solution**: Recreate account with new account example if no backup exists

---

## 📚 Additional Resources

### **🔗 Related Concepts**
- **[Cryptographic Foundations](../concepts/cryptographic-foundations.md)** - Deep dive into the mathematics
- **[Secure Connections](../concepts/secure-connections.md)** - How identities connect safely

### **🛠️ Developer Tools**
- **Self SDK Documentation** - Complete API reference
- **Example Applications** - Full applications using Self identities
- **Testing Utilities** - Tools for testing identity workflows

### **🌐 Community**
- **Self Developer Forum** - Get help and share experiences  
- **GitHub Issues** - Report bugs and request features
- **Discord Community** - Real-time chat with other developers

---

**🎉 Congratulations!** You've built your first decentralized identity and experienced the future of digital identity firsthand. Your Self account is now ready for secure communication, credential exchange, and passwordless authentication across the Self ecosystem.

**Ready to connect with others?** Continue with [Connection Examples](connections.md) to establish secure relationships with other Self users! 🚀
