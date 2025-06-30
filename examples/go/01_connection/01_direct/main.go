// Package main demonstrates how to establish direct connections using inbox addresses.
//
// 🎯 What you'll learn:
// - How to create a Self account that accepts all incoming connections
// - How to display an inbox address for direct connections
// - How direct address-based connections work (without QR codes)
// - How to handle incoming connection requests automatically
//
// This example shows the SERVER SIDE of a direct connection. Other parties can
// connect directly using the displayed inbox address.
//
// 🔑 KEY CONCEPTS:
// 1. Direct Address Connection - Use inbox addresses instead of QR codes
// 2. Automatic Connection Acceptance - Accept all incoming connection requests
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joinself/academy/examples/go/common"
	"github.com/joinself/self-go-sdk/account"
	"github.com/joinself/self-go-sdk/event"
)

func main() {
	fmt.Println("🔗 Direct Connection Example - Server Side")
	fmt.Println("==========================================")

	// Setup: Create Self account with connection handling
	fmt.Println("Setting up Self account...")

	selfAccount := common.SetupAccount(common.AccountConfig{
		Callbacks: account.Callbacks{
			// Automatically accept all incoming connections via key package
			OnKeyPackage: handleKeyPackageCallback,
			// Required but unused for this example
			OnMessage: func(selfAccount *account.Account, msg *event.Message) {},
		},
	})

	fmt.Println("✅ Account ready!")
	common.DisplayAccountInfo(selfAccount, "Server Account")
	defer selfAccount.Close()

	// CONCEPT 1: Display Inbox Address for Direct Connection
	fmt.Println("\n🔑 CONCEPT 1: Creating Direct Connection Address")
	fmt.Println("===============================================")
	if !displayConnectionAddress(selfAccount) {
		fmt.Println("❌ Failed to create connection address. Please try again.")
		return
	}

	// CONCEPT 2: Wait for and Accept All Incoming Connections
	fmt.Println("\n🔑 CONCEPT 2: Accepting All Incoming Connections")
	fmt.Println("==============================================")
	fmt.Println("📧 Share the address above with other parties for direct connection")
	fmt.Println("⏳ Waiting for connections... (Press Ctrl+C to exit)")
	fmt.Println("🤖 All incoming connections will be accepted automatically")

	// Keep running to handle connections
	select {}
}

// ============================================================================
// 🔑 KEY CONCEPT 1: Direct Address Connection
// ============================================================================
//
// displayConnectionAddress demonstrates how to create and display an inbox
// address that other parties can use to connect directly (without QR codes).
//
// Key steps:
// 1. Open an inbox (temporary address for receiving connections)
// 2. Extract the inbox address from the response
// 3. Display the inbox address for direct sharing
//
// Returns true on success, false if inbox creation fails.
func displayConnectionAddress(selfAccount *account.Account) bool {
	// Step 1: Open inbox for receiving connection requests
	// This creates a temporary "mailbox" where others can send connection requests
	inboxAddress, err := selfAccount.InboxOpen()
	if err != nil {
		log.Printf("❌ Failed to open inbox: %v", err)
		log.Printf("💡 This might happen if network connectivity is poor or account setup failed")
		return false
	}

	// Step 2: Extract the inbox address (returned directly from InboxOpen)
	// The address is a Self DID that others can use to connect directly

	// Step 3: Display the inbox address for direct connection
	fmt.Println("\n📧 DIRECT CONNECTION ADDRESS:")
	fmt.Println("=============================")
	fmt.Printf("Address: %s\n", inboxAddress)
	fmt.Println("\n💡 Other parties can connect using this address directly")
	fmt.Println("   (no QR code scanning required)")
	fmt.Println("📋 How others connect:")
	fmt.Println("   1. Copy the address above")
	fmt.Println("   2. Use it in their Self SDK connection method")
	fmt.Println("   3. Send a connection request to this address")

	return true
}

// ============================================================================
// 🔑 KEY CONCEPT 2: Automatic Connection Acceptance
// ============================================================================
//
// handleKeyPackageCallback handles incoming connection requests automatically.
// This function is called when another party sends a connection request to our
// inbox address.
//
// What happens:
// 1. Receives a KeyPackage from the connecting party
// 2. Establishes an encrypted connection using the key package
// 3. Creates a secure communication group
// 4. Exits the program (for demo purposes)
//
// In production, you might want to:
// - Validate the connection request before accepting
// - Store connection information for later use
// - Continue running to handle multiple connections
func handleKeyPackageCallback(selfAccount *account.Account, kpg *event.KeyPackage) {
	fmt.Printf("\n🎉 Connection request received from: %s\n", kpg.FromAddress().String())

	// Establish a new encrypted group with the connecting party using their key package
	// This creates the secure communication channel between the two parties
	groupAddress, err := selfAccount.ConnectionEstablish(
		kpg.ToAddress(),  // Our address (where they sent the request)
		kpg.KeyPackage(), // Their cryptographic key package for encryption
	)

	if err != nil {
		log.Printf("❌ Failed to establish connection: %v", err)
		log.Println("💡 This might happen if the key package is invalid or network issues occur")
		return
	}

	fmt.Printf("✅ Successfully established encrypted connection!\n")
	fmt.Printf("📱 Connected to: %s\n", kpg.FromAddress().String())
	fmt.Printf("🔐 Secure group created: %s\n", groupAddress.String())
	fmt.Println("🚀 Connection is now ready for secure messaging!")

	// Exit the program for this demo (in production, you'd continue running)
	fmt.Println("\n🏁 Demo completed - connection established successfully!")
	os.Exit(0)
}

// ============================================================================
// 💡 APPLYING THESE CONCEPTS IN YOUR APPLICATIONS
// ============================================================================
//
// The direct address connection approach demonstrated above is useful for:
//
// 🔑 Direct Address Connection (displayConnectionAddress):
// - Server-to-server connections where QR codes aren't practical
// - API endpoints that need to establish Self connections
// - Automated systems that share connection addresses programmatically
// - Integration with existing address-based communication systems
//
// 🔑 Automatic Connection Acceptance (handleKeyPackageCallback):
// - Development and testing environments
// - Trusted network scenarios
// - Automated connection workflows
// - Systems that need to accept connections without user intervention
//
// 🚀 Next Steps:
// - Add connection validation and filtering logic
// - Implement selective connection acceptance based on criteria
// - Build address sharing mechanisms (email, API, etc.)
// - Add connection management and tracking features
