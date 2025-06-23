// Package main demonstrates simple discovery and connection using the core Self SDK.
//
// This example shows how to generate a discovery QR code, wait for a peer
// to connect, send them a message, and complete the demo. It focuses on
// the essential discovery workflow without complexity.
//
// 🎯 Key concepts: QR generation, connection handling, message sending
// 📚 For detailed explanations, see the README.md file
package main

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joinself/self-go-sdk/account"
	"github.com/joinself/self-go-sdk/event"
	"github.com/joinself/self-go-sdk/keypair/signing"
	"github.com/joinself/self-go-sdk/message"
)

// SimpleDiscoveryDemo manages the simplified discovery demonstration
type SimpleDiscoveryDemo struct {
	account   *account.Account
	connected bool
	peerDID   string
	ctx       context.Context
	cancel    context.CancelFunc
	done      chan bool
}

func main() {
	fmt.Println("🔍 Simple Discovery Demo")
	fmt.Println("=========================")
	fmt.Println("This demo shows basic discovery workflow:")
	fmt.Println("• Generate one QR code for connection")
	fmt.Println("• Wait for a peer to scan and connect")
	fmt.Println("• Send a welcome message")
	fmt.Println("• Complete the demo")
	fmt.Println()

	// Create and run the demo
	demo := NewSimpleDiscoveryDemo()
	defer demo.Close()

	demo.Run()
}

// NewSimpleDiscoveryDemo creates a new simplified discovery demonstration
func NewSimpleDiscoveryDemo() *SimpleDiscoveryDemo {
	ctx, cancel := context.WithCancel(context.Background())

	return &SimpleDiscoveryDemo{
		ctx:    ctx,
		cancel: cancel,
		done:   make(chan bool, 1),
	}
}

// Close cleans up the demo resources
func (d *SimpleDiscoveryDemo) Close() {
	d.cancel()
	if d.account != nil {
		d.account.Close()
	}
}

// Run executes the simplified discovery demonstration
func (d *SimpleDiscoveryDemo) Run() {
	// Setup graceful shutdown
	d.setupGracefulShutdown()

	// Step 1: Create account with event handlers
	d.createAccount()

	// Step 2: Generate and display QR code
	d.generateQRCode()

	// Step 3: Wait for connection and complete demo
	d.waitForConnection()
}

// setupGracefulShutdown handles Ctrl+C gracefully
func (d *SimpleDiscoveryDemo) setupGracefulShutdown() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		fmt.Println("\n🛑 Demo interrupted")
		d.cancel()
	}()
}

// generateStorageKey creates a cryptographically secure 32-byte key
func generateStorageKey(seed string) []byte {
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		h := sha256.Sum256([]byte(fmt.Sprintf("self-sdk-%s-%d", seed, time.Now().UnixNano())))
		return h[:]
	}
	return key
}

// createAccount sets up the account with event handlers
func (d *SimpleDiscoveryDemo) createAccount() {
	fmt.Println("🔧 Setting up discovery account...")

	cfg := &account.Config{
		StorageKey:  generateStorageKey("simple_discovery"),
		StoragePath: "./simple_discovery_storage",
		Environment: account.TargetSandbox,
		LogLevel:    account.LogWarn,
		Callbacks: account.Callbacks{
			OnConnect: func(acc *account.Account) {
				fmt.Println("🔗 Connected to Self network")
			},
			OnWelcome: d.onWelcome,
		},
	}

	var err error
	d.account, err = account.New(cfg)
	if err != nil {
		log.Fatal("Failed to create account:", err)
	}

	inboxAddress, err := d.account.InboxOpen()
	if err != nil {
		log.Fatal("Failed to open inbox:", err)
	}

	fmt.Println("✅ Account created successfully")
	fmt.Printf("🆔 Your DID: %s\n", inboxAddress.String())
	fmt.Println()
}

// onWelcome handles incoming connection requests
func (d *SimpleDiscoveryDemo) onWelcome(acc *account.Account, wlc *event.Welcome) {
	fmt.Printf("🤝 Connection request from: %s\n", wlc.FromAddress().String())

	// Accept the connection
	_, err := acc.ConnectionAccept(wlc.ToAddress(), wlc.Welcome())
	if err != nil {
		fmt.Printf("❌ Failed to accept connection: %v\n", err)
		return
	}

	d.peerDID = wlc.FromAddress().String()
	d.connected = true

	fmt.Printf("🎉 Connected to peer: %s\n", d.peerDID)

	// Send welcome message and complete demo
	d.sendWelcomeMessage(wlc.FromAddress())

	// Signal completion
	d.done <- true
}

