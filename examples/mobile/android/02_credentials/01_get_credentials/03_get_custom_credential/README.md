# 🔗 Custom Credential Demo

> **📖 Learn the concepts first:** [Verifiable Credentials Concepts](../../../../../../docs/concepts/verifiable-credentials.md)
> **🎯 What you'll learn:** How to get verified credentials from server.

This example demonstrates **Get Credentials** with Self SDK with a server.

## 🟢 Complexity: Beginner

## 🚀 Quick Start

1. Install the app to your phone
```bash
./gradlew :02_credentials:01_get_credentials:03_get_custom_credential:installDebug
```
2. Open `Self Academy` app and register a new account
3. Need a server to connect, run this command on a terminal
```
docker run --pull=always --rm -it ghcr.io/joinself/self-sdk-demo:java
```
4. Scan the QRCode to connect to server
5. Click on `Get Credentials`

### Expected Output

A screen with a `Create Account` button.
Once an account is registered, need to scan the QRCode to connect to server.
Then a `Get Credentials` button become available.

logs in logcat:
```
🔧 Initialize Self SDK...
✅ Connected to Self network

🔧 Open registration flow...
✅ Registration successfully!

🔧 Open QRCode flow...
✅ Server connected!!

🔧 Start getting credentials...
✅ Received credential message
✅ Get Credentials successfully!
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

### Step 5: Get Custom Credentials
```
The server sends a Credential Message to the mobile app, which includes custom credentials. 
The SDK then displays a UI to prompt the user to confirm or reject storing these credentials. 
If the user accepts, the credentials are stored in the local database. Otherwise, they are discarded.
```

Listen to the message from the server.
```kotlin
override fun onMessage(message: Message) {
    // check if it is custom credential message.
    if (message is CredentialMessage) {
        credentialMessage = message
    }
}
```

Users confirm storing the credentials in the SDK.
```kotlin
account?.DisplayRequestUI(selfModifier, credentialMessage, 
    onFinish = { isSent, status ->
        Log.d(LOGTAG, "Custom credentials are stored!!")
    }
)
```

## 🚀 Next Steps

1. **Share credentials**: Try `../02_credentials/02_share_credentials` to share verified credentials
2. **Send messages**: Explore `../../03_chat` for messaging capabilities

