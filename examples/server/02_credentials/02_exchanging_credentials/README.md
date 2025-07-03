# Credential Exchange Examples 🔄

> **📖 Learn the concepts first:** [Verifiable Credentials Concepts](../../../../../docs/concepts/verifiable-credentials.md)

Transform issued credentials into exchangeable digital assets through verifiable presentations and request/response workflows.

## 🎯 Learning Objectives

By completing these examples, you'll understand:
- **Exchange Foundation**: How credential issuance enables exchange workflows
- **Request Processing**: Matching credentials against verification requirements
- **Presentation Creation**: Packaging multiple credentials for sharing
- **Real-world Applications**: Employment, education, and identity verification scenarios

## 📚 Prerequisites

Complete credential issuance fundamentals first:
- `../01_issuing_credentials/01_basic/` - Foundation credential creation
- `../01_issuing_credentials/02_multi_claim/` - Multiple claims patterns
- `../01_issuing_credentials/03_with_evidence/` - Evidence-based credentials

## 🏗️ Core Concepts

### The Exchange Process
```
Issuance → Storage → Request → Matching → Presentation → Verification
```

### Key Components
- **Credential Store**: Searchable collection of issued credentials
- **Exchange Requests**: Type-based credential requirements
- **Presentation Assembly**: Multi-credential packaging for sharing
- **Verification Workflow**: Cryptographic trust validation

### Architecture Pattern
```
┌─────────────┐    Request    ┌─────────────┐    Presentation    ┌─────────────┐
│   Verifier  │ ──────────── │   Holder    │ ──────────────── │   Issuer    │
│ (Employer)  │              │ (Employee)  │                  │ (University)│
└─────────────┘              └─────────────┘                  └─────────────┘
      ↑                            ↑                                 ↓
      │                     Searches Store                    Originally Issued
      │                            ↓                                 ↓
      └─── Verifies ←── Presentation Created ←────────── Stored Credentials
```

## 📖 Examples Overview

| Example | Complexity | Focus Area | Key Learning |
|---------|------------|------------|--------------|
| **`presentation_request/`** | 🟢 Foundation | Basic exchange patterns | Request/response workflows |
| **`email_verification/`** | 🟡 Real-world | Email credential verification | Mobile credential delivery |

## 🚀 Quick Start

```bash
# Navigate to exchange examples
cd examples/go/02_credentials/02_exchanging_credentials

# Run the foundation example
cd presentation_request && go run main.go

# Run the real-world example
cd ../email_verification && go run main.go
```

## 📋 Example Details

### 🟢 `presentation_request/` - Foundation Exchange Patterns
**Scenario**: University graduate applying for employment

**What you'll build:**
- Multi-party credential exchange system
- Type-based credential matching
- Verifiable presentation creation
- Educational credential workflow

**Key components:**
```go
type ExchangeParty struct {
    name        string
    account     *account.Account
    credentials map[string]*credential.VerifiableCredential
}
```

**Workflow demonstrated:**
1. **Setup**: Create issuer (university) and holder (student) identities
2. **Issuance**: Issue email, student ID, and degree credentials
3. **Storage**: Organize credentials in searchable store
4. **Request**: Employer requests education proof
5. **Matching**: Find credentials matching requested types
6. **Presentation**: Package matching credentials for sharing

**Expected output:**
```
Credential Exchange Demo
========================
Setting up exchange parties...
Issuer: 1:ABC123...
Holder: 1:DEF456...

Issuing credentials...
✅ Email credential issued
✅ Student ID credential issued  
✅ Degree credential issued
Holder has 3 credentials

Exchange workflow:
Scenario: Employer requests proof of education
Requested types: [StudentCredential DegreeCredential]
✅ Found matching: student_id (StudentCredential)
✅ Found matching: degree (DegreeCredential)
✅ Presentation created with 2 credentials
✅ Exchange completed successfully
```

### 🟡 `email_verification/` - Real-world Mobile Exchange
**Scenario**: Email service provider issues verification credentials to users

**What you'll build:**
- Email verification credential workflow
- Service-to-user credential issuance
- Mobile credential delivery concepts
- Real-world verification scenarios

**Key components:**
```go
type VerificationParty struct {
    name    string
    account *account.Account
}
```

**Workflow demonstrated:**
1. **Service Setup**: Create email service provider and user identities
2. **Email Verification**: User completes email verification process
3. **Credential Issuance**: Service creates verifiable email credential
4. **Mobile Delivery**: Conceptual delivery to user's mobile device
5. **Portable Proof**: User owns permanent verification credential

