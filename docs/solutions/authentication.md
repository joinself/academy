# Authentication

> **🎯 What you'll learn:** How to replace traditional logins with secure, passwordless authentication using verifiable credentials for both server-side and mobile applications.

Self's authentication solution enables you to build secure, passwordless login experiences. Instead of relying on vulnerable usernames and passwords, users authenticate using their unique SelfID, backed by biometrics. This approach provides higher security and a more seamless user experience.

## How It Works

The typical authentication flow involves these steps:

1.  **Connection**: The user scans a QR code or accepts a notification to establish a secure connection with your application.
2.  **Presentation Request**: Your application requests proof of identity from the user. This is done by requesting a presentation of one or more verifiable credentials.
3.  **Presentation Verification**: The user approves the request on their mobile device, and your application cryptographically verifies the received credentials.
4.  **Session Establishment**: Upon successful verification, the user is granted access and a session is created.

## Core Concepts

Authentication is achieved by verifying control over a decentralized identifier (DID) and exchanging verifiable credentials. To understand the foundations, please review these concepts:

- **[Decentralized Identity](../concepts/decentralized-identity.md)**: The core paradigm behind Self's authentication.
- **[Secure Connections](../concepts/secure-connections.md)**: How two parties establish a trusted communication channel.
- **[Verifiable Credentials](../concepts/verifiable-credentials.md)**: The data format used to prove identity attributes.

## Server-Side Examples

The following examples demonstrate how to implement the core authentication workflows from your backend.

### 1. Account Setup and Connections

Before you can authenticate a user, you need to set up your own application's identity and be able to establish a connection with them.

- **[Setup a new account](https://github.com/joinself/academy/tree/main/examples/server/00_setup/01_new_account/):** Learn how to create a new identity for your application.
- **[Connect with a user via QR code](https://github.com/joinself/academy/tree/main/examples/server/01_connection/02_qr/)**: The most common way to initiate an interaction with a user is by displaying a QR code they can scan.

<div data-github-embed="https://github.com/joinself/academy/blob/main/examples/server/01_connection/02_qr/go/main.go#L80-L112"
     data-style="github-dark-dimmed"
     data-show-border="true"
     data-show-line-numbers="true"
     data-show-file-meta="true"
     data-show-full-path="true"
     data-show-copy="true"></div>

and accept the connection response

<div data-github-embed="https://github.com/joinself/academy/blob/main/examples/server/01_connection/02_qr/go/main.go#L140-L142"
     data-style="github-dark-dimmed"
     data-show-border="true"
     data-show-line-numbers="true"
     data-show-file-meta="true"
     data-show-full-path="true"
     data-show-copy="true"></div>


### 2. Requesting and Verifying Credentials

Once a connection is established, you can request credentials to authenticate the user. The `email_verification` example is a great starting point, as it shows credential presentation request in a real-world scenario.

- **[Email Verification](https://github.com/joinself/academy/tree/main/examples/server/02_credentials/02_exchanging_credentials/email_verification/):** This example shows a complete flow where a service provider issues an email credential, and a user presents it for verification.

<div data-github-embed="https://github.com/joinself/academy/blob/main/examples/server/02_credentials/02_exchanging_credentials/email_verification/go/main.go#L179-L218"
     data-style="github-dark-dimmed"
     data-show-border="true"
     data-show-line-numbers="true"
     data-show-file-meta="true"
     data-show-full-path="true"
     data-show-copy="true"></div>

## Mobile Implementation

From the user's perspective, authentication happens on their mobile device. The Self SDK provides UI components to handle these interactions securely and seamlessly.

- **[New Account Creation](https://github.com/joinself/academy/tree/main/examples/mobile/android/00_setup/01_new_account/):** Before a user can authenticate, they need a Self identity. This example shows how a user can create one within your mobile app.
- **[Credential Presentation](https://github.com/joinself/academy/tree/main/examples/mobile/android/02_credentials/):** These examples showcase how a user is prompted to share their credentials for authentication. The pre-built UI handles the secure presentation of credentials to your application.

## Next Steps

- **[Identity Verification](./identity-verification.md)**: Learn how to issue your own credentials to verify user identities.
- **[Chat & Messaging](https://github.com/joinself/academy/tree/main/examples/chat.md)**: Build secure, in-app communication channels with authenticated users. 
