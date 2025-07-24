# **Quick Start:** Advanced Features Examples

## What You'll Learn

Welcome to the final set of examples, where we transition from foundational concepts to **production-ready patterns**. These advanced examples demonstrate how to build robust, scalable, and feature-rich applications using the Self SDK.

****What you will learn:** Learning Goals:**
- Master production patterns for account and storage management
- Implement multi-device synchronization with account pairing
- Build real-time user engagement with push notifications
- Understand graceful shutdown, resource management, and monitoring
- Connect all foundational concepts into production applications

## **Architecture:** Progressive Learning Path

These examples build upon everything you've learned, applying foundational knowledge to advanced use cases.

| **What you will learn:** Example | **Learning:** Complexity | **Resources:** Focus |  
|------------|---------------|----------|
| **🔴 [Core Production Features](#core-production-features)** | Advanced (7/10) | Encrypted storage, production patterns, core SDK control |
| **🟡 [Device Pairing](#device-pairing--synchronization)** | Intermediate (5/10) | Multi-device account synchronization |
| **🟡 [Push Notifications](#push-notifications--user-engagement)** | Intermediate (4/10) | Real-time user engagement and alerts |

---

## 🔴 Core Production Features

**Path:** [`examples/server/04_advanced_features/01_core_features/`](../../examples/server/04_advanced_features/01_core_features/)

This is the capstone example for core SDK functionality, demonstrating how to build a robust service with production-ready patterns for storage, messaging, and resource management.

### Quick Start

```bash
cd examples/server/04_advanced_features/01_core_features/go
go run main.go
```

**What happens:**
1. **Advanced Account Setup**: Creates an account with production-focused configuration.
2. **Encrypted Storage Demo**: Stores and retrieves various data types using the SDK's secure storage.
3. **Discovery QR**: Generates a QR code for connections.
4. **Real-time Monitoring**: Listens for connections and messages, providing status updates.
5. **Graceful Shutdown**: Catches `Ctrl+C` to display a summary before exiting.

### How It Works

#### 1. Production-Ready Account Configuration

This example demonstrates advanced configuration for production environments:

```go
// From: 01_core_features/go/main.go
func (d *AdvancedDemo) setupAdvancedAccount() {
    cfg := &account.Config{
        StorageKey:  generateStorageKey("advanced_demo"), // Secure random key
        StoragePath: "./advanced_demo_storage",
        Environment: account.TargetSandbox, // Use TargetProduction in production
        LogLevel:    account.LogWarn, // Reduce log noise
        Callbacks: account.Callbacks{
            OnWelcome:    d.onWelcomeMessage,
            OnKeyPackage: d.onKeyPackage,
            OnMessage:    d.onMessage,
        },
    }
    // ...
}
```

****Key Concepts:** Key Concepts:**
- **Secure Key Generation**: `generateStorageKey` uses `crypto/rand` for secure key material.
- **Environment Configuration**: Easily switch between `sandbox` and `production` environments.
- **Logging Control**: Set log levels to reduce noise in production.
- **Comprehensive Callbacks**: Handle all core event types for a robust application.

#### 2. Advanced Encrypted Storage

Go beyond simple key-value storage by using JSON to store complex, structured data securely:

```go
// From: 01_core_features/go/main.go
func (d *AdvancedDemo) demonstrateAdvancedStorage() {
    // Storing a complex object
    userProfile := map[string]interface{}{
        "name":     "Advanced User",
        "role":     "Developer",
        "verified": true,
    }
    jsonData, _ := json.Marshal(userProfile)
    d.account.ValueStore("user_profile", jsonData)

    // Retrieving and unmarshaling
    retrievedData, _ := d.account.ValueLookup("user_profile")
    var retrievedProfile map[string]interface{}
    json.Unmarshal(retrievedData, &retrievedProfile)
}
```

****Key Concepts:** Key Concepts:**
- **Structured Data**: Store complex objects, not just strings.
- **JSON Serialization**: A standard, flexible format for data exchange.
- **Encrypted at Rest**: The SDK automatically encrypts all data before writing to storage.
- **Data Integrity**: Storage is protected from tampering.

#### 3. Graceful Shutdown & Resource Management

Production services must shut down cleanly. This example shows how to catch system signals to perform cleanup and report final statistics:

```go
// From: 01_core_features/go/main.go
func (d *AdvancedDemo) setupGracefulShutdown() {
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt, syscall.SIGTERM)

    go func() {
        <-c
        fmt.Println("\n🏁 Demo completed by user")
        d.done <- true // Signal the main loop to exit
    }()
}

func (d *AdvancedDemo) waitForInteractions() {
    // ... main loop ...
    <-d.done // Wait for shutdown signal
}
```

****Key Concepts:** Key Concepts:**
- **Signal Handling**: Intercept `Ctrl+C` (`SIGINT`) and `SIGTERM`.
- **Clean Cleanup**: Ensures `defer d.account.Close()` is called.
- **Final Reporting**: `displaySummary()` provides metrics on exit.
- **Resource Safety**: Prevents corrupted state from abrupt termination.

---

## 🟡 Device Pairing & Synchronization

**Path:** [`examples/server/04_advanced_features/02_pairing/`](../../examples/server/04_advanced_features/02_pairing/)

This example demonstrates how to manage a single identity across multiple devices, a crucial feature for modern applications.

### What You'll Learn
- **Multi-Device Identity**: How one DID can be controlled by multiple synchronized devices.
- **Pairing Codes & QR**: Generate secure codes for adding new devices.
- **State Synchronization**: How account state (connections, credentials) is managed across devices.
- **Device Management**: Patterns for listing and removing paired devices.

**How it works:** The SDK uses cryptographic protocols to securely "pair" a new device to an existing account. The new device generates its own keypair, and through a consent workflow (usually involving a QR code scan on an existing device), it's added to the account's MLS group. This allows it to send and receive messages on behalf of the user's DID.

---

## 🟡 Push Notifications & User Engagement

**Path:** [`examples/server/04_advanced_features/03_notifications/`](../../examples/server/04_advanced_features/03_notifications/)

Learn how to build real-time user engagement systems with push notifications for various events.

### What You'll Learn
- **Notification Systems**: How to trigger alerts for chat messages, credential requests, and custom events.
- **Targeted Messaging**: Send notifications for specific scenarios like user onboarding or re-engagement.
- **Delivery Tracking**: Use SDK features to monitor notification status.
- **Custom Payloads**: Craft custom notification content for different use cases.

**How it works:** This example uses the core SDK messaging capabilities to send specially formatted messages that a mobile or client application can interpret as push notifications. It demonstrates patterns for managing different notification types, tracking their delivery, and storing notification history, all over the secure, end-to-end encrypted channels you've learned about.

## **Learning:** What Just Happened? Production-Readiness Achieved

### **Success:** **Advanced Pattern Mastery**

You now understand the patterns required to move from prototypes to production:

- ****Architecture:** Robust Services**: Building applications that handle real-world conditions like graceful shutdowns.
- ****Security:** Secure Data Management**: Using the SDK's encrypted storage for complex application data.
- **📱 Modern User Features**: Implementing multi-device support and push notifications.

### **Success:** **Connecting Foundations to Production**

These examples show how foundational concepts become production features:

| Foundational Concept | Production Feature | Advanced Example |
|----------------------|--------------------|------------------|
| Encrypted Storage | Storing JSON Objects | Core Features |
| Event Callbacks | Graceful Shutdown | Core Features |
| Cryptographic Identity (DID) | Multi-Device Sync | Pairing |
| Secure Messaging (MLS) | Push Notifications | Notifications |

## **Resources:** Next Steps

You have now completed the entire **Examples** section of the Academy!

1.  ****Success:** [Setup](setup.md)**
2.  ****Success:** [Connections](connections.md)**
3.  ****Success:** [Credentials](credentials.md)**
4.  ****Success:** [Chat](chat.md)**
5.  ****Success:** [Advanced Features](advanced.md)**

With this comprehensive knowledge, you are fully equipped to:
- **Build Your Own Applications**: Use these patterns as a starting point.
- **Contribute to the Ecosystem**: Your expertise can help improve the SDK and documentation.

## 💡 Key Takeaways

- **Production is About More Than Just Logic**: It requires robust error handling, resource management, and secure configuration.
- **The SDK Provides Powerful Primitives**: Features like encrypted storage are building blocks for complex application state.
- **User Experience Matters**: Multi-device support and notifications are key to modern applications.
- **You Are Ready to Build**: You have progressed from "Hello World" to production-ready patterns.

Congratulations on completing the Joinself Academy examples! You now have the knowledge and practical experience to build the next generation of secure, decentralized applications. **Quick Start:**
