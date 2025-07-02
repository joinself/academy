# Common Utilities for Self SDK Examples

This directory provides shared utilities for Self SDK examples to reduce code duplication and provide consistent patterns across all examples.

## Overview

The common utilities centralize the account setup patterns that were previously repeated across multiple examples. Instead of each example implementing its own account setup functions with slightly different configurations, they can now use the standardized patterns provided here.

## Key Features

### Account Setup Patterns

- **Basic Account Setup** - Simple account setup with minimal configuration
- **Custom Configuration** - Account setup with custom callbacks and storage locations
- **Issuer/Holder Pattern** - Creates both issuer and holder accounts (common in credential examples)
- **Display Utilities** - Consistent information display across examples

### Utility Functions

- **Account Information Display** - Shows connection information for single accounts
- **Account Pair Display** - Shows information for issuer/holder account pairs
- **Storage Management** - Consistent storage directory patterns

## Language-Specific Implementations

### Go Implementation (`account.go`)
The Go implementation provides standardized functions for account setup and utilities:

```go
// Available in Go implementation
SetupAccount(config AccountConfig)           // Main setup with flexible config
SetupIssuerHolder(issuerCallbacks, holderCallbacks) // Credential workflow pattern  
SetupBasicAccount()                          // Simple default setup
DisplayAccountInfo(acc, name)                // Single account info
DisplayAccountPair(issuer, holder)           // Issuer/holder info
```

### Java Implementation
```
Coming soon - Java implementations will provide equivalent functionality
```

## Usage Patterns

### Basic Account Setup

**Conceptual Flow:**
```
1. Initialize account with default configuration
2. Set up storage in standard location
3. Connect to Self network
4. Return ready-to-use account
```

**Go Example:**
```go
import "github.com/joinself/academy/examples/server/common"

// Simple account with default settings
account := common.SetupBasicAccount()
defer account.Close()
```

### Custom Configuration

**Conceptual Flow:**
```
1. Define custom callbacks for events
2. Specify custom storage location (optional)
3. Initialize account with configuration
4. Connect to network
```

**Go Example:**
```go
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

**Conceptual Flow:**
```
1. Create issuer account with signing capabilities
2. Create holder account with receiving capabilities
3. Set up separate storage for each role
4. Display both account addresses for testing
```

**Go Example:**
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

## Migration Benefits

### Before Common Utilities
Each example contained duplicated account setup code:
- Repeated configuration patterns
- Inconsistent storage locations
- Duplicated error handling
- Maintenance burden across multiple files

### After Common Utilities
Examples use standardized patterns:
- **Reduced Duplication**: Account setup logic centralized
- **Consistency**: All examples use same patterns and storage naming
- **Maintainability**: Changes made in one place
- **Flexibility**: Support for different configurations while maintaining consistency

## Configuration Options

### Standard Configuration Fields

- **Storage Directory**: Custom base directory for account storage
- **Event Callbacks**: Functions to handle account events (connections, messages, etc.)
- **Environment**: Sandbox vs production network selection
- **Logging Level**: Control amount of debug information

### Default Settings

**Environment**: Sandbox network (suitable for development and testing)
**Storage**: Consistent, predictable paths across all examples
**Security**: Cryptographically secure storage keys
**Logging**: Minimal logging to reduce noise during examples

## Storage Layout Patterns

### Single Account Examples
```
./storage/                    # Standard single account storage
├── account.db               # Account identity and keys  
├── connections.db           # Connection data
└── messages.db              # Message history
```

### Issuer/Holder Examples  
```
./issuer_storage/            # Issuer account storage
├── account.db
├── connections.db
└── messages.db

./holder_storage/            # Holder account storage  
├── account.db
├── connections.db
└── messages.db
```

This ensures that:
- All single-account examples share consistent storage patterns
- Issuer and holder accounts in credential examples are kept separate
- Storage paths are predictable across all examples
- Examples can be run multiple times without conflicts

## Integration in Examples

### Adding to Your Example

**Go Implementation:**
```go
// In your go.mod file
require (
    github.com/joinself/academy/examples/server/common v0.0.0-00010101000000-000000000000
)

replace github.com/joinself/academy/examples/server/common => ../../../common
```

**Java Implementation:**
```
Coming soon - Java dependency management patterns
```

### Using in Your Code

Replace custom account setup functions with calls to common utilities:

1. **Identify** your account setup pattern
2. **Choose** the appropriate common utility function  
3. **Replace** custom setup code with common utility call
4. **Test** that your example works with the new pattern

## Best Practices

### For Example Authors

- **Use Common Patterns**: Always prefer common utilities over custom setup
- **Follow Storage Conventions**: Use standard storage directory patterns
- **Document Deviations**: If you need custom behavior, document why
- **Test Consistency**: Ensure your example works with standard patterns

### For Common Utility Maintenance

- **Maintain Backward Compatibility**: Don't break existing examples
- **Document Changes**: Update README when adding new utilities
- **Test Across Examples**: Verify changes work across multiple examples
- **Keep It Simple**: Focus on common patterns, not complex edge cases

---

**Need to add a new utility pattern?** Follow the existing patterns and update this documentation to help other example authors! 🛠️
