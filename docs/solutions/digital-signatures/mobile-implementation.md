# Digital Signatures: Mobile Implementation

> **🎯 What you'll learn:** How to integrate digital signature capabilities into your mobile app with agreement review, biometric approval, and credential verification response handling.

This guide shows you how to build mobile applications that can receive signing requests, display agreement details, and approve digital signatures with biometric authentication.

## Architecture Overview

The mobile digital signature implementation handles **agreement review and approval** with three core components:

**Request Reception**
- Receives credential verification requests from server
- Parses agreement details and evidence
- Displays agreement information to user

**User Review**
- Shows agreement terms and conditions
- Displays document metadata and hash
- Provides clear approval/rejection options

**Biometric Approval**
- Handles biometric authentication for signature approval
- Sends credential verification response with signed credential
- Maintains user privacy and control

## Core Implementation

### 1. Receiving Signing Requests

The mobile app receives credential verification requests containing agreement details and evidence.

**Android Implementation:**
```kotlin
// Handle credential verification requests
override fun onCredentialVerificationRequest(request: CredentialVerificationRequest) {
    // Extract agreement details from request
    val evidence = request.evidence["agreement"]
    val agreementType = request.type.firstOrNull { it.contains("AgreementCredential") }
    
    // Parse agreement metadata
    val agreementDetails = parseAgreementDetails(evidence)
    
    // Show agreement review UI
    showAgreementReview(agreementDetails)
}
```

**iOS Implementation:**
```swift
// Handle credential verification requests
func handleCredentialVerificationRequest(_ request: CredentialVerificationRequest) {
    // Extract agreement details from request
    guard let evidence = request.evidence["agreement"],
          let agreementType = request.type.first(where: { $0.contains("AgreementCredential") }) else {
        return
    }
    
    // Parse agreement metadata
    let agreementDetails = parseAgreementDetails(evidence)
    
    // Show agreement review UI
    showAgreementReview(agreementDetails)
}
```

### 2. Agreement Review Interface

The mobile app displays agreement details for user review before approval.

**Android UI Example:**
```kotlin
private fun showAgreementReview(agreementDetails: AgreementDetails) {
    val intent = Intent(this, AgreementReviewActivity::class.java).apply {
        putExtra("agreement_type", agreementDetails.type)
        putExtra("agreement_id", agreementDetails.id)
        putExtra("document_hash", agreementDetails.documentHash)
        putExtra("signing_date", agreementDetails.signingDate)
        putExtra("tenant_name", agreementDetails.tenantName)
        putExtra("property_address", agreementDetails.propertyAddress)
        putExtra("rent_amount", agreementDetails.rentAmount)
    }
    startActivity(intent)
}
```

**iOS UI Example:**
```swift
private func showAgreementReview(_ agreementDetails: AgreementDetails) {
    let reviewVC = AgreementReviewViewController()
    reviewVC.agreementDetails = agreementDetails
    reviewVC.delegate = self
    present(reviewVC, animated: true)
}
```

### 3. Biometric Signature Approval

The mobile app handles biometric authentication for signature approval.

**Android Biometric Implementation:**
```kotlin
private fun approveSignatureWithBiometrics(agreementDetails: AgreementDetails) {
    val biometricPrompt = BiometricPrompt.PromptInfo.Builder()
        .setTitle("Approve Digital Signature")
        .setSubtitle("Sign the tenancy agreement with biometric authentication")
        .setNegativeButtonText("Cancel")
        .build()

    biometricPrompt.authenticate(biometricPrompt) { result ->
        when (result.authenticationResult) {
            BiometricPrompt.AuthenticationResult.SUCCESS -> {
                // Send credential verification response
                sendCredentialVerificationResponse(agreementDetails)
            }
            BiometricPrompt.AuthenticationResult.ERROR -> {
                // Handle biometric error
                handleBiometricError(result.errorCode)
            }
        }
    }
}
```

**iOS Biometric Implementation:**
```swift
private func approveSignatureWithBiometrics(_ agreementDetails: AgreementDetails) {
    let context = LAContext()
    var error: NSError?
    
    if context.canEvaluatePolicy(.deviceOwnerAuthenticationWithBiometrics, error: &error) {
        context.evaluatePolicy(.deviceOwnerAuthenticationWithBiometrics,
                              localizedReason: "Approve digital signature for tenancy agreement") { success, error in
            DispatchQueue.main.async {
                if success {
                    // Send credential verification response
                    self.sendCredentialVerificationResponse(agreementDetails)
                } else {
                    // Handle biometric error
                    self.handleBiometricError(error)
                }
            }
        }
    }
}
```

