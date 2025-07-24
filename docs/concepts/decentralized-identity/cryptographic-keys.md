# Cryptographic Keys: Your Digital DNA

### The Key Pair Concept

Every Self identity consists of two mathematically linked keys:

```
🔒 Private Key (Secret)          🔓 Public Key (Shareable)
├─ Never shared with anyone      ├─ Shared with everyone
├─ Used to sign messages         ├─ Used to verify signatures  
├─ Proves you own the identity   ├─ Part of your DID
└─ Like your digital DNA         └─ Like your digital fingerprint
```

### How This Works in Practice

#### 1. **Identity Creation**
```go
// When you run: go run main.go
account := createNewAccount()
// **Success:** Generates cryptographic key pair automatically
// **Success:** Creates unique DID from public key
// **Success:** Stores private key securely on device
```

#### 2. **Message Signing**
```go
// When you send a message
account.SendMessage("Hello!", recipientDID)
// **Success:** Automatically signs with your private key
// **Success:** Recipients verify with your public key
// **Success:** Proves the message came from you
```

#### 3. **Identity Verification**
```go
// When someone receives your message
if (verifySignature(message, senderPublicKey)) {
// **Success:** Message is authentic
// **Success:** Sender identity is verified
// **Success:** Message hasn't been tampered with
}
``` 
