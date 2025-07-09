# 🎫 Credential Examples

> **📖 Learn the concepts first:** [Verifiable Credentials Concepts](../concepts/verifiable-credentials.md)

## Overview

*Content coming in Phase 2*

This page will provide comprehensive guidance for credential examples, including:
- Credential issuance patterns
- Exchange and verification workflows  
- Storage and management strategies
- Real-world application scenarios
- Integration with existing systems

## Example Categories

### 🟢 Basic Credential Issuance
Learn fundamental credential creation and signing.

**Examples:**
- [Basic Credential Issuance](../../examples/server/02_credentials/01_issuing_credentials/01_basic/)
- [Multi-Claim Credentials](../../examples/server/02_credentials/01_issuing_credentials/02_multi_claim/)

### 🟡 Advanced Credential Features
Master evidence attachment and complex data structures.

**Examples:**
- [Evidence-Based Credentials](../../examples/server/02_credentials/01_issuing_credentials/03_with_evidence/)
- [Complex Data Credentials](../../examples/server/02_credentials/01_issuing_credentials/04_complex_data/)
- [Comprehensive Credentials](../../examples/server/02_credentials/01_issuing_credentials/05_comprehensive/)

### 🟡 Credential Exchange
Understand presentation requests and verification workflows.

**Examples:**
- [Presentation Requests](../../examples/server/02_credentials/02_exchanging_credentials/presentation_request/)
- [Email Verification](../../examples/server/02_credentials/02_exchanging_credentials/email_verification/)

## Key Learning Outcomes

After completing these examples, you'll understand:
- W3C Verifiable Credentials standard implementation
- Issuer-Holder-Verifier relationship patterns
- Credential lifecycle management
- Proof mechanisms and verification
- Real-world credential use cases

## Credential Lifecycle

### Issuance Phase
- Identity verification
- Claim definition and validation
- Cryptographic signing
- Evidence attachment
- Credential delivery

### Storage Phase
- Secure credential storage
- Organization and categorization
- Search and retrieval
- Backup and recovery

### Exchange Phase
- Presentation requests
- Selective disclosure
- Verification workflows
- Trust establishment

## Use Case Patterns

### Identity Verification
- Email verification credentials
- Identity document verification
- Multi-factor authentication
- Account linking and validation

### Educational Credentials
- Course completion certificates
- Degree and diploma verification
- Skill attestations
- Professional certifications

### Professional Credentials
- Employment verification
- Role and responsibility attestation
- Performance evaluations
- Industry certifications

## Prerequisites

- Completed [Setup Examples](setup.md)
- Understanding of [Connection Examples](connections.md)
- Knowledge of [Verifiable Credentials Concepts](../concepts/verifiable-credentials.md)

## Next Steps

After mastering credentials, continue with:
- [Chat Examples](chat.md) - Credential-aware messaging
- [Advanced Examples](advanced.md) - Production credential systems 

# Credential Examples

> **Theory foundation:** [Verifiable Credentials Concepts](../concepts/verifiable-credentials.md)  
> **What you'll learn:** How to issue, exchange, and verify cryptographically secure credentials

Welcome to **verifiable credential implementation**! These examples transform the credential concepts you've learned into working code that creates, exchanges, and verifies digital credentials with cryptographic proof.

---

## What You'll Learn

By completing these credential examples, you'll master:

- **🏭 Credential Issuance**: How to create and cryptographically sign verifiable credentials
- **🔄 Credential Exchange**: How to request, present, and verify credential proofs
- **💾 Credential Storage**: How to organize and manage credential collections
- **🔐 Cryptographic Verification**: How to validate credential authenticity instantly
- **📱 Real-world Applications**: How to build complete credential workflows

**Time Investment**: 45 minutes to complete all credential patterns  
**Immediate Result**: Working credential system ready for production use

---

## The Big Picture: From Theory to Practice

In [Verifiable Credentials Concepts](../concepts/verifiable-credentials.md), you learned **why** traditional credentials are broken and **how** the Issuer-Holder-Verifier triangle works. Now you'll **build** complete credential workflows and experience cryptographic verification.

### Theory → Practice Connection

```
THEORY                           PRACTICE
Issuer-Holder-Verifier triangle → Implement all three roles with working code
Claims provide verifiable proof  → Create credentials with real claims and signatures
Cryptographic signatures work   → Verify credentials instantly without phone calls
Selective disclosure protects   → Build privacy-preserving credential presentations
```

