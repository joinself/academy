# Identity Verification

> **🎯 What you'll learn:** How to issue and verify credentials to confirm a user's identity with cryptographic certainty, streamlining onboarding and enhancing trust.

Identity verification is the process of confirming that a person is who they claim to be. With Self, this is achieved by issuing and verifying credentials that contain verified information about a user, such as their name, age, or email address. This process is more secure and privacy-preserving than traditional methods.

## How It Works

A typical identity verification flow looks like this:

1.  **Identity Proofing**: An authoritative source (like your application or a trusted third party) verifies the user's identity through some offline or online process (e.g., checking a government ID, verifying an email address).
2.  **Credential Issuance**: The authority issues a verifiable credential to the user containing the verified information (e.g., an `EmailCredential` or a `KYCCredential`).
3.  **Presentation Request**: When the user needs to prove their identity to another service (a "verifier"), that service requests a presentation of the relevant credential.
4.  **Verification**: The verifier receives the credential and cryptographically verifies its authenticity and integrity, confirming the user's identity.

## Core Concepts

The foundation of identity verification is the issuance and exchange of verifiable credentials. These concepts are essential to understanding the process:

- **[Verifiable Credentials](../concepts/verifiable-credentials.md)**: The core data structure for holding and exchanging identity information.
- **[Cryptographic Foundations](../concepts/cryptographic-foundations.md)**: Learn about the public-key cryptography that secures the entire process.

## Server-Side Examples

These examples walk you through the process of issuing and verifying credentials for identity verification purposes.

### 1. Issuing a Credential

The first step is to issue a credential to a user after verifying some piece of information about them.

- **[Basic Credential Issuance](https://github.com/joinself/academy/tree/main/examples/server/02_credentials/01_issuing_credentials/01_basic/)**: Learn the fundamentals of creating and issuing a simple credential.
- **[Issuing with Evidence](https://github.com/joinself/academy/tree/main/examples/server/02_credentials/01_issuing_credentials/03_with_evidence/)**: For higher-stakes verification, you can attach evidence (like a scanned document) to a credential.

<div data-github-embed="https://github.com/joinself/academy/blob/main/examples/server/02_credentials/02_exchanging_credentials/email_verification/go/main.go#L179-L218"
     data-style="github-dark-dimmed"
     data-show-border="true"
     data-show-line-numbers="true"
     data-show-file-meta="true"
     data-show-full-path="true"
     data-show-copy="true"></div>

### 2. Exchanging and Verifying Credentials

Once a user has a credential, they can present it to other services to prove their identity.

- **[Email Verification Exchange](https://github.com/joinself/academy/tree/main/examples/server/02_credentials/02_exchanging_credentials/email_verification/)**: This comprehensive example demonstrates the full lifecycle: issuing an email credential and then using it for verification.
- **[Presentation Request](https://github.com/joinself/academy/tree/main/examples/server/02_credentials/02_exchanging_credentials/presentation_request/)**: This example focuses on the verifier's side, showing how to request and handle a presentation of credentials.

<div data-github-embed="https://github.com/joinself/academy/blob/main/examples/server/02_credentials/02_exchanging_credentials/email_verification/go/main.go#L221-L258"
     data-style="github-dark-dimmed"
     data-show-border="true"
     data-show-line-numbers="true"
     data-show-file-meta="true"
     data-show-full-path="true"
     data-show-copy="true"></div>

## Mobile Implementation

The user's mobile device is where they manage their digital identity. The Self mobile SDKs provide the necessary UI and functionality for a user to receive and present credentials.

- **[Receiving and Storing Credentials](https://github.com/joinself/academy/tree/main/examples/mobile/android/02_credentials/)**: This example shows how a user's mobile application can be notified of a new credential, and how it is securely stored in their digital wallet.
- **[Presenting Credentials for Verification](https://github.com/joinself/academy/tree/main/examples/mobile/android/02_credentials/)**: This example demonstrates the user flow for when they are asked to present a credential. The UI components handle the user's approval and the secure sharing of the credential with the verifier.

## Next Steps

- **[Digital Signatures](./digital-signatures.md)**: Apply credential issuance to create legally binding digital signatures.
- **[Advanced Credential Examples](https://github.com/joinself/academy/tree/main/examples/credentials.md)**: Explore more complex credential structures and workflows. 
