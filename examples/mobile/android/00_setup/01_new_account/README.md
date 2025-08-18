# 🆕 Mobile New Account - Your First Self Account

> **📖 Learn the concepts first:** [Decentralized Identity Concepts](../../../../docs/concepts/decentralized-identity.md)  
> **🎯 What you'll learn:** How to create a brand new Self account from scratch with proper storage and network registration

This example demonstrates the **first-time setup** process for a Self SDK account. Perfect for new users, development environments, and clean demo setups.

## 🟢 Complexity: Beginner
**Perfect starting point** - This is the foundation for all other Self SDK operations.

## 🚀 Quick Start

1. Install the app to your phone
2. Open Self Academy app in your phone

```bash
./gradlew :00_setup:01_new_account:installDebug
```

### Expected Output

The liveness screen opens and guides you through completing the liveness check.

```
🔍 Checking for account storage path...
🔧 Creating new Self account...

✅ Connected to Self network

🔧 Open registration flow...
✅ Registration flow finished
✅ New Self account ready!
```


## 🏗️ How It Works

### Step 1: SDK initialization
First, the Self SDK needs to be initialized.
```
Set up the pushToken to receive push notifications
Append the SDK logs to your logger
```

### Step 2: Account Storage Check
Next, check and create a directory for account storage
```
Check for existing account storage
If exists: existing account will be loaded
If not: new account will be created
```

### Step 3: Account Initialization and UI Flow Integration
The SDK creates a completely new Self account:
```
Initialize new account with default configuration
Register with Self network

Integrate the built-in Self-UI Flows into the main app’s navigation.
```

### Step 4: Account Registration Status
Once initialized, check registration status
```
Retrieve the registration status of the user using `account.registered()` to display the UI accordingly. 
If the user is not registered, display a button that allows them to register the app. 
If the user is registered, disable the registration function.
```

### Step 5: Account Registration Flow
The liveness will be open, and guide user finish the flow
```
To ensure that users are real people, a live check is conducted. 
An image is captured and submitted to the server for secure verification. 
This image will be used to verify the user’s identity for subsequent operations.
```

### Step 6: Account Registered
```
After the user completes registration, handle the UI accordingly. 
Then, start using other functions, such as connecting to the server.
```

## 🚀 Next Steps

After creating your account:

1. **Connect with a server**: Try `../01_connection` to connect with server.
2. **Send messages**: Explore `../../03_chat` for messaging capabilities


