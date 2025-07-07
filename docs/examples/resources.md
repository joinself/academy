# 📚 Resources & Community

This page consolidates all developer resources, documentation links, and community support information for Self SDK examples. Whether you're just getting started or building production applications, you'll find the tools and support you need here.

## Quick Navigation

- [📖 Related Concepts](#related-concepts)
- [🛠️ Developer Tools](#developer-tools)
- [📋 Standards & Specifications](#standards--specifications)
- [👥 Community Support](#community-support)
- [🔗 SDK Resources](#sdk-resources)
- [📱 Integration Resources](#integration-resources)

---

## 📖 Related Concepts

### **Core Concepts**
Essential theoretical foundations for understanding Self SDK:

- **[Decentralized Identity](../concepts/decentralized-identity.md)** - Why traditional identity is broken and how decentralized identity fixes it
- **[Cryptographic Foundations](../concepts/cryptographic-foundations.md)** - Mathematical basis of Self SDK security (Ed25519, MLS, etc.)
- **[Secure Connections](../concepts/secure-connections.md)** - How identities connect safely using cryptographic handshakes
- **[Message Layer Security](../concepts/message-layer-security.md)** - Protocol details and security properties of MLS
- **[Verifiable Credentials](../concepts/verifiable-credentials.md)** - W3C standard implementation for digital credentials

### **Architecture & Design**
Deep dives into system design and integration patterns:

- **[System Overview](../architecture/system-overview.md)** - Complete technical architecture and component interaction
- **[Security Model](../architecture/security-model.md)** - Threat analysis, mitigations, and security guarantees
- **[Integration Patterns](../architecture/integration-patterns.md)** - Production deployment patterns and best practices

---

## 🛠️ Developer Tools

### **SDK Documentation**
Official documentation and API references:

- **Self Go SDK Documentation** - Complete API reference and advanced usage patterns
- **Self Java SDK Documentation** - Comprehensive guide for Java/Kotlin applications
- **Self Mobile SDK Documentation** - iOS and Android integration guides

### **Development Utilities**
Tools to enhance your development experience:

- **Debug Logging** - Enable detailed SDK logging for troubleshooting:
  ```go
  cfg := &account.Config{
      LogLevel: account.LogDebug, // Enable detailed logging
  }
  ```

- **Connection Testing** - Use client examples to test any server implementation
- **Network Monitoring** - Monitor connection establishment and key exchange
- **Credential Validation** - Built-in SDK validation for credential integrity

### **Testing Tools**
Utilities for testing Self SDK applications:

- **Testing Utilities** - Tools for testing identity workflows and connection patterns
- **Mock Issuers and Holders** - Development utilities for credential testing
- **Example Applications** - Full applications demonstrating Self SDK usage patterns
- **Integration Test Suites** - Comprehensive testing patterns for production applications

### **Code Examples Repository**
All examples are available in the academy repository:

```
examples/server/
├── 00_setup/                # Account creation and management
├── 01_connection/           # Secure connections
├── 02_credentials/          # Verifiable credentials  
├── 03_chat/                 # Messaging and communication
├── 04_advanced_features/    # Production patterns
└── common/                  # Shared utilities
```

---

## 📋 Standards & Specifications

### **W3C Standards**
Self SDK implements these official standards:

- **[W3C Verifiable Credentials](https://www.w3.org/TR/vc-data-model/)** - Official standard for credential data model and verification
- **[W3C DID Core Specification](https://www.w3.org/TR/did-core/)** - Decentralized identifiers for credential subjects and issuers
- **[JSON-LD Specification](https://json-ld.org/)** - Data format used in verifiable credentials for semantic interoperability

### **Cryptographic Standards**
Security foundations based on proven cryptographic standards:

- **[Ed25519 Signature Scheme](https://tools.ietf.org/html/rfc8032)** - Fast and secure digital signatures
- **[Message Layer Security (MLS)](https://datatracker.ietf.org/doc/draft-ietf-mls-protocol/)** - Modern group messaging security protocol
- **[X3DH Key Agreement](https://signal.org/docs/specifications/x3dh/)** - Secure key exchange for initial connection establishment

### **Protocol Documentation**
Technical specifications for Self SDK protocols:

- **Self Protocol Specification** - Complete protocol documentation and wire formats
- **Network Layer Documentation** - P2P communication and discovery mechanisms
- **Storage Format Specification** - Encrypted storage formats and key management

---

## 👥 Community Support

### **Getting Help**
Multiple channels for getting assistance with Self SDK:

- **[Self Developer Forum](https://forum.joinself.com/)** - Get help from community and Self team members
- **[GitHub Issues](https://github.com/joinself/self-go-sdk/issues)** - Report bugs, request features, and track development
- **[Discord Community](https://discord.gg/joinself)** - Real-time chat with other developers building with Self SDK

### **Contributing**
Ways to contribute to the Self SDK ecosystem:

- **Bug Reports** - Help improve SDK stability by reporting issues
- **Feature Requests** - Suggest new capabilities and improvements
- **Documentation** - Contribute to guides, examples, and API documentation
- **Code Contributions** - Submit pull requests for SDK improvements
- **Community Support** - Help other developers in forums and Discord

### **Issue Reporting Best Practices**
When reporting issues, include:

1. **SDK version** and language (Go/Java/etc.)
2. **Operating system** and version
3. **Environment** (sandbox/production)
4. **Complete error message** with stack trace
5. **Steps to reproduce** the issue
6. **Expected vs actual behavior**
7. **Sample code** demonstrating the issue

### **Community Guidelines**
- Be respectful and constructive in all interactions
- Search existing issues before creating new ones
- Provide complete information when asking for help
- Follow up on issues and discussions
- Share knowledge and help other community members

---

## 🔗 SDK Resources

### **Installation Guides**
Getting started with Self SDK in your preferred language:

- **Go Installation** - `go get github.com/joinself/self-go-sdk`
- **Java/Kotlin Installation** - Gradle and Maven dependency configuration
- **Mobile SDKs** - iOS and Android installation and setup guides

### **Version Compatibility**
Understanding SDK versions and compatibility:

- **Version Checking** - `go list -m github.com/joinself/self-go-sdk`
- **Migration Guides** - Upgrade paths between major versions
- **Compatibility Matrix** - Cross-platform and cross-version compatibility
- **Release Notes** - Changes, improvements, and breaking changes in each release

### **Configuration References**
Complete configuration options for different environments:

- **Environment Settings** - Sandbox vs Production configuration
- **Network Configuration** - Custom network settings and proxies
- **Security Configuration** - Key management and storage encryption options
- **Performance Tuning** - Optimization settings for production deployments

### **Advanced Usage Patterns**
Beyond basic examples:

- **Multi-Device Synchronization** - Account pairing and device management
- **Enterprise Integration** - Large-scale deployment patterns
- **Custom Storage Backends** - Implementing custom storage providers
- **Network Customization** - Custom transport and discovery mechanisms

---

## 📱 Integration Resources

### **Mobile Development**
Resources for mobile app integration:

- **iOS Swift SDK** - Native iOS development with Self SDK
- **Android Kotlin SDK** - Native Android development patterns
- **React Native Bridge** - Cross-platform mobile development
- **Flutter Plugin** - Self SDK integration for Flutter applications

### **Web Development**
Web application integration options:

- **JavaScript SDK** - Browser-based Self SDK usage
- **WebAssembly Bridge** - High-performance web integration
- **Server-Side Integration** - Backend services using Self SDK
- **API Gateway Patterns** - Exposing Self SDK functionality via REST APIs

### **Backend Integration**
Server and infrastructure integration:

- **Microservices Architecture** - Self SDK in distributed systems
- **Database Integration** - Storing credentials and connection data
- **Message Queue Integration** - Asynchronous message processing
- **Monitoring and Observability** - Production monitoring patterns

### **Enterprise Features**
Advanced features for enterprise deployments:

- **SSO Integration** - Single Sign-On with existing identity systems
- **LDAP/Active Directory** - Enterprise directory integration
- **Audit and Compliance** - Meeting regulatory requirements
- **High Availability** - Production deployment patterns for reliability

---

## 🔧 Quick References

### **Common Commands**
Frequently used commands during development:

```bash
# Check SDK version
go list -m github.com/joinself/self-go-sdk

# Run basic examples
cd examples/server/00_setup/01_new_account/go
go run main.go

# Enable debug logging
export SELF_LOG_LEVEL=debug

# Clean restart (remove all storage)
rm -rf *_storage/
```

### **Configuration Templates**
Common configuration patterns:

```go
// Development configuration
cfg := &account.Config{
    StorageKey:  generateStorageKey("dev"),
    StoragePath: "./dev_storage",
    Environment: account.TargetSandbox,
    LogLevel:    account.LogDebug,
}

// Production configuration
cfg := &account.Config{
    StorageKey:  loadSecureStorageKey(),
    StoragePath: "/var/lib/myapp/storage",
    Environment: account.TargetProduction,
    LogLevel:    account.LogWarn,
}
```

### **Error Patterns**
Common error patterns and solutions:

- **Account Creation Errors** - See [Setup Issues](troubleshooting.md#setup--account-issues)
- **Connection Failures** - See [Connection Issues](troubleshooting.md#connection-issues)
- **Network Problems** - See [Network Issues](troubleshooting.md#network--connectivity-issues)
- **Credential Issues** - See [Credential Issues](troubleshooting.md#credential-issues)

---

## 📞 Need More Help?

### **Escalation Path**
If you can't find what you need:

1. **Check Documentation** - Start with concept pages and examples
2. **Search Community** - Forum and Discord for similar questions
3. **Review Issues** - GitHub issues for known problems and solutions
4. **Ask Community** - Post in forum or Discord with specific details
5. **Report Bug** - Create GitHub issue with reproduction steps

### **Professional Support**
For enterprise customers:

- **Professional Services** - Implementation assistance and consultation
- **Priority Support** - Dedicated support channels for enterprise customers
- **Custom Development** - Specialized features and integrations
- **Training Services** - Team training and best practices workshops

---

**Everything you need to build with Self SDK!** 🚀

*This resource guide is continuously updated based on community feedback and SDK development. Please contribute improvements via GitHub issues or community forums.* 
