# Self Authentication Minimal Demo

A plug-and-play, production-grade example of authentication using the Self SDK.  
**Scan a QR code, authenticate with your Self mobile app, and see your cryptographic identity and claims—no passwords, no sessions, no complexity.**

---

## 🎯 What You'll Learn
- How to implement an authentication workflow with Self SDK
- How QR code-based authentication works
- How to extract a user's DID and claims

---

## 🚀 Quick Start

1. **Install Go 1.24+**
2. **Clone this repo and enter the directory:**
   ```sh
   git clone <repo-url>
   cd <repo-dir>
   ```
3. **Generate a secure storage key:**
   ```sh
   openssl rand -base64 32
   ```
4. **Run the server:**
   ```sh
   SELF_AUTH_STORAGE_KEY="your-base64-key" go run cmd/server/main.go
   ```
5. **Open [http://localhost:8081](http://localhost:8081) in your browser.**
6. **Scan the QR code with your Self mobile app and authenticate!**

---

## 🔍 How It Works

- The server generates a unique QR code for each authentication request.
- You scan the QR code with your Self mobile app.
- The server verifies your identity and displays your DID and verified claims.
- No passwords, no persistent sessions, no personal data processed.
- Refresh or click "Start Over" to try again!

---

## 📝 What Just Happened
- You scanned a QR code, which contained a unique cryptographic request.
- Your Self app responded, proving your identity.
- The server displayed your DID and claims—no passwords, no session state.

---

## ✨ Features

- Passwordless authentication with Self SDK
- Cryptographically unique QR codes for each request
- User isolation and perfect request/response correlation
- Minimal, stateless backend—no session or cookie logic
- Clean, modern UI for easy demos and integration

---

## ⚙️ Configuration

- `SELF_AUTH_STORAGE_KEY` (required): 32-byte base64 key for secure storage
- `SELF_AUTH_STORAGE_PATH` (optional): Path for persistent storage (default: `./auth_service_storage`)
- `SELF_SERVER_PORT` (optional): Port to run the server (default: `8081`)

---

## 🔌 Integration

The authentication logic is fully isolated in the `internal/auth` package.  
You can easily plug it into your own Go server:

```go
import "github.com/joinself/academy/examples/server/auth-system/internal/auth"

authService, err := auth.NewAuthService(authConfig, logger)
```

---

## ❓ FAQ

- **Q: Can I use this in production?**  
  A: Yes! The core authentication logic is production-grade. Just add your own session or user management as needed.

- **Q: What if I want to add sessions or protected data?**  
  A: See the [full-featured example](link-to-advanced-example) for session management and protected endpoints.

---

## 📖 Theory Meets Practice
- Learn more about [Decentralized Identity](../../docs/concepts/decentralized-identity.md)
- See advanced [production patterns](../../docs/examples/advanced.md)

---

> 🟢 This is a beginner example. For session management, protected data, and advanced flows, see the [full-featured example](link-to-advanced-example). 
