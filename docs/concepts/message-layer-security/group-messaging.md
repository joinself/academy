# 👥 Group Messaging Security

MLS shines in group scenarios where traditional encryption becomes complex and inefficient.

### The Group Challenge

**Without MLS (Traditional Approach):**
```
Group of 4 people = 6 pairwise connections
├── A ↔ B, A ↔ C, A ↔ D
├── B ↔ C, B ↔ D  
└── C ↔ D

Problems:
❌ Message sent 3 times (inefficient)
❌ Keys managed separately (complex)
❌ Adding member requires 4 new connections
❌ No forward secrecy across group
```

**With MLS (Efficient Approach):**
```
Group of 4 people = 1 secure group
├── Shared group encryption key
├── Individual identity verification
├── Efficient key derivation tree
└── Automatic member management

Benefits:
✅ Message sent once (efficient)
✅ One key derivation for all (simple)
✅ Adding member requires one operation
✅ Forward secrecy for entire group
```

### Group Operations in Self SDK

**Creating Secure Groups:**
```go
// Groups are created automatically when connections are established
// Every connection in Self SDK is actually a "group" (even 1:1 conversations)

groupAddress, err := selfAccount.ConnectionEstablish(
    kpg.ToAddress(),  // Group coordination address
    kpg.KeyPackage(), // Cryptographic material
)
// Result: groupAddress represents a secure MLS group
```

**Group State Management:**
- **Membership**: Who can send/receive messages
- **Keys**: Current encryption material for the group
- **History**: Secure audit trail of group changes
- **Epochs**: Security periods with distinct keys

### Real Group Example

<div data-github-embed="https://github.com/joinself/academy/blob/main/examples/server/04_advanced_features/01_core_features/go/main.go#L241-L283"
     data-style="github-dark-dimmed"
     data-show-border="true"
     data-show-line-numbers="true"
     data-show-file-meta="true"
     data-show-full-path="true"
     data-show-copy="true"></div>


**What Just Happened:**
1. **Welcome Event**: Someone wants to join secure communication
2. **Group Update**: MLS adds them to the secure group  
3. **Key Derivation**: New group keys derived for all members
4. **Message Encryption**: Welcome message encrypted with new group key
5. **Forward Secrecy**: Previous keys deleted, old messages stay secure 
