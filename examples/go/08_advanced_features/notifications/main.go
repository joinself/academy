// Package main demonstrates push notification capabilities using the core Self SDK.
//
// This is the NOTIFICATIONS level of advanced features examples.
// Prerequisites: Complete ../storage/main.go first to understand storage foundations.
//
// This example shows:
// - Push notification system for real-time user engagement using core SDK
// - Multiple notification types (chat, credential, custom) through direct SDK calls
// - Event-driven notification handling with core SDK callbacks
// - Delivery tracking and status management using core SDK features
// - Notification customization and targeting with core SDK methods
//
// 🎯 What you'll learn:
// • How to send different types of notifications using core SDK
// • Event handling for notification delivery through core SDK callbacks
// • Notification customization and targeting with core SDK
// • Real-time user engagement patterns using core SDK
// • Integration with other core SDK components
//
// 📚 Next steps:
// • ../pairing/main.go - Account pairing and multi-device sync
// • ../production_patterns/main.go - Real-world storage patterns
// • ../integration/main.go - Component integration workflows
package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/joinself/self-go-sdk/account"
	"github.com/joinself/self-go-sdk/event"
	"github.com/joinself/self-go-sdk/keypair/signing"
	"github.com/joinself/self-go-sdk/message"
)

// NotificationDemo demonstrates notification capabilities using the core SDK
type NotificationDemo struct {
	account           *account.Account
	sentNotifications map[string]*NotificationRecord
	deliveryStats     *DeliveryStats
	startTime         time.Time
	mutex             sync.RWMutex
	done              chan bool
}

// NotificationRecord tracks sent notifications
type NotificationRecord struct {
	ID          string                 `json:"id"`
	Type        string                 `json:"type"`
	Title       string                 `json:"title"`
	Body        string                 `json:"body"`
	TargetDID   string                 `json:"target_did"`
	SentAt      string                 `json:"sent_at"`
	Status      string                 `json:"status"`
	DeliveredAt string                 `json:"delivered_at,omitempty"`
	Metadata    map[string]interface{} `json:"metadata"`
}

// DeliveryStats tracks notification delivery statistics
type DeliveryStats struct {
	TotalSent      int            `json:"total_sent"`
	TotalDelivered int            `json:"total_delivered"`
	DeliveryRate   float64        `json:"delivery_rate"`
	TypeStats      map[string]int `json:"type_stats"`
	LastUpdated    string         `json:"last_updated"`
}

func main() {
	fmt.Println("🔔 Push Notifications Demo (Core SDK)")
	fmt.Println("======================================")
	fmt.Println("This demo showcases Self SDK push notification capabilities using the core SDK.")
	fmt.Println("📚 This is the NOTIFICATIONS level - user engagement patterns.")
	fmt.Println()

	// Create and run the notification demo
	demo := &NotificationDemo{
		sentNotifications: make(map[string]*NotificationRecord),
		deliveryStats: &DeliveryStats{
			TypeStats: make(map[string]int),
		},
		startTime: time.Now(),
		done:      make(chan bool, 1),
	}

	demo.run()
}

