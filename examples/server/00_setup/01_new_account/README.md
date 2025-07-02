# 🆕 New Account Creation - Your First Self Account

> **🎯 What you'll learn:** How to create a brand new Self account from scratch with proper storage and network registration

This example demonstrates the **first-time setup** process for a Self SDK account. Perfect for new users, development environments, and clean demo setups.

## 🟢 Complexity: Beginner
**Perfect starting point** - This is the foundation for all other Self SDK operations.

---

## 🚀 Quick Start

### Go
```bash
cd go
go run main.go
```

### Java
```bash
cd java
gradle run                      # Run the example
gradle clean > /dev/null 2>&1   # Clean up build artifacts
```

### Expected Output
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
First, the example checks if an account already exists to prevent accidental overwrites:
```
Check for existing account storage
If exists: Guide user to existing account example
If not: Proceed with new account creation
```

### Step 2: New Account Creation
The SDK creates a completely new Self account:
```
Initialize new account with default configuration
Generate cryptographic key pairs
Create unique Decentralized Identifier (DID)
Register with Self network
```

### Step 3: Account Information Display
Once created, the account provides essential information:
```
Display the unique DID (account identifier)
Show the inbox address (for receiving connections)
Confirm encrypted storage location
```

### Step 4: Storage Information
The account data is securely stored locally:
```
Create encrypted storage directory
Save cryptographic keys securely
Store account identity and DID information
```

---

## 📊 Expected Results

### 💾 Account Storage
After running this example, you'll have a storage directory containing:

```
./storage/ (or language-specific equivalent)
├── account.db          # Account identity and keys
├── connections.db      # Connection data
└── messages.db         # Message history
```

**🔒 Security Note**: The storage directory contains encrypted account data. Keep it secure and backed up!

### 🆔 Account Information
You'll receive:
- **DID (Decentralized Identifier)**: Your unique identity on the Self network
- **Inbox Address**: Where others can send you connection requests and messages
- **Storage Location**: Where your encrypted account data is saved

---

## 💡 When to Use This Pattern

### ✅ Perfect For:
- **First-time users**: Creating their initial Self account
- **Development**: Fresh accounts for testing and development
- **Onboarding**: Guiding new users through account setup
- **Demos**: Clean slate for demonstrations and tutorials

### ⚠️ Consider Alternatives For:
- **Existing accounts**: Use `../02_existing_account` instead
- **Make connections**: Use `../../01_connection` to connect with other accounts

---

## 🚀 Next Steps

After creating your account:

1. **Load your account**: Try `../02_existing_account` to reload this account
2. **Check messages**: Use `../03_inbox_access` to explore inbox management
3. **Make connections**: Move to `../../01_connection` to connect with others
4. **Send messages**: Explore `../../03_chat` for messaging capabilities

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
⚠️  Account already exists in storage directory
```
**Solution**: Delete storage directory or use `../02_existing_account`

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
