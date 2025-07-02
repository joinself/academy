# Email Credential Verification 📧

🎯 **What you'll learn:**
- Email verification credential workflows
- Real-world credential exchange scenarios
- Service-to-user credential issuance patterns
- Mobile credential delivery concepts
- Email domain verification use cases

📚 **Prerequisites:** Complete foundation patterns first:
- `../presentation_request/` - Basic exchange patterns
- `../../01_issuing_credentials/01_basic/` - Credential creation fundamentals

## Overview

This example demonstrates email verification through verifiable credentials - a common real-world use case where services need to verify users' email addresses and provide cryptographic proof of verification.

## Key Concepts

### Email Verification Workflow
Traditional email verification relies on centralized systems and temporary tokens. Verifiable credentials provide:
- **Permanent proof** - Credentials don't expire like email tokens
- **Portable verification** - Users own their verification proof
- **Cryptographic integrity** - Tamper-evident verification status
- **Privacy preservation** - No need to contact original verifier

### Real-world Applications
- **Account Registration**: Prove email ownership during signup
- **Identity Verification**: Email as part of multi-factor authentication
- **Service Access**: Email verification for premium features
- **Compliance**: Regulatory requirements for verified communication

## Code Breakdown

### 1. Verification Party Setup
```go
type VerificationParty struct {
    name    string
    account *account.Account
}

func createVerificationParties() (*VerificationParty, *VerificationParty) {
    // Create email service issuer
    issuer := &VerificationParty{name: "Email Service Provider"}
    issuerCfg := &account.Config{
        StorageKey:  generateStorageKey("email_issuer"),
        StoragePath: "./email_issuer_storage",
        Environment: account.TargetSandbox,
        LogLevel:    account.LogWarn,
    }
    
    // Create user holder
    holder := &VerificationParty{name: "User"}
    // ... similar configuration
}
```

**Party roles:**
- **Email Service Provider**: Issues verification credentials after email validation
- **User**: Receives and stores email verification credentials
- **Verifying Services**: Accept credentials as proof of email ownership

### 2. Email Verification Credential
```go
func createEmailVerificationCredential(issuerAccount *account.Account, issuerAddress, holderAddress *signing.PublicKey) *credential.VerifiableCredential {
    claims := map[string]interface{}{
        "emailAddress":       "user@example.com",
        "verified":          true,
        "verificationDate":  time.Now().Format("2006-01-02"),
        "domain":            "example.com",
        "verificationMethod": "email_link",
        "issuerName":        "Example Email Service",
    }
    
    credentialBuilder := credential.NewCredential().
        CredentialType([]string{"VerifiableCredential", "EmailCredential"}).
        CredentialSubject(credential.AddressKey(holderAddress)).
        CredentialSubjectClaims(claims).
        Issuer(credential.AddressKey(issuerAddress)).
        ValidFrom(time.Now()).
        SignWith(issuerAddress, time.Now())
}
```

**Credential claims:**
- **emailAddress**: The verified email address
- **verified**: Boolean verification status
- **verificationDate**: When verification occurred
- **domain**: Email domain for organizational verification
- **verificationMethod**: How verification was performed
- **issuerName**: Human-readable issuer identification

### 3. Production Delivery Concepts
```go
fmt.Println("📱 In production:")
fmt.Println("   • Credential would be sent via QR code or deep link")
fmt.Println("   • Mobile app would receive and store credential")
fmt.Println("   • User can present verification proof when needed")
```

**Delivery mechanisms:**
- **QR Code**: User scans code to receive credential
- **Deep Link**: Direct app-to-app credential transfer
- **Push Notification**: Alert user to pending credential
- **Email Link**: Ironically, email can deliver email verification credentials

## Running the Example

```bash
cd examples/go/02_credentials/02_exchanging_credentials/email_verification
go run main.go
```

## Expected Output

```
Email Credential Verification Demo
=================================
Setting up verification parties...
Email Service: 1:ABC123...
User: 1:DEF456...

Email verification workflow:
Scenario: User verifies email for service account
Creating email verification credential...
✅ Email verification credential created
📧 Credential details:
   • Email: user@example.com
   • Status: verified
   • Domain: example.com
   • Method: email_link

📱 In production:
   • Credential would be sent via QR code or deep link
   • Mobile app would receive and store credential
   • User can present verification proof when needed
✅ Email verification demo completed!
```

## What Just Happened

1. **Created verification parties** - Email service provider and user with separate identities
2. **Issued email credential** - Service creates verifiable proof of email verification
3. **Packaged verification data** - Multiple claims bundled into tamper-evident credential
4. **Demonstrated delivery concepts** - How credentials reach mobile devices in production

## Email Verification Architecture Patterns

