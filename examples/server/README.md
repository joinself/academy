# Self SDK Server Examples

Welcome to the Self SDK server examples! This directory contains a comprehensive learning journey from basic concepts to production-ready applications. Whether you're new to Self SDK or building advanced features, we've got you covered.

These examples demonstrate the full capabilities of the Self SDK using direct core SDK integration. Each example is designed to teach specific concepts while building towards production-ready patterns.

## 🚀 Quick Start

**Just want to see something work right now?**

### Go
```bash
# Try the simplest example first
cd 00_setup/01_new_account/go && go run main.go
```

### Java
```bash
# Java implementations coming soon!
```

**Want to jump to a specific feature?**

| I want to... | Go here | Complexity |
|--------------|---------|------------|
| 🆕 **Create an account** | [`00_setup/01_new_account/`](00_setup/01_new_account/) | 🟢 Beginner |
| 🔗 **Connect clients** | [`01_connection/01_direct/`](01_connection/01_direct/) | 🟢 Beginner |
| 💬 **Send messages** | [`03_chat/01_basic/`](03_chat/01_basic/) | 🟢 Beginner |
| 🆔 **Create credentials** | [`02_credentials/01_issuing_credentials/01_basic/`](02_credentials/01_issuing_credentials/01_basic/) | 🟢 Beginner |
| 🔄 **Share credentials** | [`02_credentials/02_exchanging_credentials/`](02_credentials/02_exchanging_credentials/) | 🟡 Intermediate |
| 🔔 **Send notifications** | [`04_advanced_features/03_notifications/`](04_advanced_features/03_notifications/) | 🟡 Intermediate |
| 💾 **Store data securely** | [`04_advanced_features/04_storage/`](04_advanced_features/04_storage/) | 🟠 Advanced |
| 🔗 **Sync across devices** | [`04_advanced_features/02_pairing/`](04_advanced_features/02_pairing/) | 🟠 Advanced |

## 🎓 Learning Paths

### 🌱 New to Self SDK? Start Here!

**Path 1: Account Setup & Connections** (30-45 minutes)
1. **[New Account](00_setup/01_new_account/)** (🟢 2/10) → Create your first Self identity
2. **[Direct Connection](01_connection/01_direct/)** (🟢 3/10) → Establish secure connections
3. **[QR Connections](01_connection/02_qr/)** (🟢 3/10) → Visual connection discovery
4. **[Basic Chat](03_chat/01_basic/)** (🟢 4/10) → Learn secure messaging

**Path 2: Credential Fundamentals** (45-60 minutes)
1. **[Basic Credential Issuance](02_credentials/01_issuing_credentials/01_basic/)** (🟢 3/10) → Create digital credentials
2. **[Multi-Claim Credentials](02_credentials/01_issuing_credentials/02_multi_claim/)** (🟢 4/10) → Complex credential data
3. **[Credential Exchange](02_credentials/02_exchanging_credentials/)** (🟡 5/10) → Share and verify credentials

### 🚀 Ready for Advanced Features?

**Path 3: Production Applications** (60-90 minutes)
1. **[Core Features](04_advanced_features/01_core_features/)** (🟡 4/10) → Essential production patterns
2. **[Notifications](04_advanced_features/03_notifications/)** (🟡 4/10) → User engagement
3. **[Storage](04_advanced_features/04_storage/)** (🟠 5/10) → Data persistence
4. **[Pairing](04_advanced_features/02_pairing/)** (🟠 5/10) → Multi-device sync

### 🎯 Goal-Oriented Learning

**I want to build a messaging app:**
`00_setup/01_new_account` → `01_connection/01_direct` → `03_chat/01_basic` → `04_advanced_features/04_storage` → `04_advanced_features/03_notifications`

**I want to work with credentials:**
`00_setup/01_new_account` → `02_credentials/01_issuing_credentials/01_basic` → `02_credentials/01_issuing_credentials/02_multi_claim` → `02_credentials/02_exchanging_credentials`

**I want production-ready patterns:**
`04_advanced_features/01_core_features` → `04_advanced_features/03_notifications` → `04_advanced_features/04_storage` → `04_advanced_features/02_pairing`

## 📁 All Examples Overview

### 🟢 Beginner Examples (Perfect for getting started)

