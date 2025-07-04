# Evidence-Based Credential Issuance

🎯 **What you'll learn:**
- How to attach evidence files to credentials
- Asset management and secure storage
- Creating verifiable presentations
- Linking evidence to claims with hashes
- Custom credential types

📚 **Prerequisites:** Complete `../01_basic/` and `../02_multi_claim/` examples first.

## Overview

This example demonstrates creating credentials with attached file evidence using the Self SDK. You'll learn how to upload supporting documents, reference them in credentials, and create verifiable presentations that include both claims and evidence.

## Key Concepts

### Evidence Assets
Evidence files provide supporting documentation for credential claims:
- **Certificates**: Diplomas, course completion certificates
- **Documents**: Transcripts, official records, licenses
- **Images**: Photos, scanned documents, signatures
- **Data files**: JSON records, XML documents

### Asset Management
The Self SDK provides secure storage for evidence:
```go
// Create evidence object
evidence, err := object.New("certificate.pdf", certificateData)

// Upload to secure storage
err = issuer.ObjectUpload(evidence, true)

// Reference in credential claims
claims["evidenceId"] = fmt.Sprintf("%x", evidence.Id())
```

### Verifiable Presentations
Presentations bundle credentials for sharing:
- **Selective disclosure**: Choose which credentials to include
- **Context preservation**: Maintain relationship between credentials
- **Cryptographic integrity**: All included credentials remain verifiable

## Code Breakdown

### 1. Creating Evidence Assets
```go
func createEvidence(issuer *account.Account) *object.Object {
    // Create document content
    certificateData := []byte(`Certificate of Completion
Advanced Go Programming Course
Student: John Doe
...`)

    // Create evidence object
    evidence, err := object.New("certificate.pdf", certificateData)
    if err != nil {
        log.Printf("Failed to create evidence object: %v", err)
        return nil
    }

    // Upload to secure storage
    err = issuer.ObjectUpload(evidence, true)
    if err != nil {
        log.Printf("Failed to upload evidence: %v", err)
        return nil
    }

    return evidence
}
```

**What happens:**
1. **Document creation**: Generates certificate content as byte array
2. **Object creation**: Creates Self SDK object with filename and data
3. **Secure upload**: Stores evidence in encrypted, distributed storage
4. **ID generation**: Returns unique identifier for referencing

### 2. Linking Evidence to Credentials
```go
claims := map[string]interface{}{
    "certificationType": "Professional Development",
    "courseName":        "Advanced Go Programming",
    "completionDate":    time.Now().Format("2006-01-02"),
    "grade":             "A+",
    "institution":       "Self SDK Academy",
    "courseHours":       40,
    "instructor":        "Jane Smith",
    "evidenceId":        fmt.Sprintf("%x", evidence.Id()),  // Link to evidence
}
```

**Key points:**
- `evidenceId` claim links credential to supporting evidence
- Evidence ID is cryptographically unique and tamper-proof
- Claims describe the achievement, evidence proves it
- Multiple evidence files can be referenced if needed

### 3. Custom Credential Types
```go
credentialBuilder := credential.NewCredential().
    CredentialType([]string{"VerifiableCredential", "CertificationCredential"}).
    CredentialSubject(credential.AddressKey(holderAddress)).
    CredentialSubjectClaims(claims).
    Issuer(credential.AddressKey(issuerAddress)).
    ValidFrom(time.Now()).
    SignWith(issuerAddress, time.Now())
```

**Key points:**
- `CertificationCredential` is a custom type for course certifications
- Multiple types can be specified in an array
- Custom types help categorize and filter credentials
- Standard `VerifiableCredential` type is always included

### 4. Creating Verifiable Presentations
```go
presentationBuilder := credential.NewPresentation().
    PresentationType([]string{"VerifiablePresentation", "CertificationPresentation"}).
    Holder(credential.AddressKey(issuerAddress)).
    CredentialAdd(cred)

presentation, err := issuer.PresentationIssue(unsignedPresentation)
```

**What happens:**
1. **Presentation builder**: Creates container for credentials
2. **Type specification**: Defines presentation purpose and context
3. **Credential addition**: Includes one or more credentials
4. **Cryptographic signing**: Signs the entire presentation

## Running the Example

### Go
```bash
cd examples/go/02_credentials/01_issuing_credentials/03_with_evidence
go run main.go
```

### Java
```bash
cd java
gradle run
```

## Expected Output

```
Evidence-Based Credential Issuance Demo
=======================================
Issuer: 1:ABC123...
Holder: 1:DEF456...

Creating certification with evidence...
Creating evidence asset...
✅ Evidence uploaded: a1b2c3d4... (542 bytes)
Building certification credential...
✅ Certification credential issued (Type: [VerifiableCredential CertificationCredential])
Creating verifiable presentation...
✅ Presentation created (Type: [VerifiablePresentation CertificationPresentation], Credentials: 1)
✅ Demo completed successfully!
```

## What Just Happened

