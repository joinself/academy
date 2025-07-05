# Credential Exchange - Presentation Requests

> **📖 Learn the concepts first:** [Verifiable Credentials Concepts](../../../../../../docs/concepts/verifiable-credentials.md)

🎯 **What you'll learn:**
- How credential issuance enables exchange workflows
- Credential storage and organization patterns
- Request matching and filtering mechanisms
- Verifiable presentation creation
- Foundation concepts for real exchange applications

📚 **Prerequisites:** Complete credential issuance examples first:
- `../../01_issuing_credentials/01_basic/` - Foundation credential creation
- `../../01_issuing_credentials/02_multi_claim/` - Multiple claims concepts

## Overview

This example demonstrates the bridge between credential issuance and credential exchange. You'll learn how issued credentials become exchangeable digital assets and understand the fundamental patterns for building real exchange applications.

## Key Concepts

### Exchange Workflow Foundation
The exchange process builds directly on credential issuance:
1. **Issuance creates assets** - Credentials with cryptographic integrity
2. **Storage organizes assets** - Searchable credential collections
3. **Exchange shares assets** - Request matching and presentation
4. **Verification validates assets** - Cryptographic signature verification

### ExchangeParty Structure
```go
type ExchangeParty struct {
    name        string
    account     *account.Account
    credentials map[string]*credential.VerifiableCredential
}
```

**Components:**
- **name**: Human-readable identifier for the party
- **account**: Self SDK account with cryptographic identity
- **credentials**: Local credential store for search and matching

## Code Breakdown

### 1. Exchange Party Setup
```go
func createExchangeParties() (*ExchangeParty, *ExchangeParty) {
    issuer := &ExchangeParty{
        name:        "University Issuer",
        credentials: make(map[string]*credential.VerifiableCredential),
    }
    
    issuerCfg := &account.Config{
        StorageKey:  generateStorageKey("demo_issuer"),
        StoragePath: "./demo_issuer_storage",
        Environment: account.TargetSandbox,
        LogLevel:    account.LogWarn,
    }
    
    issuerAccount, err := account.New(issuerCfg)
    issuer.account = issuerAccount
    
    // Similar setup for holder...
}
```

**What happens:**
- Creates separate identities for issuer and holder
- Establishes encrypted storage for each party
- Initializes credential stores for organization
- Generates unique DIDs for cryptographic identity

### 2. Credential Issuance for Exchange
```go
func issueCredentialsForExchange(issuer, holder *ExchangeParty) {
    // Issue email credential
    emailCred := createEmailCredential(issuer.account, issuerAddress, holderAddress)
    if emailCred != nil {
        holder.credentials["email"] = emailCred
    }
    
    // Issue student credential
    studentCred := createStudentCredential(issuer.account, issuerAddress, holderAddress)
    if studentCred != nil {
        holder.credentials["student_id"] = studentCred
    }
    
    // Issue degree credential
    degreeCred := createDegreeCredential(issuer.account, issuerAddress, holderAddress)
    if degreeCred != nil {
        holder.credentials["degree"] = degreeCred
    }
}
```

**Key patterns:**
- **Multiple credential types**: Email, student ID, and degree credentials
- **Organized storage**: Credentials stored with descriptive keys
- **Error handling**: Graceful handling of issuance failures
- **Exchange preparation**: Credentials ready for search and matching

### 3. Exchange Request Processing
```go
func demonstrateExchangeWorkflow(holder *ExchangeParty) {
    // Simulate exchange request
    requestedTypes := []string{"StudentCredential", "DegreeCredential"}
    
    // Search credential store
    var matchingCreds []*credential.VerifiableCredential
    
    for credName, cred := range holder.credentials {
        credTypes := cred.CredentialType()
        for _, credType := range credTypes {
            for _, requestedType := range requestedTypes {
                if credType == requestedType {
                    matchingCreds = append(matchingCreds, cred)
                    break
                }
            }
        }
    }
}
```

**Exchange mechanics:**
- **Request specification**: Verifier specifies needed credential types
- **Store search**: Holder searches their credential collection
- **Type matching**: Credentials matched against request criteria
- **Collection building**: Matching credentials gathered for presentation

### 4. Credential Types Demonstrated

#### Email Credential
```go
claims := map[string]interface{}{
    "emailAddress": "student@university.edu",
    "verified":     true,
    "domain":       "university.edu",
}

credentialBuilder := credential.NewCredential().
    CredentialType(credential.CredentialTypeEmail).
    CredentialSubject(credential.AddressKey(holderAddress)).
    CredentialSubjectClaims(claims).
    Issuer(credential.AddressKey(issuerAddress)).
    ValidFrom(time.Now()).
    SignWith(issuerAddress, time.Now())
```

**Use cases:**
- Institutional email verification
- Domain ownership proof
- Communication channel validation