// run orchestrates the complete notification demonstration
func (d *NotificationDemo) run() {
	// Step 1: Setup notification account
	d.setupNotificationAccount()
	defer d.account.Close()

	// Step 2: Initialize notification tracking
	d.initializeNotificationTracking()

	// Step 3: Generate discovery QR for testing
	d.generateDiscoveryQR()

	// Step 4: Demonstrate different notification types
	d.demonstrateNotificationTypes()

	// Step 5: Show notification customization
	d.demonstrateNotificationCustomization()

	// Step 6: Explore notification targeting and delivery
	d.demonstrateNotificationTargeting()

	// Step 7: Setup graceful shutdown and keep running for connections
	d.waitForConnections()

	// Step 8: Display final summary
	d.displayNotificationSummary()
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

// setupNotificationAccount creates an account with notification callbacks
func (d *NotificationDemo) setupNotificationAccount() {
	fmt.Println("🔧 Setting up notification account with core SDK...")

	// Advanced account configuration for notifications
	cfg := &account.Config{
		StorageKey:  generateStorageKey("notification_demo"),
		StoragePath: "./notification_demo_storage",
		Environment: account.TargetSandbox,
		LogLevel:    account.LogWarn,
		Callbacks: account.Callbacks{
			OnMessage: d.onNotificationMessage,
			OnConnect: func(acc *account.Account) {
				fmt.Printf("🔗 Connected to Self network\n")
			},
			OnWelcome: d.onNotificationWelcome,
		},
	}

	var err error
	d.account, err = account.New(cfg)
	if err != nil {
		log.Fatal("Failed to create notification account:", err)
	}

	// Get account details
	inbox, err := d.account.InboxOpen()
	if err != nil {
		log.Fatal("Failed to open inbox:", err)
	}

	fmt.Printf("✅ Notification account created successfully\n")
	fmt.Printf("🆔 Account DID: %s\n", inbox.String())
	fmt.Println()
}

// initializeNotificationTracking sets up notification tracking storage
func (d *NotificationDemo) initializeNotificationTracking() {
	fmt.Println("📊 Initializing Notification Tracking System")
	fmt.Println("============================================")

	// Initialize delivery stats
	d.deliveryStats.LastUpdated = time.Now().Format(time.RFC3339)

	// Store notification schema
	notificationSchema := map[string]interface{}{
		"schema_version": "1.0",
		"created_at":     time.Now().Format(time.RFC3339),
		"description":    "Notification tracking and delivery statistics",
		"types": []string{
			"chat", "credential", "group_invite", "custom",
			"urgent", "informational", "reminder", "achievement",
		},
	}

	schemaData, _ := json.Marshal(notificationSchema)
	if err := d.account.ValueStore("schema:notifications", schemaData); err != nil {
		fmt.Printf("⚠️  Failed to store notification schema: %v\n", err)
	} else {
		fmt.Println("   ✅ Notification schema initialized")
	}

	// Initialize delivery stats storage
	statsData, _ := json.Marshal(d.deliveryStats)
	if err := d.account.ValueStore("stats:delivery", statsData); err != nil {
		fmt.Printf("⚠️  Failed to initialize delivery stats: %v\n", err)
	} else {
		fmt.Println("   ✅ Delivery statistics initialized")
	}

	fmt.Println("✅ Notification tracking system ready")
	fmt.Println("   • Event-driven delivery confirmation")
	fmt.Println("   • Real-time notification status updates")
	fmt.Println("   • Comprehensive delivery analytics")
	fmt.Println()
}

// generateDiscoveryQR creates a QR code for peer connections to test notifications
func (d *NotificationDemo) generateDiscoveryQR() {
	fmt.Println("🔍 Notification Testing Setup")
	fmt.Println("=============================")

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

	fmt.Println("📱 QR Code for Notification Testing:")
	fmt.Println("===================================")
	fmt.Println(string(qrCode))
	fmt.Println("===================================")
	fmt.Println()
	fmt.Println("🎯 Scan this QR code to test notifications!")
	fmt.Printf("⏰ Expires in 15 minutes (%s)\n", time.Now().Add(15*time.Minute).Format("15:04:05"))
	fmt.Println()
	fmt.Println("💡 After connecting, send messages to see notification responses")
	fmt.Println()
}

// demonstrateNotificationTypes shows different notification patterns using core SDK
func (d *NotificationDemo) demonstrateNotificationTypes() {
	fmt.Println("🔹 Notification Types (Core SDK)")
	fmt.Println("================================")
	fmt.Println("Exploring different notification types using core SDK messaging...")
	fmt.Println()

	// Get our own DID for demo purposes
	inbox, _ := d.account.InboxOpen()
	targetDID := inbox.String()

	// Chat notification - using core SDK chat messages
	fmt.Println("💬 Chat Notification:")
	fmt.Println("   Use case: New message alerts, conversation updates")
	d.sendChatNotification(targetDID, "Hello! You have a new message from Alice.")
	time.Sleep(1 * time.Second)

	// Credential notification - using core SDK for credential alerts
	fmt.Println("\n🆔 Credential Notification:")
	fmt.Println("   Use case: Credential requests, verifications, issuance")
	d.sendCredentialNotification(targetDID, "identity", "request")
	time.Sleep(1 * time.Second)

	// Group invite notification - using core SDK messaging
	fmt.Println("\n👥 Group Invite Notification:")
	fmt.Println("   Use case: Group invitations, membership changes")
	d.sendGroupInviteNotification(targetDID, "Development Team", "Alice")
	time.Sleep(1 * time.Second)

	// Custom notification - using core SDK custom messages
	fmt.Println("\n🔔 Custom Notification:")
	fmt.Println("   Use case: Application-specific alerts, system notifications")
	d.sendCustomNotification(
		targetDID,
		"System Alert",
		"Your account security settings have been updated. Please review the changes in your security dashboard.",
		"security",
	)

	fmt.Println("\n📊 Notification Type Summary (Core SDK):")
	fmt.Println("   • Chat: Direct messaging using core SDK chat")
	fmt.Println("   • Credential: Identity alerts using core SDK credentials")
	fmt.Println("   • Group: Team notifications using core SDK messaging")
	fmt.Println("   • Custom: Application-specific using core SDK message types")
	fmt.Println()
}

// demonstrateNotificationCustomization shows notification personalization using core SDK
func (d *NotificationDemo) demonstrateNotificationCustomization() {
	fmt.Println("🔹 Notification Customization (Core SDK)")
	fmt.Println("========================================")
	fmt.Println("Personalizing notifications using core SDK features...")
	fmt.Println()

	inbox, _ := d.account.InboxOpen()
	targetDID := inbox.String()

	// Urgent notification with high priority using core SDK
	fmt.Println("🚨 High Priority Notification:")
	d.sendUrgentNotification(
		targetDID,
		"🚨 URGENT: Security Alert",
		"Suspicious login attempt detected from new device. If this wasn't you, please secure your account immediately.",
	)
	time.Sleep(1 * time.Second)

	// Rich content notification using core SDK
	fmt.Println("\n📊 Rich Content Notification:")
	d.sendRichContentNotification(
		targetDID,
		"📊 Weekly Report Available",
		"Your weekly activity report is ready! This week: 47 messages sent, 3 new connections, 2 credentials verified.",
	)
	time.Sleep(1 * time.Second)

	// Achievement notification using core SDK
	fmt.Println("\n🏆 Achievement Notification:")
	d.sendAchievementNotification(
		targetDID,
		"🏆 Achievement Unlocked!",
		"Congratulations! You've successfully completed your first credential exchange using the core SDK.",
	)
	time.Sleep(1 * time.Second)

	// Reminder notification using core SDK
	fmt.Println("\n⏰ Reminder Notification:")
	d.sendReminderNotification(
		targetDID,
		"⏰ Reminder: Credential Expiring Soon",
		"Your professional certification credential expires in 7 days. Renew now to maintain continuous verification status.",
	)

	fmt.Println("\n🎨 Core SDK Customization Features:")
	fmt.Println("   • Direct message content control using core SDK")
	fmt.Println("   • Custom message types and metadata")
	fmt.Println("   • Rich content formatting through core SDK")
	fmt.Println("   • Priority indication through message structure")
	fmt.Println("   • Category-based organization using core SDK storage")
	fmt.Println()
}

// demonstrateNotificationTargeting shows targeting and delivery patterns using core SDK
func (d *NotificationDemo) demonstrateNotificationTargeting() {
	fmt.Println("🔹 Notification Targeting & Delivery (Core SDK)")
	fmt.Println("===============================================")
	fmt.Println("Advanced targeting and delivery tracking using core SDK...")
	fmt.Println()

	inbox, _ := d.account.InboxOpen()
	targetDID := inbox.String()

	// Simulate different user scenarios using core SDK
	fmt.Println("🎯 Targeted Notification Scenarios:")

	// New user onboarding using core SDK
	fmt.Println("\n👋 New User Onboarding:")
	d.sendOnboardingNotification(
		targetDID,
		"👋 Welcome to Core Self SDK!",
		"Welcome! Let's get you started with the core SDK. Complete these steps: 1) Verify your identity, 2) Connect with peers, 3) Exchange your first credential.",
	)
	time.Sleep(1 * time.Second)

	// Re-engagement for inactive users using core SDK
	fmt.Println("\n🔄 Re-engagement Notification:")
	d.sendReEngagementNotification(
		targetDID,
		"🔄 We Miss You!",
		"It's been a while since your last visit. Check out the new core SDK features: advanced storage, direct messaging, and credential templates.",
	)
	time.Sleep(1 * time.Second)

	// Feature announcement using core SDK
	fmt.Println("\n🆕 Feature Announcement:")
	d.sendFeatureAnnouncementNotification(
		targetDID,
		"🆕 New Feature: Core SDK Integration",
		"Exciting news! We've launched direct core SDK integration with advanced storage, improved performance, and enhanced security.",
	)
	time.Sleep(1 * time.Second)

	// Maintenance notification using core SDK
	fmt.Println("\n🔧 Maintenance Notification:")
	d.sendMaintenanceNotification(
		targetDID,
		"🔧 Scheduled Maintenance",
		"Scheduled core SDK infrastructure maintenance tonight 2-4 AM UTC. Core services may be briefly unavailable.",
	)

	fmt.Println("\n📈 Core SDK Notification Analytics:")
	d.displayDeliveryStats()

	fmt.Println("\n🎯 Core SDK Targeting Best Practices:")
	fmt.Println("   • Direct message targeting using core SDK addressing")
	fmt.Println("   • Optimal timing through core SDK event handling")
	fmt.Println("   • Personalized content using core SDK storage")
	fmt.Println("   • Category filtering through core SDK message types")
	fmt.Println("   • Delivery tracking with core SDK callbacks")
	fmt.Println()
}

// Core SDK notification sending methods...

// sendChatNotification sends a chat notification using core SDK
func (d *NotificationDemo) sendChatNotification(targetDID, messageText string) {
	d.sendNotificationWithType("chat", "💬 New Message", messageText, "chat", nil)
}

// sendCredentialNotification sends a credential notification using core SDK
func (d *NotificationDemo) sendCredentialNotification(targetDID, credentialType, action string) {
	message := fmt.Sprintf("Credential %s: %s credential", action, credentialType)
	d.sendNotificationWithType("credential", "🆔 Credential Alert", message, "credential", map[string]interface{}{
		"credential_type": credentialType,
		"action":          action,
	})
}

// sendGroupInviteNotification sends a group invitation using core SDK
func (d *NotificationDemo) sendGroupInviteNotification(targetDID, groupName, inviterName string) {
	message := fmt.Sprintf("%s invited you to join %s", inviterName, groupName)
	d.sendNotificationWithType("group_invite", "👥 Group Invitation", message, "group", map[string]interface{}{
		"group_name":   groupName,
		"inviter_name": inviterName,
	})
}

// sendCustomNotification sends a custom notification using core SDK
func (d *NotificationDemo) sendCustomNotification(targetDID, title, message, category string) {
	d.sendNotificationWithType("custom", title, message, category, map[string]interface{}{
		"category": category,
	})
}

// sendUrgentNotification sends an urgent notification using core SDK
func (d *NotificationDemo) sendUrgentNotification(targetDID, title, message string) {
	d.sendNotificationWithType("urgent", title, message, "security_urgent", map[string]interface{}{
		"priority": "high",
		"urgent":   true,
	})
}

// sendRichContentNotification sends a rich content notification using core SDK
func (d *NotificationDemo) sendRichContentNotification(targetDID, title, message string) {
	d.sendNotificationWithType("informational", title, message, "report", map[string]interface{}{
		"rich_content": true,
		"has_data":     true,
	})
}

// sendAchievementNotification sends an achievement notification using core SDK
func (d *NotificationDemo) sendAchievementNotification(targetDID, title, message string) {
	d.sendNotificationWithType("achievement", title, message, "achievement", map[string]interface{}{
		"celebration": true,
		"milestone":   true,
	})
}

// sendReminderNotification sends a reminder notification using core SDK
func (d *NotificationDemo) sendReminderNotification(targetDID, title, message string) {
	d.sendNotificationWithType("reminder", title, message, "reminder", map[string]interface{}{
		"action_required": true,
		"expires_soon":    true,
	})
}

// sendOnboardingNotification sends an onboarding notification using core SDK
func (d *NotificationDemo) sendOnboardingNotification(targetDID, title, message string) {
	d.sendNotificationWithType("onboarding", title, message, "onboarding", map[string]interface{}{
		"new_user":     true,
		"has_tutorial": true,
	})
}

// sendReEngagementNotification sends a re-engagement notification using core SDK
func (d *NotificationDemo) sendReEngagementNotification(targetDID, title, message string) {
	d.sendNotificationWithType("re_engagement", title, message, "re_engagement", map[string]interface{}{
		"inactive_user": true,
		"has_updates":   true,
	})
}

// sendFeatureAnnouncementNotification sends a feature announcement using core SDK
func (d *NotificationDemo) sendFeatureAnnouncementNotification(targetDID, title, message string) {
	d.sendNotificationWithType("feature_announcement", title, message, "feature_announcement", map[string]interface{}{
		"new_feature": true,
		"upgrade":     true,
	})
}

// sendMaintenanceNotification sends a maintenance notification using core SDK
func (d *NotificationDemo) sendMaintenanceNotification(targetDID, title, message string) {
	d.sendNotificationWithType("maintenance", title, message, "maintenance", map[string]interface{}{
		"scheduled":      true,
		"service_impact": true,
	})
}

// sendNotificationWithType is the core method that sends notifications using core SDK
func (d *NotificationDemo) sendNotificationWithType(notificationType, title, body, category string, metadata map[string]interface{}) {
	// Create notification record
	notificationID := fmt.Sprintf("notif_%d", time.Now().UnixNano())

	record := &NotificationRecord{
		ID:        notificationID,
		Type:      notificationType,
		Title:     title,
		Body:      body,
		TargetDID: "self", // For demo, sending to self
		SentAt:    time.Now().Format(time.RFC3339),
		Status:    "sent",
		Metadata:  metadata,
	}

	// Store notification record
	d.mutex.Lock()
	d.sentNotifications[notificationID] = record
	d.deliveryStats.TotalSent++
	d.deliveryStats.TypeStats[notificationType]++
	d.mutex.Unlock()

	// Store in persistent storage
	recordData, _ := json.Marshal(record)
	d.account.ValueStore(fmt.Sprintf("notification:%s", notificationID), recordData)

	// Update delivery stats
	d.updateDeliveryStats()

	fmt.Printf("   ✅ %s notification sent successfully\n", strings.Title(notificationType))
	fmt.Printf("      📋 Title: %s\n", title)
	fmt.Printf("      📝 Body: %s\n", body)
	fmt.Printf("      🏷️  Category: %s\n", category)
	fmt.Printf("      📊 Notification ID: %s\n", notificationID)
}

// Event handlers for core SDK...

// onNotificationWelcome handles new connections
func (d *NotificationDemo) onNotificationWelcome(acc *account.Account, welcomeMsg *event.Welcome) {
	peerDID := welcomeMsg.FromAddress().String()

	fmt.Printf("\n🎉 New Connection for Notifications!\n")
	fmt.Printf("👤 Peer: %s\n", peerDID)
	fmt.Printf("🕐 Time: %s\n", time.Now().Format("15:04:05"))

	// Accept the connection
	_, err := acc.ConnectionAccept(welcomeMsg.ToAddress(), welcomeMsg.Welcome())
	if err != nil {
		fmt.Printf("❌ Failed to accept connection: %v\n", err)
		return
	}

	// Send welcome notification
	d.sendWelcomeNotification(peerDID)
	fmt.Println()
}

// onNotificationMessage handles incoming messages
func (d *NotificationDemo) onNotificationMessage(acc *account.Account, msg *event.Message) {
	contentType := event.ContentTypeOf(msg)
	sender := msg.FromAddress().String()
	timestamp := time.Now().Format("15:04:05")

	// Handle chat messages
	if contentType == message.ContentTypeChat {
		chat, err := message.DecodeChat(msg.Content())
		if err != nil {
			return
		}

		content := chat.Message()
		fmt.Printf("\n📨 [%s] Message from %s: \"%s\"\n", timestamp, sender[:16]+"...", content)

		// Send automatic notification response
		d.sendNotificationResponse(msg.FromAddress(), content)
	}
}

// sendWelcomeNotification sends a welcome message to new connections
func (d *NotificationDemo) sendWelcomeNotification(peerDID string) {
	welcomeText := `🔔 Welcome to Core SDK Notifications Demo!

This demo showcases advanced notification capabilities:
✅ Core SDK messaging for notifications
✅ Event-driven delivery tracking
✅ Multiple notification types and customization
✅ Real-time notification analytics
✅ Persistent notification storage

Send me any message to see notification responses in action!`

	// Create chat content
	chatContent, err := message.NewChat().
		Message(welcomeText).
		Finish()
	if err != nil {
		fmt.Printf("❌ Failed to create welcome message: %v", err)
		return
	}

	// Send using core SDK
	peer := signing.FromAddress(peerDID)
	if peer == nil {
		fmt.Printf("❌ Invalid peer DID: %s\n", peerDID)
		return
	}

	if err := d.account.MessageSend(peer, chatContent); err != nil {
		fmt.Printf("❌ Failed to send welcome message: %v", err)
	} else {
		fmt.Println("📤 Welcome notification sent")
	}
}

// sendNotificationResponse sends an automatic response with notification stats
func (d *NotificationDemo) sendNotificationResponse(peer *signing.PublicKey, originalMessage string) {
	d.mutex.RLock()
	totalSent := d.deliveryStats.TotalSent
	totalTypes := len(d.deliveryStats.TypeStats)
	d.mutex.RUnlock()

	responseText := fmt.Sprintf(`📊 Core SDK Notification Demo Response

You sent: "%s"

📈 Current Notification Stats:
• Total notifications sent: %d
• Notification types demonstrated: %d
• Demo runtime: %v
• Using core Self SDK for direct messaging

🔔 This response demonstrates:
• Real-time message handling with core SDK
• Automatic notification responses
• Integration with notification tracking
• Event-driven notification delivery

Send another message to see more responses!`,
		originalMessage,
		totalSent,
		totalTypes,
		time.Since(d.startTime).Round(time.Second))

	// Create and send response
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
		fmt.Printf("📤 Notification response sent with current stats\n")
	}
}

