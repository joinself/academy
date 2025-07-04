# Multi-Claim Credential Issuance

🎯 **What you'll learn:**
- How to add multiple claims to a single credential
- Different data types in claims (strings, booleans, numbers)
- Organizing related identity information
- Efficient credential structuring

📚 **Prerequisites:** Complete `../01_basic/` example first.

## Overview

This example demonstrates creating credentials with multiple claims using the Self SDK. Instead of single-value credentials, you'll learn to group related information efficiently in one cryptographically signed credential.

## Key Concepts

### Multi-Claim Credentials
A credential can contain multiple related pieces of information:
- **Profile credentials**: Name, age, country, verification status
- **Education credentials**: Institution, degree, GPA, graduation details
- **Employment credentials**: Company, role, start date, salary

### Data Types in Claims
Claims support various data types:
```go
claims := map[string]interface{}{
    "firstName":        "John",              // string
    "age":              30,                  // integer
    "isActive":         true,                // boolean
    "registrationDate": "2006-01-02",       // string (formatted date)
    "gpa":              3.8,                 // float64
}
```

## Code Breakdown

### 1. Account Setup
```go
func createAccounts() (*account.Account, *account.Account) {
    issuerCfg := &account.Config{
        StorageKey:  generateStorageKey("multi_issuer"),
        StoragePath: "./multi_issuer_storage",
        Environment: account.TargetSandbox,
        LogLevel:    account.LogWarn,
    }
    
    issuer, err := account.New(issuerCfg)
    // ... error handling
}
```

**What happens:**
1. Creates separate storage for issuer and holder
2. Generates unique cryptographic keys for each
3. Configures sandbox environment for safe testing

### 2. Profile Credential Creation
```go
claims := map[string]interface{}{
    "firstName":        "John",
    "lastName":         "Doe",
    "displayName":      "John Doe",
    "profileLevel":     "verified",
    "country":          "United States",
    "age":              30,
    "isActive":         true,
    "registrationDate": time.Now().Format("2006-01-02"),
}

credentialBuilder := credential.NewCredential().
    CredentialType(credential.CredentialTypeProfileName).
    CredentialSubject(credential.AddressKey(holderAddress)).
    CredentialSubjectClaims(claims).
    Issuer(credential.AddressKey(issuerAddress)).
    ValidFrom(time.Now()).
    SignWith(issuerAddress, time.Now())
```

**Key points:**
- `CredentialType`: Uses standard profile name type
- `CredentialSubjectClaims`: Accepts the entire claims map
- All claims are cryptographically signed together
- Mixed data types are supported seamlessly

### 3. Education Credential Creation
```go
claims := map[string]interface{}{
    "institution":      "University of Technology",
    "degree":           "Bachelor of Science",
    "major":            "Computer Science",
    "graduationYear":   2020,
    "gpa":              3.8,
    "honors":           true,
    "creditsCompleted": 120,
    "thesis":           "Machine Learning Applications",
    "graduationDate":   "2020-05-15",
}

credentialBuilder := credential.NewCredential().
    CredentialType([]string{"VerifiableCredential", "EducationCredential"}).
    // ... rest of the configuration
```

**Key points:**
- Custom credential types using string arrays
- Academic-specific claims structure
- Numeric values for quantitative data (GPA, credits)
- Boolean flags for achievements (honors)

## 🚀 Quick Start

### Go

```bash
cd examples/go/02_credentials/01_issuing_credentials/02_multi_claim
go run main.go
```

### Java
```bash
cd java
gradle run
```

## Expected Output

```
Multi-Claim Credential Issuance Demo
====================================
Issuer: 1:ABC123...
Holder: 1:DEF456...

Creating profile credential...
✅ Profile credential issued (Type: [VerifiableCredential ProfileName])
Creating education credential...
✅ Education credential issued (Type: [VerifiableCredential EducationCredential])
✅ Demo completed successfully!
```

## What Just Happened

1. **Created two accounts** with separate encrypted storage
2. **Built profile credential** with 8 different claims including strings, numbers, and booleans
3. **Built education credential** with 9 academic-related claims
4. **Issued both credentials** with cryptographic signatures
5. **Each credential maintains integrity** - all claims are signed together

## Benefits of Multi-Claim Credentials

### Efficiency
- **Fewer transactions**: One credential instead of multiple single-claim credentials
- **Reduced overhead**: Single signature covers all related claims
- **Atomic operations**: All claims are verified together

### Organization
- **Logical grouping**: Related information stays together
- **Context preservation**: Claims make sense as a unit
- **Easier management**: One credential to track instead of many

### Security
- **Cryptographic integrity**: All claims protected by single signature
- **Tamper evidence**: Any change to any claim invalidates the credential
- **Selective disclosure**: Can reveal subsets of claims when needed

## Best Practices

### Claim Grouping
```go
// ✅ Good: Related claims together
profileClaims := map[string]interface{}{
    "firstName": "John",
    "lastName":  "Doe",
    "country":   "US",
}

// ❌ Avoid: Unrelated claims together
mixedClaims := map[string]interface{}{
    "firstName":     "John",
    "bankAccount":   "123456789",  // Different context
    "medicalRecord": "private",    // Different privacy level
}
```

### Data Types
```go
// ✅ Use appropriate types
claims := map[string]interface{}{
    "age":        30,                    // int for numbers
    "isActive":   true,                  // bool for flags
    "joinDate":   "2023-01-15",         // string for dates
    "score":      95.7,                  // float64 for decimals
}
```

## Troubleshooting

### Common Issues

**"Failed to build credential"**
- Check that all required fields are provided
- Ensure claim values are valid JSON types
- Verify addresses are properly opened

**"Failed to issue credential"**
- Account permissions may be insufficient
- Network connectivity issues
- Storage path conflicts

### Debugging Tips
- Enable debug logging: `LogLevel: account.LogDebug`
- Check storage directory permissions
- Verify environment configuration (Sandbox vs Production)

## Next Steps

📚 **Continue learning:**
- `../03_with_evidence/` - Evidence and asset management
- `../04_complex_data/` - Complex nested data structures  
- `../05_comprehensive/` - All features combined

🔍 **Related examples:**
- `../01_basic/` - Single-claim credentials
- `../../02_exchanging_credentials/` - Using credentials in presentations

## When to Use Multi-Claim Credentials

### Perfect for:
- **Identity profiles**: Name, age, location, verification status
- **Academic records**: Institution, degree, GPA, graduation date
- **Employment verification**: Company, role, start date, salary
- **Medical records**: Patient info, conditions, treatments, dates

### Consider alternatives for:
- **Unrelated data**: Use separate credentials for different contexts
- **Different privacy levels**: Separate sensitive from public claims
- **Different lifespans**: Claims that expire at different times 
