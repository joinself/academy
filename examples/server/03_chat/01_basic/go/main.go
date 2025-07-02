// Package main demonstrates simple chat messaging using the Self SDK
package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/joinself/academy/examples/go/common"
	"github.com/joinself/self-go-sdk/account"
	"github.com/joinself/self-go-sdk/event"
	"github.com/joinself/self-go-sdk/keypair/signing"
	"github.com/joinself/self-go-sdk/message"
)

func main() {
	fmt.Println("💬 Simple Chat Demo")
	fmt.Println("===================")

	// Create and configure Self account
	selfAccount := common.SetupAccount(common.AccountConfig{
		Callbacks: account.Callbacks{
			OnMessage: handleMessage,
			OnConnect: func(acc *account.Account) {
				fmt.Println("✅ Connected to Self network")
			},
		},
	})

	// Display connection information
	common.DisplayAccountInfo(selfAccount, "Chat Account")

	fmt.Println("✅ Chat demo ready! Press Ctrl+C to exit.")

	// Keep running to handle incoming messages
	select {}
}

// handleMessage processes all incoming messages
func handleMessage(acc *account.Account, msg *event.Message) {
	timestamp := time.Now().Format("15:04:05")

	switch event.ContentTypeOf(msg) {
	case message.ContentTypeChat:
		handleChatMessage(acc, msg, timestamp)
	default:
		fmt.Printf("📨 [%s] Received message from %s\n",
			timestamp, msg.FromAddress().String())
	}
}

// handleChatMessage processes incoming chat messages and sends responses
func handleChatMessage(acc *account.Account, msg *event.Message, timestamp string) {
	chat, err := message.DecodeChat(msg.Content())
	if err != nil {
		fmt.Printf("❌ [%s] Failed to decode chat message: %v\n", timestamp, err)
		return
	}

	fmt.Printf("📨 [%s] %s: \"%s\"\n", timestamp, msg.FromAddress().String(), chat.Message())

	response := generateResponse(chat.Message(), timestamp)
	sendResponse(acc, msg.FromAddress(), response, timestamp)
}

// generateResponse creates appropriate responses based on the message content
func generateResponse(messageText, timestamp string) string {
	message := strings.ToLower(strings.TrimSpace(messageText))

	switch {
	case strings.Contains(message, "hello") || strings.Contains(message, "hi"):
		return fmt.Sprintf("👋 Hello! Message received at %s", timestamp)
	case strings.Contains(message, "how are you"):
		return "🤖 I'm doing great! I'm a Self SDK chat demo."
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
	chatContent, err := message.NewChat().
		Message(responseText).
		Finish()
	if err != nil {
		fmt.Printf("❌ [%s] Failed to build response: %v\n", timestamp, err)
		return
	}

	err = acc.MessageSend(toAddress, chatContent)
	if err != nil {
		fmt.Printf("❌ [%s] Failed to send response: %v\n", timestamp, err)
	} else {
		fmt.Printf("📤 [%s] Sent: \"%s\"\n", timestamp, responseText)
	}
}
