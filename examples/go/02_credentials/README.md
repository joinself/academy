# Verifiable Credentials Examples

This directory contains comprehensive examples for working with verifiable credentials using the Self SDK, organized by functionality for better learning and implementation.

> **📢 Updated Organization**: Examples are now organized by functionality rather than complexity, making it easier to find relevant patterns for your use case.

## 🗂️ Functional Organization

| 🎯 Category | 📁 Directory | 🎓 Focus Area |
|-------------|--------------|---------------|
| **Issue Credentials** | `01_issuing_credentials/` | Creating and signing credentials |
| **Exchange Credentials** | `02_exchanging_credentials/` | Presentation requests and verification |
| **Store Credentials** | `03_credential_storage/` | Management and organization patterns |

## 🚀 Quick Start by Use Case

### 🏭 **Credential Issuers** (Organizations creating credentials)
```bash
cd 01_issuing_credentials/
# Start with basic, then progress through complexity
cd 01_basic && go run main.go
```

### 🔍 **Credential Verifiers** (Applications requesting proof)
```bash
cd 02_exchanging_credentials/
# Learn presentation request patterns
cd presentation_request && go run main.go
```

### 👤 **Credential Holders** (Individuals managing credentials)
```bash
cd 03_credential_storage/
# Understand storage and organization
cd basic_storage && go run main.go
```

## 📚 Learning Paths

### 🎯 Path 1: Complete Issuer Journey
1. `01_issuing_credentials/01_basic/` - Foundation
2. `01_issuing_credentials/02_multi_claim/` - Multiple claims
3. `01_issuing_credentials/03_with_evidence/` - File attachments
4. `01_issuing_credentials/04_complex_data/` - Nested structures
5. `01_issuing_credentials/05_comprehensive/` - All features

### 🎯 Path 2: Exchange Implementation
1. `01_issuing_credentials/01_basic/` - Understand credentials first
2. `02_exchanging_credentials/presentation_request/` - Basic exchange
3. `03_credential_storage/basic_storage/` - Storage integration

### 🎯 Path 3: Full Stack Understanding
Complete all directories in order: `01_issuing_credentials/` → `02_exchanging_credentials/` → `03_credential_storage/`

## 🏗️ Architecture Overview

```
┌─────────────────┐    ┌──────────────────┐    ┌─────────────────┐
│   01_issuing    │───▶│  02_exchanging   │───▶│ 03_credential   │
│   _credentials  │    │   _credentials   │    │    _storage     │
└─────────────────┘    └──────────────────┘    └─────────────────┘
       │                        │                        │
       ▼                        ▼                        ▼
   Create & Sign          Request & Verify         Store & Manage
   Credentials            Presentations            Credentials
```

## 🔧 Prerequisites

1. Go 1.19 or later
2. Self SDK dependencies (handled by go.mod in each example)

## 📖 Legacy Structure Reference

If you're familiar with the previous complexity-based organization:

| Old Structure | New Location |
|---------------|--------------|
| `basic/` | `01_issuing_credentials/01_basic/` |
| `multi_claim/` | `01_issuing_credentials/02_multi_claim/` |
| `evidence/` | `01_issuing_credentials/03_with_evidence/` |
| `complex/` | `01_issuing_credentials/04_complex_data/` |
| `advanced/` | `01_issuing_credentials/05_comprehensive/` |
| `exchange_demo/` | `02_exchanging_credentials/presentation_request/` |

## 🎓 Educational Philosophy

**Functionality over Complexity**: Learn by use case rather than feature complexity, making it easier to apply knowledge to real-world scenarios.

## 🔗 Related Examples

- **Setup**: `../00_setup/` - Account creation and configuration
- **Connections**: `../01_connection/` - Establishing peer connections  
- **Chat**: `../04_chat/` - Messaging with credential context
