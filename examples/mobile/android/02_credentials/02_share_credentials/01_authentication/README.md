# 🔗 Custom Credential Demo

> **📖 Learn the concepts first:** [Verifiable Credentials Concepts](../../../../../../docs/concepts/verifiable-credentials.md)
> **🎯 What you'll learn:** How to share verified credentials that stored in SDK.

This example demonstrates **Share Credentials** with Self SDK.

## 🟢 Complexity: Beginner

## 🚀 Quick Start

1. Install the app to your phone
```bash
./gradlew :02_credentials:02_share_credentials:01_authentication:installDebug
```
2. Open `Self Academy` app and register a new account
3. Need a server to connect, run this command on a terminal
```
docker run --pull=always --rm -it ghcr.io/joinself/self-sdk-demo:java
```
4. Scan the QRCode to connect to server
5. Click on `Start Authentication`

### Expected Output

A screen with a `Create Account` button.
Once an account is registered, need to scan the QRCode to connect to server.
Then a `Start Authentication` button become available.

```
🔧 Initialize Self SDK...
🔍 Checking for account storage path...
🔧 Creating new Self account...

✅ Connected to Self network

🔧 Open registration flow...
✅ Registration successfully!

🔧 Start connecting...
✅ Connection established successfully!

🔧 Start sharing credentials...
✅ Authentication successfully!
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

### Step 4: Scan a QRCode to connect to Server
```
A QRCode flow is initiated, and the SDK scans and parses the provided QRCode data. 
Using the parsed data, the SDK negotiates the connection. 
```

### Step 5: Authenticate
```
The server sends a Credential Request message to the mobile app, which includes a liveness query. 
The SDK then displays a UI to prompt the user to confirm or reject the liveness check. 
If the user accepts, the Liveness check flow begins, verifying the user’s identity. 
The verified credentials are then returned to the requester.
```

## 🚀 Next Steps

1. **Digital Signature**: Try `../03_digital_signatures` to learn how to sign a signature.
2. **Send messages**: Explore `../../03_chat` for messaging capabilities

