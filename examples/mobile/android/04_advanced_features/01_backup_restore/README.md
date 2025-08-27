# 🔗 Backup & Restore Demo

> **🎯 What you'll learn:** How to back up and restore data that stored in the SDK.

This example demonstrates **Backup and Restore** with Self SDK.

## 🟢 Complexity: Advanced

## 🚀 Quick Start

The backup data is stored in the user’s Google Drive, so users need to sign in to their Google account.
To sign in, Google Drive requires setting up the `Android OAuth Client ID` in the [Google Cloud Console](https://console.cloud.google.com/apis/credentials) with the SHA1 signing hash and the package name. 

1. Install the app to your phone
```bash
./gradlew :04_advanced_features:01_backup_restore:installDebug
```
2. Open `Self Academy` app and register a new account
3. Click on `Backup`

### Expected Output

A screen with a `Create Account` button.
Once an account is registered, `Backup` button become available.

logs in logcat:
```
🔧 Initialize Self SDK...
✅ Connected to Self network

🔧 Open registration flow...
✅ Registration successfully!

🔧 Start backup flow...
✅ Backup successfully!
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

### Step 4: Back up
After registration, users can start back up flow
```
The SDK guides the user to sign in to Google Drive. 
Then, it serializes and compresses the data. 
The backup is encrypted with the user’s identity credential, allowing the user to restore the data with only a livenesss check. 
Finally, the SDK uploads the encrypted data to Google Drive. 
```

Open the flow in Jetpack Navigation
```kotlin
account?.openBackupFlow(onFinish = { isSuccess, error ->
    if (isSuccess) {
        Log.d(LOGTAG, "✅ Backup successfully")
    }
})
```

### Step 5: Restore
Users can restore the data without registration.
```
The SDK guides the user to sign in to Google Drive and download the encrypted data. 
Subsequently, users must complete and pass a live check to verify that the backup belongs to them.
The SDK then utilizes the liveness credential to decrypt the backup. 
Finally, if the check passes, the SDK will restore the backup.
```

Open the flow in Jetpack Navigation
```kotlin
account?.openRestoreFlow(onFinish = { isSuccess, error ->
    if (isSuccess) {
        Log.d(LOGTAG, "✅ Restore successfully")
    }
})
```

## 🚀 Next Steps