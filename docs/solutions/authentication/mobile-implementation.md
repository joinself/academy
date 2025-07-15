# Mobile Implementation

> **🎯 What you'll learn:** How to integrate Self authentication into your mobile applications using Self SDKs for iOS and Android.

This guide covers mobile app integration with Self authentication, including SDK setup, QR code scanning, biometric verification, and production deployment considerations.

## Building Your Mobile App

You can build your own mobile application with Self authentication capabilities using our mobile SDKs.

### Current Testing Approach

While mobile SDK examples are being finalized, you can test your authentication backend immediately:

**[Self Developer App](https://www.joinself.com/developers/developers)** - A complete mobile testing tool that:

- Scans authentication QR codes from your backend
- Handles the complete authentication flow  
- Provides biometric verification capabilities
- Works with any Self-enabled authentication backend
- Perfect for development and testing phases

> **💡 Production Ready:** The backend authentication system is production-ready. Once mobile SDK examples are available, you can seamlessly integrate Self authentication into your own mobile applications.

## Mobile SDK Integration _(Coming Soon)_

### iOS Implementation

The iOS SDK will provide:

- Native Swift/Objective-C integration
- QR code scanning capabilities
- Biometric authentication (Face ID/Touch ID)
- Secure credential management
- Production-ready UI components

### Android Implementation

The Android SDK will provide:

- Native Java/Kotlin integration
- Camera-based QR code scanning
- Biometric authentication (Fingerprint/Face unlock)
- Secure storage for credentials
- Material Design components

### Platform-Specific Examples

**iOS Example Structure:**
```
examples/mobile/ios/authentication/
├── basic-auth/           # Simple authentication flow
├── biometric-setup/      # Biometric configuration
├── production-ready/     # Complete production example
└── README.md
```

**Android Example Structure:**
```
examples/mobile/android/authentication/
├── basic-auth/           # Simple authentication flow
├── biometric-setup/      # Biometric configuration  
├── production-ready/     # Complete production example
└── README.md
```

## Mobile Authentication Flow

From the mobile app perspective, the authentication process involves:

1. **QR Code Scanning**
   - Integrate camera for QR code detection
   - Parse Self discovery request from QR data
   - Establish secure connection with backend

2. **User Consent**
   - Display authentication request details
   - Show requesting application information
   - Obtain user confirmation to proceed

3. **Biometric Verification**
   - Request device biometric authentication
   - Generate cryptographic proof of liveness
   - Present verifiable credential to backend

4. **Authentication Completion**
   - Receive authentication result from backend
   - Handle success/failure states appropriately
   - Provide user feedback on authentication status

## SDK Integration Patterns

### Basic Authentication Flow

```swift
// iOS Example (Coming Soon)
import SelfSDK

class AuthenticationViewController: UIViewController {
    func scanQRCode() {
        // QR scanning implementation
    }
    
    func handleAuthRequest(_ request: AuthRequest) {
        // Process authentication request
    }
    
    func performBiometricAuth() {
        // Biometric verification
    }
}
```

```kotlin
// Android Example (Coming Soon)
import com.joinself.sdk

class AuthenticationActivity : AppCompatActivity() {
    fun scanQRCode() {
        // QR scanning implementation
    }
    
    fun handleAuthRequest(request: AuthRequest) {
        // Process authentication request
    }
    
    fun performBiometricAuth() {
        // Biometric verification
    }
}
```

## Next Steps

- **[Server Implementation](./server-implementation.md)**: Set up your authentication backend  
- **[Conclusions & Best Practices](./conclusions.md)**: Production deployment guide, security best practices, and development workflow
- **[Authentication Overview](../authentication.md)**: Return to the main authentication guide 
