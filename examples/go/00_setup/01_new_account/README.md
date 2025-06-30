# 🆕 New Account Creation - Your First Self Account

> **🎯 What you'll learn:** How to create a brand new Self account from scratch with proper storage and network registration

This example demonstrates the **first-time setup** process for a Self SDK account. Perfect for new users, development environments, and clean demo setups.

## 🟢 Complexity: Beginner
**Perfect starting point** - This is the foundation for all other Self SDK operations.

---

## 🚀 Quick Start

```bash
# Create your first Self account
go run main.go

# Expected output: New account created with DID and storage info
```

## 🎯 What You'll Learn

### 🔑 Core Concepts Demonstrated
1. **Fresh Account Creation** - Generate new identity and cryptographic keys
2. **Storage Initialization** - Set up persistent storage for account data  
3. **Network Registration** - Connect and register with Self network
4. **Account Information** - Understanding DIDs and inbox addresses

### 📋 Key Features Covered
- ✅ **New Identity Generation**: Creates fresh DID and cryptographic keys
- ✅ **Automatic Storage**: Initializes encrypted storage directory
- ✅ **Network Connection**: Registers with Self network automatically
- ✅ **Ready for Use**: Account immediately ready for connections and messaging

---

## 🏗️ How It Works

### Step 1: Account Existence Check
```go
if accountExists() {
    // Prevents overwriting existing accounts
    // Guides user to load existing account instead
}
```

### Step 2: New Account Creation
```go
selfAccount := common.SetupAccount(common.AccountConfig{
    // Uses default sandbox environment
    // Creates storage in ./storage/ directory
    // Generates fresh identity automatically
})
```

### Step 3: Account Information Display
```go
inboxAddress, err := selfAccount.InboxOpen()
// Shows your unique DID and inbox address
// Explains what each piece of information means
```

### Step 4: Storage Information
```go
// Explains how account data is stored
// Shows security considerations
// Guides next steps for account reuse
```

---

## 📊 Expected Output

```
🆕 New Account Creation Example
===============================
🔍 Checking for existing account...
🔧 Creating new Self account...
✅ Connected to Self network
✅ New account created successfully!

📋 New Account Information
==========================
🆔 Account DID: did:self:123...
📬 Inbox Address: did:self:123...
🔐 Status: Encrypted and ready for secure communication

💾 Account Storage
==================
📁 Storage Location: ./storage/
🔑 Storage Contents:
   • Cryptographic keys (securely encrypted)
   • Account identity and DID information

🎉 Success! Your new Self account is ready!
```

---

## 🔍 Key Files Created

After running this example, you'll have:

```
./storage/
├── account.db          # Account identity and keys
├── connections.db      # Connection data
└── messages.db         # Message history
```

**🔒 Security Note**: The storage directory contains encrypted account data. Keep it secure and backed up!

---

## 💡 When to Use This Pattern

### ✅ Perfect For:
- **First-time users**: Creating their initial Self account
- **Development**: Fresh accounts for testing and development
- **Onboarding**: Guiding new users through account setup
- **Demos**: Clean slate for demonstrations and tutorials

### ⚠️ Consider Alternatives For:
- **Existing accounts**: Use `../02_existing_account` instead
- **Make connections**: Use `../01_connection` to connect with other accounts

---

## 🚀 Next Steps

After creating your account:

1. **Load your account**: Try `../02_existing_account` to reload this account
2. **Check messages**: Use `../03_inbox_access` to explore inbox management
3. **Make connections**: Move to `../01_connection` to connect with others
4. **Send messages**: Explore `../04_chat` for messaging capabilities

---

## 🎓 Learning Path Integration

This example is **Step 1** in the complete account management learning path:

1. **🆕 New Account** ← *You are here*
2. **📂 Existing Account** ← *Next: Load your saved account*
3. **📧 Inbox Access** ← *Then: Manage messages and data*
4. **🔄 Import Backup** ← *Later: Account recovery*
5. **⚙️ Custom Storage** ← *Advanced: Production configurations*

Each step builds on the previous one to give you complete account management mastery!

---

## 🔧 Troubleshooting

### Account Already Exists
```
⚠️  Account already exists in ./storage/
```
**Solution**: Delete `./storage/` directory or use `../02_existing_account`

### Network Connection Issues
```
❌ Failed to create Self account: network error
```
**Solution**: Check internet connection and try again

### Permission Issues
```
❌ Failed to create storage directory
```
**Solution**: Ensure write permissions in current directory 
