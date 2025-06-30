# Common Library for Self SDK Examples

This package provides shared utilities for Self SDK examples to reduce code duplication and provide consistent patterns across all examples.

## Overview

The common library centralizes the account setup patterns that were previously repeated across multiple examples. Instead of each example implementing its own `setupAccount()` function with slightly different configurations, they can now use the standardized functions provided here.

## Key Features

### Account Setup Functions

- **`SetupAccount(config AccountConfig)`** - Main account setup function with flexible configuration
- **`SetupIssuerHolder(issuerCallbacks, holderCallbacks)`** - Creates both issuer and holder accounts (common in credential examples)  
- **`SetupBasicAccount()`** - Simple account setup with minimal configuration

### Utility Functions

- **`DisplayAccountInfo(acc, name)`** - Shows connection information for a single account
- **`DisplayAccountPair(issuer, holder)`** - Shows information for issuer/holder account pairs

## Usage Examples

### Basic Account Setup

```go
import "github.com/joinself/academy/examples/go/common"

// Simple account with default settings
account := common.SetupBasicAccount()
defer account.Close()

// Account with custom callbacks
account := common.SetupAccount(common.AccountConfig{
    Callbacks: account.Callbacks{
        OnMessage: handleMessage,
        OnConnect: func(acc *account.Account) {
            fmt.Println("Connected!")
        },
    },
})
```

### Credential Examples (Issuer/Holder Pattern)

```go
// Create both issuer and holder accounts
issuer, holder := common.SetupIssuerHolder(
    account.Callbacks{}, // issuer callbacks
    account.Callbacks{}, // holder callbacks  
)
defer issuer.Close()
defer holder.Close()

// Display account information
common.DisplayAccountPair(issuer, holder)
```

### Custom Storage Directory

```go
account := common.SetupAccount(common.AccountConfig{
    StorageDir: "./custom_storage",
    Callbacks: account.Callbacks{
        OnConnect: onConnect,
    },
})
```

## Migration from Existing Examples

### Before (duplicated code)

```go
func setupAccount() *account.Account {
    cfg := &account.Config{
        StorageKey:  make([]byte, 32),
        StoragePath: "./storage", 
        Environment: account.TargetSandbox,
        LogLevel:    account.LogWarn,
        Callbacks: account.Callbacks{
            OnConnect: func(acc *account.Account) {
                fmt.Println("Connected")
            },
        },
    }
    
    selfAccount, err := account.New(cfg)
    if err != nil {
        log.Fatal("Failed to create account:", err)
    }
    
    return selfAccount
}
```

### After (using common library)

```go
import "github.com/joinself/academy/examples/go/common"

account := common.SetupAccount(common.AccountConfig{
    Callbacks: account.Callbacks{
        OnConnect: func(acc *account.Account) {
            fmt.Println("Connected")
        },
    },
})
```

## Adding to Your go.mod

Add the common library to your example's `go.mod`:

```go
module your_example

go 1.22

require (
    github.com/joinself/academy/examples/go/common v0.0.0-00010101000000-000000000000
    github.com/joinself/self-go-sdk v0.59.0
)

replace github.com/joinself/academy/examples/go/common => ../common
```

## Benefits

1. **Reduced Duplication** - Eliminates repeated `setupAccount` functions across examples
2. **Consistency** - All examples use the same account setup patterns and storage naming
3. **Maintainability** - Changes to account setup logic only need to be made in one place
4. **Flexibility** - Supports different callback configurations while maintaining consistency
5. **Documentation** - Centralized documentation for account setup patterns

## Configuration Details

### AccountConfig Fields

- **StorageDir** (optional) - Custom base directory for storage, defaults to current directory  
- **Callbacks** (optional) - Event callbacks for the account

### Default Settings

- **Environment**: `account.TargetSandbox` (suitable for development and testing)
- **LogLevel**: `account.LogWarn` (minimal logging)
- **Storage**: All accounts use `./storage` for consistency. For issuer/holder pairs, they use `./issuer/storage` and `./holder/storage` respectively.

All accounts created use:
- Cryptographically secure 32-byte storage keys
- Sandbox environment (appropriate for examples and development)
- Warning-level logging to reduce noise
- Consistent storage paths across all examples

## Storage Layout

### Single Account Examples
- Storage path: `./storage`

### Issuer/Holder Examples
- Issuer storage: `./issuer/storage`  
- Holder storage: `./holder/storage`

This ensures that:
- All single-account examples share the same storage location
- Issuer and holder accounts in credential examples are kept separate
- Storage paths are predictable and consistent across examples
