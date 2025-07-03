# Advanced Notifications Example

> **📖 Learn the concepts first:** [System Architecture](../../../../docs/architecture/system-overview.md)

A comprehensive demonstration of the Self SDK's push notification system for real-time user engagement.

## 🚀 Quick Start

```bash
# Run the notifications demo
go run main.go
```

The demo automatically showcases all Self SDK notification features in a structured, educational format!

## 📊 Complexity Rating

**4/10** (Intermediate) - Perfect for learning notification systems

- 🟡 **Notification concepts**: Push notifications, event handling, delivery tracking
- 🟡 **User engagement**: Real-time alerts and communication
- 🟢 **Event management**: Callbacks and status monitoring
- 🟢 **Integration**: Seamless integration with other SDK components

## 🎯 What This Example Demonstrates

### Core Notification Features
- ✅ **Chat Notifications** - Message alerts and conversation updates
- ✅ **Credential Notifications** - Identity and verification alerts
- ✅ **Group Notifications** - Team and group activity alerts
- ✅ **Custom Notifications** - Application-specific notifications
- ✅ **Event Handling** - Delivery status and response management

### Educational Learning Path
1. **Notification Setup** - Configure notification handlers and events
2. **Notification Types** - Explore different notification categories
3. **Delivery Tracking** - Monitor notification status and delivery
4. **Event Management** - Handle notification responses and callbacks

## 🏃‍♂️ How to Run

### Single Command Demo
```bash
go run main.go
```

The demo runs automatically and demonstrates:
- Setting up notification event handlers
- Sending various types of notifications
- Tracking notification delivery status
- Managing notification responses and callbacks

### What Happens Automatically
1. **Client Creation**: Notification-focused Self client setup
2. **Handler Setup**: Configure notification event handlers
3. **Notification Types**: Demonstrate chat, credential, group, and custom notifications
4. **Delivery Tracking**: Monitor notification status and responses
5. **Event Management**: Handle notification callbacks and events

## 📋 What You'll See

```
🔔 Advanced Notifications Demo
==============================
This demo showcases Self SDK notification capabilities.

🔧 Setting up notification client...
✅ Notification client created successfully
🆔 Client DID: did:self:notifications123...

🔹 Setting up Notification Handlers
===================================
📨 Configuring notification event handlers...
   ✅ Notification sent handler configured
   ✅ Notification delivered handler configured
   ✅ Notification response handler configured

🔹 Chat Notifications
=====================
💬 Sending chat notification...
   ✅ Chat notification sent successfully
   📨 Notification delivered to peer
   ⏰ Delivery time: 15:30:45

🔹 Credential Notifications
===========================
🆔 Sending credential notification...
   ✅ Credential notification sent successfully
   📨 Notification delivered to peer
   ⏰ Delivery time: 15:30:46

🔹 Group Notifications
======================
👥 Sending group invite notification...
   ✅ Group notification sent successfully
   📨 Notification delivered to peer
   ⏰ Delivery time: 15:30:47

🔹 Custom Notifications
=======================
🔔 Sending custom notification...
   ✅ Custom notification sent successfully
   📨 Notification delivered to peer
   ⏰ Delivery time: 15:30:48
```

## 🔍 Key Code Sections

| Function | Purpose |
|----------|---------|
| `main()` | Step-by-step notification demo orchestration |
| `createNotificationClient()` | Notification-focused Self SDK client setup |
| `setupNotificationHandlers()` | Configure notification event handlers |
| `demonstrateChatNotifications()` | Chat and messaging notifications |
| `demonstrateCredentialNotifications()` | Identity and verification notifications |
| `demonstrateGroupNotifications()` | Group and team activity notifications |
| `demonstrateCustomNotifications()` | Application-specific notifications |

## 🎓 Educational Notes

### Notification Concepts
- **Push Notifications**: Real-time alerts delivered to users
- **Event Handling**: Callbacks for notification delivery and responses
- **Delivery Tracking**: Monitor notification status and success rates
- **User Engagement**: Keep users informed and engaged with your application

### Notification Types
- **Chat Notifications**: Message alerts and conversation updates
- **Credential Notifications**: Identity verification and credential alerts
- **Group Notifications**: Team invitations and group activity
- **Custom Notifications**: Application-specific alerts and updates

### Benefits
- **Real-Time Communication**: Instant alerts and updates
- **User Engagement**: Keep users informed and active
- **Delivery Assurance**: Track notification success and failures
- **Flexible Integration**: Works seamlessly with other SDK components

## 🔧 Customization Ideas

Try modifying the code to:
- Create custom notification types for your application
- Implement notification scheduling and timing
- Add notification preferences and user settings
- Create notification templates and formatting
- Implement notification analytics and tracking

## 🚀 Next Steps

After understanding this example, continue with:

| Next Example | Complexity | Description |
|-------------|------------|-------------|
| **Pairing** | 5/10 | Multi-device synchronization |
| **Production Patterns** | 6/10 | Real-world storage patterns |
| **Integration** | 7/10 | Multi-component workflows |

## 🛠️ Prerequisites

- Go 1.19 or later
- Self SDK dependencies (handled by go.mod)
- Basic understanding of event-driven programming
- Completion of simple_chat and group_chat examples
- Understanding of storage example (recommended)

## 💡 Troubleshooting

**Notification Issues:**
- Ensure notification handlers are set up before sending
- Check network connectivity for notification delivery
- Verify target DIDs are valid and reachable

**Event Handling Issues:**
- Confirm event handlers are properly configured
- Check callback function implementations
- Verify event handler registration timing

**Build Issues:**
- Run `go mod tidy` to ensure dependencies
- Check Go version with `go version`
- Verify you're in the correct directory

## 🎯 Key Differences from Other Examples

| Feature | Simple Chat | Group Chat | Storage | **Notifications** |
|---------|-------------|------------|---------|-------------------|
| **Focus** | Basic messaging | Group coordination | Data persistence | **User engagement** |
| **Real-time** | Basic | Group events | None | **Push notifications** |
| **Complexity** | 4/10 | 5/10 | 5/10 | **4/10** |
| **Event Handling** | Basic | Group events | None | **Advanced callbacks** |
| **User Experience** | Chat only | Group chat | Data storage | **Proactive alerts** |

## 🔔 Notification Architecture

### Notification Flow
```
Application → Self SDK → Notification System → Target Device
     ↓              ↓              ↓              ↓
Event Handler ← Delivery Status ← Network ← User Response
```

### Event Types
- **Sent**: Notification was sent from your application
- **Delivered**: Notification reached the target device
- **Response**: User interacted with the notification
- **Failed**: Notification delivery failed

This example provides the foundation for real-time user engagement in Self SDK applications! 🔔 
