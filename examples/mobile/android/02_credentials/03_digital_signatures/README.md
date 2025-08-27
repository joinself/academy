# 🔗 Digital Signature Demo

> **📖 Learn the concepts first:** [Verifiable Credentials Concepts](../../../../../../docs/concepts/verifiable-credentials.md)
> **🎯 What you'll learn:** How to use digital signatures with Self SDK.

This example demonstrates **Digital Signature Credential** with Self SDK.

## 🟢 Complexity: Beginner

## 🚀 Quick Start

1. Install the app to your phone
```bash
./gradlew :02_credentials:03_digital_signatures:installDebug
```
2. Open `Self Academy` app and register a new account
3. Need a server to connect, run this command on a terminal
```
docker run --pull=always --rm -it ghcr.io/joinself/self-sdk-demo:java
```
4. Scan the QRCode to connect to server
5. Click on `Start Document Sign`

### Expected Output

A screen with a `Create Account` button.
Once an account is registered, need to scan the QRCode to connect to server.
Then a `Start Document Sign` button become available.

```
🔧 Initialize Self SDK...
✅ Connected to Self network

🔧 Open registration flow...
✅ Registration successfully!

🔧 Open QRCode flow...
✅ Server connected!!

🔧 Start signing...
✅ Received request message
✅ Display signing request
✅ Digital signature successfully!
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

### Step 5: Sign Document
```
The server sends a Verification Request message to the mobile app, which includes a document to sign.
The app needs to display the document to the user to read. 
The SDK then displays a UI to prompt the user to confirm or reject the signing process. 
If the user accepts, the Liveness check flow begins, verifying the user’s identity.
Then verified credentials are returned to the requester.
```

Listen to the request message from the server.
```kotlin
override fun onMessage(message: Message) {
    // check if it is verification request message.
    if (message is VerificationRequest) {
        requestMessage = message
    }
}
```

Get the document data and display to user
```kotlin
    val data = requestMessage.evidence()
```

Utilize this Jetpack Compose component to display the request message, and the SDK will handle the rest.
```kotlin
account?.DisplayRequestUI(selfModifier, requestMessage, 
    onFinish = { isSent, status ->
        Log.d(LOGTAG, "✅ Digital signature successfully!")
    }
)
```

## 🚀 Next Steps

2. **Send messages**: Explore `../../03_chat` for messaging capabilities

