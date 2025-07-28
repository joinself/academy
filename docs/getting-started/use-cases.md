# Use Cases

Self provides three core use cases that demonstrate the power of decentralized identity technology. Each use case is designed to solve real-world problems with secure, privacy-preserving solutions.

## Authentication

Self provides a secure authentication workflow that allows users to authenticate themselves using their digital identity. This workflow is designed to be scalable and can be used in a variety of scenarios, including web applications, mobile apps, and enterprise systems.

### Key Features

- **Passwordless Authentication**: Replace traditional passwords with biometric credentials
- **Liveness Detection**: Prevent spoofing attacks with advanced biometric verification
- **Privacy-Preserving**: No personal data stored on servers
- **Cross-Platform**: Works on iOS, Android, and Web applications
- **Scalable**: Designed for both small apps and enterprise systems

### How It Works

1. **User Registration**: Users create a Self account with biometric verification
2. **Credential Issuance**: The system issues liveness credentials to the user
3. **Authentication Request**: Server requests liveness verification from user
4. **Biometric Check**: User completes liveness check on their device
5. **Credential Presentation**: User presents verified credentials to server
6. **Verification**: Server validates credentials and grants access

### Use Cases

- **Web Application Login**: Secure login for SaaS applications
- **Mobile App Authentication**: Passwordless access to mobile apps
- **Enterprise Access Control**: Secure access to corporate systems
- **Financial Services**: Banking and payment authentication
- **Healthcare Applications**: Secure access to patient data

[**Learn More →**](../solutions/authentication/)

## Identity Verification

The Self SDK provides streamlined identity verification flows for mobile applications. These flows support the capture and validation of various identity documents including passports, ID cards, and driving licenses.

### Key Features

- **NFC Chip Reading**: Enhanced passport verification with NFC technology
- **Pre-built UI Components**: Ready-to-use document capture interfaces
- **Secure Local Storage**: Verified credentials stored securely on device
- **Multiple Document Types**: Support for passports, ID cards, driving licenses
- **Real-time Verification**: Instant validation of document authenticity

### How It Works

1. **Document Capture**: User captures identity document using device camera
2. **NFC Verification**: For passports, NFC chip is read for additional verification
3. **Document Analysis**: AI-powered analysis of document authenticity
4. **Credential Creation**: Verified identity information converted to credentials
5. **Secure Storage**: Credentials stored locally on user's device
6. **Selective Sharing**: User controls what information to share

### Use Cases

- **KYC Processes**: Know Your Customer verification for financial services
- **Age Verification**: Prove age without revealing exact date of birth
- **Employment Verification**: Verify work eligibility and qualifications
- **Travel Applications**: Passport verification for travel services
- **Government Services**: Identity verification for public services

[**Learn More →**](../solutions/identity-verification/)

## Digital Signatures

Self enables secure digital document signing workflows that provide legally-compliant electronic signatures. Users can cryptographically sign documents using their verified digital identity, ensuring document integrity and non-repudiation.

### Key Features

- **Cryptographic Signing**: Documents signed with verified digital identity
- **Tamper-Proof**: Cryptographic verification prevents document alteration
- **Legal Compliance**: Meets electronic signature legal requirements
- **Audit Trails**: Complete record of signing process and verification
- **Document Integrity**: Hash-based verification ensures document authenticity

### How It Works

1. **Document Preparation**: Server prepares document for signing
2. **Signing Request**: Server sends signing request to user
3. **Identity Verification**: User verifies their identity through Self
4. **Document Review**: User reviews document content
5. **Cryptographic Signing**: Document signed with user's private key
6. **Verification**: Server verifies signature and document integrity

### Use Cases

- **Legal Contracts**: Business agreement signing
- **Terms of Service**: Accepting terms and conditions
- **Financial Agreements**: Loan documents and financial contracts
- **Healthcare Consent**: Medical consent forms and agreements
- **Employment Documents**: Employment contracts and agreements
- **Real Estate**: Property purchase and rental agreements

[**Learn More →**](../solutions/digital-signatures/)

## Working Code Examples

All use cases are demonstrated through **open source implementations** that showcase real-world usage of Self SDK. These examples are designed as complete, functional applications that you can run and test immediately.

### Self SDK Examples Repository

Our examples repository provides fully functional implementations across multiple platforms:

**Mobile Peers (Frontend):**
- **Android:** Complete Android SDK implementations with UI flows
- **iOS:** Swift examples with proper async/await patterns

**Server Peers (Backend):**
- **Golang:** Backend integration examples with full workflows
- **Java:** Enterprise-ready server implementations

### Quick Start - Try Examples Immediately

For instant testing without building from source, use our pre-compiled binaries and applications:

| Platform | Quick Install | Source Code |
|----------|---------------|-------------|
| **Mobile - Android** | [Download from Play Store](https://play.google.com/store/apps/details?id=com.joinself.app.demo) | [View Source](https://github.com/joinself/demo-android) |
| **Mobile - iOS** | [Download from App Store](PLACEHOLDER_IOS_APP_STORE_LINK) | [View Source](https://github.com/joinself/demo-ios) |
| **Server - Golang** | `docker run -it ghcr.io/joinself/self-sdk-demo:go` | [View Source](https://github.com/joinself/academy/blob/main/examples/server/) |
| **Server - Java** | `docker run -it ghcr.io/joinself/self-sdk-demo:java` | [View Source](https://github.com/joinself/academy/blob/main/examples/server/) |

### Build from Source

To customize and build the examples yourself:

```bash
# Clone with all submodules
git clone --recurse-submodules https://github.com/joinself/academy.git

# Or if already cloned, initialize submodules
git submodule update --init --recursive
```

### Repository Structure

Each platform's examples are organized in dedicated folders within the repository:

| Platform | Example Location | Type | Description |
|----------|-----------------|------|-------------|
| **Android** | [`demo-android`](https://github.com/joinself/demo-android) | Mobile Peer | Complete Android apps demonstrating all SDK features |
| **iOS** | [`demo-ios`](https://github.com/joinself/demo-ios) | Mobile Peer | Native iOS implementations with SwiftUI |
| **Golang** | [`examples/server/`](https://github.com/joinself/academy/blob/main/examples/server/) | Server Peer | Backend Go implementations with full workflows |
| **Java** | [`examples/server/`](https://github.com/joinself/academy/blob/main/examples/server/) | Server Peer | Enterprise Java server examples |

This structure provides both **server peers** (backend) and **mobile peers** (frontend) that can interact with each other, giving you a complete Self network for testing and development.

## Next Steps

**Tip:** Start by exploring the examples repository to see Self SDK in action, then follow the detailed guides above to implement the features in your own application.

Ready to master digital identity? Join thousands of developers learning to build the future of authentication:

- **[Start Your Learning Journey](../examples/setup/)**: Begin with identity management
- **[Join Developer Community](../examples/resources/)**: Connect with other developers
- **[Explore Advanced Features](../examples/advanced/)**: Learn production-ready patterns 
