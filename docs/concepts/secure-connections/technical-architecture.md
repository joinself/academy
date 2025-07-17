# Technical Architecture: How It All Fits Together

### **Complete Connection Architecture**

Through the examples, you've built this **cryptographically secure system**:

```
Self Connection Architecture
├── Identity Layer (DIDs + Cryptographic Keys)
├── Connection Layer (Handshakes + Key Exchange)  
├── Encryption Layer (MLS + Forward Secrecy)
├── Transport Layer (Network + Routing)
└── Application Layer (Messages + Credentials)
```

### **Network Security Model**

- **Discovery**: Inbox addresses enable secure discovery
- **Handshake**: MLS protocol ensures secure key exchange
- **Encryption**: All communication is end-to-end encrypted
- **Verification**: Every message is cryptographically authenticated
- **Rotation**: Keys rotate automatically for forward secrecy 