| Category | Example | What it teaches | Time | Key concepts |
|----------|---------|----------------|------|--------------|
| **Setup** | [New Account](00_setup/01_new_account/) | Creating Self identities | 10 min | Identity generation, network registration |
| **Setup** | [Existing Account](00_setup/02_existing_account/) | Loading saved accounts | 10 min | Account persistence, storage |
| **Setup** | [Inbox Access](00_setup/03_inbox_access/) | Managing account inboxes | 10 min | Message routing, inbox management |
| **Connection** | [Direct Connection](01_connection/01_direct/) | Programmatic connections | 15 min | Address-based connections, server patterns |
| **Connection** | [QR Connection](01_connection/02_qr/) | Visual discovery | 15 min | QR codes, mobile integration |
| **Connection** | [Client Connection](01_connection/03_client/) | Client-side connections | 15 min | Client patterns, connection establishment |
| **Credentials** | [Basic Issuance](02_credentials/01_issuing_credentials/01_basic/) | Creating credentials | 20 min | Credential creation, signing, claims |
| **Chat** | [Basic Chat](03_chat/01_basic/) | Secure messaging | 15 min | P2P messaging, encryption, responses |

### 🟡 Intermediate Examples (Building on the basics)

| Category | Example | What it teaches | Time | Key concepts |
|----------|---------|----------------|------|--------------|
| **Credentials** | [Multi-Claim](02_credentials/01_issuing_credentials/02_multi_claim/) | Complex credential data | 25 min | Multiple claims, structured data |
| **Credentials** | [With Evidence](02_credentials/01_issuing_credentials/03_with_evidence/) | Evidence-backed credentials | 25 min | Supporting evidence, verification |
| **Credentials** | [Complex Data](02_credentials/01_issuing_credentials/04_complex_data/) | Advanced structures | 30 min | Nested data, complex schemas |
| **Credentials** | [Email Verification](02_credentials/02_exchanging_credentials/email_verification/) | Real-world verification | 30 min | Email verification workflows |
| **Advanced** | [Core Features](04_advanced_features/01_core_features/) | Essential patterns | 20 min | Production fundamentals |
| **Advanced** | [Notifications](04_advanced_features/03_notifications/) | Push notifications | 15 min | User engagement, alerts |

### 🟠 Advanced Examples (Production-ready patterns)

| Category | Example | What it teaches | Time | Key concepts |
|----------|---------|----------------|------|--------------|
| **Credentials** | [Comprehensive](02_credentials/01_issuing_credentials/05_comprehensive/) | Full workflows | 35 min | End-to-end credential flows |
| **Credentials** | [Presentation Request](02_credentials/02_exchanging_credentials/presentation_request/) | Credential sharing | 30 min | Verification workflows |
| **Advanced** | [Pairing](04_advanced_features/02_pairing/) | Multi-device sync | 20 min | Device verification, state sync |
| **Advanced** | [Storage](04_advanced_features/04_storage/) | Data persistence | 25 min | Encryption, caching, TTL |

## 🔍 Find Examples by Feature

<details>
<summary><strong>🔧 Account Management</strong></summary>

- **[New Account](00_setup/01_new_account/)** - Creating fresh Self identities
- **[Existing Account](00_setup/02_existing_account/)** - Loading saved accounts
- **[Inbox Access](00_setup/03_inbox_access/)** - Managing message routing

</details>

<details>
<summary><strong>🔗 Connections & Discovery</strong></summary>

- **[Direct Connection](01_connection/01_direct/)** - Programmatic address-based connections
- **[QR Connection](01_connection/02_qr/)** - Visual discovery with QR codes
- **[Client Connection](01_connection/03_client/)** - Client-side connection patterns

</details>

<details>
<summary><strong>💬 Messaging & Communication</strong></summary>

- **[Basic Chat](03_chat/01_basic/)** - Secure 1-to-1 messaging with auto-responses

</details>

<details>
<summary><strong>🆔 Credentials & Identity</strong></summary>

- **[Basic Issuance](02_credentials/01_issuing_credentials/01_basic/)** - Creating and signing credentials
- **[Multi-Claim](02_credentials/01_issuing_credentials/02_multi_claim/)** - Complex credential data
- **[With Evidence](02_credentials/01_issuing_credentials/03_with_evidence/)** - Evidence-backed credentials
- **[Complex Data](02_credentials/01_issuing_credentials/04_complex_data/)** - Advanced data structures
- **[Comprehensive](02_credentials/01_issuing_credentials/05_comprehensive/)** - Full credential workflows
- **[Email Verification](02_credentials/02_exchanging_credentials/email_verification/)** - Email verification patterns
- **[Presentation Request](02_credentials/02_exchanging_credentials/presentation_request/)** - Credential sharing workflows

