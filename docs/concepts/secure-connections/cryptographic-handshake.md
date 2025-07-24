# The Cryptographic Handshake: How It Really Works

### **The Three-Act Drama**

Every Self connection follows this **cryptographic protocol**:

```mermaid
graph LR
    A[DISCOVERY<br/>Find each other]
    B[NEGOTIATION<br/>Exchange keys]
    C[ESTABLISHMENT<br/>Secure channel ready]
    
    A --> B --> C
    
    style A fill:#2E3138
    style B fill:#2E3138
    style C fill:#2E3138
```


Let's see this in **actual working code**:

### **Act 1: Discovery**
**How parties find each other**

```go
// Server creates discoverable inbox address
inboxAddress, err := selfAccount.InboxOpen()
// Creates temporary address: did:self:inbox:ABC123...
// Ready to receive connection requests
// Address can be shared via QR code or direct sharing
```

****Key Concepts:** Key Concept**: Inbox addresses are **temporary** and **secure** - they expire automatically and can't be guessed or brute-forced.

### **Act 2: Negotiation** 
**How cryptographic material is exchanged**

#### Option A: QR Code Pattern (Mobile-Friendly)

<div data-github-embed="https://github.com/joinself/academy/blob/main/examples/server/01_connection/02_qr/go/main.go#L80-L112"
     data-style="github-dark-dimmed"
     data-show-border="true"
     data-show-line-numbers="true"
     data-show-file-meta="true"
     data-show-full-path="true"
     data-show-copy="true"></div>

#### Option B: Direct Address Pattern (Programmatic)

<div data-github-embed="https://github.com/joinself/academy/blob/main/examples/server/01_connection/03_client/go/main.go#L94-L110"
     data-style="github-dark-dimmed"
     data-show-border="true"
     data-show-line-numbers="true"
     data-show-file-meta="true"
     data-show-full-path="true"
     data-show-copy="true"></div>

> **Key Concept**: Both patterns exchange **cryptographic key packages** that contain the mathematical material needed for secure communication.

### **Act 3: Establishment**
**How the secure channel is created**

<div data-github-embed="https://github.com/joinself/academy/blob/main/examples/server/01_connection/01_direct/go/main.go#L127-L132"
     data-style="github-dark-dimmed"
     data-show-border="true"
     data-show-line-numbers="true"
     data-show-file-meta="true"
     data-show-full-path="true"
     data-show-copy="true"></div>

<div data-github-embed="https://github.com/joinself/academy/blob/main/examples/server/01_connection/02_qr/go/main.go#L141-L146"
     data-style="github-dark-dimmed"
     data-show-border="true"
     data-show-line-numbers="true"
     data-show-file-meta="true"
     data-show-full-path="true"
     data-show-copy="true"></div>


> **Key Concept**: The `groupAddress` represents a **secure communication group** where all messages are end-to-end encrypted. 
