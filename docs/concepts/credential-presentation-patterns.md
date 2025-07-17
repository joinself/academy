# Credential Presentation Patterns

> **🛠️ Practical Implementation:** [Credential Examples](../examples/credentials.md#phase-25-master-presentation-request-patterns)  
> **📖 Learn the concepts first:** [Verifiable Credentials Concepts](verifiable-credentials.md)

## 🎯 What You'll Learn

- The two fundamental patterns for requesting credential data
- When to use simple retrieval vs zero-knowledge verification
- How privacy-preserving credential requests work mathematically
- **Working code** to implement both patterns in your applications
- Real-world decision making for credential request design

**Time Investment**: 10 minutes to understand patterns  
**Immediate Result**: Clear strategy for designing credential requests

---

## 🔑 The Two Fundamental Patterns

### **Pattern 1: Simple Retrieval** (`nil` parameters)

**Mathematical Model:**
```
Verifier → "Please share your email credential" → Holder
Holder → Returns complete credential with all claims → Verifier
Result: Verifier receives full credential data
```

**Implementation:**
```go
// Request complete credential data
content, err := message.NewCredentialPresentationRequest().
    Type([]string{"VerifiablePresentation"}).
    Details([]string{"EmailCredential"}, nil).  // ← nil = full disclosure
    Finish()

// Response contains complete credential:
// {
//   "emailAddress": "alice@company.com",
//   "verified": true,
//   "domain": "company.com",
//   "issuer": "did:self:email_service",
//   "issuanceDate": "2024-01-15"
// }
```

**Use Cases:**
- User account registration (need actual email)
- Professional verification (need specific qualifications)
- System integration (need structured data)
- Identity establishment (need verifiable information)

---

### **Pattern 2: Zero-Knowledge Verification** (with parameters)

**Mathematical Model:**
```
Verifier → "Prove condition X without revealing data Y" → Holder
Holder's Device → Evaluates condition locally → Returns boolean only
Result: Verifier gets proof without seeing sensitive data
```

**Implementation:**
```go
// Request age verification without birth date
eighteenYearsAgo := time.Now().AddDate(-18, 0, 0).Format("2006-01-02")

content, err := message.NewCredentialPresentationRequest().
    Type([]string{"VerifiablePresentation"}).
    Details(
        []string{"DateOfBirthCredential"},
        []*message.CredentialPresentationDetailParameter{
            message.NewCredentialPresentationDetailParameter(
                message.OperatorLowerThan,
                "dateOfBirth",
                eighteenYearsAgo,
            ),
        },
    ).
    Finish()

// Response contains only evaluation result:
// {
//   "result": true,
//   "operator": "LowerThan",
//   "field": "dateOfBirth"
// }
// Birth date never leaves user's device!
```

**Use Cases:**
- Age verification for restricted content
- Eligibility checks without data retention  
- Privacy-first verification scenarios
- Regulatory compliance (data minimization)

---

## 🔐 Privacy-Security Spectrum

The two patterns represent different points on the privacy-security spectrum:

```
Full Disclosure ←→ Zero-Knowledge
     ↓                    ↓
Simple Retrieval    Filtered Requests
(Details(type, nil)) (Details(type, [params]))
     ↓                    ↓
High Data Utility   Maximum Privacy
```

### **Privacy Analysis**

| Aspect | Simple Retrieval | Zero-Knowledge |
|--------|------------------|----------------|
| **Data Exposure** | Full credential content | Minimal (boolean results) |
| **Verifier Storage** | Can store all claims | Only stores verification results |
| **User Privacy** | Standard | Maximum |
| **Regulatory Compliance** | Standard | GDPR/data minimization friendly |
| **Integration Complexity** | Simple (direct data use) | Moderate (result interpretation) |

---

## 🔬 Cryptographic Foundations

### **How Zero-Knowledge Verification Works**

1. **Local Evaluation**: User's device evaluates the condition locally
2. **Cryptographic Proof**: Creates mathematical proof that evaluation is correct
3. **Result Transmission**: Only the boolean result (and proof) is sent
4. **Verification**: Verifier confirms the proof without seeing original data

**Mathematical Properties:**
- **Completeness**: If the statement is true, an honest prover can convince the verifier
- **Soundness**: If the statement is false, no dishonest prover can convince the verifier  
- **Zero-Knowledge**: The verifier learns nothing beyond the truth of the statement

### **Operator Types Available**

```go
// Comparison operators for numerical/date fields
message.OperatorEquals        // field == value
message.OperatorNotEquals     // field != value  
message.OperatorGreaterThan   // field > value
message.OperatorLowerThan     // field < value

// Existence operators
message.OperatorNotEquals     // Check field exists (use empty value)
```

---

## 🎯 Pattern Selection Decision Tree

```
Do you need the actual data?
├─ YES → Use Simple Retrieval (nil)
│   ├─ Account setup → Need email address
│   ├─ Profile display → Need name, photo
│   └─ Integration → Need structured data
│
└─ NO → Use Zero-Knowledge (params)
    ├─ Age verification → Just need "over 18"
    ├─ Eligibility → Just need "qualified" 
    └─ Privacy compliance → Minimize data
```

### **Real-World Decision Examples**

**E-commerce Platform:**
- **Login**: Simple retrieval for email (need actual address)
- **Age-restricted products**: Zero-knowledge for age (just need eligibility)

**Financial Services:**
- **Account opening**: Simple retrieval for identity data (KYC requirements)
- **Credit eligibility**: Zero-knowledge for income verification (privacy + compliance)

**Healthcare Platform:**
- **Registration**: Simple retrieval for insurance details (need actual data)
- **Treatment eligibility**: Zero-knowledge for condition verification (privacy protection)

---

## 🚀 Try It Yourself

### **Immediate Testing**

Test both patterns with these working examples:

```bash
# Simple retrieval pattern
cd examples/server/02_credentials/02_exchanging_credentials/email_verification/go
# Edit main.go to use Details(credential.CredentialTypeEmail, nil)
go run main.go

# Zero-knowledge pattern  
cd ../../solutions/age-verifier/
# Examine how age verification uses operators
go run cmd/server/main.go
```

### **Pattern Comparison Exercise**

Implement the same verification scenario using both patterns:

```go
// Pattern 1: Get actual email for account setup
emailRequest := message.NewCredentialPresentationRequest().
    Details(credential.CredentialTypeEmail, nil)

// Pattern 2: Just verify email exists (privacy-first)
emailExistsRequest := message.NewCredentialPresentationRequest().
    Details(credential.CredentialTypeEmail, []*message.CredentialPresentationDetailParameter{
        message.NewCredentialPresentationDetailParameter(
            message.OperatorNotEquals,
            "emailAddress",
            "",  // Empty value checks for existence
        ),
    })
```

---

## 📚 What Just Happened?

You've learned the fundamental choice in credential verification:

### **🔓 Simple Retrieval Pattern**
- **When**: You need actual data for application functionality
- **Privacy**: Standard data sharing model
- **Result**: Full credential content for integration

### **🔐 Zero-Knowledge Pattern**  
- **When**: You need proof without data exposure
- **Privacy**: Maximum privacy preservation
- **Result**: Boolean verification results only

### **🎯 Strategic Decision**
The pattern you choose determines:
- User privacy level
- Data retention obligations  
- Integration complexity
- Regulatory compliance approach

---

## 🚀 Next Steps

**Ready to implement credential requests?** Continue with:

- **[Credential Examples](../examples/credentials.md#phase-25-master-presentation-request-patterns)** - Implement both patterns in working code
- **[Age Verification Solution](../../examples/solutions/age-verifier/)** - See zero-knowledge pattern in production
- **[Email Verification Example](../../examples/server/02_credentials/02_exchanging_credentials/email_verification/)** - See simple retrieval pattern in action

**Building production systems?** See:
- **[Production Deployment](../examples/production.md)** - Scale credential verification systems
- **[Security Model](../architecture/security-model.md)** - Threat analysis for credential applications

---

**Congratulations!** You now understand the fundamental choice in credential verification design. Choose simple retrieval when you need data, zero-knowledge when you need privacy. Both patterns provide cryptographic security - the difference is in data disclosure strategy. 
