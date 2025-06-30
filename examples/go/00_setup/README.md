# 🔧 Account Setup Examples - Master Self SDK Account Management

> **🎯 What you'll learn:** Complete account management patterns for Self SDK, from account creation to recovery to advanced configuration

This directory contains **five complementary examples** that demonstrate the full spectrum of Self SDK account setup and management patterns. Together, they show how to handle accounts throughout their complete lifecycle.

## 🎯 Perfect Learning Progression

**🎯 Best Learning Experience:** Progress through examples 01→02→03→04→05
- Start with **`01_new_account`** to understand account creation
- Learn **`02_existing_account`** to open previously created accounts  
- Master **`03_inbox_access`** to check messages and manage account data



## 🟢 Complexity: Beginner to Intermediate
**Perfect foundation** - Examples progress from beginner account creation to advanced storage patterns.

---

## 🎯 Quick Decision Guide

**Choose your learning path based on your scenario:**

| I want to... | Start with | Then try | Why this path? |
|--------------|------------|----------|----------------|
| 🆕 **Create my first account** | [01_new_account/](01_new_account/) | [03_inbox_access/](03_inbox_access/) | Fresh start workflow |
| 📂 **Open existing account** | [02_existing_account/](02_existing_account/) | [03_inbox_access/](03_inbox_access/) | Load saved account data |
| 📧 **Check account messages** | [03_inbox_access/](03_inbox_access/) | [02_existing_account/](02_existing_account/) | Message management |


| 📚 **Understand everything** | [01_new_account/](01_new_account/) → [02_existing_account/](02_existing_account/) → [03_inbox_access/](03_inbox_access/) | - | Complete core setup progression |

---

## 📊 The Five Setup Patterns

### 🆕 New Account (01_new_account/) - "First Time Setup"
**Creates a brand new Self account from scratch**

```
🖥️  New Account Creation
├── Generates fresh identity
├── Creates storage directory
├── Saves account credentials
└── Ready for connections
```

**Key Features:**
- ✅ **Fresh identity**: Creates new DID and cryptographic keys
- ✅ **Storage setup**: Initializes account storage directory
- ✅ **Account registration**: Registers with Self network
- ✅ **Ready to use**: Account ready for connections and messaging

### 📂 Existing Account (02_existing_account/) - "Load Saved Account"
**Opens an account that was previously created and saved**

```
🖥️  Existing Account Loading
├── Locates storage directory
├── Loads saved credentials
├── Reconnects to Self network
└── Restores full functionality
```

**Key Features:**
- ✅ **Storage discovery**: Finds and validates existing account data
- ✅ **Credential loading**: Restores cryptographic keys and identity
- ✅ **Network reconnection**: Re-establishes connection to Self network
- ✅ **State restoration**: Maintains all previous connections and data

### 📧 Inbox Access (03_inbox_access/) - "Check Messages & Data"
**Demonstrates how to access and manage account inbox and data**

```
🖥️  Inbox Management
├── Opens account inbox
├── Lists pending messages
├── Processes credentials
└── Manages account state
```

**Key Features:**
- ✅ **Message checking**: Lists and processes pending messages
- ✅ **Credential management**: Handles received credentials
- ✅ **Connection status**: Shows active connections
- ✅ **Account health**: Validates account state and storage





---

## 🚀 Quick Start

### 🆕 Create Your First Account
```bash
cd 01_new_account && go run main.go
# Creates fresh account and shows how to save it
```

### 📂 Open Previously Created Account
```bash
# First, create an account with example 01
cd 01_new_account && go run main.go

# Then load it with example 02
cd ../02_existing_account && go run main.go
# Loads the account created in step 1
```

### 📧 Check Your Account Inbox
```bash
cd 03_inbox_access && go run main.go
# Shows how to check messages and manage account data
```





---

## 🎓 Learning Path

### 🌱 New to Self SDK? Start Here!
All examples teach essential account management concepts:
- **Self account lifecycle** and storage patterns
- **Identity management** and cryptographic keys
- **Network connectivity** and account registration
- **Data persistence** and recovery strategies

**Recommended starting path:**
1. **[01_new_account/](01_new_account/)** - Create your first Self account
2. **[02_existing_account/](02_existing_account/)** - Learn to load saved accounts
3. **[03_inbox_access/](03_inbox_access/)** - Master inbox and message management

### 📋 Account Management Best Practices
**Production-ready patterns:**
1. **Account Creation** - Secure initial setup with proper storage
2. **Account Loading** - Reliable access to existing accounts
3. **Health Monitoring** - Check account status and connectivity


---

## 🔍 Technical Differences

| Aspect | New Account | Existing Account | Inbox Access |
|--------|-------------|------------------|--------------|
| **Purpose** | Create fresh | Load saved | Manage data |
| **Input Required** | None | Storage path | Account |
| **Storage** | Creates new | Uses existing | Reads existing |
| **Network** | Registers new | Reconnects | Accesses |
| **Use Case** | First time setup | Daily operations | Message checking |
| **Complexity** | 🟢 Beginner | 🟢 Beginner | 🟡 Intermediate |

### Code Comparison

**New Account Creation:**
```go
// Fresh account with default settings
selfAccount, err := account.New(&account.Config{
    StorageKey:  []byte(storageKey),
    StoragePath: "./storage",
    Environment: account.TargetSandbox,
})
```

**Existing Account Loading:**
```go
// Load from existing storage
selfAccount, err := account.New(&account.Config{
    StorageKey:  []byte(storageKey),
    StoragePath: "./storage", // Points to existing data
    Environment: account.TargetSandbox,
})
```

**Inbox Access:**
```go
// Open inbox and check messages
inboxAddress, err := selfAccount.InboxOpen()
messages, err := selfAccount.InboxMessages()
```





---

## 💡 When to Use Each Pattern

### ✅ Use New Account (01_new_account/) For:
- **First-time application setup**: When users install your app
- **Development and testing**: Creating fresh test accounts
- **User onboarding**: Guiding new users through account creation
- **Demo applications**: Clean slate for demonstrations

### ✅ Use Existing Account (02_existing_account/) For:
- **Application startup**: Loading user's saved account
- **Session management**: Restoring account state after app restart
- **Multi-device sync**: Accessing same account from different devices
- **Production applications**: Standard account loading pattern

### ✅ Use Inbox Access (03_inbox_access/) For:
- **Message management**: Checking for new messages and credentials
- **Account monitoring**: Verifying account health and connectivity
- **Data synchronization**: Syncing account state with server
- **User dashboards**: Displaying account status and activity





---

## 🚀 Next Steps

After mastering account setup, you're ready for:
- **[01_connection/](../01_connection/)** - Establish connections with other accounts
- **[02_credentials/](../02_credentials/)** - Issue and manage verifiable credentials
- **[04_chat/](../04_chat/)** - Build messaging applications
- **[08_advanced_features/](../08_advanced_features/)** - Explore advanced Self SDK features

Each section builds on the account management skills learned here!

---

## 📚 Additional Resources

- **Common utilities**: See [../common/](../common/) for shared account management functions
- **Production guide**: Check [../../docs/production-readiness/](../../docs/production-readiness/) for deployment best practices
- **API reference**: Full account management API documentation
- **Security guide**: Best practices for secure account storage and management 
