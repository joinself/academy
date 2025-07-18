# Credential Issuance Examples

> **Learn the concepts first:** [Verifiable Credentials Concepts](../../../../../docs/concepts/verifiable-credentials.md)

Learn how to create, configure, and issue verifiable credentials using the Self SDK.

## Purpose

This section focuses exclusively on **credential issuance** - the process of creating, signing, and distributing verifiable credentials. Master these patterns before moving to exchange and storage.

## Learning Progression

| Example | Complexity | Time | Focus |
|---------|------------|------|-------|
| **`01_basic/`** | Beginner | 5-10min | Foundation concepts |
| **`02_multi_claim/`** | Intermediate | 10-15min | Multiple claims |
| **`03_with_evidence/`** | Advanced | 15-20min | File attachments |
| **`04_complex_data/`** | Advanced | 20-25min | Nested structures |
| **`05_comprehensive/`** | Expert | 30-45min | All features |

## Quick Start

```bash
# Start with foundation concepts
cd 01_basic && go run main.go

# Progress through complexity
cd ../02_multi_claim && go run main.go
cd ../03_with_evidence && go run main.go
cd ../04_complex_data && go run main.go
cd ../05_comprehensive && go run main.go
```

## What You'll Learn

### Core Concepts
- Account setup for issuers and holders
- Credential builder patterns
- Cryptographic signing workflows
- Validation and error handling

### Advanced Patterns
- Multiple claims in single credentials
- File evidence and asset management
- Complex nested data structures
- Production-ready configurations

## Use Cases Covered

- **Email Verification**: Basic identity claims
- **Profile Management**: Personal information credentials
- **Document Certification**: Credentials with file evidence
- **Organizational Data**: Complex institutional credentials
- **Educational Records**: Academic achievement credentials

## Next Steps

After mastering issuance:
- **Exchange**: `../02_exchanging_credentials/` - Request and verify credentials
- **Storage**: `../03_credential_storage/` - Manage credential collections 
