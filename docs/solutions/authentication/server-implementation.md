# Server Implementation

> ****What you will learn:** What you'll learn:** Complete backend implementation of Self authentication with detailed flow breakdown, correlation system, and code examples.

This guide covers the technical implementation of Self authentication on the server side, including the complete authentication flow, request-response correlation, and session management.

## Quick Start

Experience the full authentication workflow with our **[production-ready](../examples/solutions/auth-system/)** example featuring:

- QR code generation with automatic expiration
- Real-time request tracking and correlation  
- Biometric liveness verification
- Session management with configurable timeouts
- Clean integration interface for production use

**Try it now:**
```bash
cd examples/solutions/auth-system
SELF_AUTH_STORAGE_KEY="$(openssl rand -base64 32)" go run cmd/solutions/main.go
# Open http://localhost:8081 and scan the QR code with the Self Demo App!
```

> ****Troubleshooting:** Testing:** While mobile SDK examples are in development, you can test the authentication flow using the **[Self Demo App](https://play.google.com/store/apps/details?id=com.joinself.app.demo)** - a ready-to-use mobile app for testing Self authentication backends.


## Implementation Phases

### 1. QR Code Generation & Content ID Mapping
   
   Your server creates a unique discover request embedded on a QR code with a trackable content ID:

<div data-github-embed="https://github.com/joinself/academy/blob/main/examples/solutions/auth-system/internal/auth/service.go#L279-L313"
     data-style="github-dark-dimmed"
     data-show-border="true"
     data-show-line-numbers="true"
     data-show-file-meta="true"
     data-show-full-path="true"
     data-show-copy="true"></div>

### 2. User Connection & Discovery Response
   
   When users scan the QR code with their mobile app, the SDK automatically correlates their response:

<div data-github-embed="https://github.com/joinself/academy/blob/main/examples/solutions/auth-system/internal/auth/service.go#L366-L410"
     data-style="github-dark-dimmed"
     data-show-border="true"
     data-show-line-numbers="true"
     data-show-file-meta="true"
     data-show-full-path="true"
     data-show-copy="true"></div>

### 3. Biometric Verification Request
   
   Once connected, the server requests liveness verification for authentication:

<div data-github-embed="https://github.com/joinself/academy/blob/main/examples/solutions/auth-system/internal/auth/service.go#L415-L461"
     data-style="github-dark-dimmed"
     data-show-border="true"
     data-show-line-numbers="true"
     data-show-file-meta="true"
     data-show-full-path="true"
     data-show-copy="true"></div>

### 4. Credential Verification & Session Creation
   
   The server verifies the biometric proof and creates an authenticated session:

<div data-github-embed="https://github.com/joinself/academy/blob/main/examples/solutions/auth-system/internal/auth/service.go#L462-L520"
     data-style="github-dark-dimmed"
     data-show-border="true"
     data-show-line-numbers="true"
     data-show-file-meta="true"
     data-show-full-path="true"
     data-show-copy="true"></div>

> ****Documentation:** Full Implementation:** See the complete backend implementation with HTTP server, UI, and production-ready architecture at **[examples/solutions/auth-system/](https://github.com/joinself/academy/blob/main/examples/solutions/auth-system/)**

### Request-Response Correlation System

A critical aspect of Self authentication is the sophisticated correlation system that tracks requests through their complete lifecycle using cryptographic content IDs.

#### Content ID-Based Tracking

The Self SDK provides automatic correlation between QR codes, user responses, and authentication completion using unique content identifiers embedded in each message.

#### Three-Phase Correlation Process

**Phase 1: QR Code → Discovery Response**

- QR code contains discovery request with unique content ID
- Your mobile app scans QR and responds with discovery response containing correlation ID
- Server matches response to original request using content ID

**Phase 2: Discovery → Credential Request**  

- Server sends liveness verification request to connected user
- New content ID generated for credential request
- Auth request updated with credential request ID for tracking

**Phase 3: Credential Request → Authentication Result**

- User provides biometric proof via credential presentation response
- Server correlates credential response to original credential request
- Authentication completed and session created upon successful verification

This multi-phase correlation ensures perfect request isolation and enables concurrent authentication flows without interference.

## Next Steps

- **[Mobile Implementation](./mobile-implementation.md)**: Integrate Self authentication into your mobile app
- **[Conclusions & Best Practices](./conclusions.md)**: Production deployment guide, session management strategies, and security best practices
- **[Authentication Overview](../authentication.md)**: Return to the main authentication guide 
