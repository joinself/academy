# Credential Issuance Examples

This directory contains a comprehensive tutorial series for learning credential issuance using the Self SDK. The examples are designed with educational progression, starting from basic concepts and building up to advanced real-world scenarios.

> **📢 Updated Approach**: These examples now use the underlying Self SDK directly (v0.59.0+) instead of the deprecated client package facade. This provides better understanding of the core SDK architecture and patterns.

## 🚀 Quick Start

| 🎯 Goal | 📁 Example | 🏃‍♂️ Command |
|---------|------------|-------------|
| **Just want to see it work?** | Root | `go run main.go` |
| **Learn the basics** | Basic | `cd basic && go run main.go` |
| **Multiple claims** | Multi-Claim | `cd multi_claim && go run main.go` |
| **Evidence & assets** | Evidence | `cd evidence && go run main.go` |
| **Complex data** | Complex | `cd complex && go run main.go` |
| **All features** | Advanced | `cd advanced && go run main.go` |
| **Exchange foundation** | Exchange Demo | `cd exchange_demo && go run main.go` |

## 📚 Educational Progression

### 🎯 Learning Path

Complete the examples in this order for the best learning experience:

| Example | Complexity | Description | Time |
|---------|------------|-------------|------|
| **`basic/main.go`** | 🟢 **9/10** (Very Simple) | Foundation concepts | 5-10 min |
| **`multi_claim/main.go`** | 🟡 **7/10** (Intermediate) | Multiple claims in credentials | 10-15 min |
| **`evidence/main.go`** | 🟠 **5/10** (Advanced) | Evidence and asset management | 15-20 min |
| **`complex/main.go`** | 🟠 **4/10** (Advanced) | Complex nested data structures | 20-25 min |
| **`advanced/main.go`** | 🔴 **3/10** (Expert) | All features combined | 30-45 min |
| **`exchange_demo/main.go`** | 🟡 **6/10** (Intermediate) | Bridge to credential exchange | 15-20 min |

### 🔄 Current main.go

The current `main.go` contains the simplified basic issuance example using the underlying Self SDK directly (**9/10 simplicity**). This demonstrates how to use the core SDK without the client package facade. For the full educational progression, use the individual files above.

### 📊 Complexity Guide

- 🟢 **8-10/10** (Very Simple) - Perfect for beginners, minimal concepts
- 🟡 **6-7/10** (Intermediate) - Some complexity, builds on basics  
- 🟠 **4-5/10** (Advanced) - Complex concepts, production patterns
- 🔴 **1-3/10** (Expert) - Very complex, requires deep understanding

## 📖 Example Descriptions

### 1️⃣ Basic Issuance (`basic/main.go`)
**Start here if you're new to credential issuance**

- **What you'll learn:**
  - How credential issuance works
  - Basic credential creation patterns
  - Simple claim addition
  - Client setup and configuration
  - Cryptographic signing basics
  
- **Key concepts:**
  - Account setup using core SDK (issuer and holder)
  - Direct credential creation with credential.NewCredential()
  - Core SDK builder pattern usage
  - Direct credential signing and issuance

- **Complexity:** 🟢 **9/10** (Very Simple)
- **Time to complete:** 5-10 minutes

### 2️⃣ Multi-Claim Issuance (`multi_claim/main.go`)
**Prerequisites: Complete basic/main.go first**

- **What you'll learn:**
  - How to add multiple claims to a single credential
  - Different data types in claims (strings, booleans, numbers)
  - Organizing related identity information
  - Efficient credential structuring

- **Key concepts:**
  - Multiple claims in one credential
  - Different data types
  - Profile and education credentials
  - Related information grouping

- **Complexity:** 🟡 **7/10** (Intermediate)
- **Time to complete:** 10-15 minutes

### 3️⃣ Evidence-Based Issuance (`evidence/main.go`)
**Prerequisites: Complete basic/main.go and multi_claim/main.go**

- **What you'll learn:**
  - How to attach evidence files to credentials
  - Asset management and secure storage
  - Creating verifiable presentations
  - Linking evidence to claims with hashes
  - Custom credential types

- **Key concepts:**
  - Evidence and asset management
  - File attachments to credentials
  - Verifiable presentations
  - Hash-based evidence linking

- **Complexity:** 🟠 **5/10** (Advanced)
- **Time to complete:** 15-20 minutes

### 4️⃣ Complex Data Issuance (`complex/main.go`)
**Prerequisites: Complete all previous examples**

- **What you'll learn:**
  - How to structure complex nested data in credentials
  - Arrays and collections in claims
  - Hierarchical data organization
  - Real-world data modeling patterns
  - Advanced claim structuring

- **Key concepts:**
  - Complex nested objects
  - Arrays and collections
  - Hierarchical data structures
  - Real-world organizational data

- **Complexity:** 🟠 **4/10** (Advanced)
- **Time to complete:** 20-25 minutes

### 5️⃣ Advanced Issuance (`advanced/main.go`)
**Prerequisites: Complete all previous examples**

- **What you'll learn:**
  - All credential issuance features combined
  - Production-ready patterns
  - Comprehensive credential workflows
  - Request/response handling
  - Discovery integration

- **Key concepts:**
  - All features combined
  - Production patterns
  - Event-driven workflows
  - Discovery integration

- **Complexity:** 🔴 **3/10** (Expert)
- **Time to complete:** 30-45 minutes

