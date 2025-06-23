// Package main demonstrates simple chat messaging using the underlying Self SDK.
//
// This example shows the basics of:
// - Setting up a Self account for messaging
// - Handling incoming chat messages
// - Sending automatic responses
// - Understanding the core SDK architecture
//
// 🎯 What you'll learn:
// • How to create and configure a Self account
// • How to handle incoming messages with callbacks
// • How to send chat messages using the core SDK
// • Basic message processing and response generation
//
// 💬 CORE FUNCTIONALITY DEMONSTRATED:
// • Account creation and configuration
// • Event-driven message handling
// • Chat message sending and receiving
// • End-to-end encryption (automatic)
package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/joinself/self-go-sdk/account"
	"github.com/joinself/self-go-sdk/event"
	"github.com/joinself/self-go-sdk/keypair/signing"
	"github.com/joinself/self-go-sdk/message"
)

func main() {
	fmt.Println("💬 Simple Chat Demo (Core SDK)")
	fmt.Println("===============================")
	fmt.Println("This demo shows basic chat messaging using the underlying Self SDK.")
	fmt.Println()

	// Step 1: Create and configure a Self account
	selfAccount := setupAccount()

	// Step 2: Display connection information
	showConnectionInfo(selfAccount)

	fmt.Println("✅ Chat demo ready!")
	fmt.Println()
	fmt.Println("🎓 What's running:")
	fmt.Println("   • Self account with event-driven message handling")
	fmt.Println("   • Automatic chat responses to incoming messages")
	fmt.Println("   • End-to-end encrypted communication")
	fmt.Println()
	fmt.Println("💡 To test: Connect from another Self SDK instance and send messages!")
	fmt.Println("Press Ctrl+C to exit.")

	// Keep running to handle incoming messages
	select {}
}

// setupAccount creates and configures the Self account with message handling
func setupAccount() *account.Account {
	fmt.Println("🔧 Setting up Self account...")

	// Simple account configuration focused on chat
	cfg := &account.Config{
		StorageKey:  make([]byte, 32),      // Storage encryption key
		StoragePath: "./storage",           // Local storage path
		Environment: account.TargetSandbox, // Use sandbox environment
		LogLevel:    account.LogWarn,       // Minimal logging
		Callbacks: account.Callbacks{
			// Handle incoming messages
			OnMessage: handleMessage,
			// Handle connection events (optional)
			OnConnect: func(acc *account.Account) {
				fmt.Println("🔗 Connected to Self network")
			},
		},
	}

	selfAccount, err := account.New(cfg)
	if err != nil {
		log.Fatal("❌ Failed to create Self account:", err)
	}

	fmt.Println("✅ Self account created successfully")
	return selfAccount
}

// showConnectionInfo displays information for connecting to this instance
func showConnectionInfo(selfAccount *account.Account) {
	// Open an inbox for receiving messages
	inboxAddress, err := selfAccount.InboxOpen()
	if err != nil {
		log.Fatal("❌ Failed to open inbox:", err)
	}

	fmt.Println("\n📬 Connection Information:")
	fmt.Printf("   🆔 Inbox Address: %s\n", inboxAddress.String())
	fmt.Println("   📱 To connect: Use this address in another Self SDK instance")
	fmt.Println("   🔐 All messages will be automatically encrypted")
	fmt.Println()
}

// handleMessage processes all incoming messages
func handleMessage(acc *account.Account, msg *event.Message) {
	timestamp := time.Now().Format("15:04:05")

	// Focus on chat messages, but handle others gracefully
	switch event.ContentTypeOf(msg) {
	case message.ContentTypeChat:
		handleChatMessage(acc, msg, timestamp)
	default:
		// For other message types, just acknowledge receipt
		fmt.Printf("📨 [%s] Received message from %s (type: %v)\n",
			timestamp, msg.FromAddress().String(), event.ContentTypeOf(msg))
	}
}

// handleChatMessage processes incoming chat messages and sends responses
func handleChatMessage(acc *account.Account, msg *event.Message, timestamp string) {
	// Decode the chat message
	chat, err := message.DecodeChat(msg.Content())
	if err != nil {
		fmt.Printf("❌ [%s] Failed to decode chat message: %v\n", timestamp, err)
		return
	}

	// Display the received message
	fmt.Printf("\n📨 [%s] Chat message from %s:\n", timestamp, msg.FromAddress().String())
	fmt.Printf("   💬 \"%s\"\n", chat.Message())

	// Generate and send a response
	response := generateResponse(chat.Message(), timestamp)
	sendResponse(acc, msg.FromAddress(), response, timestamp)
}

// generateResponse creates appropriate responses based on the message content
func generateResponse(messageText, timestamp string) string {
	message := strings.ToLower(strings.TrimSpace(messageText))

	switch {
	case strings.Contains(message, "hello") || strings.Contains(message, "hi"):
		return fmt.Sprintf("👋 Hello! Message received at %s via Self SDK", timestamp)
	case strings.Contains(message, "how are you"):
		return "🤖 I'm doing great! I'm a Self SDK chat demo running on the core SDK."
	case strings.Contains(message, "help"):
		return "💡 I'm a simple chat bot. Try saying 'hello', 'how are you', or send any message for an echo!"
	case strings.Contains(message, "time"):
		return fmt.Sprintf("🕐 Current time is %s", timestamp)
	default:
		return fmt.Sprintf("🔄 Echo: %s", messageText)
	}
}

// sendResponse sends a chat response to the peer
func sendResponse(acc *account.Account, toAddress *signing.PublicKey, responseText, timestamp string) {
	// Build the chat message
	chatContent, err := message.NewChat().
		Message(responseText).
		Finish()
	if err != nil {
		fmt.Printf("❌ [%s] Failed to build response: %v\n", timestamp, err)
		return
	}

	// Send the message
	fmt.Printf("📤 [%s] Sending response...\n", timestamp)
	err = acc.MessageSend(toAddress, chatContent)
	if err != nil {
		fmt.Printf("❌ [%s] Failed to send response: %v\n", timestamp, err)
	} else {
		fmt.Printf("✅ [%s] Response sent: \"%s\"\n", timestamp, responseText)
	}
	fmt.Println()
}
