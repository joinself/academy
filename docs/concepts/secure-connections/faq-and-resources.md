# Common Security Questions & Troubleshooting

### **"How secure are Self connections really?"**
**Answer**: Self connections use **Message Layer Security (MLS)**, the same protocol being standardized by the IETF for next-generation secure messaging. This provides:
- Perfect Forward Secrecy (PFS)
- Post-Compromise Security (PCS)  
- Cryptographic authentication
- End-to-end encryption

### **"Why do connections expire?"**
**Answer**: Expiration provides **forward secrecy** - if keys are compromised later, past communications remain secure. It also prevents indefinite accumulation of stale connection attempts.

### **"Can connections be intercepted?"**
**Answer**: Self connections use **end-to-end encryption** with **identity verification**. Even if network traffic is intercepted, attackers cannot:
- Decrypt the messages (no keys)
- Impersonate parties (no private keys)
- Modify messages (cryptographic integrity)

### **"What happens if connection establishment fails?"**
```bash
❌ Failed to establish connection
```
**Solution**: 
- Check network connectivity
- Verify address format and validity
- Ensure both parties are online
- Check key package expiration times

---

## Additional Security Resources

### **Cryptographic Concepts**
- **[Cryptographic Foundations](cryptographic-foundations.md)** - Mathematical primitives explained
- **[Message Layer Security](message-layer-security.md)** - Deep dive into MLS protocol

### **Implementation Guides**  
- **[Connection Examples](../examples/connections.md)** - Hands-on connection patterns

### **Standards & Specifications**
- **[IETF MLS Working Group](https://datatracker.ietf.org/wg/mls/)** - Message Layer Security standard
- **[RFC 9420](https://tools.ietf.org/rfc/rfc9420.txt)** - The MLS Protocol specification
- **[Double Ratchet Algorithm](https://signal.org/docs/specifications/doubleratchet/)** - Forward secrecy foundations 