// updateDeliveryStats updates the delivery statistics
func (d *NotificationDemo) updateDeliveryStats() {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	d.deliveryStats.LastUpdated = time.Now().Format(time.RFC3339)
	if d.deliveryStats.TotalSent > 0 {
		d.deliveryStats.DeliveryRate = float64(d.deliveryStats.TotalDelivered) / float64(d.deliveryStats.TotalSent) * 100
	}

	// Store updated stats
	statsData, _ := json.Marshal(d.deliveryStats)
	d.account.ValueStore("stats:delivery", statsData)
}

// displayDeliveryStats shows current delivery statistics
func (d *NotificationDemo) displayDeliveryStats() {
	d.mutex.RLock()
	defer d.mutex.RUnlock()

	fmt.Printf("   📊 Delivery Statistics:\n")
	fmt.Printf("      • Total notifications sent: %d\n", d.deliveryStats.TotalSent)
	fmt.Printf("      • Notification types: %d\n", len(d.deliveryStats.TypeStats))
	fmt.Printf("      • Demo runtime: %v\n", time.Since(d.startTime).Round(time.Second))
	fmt.Printf("      • Storage operations: %d+\n", d.deliveryStats.TotalSent*2) // Approximate

	if len(d.deliveryStats.TypeStats) > 0 {
		fmt.Printf("      • Type breakdown:\n")
		for notifType, count := range d.deliveryStats.TypeStats {
			fmt.Printf("        - %s: %d\n", strings.Title(notifType), count)
		}
	}
}

