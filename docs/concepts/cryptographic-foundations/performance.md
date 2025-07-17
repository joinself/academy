# Performance Characteristics

### Ed25519 Performance

**Speed Advantages:**

- **Signing**: ~40,000 signatures/second on modern CPUs
- **Verification**: ~15,000 verifications/second  
- **Key Generation**: ~10,000 keypairs/second
- **Memory**: Minimal RAM requirements (~1KB per operation)

**Comparison with RSA:**

| Operation | Ed25519 | RSA-2048 | Speedup |
|-----------|---------|----------|---------|
| Sign      | 68μs    | 1100μs   | 16x     |
| Verify    | 201μs   | 33μs     | 0.16x*  |
| Key Gen   | 51μs    | 226ms    | 4400x   |

*Note: Ed25519 verification is slower than RSA but still very fast for most applications

### Network Efficiency

**Bandwidth Optimization:**

- **Public Keys**: 32 bytes vs 256+ bytes for RSA
- **Signatures**: 64 bytes vs 256+ bytes for RSA  
- **Total Overhead**: ~50% reduction in cryptographic data size 
