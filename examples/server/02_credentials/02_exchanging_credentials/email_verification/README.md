# Real Credential Presentation Request Demo 📱

> **📖 Learn the concepts first:** [Verifiable Credentials Concepts](../../../../../../docs/concepts/verifiable-credentials.md)

🎯 **What you'll learn:**
- Requesting real credentials from mobile devices
- Validating actual email verification and liveness proofs
- Building credential verifier services
- Handling presentation request workflows
- Processing authentic mobile credential responses

📚 **Prerequisites:** Complete foundation patterns first:
- `../presentation_request/` - Basic exchange patterns
- `../../01_issuing_credentials/01_basic/` - Credential creation fundamentals

## Overview

This example demonstrates how to request and validate **real credentials** from mobile devices. Unlike other examples that create mock credentials, this connects to actual Self mobile apps and requests genuine credentials like verified email addresses and liveness proofs.

## Key Concepts

### Real Credential Validation
This example shows how to:
- **Connect to mobile devices** - Establish secure connections via QR codes
- **Request actual credentials** - Ask for real email and liveness verification
- **Validate authentic data** - Process credentials issued by trusted authorities
- **Build verifier services** - Create applications that consume real credentials

### Mobile Credential Ecosystem
- **Trusted Issuers**: Email services, government agencies, and identity providers
- **Mobile Wallets**: Self mobile apps storing real credentials
- **Verifier Services**: Applications that request and validate credentials
- **Authentic Claims**: Real email addresses, liveness proofs, identity documents

## Code Breakdown

### 1. Credential Verifier Setup
```go
func createVerifierAccount() *account.Account {
    fmt.Println("Setting up credential verifier...")
    
    verifierService := common.SetupAccount(common.AccountConfig{
        StorageDir: "verifier_service",
        Callbacks: account.Callbacks{
            OnWelcome: handleMobileConnection,
            OnMessage: handleMessage,
        },
    })
    
    return verifierService
}
```

**Verifier role:**
- **Service Provider**: Application requesting credentials from users
- **Trust Evaluator**: Validates credentials against trusted issuers
- **Decision Maker**: Uses credential data for access control or verification

### 2. Mobile Connection via QR Code
```go
func generateConnectionQR(verifierService *account.Account) bool {
    // Generate key package for secure communication
    keyPackage, err := verifierService.ConnectionNegotiateOutOfBand(
        inboxAddress,
        time.Now().Add(30*time.Minute),
    )
    
    // Build discovery request for mobile connection
    content, err := message.NewDiscoveryRequest().
        KeyPackage(keyPackage).
        Expires(time.Now().Add(30 * time.Minute)).
        Finish()
    
    // Create QR code for mobile scanning
    qrCode, err := event.NewAnonymousMessage(content).
        SetFlags(event.MessageFlagTargetSandbox).
        EncodeToQR(event.QREncodingUnicode)
}
```

**Connection flow:**
1. Service generates QR code with connection request
2. User scans QR code with Self mobile app
3. Secure encrypted channel established
4. Service can now request credentials

### 3. Real Credential Request
```go
func requestCredentialsFromMobile(verifierService *account.Account) {
    // Create presentation request for multiple credential types
    content, err := message.NewCredentialPresentationRequest().
        Type([]string{"VerifiablePresentation", "MobileCredentialPresentation"}).
        Details(
            credential.CredentialTypeLiveness,
            []*message.CredentialPresentationDetailParameter{
                message.NewCredentialPresentationDetailParameter(
                    message.OperatorNotEquals,
                    "sourceImageHash",
                    "",
                ),
            },
        ).
        Details(
            credential.CredentialTypeEmail,
            []*message.CredentialPresentationDetailParameter{
                message.NewCredentialPresentationDetailParameter(
                    message.OperatorNotEquals,
                    "emailAddress",
                    "",
                ),
            },
        ).
        Finish()
    
    // Send request to mobile device
    err = verifierService.MessageSend(connectedMobile, content)
}
```

**Request types:**
- **Liveness Credentials**: Biometric proofs from trusted verification services
- **Email Credentials**: Verified email addresses from email providers
- **Custom Requirements**: Specific claim parameters using operators

