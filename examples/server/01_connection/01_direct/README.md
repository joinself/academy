# Direct Connection Example - Address-Based Connections

> **Learn the concepts first:** [Secure Connection Concepts](../../../../docs/concepts/secure-connections.md)
> **What you'll learn:** How to establish secure connections using inbox addresses for programmatic server-to-server communication

This example demonstrates **DIRECT ADDRESS-BASED CONNECTIONS** with Self SDK. Perfect for server-to-server communication, APIs, and automated systems.

## Complexity: Beginner

---

## Quick Start

### Go
```bash
# Terminal 1: Start this server
cd go
go run main.go
# Copy the displayed inbox address

# Terminal 2: Connect from the client example
cd ../../03_client/go
go run main.go <paste-inbox-address-here>
# Watch the connection establish in both terminals!
```

### Java
```bash
# Terminal 1: Start this server
cd java
gradle run
# Copy the displayed inbox address

# Terminal 2: Connect from the client example
cd ../../03_client/java
gradle run <paste-inbox-address-here>
# Watch the connection establish in both terminals!
```

### Complete Server-Client Test (Recommended!)
The `../03_client` example is specifically designed to connect TO this server! This creates a full end-to-end connection demonstration.

---

## Direct vs QR Approach

| **Direct Addresses** (This Example) | **QR Codes** ([../02_qr/](../02_qr/)) |
|-------------------------------------|--------------------------------------|
| **Server-to-server** connections | **Mobile app** connections |
| **API integrations** | **User-facing** applications |
| **Automated systems** | **Visual discovery** |
| **Programmatic workflows** | **Camera-based** interaction |

---

## How It Works

### Step 1: Address Creation
The server creates a shareable "mailbox" for receiving connection requests:
```
Open account inbox for connections
Generate shareable inbox address
Display address for others to use
Wait for incoming connection requests
```

### Step 2: Connection Request Handling
When someone wants to connect, they send a key package:
```
Receive incoming key package from potential connection
Extract their connection details
Prepare to establish secure channel
```

### Step 3: Connection Establishment
The server accepts and establishes the connection:
```
Accept the connection request
Create secure communication channel
Exchange cryptographic keys
Confirm connection is ready for messaging
```

### Core Connection Flow
```
1. Server: Create inbox address
2. Server: Share address with potential connections
3. Client: Use address to send connection request
4. Server: Receive and process request automatically
5. Both: Establish secure encrypted channel
6. Both: Ready for secure messaging
```

---

## What You'll See

### Server Side Output
```
Direct Connection Example - Server Side
==========================================
DIRECT CONNECTION ADDRESS:
Address: did:self:inbox:9876543210fedcba...

Share the address above with other parties for direct connection
Waiting for connections... (Press Ctrl+C to exit)
```

### When Someone Connects
```
Connection request received from: did:self:connecting_party...
Successfully established encrypted connection!
Connection is now ready for secure messaging!
```

---

## Works with Client Example

**This server pairs with [../03_client/](../03_client/) for complete testing:**

| This Server Does | Client Example Does |
|------------------|-------------------|
| Creates inbox address | Takes address as input |
| Waits for connection requests | Sends connection requests |
| Accepts incoming connections | Establishes outgoing connections |
| Uses connection acceptance callbacks | Uses connection welcome callbacks |

---

## Use Cases

### Perfect for:
- **API endpoints and backend services**: Allow other services to connect securely
- **Microservice communication**: Establish secure channels between services
- **Command-line tools and automation**: Programmatic connection establishment
- **Email/chat address sharing**: Share addresses through existing channels
- **Development servers**: Test environments and local development

### Technical Benefits:
- **No visual UI required**: Pure programmatic connection
- **Automated acceptance**: Server can auto-accept connections
- **Scalable**: Handle multiple concurrent connection requests
- **Integration-friendly**: Easy to embed in existing systems

---

## Security Features

### Cryptographic Protection
- **End-to-end encryption**: All communication automatically encrypted
- **Key exchange**: Secure cryptographic key negotiation
- **Identity verification**: Each party's identity is cryptographically verified
- **Forward secrecy**: Past communications remain secure even if keys are compromised

### Access Control
- **Address-based**: Only parties with the address can initiate connections
- **Acceptance control**: Server can implement custom acceptance logic
- **Isolation**: Each connection is cryptographically isolated

---

## Next Steps

1. **Add Messaging** → [Chat Example](../../03_chat)
2. **Issue Credentials** → [Credentials Example](../../02_credentials)  
3. **Try QR Approach** → [QR Connection Example](../02_qr/)
4. **Multi-party Setup** → [Advanced Features](../../04_advanced_features)

---

## Troubleshooting

### Address Not Displaying
```
Failed to create inbox address
```
**Solution**: Ensure account is properly initialized and connected to network

### No Connections Received
```
Waiting for connections... (no activity)
```
**Solution**: 
- Verify the client is using the correct address
- Check network connectivity
- Ensure firewall allows connections

### Connection Failed
```
Failed to establish connection
```
**Solution**: 
- Verify both parties are on compatible SDK versions
- Check network connectivity between parties
- Ensure storage permissions are correct

---

**Ready to build programmatic Self connections?** This is your starting point!