**Expected output:**
```
Email Credential Verification Demo
=================================
Setting up email service provider...
✅ Email service provider ready
📬 Email Service Provider Address: 1:ABC123...
Generating QR code for mobile email verification...

[QR CODE DISPLAYED HERE]

⏱️  Expires: 09:16:14

📱 SCAN QR CODE with Self mobile app to verify email
⏳ Waiting for mobile device connection...
🔐 Once connected, email verification credential will be created

[After mobile device scans QR code:]

📱 Mobile device connected: 1:DEF456...
✅ Mobile connection established!

📧 Creating email verification credential...
✅ Email verification credential created successfully!
📱 Credential details:
   • Email: user@example.com
   • Status: verified
   • Domain: example.com
   • Method: email_link_clicked
   • Issuer: Example Email Service Provider

🎉 Email verification workflow completed!
💡 In production, this credential would be:
   • Sent directly to the mobile device
   • Stored in the user's credential wallet
   • Available for proving email ownership
   • Reusable across different services
```

**Real-world applications:**
- **Account Registration**: Streamlined user onboarding
- **KYC Compliance**: Financial services verification
- **Student Verification**: Academic email credentials
- **Enterprise Access**: Corporate email validation

## 🌐 Real-World Applications

### Employment Verification
- **Traditional**: Phone calls, email verification, delayed hiring
- **With Credentials**: Instant verification, tamper-evident proof, privacy-preserving

### Academic Verification
- **Traditional**: Transcript requests, manual validation, fraud risk
- **With Credentials**: Cryptographic proof, real-time verification, cost reduction

### Identity Verification
- **Traditional**: Multiple documents, manual checking, error-prone
- **With Credentials**: Consolidated proof, automated validation, enhanced security

## 🏗️ Exchange Architecture Patterns

### 1. Store-Based Exchange
```go
// Credential organization
holder.credentials["email"] = emailCredential
holder.credentials["student_id"] = studentCredential
holder.credentials["degree"] = degreeCredential

// Type-based search
for credName, cred := range holder.credentials {
    if matchesRequestedType(cred, requestedTypes) {
        matches = append(matches, cred)
    }
}
```

### 2. Request Processing Pipeline
```go
type ExchangeRequest struct {
    RequestedTypes []string
    Purpose        string
    Requester      *account.Account
}

// Process request → search store → create presentation
func processExchange(holder *ExchangeParty, request ExchangeRequest) (*Presentation, error) {
    matches := searchCredentials(holder.credentials, request.RequestedTypes)
    return createPresentation(matches), nil
}
```

### 3. Multi-Credential Presentations
```go
// Bundle multiple credentials
presentation := credential.NewPresentation().
    PresentationType([]string{"VerifiablePresentation"}).
    Holder(holderAddress)

for _, cred := range matchingCredentials {
    presentation.CredentialAdd(cred)
}
```

## 🔧 Technical Patterns

### Credential Matching Logic
```go
// Flexible type matching
func matchesRequestedType(cred *credential.VerifiableCredential, requestedTypes []string) bool {
    credTypes := cred.CredentialType()
    for _, credType := range credTypes {
        for _, requestedType := range requestedTypes {
            if credType == requestedType {
                return true
            }
        }
    }
    return false
}
```

### Error Handling
```go
// Graceful degradation
if len(matchingCreds) == 0 {
    return &ExchangeResponse{
        Status: "partial",
        Message: "Some requested credentials not available",
        AvailableTypes: getAvailableTypes(holder.credentials),
    }
}
```

### Security Considerations
```go
// Validate requests
func validateExchangeRequest(req ExchangeRequest) error {
    if len(req.RequestedTypes) == 0 {
        return errors.New("no credential types requested")
    }
    
    if !isAuthorizedRequester(req.Requester) {
        return errors.New("unauthorized requester")
    }
    
    return nil
}
```

## 🚀 Production Considerations

### Performance Optimization
- **Indexed Storage**: Fast credential lookup by type, issuer, date
- **Caching**: Frequently accessed credentials and presentations
- **Lazy Loading**: Load credential details only when needed
- **Batch Processing**: Handle multiple exchange requests efficiently

### Privacy Protection
- **Selective Disclosure**: Share only requested claims
- **Zero-Knowledge Proofs**: Prove properties without revealing data
- **Consent Management**: User control over credential sharing
- **Data Minimization**: Only collect necessary information

### Security Best Practices
- **Credential Encryption**: Protect stored credentials
- **Audit Logging**: Track all exchange activities
- **Request Validation**: Prevent malicious or malformed requests
- **Rate Limiting**: Protect against abuse

