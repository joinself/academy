# Tenancy Agreement Signing System Demo

A complete, production-ready example of digital document signing using the Self SDK.  
**Authenticate with your Self mobile app, fill out a tenancy agreement, and sign it digitally with biometric security—no passwords, no complexity, just secure digital signatures.**

---

## 🎯 What You'll Learn
- How to implement a complete digital signing workflow with Self SDK
- How to create and manage PDF agreements for signing
- How to extend authentication workflows with signing capabilities
- How to handle credential issuance for digital signatures
- How to build a production-ready signing system

---

## 🚀 Quick Start

1. **Install Go 1.24+**
2. **Clone this repo and enter the directory:**
   ```sh
   git clone <repo-url>
   cd <repo-dir>
   ```
3. **Generate a secure storage key:**
   ```sh
   openssl rand -base64 32
   ```
4. **Run the server:**
   ```sh
   SELF_AUTH_STORAGE_KEY="your-base64-key" go run cmd/server/main.go
   ```
5. **Open [http://localhost:8081](http://localhost:8081) in your browser.**
6. **Authenticate and sign a tenancy agreement with your Self mobile app!**

---

## 🔍 How It Works

### Authentication Workflow
1. **QR Code Generation**: Server creates a unique QR code for authentication
2. **Mobile Connection**: User scans QR code with Self mobile app
3. **Biometric Verification**: App requests liveness check for authentication
4. **Session Creation**: Server creates authenticated session for user

### Signing Workflow
1. **Agreement Creation**: User fills out tenancy agreement form
2. **PDF Generation**: Server creates PDF agreement with user details
3. **Credential Issuance**: Server creates and issues a signing credential with agreement metadata
4. **Verification Request**: Server sends credential verification request with PDF evidence
5. **User Review**: User reviews agreement details on mobile app
6. **Digital Signature**: User approves with biometric authentication
7. **Response Processing**: Server decodes credential verification response and extracts real claims
8. **Signature Verification**: Server validates the signed credential as proof of signature

---

## 📝 What Just Happened

### Authentication Phase
- You initiated authentication by clicking "Start Authentication"
- Server generated a unique QR code containing cryptographic connection details
- You scanned the QR code with your Self mobile app
- App performed liveness verification to prove your identity
- Server verified your credentials and created a secure session

### Signing Phase
- You filled out the tenancy agreement form with property details
- Server created a PDF agreement and uploaded it to secure storage
- Server issued a signing credential containing agreement metadata (type, ID, hash, etc.)
- Server sent a credential verification request with the PDF as evidence
- You reviewed the agreement terms and approved with biometric authentication
- Server received and decoded the credential verification response
- Server extracted real claims from the signed credential (agreement type, signing date, document hash)
- The credential was validated and stored as a permanent, verifiable digital signature

---

## 🏗️ Architecture

### Core Components

**Authentication Service (`internal/auth/service.go`)**
- QR code generation for mobile connections
- Credential verification and session management
- PDF agreement creation and storage
- Digital signature request handling

**Signing Service (`internal/auth/sign.go`)**
- Credential issuance for digital signatures
- Credential verification request creation with evidence
- Real credential verification response processing
- Claims extraction from signed credentials

**HTTP Server (`internal/server/server.go`)**
- REST API endpoints for authentication and signing
- Session management and request routing
- Error handling and response formatting

**Web Interface (`web/index.html`)**
- Modern, responsive UI for user interaction
- Real-time status updates and progress tracking
- Form handling and API integration

### Key Features

**🔐 Biometric Security**
- Liveness verification for authentication
- Biometric approval for digital signatures
- No passwords or traditional credentials

**📄 Digital Signatures**
- Legally binding digital signatures with real credential verification
- PDF agreement generation and storage with evidence attachment
- Cryptographic proof of signature with extracted claims validation
- Robust response processing for different credential structures

**🔄 Real-time Updates**
- Live status checking for authentication
- Progress tracking for signing workflow
- Automatic retry and error handling

**🎨 Modern UI**
- Responsive design for all devices
- Intuitive user experience
- Professional styling and animations

---

## 🔧 API Endpoints

### Authentication
- `POST /api/v1/auth/request` - Start authentication
- `GET /api/v1/auth/status/{requestId}` - Check authentication status

### Signing
- `POST /api/v1/sign/request` - Create signing request
- `GET /api/v1/sign/status/{requestId}` - Check signing status

### Request Examples

**Start Authentication:**
```bash
curl -X POST http://localhost:8081/api/v1/auth/request \
  -H "Content-Type: application/json" \
  -d '{}'
```

**Create Signing Request:**
```bash
curl -X POST http://localhost:8081/api/v1/sign/request \
  -H "Content-Type: application/json" \
  -d '{
    "tenant_name": "John Doe",
    "property_address": "123 Main Street, City, State 12345",
    "rent_amount": "1500",
    "start_date": "2024-02-01",
    "end_date": "2025-01-31"
  }'
```

---

## 🎓 Key Learning Points

### Digital Signature Fundamentals

**What is a Digital Signature?**
A cryptographically secure credential that represents a legally binding agreement, issued by a trusted authority and verified through real credential processing.

**Core Components:**
- **Agreement**: The document being signed (PDF, contract, etc.)
- **Evidence**: The PDF document attached as evidence to the credential
- **Signer**: The person providing the signature
- **Claims**: Real metadata extracted from the signed credential (timestamp, document hash, agreement type, etc.)
- **Credential**: The cryptographically signed proof of agreement with verifiable claims

### Self SDK Architecture

**Credential-Based Signatures:**
- Digital signatures are implemented as verifiable credentials with real verification
- Each signature is a credential with claims representing the agreement
- Cryptographic signatures ensure integrity and authenticity
- Evidence attachments link signatures to original documents
- Real credential verification responses are decoded and processed
- Claims are extracted from actual signed credentials, not hardcoded values

**Workflow Integration:**
- Authentication and signing workflows are seamlessly integrated
- Same mobile app handles both authentication and signing
- Consistent user experience across all interactions
- Unified session management and security model
- Real-time credential verification response processing
- Robust error handling for different response structures

---

## 🔧 Customization Options

### Different Agreement Types

**Standard Types:**
- `TenancyAgreement`: For rental contracts
- `EmploymentContract`: For job agreements
- `ServiceAgreement`: For service contracts
- `CustomAgreement`: For domain-specific use cases

### Adding More Fields

**Extended Tenancy Agreement:**
```json
{
  "tenant_name": "John Doe",
  "property_address": "123 Main Street",
  "rent_amount": "1500",
  "start_date": "2024-02-01",
  "end_date": "2025-01-31",
  "security_deposit": "1500",
  "utilities_included": true,
  "pet_policy": "No pets allowed",
  "parking_spaces": 1
}
```

### Custom Credential Claims

**Extended Signing Credential:**
```go
claims := map[string]interface{}{
    "agreementType":    "Tenancy Agreement",
    "agreementId":      fmt.Sprintf("%x", agreementPDF.Id()),
    "documentHash":     fmt.Sprintf("%x", agreementPDF.Hash()),
    "signingDate":      time.Now().Format("2006-01-02"),
    "signingTimestamp": time.Now().Unix(),
    "documentType":     "application/pdf",
    "evidenceId":       fmt.Sprintf("%x", agreementPDF.Id()),
    "tenantName":       tenantName,
    "propertyAddress":  propertyAddress,
    "rentAmount":       rentAmount,
    "securityDeposit":  securityDeposit,
}
```

### Custom Storage Paths

**Storage Considerations:**
- Use secure, backed-up locations for production
- Separate storage per environment (dev/staging/prod)
- Ensure proper access permissions for PDF storage

### Credential Verification Response Handling

**Real Implementation Features:**
- Decodes actual credential verification responses using `DecodeCredentialVerificationResponse()`
- Extracts real claims from signed credentials, not hardcoded values
- Handles different response structures with robust fallback mechanisms
- Provides comprehensive logging for debugging and monitoring
- Gracefully handles cases where credential data isn't available

---

## 🎯 Real-World Applications

### Property Management
- **Use Case**: Digital tenancy agreement signing
- **Benefits**: Faster processing, reduced paperwork, secure storage
- **Integration**: Property management software integration

### Legal Services
- **Use Case**: Contract signing and verification
- **Benefits**: Legally binding signatures, audit trails, compliance
- **Integration**: Legal document management systems

### Financial Services
- **Use Case**: Loan agreements and financial contracts
- **Benefits**: Regulatory compliance, secure verification, cost reduction
- **Integration**: Banking and financial platforms

---

## 🚀 Production Deployment

### Environment Variables
```bash
SELF_AUTH_STORAGE_KEY="your-base64-encrypted-storage-key"
SELF_AUTH_STORAGE_PATH="./secure_storage"
SELF_SERVER_ADDRESS="0.0.0.0"
SELF_SERVER_PORT="8081"
```

### Security Considerations
- Use strong, randomly generated storage keys
- Secure storage locations with proper access controls
- HTTPS in production for all API communications
- Regular security audits and updates

### Scaling Considerations
- Database integration for session management
- Load balancing for high-traffic scenarios
- Monitoring and logging for production environments
- Backup and disaster recovery procedures

---

## 📚 Next Steps

**Ready to build your own signing system?**
- **Custom Agreements**: Adapt the PDF generation for your document types
- **Multi-Party Signing**: Extend for multiple signers on single documents
- **Integration**: Connect with existing document management systems
- **Compliance**: Add specific compliance features for your industry
- **Real Credential Processing**: Implement actual credential verification response handling
- **Claims Extraction**: Build robust systems for extracting real claims from signed credentials

**Learn More:**
- **[Verifiable Credentials](../../docs/concepts/verifiable-credentials.md)** - Understand credential fundamentals
- **[Digital Signatures](../../docs/solutions/digital-signatures.md)** - Learn about digital signature patterns
- **[Authentication System](../auth-system/)** - See the base authentication system
- **[Age Verification](../age-verifier/)** - Explore another credential verification example
- **[Credential Presentation Patterns](../../docs/concepts/credential-presentation-patterns.md)** - Learn about credential request patterns
- **[Credential Exchange Examples](../examples/credentials.md)** - See real credential exchange implementations

---

## 🤝 Contributing

This example demonstrates production-ready patterns for digital signing with Self SDK, including real credential verification response processing. Feel free to:

- **Extend the functionality** with additional agreement types
- **Improve the UI** with better styling and user experience
- **Add more features** like multi-party signing or document templates
- **Implement real credential processing** in your own applications
- **Share your use cases** and how you've adapted this example

---

**Built with Self SDK • Production-ready digital signing system with real credential verification** 