### 1. Traditional vs. Credential-Based Verification

#### Traditional Email Verification
```
User → Service → Email Token → User Clicks → Service Validates → Temporary Status
```
**Issues:**
- Tokens expire and become useless
- No portable proof of verification
- Requires contacting original service for validation
- Vulnerable to email interception

#### Credential-Based Verification
```
User → Service → Email Verification → Credential Issued → Permanent Portable Proof
```
**Benefits:**
- Permanent, reusable verification proof
- Cryptographically secured and tamper-evident
- No need to contact original verifier
- User owns and controls their verification

### 2. Verification Claim Structure
```go
// Basic email verification
claims := map[string]interface{}{
    "emailAddress": "user@example.com",
    "verified":     true,
}

// Enhanced verification with metadata
claims := map[string]interface{}{
    "emailAddress":       "user@example.com",
    "verified":          true,
    "verificationDate":  "2024-01-15",
    "domain":            "example.com",
    "verificationMethod": "email_link",
    "domainOwnership":   "verified", // For business emails
    "mxRecordCheck":     true,       // Technical validation
    "disposableEmail":   false,      // Anti-fraud measure
}
```

### 3. Multi-Level Verification
```go
// Personal email verification
personalClaims := map[string]interface{}{
    "emailAddress": "john@gmail.com",
    "verified":     true,
    "level":        "personal",
}

// Business email verification  
businessClaims := map[string]interface{}{
    "emailAddress": "john@company.com",
    "verified":     true,
    "level":        "business",
    "organization": "Company Inc.",
    "domainControl": "verified",
}
```

## Real-World Implementation Scenarios

### 1. SaaS Platform Registration
**Use case**: Users sign up for SaaS service and need verified email

**Traditional flow:**
1. User enters email
2. Service sends verification email
3. User clicks link
4. Service marks email as verified
5. **Problem**: Verification tied to this specific service

**Credential flow:**
1. User enters email
2. Service sends verification email
3. User clicks link
4. Service issues verifiable email credential
5. **Benefit**: User can reuse verification across services

### 2. Financial Services KYC
**Use case**: Bank requires verified email for compliance

**Credential approach:**
```go
kycEmailClaims := map[string]interface{}{
    "emailAddress":    "customer@email.com",
    "verified":       true,
    "verificationDate": "2024-01-15",
    "kycCompliant":   true,
    "verifierLicense": "FIN-12345",
    "complianceLevel": "enhanced",
}
```

**Benefits:**
- Permanent compliance record
- Auditable verification trail
- Transferable between compliant institutions
- Reduces repeated KYC friction

### 3. Educational Institution Access
**Use case**: University provides email credentials to students

**Implementation:**
```go
academicEmailClaims := map[string]interface{}{
    "emailAddress":   "student@university.edu",
    "verified":      true,
    "institution":   "State University",
    "studentStatus": "enrolled",
    "academicYear":  "2024",
    "emailType":     "institutional",
}
```

**Applications:**
- Student discount verification
- Academic software access
- Research collaboration proof
- Alumni status verification

## Advanced Patterns

### 1. Verification Levels
```go
// Different verification rigor levels
type VerificationLevel struct {
    Level       string   `json:"level"`
    Methods     []string `json:"methods"`
    Confidence  int      `json:"confidence"`
    Requirements []string `json:"requirements"`
}

basicVerification := VerificationLevel{
    Level:       "basic",
    Methods:     []string{"email_link"},
    Confidence:  70,
    Requirements: []string{"valid_email_format", "deliverable"},
}

enhancedVerification := VerificationLevel{
    Level:       "enhanced", 
    Methods:     []string{"email_link", "sms_backup", "document_proof"},
    Confidence:  95,
    Requirements: []string{"mx_record_valid", "domain_ownership", "identity_document"},
}
```

### 2. Domain-Based Verification
```go
// Verify email domain ownership for business use
func createDomainVerifiedEmailCredential(domain string, adminEmail string) *credential.VerifiableCredential {
    claims := map[string]interface{}{
        "emailAddress":     adminEmail,
        "domain":          domain,
        "domainOwnership": "verified",
        "verificationMethod": "dns_txt_record",
        "adminRole":       true,
        "canIssueForDomain": true,
    }
    
    // This admin can now issue credentials for other emails in this domain
}
```

### 3. Expiring vs. Permanent Verification
```go
// Temporary verification (expires)
temporaryCredential := credential.NewCredential().
    ValidFrom(time.Now()).
    ValidUntil(time.Now().Add(30 * 24 * time.Hour)). // 30 days
    Claim("temporary", true)

// Permanent verification (until revoked)
permanentCredential := credential.NewCredential().
    ValidFrom(time.Now()).
    // No ValidUntil = permanent
    Claim("permanent", true)
```