### 4. Sending Credential Verification Response

The mobile app sends the credential verification response with the signed credential.

**Android Response Implementation:**
```kotlin
private fun sendCredentialVerificationResponse(agreementDetails: AgreementDetails) {
    // Create credential verification response
    val response = CredentialVerificationResponse.Builder()
        .responseTo(requestId)
        .addPresentation(signedCredentialPresentation)
        .build()
    
    // Send response to server
    selfAccount.sendMessage(serverDID, response) { result ->
        when (result) {
            is Result.Success -> {
                showSignatureComplete()
            }
            is Result.Error -> {
                showSignatureError(result.exception)
            }
        }
    }
}
```

**iOS Response Implementation:**
```swift
private func sendCredentialVerificationResponse(_ agreementDetails: AgreementDetails) {
    // Create credential verification response
    let response = CredentialVerificationResponse(
        responseTo: requestId,
        presentations: [signedCredentialPresentation]
    )
    
    // Send response to server
    selfAccount.sendMessage(to: serverDID, content: response) { result in
        DispatchQueue.main.async {
            switch result {
            case .success:
                self.showSignatureComplete()
            case .failure(let error):
                self.showSignatureError(error)
            }
        }
    }
}
```

## User Experience Design

### Agreement Review Screen

**Key UI Elements:**
- **Agreement Type**: Clear indication of document type (e.g., "Tenancy Agreement")
- **Document Hash**: Cryptographic hash for verification
- **Signing Date**: Timestamp of signature
- **Agreement Details**: Formatted display of key terms
- **Approve/Reject Buttons**: Clear action options
- **Biometric Prompt**: Secure authentication flow

**Android Layout Example:**
```xml
<LinearLayout android:orientation="vertical">
    <TextView android:text="Tenancy Agreement" style="@style/AgreementTitle"/>
    <TextView android:text="Document Hash: a1b2c3..." style="@style/AgreementHash"/>
    <TextView android:text="Signing Date: 2024-02-01" style="@style/AgreementDate"/>
    
    <ScrollView>
        <TextView android:text="[Agreement Terms]" style="@style/AgreementTerms"/>
    </ScrollView>
    
    <Button android:text="Approve with Biometrics" 
            android:onClick="approveSignature"/>
    <Button android:text="Reject" 
            android:onClick="rejectSignature"/>
</LinearLayout>
```

**iOS Layout Example:**
```swift
class AgreementReviewViewController: UIViewController {
    @IBOutlet weak var agreementTitleLabel: UILabel!
    @IBOutlet weak var documentHashLabel: UILabel!
    @IBOutlet weak var signingDateLabel: UILabel!
    @IBOutlet weak var agreementTermsTextView: UITextView!
    @IBOutlet weak var approveButton: UIButton!
    @IBOutlet weak var rejectButton: UIButton!
    
    override func viewDidLoad() {
        super.viewDidLoad()
        setupAgreementDisplay()
    }
}
```

### Signature Confirmation Flow

**Success State:**
- Clear confirmation message
- Agreement details summary
- Option to view signed document
- Return to main app flow

**Error State:**
- Clear error message
- Retry option
- Fallback authentication methods
- Support contact information

## Security Considerations

### Biometric Security

**Best Practices:**
- Use device-native biometric authentication
- Implement proper error handling for biometric failures
- Provide fallback authentication methods
- Never store biometric data locally

**Android Security:**
```kotlin
private fun setupBiometricSecurity() {
    val keyGenerator = KeyGenerator.getInstance(KeyProperties.KEY_ALGORITHM_AES, "AndroidKeyStore")
    val keyGenParameterSpec = KeyGenParameterSpec.Builder(
        "signature_key",
        KeyProperties.PURPOSE_ENCRYPT or KeyProperties.PURPOSE_DECRYPT
    )
    .setBlockModes(KeyProperties.BLOCK_MODE_GCM)
    .setEncryptionPaddings(KeyProperties.ENCRYPTION_PADDING_NONE)
    .setUserAuthenticationRequired(true)
    .setUserAuthenticationValidityDurationSeconds(30)
    .build()
    
    keyGenerator.init(keyGenParameterSpec)
    keyGenerator.generateKey()
}
```

**iOS Security:**
```swift
private func setupBiometricSecurity() {
    let context = LAContext()
    context.localizedFallbackTitle = "Use Passcode"
    context.localizedCancelTitle = "Cancel"
    
    // Configure biometric policy
    let policy = LAPolicy.deviceOwnerAuthenticationWithBiometrics
    context.evaluatePolicy(policy, localizedReason: "Approve digital signature") { success, error in
        // Handle authentication result
    }
}
```

