# Digital Signatures: Mobile Implementation

> **🎯 What you'll learn:** How to integrate digital signature capabilities into your mobile app with agreement review, biometric approval, and credential verification response handling.

This guide covers mobile app integration with Self digital signatures, including SDK setup, agreement review interfaces, biometric verification, and production deployment considerations.

## Mobile Digital Signature Flow

From the mobile app perspective, the digital signature process involves:

1. **Request Reception**
    - Receive credential verification requests from server
    - Parse agreement details and evidence
    - Display agreement information to user

2. **Agreement Review**
    - Show agreement terms and conditions
    - Display document metadata and hash
    - Provide clear approval/rejection options

3. **Biometric Approval**
    - Request device biometric authentication
    - Generate cryptographic proof of approval
    - Send credential verification response to backend

4. **Signature Completion**
    - Receive signature confirmation from backend
    - Handle success/failure states appropriately
    - Provide user feedback on signature status

## Current Testing Approach

While mobile SDK examples are being finalized, you can test your digital signature backend immediately:

**[Self Developer App](https://www.joinself.com/developers/developers)** - A complete mobile testing tool that:

- Receives signing requests from your backend
- Handles the complete digital signature flow
- Provides biometric verification capabilities
- Works with any Self-enabled digital signature backend
- Perfect for development and testing phases

> **💡 Production Ready:** The backend digital signature system is production-ready. Once mobile SDK examples are available, you can seamlessly integrate Self digital signatures into your own mobile applications.

## Mobile SDK Integration _(Coming Soon)_

> Coming soon

## Next Steps

- **[Server Implementation](./server-implementation.md)**: Set up your digital signature backend  
- **[Conclusions & Best Practices](./conclusions.md)**: Production deployment guide, security best practices, and development workflow
- **[Digital Signatures Overview](../digital-signatures.md)**: Return to the main digital signatures guide 
