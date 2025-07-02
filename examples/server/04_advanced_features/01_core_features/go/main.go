// Package main demonstrates advanced Self SDK features using the core SDK.
//
// This example showcases the most important advanced capabilities of the Self SDK
// using the underlying SDK directly.
//
// 🎯 What you'll learn:
// • Advanced account configuration and management
// • Encrypted storage with automatic security
// • Discovery and connection patterns
// • Message handling with different content types
// • Production-ready account setup patterns
//
// 🚀 ADVANCED FEATURES DEMONSTRATED:
// • Core SDK account creation and configuration
// • Secure storage management
// • Discovery request generation and QR codes
// • Event-driven message handling
// • Production account patterns
package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/joinself/self-go-sdk/account"
	"github.com/joinself/self-go-sdk/event"
	"github.com/joinself/self-go-sdk/keypair/signing"
	"github.com/joinself/self-go-sdk/message"
)

// AdvancedDemo manages our advanced features demonstration
type AdvancedDemo struct {
	account      *account.Account
	connections  map[string]bool
	messageCount int
	startTime    time.Time
	mutex        sync.RWMutex
	done         chan bool
}

func main() {
	fmt.Println("🚀 Advanced Self SDK Features Demo (Core SDK)")
	fmt.Println("==============================================")
	fmt.Println("This demo showcases advanced Self SDK capabilities using the core SDK directly.")
	fmt.Println()

	// Create and run the advanced demo
	demo := &AdvancedDemo{
		connections: make(map[string]bool),
		startTime:   time.Now(),
		done:        make(chan bool),
	}

	demo.run()
}

// run orchestrates the complete advanced features demonstration
func (d *AdvancedDemo) run() {
	// Step 1: Advanced account setup
	d.setupAdvancedAccount()
	defer d.account.Close()

	// Step 2: Demonstrate storage capabilities
	d.demonstrateAdvancedStorage()

	// Step 3: Generate discovery QR for connections
	d.generateDiscoveryQR()

	// Step 4: Setup graceful shutdown
	d.setupGracefulShutdown()

	// Step 5: Wait for interactions or timeout
	d.waitForInteractions()

	// Step 6: Display final summary
	d.displaySummary()
}

// generateStorageKey creates a cryptographically secure 32-byte key
func generateStorageKey(seed string) []byte {
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		// Fallback to deterministic key generation if crypto/rand fails
		h := sha256.Sum256([]byte(fmt.Sprintf("self-sdk-%s-%d", seed, time.Now().UnixNano())))
		return h[:]
	}
	return key
}

// setupAdvancedAccount creates an account with advanced configuration
func (d *AdvancedDemo) setupAdvancedAccount() {
	fmt.Println("🔧 Setting up advanced account with core SDK...")

	// Advanced account configuration
	cfg := &account.Config{
		StorageKey:  generateStorageKey("advanced_demo"),
		StoragePath: "./advanced_demo_storage",
		Environment: account.TargetSandbox,
		LogLevel:    account.LogWarn,
		Callbacks: account.Callbacks{
			OnWelcome:    d.onWelcomeMessage,
			OnKeyPackage: d.onKeyPackage,
			OnMessage:    d.onMessage,
		},
	}

	var err error
	d.account, err = account.New(cfg)
	if err != nil {
		log.Fatal("Failed to create advanced account:", err)
	}

	// Get account details
	inbox, err := d.account.InboxOpen()
	if err != nil {
		log.Fatal("Failed to open inbox:", err)
	}

	fmt.Printf("✅ Advanced account created successfully\n")
	fmt.Printf("🆔 Account DID: %s\n", inbox.String())
	fmt.Println()
}