---

## Complete Credential Journey

Follow this **progressive learning path** to master Self credential workflows:

### **Phase 1:** Master Credential Issuance
Learn to create and sign verifiable credentials with increasing complexity:

#### **Step 1:** Basic Credential Creation
**[Basic Credential Issuance](../../examples/server/02_credentials/01_issuing_credentials/01_basic/)**

**What it demonstrates:**
- Creating issuer and holder accounts
- Building credentials with simple claims
- Cryptographic signing and issuance process
- Foundation for all credential workflows

```go
// Create simple email verification credential
claims := map[string]interface{}{
    "emailAddress": "alice@example.com",
    "verified":     true,
    "domain":       "example.com",
}

credential := credential.NewCredential().
    CredentialType(credential.CredentialTypeEmail).
    CredentialSubject(credential.AddressKey(holderAddress)).
    CredentialSubjectClaims(claims).
    Issuer(credential.AddressKey(issuerAddress)).
    SignWith(issuerAddress, time.Now())

emailCredential, err := issuer.CredentialIssue(credential)
// ✅ Cryptographically signed credential created
// ✅ Cannot be forged or tampered with
// ✅ Instantly verifiable by anyone
```

**Key Concept**: Every credential contains cryptographic proof of the issuer's signature, making forgery impossible.

**Time**: 5 minutes to complete  
**Success**: Working email verification credential with cryptographic signature

---

#### **Step 2:** Multi-Claim Credentials
**[Multi-Claim Credentials](../../examples/server/02_credentials/01_issuing_credentials/02_multi_claim/)**

**What it demonstrates:**
- Grouping related claims in single credentials
- Different data types in claims (strings, booleans, numbers)
- Efficient credential structuring patterns
- Building comprehensive identity proofs

```go
// Profile credential with multiple related claims
profileClaims := map[string]interface{}{
    "firstName":        "John",
    "lastName":         "Doe", 
    "country":          "United States",
    "age":              30,
    "isActive":         true,
    "profileLevel":     "verified",
    "registrationDate": time.Now().Format("2006-01-02"),
}

// Education credential with academic information
educationClaims := map[string]interface{}{
    "institution":      "University of Technology",
    "degree":           "Bachelor of Science",
    "major":            "Computer Science",
    "graduationYear":   2020,
    "gpa":              3.8,
    "honors":           true,
    "creditsCompleted": 120,
}
```

**Key Concept**: Multiple related claims in one credential reduce verification overhead and improve user experience.

**Time**: 10 minutes to complete  
**Success**: Rich credentials containing multiple verified claims

---

#### **Step 3:** Evidence-Based Credentials
**[Credentials with Evidence](../../examples/server/02_credentials/01_issuing_credentials/03_with_evidence/)**

**What it demonstrates:**
- Attaching file evidence to credentials
- Secure asset management and upload
- Creating verifiable presentations
- Linking evidence to credential claims

```go
// Create evidence file (certificate, diploma, etc.)
certificateData := []byte("Certificate of Completion...")
evidence, err := object.New("certificate.pdf", certificateData)
err = issuer.ObjectUpload(evidence, true)

// Reference evidence in credential claims
claims := map[string]interface{}{
    "certificationType": "Professional Development",
    "courseName":        "Advanced Go Programming", 
    "completionDate":    time.Now().Format("2006-01-02"),
    "grade":             "A+",
    "evidenceId":        fmt.Sprintf("%x", evidence.Id()),
    "institution":       "Self SDK Academy",
}
// ✅ Credential linked to tamper-evident file evidence
// ✅ Recipients can verify both claims and supporting documents
```

**Key Concept**: Evidence files provide additional verification material while maintaining cryptographic integrity.

**Time**: 15 minutes to complete  
**Success**: Credentials with attached file evidence and presentations

---

#### **Step 4:** Complex Data Structures
**[Complex Data Credentials](../../examples/server/02_credentials/01_issuing_credentials/04_complex_data/)**

**What it demonstrates:**
- Nested objects and hierarchical data
- Arrays and collections in credentials
- Real-world organizational data modeling
- Advanced claim structuring techniques

