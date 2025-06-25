// Package main demonstrates how to establish secure connections with Self SDK.
//
// 🎯 What you'll learn:
// - How to create a Self account that can accept connections
// - How to generate QR codes for mobile apps to scan
// - How the connection handshake works between server and mobile
// - How to handle incoming connection requests
//
// This example shows the SERVER SIDE of a connection. A mobile app will scan
// the QR code generated here to establish an encrypted connection.
//
// 🔑 KEY CONCEPTS:
// 1. QR Code Generation - Create scannable codes for mobile discovery
// 2. Connection Acceptance - Handle and accept incoming connection requests
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joinself/academy/examples/go/common"
	"github.com/joinself/self-go-sdk/account"
	"github.com/joinself/self-go-sdk/event"
	"github.com/joinself/self-go-sdk/message"
)

func main() {
	fmt.Println("🔗 Connection Example - Server Side")
	fmt.Println("====================================")

	// Setup: Create Self account with connection handling
	fmt.Println("Setting up Self account...")

	selfAccount := common.SetupAccount(common.AccountConfig{
		Callbacks: account.Callbacks{
			// This is where connection acceptance happens
			OnWelcome: handleIncomingConnection,
			// Required but unused for this example
			OnMessage: func(selfAccount *account.Account, msg *event.Message) {},
		},
	})

	fmt.Println("✅ Account ready!")
	common.DisplayAccountInfo(selfAccount, "Server Account")
	defer selfAccount.Close()

	// CONCEPT 1: Generate QR Code for Mobile Discovery
	fmt.Println("\n🔑 CONCEPT 1: Generating Connection QR Code")
	fmt.Println("============================================")
	if !generateConnectionQR(selfAccount) {
		fmt.Println("❌ Failed to generate QR code. Please try again.")
		return
	}

	// CONCEPT 2: Wait for and Accept Incoming Connections
	fmt.Println("\n🔑 CONCEPT 2: Accepting Incoming Connections")
	fmt.Println("==============================================")
	fmt.Println("📱 Scan the QR code above with your Self mobile app")
	fmt.Println("⏳ Waiting for connection... (Press Ctrl+C to exit)")

	// Keep running to handle connections
	select {}
}

// ============================================================================
// 🔑 KEY CONCEPT 1: QR Code Generation for Mobile Discovery
// ============================================================================
//
// generateConnectionQR demonstrates how to create scannable QR codes that
// mobile apps can use to initiate secure connections.
//
// Key steps:
// 1. Open an inbox (temporary address for receiving connections)
// 2. Generate a key package (cryptographic material for secure communication)
// 3. Create a discovery request (message mobile apps understand)
// 4. Encode everything into a QR code
func generateConnectionQR(selfAccount *account.Account) bool {
	// Step 1: Open inbox for receiving connection requests
	inboxAddress, err := selfAccount.InboxOpen()
	if err != nil {
		log.Printf("❌ Failed to open inbox: %v", err)
		return false
	}

	// Step 2: Generate cryptographic key package for secure communication
	keyPackage, err := selfAccount.ConnectionNegotiateOutOfBand(
		inboxAddress,
		time.Now().Add(30*time.Minute), // Valid for 30 minutes
	)
	if err != nil {
		log.Printf("❌ Failed to generate key package: %v", err)
		fmt.Println("💡 This may happen in demo environments or with network issues")
		return false
	}

	// Step 3: Build discovery request message
	content, err := message.NewDiscoveryRequest().
		KeyPackage(keyPackage).
		Expires(time.Now().Add(30 * time.Minute)).
		Finish()
	if err != nil {
		log.Printf("❌ Failed to build discovery request: %v", err)
		return false
	}

	// Step 4: Create anonymous message and encode to QR
	anonymousMsg := event.NewAnonymousMessage(content)
	anonymousMsg.SetFlags(event.MessageFlagTargetSandbox)

	qrCode, err := anonymousMsg.EncodeToQR(event.QREncodingUnicode)
	if err != nil {
		log.Printf("❌ Failed to generate QR code: %v", err)
		return false
	}

	// Display the scannable QR code
	fmt.Println("\n" + string(qrCode))
	fmt.Printf("⏱️  Expires: %s\n", time.Now().Add(30*time.Minute).Format("15:04:05"))

	return true
}

// ============================================================================
// 🔑 KEY CONCEPT 2: Accepting Incoming Connections
// ============================================================================
//
// handleIncomingConnection demonstrates how to accept connection requests
// from mobile apps that have scanned your QR code.
//
// Key steps:
// 1. Receive a Welcome message (connection request from mobile app)
// 2. Accept the connection using ConnectionAccept
// 3. Establish encrypted communication channel
func handleIncomingConnection(acc *account.Account, welcome *event.Welcome) {
	// Step 1: Connection request received from mobile app
	fmt.Printf("\n🎉 Connection received from: %s\n", welcome.FromAddress().String())

	// Step 2: Accept the connection request
	// This is the critical call that establishes the secure connection
	_, err := acc.ConnectionAccept(welcome.ToAddress(), welcome.Welcome())
	if err != nil {
		fmt.Printf("❌ Failed to accept connection: %v\n", err)
		return
	}

	// Step 3: Connection established - ready for secure communication
	fmt.Println("✅ Connection established successfully!")
	fmt.Println("🚀 Ready to exchange messages and credentials")

	// Exit gracefully after successful connection
	os.Exit(0)
}

// ============================================================================
// 💡 APPLYING THESE CONCEPTS IN YOUR APPLICATIONS
// ============================================================================
//
// The two key concepts demonstrated above can be applied in many scenarios:
//
// 🔑 QR Code Generation (generateConnectionQR):
// - Desktop apps that mobile users connect to
// - IoT devices that need mobile configuration
// - Web services with mobile authentication
// - Developer tools with mobile companion apps
//
// 🔑 Connection Acceptance (handleIncomingConnection):
// - Automatically accepting all connections (like this example)
// - Manual approval with user confirmation
// - Role-based connection acceptance
// - Rate limiting or connection management
//
// 🚀 Next Steps:
// - Customize the OnWelcome callback for your needs
// - Add message handling with OnMessage callback
// - Implement credential exchange workflows
// - Build real-time communication features