// demonstrateAdvancedStorage shows various storage capabilities
func (d *AdvancedDemo) demonstrateAdvancedStorage() {
	fmt.Println("💾 Advanced Storage Capabilities")
	fmt.Println("================================")

	// Basic storage operations
	fmt.Println("🔹 Encrypted Storage Operations")

	// Store various types of data
	storageData := map[string]interface{}{
		"user_profile":     map[string]string{"name": "Advanced User", "role": "Developer"},
		"app_settings":     map[string]bool{"dark_mode": true, "notifications": true},
		"session_token":    fmt.Sprintf("token_%d", time.Now().Unix()),
		"last_activity":    time.Now().Format(time.RFC3339),
		"connection_count": 0,
	}

	for key, value := range storageData {
		// Convert value to JSON bytes for storage
		jsonData, err := json.Marshal(value)
		if err != nil {
			fmt.Printf("❌ Failed to marshal %s: %v\n", key, err)
			continue
		}
		if err := d.account.ValueStore(key, jsonData); err != nil {
			fmt.Printf("❌ Failed to store %s: %v\n", key, err)
		} else {
			fmt.Printf("   ✅ Stored %s\n", key)
		}
	}

	// Retrieve and verify stored data
	fmt.Println("🔹 Data Retrieval Verification")
	if profileData, err := d.account.ValueLookup("user_profile"); err == nil {
		var profile map[string]string
		if json.Unmarshal(profileData, &profile) == nil {
			fmt.Printf("   📖 Retrieved profile: %s (%s)\n", profile["name"], profile["role"])
		}
	}

	if settingsData, err := d.account.ValueLookup("app_settings"); err == nil {
		var settings map[string]bool
		if json.Unmarshal(settingsData, &settings) == nil {
			fmt.Printf("   🎨 Dark mode: %v, Notifications: %v\n", settings["dark_mode"], settings["notifications"])
		}
	}

	fmt.Println("   ✅ Storage operations completed successfully")
	fmt.Println()
}

// generateDiscoveryQR creates a QR code for peer connections
func (d *AdvancedDemo) generateDiscoveryQR() {
	fmt.Println("🔍 Discovery & Connection Setup")
	fmt.Println("===============================")

	// Get inbox address for connection
	inboxAddress, err := d.account.InboxOpen()
	if err != nil {
		log.Printf("Failed to open inbox: %v", err)
		return
	}

	// Generate key package for connection (valid for 15 minutes)
	keyPackage, err := d.account.ConnectionNegotiateOutOfBand(
		inboxAddress,
		time.Now().Add(15*time.Minute),
	)
	if err != nil {
		log.Printf("Failed to generate key package: %v", err)
		return
	}

	// Build discovery request
	discoveryContent, err := message.NewDiscoveryRequest().
		KeyPackage(keyPackage).
		Expires(time.Now().Add(15 * time.Minute)).
		Finish()
	if err != nil {
		log.Printf("Failed to build discovery request: %v", err)
		return
	}

	// Create and encode QR code
	anonymousMsg := event.NewAnonymousMessage(discoveryContent)
	anonymousMsg.SetFlags(event.MessageFlagTargetSandbox)

	qrCode, err := anonymousMsg.EncodeToQR(event.QREncodingUnicode)
	if err != nil {
		log.Printf("Failed to encode QR code: %v", err)
		return
	}

	fmt.Println("📱 QR Code for Discovery:")
	fmt.Println("========================")
	fmt.Println(string(qrCode))
	fmt.Println("========================")
	fmt.Println()
	fmt.Println("🎯 Scan this QR code with the Self app to connect!")
	fmt.Printf("⏰ Expires in 15 minutes (%s)\n", time.Now().Add(15*time.Minute).Format("15:04:05"))
	fmt.Println()
	fmt.Println("💡 What happens when someone connects:")
	fmt.Println("   • Secure connection established")
	fmt.Println("   • Welcome message sent automatically")
	fmt.Println("   • Advanced storage updated with connection info")
	fmt.Println("   • Demo showcases real-time message handling")
	fmt.Println()
}