```go
// Complex organizational credential with nested data
complexClaims := map[string]interface{}{
    "organizationName": "TechCorp Inc.",
    "employeeId":       "EMP-2024-001",
    "position": map[string]interface{}{
        "title":      "Senior Software Engineer",
        "department": "Engineering",
        "level":      "L5",
        "startDate":  "2024-01-15",
        "manager":    "jane.smith@techcorp.com",
    },
    "permissions": []string{
        "read:repositories",
        "write:code", 
        "deploy:staging",
        "review:pull-requests",
    },
    "certifications": []map[string]interface{}{
        {
            "name":       "AWS Solutions Architect",
            "level":      "Professional",
            "issueDate":  "2023-06-15",
            "verified":   true,
        },
    },
}
```

**Key Concept**: Complex nested structures handle real-world organizational data while maintaining cryptographic integrity.

**Time**: 20 minutes to complete  
**Success**: Enterprise-grade credentials with complex hierarchical data

---

#### **Step 5:** Comprehensive Credential System
**[Comprehensive Credentials](../../examples/server/02_credentials/01_issuing_credentials/05_comprehensive/)**

**What it demonstrates:**
- Complete credential issuance system
- All features combined in production workflow
- Multiple credential types and patterns
- End-to-end credential management

**Key Concept**: Production-ready credential issuance combining all learned patterns.

**Time**: 30 minutes to complete  
**Success**: Complete credential issuance system ready for deployment

---

### **Phase 2:** Master Credential Exchange
Learn to request, present, and verify credentials:

#### **Step 6:** Presentation Requests
**[Presentation Request Workflow](../../examples/server/02_credentials/02_exchanging_credentials/presentation_request/)**

**What it demonstrates:**
- Complete Issuer-Holder-Verifier workflow
- Credential storage and organization patterns
- Type-based credential matching
- Verifiable presentation creation

```go
// Verifier requests specific credential types
requestedTypes := []string{"EducationCredential", "StudentCredential"}

// Holder searches credential store for matches
var matchingCreds []*credential.VerifiableCredential
for credName, cred := range holder.credentials {
    for _, credType := range cred.CredentialType() {
        for _, requestedType := range requestedTypes {
            if credType == requestedType {
                matchingCreds = append(matchingCreds, cred)
                break
            }
        }
    }
}

// Create verifiable presentation with matching credentials
presentation := createCredentialPresentation(matchingCreds)
// ✅ Multiple credentials packaged for sharing
// ✅ Holder maintains control over disclosure
// ✅ Cryptographic integrity preserved
```

**Key Concept**: Holders search their credential stores and create presentations containing only requested proof types.

**Time**: 15 minutes to complete  
**Success**: Complete credential exchange workflow with request/response patterns

---

#### **Step 7:** Real-world Mobile Integration
**[Email Verification Workflow](../../examples/server/02_credentials/02_exchanging_credentials/email_verification/)**

**What it demonstrates:**
- Mobile app credential delivery
- QR code connection for credential exchange
- Service provider credential issuance
- Real-world verification scenarios

```go
// Generate QR code for mobile connection
keyPackage, err := emailService.ConnectionNegotiateOutOfBand(inboxAddress, expiration)
qrCode, err := anonymousMsg.EncodeToQR(event.QREncodingUnicode)
// Mobile app scans QR code to connect

// Handle mobile connection and issue credential
func handleMobileConnection(emailService *account.Account, welcome *event.Welcome) {
    // Create email verification credential for mobile user
    emailCredential := createEmailVerificationCredential(emailService, issuerAddress, mobileAddress)
    // ✅ Credential delivered directly to mobile device
    // ✅ User owns portable email verification proof
}
```

**Key Concept**: QR codes enable seamless credential delivery to mobile devices for real-world applications.

**Time**: 10 minutes to complete  
**Success**: Mobile-friendly credential delivery system

---

### **Phase 3:** Master Credential Storage
Learn to organize and manage credential collections:

#### **Step 8:** Credential Storage Patterns
**[Basic Credential Storage](../../examples/server/02_credentials/03_credential_storage/basic_storage/)**

**What it demonstrates:**
- Credential organization and categorization
- Search and retrieval patterns
- Storage optimization techniques
- Credential lifecycle management

**Key Concept**: Efficient credential storage enables fast search and presentation workflows.

**Time**: 10 minutes to complete  
**Success**: Organized credential storage system

---

## What Just Happened? Credential Revolution in Action

### **Credential Security Achieved**
You've just implemented the **credential revolution**:

- **🚫 No more forgery** - credentials contain cryptographic proof impossible to fake
- **⚡ Instant verification** - no phone calls to institutions or manual checking required
- **🔐 Tamper detection** - any modification of credential data is immediately detectable
- **👤 User control** - holders decide what information to share and when

### **Traditional vs Self Credentials: You Built Both**

| Traditional Credentials | Your Self Credentials |
|------------------------|----------------------|
| 📄 "PDF certificate" | 🔐 **Cryptographic proof** |
| 📞 "Call to verify" | ⚡ **Instant verification** |
| 📋 "Share everything" | 🎯 **Selective disclosure** |
| ❌ "Trust the paper" | ✅ **Mathematical guarantee** |

### **Real-World Impact Delivered**
Your credential system now provides:

- **🏛️ Institutional trust** - universities, employers, and agencies can issue unforgeable credentials
- **👤 Personal control** - individuals own and control their verified information
- **🔍 Business efficiency** - employers verify qualifications instantly without delays
- **🌐 Universal acceptance** - credentials work across all Self-enabled applications

---

## Technical Foundation Mastered

### **Architecture Understanding**
Through these examples, you've built this complete credential system:

```
Self Credential System
├── 🏭 Issuance Layer (create and sign credentials)
├── 💾 Storage Layer (organize and search credentials)  
├── 🔄 Exchange Layer (request and present credentials)
├── 🔐 Verification Layer (validate cryptographic proofs)
└── 📱 Integration Layer (mobile and API integration)
```

### **Credential Lifecycle Implemented**
- **Issuance**: Authorities create signed credentials for individuals
- **Storage**: Holders organize credentials in searchable collections
- **Presentation**: Holders share specific credentials in response to requests
- **Verification**: Verifiers validate cryptographic proofs instantly
- **Management**: Complete lifecycle from creation to expiration

### **Security Model Mastered**
- **Cryptographic Signatures**: Mathematical proof of issuer authenticity
- **Tamper Evidence**: Any modification breaks cryptographic integrity
- **Selective Disclosure**: Share only necessary information for each context
- **Decentralized Verification**: No need to contact original issuers

---

## Credential Pattern Selection Guide

### **When to Use Basic Credentials**
```go
// Perfect for:
emailClaims := map[string]interface{}{
    "emailAddress": "alice@example.com",
    "verified":     true,
}
// ✅ Simple identity verification
// ✅ Account registration workflows
// ✅ Basic trust establishment
// ✅ Development and testing
```

### **When to Use Multi-Claim Credentials**
```go
// Perfect for:
profileClaims := map[string]interface{}{
    "firstName": "Alice", "lastName": "Smith",
    "country": "US", "age": 30, "verified": true,
}
// ✅ Comprehensive identity profiles
// ✅ Reducing verification overhead
// ✅ Related information grouping
// ✅ User experience optimization
```

### **When to Use Evidence-Based Credentials**
```go
// Perfect for:
certificationClaims := map[string]interface{}{
    "certification": "AWS Solutions Architect",
    "evidenceId":    fmt.Sprintf("%x", evidenceFile.Id()),
}
// ✅ Academic diplomas and certificates
// ✅ Professional licensing
// ✅ Legal document verification
// ✅ High-assurance use cases
```

### **When to Use Complex Data Credentials**
```go
// Perfect for:
organizationalClaims := map[string]interface{}{
    "employee": map[string]interface{}{
        "id": "EMP-001", "department": "Engineering",
        "permissions": []string{"deploy", "review"},
    },
}
// ✅ Enterprise employee credentials
// ✅ Complex organizational hierarchies
// ✅ Multi-faceted professional profiles
// ✅ Government and institutional use cases
```

---

## Complete Credential Testing

### **End-to-End Credential Workflow**
Experience the complete Issuer-Holder-Verifier triangle:

```bash
# Phase 1: Master credential issuance
cd examples/server/02_credentials/01_issuing_credentials/01_basic/go
go run main.go
# Creates issuer and holder accounts, issues email credential

# Phase 2: Test credential exchange  
cd ../../02_exchanging_credentials/presentation_request/go
go run main.go
# Demonstrates credential storage, search, and presentation

# Phase 3: Try mobile integration
cd ../email_verification/go
go run main.go
# Generates QR code for mobile credential delivery
```

### **Progressive Complexity Testing**
Experience the learning progression:

```bash
# Basic → Multi-claim → Evidence → Complex → Comprehensive
cd 01_basic/go && go run main.go
cd ../02_multi_claim/go && go run main.go  
cd ../03_with_evidence/go && go run main.go
cd ../04_complex_data/go && go run main.go
cd ../05_comprehensive/go && go run main.go
```

---

## Next Steps: Build On Your Credential Foundation

With verifiable credentials mastered, you're ready for advanced applications:

### **Level 1: Secure Messaging** 🟡
- **[Chat Examples](chat.md)** - Send credential-aware messages over secure connections
- **Credential Context** - Include verified identity information in messages

### **Level 2: Advanced Features** 🟠
- **[Advanced Examples](advanced.md)** - Production credential systems and optimization
- **Multi-device Sync** - Credential portability across devices
- **Enterprise Integration** - Large-scale credential management


---

## 🏭 Production Deployment

**Ready for production credentials?** Check our comprehensive **[Production Deployment Guide](production.md)** for everything you need to deploy credential systems at scale:

- 🚀 **[Moving to Production](production.md#moving-to-production)** - Issuer key management, credential revocation workflows, and performance optimization
- 🔐 **[Security Hardening](production.md#security-hardening)** - Issuer validation, credential filtering, access control, and audit trails
- ⚡ **[Performance Optimization](production.md#performance-optimization)** - Credential indexing, batch operations, storage optimization, and caching strategies
- 📊 **[Monitoring & Observability](production.md#monitoring--observability)** - Issuance analytics, verification tracking, and usage monitoring
- 📈 **[Scalability Patterns](production.md#scalability-patterns)** - Large-scale credential collections, distributed storage, and high-throughput processing

The production guide includes credential-specific security patterns, enterprise deployment strategies, and scalability optimization techniques.

---

## Success Checklist

Confirm you've mastered verifiable credentials:

**✅ Credential Issuance**
- [ ] Can create credentials with various claim types and complexity levels
- [ ] Understand cryptographic signing and verification processes
- [ ] Can attach evidence files and create verifiable presentations
- [ ] Master the complete credential builder pattern

**✅ Credential Exchange**  
- [ ] Know how to implement request/response workflows
- [ ] Understand credential search and matching algorithms
- [ ] Can create verifiable presentations for selective disclosure
- [ ] Handle mobile credential delivery scenarios

**✅ Credential Storage**
- [ ] Can organize credentials in searchable collections
- [ ] Understand credential lifecycle management
- [ ] Know how to optimize storage and retrieval performance
- [ ] Master credential backup and recovery patterns

**✅ Real-World Applications**
- [ ] Can choose appropriate credential patterns for use cases
- [ ] Understand privacy implications and selective disclosure
- [ ] Know how to integrate credentials with existing systems
- [ ] Ready to build production credential applications

---

## 🔧 Need Help?

**Having credential issues?** Check our comprehensive **[Troubleshooting Guide](troubleshooting.md)** for solutions to credential problems, including:

- 🎫 **[Credential Issues](troubleshooting.md#credential-issues)** - Issuance failures, evidence upload problems, verification errors, and presentation creation issues
- 🏗️ **[Setup Issues](troubleshooting.md#setup--account-issues)** - Account initialization for issuers and holders
- 🌐 **[Network Issues](troubleshooting.md#network--connectivity-issues)** - Connectivity problems affecting credential operations

The troubleshooting guide includes detailed solutions, common causes, and debugging tips for all Self SDK examples.

---

## 📚 Resources & Next Steps

**Building credential systems?** Check our comprehensive **[Resources & Community Guide](resources.md)** for everything you need to work with verifiable credentials:

- 📋 **[Standards & Specifications](resources.md#standards--specifications)** - W3C Verifiable Credentials, JSON-LD, and DID Core specifications
- 🛠️ **[Developer Tools](resources.md#developer-tools)** - Credential validation, testing utilities, and mock issuers/holders for development
- 📖 **[Related Concepts](resources.md#related-concepts)** - Cryptographic foundations and secure connection patterns for credentials
- 👥 **[Community Support](resources.md#community-support)** - Get help with credential implementation and connect with other developers

The resources guide includes complete documentation for credential standards, development tools, and community support.

---

**Congratulations!** You've mastered verifiable credentials and built a complete digital trust system. Your credentials now provide cryptographic proof, instant verification, and user privacy - transforming how digital identity works.

**Ready to add secure messaging?** Continue with [Chat Examples](chat.md) to build credential-aware communication over your secure connections! 🚀