### 4. Processing Real Responses
```go
func displayReceivedCredentials(verifierService *account.Account, response *message.CredentialPresentationResponse) {
    presentations := response.Presentations()
    
    for i, presentation := range presentations {
        credentials := presentation.Credentials()
        
        for j, cred := range credentials {
            // Display actual credential information
            fmt.Printf("      • Type: %v\n", cred.CredentialType())
            fmt.Printf("      • Issuer: %s\n", cred.Issuer().String())
            fmt.Printf("      • Subject: %s\n", cred.CredentialSubject().String())
            
            // Show real claims from mobile device
            if claims, err := cred.CredentialSubjectClaims(); err == nil {
                for key, value := range claims {
                    fmt.Printf("        - %s: %v\n", key, value)
                }
            }
        }
    }
}
```

**Real data examples:**
- **Email**: `adria.cidre@gmail.com` (actual verified email)
- **Liveness**: Real image hashes from biometric verification
- **Issuers**: `did:aure:0088a8de...` (trusted verification services)

## Running the Example

### Go
```bash
cd examples/server/02_credentials/02_exchanging_credentials/email_verification/go
go run main.go
```

### Java
```bash
cd java
gradle run
```

## Expected Output

```
Real Credential Presentation Request Demo
=======================================
Setting up credential verifier...
✅ Credential verifier ready
📬 Credential Verifier Address: did:key:z6Mk...

Generating QR code for mobile connection...

[QR CODE DISPLAYED]

📱 SCAN QR CODE with Self mobile app
⏳ Waiting for mobile device connection...
📋 Once connected, will request real credentials from your device

📱 Mobile device connected: 001d56310f8462280d193906e15b776b564dba9169977bf15d8c27135bf73edc0d
✅ Mobile connection established!
📋 Now requesting real credentials from your mobile device...

📋 Requesting real credentials from mobile device...
💡 This demonstrates validation of actual mobile credentials
📤 Sending credential request to mobile: 001d56310f8462280d193906e15b776b564dba9169977bf15d8c27135bf73edc0d
   • Requesting: Liveness and Email credentials
   • This will show your actual verified credentials
⏳ Waiting for credential presentation response...

✅ Received real credentials from mobile device!
✅ Received 1 presentation(s) from mobile device:

📜 Presentation #1:
   • Type: [VerifiablePresentation MobileCredentialPresentation]
   • Holder: did:key:z6MkgRn6ahSTaBthvjJAcgDie7qzHAQFtJ36d4S2Rb24ypQk
   • Contains 2 real credential(s):

   📋 Real Credential #1:
      • Type: [VerifiableCredential LivenessCredential]
      • Issuer: did:aure:0088a8de42c392bf92711c7cffcfcdf6e8e79b7e1af556c97cf095b3ab15cb3c69
      • Subject: did:key:z6MkmPP9VeB4bEuEaZg96ZMjJFXaDvxLx89LNRnZRCVBs1Jc
      • Valid From: 2025-07-09 12:09:09
      • Actual Claims:
        - sourceImageHash: 210482c4c7ff1932fbdfe575bc250de3501bffc2137fbf87d5e2d43bcc87829d
        - targetImageHash: 9231a0de04edff82b9ad0f949309b290260bc21302d9a101a279a044e2813485
        - id: did:key:z6MkmPP9VeB4bEuEaZg96ZMjJFXaDvxLx89LNRnZRCVBs1Jc

   📋 Real Credential #2:
      • Type: [VerifiableCredential EmailCredential]
      • Issuer: did:aure:0088a8de42c392bf92711c7cffcfcdf6e8e79b7e1af556c97cf095b3ab15cb3c69
      • Subject: did:key:z6MkrzTe1WcSY7k58WCcVRZ442FAQnREtosdSqsXzutNaqri
      • Valid From: 2025-07-09 12:02:13
      • Actual Claims:
        - emailAddress: adria.cidre@gmail.com
        - id: did:key:z6MkrzTe1WcSY7k58WCcVRZ442FAQnREtosdSqsXzutNaqri

🎉 Real credential validation completed!
💡 This demonstrates:
   • Requesting real credentials from mobile devices
   • Receiving actual verified email addresses and liveness proofs
   • Validating authentic credentials issued by trusted authorities
   • Building trust through real identity verification

🏁 Demo completed successfully!
```

## What Just Happened

1. **Created verifier service** - Application that requests credentials from mobile devices
2. **Generated QR connection** - Secure method for mobile devices to connect
3. **Requested real credentials** - Asked mobile for actual email and liveness proofs
4. **Received authentic data** - Got real email address and biometric verification
5. **Processed valid credentials** - Displayed actual claims from trusted issuers