## Security Considerations

### 1. Email Hijacking Protection
```go
// Include verification metadata to prevent replay attacks
claims := map[string]interface{}{
    "emailAddress":     "user@example.com",
    "verificationCode": "ABC-123-XYZ", // One-time code used
    "clientIP":        "192.168.1.1",  // Where verification occurred
    "userAgent":       "Mozilla/5.0...", // Browser fingerprint
    "timestamp":       time.Now().Unix(),
}
```

### 2. Domain Validation
```go
// Validate email domain before issuing credential
func validateEmailDomain(email string) error {
    domain := strings.Split(email, "@")[1]
    
    // Check MX records
    mxRecords, err := net.LookupMX(domain)
    if err != nil || len(mxRecords) == 0 {
        return errors.New("invalid email domain")
    }
    
    // Check against disposable email list
    if isDisposableEmailDomain(domain) {
        return errors.New("disposable email not allowed")
    }
    
    return nil
}
```

### 3. Verification Rate Limiting
```go
// Prevent abuse of verification system
type VerificationLimiter struct {
    attempts map[string]int
    lastAttempt map[string]time.Time
}

func (vl *VerificationLimiter) AllowVerification(email string) bool {
    if vl.attempts[email] >= 5 {
        return time.Since(vl.lastAttempt[email]) > time.Hour
    }
    return true
}
```

## Integration Patterns

### 1. REST API Integration
```go
// Email verification endpoint
func handleEmailVerification(w http.ResponseWriter, r *http.Request) {
    var req EmailVerificationRequest
    json.NewDecoder(r.Body).Decode(&req)
    
    // Validate email format and domain
    if err := validateEmail(req.Email); err != nil {
        http.Error(w, "Invalid email", http.StatusBadRequest)
        return
    }
    
    // Send verification email
    verificationCode := generateVerificationCode()
    sendVerificationEmail(req.Email, verificationCode)
    
    // Store pending verification
    pendingVerifications[req.Email] = verificationCode
    
    w.WriteHeader(http.StatusAccepted)
    json.NewEncoder(w).Encode(map[string]string{
        "status": "verification_sent",
        "email":  req.Email,
    })
}

// Verification confirmation endpoint
func handleEmailConfirmation(w http.ResponseWriter, r *http.Request) {
    var req EmailConfirmationRequest
    json.NewDecoder(r.Body).Decode(&req)
    
    // Validate verification code
    if pendingVerifications[req.Email] != req.Code {
        http.Error(w, "Invalid code", http.StatusBadRequest)
        return
    }
    
    // Issue credential
    credential, err := issueEmailCredential(req.Email, req.UserDID)
    if err != nil {
        http.Error(w, "Failed to issue credential", http.StatusInternalServerError)
        return
    }
    
    // Return credential or delivery confirmation
    json.NewEncoder(w).Encode(map[string]interface{}{
        "status": "verified",
        "credential": credential,
    })
}
```

### 2. Mobile App Integration
```go
// QR code generation for mobile delivery
func generateEmailCredentialQR(credential *credential.VerifiableCredential) (string, error) {
    // Package credential for mobile delivery
    delivery := CredentialDelivery{
        Type:       "email_verification",
        Credential: credential,
        Issuer:     "Email Service Provider",
        Timestamp:  time.Now(),
    }
    
    // Encode as QR code
    qrData, err := json.Marshal(delivery)
    if err != nil {
        return "", err
    }
    
    return generateQRCode(string(qrData))
}

// Deep link generation
func generateEmailCredentialDeepLink(credential *credential.VerifiableCredential) string {
    credentialJSON, _ := json.Marshal(credential)
    encodedCredential := base64.URLEncoding.EncodeToString(credentialJSON)
    
    return fmt.Sprintf("selfapp://credential/receive?data=%s&type=email", encodedCredential)
}
```

### 3. Webhook Integration
```go
// Notify services when email verification completes
func notifyEmailVerified(email string, credential *credential.VerifiableCredential) {
    notification := EmailVerificationNotification{
        Email:      email,
        Verified:   true,
        Timestamp:  time.Now(),
        Credential: credential,
    }
    
    // Send to configured webhooks
    for _, webhook := range configuredWebhooks {
        go sendWebhookNotification(webhook, notification)
    }
}
```

## Testing Strategies

### 1. Unit Testing Email Validation
```go
func TestEmailValidation(t *testing.T) {
    testCases := []struct {
        email    string
        valid    bool
        reason   string
    }{
        {"user@example.com", true, "valid email"},
        {"user@disposable.com", false, "disposable email"},
        {"invalid-email", false, "invalid format"},
        {"user@nonexistent.domain", false, "no MX record"},
    }
    
    for _, tc := range testCases {
        err := validateEmail(tc.email)
        if tc.valid && err != nil {
            t.Errorf("Expected %s to be valid: %v", tc.email, err)
        }
        if !tc.valid && err == nil {
            t.Errorf("Expected %s to be invalid", tc.email)
        }
    }
}
```

