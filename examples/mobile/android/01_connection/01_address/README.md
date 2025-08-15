# 🆕 Direct Connection Example - Address-Based Connections

> **Learn the concepts first:** [Secure Connection Concepts](../../../../docs/concepts/secure-connections.md)
> **What you'll learn:** How to establish secure connections using inbox addresses for programmatic client-to-server communication

This example demonstrates **DIRECT ADDRESS-BASED CONNECTIONS** with Self SDK. Perfect for client-to-server communication, APIs, and automated systems.

## 🟢 Complexity: Beginner

## 🚀 Quick Start

1. Install the app to your phone
```bash
./gradlew :01_connection:01_address:installDebug
```
2. Open `Self Academy` app and register a new account
3. Need a server to connect, run this command on a terminal
```
docker run --pull=always --rm -it ghcr.io/joinself/self-sdk-demo:java
```
4. Copy the address and paste it into the mobile app
5. Click `Connect` button

### Expected Output

A screen with a `Create Account` button. 

Once an account is registered, an input field for `address` and a `Connect` button become available.


```
🔧 Initialize Self SDK...
🔍 Checking for account storage path...
🔧 Creating new Self account...

✅ Connected to Self network

🔧 Open registration flow...
✅ Registration successfully!

🔧 Start connecting...
✅ Connection established successfully!
```


## 🏗️ How It Works

### Step 1: SDK initialization
First, the Self SDK needs to be initialized.
```
Set up the pushToken to receive push notifications
Append the SDK logs to your logger
```

### Step 2: Account Initialization and UI Flows Integration
Before using the account and UI flows, initialize them.
```
Initialize new account with default configuration
Register with Self network

Integrate the built-in Self-UI Flows into the main app’s navigation.
```

### Step 3: Account Registration
Once initialized, need to register an account before connecting to server
```
Users must pass the liveness check to be considered real people.
The server returns verified liveness credentials.
```

### Step 4: Connect to another Self Account
```
Enter an address string manually. 
The SDK internally negotiates the connection using the provided address. 
Once the negotiation is successful, the SDK returns a group address.
```

### Step 5: Connection Connected
```
Start sending chat message
Receive credentials requests, document sign,...
```


## 🚀 Next Steps

After creating your account:

1. **Verify credentials**: Try `../02_credentials` to verify identity documents
2. **Send messages**: Explore `../../03_chat` for messaging capabilities

