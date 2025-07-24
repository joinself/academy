# Mobile Implementation

> ****What you will learn:** What you'll learn:** How the Self mobile app handles age verification through document capture, credential issuance, and privacy-preserving presentation.

This guide covers the mobile application's role in the identity verification process, focusing on how the Self SDK simplifies document verification while maintaining user privacy and security.

## Mobile Verification Workflow

The age verification process from the mobile perspective involves these key phases:

### 1. Initial Connection via QR Code
- **User scans QR code** displayed by the web application
- **Secure channel established** between mobile app and server
- **Connection confirmed** on both ends with cryptographic handshake

### 2. Receiving the Credential Request
- Mobile app receives a **credential presentation request** for `dateOfBirth`
- Request specifies **zero-knowledge verification** parameters (age > 18)
- App checks if user already has the required credential

### 3. Automatic Credential Issuance (If Needed)
If the user doesn't have a date of birth credential, the Self SDK automatically initiates:

#### Document Capture Process
- **Camera activation** with guided document positioning
- **Real-time quality checks** for blur, glare, and document type
- **Multi-frame capture** to ensure optimal image quality
- **Document type detection** (driver's license, passport, etc.)

#### On-Device Processing
- **Text extraction** using OCR technology
- **Data validation** and format checking
- **Privacy-preserving analysis** (data never leaves device unencrypted)
- **User confirmation** of extracted information

#### Credential Issuance
- **Secure submission** to Self's verification services
- **Identity document validation** through multiple checks
- **Verifiable credential creation** with cryptographic proof
- **Local storage** on user's device

### 4. Zero-Knowledge Credential Presentation

The mobile app performs privacy-preserving verification:

```
┌─────────────────────┐    ┌─────────────────────┐    ┌─────────────────────┐
│   Server Request    │    │   Mobile App        │    │   User Device       │
│                     │    │                     │    │                     │
│ "Prove age > 18"    │───►│ 1. Check local      │───►│ Credential Storage  │
│ (zero-knowledge)    │    │    credentials      │    │                     │
│                     │    │                     │    │                     │
│ Receives: true/false│◄───│ 2. Evaluate locally │◄───│ Date of birth:      │
│ (not actual age)    │    │    dateOfBirth      │    │ "1985-03-15"        │
│                     │    │    < cutoff_date    │    │                     │
└─────────────────────┘    └─────────────────────┘    └─────────────────────┘
```

**Key Privacy Features:**
- **Local Evaluation**: Age check performed on device
- **Boolean Response**: Only "verified" or "not verified" sent to server
- **No Data Leakage**: Actual birth date never transmitted
- **Cryptographic Proof**: Response includes tamper-evident signature

### 5. User Experience Flow

#### Happy Path (User has credential)
1. Scan QR code → "Connected"
2. Request received → "Age verification requested"
3. Local check → Age confirmed ≥ 18
4. Send verification → "Access granted"
5. Server confirmation → "Welcome to restricted content"

#### Issuance Path (New user)
1. Scan QR code → "Connected"
2. Request received → "Please verify your identity"
3. Document prompt → "Scan your ID document"
4. Capture process → "Hold steady... Processing..."
5. Data extraction → "Please confirm: Date of birth: MM/DD/YYYY"
6. User approval → "Creating secure credential..."
7. Credential saved → "Identity verified!"
8. Age check → "Access granted"

#### Failure Scenarios
- **Under 18**: "Sorry, you must be 18+ to access this service"
- **Invalid document**: "Document not recognized, please try again"
- **Poor image quality**: "Please retake photo with better lighting"
- **Network issues**: "Connection lost, please scan QR code again"

## Technical Implementation Details

### Document Verification Capabilities

The Self mobile SDK supports various identity documents:

#### Supported Document Types
- **Driver's Licenses**: All US states, EU countries
- **Passports**: International machine-readable travel documents
- **National ID Cards**: Government-issued photo identification
- **State ID Cards**: Non-driver identification cards

#### Security Features
- **Anti-Spoofing**: Detection of printed photos, screens, masks
- **Tamper Detection**: Identification of altered or fraudulent documents
- **Hologram Verification**: Advanced optical security feature validation
- **Biometric Matching**: Face comparison between photo and live selfie

### Privacy Architecture

#### On-Device Processing
```
Document Image → OCR Engine → Data Extraction → User Validation → Encrypted Storage
       ↓              ↓             ↓              ↓               ↓
   Never sent    Local only    User confirms   Cryptographic   Secure vault
   to server     processing    before storage    signing       on device
```

#### Zero-Knowledge Proofs
```go
// Mobile app evaluation (pseudocode)
func evaluateAgeRequirement(birthDate string, cutoffDate string) bool {
    userBirthDate := parseDate(birthDate)
    requiredDate := parseDate(cutoffDate)
    
    // Only the boolean result is shared
    return userBirthDate.Before(requiredDate)
}

// Cryptographic proof generation
proof := generateZKProof(result, userCredential, serverChallenge)
response := createSignedResponse(result, proof)
```

### Credential Management

#### Storage Security
- **Hardware Security Module**: Credentials stored in device secure enclave
- **Encryption**: AES-256 encryption with device-specific keys
- **Access Control**: Biometric or PIN protection for credential access
- **Backup Protection**: Credentials never included in device backups

#### Credential Lifecycle
1. **Issuance**: Document verification → credential creation
2. **Storage**: Secure vault on user device
3. **Presentation**: Zero-knowledge proofs for verification
4. **Expiration**: Automatic credential renewal based on document validity
5. **Revocation**: User can delete credentials at any time

## User Interface Patterns

### Document Capture UI
```
┌─────────────────────────────────┐
│  📷 Scan Your Driver's License   │
│  ─────────────────────────────  │
│                                 │
│    ┌─────────────────────┐      │
│    │                     │      │  ← Camera viewfinder
│    │   [ID Document]     │      │
│    │                     │      │
│    └─────────────────────┘      │
│                                 │
│  • Hold steady                  │
│  • Ensure good lighting         │
│  • Align with guides            │
│  ─────────────────────────────  │
│         [Capture]               │
└─────────────────────────────────┘
```

### Verification Status UI
```
┌─────────────────────────────────┐
│       **Process:** Processing...          │
│  ─────────────────────────────  │
│                                 │
│  Extracting information from    │
│  your document...               │
│                                 │
│  ✓ Document detected            │
│  ✓ Image quality verified       │
│  **Process:** Reading text...             │
│  ⏳ Validating data...          │
│                                 │
│  This may take a few moments    │
└─────────────────────────────────┘
```

### Confirmation UI
```
┌─────────────────────────────────┐
│      ✓ Information Extracted    │
│  ─────────────────────────────  │
│                                 │
│  Please confirm the following:  │
│                                 │
│  Name: John Doe                 │
│  Date of Birth: 03/15/1985      │
│  Document: Driver's License     │
│  State: California              │
│                                 │
│  ⚠️  This will be used for age  │
│     verification only           │
│  ─────────────────────────────  │
│   [Cancel]      [Confirm]       │
└─────────────────────────────────┘
```

## Integration Best Practices

### Error Handling
```swift
// iOS Swift example
func handleDocumentCapture() {
    selfSDK.captureDocument { result in
        switch result {
        case .success(let credential):
            self.presentCredential(credential)
        case .failure(let error):
            switch error {
            case .poorImageQuality:
                self.showRetryWithImprovedLighting()
            case .documentNotSupported:
                self.showSupportedDocumentTypes()
            case .userCancelled:
                self.returnToMainFlow()
            default:
                self.showGenericError()
            }
        }
    }
}
```

### User Guidance
```kotlin
// Android Kotlin example
class DocumentCaptureActivity {
    fun startDocumentCapture() {
        // Show tutorial for first-time users
        if (isFirstTimeUser()) {
            showDocumentCaptureTutorial()
        }
        
        // Configure capture with guidance
        val config = DocumentCaptureConfig.builder()
            .showRealTimeGuidance(true)
            .enableQualityChecks(true)
            .supportedTypes(listOf(DRIVERS_LICENSE, PASSPORT))
            .build()
            
        SelfSDK.captureDocument(config) { result ->
            handleCaptureResult(result)
        }
    }
}
```

## Security Considerations

### Privacy Protection
- **Minimal Data Exposure**: Only required claims shared with servers
- **User Consent**: Explicit approval before any data transmission
- **Selective Disclosure**: Users choose which credentials to present
- **Data Sovereignty**: Users maintain control over their credentials

### Anti-Fraud Measures
- **Liveness Detection**: Ensures real person, not photo/video
- **Document Authentication**: Multiple security feature validation
- **Behavioral Analysis**: Detection of automated or suspicious behavior
- **Geolocation Checks**: Optional location-based verification

### Compliance Features
- **GDPR Compliance**: Right to deletion, data minimization
- **CCPA Compliance**: California privacy law adherence
- **Industry Standards**: W3C Verifiable Credentials, DIF standards
- **Accessibility**: WCAG 2.1 AA compliance for inclusive design

## Next Steps

- **[Server Implementation](./server-implementation.md)**: Understand how the backend processes verification results
- **[Conclusions & Best Practices](./conclusions.md)**: Learn about production deployment and security
- **[Complete Example](../../examples/solutions/age-verifier/)**: See the server-side implementation that works with these mobile flows
- **[Identity Verification Overview](../identity-verification.md)**: Return to the main identity verification guide 