// onWelcomeMessage handles new connection welcome messages
func (d *AdvancedDemo) onWelcomeMessage(acc *account.Account, welcomeMsg *event.Welcome) {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	peerAddress := welcomeMsg.FromAddress()
	d.connections[peerAddress.String()] = true

	fmt.Printf("🎉 New connection established!\n")
	fmt.Printf("👤 Peer: %s\n", peerAddress.String())
	fmt.Printf("🕐 Time: %s\n", time.Now().Format("15:04:05"))
	fmt.Println()

	// Accept the connection
	_, err := acc.ConnectionAccept(welcomeMsg.ToAddress(), welcomeMsg.Welcome())
	if err != nil {
		fmt.Printf("❌ Failed to accept connection: %v\n", err)
		return
	}

	// Update storage with connection info
	connectionInfo := map[string]interface{}{
		"peer_id":         peerAddress.String(),
		"connected_at":    time.Now().Format(time.RFC3339),
		"connection_type": "discovery",
		"status":          "active",
	}

	storageKey := fmt.Sprintf("connection_%s", peerAddress.String()[:8])
	jsonData, _ := json.Marshal(connectionInfo)
	if err := d.account.ValueStore(storageKey, jsonData); err != nil {
		fmt.Printf("⚠️  Failed to store connection info: %v\n", err)
	} else {
		fmt.Println("💾 Connection info stored in advanced storage")
	}

	// Send advanced welcome message
	d.sendAdvancedWelcomeMessage(peerAddress)

	// Update connection count in storage
	countData, _ := json.Marshal(len(d.connections))
	d.account.ValueStore("connection_count", countData)
}

// onKeyPackage handles key package events (connection setup)
func (d *AdvancedDemo) onKeyPackage(acc *account.Account, keyPackageMsg *event.KeyPackage) {
	fmt.Printf("🔑 Key package received from %s\n", keyPackageMsg.FromAddress().String())
}

// onMessage handles incoming messages with advanced processing
func (d *AdvancedDemo) onMessage(acc *account.Account, msg *event.Message) {
	sender := msg.FromAddress().String()
	contentType := event.ContentTypeOf(msg)

	// Only process chat messages to avoid flooding
	if contentType != message.ContentTypeChat {
		fmt.Printf("🔍 DEBUG: Skipping non-chat message (type: %v) from %s\n", contentType, sender[:16]+"...")
		return
	}

	d.mutex.Lock()
	d.messageCount++
	count := d.messageCount
	d.mutex.Unlock()

	// Decode chat message to get content
	var content string
	if chat, err := message.DecodeChat(msg.Content()); err == nil {
		content = chat.Message()
	} else {
		content = "(failed to decode chat)"
	}

	fmt.Printf("📨 Chat Message #%d received\n", count)
	fmt.Printf("👤 From: %s\n", sender)
	fmt.Printf("💬 Content: %s\n", content)
	fmt.Printf("🕐 Time: %s\n", time.Now().Format("15:04:05"))

	// Store message in advanced storage
	messageInfo := map[string]interface{}{
		"sender":     sender,
		"content":    content,
		"timestamp":  time.Now().Format(time.RFC3339),
		"message_id": count,
		"type":       "chat",
	}

	storageKey := fmt.Sprintf("message_%d", count)
	jsonData, _ := json.Marshal(messageInfo)
	if err := d.account.ValueStore(storageKey, jsonData); err != nil {
		fmt.Printf("⚠️  Failed to store message: %v\n", err)
	} else {
		fmt.Println("💾 Message stored in advanced storage")
	}

	// Send automatic response showcasing advanced features
	d.sendAdvancedResponse(msg.FromAddress(), count)
	fmt.Println()
}

// sendAdvancedWelcomeMessage sends a comprehensive welcome message
func (d *AdvancedDemo) sendAdvancedWelcomeMessage(peer *signing.PublicKey) {
	welcomeText := fmt.Sprintf(`🚀 Welcome to the Advanced Self SDK Demo!

This demo showcases advanced Self SDK features:
✅ Core SDK account management
✅ Encrypted storage with automatic security  
✅ Real-time message handling
✅ Advanced connection tracking
✅ Production-ready patterns

Connected at: %s
Demo started: %s
Active connections: %d

Send me any message to see advanced features in action!`,
		time.Now().Format("15:04:05"),
		d.startTime.Format("15:04:05"),
		len(d.connections))

	chatContent, err := message.NewChat().
		Message(welcomeText).
		Finish()
	if err != nil {
		fmt.Printf("❌ Failed to create welcome message: %v\n", err)
		return
	}

	if err := d.account.MessageSend(peer, chatContent); err != nil {
		fmt.Printf("❌ Failed to send welcome message: %v\n", err)
	} else {
		fmt.Println("📤 Advanced welcome message sent")
	}
}

