# Cryptographic Deep Dive: The Math Behind The Magic

### **Key Exchange Protocol**

Self SDK uses **Message Layer Security (MLS)** protocol for group key management:

```
Mathematical Guarantees:
├── Perfect Forward Secrecy (PFS)
├── Post-Compromise Security (PCS)  
├── Authentication of all participants
└── End-to-end encryption by default
```

### **Forward Secrecy in Action**

<div data-github-embed="https://github.com/joinself/academy/blob/main/examples/server/01_connection/02_qr/go/main.go#L87-L91"
     data-style="github-dark-dimmed"
     data-show-border="true"
     data-show-line-numbers="true"
     data-show-file-meta="true"
     data-show-full-path="true"
     data-show-copy="true"></div>


**What this means**:
- **Each connection uses different keys** - compromising one doesn't affect others
- **Keys expire automatically** - old keys become useless over time  
- **Key rotation happens automatically** - SDK handles key refresh behind the scenes
- **Past messages stay secure** - even if current keys are stolen

### **Identity Verification Process**

```go
// Every connection cryptographically verifies identity
fmt.Printf("Connection received from: %s\n", kpg.FromAddress().String())
// FromAddress is cryptographically verified
// Cannot be spoofed or impersonated
// Backed by their private key signature
```

**Security guarantee**: When you see a `FromAddress`, you can **mathematically prove** that message came from the holder of that DID's private key. 
