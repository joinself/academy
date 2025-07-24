# **Troubleshooting:** Troubleshooting Guide

This guide covers common issues and solutions across all Self SDK examples. Each section includes symptoms, common causes, and step-by-step solutions.

## Quick Navigation

- [**Architecture:** Setup & Account Issues](#setup--account-issues)
- [**Network:** Network & Connectivity Issues](#network--connectivity-issues)
- [**Connections:** Connection Issues](#connection-issues)
- [**Credentials:** Credential Issues](#credential-issues)
- [**Chat:** Chat & Messaging Issues](#chat--messaging-issues)
- [**Storage:** Storage & Permission Issues](#storage--permission-issues)
- [**Troubleshooting:** General Debugging Tips](#general-debugging-tips)

---

## **Architecture:** Setup & Account Issues

### Account Already Exists in Storage Directory

```bash
⚠️ Account already exists in storage directory
```

**Common Causes:**
- Trying to create a new account when one already exists in the storage path
- Previous demo run left storage files in place
- Running multiple examples with the same storage directory

**Solutions:**
1. **Use existing account**: Switch to the existing account example instead
2. **Fresh start**: Delete the storage directory: `rm -rf ./storage/` 
3. **Different path**: Use a different storage directory for new accounts

### Failed to Load Account

```bash
❌ Failed to load account: invalid storage format
```

**Common Causes:**
- Storage files corrupted due to incomplete shutdown
- SDK version mismatch between creation and loading
- Storage directory manually modified or corrupted

**Solutions:**
1. **Recreate account**: Delete storage and create new account with new account example
2. **Check SDK version**: Ensure consistent SDK version across runs
3. **Backup check**: Restore from backup if available (production environments)

### Account Initialization Failed

```bash
❌ Failed to create account: configuration error
```

**Common Causes:**
- Invalid storage key or configuration
- Insufficient permissions for storage directory
- Environment configuration mismatch

**Solutions:**
1. **Check configuration**: Verify all required config fields are set
2. **Check permissions**: Ensure write access to storage directory
3. **Environment**: Verify environment setting (sandbox vs production)

---

## **Network:** Network & Connectivity Issues

### Failed to Connect to Self Network

```bash
❌ Failed to connect to Self network
```

**Common Causes:**
- Internet connectivity issues
- Firewall blocking Self SDK network traffic
- Self network temporary outage
- Incorrect environment configuration

**Solutions:**
1. **Check internet**: Verify basic internet connectivity
2. **Check firewall**: Ensure firewall allows outbound HTTPS connections
3. **Environment**: Verify you're using the correct environment (sandbox/production)
4. **Retry**: Network issues are often temporary - retry after a moment

### Inbox Open Failed

```bash
❌ Failed to open inbox
```

**Common Causes:**
- Account not properly initialized
- Network connectivity issues
- Self network registration problems

**Solutions:**
1. **Account check**: Ensure account is properly created and initialized
2. **Network**: Verify connection to Self network
3. **Restart**: Try restarting the application
4. **Fresh account**: Create a new account if problem persists

---

## **Connections:** Connection Issues

### Connection Failed to Establish

```bash
❌ Failed to establish connection
```

**Common Causes:**
- SDK version incompatibility between parties
- Network connectivity issues
- Incorrect recipient address
- Account not properly initialized

**Solutions:**
1. **SDK versions**: Verify both parties use compatible SDK versions
2. **Network**: Check network connectivity on both sides
3. **Address**: Verify the recipient address is correct and current
4. **Account status**: Ensure both accounts are properly initialized

### Connection Request Failed

```bash
❌ Connection request failed
```

**Common Causes:**
- Recipient server not running or not accepting connections
- Incorrect inbox address format
- Network routing issues
- Recipient account issues

**Solutions:**
1. **Server status**: Verify the recipient server is running and accepting connections
2. **Address format**: Check inbox address format (should start with `did:self:`)
3. **Network**: Test basic network connectivity to recipient
4. **Retry**: Connection failures can be temporary

### QR Code Not Displaying

```bash
❌ Failed to generate QR code
```

**Common Causes:**
- Account setup incomplete
- Network connectivity issues during QR generation
- Terminal/environment doesn't support QR display
- Sandbox environment limitations

**Solutions:**
1. **Account setup**: Ensure account is fully initialized
2. **Network**: Check connection to Self network
3. **Terminal**: Try a different terminal or environment
4. **Environment**: QR display may be limited in some sandbox configurations

---

## **Credentials:** Credential Issues

### Credential Issuance Failed

```bash
❌ Failed to issue credential
```

**Common Causes:**
- Issuer or holder accounts not properly initialized
- Network connectivity issues
- Invalid credential claims or structure
- Cryptographic signing errors

**Solutions:**
1. **Account check**: Verify both issuer and holder accounts are initialized and connected
2. **Claims validation**: Check credential claims for invalid data or structure
3. **Network**: Ensure network connectivity for both parties
4. **Signing**: Verify issuer has proper signing capabilities

### Evidence Upload Failed

```bash
❌ Failed to upload evidence
```

**Common Causes:**
- File size exceeds limits
- Network connectivity issues during upload
- Invalid file format or object creation
- Storage space limitations

**Solutions:**
1. **File size**: Check file size limits and reduce if necessary
2. **Network**: Verify stable network connection during upload
3. **Object creation**: Ensure proper object creation with `object.New()`
4. **Storage**: Check available storage space

### Credential Verification Failed

```bash
❌ Failed to verify credential signature
```

**Common Causes:**
- Credential has been modified or tampered with
- Invalid issuer signature
- Cryptographic verification errors
- Credential format corruption

**Solutions:**
1. **Integrity check**: Ensure credential hasn't been modified since issuance
2. **Issuer validation**: Verify the issuer signature is valid and trusted
3. **Fresh copy**: Try with a fresh copy of the credential
4. **SDK version**: Ensure compatible SDK version for verification

### Presentation Creation Failed

```bash
❌ Failed to create verifiable presentation
```

**Common Causes:**
- Invalid credentials in the presentation
- Credentials not properly signed
- Presentation structure errors
- Missing required claims

**Solutions:**
1. **Credential validation**: Verify all credentials in presentation are valid and properly signed
2. **Structure check**: Ensure presentation follows correct format
3. **Claims**: Verify all required claims are present
4. **Credentials**: Try creating presentation with a single credential first

---

## **Chat:** Chat & Messaging Issues

### No Messages Received

```bash
**Success:** Chat demo ready! Press Ctrl+C to exit.
# ... no activity ...
```

**Common Causes:**
- No one has your inbox address to send messages
- Sender hasn't established a connection first
- Network connectivity issues
- Message routing problems

**Solutions:**
1. **Share address**: Share your inbox address from the application output
2. **Establish connection**: First establish a connection using connection examples
3. **Network**: Verify both sides can reach the Self network
4. **Test locally**: Try sending messages between local accounts first

### Message Send Failures

```bash
❌ [15:04:05] Failed to send response: connection not found
```

**Common Causes:**
- No active connection to the recipient
- Recipient address is invalid or expired
- Network interruption during send
- Connection was not properly established

**Solutions:**
1. **Connection check**: Verify active connection exists before sending
2. **Address validation**: Check recipient address format and validity
3. **Network**: Ensure stable network connectivity
4. **Re-establish**: Try re-establishing the connection

### Chat Decode Errors

```bash
❌ [15:04:05] Failed to decode chat message: invalid content type
```

**Common Causes:**
- Receiving non-chat message types
- Message corruption during transmission
- SDK version incompatibility
- Incorrect message handling

**Solutions:**
1. **Content type check**: Use `ContentTypeOf()` before calling `DecodeChat()`
2. **Message handling**: Add error handling for different message types
3. **SDK versions**: Verify both sides use compatible SDK versions
4. **Message validation**: Validate message structure before decoding

---

## **Storage:** Storage & Permission Issues

### Failed to Create Storage Directory

```bash
❌ Failed to create storage directory
```

**Common Causes:**
- Insufficient write permissions in working directory
- Directory path conflicts or invalid characters
- Disk space limitations
- Parent directory doesn't exist

**Solutions:**
1. **Permissions**: Ensure write permissions in the working directory
2. **Path**: Use a simpler path without special characters
3. **Disk space**: Check available disk space
4. **Create parent**: Manually create parent directories if needed

### Storage Corruption

```bash
❌ Storage corruption detected
```

**Common Causes:**
- Improper application shutdown during storage operations
- Disk space exhausted during write operations
- File system errors or hardware issues
- Concurrent access to storage files

**Solutions:**
1. **Fresh start**: Delete corrupted storage and create new account
2. **Proper shutdown**: Always ensure proper application shutdown
3. **Disk space**: Monitor disk space during operations
4. **Exclusive access**: Ensure only one process accesses storage

### Permission Denied Errors

```bash
❌ Permission denied: cannot access storage
```

**Common Causes:**
- Insufficient file system permissions
- Storage directory owned by different user
- Read-only file system
- Security software blocking access

**Solutions:**
1. **Check ownership**: Ensure current user owns storage directory
2. **Permissions**: Set appropriate permissions: `chmod 755 <storage_dir>`
3. **File system**: Verify file system is not read-only
4. **Security software**: Check if antivirus or security software is blocking access

---

## **Troubleshooting:** General Debugging Tips

### Enable Debug Logging

Add debug logging to get more detailed information:

```go
// For Go examples
cfg := &account.Config{
    LogLevel: account.LogDebug, // Enable detailed logging
    // ... other config
}
```

### Check SDK Version Compatibility

Ensure all parties use compatible SDK versions:

```bash
# Check your Go module version
go list -m github.com/joinself/self-go-sdk
```

### Environment Verification

Verify you're using the correct environment:

```go
// Development
Environment: account.TargetSandbox

// Production  
Environment: account.TargetProduction
```

### Network Diagnostics

Basic network connectivity tests:

```bash
# Test general internet connectivity
ping google.com

# Test HTTPS connectivity (Self SDK uses HTTPS)
curl -I https://httpbin.org/get
```

### Clean Restart Process

For persistent issues, try a clean restart:

1. Stop all running examples
2. Delete all storage directories: `rm -rf *_storage/`
3. Clear any temporary files
4. Restart with fresh accounts

### Resource Monitoring

Monitor system resources if experiencing performance issues:

```bash
# Check available disk space
df -h

# Check memory usage
free -h

# Check for running Self SDK processes
ps aux | grep self
```

---

## **Support:** Getting Additional Help

If you continue experiencing issues after trying these solutions:

### Community Resources
- **Self Developer Forum** - Get help from community and Self team
- **GitHub Issues** - Report bugs and request features
- **Discord Community** - Real-time chat with other developers

### When Reporting Issues
Include the following information:
1. **SDK version** and language (Go/Java/etc.)
2. **Operating system** and version
3. **Environment** (sandbox/production)
4. **Complete error message** with stack trace
5. **Steps to reproduce** the issue
6. **Expected vs actual behavior**

### Sample Issue Report Template
```
**SDK Version**: self-go-sdk v1.2.3
**OS**: macOS 12.6 / Ubuntu 20.04 / Windows 11
**Environment**: Sandbox
**Error**: Failed to establish connection
**Steps**: 
1. Created issuer account
2. Created holder account  
3. Attempted connection from holder to issuer
**Expected**: Connection established successfully
**Actual**: Connection failed with timeout error
**Additional context**: Works locally but fails on remote server
```

---

*This troubleshooting guide is continuously updated based on community feedback and common issues. Please contribute improvements via GitHub issues or community forums.* 