### 2. Integration Testing
```go
func TestEmailVerificationFlow(t *testing.T) {
    // Setup test email service
    issuer, holder := createTestVerificationParties()
    
    // Test credential creation
    credential := createEmailVerificationCredential(issuer.account, issuerAddress, holderAddress)
    assert.NotNil(t, credential)
    
    // Verify credential claims
    claims := credential.CredentialSubjectClaims()
    assert.Equal(t, "user@example.com", claims["emailAddress"])
    assert.Equal(t, true, claims["verified"])
    
    // Test credential validation
    isValid := validateCredentialSignature(credential, issuer.account)
    assert.True(t, isValid)
}
```

### 3. End-to-End Testing
```go
func TestEmailVerificationE2E(t *testing.T) {
    // Start test email server
    emailServer := startTestEmailServer()
    defer emailServer.Close()
    
    // Request email verification
    response := requestEmailVerification("test@example.com")
    assert.Equal(t, "verification_sent", response.Status)
    
    // Simulate email click
    verificationCode := extractCodeFromTestEmail()
    confirmResponse := confirmEmailVerification("test@example.com", verificationCode)
    
    // Verify credential issued
    assert.Equal(t, "verified", confirmResponse.Status)
    assert.NotNil(t, confirmResponse.Credential)
}
```

## Troubleshooting

### Common Issues

**"Invalid email domain"**
- Check email format with regex validation
- Verify domain has valid MX records
- Ensure domain is not in disposable email list
- Test DNS resolution for the domain

**"Failed to issue credential"**
- Verify issuer account is properly configured
- Check that all required claims are provided
- Ensure signing key is valid and accessible
- Validate claim data types match schema

**"Credential delivery failed"**
- Test network connectivity for delivery mechanism
- Verify recipient identity is correct
- Check mobile app can receive credentials
- Ensure QR code or deep link is properly formatted

### Debugging Tips
- Enable debug logging: `LogLevel: account.LogDebug`
- Log all verification attempts and outcomes
- Trace credential creation process step by step
- Validate email sending infrastructure

## Production Considerations

### Performance Optimization
- **Batch verification**: Process multiple emails together
- **Caching**: Cache domain validation results
- **Queue processing**: Async credential generation
- **Rate limiting**: Prevent verification spam

### Security Best Practices
- **SMTP security**: Use encrypted email transport
- **Code expiration**: Time-limited verification codes
- **Attempt limiting**: Prevent brute force attacks
- **Domain blacklisting**: Block known bad domains

### Scalability Patterns
- **Microservices**: Separate verification and credential services
- **Database sharding**: Distribute verification data
- **CDN delivery**: Fast global credential delivery
- **Load balancing**: Handle high verification volumes

## Next Steps

### Enhanced Features
- **Multi-email verification**: Support multiple email addresses per user
- **Email forwarding detection**: Identify and handle forwarded emails
- **Corporate domain verification**: Bulk verification for organizations
- **Email alias support**: Handle plus addressing and aliases

### Integration Opportunities
- **OAuth provider integration**: Link with Google, Microsoft email
- **Enterprise SSO**: Connect with corporate email systems
- **Regulatory compliance**: GDPR, CCPA, financial regulations
- **Cross-platform sync**: Share credentials across devices

### Advanced Applications
- **Reputation systems**: Email domain trust scoring
- **Fraud detection**: Anomaly detection in verification patterns
- **Privacy preservation**: Zero-knowledge email verification
- **Blockchain integration**: Immutable verification records

## Summary

Email credential verification transforms traditional email validation into portable, cryptographic proof:

### ✅ **Core Benefits**
- **Permanent verification**: Doesn't expire like email tokens
- **Portable proof**: Reusable across different services
- **Cryptographic integrity**: Tamper-evident verification status
- **User ownership**: Users control their verification credentials

### 🏗️ **Technical Patterns**
- **Service-to-user issuance**: Email providers issue verification credentials
- **Rich claim structure**: Multiple verification metadata fields
- **Flexible delivery**: QR codes, deep links, push notifications
- **Security considerations**: Domain validation, rate limiting, fraud prevention

### 🚀 **Real-world Applications**
- **Account registration**: Streamlined signup processes
- **KYC compliance**: Financial services verification
- **Academic verification**: Student and alumni access
- **Enterprise integration**: Corporate email validation

You now understand how to implement robust email verification using verifiable credentials that provide lasting value to users and services! 
