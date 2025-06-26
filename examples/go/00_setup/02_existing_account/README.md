# 📂 Existing Account Loading - Load Your Saved Account

> **🎯 What you'll learn:** How to load and use a previously created Self account from storage with proper verification

This example demonstrates how to **load existing Self accounts** that were created and saved previously. Perfect for application startup and session restoration.

## 🟢 Complexity: Beginner
**Essential pattern** - This is how applications load user accounts after initial setup.

---

## 🚀 Quick Start

```bash
# This example is self-contained - no setup required!
go run main.go

# Expected output: Account creation, closure, and loading demonstration
```

## 🎯 What You'll Learn

### 🔑 Core Concepts Demonstrated
1. **Account Creation** - Create an account for demonstration purposes
2. **Application Restart Simulation** - Close and reopen account in same script
3. **Account Discovery** - Find and validate existing account storage
4. **Account Loading** - Restore account from persistent storage
5. **Identity Persistence** - Verify that account identity is preserved exactly

### 📋 Key Features Covered
- ✅ **Storage Discovery**: Finds and validates existing account data
- ✅ **Identity Restoration**: Loads cryptographic keys and account identity
- ✅ **Network Reconnection**: Re-establishes connection to Self network
- ✅ **State Preservation**: Maintains all previous connections and data

---

## 🏗️ How It Works

### Step 1: Create Account for Demo
```go
originalAccount := createAccountForDemo()
originalAccountAddress := getAccountAddress(originalAccount)
// Creates fresh account for demonstration
// Captures original identity for comparison
```

### Step 2: Simulate Application Restart
```go
originalAccount.Close()
// Closes account to simulate application shutdown
// Tests real-world restart scenario
```

### Step 3: Load Existing Account
```go
selfAccount := loadExistingAccount()
// Loads account from persistent storage
// SDK automatically restores all data
```

### Step 4: Verify Identity Persistence
```go
verifyAccountPersistence(selfAccount, originalAccountAddress)
// Confirms identity is exactly preserved
// Demonstrates perfect data persistence
```

---

## 📊 Expected Output

```
📂 Existing Account Loading Example
===================================
🔧 STEP 1: Creating an account for demonstration...
🔧 Creating account for loading demonstration...
✅ Demo account connected to Self network
✅ Demo account created successfully!
🔄 Closing account to simulate application restart...
✅ Account closed (simulating application shutdown)

📂 STEP 2: Loading existing account from storage...
🔍 Looking for existing account storage...
✅ Found existing account storage directory
✅ Reconnected to Self network
✅ Account loaded successfully!

🔍 STEP 3: Verifying Account State and Persistence
==================================================
✅ Network connectivity: OK
✅ Account identity: did:self:abc123... (fresh inbox address)
✅ Persistence verification successful!
   • Account storage and configuration preserved
   • Cryptographic keys and identity maintained
   • Network connectivity fully restored

💡 Important Note About Inbox Addresses:
   • Inbox addresses are temporary and change between sessions
   • This is expected Self SDK behavior for security

📋 STEP 4: Demonstrating Data Persistence
==========================================
🆔 Account DID: did:self:123...
🔐 Status: Loaded and ready for secure communication
```

---

## 🔍 Account Persistence

This example demonstrates what actually persists in Self accounts across sessions:

- **🔐 Cryptographic Keys**: Account keys and signing capabilities preserved
- **🔗 Preserved Connections**: All previous connections are maintained  
- **📧 Message History**: Complete message and credential history
- **⚙️ Configuration**: All preferences and settings preserved
- **📬 Inbox Addresses**: Generated fresh each session for security

---

## 💡 When to Use This Pattern

### ✅ Perfect For:
- **Application startup**: Standard pattern for loading user accounts
- **Session restoration**: Recovering account state after restarts
- **Multi-device access**: Loading same account on different devices
- **Production applications**: Reliable account persistence

### ⚠️ Consider Alternatives For:
- **First-time setup learning**: Use `../01_new_account` to focus purely on creation
- **Send messages**: Use `../04_chat` for messaging applications

---

## 🚀 Next Steps

After loading your account:

1. **Check messages**: Try `../03_inbox_access` to explore inbox management
2. **Make connections**: Move to `../01_connection` to connect with others
3. **Send messages**: Explore `../04_chat` for messaging capabilities
4. **Send messages**: Use `../04_chat` for messaging capabilities

---

## 🎓 Learning Path Integration

This example is **Step 2** in the complete account management learning path:

1. **🆕 New Account** ← *Previous: Create your first account*
2. **📂 Existing Account** ← *You are here*
3. **📧 Inbox Access** ← *Next: Check messages and manage data*
4. **🔄 Import Backup** ← *Later: Account recovery patterns*
5. **⚙️ Custom Storage** ← *Advanced: Production configurations*

---

## 🔧 Troubleshooting

### Account Creation Failed
```
❌ This shouldn't happen - account was just created!
```
**Solution**: Check write permissions and available disk space

### Network Connection Issues
```
⚠️ Account loaded but inbox access failed
```
**Solution**: Check internet connection and Self network status

### Storage Corruption
```
❌ Failed to load account: storage error
```
**Solution**: Recreate account with `../01_new_account` if no backup exists 