#### Student Credential
```go
claims := map[string]interface{}{
    "studentId":      "STU-2024-001",
    "enrollmentDate": "2020-09-01",
    "status":         "enrolled",
    "program":        "Computer Science",
    "level":          "undergraduate",
}

credentialBuilder := credential.NewCredential().
    CredentialType([]string{"VerifiableCredential", "StudentCredential"}).
    // ... rest of configuration
```

**Use cases:**
- Academic enrollment verification
- Student status confirmation
- Program participation proof

#### Degree Credential
```go
claims := map[string]interface{}{
    "degree":         "Bachelor of Science",
    "major":          "Computer Science",
    "graduationDate": "2024-05-15",
    "gpa":            3.8,
    "honors":         "magna cum laude",
    "institution":    "University of Technology",
}

credentialBuilder := credential.NewCredential().
    CredentialType([]string{"VerifiableCredential", "DegreeCredential"}).
    // ... rest of configuration
```

**Use cases:**
- Educational achievement proof
- Professional qualification verification
- Academic credential validation

## Running the Example

```bash
cd examples/go/02_credentials/02_exchanging_credentials/presentation_request
go run main.go
```

## Expected Output

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

## What Just Happened

1. **Created exchange parties** - University issuer and student holder with separate identities
2. **Issued multiple credentials** - Email, student ID, and degree credentials using issuance patterns
3. **Organized credential store** - Holder maintains searchable collection of credentials
4. **Processed exchange request** - Employer requests education proof from holder
5. **Matched credentials** - Found credentials matching requested types
6. **Created presentation** - Packaged matching credentials for sharing

## Exchange Architecture Patterns

### 1. Credential Store Management
```go
// Organized storage with descriptive keys
holder.credentials["email"] = emailCredential
holder.credentials["student_id"] = studentCredential
holder.credentials["degree"] = degreeCredential

// Search patterns
for credName, cred := range holder.credentials {
    // Type-based matching
    if matchesRequestedType(cred, requestedTypes) {
        matchingCreds = append(matchingCreds, cred)
    }
}
```

### 2. Request Processing Pipeline
```
Request → Store Search → Type Matching → Collection → Presentation → Response
```

### 3. Type-Based Matching
```go
// Flexible type matching
credTypes := cred.CredentialType()  // ["VerifiableCredential", "StudentCredential"]
requestedTypes := []string{"StudentCredential", "DegreeCredential"}

// Match any credential type against any requested type
for _, credType := range credTypes {
    for _, requestedType := range requestedTypes {
        if credType == requestedType {
            // Found match
        }
    }
}
```

## Real-World Applications

### Employment Verification Scenario
**Parties:**
- **Employer** (Verifier): Requests education credentials
- **Job Applicant** (Holder): Provides education proof
- **University** (Issuer): Originally issued the credentials

**Workflow:**
1. Employer requests: `["DegreeCredential", "StudentCredential"]`
2. Applicant searches their credential store
3. Matching credentials found and presented
4. Employer verifies cryptographic signatures
5. Trust established without contacting university

### Benefits Over Traditional Methods
- **No phone calls** to universities for verification
- **Instant verification** through cryptographic signatures
- **Privacy preservation** through selective disclosure
- **Fraud prevention** via tamper-evident credentials
- **Cost reduction** by eliminating verification intermediaries

## Advanced Patterns

### Complex Request Filtering
```go
type CredentialRequest struct {
    Types       []string
    IssuedAfter time.Time
    RequiredClaims []string
    MaxAge      time.Duration
}

func matchCredential(cred *credential.VerifiableCredential, req CredentialRequest) bool {
    // Type matching
    if !matchesAnyType(cred.CredentialType(), req.Types) {
        return false
    }
    
    // Time-based filtering
    if cred.IssuanceDate().Before(req.IssuedAfter) {
        return false
    }
    
    // Claim requirements
    if !hasRequiredClaims(cred, req.RequiredClaims) {
        return false
    }
    
    return true
}
```

### Presentation Metadata
```go
type PresentationResponse struct {
    Presentation *credential.VerifiablePresentation
    Context      string
    Purpose      string
    CreatedAt    time.Time
    ExpiresAt    time.Time
}
```

### Credential Grouping
```go
// Group credentials by category
educationCreds := filterByCategory(holder.credentials, "education")
identityCreds := filterByCategory(holder.credentials, "identity")
professionalCreds := filterByCategory(holder.credentials, "professional")

// Create category-specific presentations
educationPresentation := createPresentation(educationCreds)
```

## Production Considerations

### Security Best Practices
- **Credential encryption** at rest in storage
- **Secure key management** for account cryptographic keys
- **Audit logging** of all exchange activities
- **Request validation** to prevent malicious requests
- **Rate limiting** on exchange operations

### Performance Optimization
- **Indexed storage** for fast credential search
- **Caching** of frequently accessed credentials
- **Lazy loading** of credential details
- **Batch operations** for multiple exchanges
- **Connection pooling** for account management

