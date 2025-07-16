# Server Implementation

> **🎯 What you'll learn:** How to implement a KYC/age-check process by requesting a verifiable credential from a user's device.

This guide covers the technical implementation of an identity verification service from the backend's perspective. The process builds upon the principles of the **[authentication server implementation](../authentication/server-implementation.md)** but focuses on requesting and verifying specific claims from a user.

## Quick Start

Experience the complete age verification workflow with our **[production-ready age verifier](../../examples/solutions/age-verifier/)** example featuring:

- Age verification with "I'm 18 or older" button workflow
- Date of birth credential requests and validation
- Automatic age calculation from birth date
- Privacy-preserving verification with minimal data exposure
- Success/failure screens with access control

**Try it now:**
```bash
cd examples/solutions/age-verifier
SELF_AUTH_STORAGE_KEY="$(openssl rand -base64 32)" go run cmd/server/main.go
# Open http://localhost:8081 and verify your age!
```

## Core Workflow

The initial steps for setting up a connection with the user's mobile device are exactly the same as in the authentication flow. Please refer to the **[Authentication Server Implementation guide](../authentication/server-implementation.md#implementation-phases)** for a detailed explanation of:

1.  **QR Code Generation & User Connection**

The key difference for identity verification lies in what happens after the connection is established.

## On-Demand Credential Issuance

Your application does not need to manage the document verification and credential issuance process. This is handled by the Self SDK on the mobile app in coordination with Self's backend services.

1.  **Your server requests a credential** containing a specific field (e.g., `dateOfBirth`).
2.  If the user does not have this credential, the **Self SDK automatically initiates the issuance flow**.
3.  The user is prompted to **capture their document** on their mobile device.
4.  The document is securely sent to **Self's services for verification**.
5.  A **verifiable credential is issued** and stored on the user's device.

Once the credential has been issued, it can be presented to your service.

## Implementation Phases

### 1. Requesting Date of Birth Credentials

After establishing a secure channel, your server sends a request for a presentation of a credential containing the required claims. For age verification, you request the `dateOfBirth` field:

<div data-github-embed="https://github.com/joinself/academy/blob/main/examples/solutions/age-verifier/internal/auth/service.go#L448-L492"
     data-style="github-dark-dimmed"
     data-show-border="true"
     data-show-line-numbers="true"
     data-show-file-meta="true"
     data-show-full-path="true"
     data-show-copy="true"></div>

### 2. Processing Credential Response & Age Verification

Once the user consents, your server receives the verifiable credential. The verification process involves extracting the `dateOfBirth` claim and calculating the user's age:

<div data-github-embed="https://github.com/joinself/academy/blob/main/examples/solutions/age-verifier/internal/auth/service.go#L615-L649"
     data-style="github-dark-dimmed"
     data-show-border="true"
     data-show-line-numbers="true"
     data-show-file-meta="true"
     data-show-full-path="true"
     data-show-copy="true"></div>

### 3. Age Calculation and Verification Logic

The system includes robust age calculation that handles multiple date formats and determines if the user meets the minimum age requirement:

<div data-github-embed="https://github.com/joinself/academy/blob/main/examples/solutions/age-verifier/internal/auth/service.go#L524-L553"
     data-style="github-dark-dimmed"
     data-show-border="true"
     data-show-line-numbers="true"
     data-show-file-meta="true"
     data-show-full-path="true"
     data-show-copy="true"></div>


## Advanced: Zero-Knowledge Verification

For enhanced privacy, you can perform a verification without the user's actual data (like their date of birth) ever leaving their device. This is a form of zero-knowledge proof where the mobile app evaluates a condition locally and returns only the boolean result to your backend.

This is achieved by sending a presentation request with a logical operator. For example, to verify if a user is older than 18, you would send the following request:

```go
eighteenYearsAgo := time.Now().AddDate(-18, 0, 0).Format("2006-01-02")

content, err := message.NewCredentialPresentationRequest().
	Type([]string{"VerifiablePresentation"}).
	Details(
		[]string{"DateOfBirthCredential"},
		[]*message.CredentialPresentationDetailParameter{
			message.NewCredentialPresentationDetailParameter(
				message.OperatorLowerThan, 
				"dateOfBirth",
				eighteenYearsAgo,
			),
		},
	).
	Finish()
```

In this scenario, your server does not receive a credential with the `dateOfBirth` field. Instead, it receives a simple `true` or `false` response, confirming whether the user's date of birth meets the condition you specified. This allows you to perform an age check without ever processing the user's sensitive PII.

> **📖 Full Implementation:** See the complete age verification implementation with HTTP server, UI, and production-ready architecture at **[examples/solutions/age-verifier/](https://github.com/joinself/academy/blob/main/examples/solutions/age-verifier/)**

## Privacy & Compliance Features

The age verification system is designed with privacy and regulatory compliance in mind:

- **Minimal Data Processing**: Only the age verification result is retained, not the actual birth date
- **Cryptographic Security**: All credentials are cryptographically signed and tamper-evident  
- **No Permanent Storage**: Personal information is not stored beyond the verification session
- **Regulatory Compliance**: Meets requirements for COPPA, GDPR, and other age verification regulations

## Next Steps

- **[Mobile Implementation](./mobile-implementation.md)**: See how the mobile app handles document capture to fulfill the credential request.
- **[Conclusions & Best Practices](./conclusions.md)**: Review production deployment strategies for identity verification.
- **[Identity Verification Overview](../identity-verification.md)**: Return to the main identity verification guide. 