// generateQRCode creates and displays a discovery QR code
func (d *SimpleDiscoveryDemo) generateQRCode() {
	fmt.Println("📱 Generating discovery QR code...")

	inboxAddress, err := d.account.InboxOpen()
	if err != nil {
		log.Fatal("Failed to open inbox:", err)
	}

	// Generate key package for connection (valid for 30 minutes)
	keyPackage, err := d.account.ConnectionNegotiateOutOfBand(
		inboxAddress,
		time.Now().Add(30*time.Minute),
	)
	if err != nil {
		log.Fatal("Failed to generate key package:", err)
	}

	// Build discovery request
	discoveryContent, err := message.NewDiscoveryRequest().
		KeyPackage(keyPackage).
		Expires(time.Now().Add(30 * time.Minute)).
		Finish()
	if err != nil {
		log.Fatal("Failed to build discovery request:", err)
	}

	// Create and encode QR code
	anonymousMsg := event.NewAnonymousMessage(discoveryContent)
	anonymousMsg.SetFlags(event.MessageFlagTargetSandbox)

	qrCodeData, err := anonymousMsg.EncodeToQR(event.QREncodingUnicode)
	if err != nil {
		log.Fatal("Failed to encode QR code:", err)
	}

	// Display QR code
	fmt.Println("--- Discovery QR Code ---")
	fmt.Println("Valid for: 30 minutes")
	fmt.Println()
	fmt.Println(string(qrCodeData))
	fmt.Println()
	fmt.Println("✅ QR code generated successfully!")
	fmt.Println()
}

// waitForConnection waits for a peer to connect
func (d *SimpleDiscoveryDemo) waitForConnection() {
	fmt.Println("⏳ Waiting for peer connection...")
	fmt.Println("📱 Scan the QR code above with a Self client to connect")
	fmt.Println("🔄 Press Ctrl+C to cancel")
	fmt.Println()

	// Wait for either connection or cancellation
	select {
	case <-d.done:
		fmt.Println("✅ Demo completed successfully!")
		d.printSummary()
	case <-d.ctx.Done():
		fmt.Println("❌ Demo cancelled")
		if !d.connected {
			fmt.Println("💡 No peer connected during this session")
		}
	case <-time.After(30 * time.Minute):
		fmt.Println("⏰ QR code expired after 30 minutes")
		fmt.Println("💡 Run the demo again to generate a new QR code")
	}
}

// sendWelcomeMessage sends a welcome message to the connected peer
func (d *SimpleDiscoveryDemo) sendWelcomeMessage(peerAddress *signing.PublicKey) {
	welcomeMsg := "🎉 Welcome! You've successfully connected to the Self SDK Discovery Demo. Connection established!"

	fmt.Printf("📤 Sending welcome message...\n")

	// Build the chat message
	chatContent, err := message.NewChat().
		Message(welcomeMsg).
		Finish()
	if err != nil {
		fmt.Printf("❌ Failed to build message: %v\n", err)
		return
	}

	// Send the message
	err = d.account.MessageSend(peerAddress, chatContent)
	if err != nil {
		fmt.Printf("❌ Failed to send message: %v\n", err)
	} else {
		fmt.Printf("✅ Welcome message sent successfully!\n")
	}
}

// printSummary shows the final demo results
func (d *SimpleDiscoveryDemo) printSummary() {
	fmt.Println()
	fmt.Println("📋 Demo Summary")
	fmt.Println("================")
	fmt.Printf("👤 Connected peer: %s\n", d.peerDID)
	fmt.Println("💬 Welcome message sent")
	fmt.Println()
	fmt.Println("🎓 What was demonstrated:")
	fmt.Println("   • QR code generation for discovery")
	fmt.Println("   • Automatic connection acceptance")
	fmt.Println("   • Direct peer messaging")
	fmt.Println("   • Clean demo completion")
	fmt.Println()
}
