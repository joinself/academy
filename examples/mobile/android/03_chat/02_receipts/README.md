# 🔗 Chat Message Demo

> **📖 Learn the concepts first:** [Secure Connections](../../../../../../docs/concepts/secure-connections.md)
> **🎯 What you'll learn:** How to make chat messaging with Self SDK.

This example demonstrates **Chat Messaging** with Self SDK.

## 🟢 Complexity: Beginner

## 🚀 Quick Start

1. Install the app to your phone
```bash
./gradlew :03_chat:02_receipts:installDebug
```
2. Open `Self Academy` app and register a new account
3. Need a server to connect, run this command on a terminal
```
docker run --pull=always --rm -it ghcr.io/joinself/self-sdk-demo:java
```
4. Scan the QRCode to connect to server
5. Input a message and click on `Send`

### Expected Output

A screen with a `Create Account` button.
Once an account is registered, need to scan the QRCode to connect to server.
Then a `Send` button become available.

logs in logcat:
```
🔧 Initialize Self SDK...
✅ Connected to Self network

🔧 Open registration flow...
✅ Registration successfully!

🔧 Open QRCode flow...
✅ Server connected!!

✅ Received chat message
🔧 Send receipt message...
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

Open the QRCode flow in Jetpack Navigation
```kotlin
account?.openQRCodeFlow(
    onFinish = { qrCode, discoverData -> 
        
    }
)
```

### Step 5: Receive chat messages
Check if the received message is a chat message or a receipt message.

```kotlin
override fun onMessage(message: Message) {
    Log.d(LOGTAG, "onMessage: ${message.id()}")
    if (message is ChatMessage) {
        sendReceipt(message.id())
    } else if (message is Receipt) {
        Log.d(LOGTAG, "Receipt" +
                "\ndelivered: ${message.delivered()}" +
                "\n:read:${message.read()}")
    }
}
```

### Step 6: Send receipt messages
After receiving a chat message, send a receipt message to the sender.

```
Construct a receipt message with messages ids that have been received.
Send the receipt message to the sender.  
```

```kotlin
fun sendReceipt(messageId: String) {
    val receipt = Receipt.Builder()
        .setDelivered(listOf(messageId))
        .build()
    coroutineScope.launch(Dispatchers.IO) {
        account?.send(toAddress = groupAddress!!, message = receipt)
    }
}
```

## 🚀 Next Steps

1. **Advanced Features**: Explore `../../04_advanced_features` to back up the SDK data.

