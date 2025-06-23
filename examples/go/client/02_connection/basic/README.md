# Basic Connection Example

This example demonstrates how to set up a Self SDK account that can receive connections from third-party applications (like mobile apps) via QR codes. It generates an actual scannable QR code using the underlying Self SDK.

## Complexity Rating: 3/10 ⭐⭐⭐

**Beginner-friendly** - Clean, simple code with functional QR generation.

## What You'll Learn

- **Account Setup**: How to create a Self account using the core SDK
- **QR Code Generation**: How to generate actual scannable QR codes
- **Connection Preparation**: How to prepare an account to receive connections
- **Real-World Flow**: Understanding how mobile apps connect via QR scanning

## Running the Example

```bash
cd examples/go/client/02_connection/basic
go run main.go
```

You'll see output like this:

```
🔗 Basic Connection Example
===========================
🔧 Setting up Self account...
✅ Account created successfully
🔗 Connected to Self network

📬 Account Address: 00321b095e3dda41452ec7ff57c257fc6fde87a186e2d48f44c2585a5137914781

📱 Generating QR code...

📱 QR CODE:
==========
██ ▄▄▄▄▄ █▀█ █▄▀▄▀ ▄▄▄▄▄ ██
██ █   █ █▀▀ █ ▀ ▀ █   █ ██
██ █▄▄▄█ █▀█ █▄▀▄▀ █▄▄▄█ ██
██▄▄▄▄▄▄▄█▄▀ ▀▄█▄▄▄▄▄▄▄██
██ ▄▄▄▄▄ █▀█ █▄▀▄▀ ▄▄▄▄▄ ██
==========
Valid for: 30 minutes

✅ Account ready! Scan the QR code above with a Self mobile app.
Press Ctrl+C to exit.
```

## How It Works

### 1. Account Creation
Creates a Self SDK account using minimal configuration:
```go
cfg := &account.Config{
    StorageKey:  make([]byte, 32),
    StoragePath: "./basic_connection_storage",
    Environment: account.TargetSandbox,
    LogLevel:    account.LogWarn,
}
```

### 2. QR Code Generation  
Uses the core SDK to generate a real connection QR code:
```go
// Generate key package for connection
keyPackage, err := selfAccount.ConnectionNegotiateOutOfBand(
    inboxAddress,
    time.Now().Add(30*time.Minute),
)

// Build discovery request
content, err := message.NewDiscoveryRequest().
    KeyPackage(keyPackage).
    Expires(time.Now().Add(30*time.Minute)).
    Finish()

// Create and encode QR code
anonymousMsg := event.NewAnonymousMessage(content)
qrCode, err := anonymousMsg.EncodeToQR(event.QREncodingUnicode)
```

### 3. Real-World Connection Flow

**Desktop/Server App (this example):**
1. Creates a Self account
2. Opens inbox to get connection address  
3. Generates QR code containing connection info
4. Waits for mobile app to scan and connect

**Mobile App (third-party):**
1. User scans QR code with Self mobile app
2. App extracts connection information
3. Initiates secure connection to desktop app
4. Encrypted communication channel established

## Code Structure

| Function | Purpose | Lines |
|----------|---------|-------|
| `main()` | Entry point and program flow | 17-32 |
| `setupAccount()` | Creates configured Self account | 34-56 |
| `showConnectionInfo()` | Displays connection details | 58-65 |
| `generateConnectionQR()` | Generates and displays QR code | 67-104 |

**Total: ~100 lines** - Clean and focused code.

## Key Features

### ✅ **Functional QR Generation**
- Generates real scannable QR codes
- Valid for 30 minutes
- Contains encrypted connection information
- Works with actual Self mobile apps

### 🔐 **Security Features**
- Cryptographic addresses ensure authenticity
- End-to-end encryption for all communication
- No central server stores messages  
- Mutual authentication between peers

### 📱 **Mobile Integration Ready**
- QR codes work with Self mobile apps
- Automatic connection handling
- Real-time bidirectional communication
- Secure messaging and credential exchange

## Use Cases

### Desktop Applications
Perfect foundation for desktop apps that mobile apps connect to:
- Web applications needing Self SDK integration
- Developer tools with mobile companion apps
- Demo applications for Self SDK features

### Server-Side Applications  
Ideal starting point for server applications:
- APIs that handle credential verification
- Services that issue credentials to mobile users
- Backend systems requiring secure mobile connections

### Learning and Development
Educational foundation for understanding:
- Core Self SDK functionality
- QR-based peer discovery
- Connection establishment principles
- Building blocks for advanced features

## Next Steps

### 1. Add Message Handling
Build on this foundation by adding chat capabilities:
```go
// Add to account callbacks
OnMessage: func(acc *account.Account, msg *event.Message) {
    // Handle incoming messages from connected peers
}
```

### 2. Credential Exchange
Issue and verify credentials with connected users:
- Build credential issuance workflows
- Verify incoming credentials
- Create trust networks

### 3. Advanced Discovery
Implement more sophisticated discovery patterns:
- Multiple QR codes with different timeouts
- Discovery subscriptions and notifications
- Connection management and monitoring

### 4. Production Features
Add production-ready capabilities:
- Proper error handling and reconnection
- Logging and monitoring for connection events
- User interface for connection management
- Testing with real Self mobile apps

## Testing with Mobile Apps

1. **Run this example** to generate a QR code
2. **Open a Self mobile app** on your phone
3. **Scan the QR code** displayed in the terminal
4. **Connection establishes automatically** 
5. **Test secure messaging** between the apps

## Best Practices Demonstrated

1. **Clean Code Structure**: Focused functions with single responsibilities
2. **Proper Resource Management**: Account cleanup with defer
3. **Error Handling**: Graceful handling of demo environment limitations
4. **Core SDK Usage**: Direct use of underlying SDK components
5. **Real Functionality**: Generates working QR codes for testing

## Files

- `main.go`: Clean connection example with QR generation (~100 lines)
- `go.mod`: Module definition with core SDK dependency
- `README.md`: This comprehensive documentation

## Dependencies

- Self SDK core library (`github.com/joinself/self-go-sdk/account`)
- Self SDK messaging (`github.com/joinself/self-go-sdk/message`) 
- Self SDK events (`github.com/joinself/self-go-sdk/event`)
- Standard Go libraries only

This example provides a clean, functional foundation for building Self SDK applications that can connect with mobile apps via QR code scanning.