1. **Created evidence asset**: Generated a PDF certificate document
2. **Uploaded to secure storage**: Stored evidence in encrypted, distributed storage
3. **Built custom credential**: Created certification credential with evidence reference
4. **Linked evidence to claims**: Connected supporting document to achievement claims
5. **Created verifiable presentation**: Bundled credential for sharing and verification

## Benefits of Evidence-Based Credentials

### Verifiability
- **Supporting documentation**: Claims backed by actual evidence
- **Tamper detection**: Any modification to evidence invalidates the credential
- **Source verification**: Evidence source can be cryptographically verified

### Trust
- **Transparency**: Verifiers can access supporting documentation
- **Audit trail**: Complete record of evidence and claims
- **Fraud prevention**: Difficult to forge with cryptographic evidence links

### Compliance
- **Regulatory requirements**: Many industries require supporting documentation
- **Audit readiness**: Evidence preserved in immutable storage
- **Legal validity**: Cryptographic proofs for legal proceedings

## Evidence Management Best Practices

### File Formats
```go
// ✅ Good: Standard formats
evidence, _ := object.New("diploma.pdf", pdfData)       // Documents
evidence, _ := object.New("transcript.json", jsonData)  // Structured data
evidence, _ := object.New("photo.jpg", imageData)       // Images

// ⚠️ Consider: Large files may impact performance
evidence, _ := object.New("video.mp4", videoData)       // Large multimedia
```

### Security Considerations
```go
// ✅ Good: Sensitive evidence
claims := map[string]interface{}{
    "evidenceId": fmt.Sprintf("%x", evidence.Id()),     // Reference only
    "evidenceHash": calculateHash(evidenceData),        // Integrity check
}

// ❌ Avoid: Sensitive data in claims
claims := map[string]interface{}{
    "ssn": "123-45-6789",           // Sensitive data exposed
    "fullDocument": string(pdfData), // Large data in claims
}
```

### Evidence Lifecycle
```go
// Create evidence with metadata
evidence, _ := object.New("certificate.pdf", certificateData)

// Upload with encryption
err := issuer.ObjectUpload(evidence, true)  // true = encrypt

// Reference in credential
evidenceId := fmt.Sprintf("%x", evidence.Id())

// Later: Retrieve evidence for verification
// retrievedEvidence := issuer.ObjectGet(evidenceId)
```

## Advanced Patterns

### Multiple Evidence Files
```go
claims := map[string]interface{}{
    "courseName": "Advanced Go Programming",
    "transcript": fmt.Sprintf("%x", transcriptEvidence.Id()),
    "certificate": fmt.Sprintf("%x", certificateEvidence.Id()),
    "portfolio": fmt.Sprintf("%x", portfolioEvidence.Id()),
}
```

### Evidence Validation
```go
// Include evidence hash for integrity checking
evidenceHash := sha256.Sum256(certificateData)
claims["evidenceHash"] = fmt.Sprintf("%x", evidenceHash)
claims["evidenceId"] = fmt.Sprintf("%x", evidence.Id())
```

### Conditional Evidence
```go
// Different evidence based on credential level
if grade == "A+" {
    claims["honorsEvidence"] = fmt.Sprintf("%x", honorsEvidence.Id())
}
```

## Troubleshooting

### Common Issues

**"Failed to create evidence object"**
- Check file data is valid byte array
- Ensure filename doesn't contain invalid characters
- Verify data size isn't excessive

**"Failed to upload evidence"**
- Network connectivity issues
- Storage quota exceeded
- Insufficient account permissions

**"Evidence not found in storage"**
- Evidence ID may be incorrect
- Upload may have failed silently
- Storage path configuration issues

### Debugging Tips
- Enable debug logging: `LogLevel: account.LogDebug`
- Check evidence ID format and validity
- Verify storage configuration and permissions
- Test with smaller files first

## Production Considerations

### Storage Costs
- Evidence files consume storage space
- Consider file size optimization
- Implement retention policies

### Performance
- Large evidence files slow credential creation
- Consider async upload patterns
- Cache frequently accessed evidence

### Privacy
- Evidence may contain sensitive information
- Implement access controls
- Consider evidence encryption levels

## Next Steps

📚 **Continue learning:**
- `../04_complex_data/` - Complex nested data structures
- `../05_comprehensive/` - All features combined
- `../../02_exchanging_credentials/` - Using credentials in presentations

🔍 **Related concepts:**
- Asset management and storage
- Cryptographic integrity verification
- Verifiable presentation patterns

## When to Use Evidence-Based Credentials

### Perfect for:
- **Educational credentials**: Diplomas, certificates, transcripts
- **Professional licenses**: Medical licenses, bar admissions, certifications
- **Identity documents**: Passports, driver's licenses, official records
- **Achievement records**: Awards, completions, assessments

### Consider alternatives for:
- **Simple assertions**: Basic claims that don't require proof
- **Real-time data**: Frequently changing information
- **Large datasets**: Consider linking to external storage
- **Privacy-sensitive**: Evidence that shouldn't be permanently stored 