## Real Credential Architecture

### 1. Trust Network
```
Trusted Issuers → Mobile Wallets → Verifier Services
     ↓               ↓               ↓
Email Providers   Self App      Your Application
Gov Agencies     Stored Creds   Validation Logic
Identity Orgs    User Control   Access Decisions
```

### 2. Credential Flow
```
Real Issuer → Real Credential → Mobile Wallet → Presentation → Verifier Service
     ↓             ↓              ↓              ↓              ↓
Gmail         Email Verified    Self App      QR Connect    Your App
Gov ID        Liveness Proof    Secure Store  Auth Request  Decision
Bank KYC      Identity Doc      User Owner    Real Claims   Access Grant
```

### 3. Verification Levels
```go
// Different types of real credentials
type CredentialType struct {
    Type       string
    Issuer     string
    TrustLevel string
}

emailCredential := CredentialType{
    Type:       "EmailCredential",
    Issuer:     "did:aure:0088a8de...", // Real email verification service
    TrustLevel: "high",
}

livenessCredential := CredentialType{
    Type:       "LivenessCredential", 
    Issuer:     "did:aure:0088a8de...", // Real biometric service
    TrustLevel: "highest",
}
```

## Real-World Use Cases

### 1. Age Verification Service
```go
// Request real age verification credentials
ageRequest := message.NewCredentialPresentationRequest().
    Details(credential.CredentialTypeAge, []*message.CredentialPresentationDetailParameter{
        message.NewCredentialPresentationDetailParameter(
            message.OperatorGreaterThan,
            "age",
            "18",
        ),
    })
```

### 2. Financial KYC Application
```go
// Request comprehensive identity verification
kycRequest := message.NewCredentialPresentationRequest().
    Details(credential.CredentialTypeEmail, nil).
    Details(credential.CredentialTypeLiveness, nil).
    Details(credential.CredentialTypeGovernmentID, nil)
```

### 3. Education Access Portal
```go
// Request student status verification
studentRequest := message.NewCredentialPresentationRequest().
    Details(credential.CredentialTypeEducation, []*message.CredentialPresentationDetailParameter{
        message.NewCredentialPresentationDetailParameter(
            message.OperatorEquals,
            "studentStatus",
            "enrolled",
        ),
    })
```

## Security Considerations

### 1. Trust Registry Management
```go
// Only accept credentials from trusted issuers
trustedIssuers := []string{
    "did:aure:0088a8de42c392bf92711c7cffcfcdf6e8e79b7e1af556c97cf095b3ab15cb3c69", // Email verification service
    "did:gov:us:dmv:...", // Government DMV
    "did:bank:chase:...", // Financial institution
}

func validateIssuerTrust(credential *credential.VerifiableCredential) bool {
    issuer := credential.Issuer().String()
    for _, trusted := range trustedIssuers {
        if issuer == trusted {
            return true
        }
    }
    return false
}
```

### 2. Credential Freshness
```go
// Check credential validity period
func validateCredentialFreshness(credential *credential.VerifiableCredential) bool {
    validFrom := credential.ValidFrom()
    now := time.Now()
    
    // Ensure credential is currently valid
    if now.Before(validFrom) {
        return false
    }
    
    // Check if credential is too old (optional)
    maxAge := time.Hour * 24 * 30 // 30 days
    if now.Sub(validFrom) > maxAge {
        return false
    }
    
    return true
}
```

### 3. Presentation Authenticity
```go
// Verify presentation is from expected holder
func validatePresentationHolder(presentation *credential.VerifiablePresentation, expectedHolder string) bool {
    actualHolder := presentation.Holder().String()
    return actualHolder == expectedHolder
}
```

## Integration Patterns

### 1. Web Application Integration
```go
// HTTP endpoint for credential verification
func handleCredentialVerification(w http.ResponseWriter, r *http.Request) {
    // Generate QR code for mobile connection
    qrCode := generateConnectionQR()
    
    // Return QR code to frontend
    json.NewEncoder(w).Encode(map[string]string{
        "qrCode": qrCode,
        "sessionId": sessionId,
    })
}

func handleCredentialResponse(w http.ResponseWriter, r *http.Request) {
    // Process received credentials
    credentials := extractCredentialsFromRequest(r)
    
    // Validate and make access decision
    if validateCredentials(credentials) {
        // Grant access
        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(map[string]bool{
            "verified": true,
        })
    }
}
```

