# 🔗 Share Custom Credential Demo

> **📖 Learn the concepts first:** [Verifiable Credentials Concepts](../../../../../../docs/concepts/verifiable-credentials.md)
> **🎯 What you'll learn:** How to share verified identity credentials that stored in SDK.

This example demonstrates **Share Credentials** with Self SDK.

## 🟢 Complexity: Intermediate

## 🚀 Quick Start

1. Install the app to your phone
```bash
./gradlew :02_credentials:02_share_credentials:04_custom_credential:installDebug
```
2. Open `Self Academy` app and register a new account
3. Need a server to connect, run this command on a terminal
```
docker run --pull=always --rm -it ghcr.io/joinself/self-sdk-demo:java
```
4. Scan the QRCode to connect to server
5. `Get Credentials` before sharing credentials
6. Click on `Start Sharing Custom Credential`

### Expected Output

A screen with a `Create Account` button.
Once an account is registered, need to scan the QRCode to connect to server.
After that, a button `Get Credentials` is available.
Finally, `Start Sharing Custom Credential` button become available.

```
🔧 Initialize Self SDK...
✅ Connected to Self network

🔧 Open registration flow...
✅ Registration successfully!

🔧 Open QRCode flow...
✅ Server connected!!

🔧 Start getting credentials...
✅ Received credential message
✅ Display credential message
✅ Store custom credential successfully!

🔧 Start sharing credentials...
✅ Received request message

✅ Display request message
✅ Sharing custom credential successfully!
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
The server verifies their liveness credentials and uses them in subsequent operations.
```

### Step 4: Scan a QRCode to connect to Server
```
A QRCode flow is initiated, and the SDK scans and parses the provided QRCode data. 
Using the parsed data, the SDK negotiates the connection. 
```

### Step 5: Get Custom Credential
Before sharing, the credentials must be present in the SDK. Users need to get them from the server.
``` 
The server verifies the custom credentials and returns CredentialMessage to the SDK, which stores them for sharing.
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

### Step 6: Share Identity Credential
The SDK will handle the request, display the UI, and then return the response to the requester.
```
The server sends a Credential Request message to the mobile app, which includes credential queries. 
The SDK then displays a UI to prompt the user to confirm or reject the request. 
If the user accepts, the stored credentials are shared with the server.
```

Utilize this Jetpack Compose component to display the request message, and the SDK will handle the rest.
```kotlin
account?.DisplayRequestUI(selfModifier, requestMessage, 
    onFinish = { isSent, status -> }
)
```

## 🚀 Next Steps

1. **Digital Signature**: Try `../03_digital_signatures` to learn how to sign a signature.
2. **Send messages**: Explore `../../03_chat` for messaging capabilities

