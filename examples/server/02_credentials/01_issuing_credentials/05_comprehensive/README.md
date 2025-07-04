# Comprehensive Credential Issuance

🎯 **What you'll learn:**
- Creating various types of verifiable credentials
- Using the credential builder pattern
- Attaching evidence and files to credentials
- Managing complex nested claims
- Creating verifiable presentations
- Handling complete credential workflows

📚 **Prerequisites:** Complete all previous examples (`01_basic/`, `02_multi_claim/`, `03_with_evidence/`, `04_complex_data/`) first.

## Overview

This comprehensive example demonstrates the full spectrum of credential issuance capabilities using the Self SDK. It combines all the concepts from previous examples into a complete demonstration of the Self SDK's credential management features.

## What This Example Demonstrates

### Core Capabilities
- **Basic credentials**: Simple email verification
- **Multi-claim credentials**: Profile information with multiple fields
- **Evidence-based credentials**: Certifications with file attachments
- **Complex data structures**: Organizational credentials with nested objects and arrays
- **Verifiable presentations**: Packaging credentials for sharing

### SDK Components Showcased
- `account.New()` - Account initialization and configuration
- `credential.NewCredential()` - Direct credential construction
- `object.New()` - Evidence and file attachment management
- `credential.NewPresentation()` - Verifiable presentation creation
- `account.CredentialIssue()` - Direct credential issuance

## Example Progression

This demo follows a progressive learning approach:

### 1. Basic Email Credential
**Foundation concepts**: Builder pattern, claims, signing, issuance

```go
claims := map[string]interface{}{
    "emailAddress":     "john.doe@example.com",
    "verified":         true,
    "verificationDate": time.Now().Format("2006-01-02"),
}

credentialBuilder := credential.NewCredential().
    CredentialType(credential.CredentialTypeEmail).
    CredentialSubject(credential.AddressKey(holderInbox)).
    CredentialSubjectClaims(claims).
    Issuer(credential.AddressKey(issuerInbox)).
    ValidFrom(time.Now()).
    SignWith(issuerInbox, time.Now())
```

**Key learning points:**
- Credentials contain claims about a subject
- Core SDK provides direct credential construction
- Cryptographic signatures ensure integrity
- Timestamps establish validity periods

### 2. Profile Credential with Multiple Claims
**Advanced claims**: Multiple related claims in one credential

```go
claims := map[string]interface{}{
    "firstName":        "John",
    "lastName":         "Doe",
    "displayName":      "John Doe",
    "profileLevel":     "verified",
    "country":          "United States",
    "registrationDate": time.Now().Format("2006-01-02"),
}
```

**Key learning points:**
- Multiple related claims can be grouped in one credential
- Claims can contain different data types (strings, booleans, dates)
- Grouping related information improves efficiency
- Each claim is cryptographically protected

### 3. Custom Credential with Evidence
**Evidence management**: File attachments and presentations

```go
// Create evidence object
evidenceObj, err := object.New("application/pdf", certificateData)

// Upload to secure storage
err = issuerAccount.ObjectUpload(evidenceObj, false)

// Reference in credential
claims := map[string]interface{}{
    "certificationType": "Professional Development",
    "courseName":        "Advanced Go Programming",
    "certificateHash":   fmt.Sprintf("%x", evidenceObj.Hash()),
    // ... other claims
}
```

**Key learning points:**
- Custom credential types support specific use cases
- Evidence provides additional verification material
- Core SDK object.New() handles secure file storage
- Presentations package credentials for sharing
- Hash references link credentials to evidence

### 4. Organization Credential with Complex Claims
**Complex data structures**: Nested objects, arrays, hierarchical data

```go
claims := map[string]interface{}{
    "organizationName": "Acme Corporation",
    "employeeID":       "EMP-12345",
    "address": map[string]interface{}{
        "street":  "123 Main Street",
        "city":    "San Francisco",
        "state":   "CA",
        "country": "United States",
    },
    "skills": []string{
        "Go Programming",
        "Microservices Architecture",
        "Docker & Kubernetes",
    },
    "certifications": []map[string]interface{}{
        {
            "name":       "AWS Solutions Architect",
            "issuer":     "Amazon Web Services",
            "level":      "Professional",
        },
    },
}
```

**Key learning points:**
- Complex nested objects represent real-world data structures
- Arrays handle collections of related information
- Hierarchical organization improves data clarity
- All nested data maintains cryptographic integrity
- Core SDK supports arbitrary JSON structures in claims

## Running the Example

```bash
cd examples/go/02_credentials/01_issuing_credentials/05_comprehensive
go run main.go
```
### Java
```bash
cd java
gradle run
```