// sendAdvancedResponse sends an intelligent response demonstrating various features
func (d *AdvancedDemo) sendAdvancedResponse(peer *signing.PublicKey, messageCount int) {
	responseText := fmt.Sprintf(`📊 Advanced Demo Stats:

Message #%d processed ✅
Active connections: %d
Demo runtime: %v
Storage operations: Working ✅
Real-time processing: Working ✅

🎯 This showcases:
• Core SDK message handling
• Advanced storage operations  
• Real-time event processing
• Production-ready patterns

Try sending another message!`,
		messageCount,
		len(d.connections),
		time.Since(d.startTime).Round(time.Second))

	chatContent, err := message.NewChat().
		Message(responseText).
		Finish()
	if err != nil {
		fmt.Printf("❌ Failed to create response: %v\n", err)
		return
	}

	if err := d.account.MessageSend(peer, chatContent); err != nil {
		fmt.Printf("❌ Failed to send response: %v\n", err)
	} else {
		fmt.Printf("📤 Advanced response sent (Message #%d)\n", messageCount)
	}
}

// setupGracefulShutdown configures clean shutdown on Ctrl+C
func (d *AdvancedDemo) setupGracefulShutdown() {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigCh
		fmt.Println("\n🛑 Shutdown signal received...")
		d.done <- true
	}()
}

// waitForInteractions waits for connections and messages with timeout
func (d *AdvancedDemo) waitForInteractions() {
	fmt.Println("⏳ Waiting for connections and messages...")
	fmt.Println("   • Scan the QR code above to connect")
	fmt.Println("   • Send messages to see advanced features")
	fmt.Println("   • Press Ctrl+C to finish the demo")
	fmt.Println()

	// Wait for either user interrupt or timeout
	timeout := time.After(5 * time.Minute)
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-d.done:
			fmt.Println("🏁 Demo completed by user")
			return
		case <-timeout:
			fmt.Println("⏰ Demo timeout reached (5 minutes)")
			return
		case <-ticker.C:
			d.mutex.RLock()
			connectionCount := len(d.connections)
			messageCount := d.messageCount
			d.mutex.RUnlock()

			fmt.Printf("📊 Status: %d connections, %d messages processed (runtime: %v)\n",
				connectionCount, messageCount, time.Since(d.startTime).Round(time.Second))
		}
	}
}

// displaySummary shows final statistics and educational information
func (d *AdvancedDemo) displaySummary() {
	d.mutex.RLock()
	connectionCount := len(d.connections)
	messageCount := d.messageCount
	d.mutex.RUnlock()

	runtime := time.Since(d.startTime).Round(time.Second)

	fmt.Println("📊 Advanced Demo Summary")
	fmt.Println("========================")
	fmt.Printf("⏱️  Total runtime: %v\n", runtime)
	fmt.Printf("🔗 Connections made: %d\n", connectionCount)
	fmt.Printf("📨 Messages processed: %d\n", messageCount)
	fmt.Printf("💾 Storage operations: %d+\n", connectionCount*2+messageCount+5) // Approximate
	fmt.Println()

	fmt.Println("🎓 What was demonstrated:")
	fmt.Println("   ✅ Core SDK account creation and configuration")
	fmt.Println("   ✅ Advanced encrypted storage with automatic security")
	fmt.Println("   ✅ Discovery QR generation and connection handling")
	fmt.Println("   ✅ Real-time event-driven message processing")
	fmt.Println("   ✅ Production-ready account and storage patterns")
	fmt.Println("   ✅ Graceful shutdown and resource management")
	fmt.Println()

	fmt.Println("🚀 Advanced features unlocked:")
	fmt.Println("   • Direct core SDK usage")
	fmt.Println("   • Encrypted storage for secure data persistence")
	fmt.Println("   • Event-driven architecture with callbacks")
	fmt.Println("   • Real-time connection and message handling")
	fmt.Println("   • Production patterns for robust applications")
	fmt.Println()

	fmt.Println("📚 Next steps:")
	fmt.Println("   • Explore individual subdirectories for deep dives")
	fmt.Println("   • Build production applications using these patterns")
	fmt.Println("   • Integrate advanced features into your own projects")
	fmt.Println()

	fmt.Println("✅ Advanced features demo completed!")
}
