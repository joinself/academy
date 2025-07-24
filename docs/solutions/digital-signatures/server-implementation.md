# Digital Signatures: Server Implementation

> ****What you will learn:** What you'll learn:** How to implement a complete digital signature backend with real credential verification response processing, PDF agreement creation, and robust claims extraction.

> ****Documentation:** Complete Example**: See the full implementation in the [Tenancy Agreement Signing System Demo](https://github.com/joinself/academy/blob/main/examples/solutions/signing-tenacy-agreement/README.md) for a complete, production-ready digital signature example.


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

<div data-github-embed="https://github.com/joinself/academy/blob/main/examples/solutions/signing-tenacy-agreement/internal/auth/service.go#L673-L685"
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

<div data-github-embed="https://github.com/joinself/academy/blob/main/examples/solutions/signing-tenacy-agreement/internal/auth/sign.go#L24-59"
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

<div data-github-embed="https://github.com/joinself/academy/blob/main/examples/solutions/signing-tenacy-agreement/internal/auth/sign.go#L142-190"
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
