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
// ✅ Generates cryptographic key pair automatically
// ✅ Creates unique DID from public key
// ✅ Stores private key securely on device
```

#### 2. **Message Signing**
```go
// When you send a message
account.SendMessage("Hello!", recipientDID)
// ✅ Automatically signs with your private key
// ✅ Recipients verify with your public key
// ✅ Proves the message came from you
```

#### 3. **Identity Verification**
```go
// When someone receives your message
if (verifySignature(message, senderPublicKey)) {
// ✅ Message is authentic
// ✅ Sender identity is verified
// ✅ Message hasn't been tampered with
}
``` 
