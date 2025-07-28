# Quick Start

Get up and running with Self in minutes. This guide will help you quickly set up and test Self's core features with working examples.

## Try Examples Immediately

Experience Self in action with our demo applications. Each app demonstrates real-world workflows with Self SDK.

### Quick Install Options

For instant testing without building from source, use our pre-compiled binaries and applications:

| Platform | Quick Install | Source Code |
|----------|---------------|-------------|
| **Mobile - Android** | [Download from Play Store](https://play.google.com/store/apps/details?id=com.joinself.app.demo) | [View Source](https://github.com/joinself/demo-android) |
| **Mobile - iOS** | [Download from App Store](PLACEHOLDER_IOS_APP_STORE_LINK) | [View Source](https://github.com/joinself/demo-ios) |
| **Server - Golang** | `docker run -it ghcr.io/joinself/self-sdk-demo:go` | [View Source](https://github.com/joinself/academy/blob/main/examples/server/) |
| **Server - Java** | `docker run -it ghcr.io/joinself/self-sdk-demo:java` | [View Source](https://github.com/joinself/academy/blob/main/examples/server/) |

### Quick Start Steps

1. **Start a server** - Use the Docker links above to run a Golang or Java server
2. **Install mobile app** - Download the Android or iOS app using the links above
3. **Register Self account** - Open the mobile app and follow the registration prompts
4. **Connect to server** - Enter the server identifier into the mobile app to establish connection
5. **Test features** - Follow the prompts in the mobile app to test authentication, verification, and signing

Each application includes complete workflows with real cryptographic operations, verification, and audit trails.

## Build from Source

To customize and build the examples yourself:

### Prerequisites

- **Git**: For cloning the repository
- **Docker**: For running server examples (optional)
- **Platform-specific tools**:
    - **Android**: Android Studio, JDK 11+
    - **iOS**: Xcode 14+, macOS
    - **Go**: Go 1.19+
    - **Java**: JDK 11+, Maven or Gradle

### Clone Repository

```bash
# Clone with all submodules
git clone --recurse-submodules https://github.com/joinself/academy.git

# Or if already cloned, initialize submodules
git submodule update --init --recursive
```

### Repository Structure

Each platform's examples are organized in dedicated folders within the repository:

| Platform | Example Location | Type | Description |
|----------|-----------------|------|-------------|
| **Android** | [`demo-android`](https://github.com/joinself/demo-android) | Mobile Peer | Complete Android apps demonstrating all SDK features |
| **iOS** | [`demo-ios`](https://github.com/joinself/demo-ios) | Mobile Peer | Native iOS implementations with SwiftUI |
| **Golang** | [`examples/server/`](https://github.com/joinself/academy/blob/main/examples/server/) | Server Peer | Backend Go implementations with full workflows |
| **Java** | [`examples/server/`](https://github.com/joinself/academy/blob/main/examples/server/) | Server Peer | Enterprise Java server examples |

This structure provides both **server peers** (backend) and **mobile peers** (frontend) that can interact with each other, giving you a complete Self network for testing and development.

## Platform-Specific Quick Start

### Android Quick Start

1. **Open in Android Studio**:
   ```bash
   cd examples/mobile/android
   # Open the project in Android Studio
   ```

2. **Configure SDK**:
   - Add your Self SDK credentials to `local.properties`
   - Sync project with Gradle files

3. **Run the App**:
   - Connect an Android device or start an emulator
   - Click "Run" in Android Studio
   - Follow the in-app instructions

### iOS Quick Start

1. **Open in Xcode**:
   ```bash
   cd examples/mobile/ios
   # Open Example.xcworkspace in Xcode
   ```

2. **Configure SDK**:
   - Add your Self SDK credentials to the project
   - Update bundle identifier if needed

3. **Run the App**:
   - Connect an iOS device or start a simulator
   - Click "Run" in Xcode
   - Follow the in-app instructions

### Go Server Quick Start

1. **Navigate to Go examples**:
   ```bash
   cd examples/server
   ```

2. **Install dependencies**:
   ```bash
   go mod download
   ```

3. **Run the server**:
   ```bash
   go run main.go
   ```

4. **Connect with mobile app**:
   - Use the server identifier displayed in the console
   - Enter it in the mobile app to establish connection

### Java Server Quick Start

1. **Navigate to Java examples**:
   ```bash
   cd examples/server
   ```

2. **Build the project**:
   ```bash
   # Using Maven
   mvn clean install
   
   # Or using Gradle
   ./gradlew build
   ```

3. **Run the server**:
   ```bash
   # Using Maven
   mvn exec:java -Dexec.mainClass="com.joinself.example.Main"
   
   # Or using Gradle
   ./gradlew run
   ```

4. **Connect with mobile app**:
   - Use the server identifier displayed in the console
   - Enter it in the mobile app to establish connection

## Testing Features

Once you have both server and mobile apps running, you can test:

### Authentication
- Request liveness verification from the server
- Complete biometric check on mobile device
- Verify credentials on server side

### Identity Verification
- Capture identity documents on mobile
- Verify document authenticity
- Store verified credentials locally

### Digital Signatures
- Prepare documents on server
- Sign documents on mobile device
- Verify signatures on server

## Troubleshooting

### Common Issues

**Connection Problems**:
- Ensure both server and mobile apps are running
- Check that server identifier is entered correctly
- Verify network connectivity

**Build Issues**:
- Check that all dependencies are installed
- Ensure you're using the correct SDK version
- Verify platform-specific requirements

**Runtime Errors**:
- Check console logs for error messages
- Verify SDK credentials are configured correctly
- Ensure device/emulator meets minimum requirements

### Getting Help

- **[Troubleshooting Guide](../examples/troubleshooting.md)**: Common issues and solutions
- **[Developer Community](../examples/resources/)**: Connect with other developers
- **[Contact Support](https://joinself.com/contact)**: Get help from the Self team

## Next Steps

After successfully running the examples:

1. **[Explore Use Cases](use-cases.md)**: Learn about authentication, verification, and signing
2. **[Choose Your Learning Path](learning-paths.md)**: Find the right path for your goals
3. **[Build Your First App](../examples/setup/)**: Create your own identity application
4. **[Deploy to Production](../examples/production.md)**: Learn about production deployment

Ready to start building? Choose your next step and begin your journey with Self! 