### 6️⃣ Exchange Foundation Demo (`exchange_demo/main.go`)
**Prerequisites: Complete basic/main.go and multi_claim/main.go**

- **What you'll learn:**
  - How credential issuance enables exchange workflows
  - Credential storage and organization patterns
  - Conceptual exchange request/response workflows
  - Foundation concepts for building exchange applications
  - Bridge between issuance and exchange use cases

- **Key concepts:**
  - Credential store management
  - Exchange request simulation
  - Presentation creation from multiple credentials
  - Foundation patterns for real exchange systems

- **Complexity:** 🟡 **6/10** (Intermediate)
- **Time to complete:** 15-20 minutes

## 🚀 Getting Started

### Prerequisites

1. Go 1.19 or later
2. Self SDK dependencies (automatically handled by go.mod)

### Running the Examples

Each example is a standalone Go program. Run them in order:

```bash
# 1. Basic Issuance (start here)
cd basic && go run main.go

# 2. Multi-Claim Issuance
cd ../multi_claim && go run main.go

# 3. Evidence-Based Issuance
cd ../evidence && go run main.go

# 4. Complex Data Issuance
cd ../complex && go run main.go

# 5. Advanced Issuance
cd ../advanced && go run main.go

# 6. Exchange Foundation Demo
cd ../exchange_demo && go run main.go

# Or run the simplified version in the root directory
cd .. && go run main.go
```

### 🔧 Build Requirements

Each subdirectory is a standalone Go module with its own `go.mod` file. The examples use the latest published version of the Self SDK (v0.59.0+), so they should build and run without any additional setup.

### What Each Example Does

All examples create two accounts (issuer and holder) using the core SDK and demonstrate different aspects of credential issuance:

- **Issuer**: Creates and signs credentials for subjects using account.CredentialIssue()
- **Holder**: Receives and stores credentials from issuers
- **Issuance**: The process of creating, signing, and delivering credentials using the underlying SDK

## 🎓 Learning Outcomes

After completing all examples, you'll understand:

### Core Concepts
- ✅ How credential issuance works between parties
- ✅ The difference between issuers, holders, and verifiers
- ✅ Credential builder patterns and best practices
- ✅ Claim structuring and organization

### Technical Skills
- ✅ Self SDK account setup and configuration using account.New()
- ✅ Direct credential creation using credential.NewCredential()
- ✅ Simple and complex claim addition with CredentialSubjectClaims()
- ✅ Evidence and asset management with object.New()
- ✅ Verifiable presentation creation
- ✅ Complex data structure modeling

### Production Readiness
- ✅ Error handling and validation
- ✅ Security considerations for credential issuance
- ✅ Scalable issuance patterns
- ✅ Asset management and storage

## 🔧 Key SDK Components Covered

### Account Management
- `account.New()` - Account initialization with core SDK
- `account.Config` - Direct account configuration options
- Storage and environment setup using core types

### Credential Operations
- `credential.NewCredential()` - Direct credential creation
- `CredentialType()`, `CredentialSubject()`, `Issuer()` - Core credential properties
- `CredentialSubjectClaims()` - Adding credential data directly
- `SignWith()`, `account.CredentialIssue()` - Direct credential signing and issuance

### Asset Management
- `CreateAsset()` - Evidence and file management
- Asset storage and retrieval
- Hash-based evidence linking

### Presentation Operations
- `CreatePresentation()` - Verifiable presentation creation
- Credential packaging for sharing
- Selective disclosure preparation

## 🛠️ Customization

Each example can be customized for your specific use case:

### Credential Types
- Modify credential types to match your domain
- Add custom claims relevant to your application
- Implement domain-specific validation logic

### Data Structures
- Customize claim structures for your data
- Implement complex nested data patterns
- Add business logic for data validation

### Evidence Integration
- Integrate with existing file storage systems
- Customize evidence types and formats
- Implement custom asset management

## 📚 Next Steps

After completing these examples:

1. **🔄 Explore credential exchange** in `../credentials_exchange/` - Learn how to request and share the credentials you've created
2. **📖 Review the Self SDK documentation** for advanced features
3. **🏗️ Build your own credential issuance application**
4. **🔗 Integrate with existing identity management systems**
5. **🎯 Combine issuance and exchange** - Create end-to-end credential workflows

### 🔄 Credential Lifecycle

Understanding the complete credential lifecycle:
- **Issuance** (this tutorial) - Creating and signing credentials
- **Exchange** (`../credentials_exchange/`) - Requesting and sharing credentials  
- **Verification** - Validating credential authenticity and claims
- **Revocation** - Managing credential lifecycle and updates

## 🤝 Contributing

If you find ways to improve these examples or have suggestions for additional educational content, please contribute back to the Self SDK project.

## 📖 Additional Resources

- [Self SDK Documentation](https://docs.joinself.com)
- [W3C Verifiable Credentials](https://w3.org/TR/vc-data-model/)
- [W3C Verifiable Presentations](https://w3.org/TR/vc-data-model/#presentations)
- [DIDComm Messaging](https://identity.foundation/didcomm-messaging/)
- [Decentralized Identity Foundation](https://identity.foundation)

---

**Happy learning! 🎉**

Start with `basic/main.go` and work your way through the progression. Each example builds upon the previous ones, providing a comprehensive understanding of credential issuance with the Self SDK.