</details>

<details>
<summary><strong>💾 Storage & Data</strong></summary>

- **[Storage](04_advanced_features/04_storage/)** - Encrypted data persistence with TTL

</details>

<details>
<summary><strong>🔗 Device & Connectivity</strong></summary>

- **[Pairing](04_advanced_features/02_pairing/)** - Multi-device synchronization

</details>

<details>
<summary><strong>🏭 Production Features</strong></summary>

- **[Core Features](04_advanced_features/01_core_features/)** - Essential production patterns
- **[Notifications](04_advanced_features/03_notifications/)** - Push notifications for engagement

</details>

## ⚡ Quick Commands

### Go
```bash
# Run any example with these commands
cd <category>/<example>/go && go run main.go

# Examples:
cd 00_setup/01_new_account/go && go run main.go           # Create your first account
cd 01_connection/01_direct/go && go run main.go          # Start with connections
cd 03_chat/01_basic/go && go run main.go                 # Try messaging
cd 02_credentials/01_issuing_credentials/01_basic/go && go run main.go  # Learn credentials
```

### Java
```bash
# Java implementations coming soon!
cd <category>/<example>/java && java Main.java
```

## 🎯 Choose Your Adventure

### 👋 "I'm completely new to Self SDK"
**Start here:** [`00_setup/01_new_account/`](00_setup/01_new_account/) → [`01_connection/01_direct/`](01_connection/01_direct/) → [`03_chat/01_basic/`](03_chat/01_basic/)

### 💼 "I want to build credential workflows"
**Start here:** [`02_credentials/01_issuing_credentials/01_basic/`](02_credentials/01_issuing_credentials/01_basic/) → [`02_credentials/01_issuing_credentials/02_multi_claim/`](02_credentials/01_issuing_credentials/02_multi_claim/) → [`02_credentials/02_exchanging_credentials/`](02_credentials/02_exchanging_credentials/)

### 🔍 "I need a specific feature"
**Use the feature finder above** ☝️ or check the [Quick Start table](#-quick-start)

### 🎓 "I want to learn everything systematically"
**Follow the complete learning path:** All setup examples → All connection examples → All credential examples → All advanced examples

## 🛠️ Prerequisites

### Go
- **Go 1.22+** - All Go examples require Go 1.22 or later
- **Self SDK** - Dependencies handled automatically by `go.mod`

### Java
- **Java 11+** - Java implementations coming soon
- **Maven/Gradle** - Build tools for dependency management

### General
- **5-10 minutes** - Most examples run in under 10 minutes
- **Network access** - Required for Self network connectivity

## 💡 Tips for Success

- ✅ **Start simple** - Begin with `00_setup/01_new_account/` even if you're experienced
- ✅ **Follow the progression** - Each category builds on previous concepts
- ✅ **Read the READMEs** - Each example has detailed documentation
- ✅ **Try both languages** - Compare implementations when Java becomes available
- ✅ **Experiment** - Modify the code to understand how it works
- ✅ **Check complexity ratings** - Don't skip ahead too quickly

## 🌍 Multi-Language Support

This restructured format supports multiple programming languages:

- **Go**: All examples currently available
- **Java**: Coming soon - implementations will be added to each example's `java/` directory
- **Future languages**: Easy to add by creating new language subdirectories

The conceptual documentation in each README is language-agnostic, while implementation-specific details are provided for each supported language.

## 🤝 Need Help?

- 📖 **Each example has a detailed README** with troubleshooting
- 🔧 **Build issues?** Run `go mod tidy` in the example's language directory
- 🐛 **Something not working?** Check the prerequisites and error messages
- 💬 **Questions?** Check the [Self SDK Documentation](https://docs.joinself.com)

## 📚 Additional Resources

- [Self SDK Documentation](https://docs.joinself.com)
- [W3C Verifiable Credentials](https://w3.org/TR/vc-data-model/)
- [Decentralized Identity Foundation](https://identity.foundation)

---

**Ready to start?** Pick an example above and run the commands for your preferred language! 🚀
