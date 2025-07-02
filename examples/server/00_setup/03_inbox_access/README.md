# 📧 Inbox Access - Get Your Account Address

> **🎯 What you'll learn:** How to access your Self account's inbox address for receiving messages and connections

This example demonstrates **inbox access** for Self accounts, showing how to get your account's inbox address that others can use to send you messages, credentials, and connection requests.

## 🟢 Complexity: Basic
**Foundation skill** - Learn to access your account's public inbox address.

---

## ❗ Why This Matters

**Your inbox address is essential** - it's your identity on the Self network:

- **🏠 Your Home Address**: Like a postal address, others need this to reach you
- **📞 Your Phone Number**: Essential for any communication on the Self network
- **🎫 Your Digital Identity**: Required to participate in the decentralized ecosystem
- **🔑 Gateway to Everything**: Without it, you can't receive messages, credentials, or connections

**💡 Key Point**: Until you access your inbox address, your account exists but cannot receive anything from others.

---

## 🚀 Quick Start

```bash
# This example is self-contained - no setup required!
go run main.go

# Expected output: Account setup and inbox address display
```

## 🎯 What You'll Learn

### 🔑 Core Concepts Demonstrated
1. **Inbox Access** - Open and access your account's inbox
2. **Address Retrieval** - Get your account's public inbox address
3. **Address Sharing** - Understand how others can reach you

### 📋 Key Features Covered
- ✅ **Inbox Opening**: Access your account's inbox securely
- ✅ **Address Display**: Get your shareable inbox address
- ✅ **Usage Explanation**: Understand what the address is used for

---

## 🏗️ How It Works

### Step 1: Load Account
```go
selfAccount := common.SetupAccount(common.AccountConfig{})
```

### Step 2: Access Inbox
```go
inboxAddress, err := selfAccount.InboxOpen()
// Opens inbox and returns your public address
```

### Step 3: Share Address
```go
fmt.Printf("📬 Your inbox address: %s\n", inboxAddress.String())
// This address can be shared with others
```

---

## 📊 Expected Output

```
📧 Inbox Access Example
========================
🔧 Setting up account...
✅ Account loaded successfully

📬 Accessing inbox...
✅ Inbox opened successfully!
📬 Your inbox address: did:self:123...

💡 This address can be shared with others to receive:
   • Messages
   • Connection requests  
   • Credentials

✅ Inbox access demonstration complete!
```

---

## 💡 What is an Inbox Address?

Your **inbox address** is like your email address for the Self network:

- **🔍 Unique Identifier**: Each account has one unique inbox address
- **📬 Receiving Point**: Others use this address to send you content
- **🔐 Secure**: Built on decentralized identity (DID) standards
- **🌐 Universal**: Works across all Self-enabled applications

### 🎯 Think of it as your:
- **📧 Email address** for decentralized messaging
- **📞 Phone number** for Self network communications  
- **🏠 Home address** for receiving digital mail
- **🆔 Digital passport** for identity verification

---

## 🔄 How Others Use Your Address

When you share your inbox address, others can:

- **📨 Send Messages**: Direct peer-to-peer messaging
- **🤝 Request Connections**: Initiate secure connections
- **🎫 Issue Credentials**: Send verifiable credentials
- **📋 Share Data**: Exchange information securely

**⚠️ Important**: Without sharing your inbox address, others cannot reach you on the Self network.

---

## 💡 When to Use This Pattern

### ✅ Perfect For:
- **Profile displays**: Show your inbox address in user interfaces
- **QR codes**: Generate QR codes for easy sharing
- **Business cards**: Include your Self address for professional networking
- **Integration setup**: Configure your address in other applications

### ✅ Integration Points:
- **Contact forms**: Let others reach you via Self
- **Networking apps**: Professional and social networking
- **Service registration**: Register your address with Self services
- **Multi-device**: Access the same inbox from multiple devices

---

## 🚀 Next Steps

After getting your inbox address:

1. **Share with others**: Give your address to contacts for connection
2. **Generate QR codes**: Use connection examples to create shareable QR codes
3. **Receive messages**: Use chat examples to start messaging
4. **Issue credentials**: Try credential examples for identity verification

---

## 🎓 Learning Path Integration

This example is **Step 3** in the complete account management learning path:

1. **🆕 New Account** ← *Foundation: Account creation*
2. **📂 Existing Account** ← *Previous: Account loading*
3. **📧 Inbox Access** ← *You are here*

---

## 📚 Advanced Topics

Ready to dive deeper? This example prepares you for:

- **Connection Management**: Using your address for connections
- **Message Processing**: Building inbox monitoring systems
- **Multi-Device Access**: Accessing your inbox from multiple devices
- **Address Management**: Advanced address handling and routing

---

## 🔧 Troubleshooting

### Account Setup Issues
```
❌ Failed to create account
```
**Solution**: Check write permissions and network connectivity

### Inbox Access Failed
```
❌ Failed to open inbox: network error
```
**Solution**: Check network connectivity and Self network status

### Address Format Questions
```
📬 Your inbox address: did:self:123...
```
**Expected**: This is a standard DID (Decentralized Identifier) format used by Self
