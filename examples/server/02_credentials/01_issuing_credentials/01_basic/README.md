# Basic Credential Issuance Demo 🟢

> **📖 Learn the concepts first:** [Verifiable Credentials Concepts](../../../../../../docs/concepts/verifiable-credentials.md)

Learn the foundation concepts of verifiable credential issuance using the Self SDK.

## 🎯 What You'll Learn

- **Account Setup**: How to create issuer and holder accounts
- **Credential Builder Pattern**: Using the SDK's fluent API for credential creation
- **Claims and Data**: Adding information to credentials
- **Cryptographic Signing**: Securing credentials with digital signatures
- **Core SDK Usage**: Working directly with the underlying Self SDK

## 🚀 Quick Start

### Go
```bash
cd go
go run main.go
```

### Java
```bash
cd java
# Coming soon - Java implementation
```

### Expected Output

```
🎓 Basic Credential Issuance Demo

🔧 Setting up Self account...
✅ Self account created successfully
🔧 Setting up Self account...
✅ Self account created successfully
📬 Issuer Address: 007d3bf33f9e3f6e2e5afba703ee8436e1c4f1d5d4ce204a1478b5a626302a6de4
📬 Holder Address: 003def6f477c09858ccb1d189f83cffcb28b6ec99e1f86662eaa71b8d084343672

📧 Creating email credential...
✅ Email credential created
   📧 Email: alice@example.com
   🔒 Type: [VerifiableCredential EmailCredential]
   🆔 Subject: did:key:z6MkesiRSMeDJVwUXFBzCotZJpmLPur1nvLLP8wdCwnFUDS9

✅ Demo completed!
```

## 🏗️ How It Works

### Step 1: Account Creation

The demo creates two separate Self accounts with isolated storage:

**Conceptual Flow:**
```
Create Issuer Account:
  - Initialize with separate storage directory
  - Generate cryptographic keys for signing
  - Register with Self network

Create Holder Account:
  - Initialize with separate storage directory  
  - Generate keys for receiving credentials
  - Register with Self network
```

**Key Concepts:**
- **Issuer Account**: Creates and signs credentials for others
- **Holder Account**: Receives and stores credentials from issuers
- **Isolated Storage**: Each account has its own encrypted storage directory
- **DIDs**: Each account gets a unique Decentralized Identifier

### Step 2: Credential Building

The credential is built using the SDK's fluent builder pattern:

**Conceptual Structure:**
```
Credential Builder:
  - Set credential type (e.g., EmailCredential)
  - Define who the credential is about (subject)
  - Add the actual data (claims)
  - Specify who issued it (issuer)
  - Set validity period
  - Apply cryptographic signature
```

**Key Concepts:**
- **Builder Pattern**: Fluent API for step-by-step credential construction
- **Credential Type**: Standard types like `EmailCredential` for interoperability
- **Subject**: The entity (person/organization) the credential describes
- **Claims**: The actual data contained in the credential
- **Issuer**: The entity that created and vouches for the credential

### Step 3: Claims Structure

Claims contain the actual information being verified:

**Example Claims Data:**
```json
{
  "emailAddress": "alice@example.com",
  "verified": true,
  "domain": "example.com"
}
```

**Key Concepts:**
- **Flexible Data**: Claims can contain any JSON-serializable data
- **Standardization**: Common claim names improve interoperability
- **Verification Status**: Claims can include metadata about verification

### Step 4: Cryptographic Security

The credential is secured through digital signatures:

**Security Process:**
```
1. Hash the credential content
2. Sign the hash with issuer's private key
3. Include signature and timestamp in credential
4. Anyone can verify using issuer's public key
```

**Key Concepts:**
- **Digital Signatures**: Prove the credential hasn't been tampered with
- **Public Key Cryptography**: Anyone can verify, only issuer can sign
- **Timestamps**: When the signature was created
- **Non-Repudiation**: Issuer cannot deny creating the credential

## 🎓 Key Learning Points

### Verifiable Credentials Fundamentals

**What is a Verifiable Credential?**
A digitally signed document that contains claims about a subject, issued by a trusted authority.

**Core Components:**
- **Issuer**: Who created the credential (university, government, employer)
- **Subject**: Who the credential is about (student, citizen, employee)
- **Claims**: What information is being verified (degree, age, skills)
- **Signature**: Cryptographic proof of authenticity

### Self SDK Architecture

**Account-Based Model:**
- Each participant has a Self account with cryptographic keys
- Accounts are identified by DIDs (Decentralized Identifiers)
- Storage is encrypted and isolated per account

**Builder Pattern Benefits:**
- **Readable**: Code clearly shows what's being built
- **Flexible**: Easy to add or modify credential properties
- **Safe**: Compile-time checking of required fields

## 🔧 Customization Options

### Different Credential Types

**Standard Types:**
- `EmailCredential`: For email verification
- `ProfileNameCredential`: For identity verification
- `CustomCredential`: For domain-specific use cases

### Adding More Claims

**Extended Email Credential:**
```json
{
  "emailAddress": "alice@example.com",
  "verified": true,
  "verificationDate": "2024-01-15",
  "verificationMethod": "email_link",
  "domain": "example.com",
  "isPrimary": true
}
```

### Custom Storage Paths

**Storage Considerations:**
- Use secure, backed-up locations for production
- Separate storage per environment (dev/staging/prod)
- Ensure proper access permissions

## 🎯 Real-World Applications

### Email Verification Service
- **Use Case**: Verify user email addresses for account registration
- **Claims**: Email address, verification status, verification method
- **Issuer**: Email verification service
- **Subject**: User creating account

### Identity Verification
- **Use Case**: Digital identity for online services
- **Claims**: Name, email, verification level, verification date
- **Issuer**: Identity verification provider
- **Subject**: Individual user

### Domain Verification
- **Use Case**: Prove ownership of email domains
- **Claims**: Domain name, ownership status, verification method
- **Issuer**: Domain verification service
- **Subject**: Domain owner

## 🔗 Next Steps

After mastering basic credential issuance:

1. **📝 Multi-Claim Credentials**: Try `../02_multi_claim` for complex credentials
2. **🔍 Evidence-Based**: Explore `../03_with_evidence` for proof-backed credentials
3. **📊 Complex Data**: Check `../04_complex_data` for advanced structures
4. **🔄 Credential Exchange**: Move to `../../02_exchanging_credentials` for full workflows

## 🔧 Troubleshooting

### Account Creation Failed
```
❌ Failed to create Self account
```
**Solution**: Check network connectivity and storage permissions

### Credential Build Failed
```
❌ Error building credential
```
**Solution**: Verify all required fields are provided and issuer account is valid

### Storage Permission Issues
```
❌ Cannot create storage directory
```
**Solution**: Ensure write permissions in working directory

---

**Ready to issue your first verifiable credentials?** This foundation will power all your credential workflows! 🎓