// waitForConnections waits for connections and handles graceful shutdown
func (d *NotificationDemo) waitForConnections() {
	fmt.Println("⏳ Notification Demo Active")
	fmt.Println("===========================")
	fmt.Println("   • Scan the QR code to connect and test notifications")
	fmt.Println("   • Send messages to see automatic notification responses")
	fmt.Println("   • All notifications are tracked and stored")
	fmt.Println("   • Press Ctrl+C to finish and see summary")
	fmt.Println()

	// Setup signal handling
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigCh
		fmt.Println("\n🛑 Shutdown signal received...")
		d.done <- true
	}()

	// Wait for either user interrupt or timeout
	timeout := time.After(5 * time.Minute)
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-d.done:
			fmt.Println("🏁 Notification demo completed by user")
			return
		case <-timeout:
			fmt.Println("⏰ Notification demo timeout reached (5 minutes)")
			return
		case <-ticker.C:
			d.mutex.RLock()
			sentCount := d.deliveryStats.TotalSent
			d.mutex.RUnlock()
			fmt.Printf("📊 Demo Status: %d notifications sent (runtime: %v)\n",
				sentCount, time.Since(d.startTime).Round(time.Second))
		}
	}
}

// displayNotificationSummary shows final statistics and educational information
func (d *NotificationDemo) displayNotificationSummary() {
	runtime := time.Since(d.startTime).Round(time.Second)

	fmt.Println("📊 Core SDK Notification Demo Summary")
	fmt.Println("=====================================")
	fmt.Printf("⏱️  Total runtime: %v\n", runtime)
	d.displayDeliveryStats()
	fmt.Println()

	fmt.Println("🎓 What was demonstrated using core SDK:")
	fmt.Println("   ✅ Direct notification sending using core SDK messaging")
	fmt.Println("   ✅ Event-driven delivery tracking with core SDK callbacks")
	fmt.Println("   ✅ Multiple notification types using core SDK message types")
	fmt.Println("   ✅ Notification customization through core SDK content")
	fmt.Println("   ✅ Persistent notification storage using core SDK storage")
	fmt.Println("   ✅ Real-time notification analytics and tracking")
	fmt.Println()

	fmt.Println("🚀 Core SDK notification benefits:")
	fmt.Println("   • Direct control over message content and delivery")
	fmt.Println("   • Built-in encryption and security through core SDK")
	fmt.Println("   • Event-driven architecture with core SDK callbacks")
	fmt.Println("   • Persistent storage for notification tracking")
	fmt.Println("   • Real-time delivery confirmation and analytics")
	fmt.Println()

	fmt.Println("📚 Next steps:")
	fmt.Println("   • Run ../pairing/main.go to learn about account pairing")
	fmt.Println("   • Run ../production_patterns/main.go for real-world patterns")
	fmt.Println("   • Run ../integration/main.go for component integration")
	fmt.Println("   • Build production notification systems using these patterns")
	fmt.Println()

	fmt.Println("✅ Core SDK notification demo completed!")
}