## Expected Output

```
Comprehensive Credential Issuance Demo
======================================
Setting up accounts...
Issuer: 1:ABC123...
Holder: 1:DEF456...

CREDENTIAL ISSUANCE EXAMPLES
============================

1. Basic Email Credential
=========================
✅ Email credential issued (Type: [VerifiableCredential Email])
   Email: john.doe@example.com (verified)

2. Profile Credential
=====================
✅ Profile credential issued (Type: [VerifiableCredential ProfileName])
   Name: John Doe, Country: United States

3. Custom Credential with Evidence
==================================
Evidence uploaded: a1b2c3d4... (542 bytes)
✅ Certification credential issued (Type: [VerifiableCredential CertificationCredential])
✅ Presentation created (Type: [VerifiablePresentation])
   Course: Advanced Go Programming, Grade: A+

4. Organization Credential
==========================
✅ Organization credential issued (Type: [VerifiableCredential EmploymentCredential])
   Employee: EMP-12345 - Senior Software Engineer
   Organization: Acme Corporation, Engineering Dept.

✅ All examples completed successfully!
```

## Code Architecture

### Account Setup
```go
func setupClients() (*account.Account, *account.Account) {
    issuerAccount, err := account.New(&account.Config{
        StorageKey:  generateStorageKey("issuer"),
        StoragePath: issuerStorageDir,
        Environment: account.TargetSandbox,
        LogLevel:    account.LogWarn,
    })
    
    holderAccount, err := account.New(&account.Config{
        StorageKey:  generateStorageKey("holder"),
        StoragePath: holderStorageDir,
        Environment: account.TargetSandbox,
        LogLevel:    account.LogWarn,
    })
    
    return issuerAccount, holderAccount
}
```

**What happens:**
- Creates separate encrypted storage for issuer and holder
- Configures sandbox environment for safe development
- Generates unique cryptographic keys for each account
- Sets up proper logging levels

### Credential Builder Pattern
```go
credentialBuilder := credential.NewCredential().
    CredentialType(credentialType).
    CredentialSubject(credential.AddressKey(holderInbox)).
    CredentialSubjectClaims(claims).
    Issuer(credential.AddressKey(issuerInbox)).
    ValidFrom(time.Now()).
    SignWith(issuerInbox, time.Now())

unsignedCredential, err := credentialBuilder.Finish()
credential, err := issuerAccount.CredentialIssue(unsignedCredential)
```

**Key components:**
- **Builder pattern**: Fluent API for credential construction
- **Type specification**: Standard or custom credential types
- **Subject assignment**: Who the credential is about
- **Claims addition**: The actual credential data
- **Issuer specification**: Who is issuing the credential
- **Validity period**: When the credential becomes valid
- **Signing**: Cryptographic signature for integrity

### Evidence Management
```go
func createEvidenceAsset(account *account.Account, data []byte) (*object.Object, error) {
    // Create evidence object
    evidenceObj, err := object.New("application/pdf", data)
    if err != nil {
        return nil, err
    }
    
    // Upload to secure storage
    err = account.ObjectUpload(evidenceObj, false)
    if err != nil {
        return nil, err
    }
    
    return evidenceObj, nil
}
```

**Evidence workflow:**
1. **Creation**: Use `object.New()` to create evidence object
2. **Upload**: Store in encrypted, distributed storage
3. **Reference**: Link evidence ID/hash in credential claims
4. **Verification**: Evidence can be retrieved for validation

### Presentation Creation
```go
func createPresentation(account *account.Account, cred *credential.VerifiableCredential) (*credential.VerifiablePresentation, error) {
    inboxAddress, err := account.InboxOpen()
    if err != nil {
        return nil, err
    }

    builder := credential.NewPresentation().
        PresentationType([]string{"VerifiablePresentation"}).
        Holder(credential.AddressKey(inboxAddress)).
        CredentialAdd(cred)

    unsignedPresentation, err := builder.Finish()
    if err != nil {
        return nil, err
    }

    return account.PresentationIssue(unsignedPresentation)
}
```

## Best Practices Demonstrated

### 1. Account Management
- **Separate storage**: Different directories for issuer and holder
- **Secure keys**: Cryptographically secure storage keys
- **Environment isolation**: Sandbox for development, production for live
- **Proper cleanup**: Defer account.Close() for resource management

### 2. Error Handling
```go
unsignedCredential, err := credentialBuilder.Finish()
if err != nil {
    log.Printf("Failed to build credential: %v", err)
    return
}

credential, err := issuerAccount.CredentialIssue(unsignedCredential)
if err != nil {
    log.Printf("Failed to issue credential: %v", err)
    return
}
```

