# Exchange Demo - From Issuance to Exchange

This example demonstrates how credential issuance creates the foundation for credential exchange workflows.

## 🎯 Purpose

This demo bridges the gap between credential issuance and credential exchange by showing:

- How issued credentials become exchangeable digital assets
- Credential storage and organization patterns for holders
- Conceptual workflow of credential presentation requests
- Foundation patterns for building real exchange applications

## 📚 Prerequisites

Complete the basic credential issuance examples first to understand the foundation:

1. `../basic/main.go` - Foundation credential creation
2. `../multi_claim/main.go` - Multiple claims in credentials

## 🔄 Exchange Workflow Demonstration

### Scenario: University to Employer Exchange

1. **University Issues Credentials** (Using issuance patterns)
   - Email credential (institutional email verification)
   - Student ID credential (enrollment status) 
   - Degree credential (academic achievement)

2. **Holder Organizes Credentials** (Credential store management)
   - Credentials stored with searchable metadata
   - Categorized by type and issuer
   - Ready for exchange scenarios

3. **Employer Requests Proof** (Exchange simulation)
   - Requests specific credential types
   - Holder searches credential store
   - Matching credentials identified

4. **Presentation Creation** (Packaging for sharing)
   - Multiple credentials bundled into presentation
   - Cryptographic integrity maintained
   - Ready for verification

## 🏗️ Key Components

### ExchangeParty Structure
```go
type ExchangeParty struct {
    name        string
    account     *account.Account
    credentials map[string]*credential.VerifiableCredential
}
```

### Credential Types Demonstrated
- **Email Credential**: Institutional email verification
- **Student Credential**: Academic enrollment status
- **Degree Credential**: Educational achievement

### Exchange Process Steps
1. **Credential Request**: Employer specifies needed credential types
2. **Store Search**: Holder searches their credential collection
3. **Match Filtering**: Credentials matched against request criteria
4. **Presentation Building**: Selected credentials packaged together
5. **Trust Verification**: Cryptographic signatures enable verification

## 🎓 Educational Value

### What You'll Learn
- **Foundation Connection**: How issuance enables exchange
- **Credential Management**: Storage and organization patterns
- **Request Matching**: Filtering credentials by type and criteria
- **Presentation Packaging**: Bundling multiple credentials
- **Trust Architecture**: Cryptographic verification without intermediaries

### Production Considerations
- **Discovery Mechanisms**: QR codes, deep links, and protocol handlers
- **Secure Messaging**: Encrypted request/response communication
- **Selective Disclosure**: Privacy-preserving data sharing
- **Advanced Filtering**: Complex matching criteria and constraints
- **Zero-Knowledge Proofs**: Privacy-preserving verification patterns

## 🔧 Running the Demo

```bash
cd exchange_demo
go run main.go
```

### Expected Output
```
🔄 Credential Exchange Demo - From Issuance to Exchange
========================================================
This demo shows how issued credentials enable exchange workflows.
📚 Building on: Basic credential issuance patterns

🔧 Setting up exchange parties...
🏢 University Issuer: [DID]
👤 Student Holder: [DID]
✅ Exchange parties ready

📋 Issuing credentials for exchange scenarios...
   (Using patterns from credential issuance examples)
   ✅ Email credential issued to holder
   ✅ Student ID credential issued to holder
   ✅ Degree credential issued to holder
📦 Issuance complete - holder has 3 credentials for exchange

🔄 Demonstrating exchange workflow...
   Scenario: Employer requests proof of education

📤 STEP 1: Employer requests education credentials
   Request: Credentials of types [StudentCredential DegreeCredential]

📨 STEP 2: Holder searches credential store
   ✅ Found matching credential: student_id (StudentCredential)
   ✅ Found matching credential: degree (DegreeCredential)

📜 STEP 3: Creating verifiable presentation
   ✅ Presentation created with 2 credentials
      📋 Credential 1: [StudentCredential]
      📋 Credential 2: [DegreeCredential]

🎉 STEP 4: Exchange completed successfully!
   ✅ Presentation would be shared with employer
   ✅ Employer verifies credentials cryptographically
   ✅ Trust established through verifiable credentials
```

## 🌐 Real-World Applications

### Use Cases Demonstrated
- **Academic Verification**: University credentials for employment
- **Professional Licensing**: Qualifications for job applications
- **Identity Verification**: Multiple identity factors bundled
- **Compliance Reporting**: Regulatory credential requirements

### Architecture Patterns
- **Credential Store**: Local credential management
- **Request Processing**: Automated credential matching
- **Presentation Assembly**: Multi-credential packaging
- **Verification Workflows**: Cryptographic trust chains

## 🔗 Next Steps

After understanding this exchange foundation:

1. **Complex Structures**: Explore `../complex/main.go` for nested credential data
2. **Advanced Features**: Try `../advanced/main.go` for comprehensive patterns
3. **Real Exchange**: Consider implementing full discovery and messaging
4. **Production Features**: Add selective disclosure, ZK proofs, and privacy

## 🎨 Design Philosophy

This demo demonstrates that **credential exchange is an extension of credential issuance**:

- **Issuance creates the assets** - credentials with cryptographic integrity
- **Exchange organizes and shares the assets** - discovery, matching, and presentation
- **Verification validates the assets** - cryptographic signature verification

The foundation is solid credential issuance patterns, which then enable sophisticated exchange workflows in real applications.

## 📖 Technical Notes

### SDK Version
- Requires Self SDK v0.59.0 or later
- Uses core SDK directly (no client facade)
- Direct account management patterns

### Dependencies
```go
require github.com/joinself/self-go-sdk v0.59.0
```

### Storage
- Creates demo storage directories
- Temporary accounts for demonstration
- Clean up after testing if desired

### Credential Types
- Uses standard Self SDK credential types
- Custom types can be defined for specific use cases
- Type matching enables flexible exchange scenarios 
