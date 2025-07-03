# Chat Examples

> **📖 Learn the concepts first:** [Message Layer Security](../../../docs/concepts/message-layer-security.md)

This directory contains comprehensive examples for implementing chat and messaging functionality using the Self SDK, organized by complexity and feature set for progressive learning.

## 🗂️ Examples Organization

| 🎯 Level | 📁 Directory | 🎓 Focus Area |
|----------|--------------|---------------|
| **Basic Chat** | `01_basic/` | Core messaging and response patterns |
| **Group Chat** | `02_group/` | Multi-participant messaging (Coming Soon) |
| **Rich Media** | `03_rich_media/` | File attachments and media sharing (Coming Soon) |
| **Advanced Features** | `04_advanced/` | Typing indicators, read receipts, etc. (Coming Soon) |

## 🚀 Quick Start by Use Case

### 💬 **Simple Messaging** (Basic chat functionality)
```bash
cd 01_basic/
go run main.go
```

### 👥 **Group Communication** (Multi-user chat) - Coming Soon
```bash
cd 02_group/
go run main.go
```

### 📎 **Rich Content** (Files and media) - Coming Soon
```bash
cd 03_rich_media/
go run main.go
```

## 📚 Learning Path

### 🎯 Recommended Progression
1. `01_basic/` - Master fundamental chat patterns
2. `02_group/` - Understand multi-participant messaging
3. `03_rich_media/` - Learn file and media handling
4. `04_advanced/` - Implement production features

## 🏗️ Architecture Overview

```
┌─────────────┐    ┌─────────────┐    ┌─────────────┐    ┌─────────────┐
│  01_basic   │───▶│  02_group   │───▶│03_rich_media│───▶│04_advanced  │
└─────────────┘    └─────────────┘    └─────────────┘    └─────────────┘
       │                   │                   │                   │
       ▼                   ▼                   ▼                   ▼
   Core Chat         Group Messaging    File Attachments   Production
   Patterns          & Management       & Media Sharing     Features
```

## 🎯 What You'll Learn

### Core Concepts
- **Event-Driven Messaging**: Handle real-time message events
- **Chat Message Types**: Text, media, and custom content
- **Response Patterns**: Automated and intelligent message responses
- **End-to-End Encryption**: Automatic secure message delivery

### Advanced Features
- **Group Management**: Create and manage chat groups
- **Rich Content**: Share files, images, and multimedia
- **Presence Indicators**: Typing, online status, read receipts
- **Message History**: Store and retrieve conversation history

## 🔧 Prerequisites

1. **Go**: Version 1.22 or later
2. **Self SDK**: Core SDK for messaging functionality
3. **Network**: Internet access for Self network connectivity
4. **Storage**: Write permissions for message and account storage

## 📖 Feature Comparison

| Feature | 01_basic | 02_group | 03_rich_media | 04_advanced |
|---------|----------|----------|---------------|-------------|
| Direct Messaging | ✅ | ✅ | ✅ | ✅ |
| Auto-Response | ✅ | ✅ | ✅ | ✅ |
| Group Chat | ❌ | ✅ | ✅ | ✅ |
| File Sharing | ❌ | ❌ | ✅ | ✅ |
| Typing Indicators | ❌ | ❌ | ❌ | ✅ |
| Read Receipts | ❌ | ❌ | ❌ | ✅ |
| Message History | ❌ | ❌ | ❌ | ✅ |

## 🎓 Educational Philosophy

**Progressive Complexity**: Start with core messaging patterns and gradually introduce advanced features, making it easier to understand each concept thoroughly.

## 🔗 Related Examples

- **Setup**: `../00_setup/` - Account creation and configuration
- **Connections**: `../01_connection/` - Establishing peer connections
- **Credentials**: `../02_credentials/` - Adding verification to conversations
- **Advanced Features**: `../04_advanced_features/` - Production-ready patterns

## 🚀 Quick Integration

Each example is designed to be:
- **Self-Contained**: Complete functionality in each directory
- **Production-Ready**: Code patterns suitable for real applications
- **Well-Documented**: Extensive comments and explanations
- **Customizable**: Easy to modify for specific use cases