### Integration Patterns
- **REST APIs**: HTTP-based exchange endpoints
- **WebSockets**: Real-time exchange protocols
- **QR Codes**: Mobile-friendly discovery mechanisms
- **Deep Links**: Seamless app-to-app exchange

## 🔍 Advanced Topics

### Coming Soon: Enhanced Examples

#### `verification_workflow/` - End-to-end Verification
- Complete issuer-holder-verifier workflows
- Cryptographic signature validation
- Revocation status checking
- Trust framework integration

#### `selective_disclosure/` - Privacy-Preserving Exchange
- Zero-knowledge proof integration
- Minimal data sharing patterns
- Consent-based credential sharing
- Privacy-first architectures

#### `multi_party_exchange/` - Complex Scenarios
- Multi-issuer credential combinations
- Conditional credential requirements
- Workflow orchestration patterns
- Enterprise integration examples

### Research Areas
- **Credential Composition**: Combining multiple credentials into complex proofs
- **Conditional Logic**: "If degree AND experience AND security clearance"
- **Reputation Systems**: Issuer trust scoring and verification
- **Revocation Management**: Real-time credential validity checking

## 🧪 Testing and Validation

### Unit Testing Patterns
```go
func TestCredentialMatching(t *testing.T) {
    testCreds := createTestCredentials()
    requested := []string{"StudentCredential"}
    matches := findMatching(testCreds, requested)
    
    assert.Equal(t, 1, len(matches))
    assert.Contains(t, matches[0].CredentialType(), "StudentCredential")
}
```

### Integration Testing
```go
func TestFullExchangeWorkflow(t *testing.T) {
    issuer, holder := createTestParties()
    issueTestCredentials(issuer, holder)
    
    request := ExchangeRequest{Types: []string{"DegreeCredential"}}
    presentation, err := processExchange(holder, request)
    
    assert.NoError(t, err)
    assert.NotNil(t, presentation)
}
```

## 🎓 Learning Progression

### After Completing These Examples
1. **Understand Exchange Fundamentals**: How issuance enables exchange
2. **Master Storage Patterns**: Credential organization and search
3. **Implement Request Processing**: Type-based matching and filtering
4. **Create Presentations**: Multi-credential packaging for sharing

### Next Steps
- **Complex Data Structures**: Advanced credential schemas
- **Real-time Messaging**: WebSocket-based exchange protocols
- **Mobile Integration**: QR codes and deep link discovery
- **Production Deployment**: Scalable exchange architectures

## 🔗 Related Examples

### Prerequisites (Complete First)
- `../01_issuing_credentials/` - Foundation credential creation patterns
- `../../01_connection/` - Basic connection and messaging
- `../../00_setup/` - Account creation and configuration

### Advanced Topics (Explore Next)
- `../../03_chat/` - Real-time messaging integration
- `../../04_advanced_features/` - Production-ready patterns
- `../../07_issue_email/` - Practical credential workflows

## 📖 Additional Resources

### Key Documentation
- **Self SDK Documentation**: Complete API reference
- **W3C Standards**: Verifiable Credentials and Presentations
- **DID Specification**: Decentralized identifier standards

### Best Practices
- **Security Guidelines**: Protecting credentials and exchanges
- **Privacy Patterns**: Selective disclosure and consent management
- **Performance Optimization**: Scaling exchange systems

## 🎯 Summary

Credential exchange transforms static issued credentials into dynamic, interactive digital assets. These examples teach you:

### ✅ **Foundation Skills**
- **Exchange Architecture**: Understanding the issuer-holder-verifier model
- **Request Processing**: Matching credentials against requirements
- **Presentation Creation**: Packaging multiple credentials for sharing
- **Verification Workflows**: Cryptographic trust validation

### 🏗️ **Production Patterns**
- **Store Management**: Efficient credential organization and search
- **Security Best Practices**: Protecting exchanges and user privacy
- **Integration Approaches**: REST APIs, WebSockets, mobile applications
- **Performance Optimization**: Caching, indexing, and batch processing

### 🚀 **Real-world Applications**
- **Employment Verification**: Instant, tamper-evident credential checking
- **Academic Verification**: Fraud-resistant education proof
- **Identity Verification**: Multi-factor, privacy-preserving authentication
- **Compliance Reporting**: Automated regulatory credential validation

Master these patterns to build sophisticated credential exchange systems that provide instant verification, enhanced privacy, and reduced costs compared to traditional methods.
