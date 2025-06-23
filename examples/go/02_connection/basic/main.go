// Package main demonstrates basic connection concepts with Self SDK.
//
// This example shows how to set up a Self SDK account that can receive
// connections from third-party applications (like mobile apps) via QR codes.
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/joinself/self-go-sdk/account"
	"github.com/joinself/self-go-sdk/event"
	"github.com/joinself/self-go-sdk/message"
)

func main() {
	fmt.Println("🔗 Basic Connection Example")
	fmt.Println("===========================")

	// Create and configure Self account
	selfAccount := setupAccount()
	defer selfAccount.Close()

	// Display connection info and generate QR code
	showConnectionInfo(selfAccount)
	generateConnectionQR(selfAccount)

	fmt.Println("✅ Account ready! Scan the QR code above with a Self mobile app.")
	fmt.Println("Press Ctrl+C to exit.")

	// Keep running to handle connections
	select {}
}

func setupAccount() *account.Account {
	fmt.Println("🔧 Setting up Self account...")

	cfg := &account.Config{
		StorageKey:  make([]byte, 32),
		StoragePath: "./basic_connection_storage",
		Environment: account.TargetSandbox,
		LogLevel:    account.LogWarn,
		Callbacks: account.Callbacks{
			OnConnect: func(acc *account.Account) {
				fmt.Println("🔗 Connected to Self network")
			},
		},
	}

	selfAccount, err := account.New(cfg)
	if err != nil {
		log.Fatal("Failed to create Self account:", err)
	}

	fmt.Println("✅ Account created successfully")
	return selfAccount
}

func showConnectionInfo(selfAccount *account.Account) {
	inboxAddress, err := selfAccount.InboxOpen()
	if err != nil {
		log.Fatal("Failed to open inbox:", err)
	}

	fmt.Printf("\n📬 Account Address: %s\n", inboxAddress.String())
}

func generateConnectionQR(selfAccount *account.Account) {
	fmt.Println("\n📱 Generating QR code...")

	inboxAddress, err := selfAccount.InboxOpen()
	if err != nil {
		log.Fatal("Failed to open inbox:", err)
	}

	// Generate key package for connection
	keyPackage, err := selfAccount.ConnectionNegotiateOutOfBand(
		inboxAddress,
		time.Now().Add(30*time.Minute),
	)
	if err != nil {
		log.Printf("Failed to generate key package: %v", err)
		fmt.Println("💡 This may happen in demo environments")
		return
	}

	// Build discovery request
	content, err := message.NewDiscoveryRequest().
		KeyPackage(keyPackage).
		Expires(time.Now().Add(30 * time.Minute)).
		Finish()
	if err != nil {
		log.Printf("Failed to build discovery request: %v", err)
		return
	}

	// Create and encode QR code
	anonymousMsg := event.NewAnonymousMessage(content)
	anonymousMsg.SetFlags(event.MessageFlagTargetSandbox)

	qrCode, err := anonymousMsg.EncodeToQR(event.QREncodingUnicode)
	if err != nil {
		log.Printf("Failed to generate QR code: %v", err)
		return
	}

	// Display QR code
	fmt.Println("\n📱 QR CODE:")
	fmt.Println("==========")
	fmt.Println(string(qrCode))
	fmt.Println("==========")
	fmt.Printf("Valid for: 30 minutes\n\n")
}
