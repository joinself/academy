# Age Verification System Demo

A plug-and-play, production-grade example of age verification using the Self SDK.  
**Click "I'm 18 or older", scan a QR code, verify your age with your Self mobile app, and access age-restricted content—no passwords, no sessions, no complexity.**

---

## 🎯 What You'll Learn
- How to implement an age verification workflow with Self SDK
- How QR code-based age verification works
- How to extract and verify a user's date of birth
- How to implement zero-knowledge age verification

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
6. **Click "I'm 18 or older" and scan the QR code with your Self mobile app to verify your age!**

---

## 🔍 How It Works

- Click the "I'm 18 or older" button to start the age verification process.
- The server generates a unique QR code for the age verification request.
- You scan the QR code with your Self mobile app.
- The app requests your date of birth credential (or prompts you to create one via document verification).
- The server verifies your age and grants or denies access based on whether you're 18 or older.
- No personal data is stored permanently, and the verification uses conditional disclosure to minimize data exposure.
- If verified, you gain access to the "super secret content"!

---

## 📝 What Just Happened
- You initiated an age verification request by declaring you're 18 or older.
- You scanned a QR code, which contained a unique cryptographic age verification request.
- Your Self app provided a date of birth credential, proving your age.
- The server calculated your age and verified you're 18+ without storing your personal data.
- You gained access to age-restricted content—no passwords, no session state, complete privacy.

---

## 🔞 Age Verification Features

- **Privacy-preserving**: Only age verification result is shared, not actual date of birth
- **Conditional disclosure**: Implementation where birth date is only revealed if verification passes
- **Cryptographic security**: Document verification and tamper-evident credentials
- **No data storage**: Age verification without storing personal information
- **Instant verification**: Real-time age checks via mobile credentials
- **Regulatory compliance**: Meets age verification requirements for various jurisdictions

---

## 🏗️ Technical Architecture

### Age Verification Flow
```
User Declaration → QR Generation → Mobile Scan → Credential Request → 
Age Calculation → Access Decision → Content Display
```

### Key Components
- **Date of Birth Credentials**: Verified through government document scanning
- **Age Calculation**: Server-side calculation from provided birth date
- **Access Control**: Binary decision based on 18+ verification
- **Conditional Disclosure**: Age verification that only reveals birth date when verification passes

---

## 🎛️ Configuration Options

### Environment Variables
- `SELF_AUTH_STORAGE_KEY`: Required encryption key for credential storage
- `SELF_AUTH_STORAGE_PATH`: Optional custom path for storage (default: `./auth_service_storage`)
- `PORT`: Optional server port (default: 8081)

### Age Verification Settings
- **Minimum Age**: Currently set to 18 (configurable in code)
- **Session Duration**: 24-hour access after successful verification
- **Credential Types**: Date of birth credentials from trusted issuers

---

## 🔐 Security Considerations

### Trust and Privacy
- **Document Verification**: Age credentials are issued after verifying government documents
- **Cryptographic Proofs**: All credentials are cryptographically signed and tamper-evident
- **No Data Retention**: Personal information is not stored on verification servers
- **Minimal Data Exposure**: Only age verification status is communicated

### Production Deployment
- Use secure, backed-up storage for account data
- Implement proper key management and rotation
- Configure appropriate session timeouts
- Monitor for unusual verification patterns
- Ensure compliance with local age verification regulations

---

## 🌐 Use Cases

### Age-Restricted Content
- **Media Platforms**: Age verification for mature content
- **E-commerce**: Age verification for restricted product purchases
- **Gaming**: Age verification for mature-rated games
- **Financial Services**: Age verification for account opening

### Regulatory Compliance
- **COPPA Compliance**: Verifying users are 13+ for data collection
- **GDPR Compliance**: Age verification for consent mechanisms
- **Tobacco/Alcohol**: Age verification for restricted product sales
- **Gambling**: Age verification for betting and gaming platforms

---

## 🚀 Advanced Features

### Conditional Disclosure Age Verification
The system uses conditional disclosure where the birth date is only revealed if the user meets the age requirement:

```go
// Server requests age verification with conditional disclosure
ageRequest := message.NewCredentialPresentationRequest().
    Details("DateOfBirthCredential", []*message.CredentialPresentationDetailParameter{
        message.NewCredentialPresentationDetailParameter(
            message.OperatorLessThan,
            "dateOfBirth",
            eighteenYearsAgo,
        ),
    })
```

This only returns the birth date credential if the user is 18+, minimizing unnecessary data exposure.

---

## 📊 Monitoring and Analytics

### Verification Metrics
- Age verification success/failure rates
- Time to complete verification
- Geographic distribution of verifications
- Device and platform usage statistics

### Compliance Reporting
- Verification attempt logs (without personal data)
- Failed verification patterns
- System uptime and availability
- Security incident tracking 
