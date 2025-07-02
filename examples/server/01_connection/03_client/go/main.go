// Package main demonstrates how to initiate connections to inbox addresses.
//
// 🎯 What you'll learn:
// - How to connect TO a known inbox address (client side)
// - How to use ConnectionNegotiate to initiate connections
// - How to handle connection responses from the server
//
// This example shows the CLIENT SIDE of a direct connection. It connects TO
// an inbox address created by the direct server example.
//
// 🔑 KEY CONCEPTS:
// 1. Connection Initiation - Connect to a known inbox address
// 2. Connection Response Handling - Handle server acceptance/rejection
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joinself/academy/examples/go/common"
	"github.com/joinself/self-go-sdk/account"
	"github.com/joinself/self-go-sdk/event"
	"github.com/joinself/self-go-sdk/keypair/signing"
)

func main() {
	fmt.Println("🔗 Connection Client Example")
	fmt.Println("============================")

	// Check if inbox address was provided
	if len(os.Args) < 2 {
		fmt.Println("❌ Usage: go run main.go <inbox-address>")
		fmt.Println("")
		fmt.Println("💡 You can get an inbox address from:")
		fmt.Println("   🖥️  Direct server example: cd ../01_direct && go run main.go")
		fmt.Println("   📱 Self mobile app: Generate address in app settings")
		fmt.Println("   🔧 Another Self SDK program: Any app using InboxOpen()")
		fmt.Println("   🌐 Self-enabled service: API endpoint or web service")
		fmt.Println("")
		fmt.Println("📚 This client connects to ANY Self inbox address!")
		fmt.Println("   Address source → This client connects → Secure channel established")
		return
	}

	inboxAddress := os.Args[1]
	fmt.Printf("🎯 Target inbox address: %s\n", inboxAddress)

	// Setup: Create Self account for client connection
	fmt.Println("Setting up client account...")

	clientAccount := common.SetupAccount(common.AccountConfig{
		Callbacks: account.Callbacks{
			// Handle connection response from server
			OnWelcome: handleConnectionResponse,
			// Required but unused for this example
			OnMessage: func(selfAccount *account.Account, msg *event.Message) {},
		},
	})

	fmt.Println("✅ Client account ready!")
	common.DisplayAccountInfo(clientAccount, "Client Account")
	defer clientAccount.Close()

	// CONCEPT 1: Initiate connection to inbox address
	fmt.Println("\n🔑 CONCEPT 1: Initiating Connection")
	fmt.Println("==================================")
	if !initiateConnection(clientAccount, inboxAddress) {
		fmt.Println("❌ Failed to initiate connection. Please try again.")
		return
	}

	// CONCEPT 2: Wait for server response
	fmt.Println("\n🔑 CONCEPT 2: Waiting for Server Response")
	fmt.Println("=======================================")
	fmt.Println("📤 Connection request sent to server")
	fmt.Println("⏳ Waiting for server to accept connection... (Press Ctrl+C to exit)")
	fmt.Println("💡 The server will either accept or reject this connection request")

	// Keep running to handle server response
	select {}
}

// ============================================================================
// 🔑 KEY CONCEPT 1: Connection Initiation
// ============================================================================
//
// initiateConnection demonstrates how to connect TO a known inbox address
// using ConnectionNegotiate to send a connection request to the server.
func initiateConnection(clientAccount *account.Account, inboxAddress string) bool {
	fmt.Printf("📞 Connecting to inbox: %s\n", inboxAddress)

	// Parse the inbox address to get the recipient's public key
	recipientKey := signing.FromAddress(inboxAddress)

	// Get our own inbox address and convert to public key for sender
	clientInboxAddress, err := clientAccount.InboxOpen()
	if err != nil {
		log.Printf("❌ Failed to get client inbox address: %v", err)
		return false
	}

	senderKey := signing.FromAddress(clientInboxAddress.String())

	// Set connection expiration to 5 minutes from now
	expirationTime := time.Now().Add(5 * time.Minute)

	// Use ConnectionNegotiate to initiate connection to the inbox address
	err = clientAccount.ConnectionNegotiate(senderKey, recipientKey, expirationTime)
	if err != nil {
		log.Printf("❌ Failed to initiate connection: %v", err)
		fmt.Println("💡 Make sure the inbox address is valid and the server is running")
		return false
	}

	fmt.Println("✅ Connection request sent successfully!")
	fmt.Println("📡 Waiting for server to process the request...")
	return true
}

// ============================================================================
// 🔑 KEY CONCEPT 2: Connection Response Handling
// ============================================================================
//
// handleConnectionResponse is called when the server responds to our
// connection request (either accepting or potentially rejecting it).
func handleConnectionResponse(clientAccount *account.Account, welcome *event.Welcome) {
	serverAddress := welcome.FromAddress().String()
	fmt.Printf("\n🎉 Connection response received from server: %s\n", serverAddress)

	// Accept the connection establishment from the server
	groupAddress, err := clientAccount.ConnectionAccept(welcome.ToAddress(), welcome.Welcome())
	if err != nil {
		fmt.Printf("❌ Failed to establish connection: %v\n", err)
		return
	}

	// Connection successful!
	fmt.Printf("✅ Connection established successfully!\n")
	fmt.Printf("🔐 Connected to server: %s\n", serverAddress)
	fmt.Printf("📱 Secure group created: %s\n", groupAddress.String())
	fmt.Println("🚀 Ready for secure communication!")

	// Exit after successful connection for this demo
	fmt.Println("\n🏁 Demo completed - connection established successfully!")
	os.Exit(0)
}

// ============================================================================
// 💡 APPLYING THESE CONCEPTS IN YOUR APPLICATIONS
// ============================================================================
//
// The client-side connection initiation demonstrated above is useful for:
//
// 🔑 Connection Initiation (initiateConnection):
// - Client applications that connect to known servers
// - Automated systems that establish connections to services
// - Mobile apps connecting to backend APIs
// - Microservices connecting to other services
//
// 🔑 Connection Response Handling (handleConnectionResponse):
// - Processing server acceptance of connection requests
// - Handling connection establishment from the client perspective
// - Building reliable client-server communication patterns
//
// 🚀 Next Steps:
// - Add retry logic for failed connections
// - Implement connection status monitoring
// - Add support for multiple server connections
// - Build message exchange after connection establishment
