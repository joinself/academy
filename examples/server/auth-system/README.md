# Self Authentication System 🔐

> **🎯 Educational Goal**: Learn to build production-ready, multi-user authentication using Self SDK with proper request/response correlation and channel management.

## 🎓 What You'll Learn

By the end of this guide, you'll understand:
- **Multi-User Authentication**: How to handle multiple concurrent users without session mixing
- **Request/Response Correlation**: Using `content.ID()` and `ResponseTo()` for perfect user isolation
- **Channel Management**: Establishing dedicated communication channels for each user
- **Production Architecture**: Building secure, scalable authentication systems

## 🚀 Quick Start (5 Minutes)

### Prerequisites
- Go 1.24+ installed
- Self Mobile App ([Android](https://play.google.com/store/apps/details?id=com.joinself.app) | [iOS](https://apps.apple.com/app/self/id1516937530))
- Self account created in the mobile app

### Run the Authentication System

```bash
# 1. Start the server
go run cmd/server/main.go

# 2. In another terminal, request authentication
curl -X POST http://localhost:8081/api/v1/auth/request \
     -H "Content-Type: application/json" \
     -d '{}'

# 3. Scan the QR code with your Self mobile app
# 4. Watch the logs - you'll see the enhanced correlation working!
```

## ⚙️ Configuration

The authentication system uses environment variables for configuration. For most users, you only need to set the storage key.

### 🔑 Essential Configuration

**Storage Key** (Required for persistent accounts):
```bash
# Generate a secure 32-byte storage key
openssl rand -base64 32

# Run with your storage key
SELF_AUTH_STORAGE_KEY="your-base64-encoded-32-byte-key" \
go run cmd/server/main.go
```

**⚠️ Security Note**: Keep your storage key secret and consistent across restarts. Never commit it to version control.

### 🔧 Advanced Configuration

For advanced users, additional environment variables are available:
- **Server settings**: `SELF_SERVER_PORT`, `SELF_SERVER_ADDRESS`, `SELF_SERVER_ENABLE_TLS`
- **Auth settings**: `SELF_AUTH_SESSION_TIMEOUT`, `SELF_AUTH_QR_EXPIRATION`, `SELF_AUTH_REQUIRED_CLAIMS`
- **Logging**: `SELF_AUTH_LOG_LEVEL`

See [`internal/config/config.go`](internal/config/config.go) for the complete list of supported environment variables and their defaults.

## 🔑 Core Innovation: Enhanced Multi-User Authentication

### The Problem This Solves

Traditional authentication systems struggle with:
- **Session Mixing**: Alice's authentication completing Bob's request
- **Race Conditions**: Multiple users scanning QR codes simultaneously
- **Poor Isolation**: Shared state causing cross-user contamination

### Our Solution: Mathematical User Isolation

This system uses **cryptographically unique request identification** combined with **SDK-native correlation** to ensure perfect user isolation:

```go
// 1. Generate unique discovery request
content, err := message.NewDiscoveryRequest().Finish()
contentID := hex.EncodeToString(content.ID()) // Cryptographically unique

// 2. Correlate discovery response
discoveryResponse, err := message.DecodeDiscoveryResponse(msg.Content())
responseToID := hex.EncodeToString(discoveryResponse.ResponseTo())
// Perfect match: responseToID == contentID

// 3. Establish dedicated channel
channel := &UserChannel{
    ID:                channelID,
    UserDID:           userDID,
    OriginalContentID: responseToID,
    Status:            ChannelActive,
}
```

### 🎯 What Just Happened

1. **Discovery Request**: Each QR code contains a unique `content.ID()`
2. **Response Correlation**: Mobile app responds with `ResponseTo()` matching the original ID
3. **Channel Establishment**: System creates dedicated communication channel for that user
4. **Credential Exchange**: All subsequent communication uses the established channel
5. **Session Creation**: Final session includes channel information for perfect isolation

## 🏗️ Architecture Overview

### Enhanced Authentication Flow

```mermaid
sequenceDiagram
    participant A as Alice's Browser
    participant B as Bob's Browser
    participant S as Server
    participant AM as Alice's Mobile
    participant BM as Bob's Mobile
    
    A->>S: POST /auth/request
    Note over S: Generate unique content.ID()
    S->>A: {qr_code, request_id}
    
    B->>S: POST /auth/request  
    Note over S: Generate different content.ID()
    S->>B: {qr_code, request_id}
    
    AM->>S: Scan Alice's QR → Discovery Response
    Note over S: Match ResponseTo() with Alice's ID<br/>Create Alice's channel
    
    BM->>S: Scan Bob's QR → Discovery Response
    Note over S: Match ResponseTo() with Bob's ID<br/>Create Bob's channel
    
    S->>AM: Credential request (Alice's channel)
    S->>BM: Credential request (Bob's channel)
    
    AM->>S: Credential response
    BM->>S: Credential response
    
    S->>A: ✅ Alice authenticated
    S->>B: ✅ Bob authenticated
```

### Key Components

**🔐 AuthService**: Core authentication logic with enhanced multi-user support
- Discovery request/response correlation
- User channel management
- Session isolation

**🌐 HTTP Server**: RESTful API with authentication middleware
- Authentication endpoints
- Channel management endpoints
- Protected resource examples

**🔗 User Channels**: Dedicated communication channels per user
- Channel lifecycle management
- Activity tracking
- Proper cleanup

## 📚 API Reference

### Authentication Endpoints

#### Start Authentication
```bash
POST /api/v1/auth/request
```

**Response**:
```json
{
  "request_id": "auth_req_abc123",
  "qr_code": "████ QR CODE ████",
  "expires_at": "2024-01-01T12:00:00Z"
}
```

#### Check Status
```bash
GET /api/v1/auth/status/{request_id}
```

**Response**:
```json
{
  "status": "completed",
  "session_id": "sess_xyz789",
  "user_did": "did:self:user123",
  "claims": {"email": "user@example.com"}
}
```

### Channel Management Endpoints

#### Get Active Channels
```bash
GET /api/v1/channels/active
```

#### Get Channel Info
```bash
GET /api/v1/channels/info
```

#### Close Channel
```bash
POST /api/v1/channels/close
```

### Protected Endpoints

#### User Profile
```bash
GET /api/v1/protected/profile
```

#### Protected Data
```bash
GET /api/v1/protected/data
```

## 🛡️ Security Guarantees

### Mathematical User Isolation

| Security Feature | Implementation | Guarantee |
|------------------|----------------|-----------|
| **Unique Request IDs** | `content.ID()` | Cryptographically unique |
| **Response Correlation** | `ResponseTo()` | SDK-native matching |
| **Channel Isolation** | Per-user channels | Zero cross-contamination |
| **Session Security** | Channel-based sessions | Perfect user isolation |

### Impossible Attack Scenarios

- ❌ **Cross-User Authentication**: Alice can't complete Bob's request
- ❌ **Session Hijacking**: Sessions tied to specific channels
- ❌ **Race Conditions**: Deterministic correlation prevents timing attacks
- ❌ **Identity Spoofing**: Cryptographic signatures verify authenticity

## 🎯 Production Integration

### Basic Integration

```go
// 1. Create authentication service
authService, err := auth.NewAuthService(auth.DefaultConfig(), logger)
if err != nil {
    log.Fatal(err)
}
defer authService.Close()

// 2. Create HTTP server
httpServer := server.NewServer(authService, server.DefaultServerConfig(), logger)

// 3. Start server
log.Fatal(httpServer.Start())
```

### Advanced Configuration

```go
// Custom authentication configuration
authConfig := &auth.Config{
    StoragePath:      "./production_storage",
    Environment:      *account.TargetProduction,
    SessionTimeout:   30 * time.Minute,
    QRCodeExpiration: 5 * time.Minute,
    RequiredClaims:   []string{"email", "name"},
}

// Custom server configuration
serverConfig := &server.Config{
    Address:         "0.0.0.0",
    Port:            8443,
    EnableTLS:       true,
    TLSCertFile:     "/path/to/cert.pem",
    TLSKeyFile:      "/path/to/key.pem",
    SessionKey:      loadSecretKey(),
    RequestTimeout:  30 * time.Second,
}
```

## 🧪 Testing Multiple Users

### Test Scenario 1: Concurrent Authentication

```bash
# Terminal 1: Alice requests authentication
curl -X POST http://localhost:8081/api/v1/auth/request

# Terminal 2: Bob requests authentication  
curl -X POST http://localhost:8081/api/v1/auth/request

# Both users can scan their respective QR codes
# System will handle them independently with zero cross-contamination
```

### Test Scenario 2: Channel Management

```bash
# After authentication, check active channels
curl -H "Cookie: session_cookie" http://localhost:8081/api/v1/channels/active

# Get specific channel info
curl -H "Cookie: session_cookie" http://localhost:8081/api/v1/channels/info

# Close channel when done
curl -X POST -H "Cookie: session_cookie" http://localhost:8081/api/v1/channels/close
```

## 🔍 Understanding the Logs

When you run the system, you'll see logs showing the enhanced correlation:

```
# Discovery request generated
Generated auth request: request_id=auth_req_123, content_id=abc123

# User connects and responds
Received discovery response: response_to=abc123

# Perfect correlation
Channel and connection established: user_did=user123, channel_id=channel_456

# Session creation with channel info
User authenticated successfully: session_id=sess_789, channel_id=channel_456
```

## 🚀 Next Steps

### 🟢 Beginner: Try the Basic Flow
1. Run the server and authenticate one user
2. Examine the logs to see the correlation working
3. Try the protected endpoints

### 🟡 Intermediate: Multiple Users
1. Test concurrent authentication with multiple users
2. Explore the channel management endpoints
3. Monitor active channels and their activity

### 🟠 Advanced: Production Integration
1. Integrate with your existing application
2. Customize the configuration for production
3. Add monitoring and alerting

### 🔴 Expert: Custom Extensions
1. Add custom credential requirements
2. Implement advanced channel management
3. Build custom authentication flows

## 📖 Learn More

- **[Self SDK Documentation](https://docs.joinself.com)**: Complete SDK reference
- **[Self Mobile Apps](https://joinself.com/download)**: Download for iOS and Android
- **[Academy Examples](../../)**: More authentication patterns and use cases

---

> **🎓 Key Takeaway**: This system demonstrates production-ready multi-user authentication with mathematical guarantees of user isolation. The enhanced architecture using proper request/response correlation and channel management eliminates all cross-user authentication vulnerabilities while maintaining the simplicity and security of passwordless authentication. 