### 3. Data Organization
- **Logical grouping**: Related claims in single credentials
- **Type consistency**: Appropriate data types for different claim values
- **Hierarchical structure**: Nested objects for complex relationships
- **Array usage**: Collections for multiple related items

### 4. Security Considerations
- **Evidence hashing**: Link evidence through cryptographic hashes
- **Signature verification**: All credentials cryptographically signed
- **Storage encryption**: All data encrypted at rest
- **Identity verification**: DIDs provide verifiable identity

## Production Considerations

### Scalability
- **Batch operations**: Consider batch processing for multiple credentials
- **Async patterns**: Use goroutines for concurrent operations
- **Connection pooling**: Manage account connections efficiently
- **Resource management**: Proper cleanup and resource disposal

### Security
- **Key management**: Secure storage and rotation of encryption keys
- **Access controls**: Implement proper authorization patterns
- **Audit logging**: Track all credential operations
- **Evidence security**: Secure handling of sensitive evidence files

### Performance
- **Credential caching**: Cache frequently accessed credentials
- **Evidence optimization**: Optimize evidence file sizes
- **Network efficiency**: Minimize unnecessary network calls
- **Storage optimization**: Efficient use of storage resources

## Integration Patterns

### Microservices Architecture
```go
type CredentialService struct {
    issuerAccount *account.Account
    config        *Config
}

func (cs *CredentialService) IssueEmailCredential(email string) (*credential.VerifiableCredential, error) {
    claims := map[string]interface{}{
        "emailAddress": email,
        "verified":     true,
        "verificationDate": time.Now().Format("2006-01-02"),
    }
    
    return cs.buildAndIssueCredential(credential.CredentialTypeEmail, claims)
}
```

### API Integration
```go
func HandleCredentialRequest(w http.ResponseWriter, r *http.Request) {
    var req CredentialRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }
    
    cred, err := credentialService.IssueCredential(req.Type, req.Claims)
    if err != nil {
        http.Error(w, "Failed to issue credential", http.StatusInternalServerError)
        return
    }
    
    json.NewEncoder(w).Encode(cred)
}
```

## Troubleshooting

### Common Issues

**"Failed to create account"**
- Check storage directory permissions
- Verify unique storage paths for different accounts
- Ensure sufficient disk space
- Validate environment configuration

**"Failed to build credential"**
- Verify all required fields are provided
- Check claim data types are JSON-serializable
- Ensure addresses are properly opened
- Validate credential type format

**"Failed to upload evidence"**
- Check network connectivity
- Verify storage quota availability
- Ensure account has upload permissions
- Validate evidence file format and size

### Debugging Tips
- Enable debug logging: `LogLevel: account.LogDebug`
- Use structured logging for better traceability
- Implement retry logic for network operations
- Add validation layers for input data

## Next Steps

### Further Learning
- **Credential Exchange**: Study `../../02_exchanging_credentials/` examples
- **Advanced Features**: Explore credential verification and validation
- **Production Deployment**: Learn about scaling and monitoring
- **Custom Types**: Implement domain-specific credential types

### Implementation Ideas
- **Identity Provider**: Build OAuth-like service with verifiable credentials
- **Document Verification**: Create system for document authenticity
- **Academic Records**: Implement digital diploma and transcript system
- **Professional Licensing**: Build licensing authority with credential issuance

### Resources
- **Self SDK Documentation**: Complete API reference and guides
- **Community Examples**: Real-world implementation patterns
- **Best Practices**: Production deployment and security guidelines
- **Support**: Community forums and technical assistance

## Summary

This comprehensive example demonstrates the full power of the Self SDK for credential issuance:

### ✅ **Features Demonstrated**
1. **Basic Email Credential** - Foundation concepts with core SDK
2. **Profile Credential** - Multiple claims in single credential  
3. **Custom Credential** - Evidence attachments and presentations
4. **Organization Credential** - Complex nested data structures

### 🔧 **Key SDK Components Utilized**
- `account.New()` - Account initialization and configuration
- `credential.NewCredential()` - Direct credential construction
- `object.New()` - Evidence and file attachment management
- `credential.NewPresentation()` - Verifiable presentation creation
- `account.CredentialIssue()` - Direct credential issuance

### 📚 **Educational Progression Complete**
- Core SDK concepts and architecture
- Direct credential building patterns
- Evidence and asset management
- Complex data modeling techniques
- Production-ready development patterns

You're now ready to build complete identity solutions using the Self SDK! 
