# Basic Credential Issuance Demo 🟢

Learn the foundation concepts of verifiable credential issuance using the Self SDK.

## 🎯 What You'll Learn

- **Account Setup**: How to create issuer and holder accounts
- **Credential Builder Pattern**: Using the SDK's fluent API for credential creation
- **Claims and Data**: Adding information to credentials
- **Cryptographic Signing**: Securing credentials with digital signatures
- **Core SDK Usage**: Working directly with the underlying Self SDK

## 🚀 Quick Start

```bash
go run main.go
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

```go
func createAccounts() (*account.Account, *account.Account) {
    issuer := common.SetupAccount(common.AccountConfig{
        StorageDir: "./basic_issuer",  // Separate storage for issuer
        Callbacks:  account.Callbacks{},
    })

    holder := common.SetupAccount(common.AccountConfig{
        StorageDir: "./basic_holder",  // Separate storage for holder
        Callbacks:  account.Callbacks{},
    })

    return issuer, holder
}
```

**Key Concepts:**
- **Issuer Account**: Creates and signs credentials for others
- **Holder Account**: Receives and stores credentials from issuers
- **Isolated Storage**: Each account has its own encrypted storage directory
- **DIDs**: Each account gets a unique Decentralized Identifier

### Step 2: Credential Building

The credential is built using the SDK's fluent builder pattern:

```go
credentialBuilder := credential.NewCredential().
    CredentialType(credential.CredentialTypeEmail).           // What type of credential
    CredentialSubject(credential.AddressKey(holderAddress)).  // Who it's about
    CredentialSubjectClaims(claims).                          // The actual data
    Issuer(credential.AddressKey(issuerAddress)).             // Who issued it
    ValidFrom(time.Now()).                                    // When it becomes valid
    SignWith(issuerAddress, time.Now())                       // Cryptographic signature
```

**Key Concepts:**
- **Builder Pattern**: Fluent API for step-by-step credential construction
- **Credential Type**: Standard types like `EmailCredential` for interoperability
- **Subject**: The entity (person/organization) the credential describes
- **Claims**: The actual data contained in the credential
- **Issuer**: The entity that created and vouches for the credential

### Step 3: Claims Structure

Claims contain the actual information being verified:

```go
claims := map[string]interface{}{
    "emailAddress": "alice@example.com",  // The email being verified
    "verified":     true,                 // Verification status
    "domain":       "example.com",        // Email domain
}
```

**Key Concepts:**
- **Flexible Data**: Claims can contain any JSON-serializable data
- **Standardization**: Common claim names improve interoperability
- **Verification Status**: Claims can include metadata about verification

### Step 4: Cryptographic Security

The credential is secured through digital signatures:

```go
SignWith(issuerAddress, time.Now())  // Signs with issuer's private key
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

```go
// Identity credential
.CredentialType(credential.CredentialTypeProfileName)

// Custom credential type
.CredentialType([]string{"CustomCredential", "VerifiableCredential"})
```

### Adding More Claims

```go
claims := map[string]interface{}{
    "emailAddress":    "alice@example.com",
    "verified":        true,
    "verificationDate": time.Now().Format("2006-01-02"),
    "verificationMethod": "email_link",
    "domain":          "example.com",
    "isPrimary":       true,
}
```

### Custom Storage Paths

```go
issuer := common.SetupAccount(common.AccountConfig{
    StorageDir: "/path/to/secure/storage",  // Custom storage location
    Callbacks:  account.Callbacks{},
})
```

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

1. **Multiple Claims**: `../02_multi_claim/` - Learn to add multiple pieces of information
2. **File Evidence**: `../03_with_evidence/` - Attach supporting documents
3. **Complex Data**: `../04_complex_data/` - Handle nested and hierarchical information
4. **All Features**: `../05_comprehensive/` - Production-ready patterns

## 💡 Production Considerations

### Security
- **Key Management**: Secure storage of cryptographic keys
- **Access Control**: Who can issue which types of credentials
- **Audit Trails**: Track all credential issuance activities

### Scalability
- **Batch Processing**: Issue multiple credentials efficiently
- **Database Integration**: Store credential metadata
- **Monitoring**: Track issuance rates and errors

### Interoperability
- **Standard Types**: Use well-known credential types
- **Consistent Claims**: Follow naming conventions
- **Version Management**: Handle credential schema evolution

## 🐛 Troubleshooting

### Storage Issues
```
Failed to create Self account: Storage could not be opened
```
**Solution**: Check directory permissions and ensure storage paths are writable.

### Network Issues
```
Failed to connect to Self network
```
**Solution**: Verify internet connection and check if using correct environment (Sandbox vs Production).

### Build Issues
```
cannot find package
```
**Solution**: Run `go mod tidy` to ensure all dependencies are downloaded. 
