# Mobile Implementation

> **🎯 What you'll learn:** How to integrate Self's mobile SDKs to handle document capture, data extraction, and credential presentation for a KYC/age-check workflow.

This guide covers the mobile application's role in the identity verification process, focusing on how the Self SDK simplifies capturing and verifying identity documents.

## Mobile Verification Flow

From the mobile app's perspective, the age verification process involves these steps:

1.  **QR Code Scanning**:
    - The app scans the QR code presented by the service.
    - It establishes a secure connection with the backend.

2.  **Receiving the Presentation Request**:
    - The app receives a request to present a credential with a `date_of_birth` field.

3.  **Document Capture and Data Extraction**:
    - The Self SDK activates the camera and guides the user to scan their identity document (e.g., driver's license).
    - The SDK's document verification engine processes the image on-device, extracting the necessary data. It performs checks for blurriness, glare, and other issues.

4.  **User Consent and Credential Presentation**:
    - The app displays the extracted data (`date_of_birth`) to the user for confirmation.
    - Upon user approval, the app generates a temporary verifiable credential containing the extracted information and presents it to the backend.

5.  **Verification Completion**:
    - The app receives a success or failure message from the backend and informs the user of the outcome.

## Mobile SDK Integration _(Coming Soon)_

> SDK examples demonstrating the complete document capture and credential presentation workflow are coming soon.

The Self mobile SDKs for iOS and Android will provide pre-built UI components to handle the entire document capture process, including:

- Camera management and user guidance.
- Image quality checks.
- On-device data extraction for various document types.
- Secure credential generation and presentation.

## Next Steps

- **[Server Implementation](./server-implementation.md)**: Understand how the backend requests and verifies the credential.
- **[Conclusions & Best Practices](./conclusions.md)**: Learn about production deployment and best practices for identity verification.
- **[Identity Verification Overview](../identity-verification.md)**: Return to the main identity verification guide. 
