# Key Management Patterns

### Storage Key Generation

In a production environment, the 32-byte storage key used to encrypt account data at rest should be managed with the utmost care. It is your responsibility to generate and store this key securely, for instance, by retrieving it from a secret management service or a secure environment variable.

For demonstration and testing purposes, the examples in this academy use a helper function `generateStorageKey` to create a key on the fly. This approach is **not suitable for production** but serves to illustrate how a key is used by the SDK.

**Example of a simplified key generation function:**

```go
// This function is a simplified example for demonstration purposes only.
// In production, use a robust and secure method to generate and manage your storage keys.
func generateStorageKey(seed string) []byte {
    key := make([]byte, 32)
    if _, err := rand.Read(key); err != nil {
        // A fallback to deterministic generation can be useful for testing,
        // but secure random generation is preferred.
        h := sha256.Sum256([]byte(fmt.Sprintf("self-sdk-%s-%d", seed, time.Now().UnixNano())))
        return h[:]
    }
    return key
}
```

**Key Properties:**

- **32-byte keys**: Standard for AES-256 encryption
- **Cryptographic quality**: Using `crypto/rand` for true randomness
- **Deterministic fallback**: SHA-256 based generation if random fails
- **Unique per account**: Isolated storage and key spaces

### Key Storage

**Encrypted Account Storage:**

- Private keys never leave encrypted storage
- Storage isolation per account prevents key mixing
- Cross-platform compatible storage format

**Storage Paths Pattern:**
```go
// From https://github.com/joinself/academy/tree/main/examples/server/04_advanced_features/01_core_features/go/main.go
cfg := &account.Config{
    // In production, the storage key must be securely generated and managed.
    // The generateStorageKey function is for demonstration purposes only.
    StorageKey:  generateStorageKey("advanced_demo"), // Key for this account
    StoragePath: "./advanced_demo_storage",           // Unique storage path
    Environment: account.TargetSandbox,
}
```

### Key Distribution

- Keys are not directly distributed between accounts.
- Public keys are shared for signature verification and encryption setup.
- Private keys remain on the device. 
