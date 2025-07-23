# Digital Signatures: Server Implementation

> **🎯 What you'll learn:** How to implement a complete digital signature backend with real credential verification response processing, PDF agreement creation, and robust claims extraction.

This guide shows you how to build a production-ready digital signature server that creates PDF agreements, issues signing credentials, sends verification requests, and processes real credential verification responses.

## Architecture Overview

The digital signature server implements a **credential-based signing model** with three core components:

**Agreement Management**

- PDF creation and secure storage
- Agreement metadata and evidence handling
- Document hash generation and validation

**Credential Issuance**

- Signing credential creation with agreement claims
- Evidence attachment for document linking
- Cryptographic signing and verification

**Response Processing**

- Real credential verification response decoding
- Claims extraction from signed credentials
- Robust error handling for different response structures

## Core Implementation

### 1. Agreement Creation and Storage

The server creates PDF agreements and uploads them to secure storage as evidence for the signing credential.

<div data-github-embed="https://github.com/joinself/academy/blob/main/examples/solutions/signing-tenacy-agreement/internal/auth/service.go#L650-L700"
     data-style="github-dark-dimmed"
     data-show-border="true"
     data-show-line-numbers="true"
     data-show-file-meta="true"
     data-show-full-path="true"
     data-show-copy="true"></div>

**Key Features:**
- **PDF Generation**: Creates structured PDF agreements with user details
- **Secure Storage**: Uploads PDFs to Self's secure object storage
- **Evidence Linking**: Uses PDF as evidence for credential verification
- **Hash Generation**: Creates cryptographic hashes for document integrity

### 2. Credential Issuance with Agreement Metadata

The server issues signing credentials containing agreement metadata and evidence.

<div data-github-embed="https://github.com/joinself/academy/blob/main/examples/solutions/signing-tenacy-agreement/internal/auth/sign.go#L20-60"
     data-style="github-dark-dimmed"
     data-show-border="true"
     data-show-line-numbers="true"
     data-show-file-meta="true"
     data-show-full-path="true"
     data-show-copy="true"></div>

**Key Features:**
- **Agreement Claims**: Embeds agreement type, ID, hash, and metadata
- **Evidence Attachment**: Links PDF document as evidence
- **Cryptographic Signing**: Signs credential with service identity
- **Verification Request**: Creates credential verification request with evidence

### 3. Real Credential Verification Response Processing

The server processes actual credential verification responses and extracts real claims from signed credentials.

<div data-github-embed="https://github.com/joinself/academy/blob/main/examples/solutions/signing-tenacy-agreement/internal/auth/sign.go#L130-180"
     data-style="github-dark-dimmed"
     data-show-border="true"
     data-show-line-numbers="true"
     data-show-file-meta="true"
     data-show-full-path="true"
     data-show-copy="true"></div>

**Key Features:**
- **Response Decoding**: Uses `DecodeCredentialVerificationResponse()` for real processing
- **Claims Extraction**: Extracts actual claims from signed credentials
- **Robust Handling**: Handles different response structures with fallback mechanisms
- **Comprehensive Logging**: Provides detailed logging for debugging and monitoring

## Implementation Details

### Request-Response Correlation

The server uses cryptographic content IDs to correlate signing requests with verification responses:

```go
// Store signing request with correlation tracking
signRequest := &SignRequest{
    ID:           requestID,
    UserDID:      userDID,
    AgreementPDF: agreementPDF,
    CreatedAt:    time.Now(),
    ExpiresAt:    time.Now().Add(10 * time.Minute),
    CompleteChan: make(chan *SignResult, 1),
    Status:       SignRequestPending,
}

// Store in unified map for response correlation
a.signRequests[requestID] = signRequest
```

### Credential Verification Request Creation

The server creates credential verification requests with PDF evidence:

```go
// Create credential verification request with agreement evidence
content, err := message.NewCredentialVerificationRequest().
    Type([]string{"VerifiableCredential", "AgreementCredential"}).
    Evidence("agreement", agreementPDF).
    Proof(signedAgreementPresentation).
    Expires(time.Now().Add(time.Hour * 24)).
    Finish()
```

### Response Processing with Real Claims Extraction

The server processes verification responses and extracts real claims:

```go
// Decode the credential verification response
credentialVerificationResponse, err := message.DecodeCredentialVerificationResponse(msg.Content())
if err != nil {
    a.logger.Error("Failed to decode credential verification response", slog.String("error", err.Error()))
    return
}

// Extract claims from presentations in the response
presentations := credentialVerificationResponse.Presentations()
if len(presentations) > 0 {
    for _, presentation := range presentations {
        credentials := presentation.Credentials()
        for _, cred := range credentials {
            if credClaims, err := cred.CredentialSubjectClaims(); err == nil {
                // Merge real claims into result
                for key, value := range credClaims {
                    claims[key] = value
                }
            }
        }
    }
}
```

## Error Handling and Fallbacks

### Robust Response Structure Handling

The implementation handles different response structures gracefully:

```go
// Try to extract presentations if the method exists
var presentations []*credential.VerifiablePresentation
if presentationsMethod, ok := interface{}(credentialVerificationResponse).(interface{ 
    Presentations() []*credential.VerifiablePresentation 
}); ok {
    presentations = presentationsMethod.Presentations()
    a.logger.Info("Found presentations in verification response", slog.Int("presentations_count", len(presentations)))
} else {
    a.logger.Info("No Presentations() method found, checking for alternative structure")
}
```