### 2. Mobile App Integration
```go
// Deep link handling for mobile-to-mobile verification
func handleDeepLinkVerification(credentialRequest string) {
    // Parse credential request from deep link
    request := parseCredentialRequest(credentialRequest)
    
    // Show user what's being requested
    showCredentialRequestUI(request)
    
    // If user approves, send credentials
    if userApproves() {
        sendCredentialsToRequester(request)
    }
}
```

### 3. API Gateway Integration
```go
// Middleware for API endpoint protection
func credentialAuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Extract credentials from request
        credentials := extractCredentialsFromHeaders(r)
        
        // Validate credentials
        if !validateCredentials(credentials) {
            http.Error(w, "Invalid credentials", http.StatusUnauthorized)
            return
        }
        
        // Add credential claims to request context
        ctx := context.WithValue(r.Context(), "claims", extractClaims(credentials))
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}
```

## Testing with Real Credentials

### 1. Development Environment
```bash
# Set up sandbox environment for testing
export SELF_ENVIRONMENT=sandbox
export SELF_STORAGE_PATH=./test_storage

# Run with debug logging
export SELF_LOG_LEVEL=debug
go run main.go
```

### 2. Mobile App Testing
1. Install Self mobile app on test device
2. Create test account in sandbox environment
3. Obtain test credentials from sandbox issuers
4. Scan QR codes to test credential exchange

### 3. Integration Testing
```go
func TestRealCredentialFlow(t *testing.T) {
    // This requires actual mobile app interaction
    // Best tested manually or with automation tools
    
    verifier := createTestVerifier()
    qrCode := generateConnectionQR(verifier)
    
    // Manual step: scan QR with mobile app
    fmt.Printf("Scan this QR code: %s\n", qrCode)
    
    // Wait for mobile connection and credential response
    credentials := waitForCredentialResponse(60 * time.Second)
    
    assert.NotEmpty(t, credentials)
    assert.Contains(t, credentials[0].CredentialType(), "EmailCredential")
}
```

## Production Deployment

### 1. Environment Configuration
```go
// Production-ready configuration
config := &account.Config{
    StorageKey:  getSecureStorageKey(), // From secure key management
    StoragePath: "/secure/credential/storage",
    Environment: account.TargetProduction, // Use production network
    LogLevel:    account.LogWarn,
}
```

### 2. Monitoring and Logging
```go
// Track credential verification metrics
type VerificationMetrics struct {
    RequestsReceived    int64
    CredentialsReceived int64
    ValidationFailures  int64
    TrustedIssuers     map[string]int64
}

func trackVerificationEvent(eventType string, details map[string]interface{}) {
    // Send to monitoring system
    metrics.Counter("credential_verification").
        Tag("event", eventType).
        Increment()
}
```

### 3. Security Best Practices
- **Network Security**: Use TLS/HTTPS for all communications
- **Storage Security**: Encrypt credential storage at rest
- **Access Control**: Implement proper authentication for verifier services
- **Audit Logging**: Log all credential verification events
- **Rate Limiting**: Prevent abuse of verification endpoints

## Next Steps

### Enhanced Features
- **Selective Disclosure**: Request only specific claims from credentials
- **Zero-Knowledge Proofs**: Verify claims without revealing actual values
- **Batch Verification**: Process multiple credentials simultaneously
- **Revocation Checking**: Verify credentials haven't been revoked

### Advanced Integration
- **Webhook Notifications**: Real-time credential verification events
- **Database Integration**: Store verification results and audit logs
- **Multi-factor Authentication**: Combine credentials with other auth methods
- **Cross-platform Support**: Support various mobile wallet implementations

## Summary

This example demonstrates real credential verification in action:

### ✅ **Real World Integration**
- **Actual mobile connections** via QR codes
- **Authentic credentials** from trusted issuers  
- **Real email addresses** and biometric proofs
- **Production-ready patterns** for credential verification

### 🔧 **Technical Implementation**
- **Secure communication** with mobile devices
- **Presentation request protocols** for specific credential types
- **Trust validation** against known issuer registries
- **Error handling** for network and validation issues

### 🚀 **Business Applications**
- **Identity verification** for high-value services
- **Age verification** for restricted content
- **Professional verification** for qualified services
- **Access control** based on real credentials

You now have a working example of requesting and validating real credentials from mobile devices - the foundation for building trust-based applications! 🎉 