### Data Privacy

**Privacy Protection:**
- Never store agreement content locally
- Clear sensitive data after processing
- Use secure communication channels
- Implement proper session management

**Android Privacy:**
```kotlin
private fun clearSensitiveData() {
    // Clear agreement details from memory
    agreementDetails = null
    signedCredentialPresentation = null
    
    // Clear any cached data
    clearCache()
    
    // Force garbage collection for sensitive objects
    System.gc()
}
```

**iOS Privacy:**
```swift
private func clearSensitiveData() {
    // Clear agreement details from memory
    agreementDetails = nil
    signedCredentialPresentation = nil
    
    // Clear any cached data
    URLCache.shared.removeAllCachedResponses()
    
    // Force memory cleanup
    autoreleasepool {
        // Sensitive data will be deallocated
    }
}
```

## Testing Strategies

### Unit Testing

**Request Processing Tests:**
```kotlin
@Test
fun testCredentialVerificationRequestProcessing() {
    // Create mock request
    val mockRequest = createMockCredentialVerificationRequest()
    
    // Test request parsing
    val agreementDetails = parseAgreementDetails(mockRequest)
    
    // Verify agreement details
    assertEquals("Tenancy Agreement", agreementDetails.type)
    assertNotNull(agreementDetails.documentHash)
    assertNotNull(agreementDetails.tenantName)
}
```

**Biometric Testing:**
```kotlin
@Test
fun testBiometricApprovalFlow() {
    // Mock biometric success
    whenever(biometricManager.authenticate(any())).thenReturn(SUCCESS)
    
    // Test approval flow
    val result = approveSignatureWithBiometrics(mockAgreementDetails)
    
    // Verify response was sent
    verify(mockSelfAccount).sendMessage(any(), any())
    assertTrue(result.isSuccess)
}
```

### Integration Testing

**End-to-End Signing Tests:**
```kotlin
@Test
fun testCompleteSigningWorkflow() {
    // Simulate receiving signing request
    val request = createSigningRequest()
    onCredentialVerificationRequest(request)
    
    // Verify agreement review UI is shown
    verify(mockActivity).startActivity(any<Intent>())
    
    // Simulate user approval
    approveSignatureWithBiometrics(agreementDetails)
    
    // Verify response is sent
    verify(mockSelfAccount).sendMessage(serverDID, any())
}
```

## Production Considerations

### Performance Optimization

**Memory Management:**
- Efficiently handle large agreement documents
- Implement proper image caching for UI elements
- Use background processing for heavy operations
- Monitor memory usage during signing flows

**Network Optimization:**
- Implement retry logic for failed requests
- Use connection pooling for server communication
- Compress data in transit when possible
- Handle offline scenarios gracefully

### Error Handling

**Comprehensive Error Handling:**
```kotlin
private fun handleSigningError(error: Throwable) {
    when (error) {
        is NetworkException -> showNetworkError()
        is BiometricException -> showBiometricError()
        is CredentialException -> showCredentialError()
        else -> showGenericError()
    }
    
    // Log error for debugging
    Log.e(TAG, "Signing error: ${error.message}", error)
    
    // Report to analytics
    analytics.reportError("signing_error", error)
}
```

### Analytics and Monitoring

**User Experience Tracking:**
```kotlin
private fun trackSigningFlow() {
    analytics.trackEvent("signing_request_received", mapOf(
        "agreement_type" to agreementDetails.type,
        "user_id" to userId
    ))
    
    analytics.trackEvent("signing_approved", mapOf(
        "agreement_type" to agreementDetails.type,
        "signing_method" to "biometric",
        "processing_time" to processingTime
    ))
}
```

## Next Steps

**Ready to implement mobile digital signatures?**
- **Complete Integration**: Follow the full implementation guide with real credential processing
- **Custom UI**: Adapt the agreement review interface for your app design
- **Advanced Features**: Add document preview, multi-language support, accessibility features
- **Production Deployment**: Implement monitoring, analytics, and error tracking

**Learn More:**
- **[Server Implementation](./server-implementation.md)**: Build your digital signature backend
- **[Conclusions & Best Practices](./conclusions.md)**: Production deployment guide
- **[Authentication](./../authentication/mobile-implementation.md)**: Mobile authentication implementation
- **[Credential Exchange](../examples/credentials.md)**: Learn about credential handling patterns 