### Privacy Protection
- **Selective disclosure** of specific claims only
- **Zero-knowledge proofs** for privacy-preserving verification
- **Minimal data sharing** - only what's requested
- **Consent management** for credential sharing
- **Data retention policies** for exchange logs

## Error Handling Patterns

### Graceful Degradation
```go
// Handle missing credentials gracefully
if len(matchingCreds) == 0 {
    return &ExchangeResponse{
        Status: "partial",
        Message: "Some requested credentials not available",
        AvailableTypes: getAvailableTypes(holder.credentials),
    }
}
```

### Validation Layers
```go
func validateExchangeRequest(req ExchangeRequest) error {
    if len(req.RequestedTypes) == 0 {
        return errors.New("no credential types requested")
    }
    
    for _, reqType := range req.RequestedTypes {
        if !isValidCredentialType(reqType) {
            return fmt.Errorf("invalid credential type: %s", reqType)
        }
    }
    
    return nil
}
```

## Integration Examples

### REST API Integration
```go
func handleExchangeRequest(w http.ResponseWriter, r *http.Request) {
    var req ExchangeRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }
    
    presentation, err := exchangeService.ProcessRequest(req)
    if err != nil {
        http.Error(w, "Exchange failed", http.StatusInternalServerError)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(presentation)
}
```

### WebSocket Real-time Exchange
```go
func handleWebSocketExchange(conn *websocket.Conn) {
    for {
        var req ExchangeRequest
        if err := conn.ReadJSON(&req); err != nil {
            break
        }
        
        presentation, err := processExchangeRequest(req)
        response := ExchangeResponse{
            Presentation: presentation,
            Error:       err,
        }
        
        conn.WriteJSON(response)
    }
}
```

## Testing Strategies

### Unit Testing
```go
func TestCredentialMatching(t *testing.T) {
    // Setup test credentials
    testCreds := createTestCredentials()
    
    // Test type matching
    requested := []string{"StudentCredential"}
    matches := findMatching(testCreds, requested)
    
    assert.Equal(t, 1, len(matches))
    assert.Contains(t, matches[0].CredentialType(), "StudentCredential")
}
```

### Integration Testing
```go
func TestFullExchangeWorkflow(t *testing.T) {
    // Setup issuer and holder
    issuer, holder := createTestParties()
    
    // Issue test credentials
    issueTestCredentials(issuer, holder)
    
    // Process exchange request
    request := ExchangeRequest{Types: []string{"DegreeCredential"}}
    presentation, err := processExchange(holder, request)
    
    assert.NoError(t, err)
    assert.NotNil(t, presentation)
}
```

## Troubleshooting

### Common Issues

**"No matching credentials found"**
- Check credential types match exactly
- Verify credentials were successfully issued
- Ensure proper storage organization
- Validate request format

**"Failed to create presentation"**
- Verify holder account is properly configured
- Check credential signatures are valid
- Ensure proper presentation builder usage
- Validate holder permissions

**"Exchange request timeout"**
- Optimize credential search performance
- Implement proper indexing
- Use connection pooling
- Add request caching

### Debugging Tips
- Enable debug logging: `LogLevel: account.LogDebug`
- Log credential store contents
- Trace request processing pipeline
- Validate credential integrity

## Next Steps

### Further Learning
- **Selective Disclosure**: Privacy-preserving credential sharing
- **Multi-party Exchange**: Complex verification scenarios
- **Real-time Messaging**: WebSocket-based exchange protocols
- **Mobile Integration**: QR code and deep link discovery

### Implementation Ideas
- **Job Application System**: Automated credential verification
- **University Admissions**: Academic record verification
- **Professional Licensing**: Qualification verification system
- **Identity Verification**: Multi-factor credential validation

### Advanced Features
- **Credential Composition**: Combining multiple credentials into complex proofs
- **Conditional Logic**: "If degree is CS AND GPA > 3.5 AND graduation within 2 years"
- **Reputation Systems**: Issuer trust scoring and verification
- **Revocation Checking**: Real-time credential validity verification

## Summary

This example demonstrates the fundamental bridge between credential issuance and exchange:

### ✅ **Core Concepts Covered**
- **Foundation Connection**: How issuance enables exchange
- **Credential Storage**: Organization and search patterns
- **Request Processing**: Matching and filtering mechanisms
- **Presentation Creation**: Packaging credentials for sharing

### 🏗️ **Architecture Patterns**
- **Exchange Party Model**: Structured approach to participant management
- **Type-based Matching**: Flexible credential filtering
- **Store Management**: Organized credential collections
- **Error Handling**: Graceful degradation and validation

### 🚀 **Production Readiness**
- **Security Considerations**: Encryption, validation, audit logging
- **Performance Optimization**: Indexing, caching, connection pooling
- **Privacy Protection**: Selective disclosure and consent management
- **Integration Patterns**: REST APIs, WebSocket, mobile applications

You now understand how credential issuance creates the foundation for sophisticated exchange workflows!
