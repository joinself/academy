# Self SDK Mobile Examples

Welcome to the Self SDK mobile examples! This directory contains examples for Android and iOS platforms that demonstrate mobile-specific patterns and UI integrations.

## 📱 Mobile Examples Overview

| Category | Example | Summary | Platform |
|----------|---------|---------|----------|
| **Setup** | [New Account](00_setup/01_new_account/) | Create Self identity with mobile UI | Android, iOS |
| **Setup** | [Existing Account](00_setup/02_existing_account/) | Load saved accounts with biometric auth | Android, iOS |
| **Connection** | [QR Scanner](01_connection/02_qr/) | Scan QR codes to connect | Android, iOS |
| **Connection** | [Client](01_connection/03_client/) | Connect to server addresses | Android, iOS |
| **Chat** | [Basic Chat](03_chat/01_basic/) | Secure messaging with mobile UI | Android, iOS |
| **Credentials** | [Email Verification](02_credentials/email_verification/) | Mobile email verification flow | Android, iOS |
| **Credentials** | [Presentation](02_credentials/presentation_request/) | Share credentials via mobile UI | Android, iOS |
| **Advanced** | [Core Features](04_advanced_features/01_core_features/) | Mobile SDK integration patterns | Android, iOS |
| **Advanced** | [Notifications](04_advanced_features/03_notifications/) | Mobile push notifications | Android, iOS |

## 🚀 Quick Start

### Android
```bash
cd <category>/<example>/android && ./gradlew run
```

### iOS
```bash
cd <category>/<example>/ios && xcodebuild -project SelfSDKExample.xcodeproj
```

## 📱 Mobile-Specific Features

- **UI Components**: Pre-built mobile UI components
- **Biometric Auth**: Fingerprint/Face ID integration
- **QR Scanning**: Camera-based connection discovery
- **Push Notifications**: Real-time mobile alerts
- **Touch Interfaces**: Mobile-optimized user experience

## 🎯 Learning Path

1. **Start with Setup** → Create and load accounts
2. **Learn Connections** → QR scanning and client connections
3. **Try Messaging** → Basic chat functionality
4. **Explore Credentials** → Mobile verification flows
5. **Advanced Features** → Production patterns

---

**Ready to build mobile apps with Self SDK?** Pick your platform and start! 🚀 
