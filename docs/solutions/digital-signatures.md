# Digital Signatures

> **🎯 What you'll learn:** How to create legally binding, biometrically-backed digital signatures by issuing a verifiable credential that represents an agreement.

With Self, you can implement legally binding, biometrically-backed digital signatures. This allows users to sign documents, contracts, or any other type of agreement directly from their phone, providing a higher level of assurance than traditional electronic signatures.

## How It Works

A digital signature in the Self ecosystem is essentially a verifiable credential that represents an agreement. The act of signing is the act of the user accepting and receiving this credential from you.

1.  **Agreement Creation**: Your application creates a document or agreement that needs to be signed.
2.  **Credential Formulation**: You create a verifiable credential where the claims represent the core terms of the agreement. For example, a claim could be `agreement_hash` with the hash of the document, `timestamp` of the signature, and `signer_identity`.
3.  **Issuance (The "Signature")**: You issue this credential to the user. The user is prompted on their device to review and accept the credential. By accepting, they are effectively signing the agreement. The credential is then stored in their mobile wallet as a permanent, verifiable record of the signature.
4.  **Verification**: Anyone who needs to verify the signature can request a presentation of this "signature credential" from the user and cryptographically verify its authenticity.

## Core Concepts

- **[Verifiable Credentials](../concepts/verifiable-credentials.md)**: The "digital signature" is a credential where the claims represent the terms of the agreement.
- **[Cryptographic Foundations](../concepts/cryptographic-foundations.md)**: The signature's integrity is guaranteed by the same cryptographic principles that secure all Self interactions.

## Server-Side Example

While there isn't a dedicated "digital signature" example yet, the process is a specific application of credential issuance. You can adapt the following example to create a signature workflow.

- **[Issuing with Evidence](https://github.com/joinself/academy/tree/main/examples/server/02_credentials/01_issuing_credentials/03_with_evidence/)**: This is the most relevant example. Imagine that instead of a "CertificationCredential", you are creating a "SignatureCredential".
    - The **claims** would describe the agreement (e.g., document name, hash, timestamp).
    - The **evidence** could be the actual PDF of the contract being signed.

<div data-github-embed="https://github.com/joinself/academy/blob/main/examples/server/02_credentials/01_issuing_credentials/03_with_evidence/go/main.go#L165-L194"
     data-style="github-dark-dimmed"
     data-show-border="true"
     data-show-line-numbers="true"
     data-show-file-meta="true"
     data-show-full-path="true"
     data-show-copy="true"></div>

By following this pattern, you can build a robust and secure digital signature solution.

## Mobile Implementation

The act of "signing" is performed by the user on their mobile device when they accept the "signature credential" that you issue. The mobile SDK handles this interaction.

- **[Receiving Credentials](https://github.com/joinself/academy/tree/main/examples/mobile/android/02_credentials/)**: The flow for receiving a signature credential on mobile is the same as for any other credential. The user is prompted with the details of the credential (which represent the agreement) and can approve or deny its acceptance with their biometrics. This approval is the legally binding signature. The credential is then stored in their wallet as a permanent record.

## Next Steps

- **[Complex Data Credentials](https://github.com/joinself/academy/tree/main/examples/server/02_credentials/01_issuing_credentials/04_complex_data/)**: Learn how to model more complex agreements with nested data structures in your credential claims.
