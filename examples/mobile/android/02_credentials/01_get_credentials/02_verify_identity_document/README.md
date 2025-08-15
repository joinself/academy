# 🔗 Identity Credential Verification Demo

> **📖 Learn the concepts first:** [Verifiable Credentials Concepts](../../../../../../docs/concepts/verifiable-credentials.md)
> **🎯 What you'll learn:** How to get verified credentials from an identity document.

This example demonstrates **Create Credentials** with Self SDK by using an identity document.

## 🟢 Complexity: Beginner

## 🚀 Quick Start

1. Install the app to your phone
```bash
./gradlew :02_credentials:01_get_credentials:02_verify_identity_document:installDebug
```
2. Open `Self Academy` app and register a new account
4. Click on `Verify Identity Document`

### Expected Output

A screen with a `Create Account` button.

Once an account is registered, a `Verify Identity Document` button become available.


```
🔧 Initialize Self SDK...
🔍 Checking for account storage path...
🔧 Creating new Self account...

✅ Connected to Self network

🔧 Open registration flow...
✅ Registration successfully!

🔧 Open document verification flow...
✅ Identity document is verified successfully!
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

### Step 4: Verify Identity Document
```
An identity document verification flow is initiated by the SDK.
The SDK guides users to capture the images of their documents and scan the NFC chip. 
Subsequently, the SDK sends captured documents to the server to verify.  
Finally, the identity credentials are verified and returned from the server.
```

## 🚀 Next Steps

After creating your account:

1. **Share credentials**: Try `../02_credentials/02_share_credentials` to share verified credentials
2. **Send messages**: Explore `../../03_chat` for messaging capabilities