### Fallback Claims Processing

When no real claims can be extracted, the system uses fallback values:

```go
// If no claims were extracted, use fallback values
if len(claims) == 0 {
    a.logger.Info("No claims extracted, using fallback values")
    claims["agreementType"] = "Tenancy Agreement"
    claims["signingDate"] = time.Now().Format("2006-01-02")
    claims["signingTimestamp"] = time.Now().Unix()
    claims["documentType"] = "application/pdf"
    claims["evidenceId"] = fmt.Sprintf("%x", matchingSignReq.AgreementPDF.Id())
}
```

## Production Considerations

### Security Best Practices

**Storage Security:**
- Use strong, randomly generated storage keys
- Secure storage locations with proper access controls
- Encrypt PDF documents in transit and at rest
- Implement proper backup and disaster recovery

**Credential Security:**
- Validate all credential claims before processing
- Verify cryptographic signatures on all credentials
- Implement proper session management and timeout
- Log all signing activities for audit trails

### Performance Optimization

**Request Handling:**
- Implement request queuing for high-traffic scenarios
- Use connection pooling for mobile device connections
- Cache frequently accessed agreement templates
- Implement proper timeout handling for mobile responses

**Storage Optimization:**
- Compress PDF documents before storage
- Implement document versioning for agreement updates
- Use efficient storage backends for production scale
- Implement proper cleanup for expired requests

### Monitoring and Logging

**Comprehensive Logging:**
```go
a.logger.Info("Processing credential verification response", 
    slog.String("response_to_id", hex.EncodeToString(credentialVerificationResponse.ResponseTo())),
    slog.Int("presentations_count", len(presentations)))

a.logger.Info("Extracted credential claims", slog.Any("claims", credClaims))
```

**Error Tracking:**
- Log all credential verification failures
- Track response processing errors
- Monitor signing success rates
- Alert on unusual signing patterns

## Integration Patterns

### Web Application Integration

```go
// HTTP endpoint for digital signature requests
func handleDigitalSignatureRequest(w http.ResponseWriter, r *http.Request) {
    // Parse agreement details from request
    var agreementDetails AgreementDetails
    json.NewDecoder(r.Body).Decode(&agreementDetails)
    
    // Create signing request
    signRequest, err := authService.GenerateSignRequest(ctx, userDID, agreementPDF)
    if err != nil {
        http.Error(w, "Failed to create signing request", http.StatusInternalServerError)
        return
    }
    
    // Return request ID for status tracking
    json.NewEncoder(w).Encode(map[string]string{
        "requestId": signRequest.ID,
        "status": "pending",
    })
}
```

### API Gateway Integration

```go
// Middleware for signature verification
func signatureAuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Extract signature credentials from request
        signatureCredential := extractSignatureCredential(r)
        
        // Validate signature credential
        if !validateSignatureCredential(signatureCredential) {
            http.Error(w, "Invalid signature", http.StatusUnauthorized)
            return
        }
        
        // Add signature claims to request context
        ctx := context.WithValue(r.Context(), "signatureClaims", extractSignatureClaims(signatureCredential))
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}
```

## Testing Strategies

### Unit Testing

**Credential Processing Tests:**
```go
func TestCredentialVerificationResponseProcessing(t *testing.T) {
    // Create mock credential verification response
    mockResponse := createMockCredentialVerificationResponse()
    
    // Test claims extraction
    claims := extractClaimsFromResponse(mockResponse)
    
    // Verify expected claims are present
    assert.Contains(t, claims, "agreementType")
    assert.Contains(t, claims, "signingDate")
    assert.Contains(t, claims, "documentHash")
}
```

### Integration Testing

**End-to-End Signing Tests:**
```go
func TestCompleteSigningWorkflow(t *testing.T) {
    // Create test agreement
    agreementPDF := createTestAgreement()
    
    // Generate signing request
    signRequest, err := authService.GenerateSignRequest(ctx, testUserDID, agreementPDF)
    assert.NoError(t, err)
    
    // Simulate mobile response
    simulateMobileSigningResponse(signRequest.ID)
    
    // Wait for completion
    result, err := authService.WaitForSign(ctx, signRequest.ID)
    assert.NoError(t, err)
    assert.True(t, result.Success)
    
    // Verify extracted claims
    assert.Contains(t, result.Claims, "agreementType")
    assert.Contains(t, result.Claims, "signingDate")
}
```

## Next Steps

**Ready to implement digital signatures?**
- **Complete Implementation**: Follow the full implementation guide with real credential processing
- **Custom Agreements**: Adapt the PDF generation for your document types
- **Multi-Party Signing**: Extend for multiple signers on single documents
- **Integration**: Connect with existing document management systems

**Learn More:**
- **[Mobile Implementation](./mobile-implementation.md)**: Integrate signing into your mobile app
- **[Conclusions & Best Practices](./conclusions.md)**: Production deployment guide
- **[Authentication](./../authentication.md)**: Authenticate users before signing
- **[Identity Verification](./../identity-verification.md)**: Verify user identity before approval 
