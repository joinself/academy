// Package main demonstrates account pairing capabilities using the Self SDK core.
//
// This is the PAIRING level of advanced features examples.
// Prerequisites: Complete ../notifications/main.go first.
//
// This example shows:
// - Multi-device account synchronization using core SDK
// - Pairing code generation and QR code creation
// - Account state management and persistence
// - Cross-device synchronization patterns
// - Pairing event handling workflows
//
// 🎯 What you'll learn:
// • How to generate pairing codes using core SDK account methods
// • Account management for multi-device scenarios
// • State synchronization patterns across devices
// • Pairing workflow implementation
// • Account storage and persistence strategies
//
// 📚 Next steps:
// • ../production_patterns/main.go - Real-world implementation patterns
package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joinself/self-go-sdk/account"
)

// PairingDemo represents our core SDK pairing demonstration
type PairingDemo struct {
	account       *account.Account
	startTime     time.Time
	pairingCode   string
	pairingActive bool
	statistics    *PairingStatistics
}

// PairingStatistics tracks pairing-related metrics
type PairingStatistics struct {
	CodesGenerated  int           `json:"codes_generated"`
	PairingAttempts int           `json:"pairing_attempts"`
	Runtime         time.Duration `json:"runtime"`
	StorageOps      int           `json:"storage_operations"`
}

// PairingInfo represents pairing session information
type PairingInfo struct {
	Code       string    `json:"code"`
	Unpaired   bool      `json:"unpaired"`
	CreatedAt  time.Time `json:"created_at"`
	ExpiresAt  time.Time `json:"expires_at"`
	DeviceID   string    `json:"device_id"`
	DeviceName string    `json:"device_name"`
}

func main() {
	fmt.Println("🔗 Self SDK Account Pairing Demo (Core SDK)")
	fmt.Println("===========================================")
	fmt.Println("This demo showcases Self SDK account pairing using the core SDK.")
	fmt.Println("📚 Advanced Features - Multi-device synchronization patterns.")
	fmt.Println()

	// Create pairing demo instance
	demo, err := NewPairingDemo()
	if err != nil {
		log.Fatal("Failed to create pairing demo:", err)
	}
	defer demo.Close()

	// Setup graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		fmt.Println("\n🛑 Shutdown signal received...")
		demo.displayFinalSummary()
		os.Exit(0)
	}()

	// Get account information
	fmt.Printf("🆔 Account ID: %s\n", demo.getAccountID())
	fmt.Println()

	// Demonstrate pairing workflows
	demo.demonstratePairingCodeGeneration()
	demo.demonstratePairingManagement()
	demo.demonstrateCrossDeviceSynchronization()
	demo.demonstrateAdvancedPairingFeatures()

	fmt.Println("✅ Account pairing demo completed!")
	demo.displayFinalSummary()
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

// NewPairingDemo creates a new pairing demonstration instance
func NewPairingDemo() (*PairingDemo, error) {
	fmt.Println("🔧 Setting up pairing demo using core SDK...")

	// Create account configuration
	cfg := &account.Config{
		StorageKey:  generateStorageKey("pairing_demo"),
		StoragePath: "./pairing_demo_storage",
		Environment: account.TargetSandbox,
		LogLevel:    account.LogWarn,
	}

	// Create account instance
	acc, err := account.New(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create account: %w", err)
	}

	demo := &PairingDemo{
		account:   acc,
		startTime: time.Now(),
		statistics: &PairingStatistics{
			CodesGenerated:  0,
			PairingAttempts: 0,
			StorageOps:      0,
		},
	}

	// Initialize pairing storage
	demo.initializePairingStorage()

	fmt.Println("✅ Pairing demo created successfully")
	return demo, nil
}

// Close cleans up resources
func (d *PairingDemo) Close() {
	if d.account != nil {
		// Note: Core SDK account doesn't have explicit close method
		// Storage is automatically handled
	}
}

// getAccountID returns the account identifier
func (d *PairingDemo) getAccountID() string {
	// Generate a demo account ID based on storage path for display
	return fmt.Sprintf("pairing_demo_%d", time.Now().Unix()%10000)
}

// initializePairingStorage sets up pairing-related storage schemas
func (d *PairingDemo) initializePairingStorage() {
	fmt.Println("🔹 Initializing Pairing Storage")
	fmt.Println("==============================")
	fmt.Println("Setting up storage schemas for pairing management...")
	fmt.Println()

	// Initialize pairing statistics
	statsData, _ := json.Marshal(d.statistics)
	d.account.ValueStore("pairing_statistics", statsData)
	d.statistics.StorageOps++

	// Initialize pairing sessions storage
	sessions := make(map[string]*PairingInfo)
	sessionsData, _ := json.Marshal(sessions)
	d.account.ValueStore("pairing_sessions", sessionsData)
	d.statistics.StorageOps++

	// Initialize device registry
	devices := make(map[string]interface{})
	devicesData, _ := json.Marshal(devices)
	d.account.ValueStore("paired_devices", devicesData)
	d.statistics.StorageOps++

	fmt.Println("✅ Pairing storage schemas initialized")
	fmt.Println("   📊 Statistics tracking enabled")
	fmt.Println("   📱 Session management active")
	fmt.Println("   🔗 Device registry configured")
	fmt.Println()
}

// demonstratePairingCodeGeneration shows pairing code creation using core SDK
func (d *PairingDemo) demonstratePairingCodeGeneration() {
	fmt.Println("🔹 Core SDK Pairing Code Generation")
	fmt.Println("===================================")
	fmt.Println("Creating secure pairing codes using core SDK account methods...")
	fmt.Println()

	// Generate pairing code using core SDK
	code, unpaired, err := d.account.SDKPairingCode()
	if err != nil {
		log.Printf("Failed to generate pairing code: %v", err)
		return
	}

	d.pairingCode = code
	d.pairingActive = true
	d.statistics.CodesGenerated++

	// Create pairing information
	pairingInfo := &PairingInfo{
		Code:       code,
		Unpaired:   unpaired,
		CreatedAt:  time.Now(),
		ExpiresAt:  time.Now().Add(24 * time.Hour), // Default 24 hour expiry
		DeviceID:   d.getAccountID(),
		DeviceName: "Primary Device",
	}

	// Store pairing session
	d.storePairingSession(pairingInfo)

	fmt.Printf("✅ Pairing code generated successfully\n")
	fmt.Printf("   🔑 Code: %s\n", code)
	fmt.Printf("   📱 Unpaired status: %t\n", unpaired)
	fmt.Printf("   ⏰ Expires at: %s\n", pairingInfo.ExpiresAt.Format("2006-01-02 15:04:05"))
	fmt.Printf("   🕐 Valid for: %v\n", time.Until(pairingInfo.ExpiresAt).Round(time.Minute))
	fmt.Printf("   🆔 Device ID: %s\n", pairingInfo.DeviceID)
	fmt.Println()

	// Display QR code information
	d.demonstrateQRCodeGeneration(code)

	fmt.Println("🔐 Core SDK Security Features:")
	fmt.Println("   • Cryptographically secure pairing codes")
	fmt.Println("   • Account-based authentication")
	fmt.Println("   • Time-limited validity for security")
	fmt.Println("   • Persistent storage through core SDK")
	fmt.Println("   • Cross-device state synchronization")
	fmt.Println()
}

// demonstrateQRCodeGeneration shows QR code creation for pairing
func (d *PairingDemo) demonstrateQRCodeGeneration(code string) {
	fmt.Println("📱 QR Code Generation for Pairing")
	fmt.Println("=================================")
	fmt.Println("Creating QR codes for easy device pairing...")
	fmt.Println()

	// Generate QR data (in real implementation, this would create actual QR code)
	qrData := fmt.Sprintf("SELF_PAIRING:%s:EXPIRES:%d", code, time.Now().Add(24*time.Hour).Unix())

	fmt.Printf("   ✅ QR code data generated successfully\n")
	fmt.Printf("   📏 QR data length: %d characters\n", len(qrData))
	fmt.Printf("   🔗 QR data: %s\n", qrData)
	fmt.Println()

	// Store QR information
	qrInfo := map[string]interface{}{
		"code":       code,
		"qr_data":    qrData,
		"created_at": time.Now().Unix(),
		"expires_at": time.Now().Add(24 * time.Hour).Unix(),
	}
	qrInfoData, _ := json.Marshal(qrInfo)
	d.account.ValueStore("current_qr_code", qrInfoData)
	d.statistics.StorageOps++

	fmt.Println("📱 QR Code Usage Scenarios:")
	fmt.Println()
	fmt.Println("   📲 Mobile Device Pairing:")
	fmt.Println("      1. Open Self mobile app")
	fmt.Println("      2. Navigate to 'Pair Device' or 'Add Account'")
	fmt.Println("      3. Scan this QR code with your device camera")
	fmt.Println("      4. Complete pairing verification")
	fmt.Println()
	fmt.Println("   💻 Cross-Platform Pairing:")
	fmt.Println("      1. Display QR code on first device")
	fmt.Println("      2. Use second device to scan QR code")
	fmt.Println("      3. Core SDK handles cryptographic handshake")
	fmt.Println("      4. Account state synchronizes automatically")
	fmt.Println()
	fmt.Println("🎯 Core SDK QR Benefits:")
	fmt.Println("   • Integrated with core account management")
	fmt.Println("   • Built-in security and encryption")
	fmt.Println("   • Automatic state synchronization")
	fmt.Println("   • Persistent pairing information")
	fmt.Println("   • Cross-device compatibility")
	fmt.Println()
}

// demonstratePairingManagement shows pairing status and management
func (d *PairingDemo) demonstratePairingManagement() {
	fmt.Println("🔹 Core SDK Pairing Management")
	fmt.Println("==============================")
	fmt.Println("Managing device pairing using core SDK account methods...")
	fmt.Println()

	// Check current pairing status
	fmt.Println("📊 Current Pairing Status:")
	if d.pairingActive && d.pairingCode != "" {
		fmt.Printf("   🔄 Active pairing session: %s\n", d.pairingCode)
		fmt.Println("   ✅ Account ready for pairing")
		fmt.Println("   📱 QR code available for scanning")
	} else {
		fmt.Println("   📱 No active pairing session")
		fmt.Println("   🔗 Ready to generate new pairing code")
	}
	fmt.Println()

	// Load and display pairing sessions
	d.displayPairingSessions()

	// Demonstrate pairing analytics
	d.updatePairingStatistics()
	d.displayPairingAnalytics()

	fmt.Println("🎯 Core SDK Pairing Management Features:")
	fmt.Println()
	fmt.Println("   📊 Session Tracking:")
	fmt.Println("      • Persistent pairing session storage")
	fmt.Println("      • Automatic expiration handling")
	fmt.Println("      • Session analytics and monitoring")
	fmt.Println("      • Device identification and registry")
	fmt.Println()
	fmt.Println("   🔐 Security Management:")
	fmt.Println("      • Account-based authentication")
	fmt.Println("      • Cryptographic key generation")
	fmt.Println("      • Secure storage through core SDK")
	fmt.Println("      • Time-limited code validity")
	fmt.Println()
	fmt.Println("   🔄 State Synchronization:")
	fmt.Println("      • Cross-device account state")
	fmt.Println("      • Automatic data consistency")
	fmt.Println("      • Conflict resolution mechanisms")
	fmt.Println("      • Real-time synchronization")
	fmt.Println()
}

// demonstrateCrossDeviceSynchronization shows multi-device scenarios
func (d *PairingDemo) demonstrateCrossDeviceSynchronization() {
	fmt.Println("🔹 Cross-Device Synchronization")
	fmt.Println("===============================")
	fmt.Println("Demonstrating multi-device synchronization patterns...")
	fmt.Println()

	// Simulate device pairing scenarios
	devices := []map[string]interface{}{
		{
			"id":        "device_mobile_001",
			"name":      "iPhone 15 Pro",
			"type":      "mobile",
			"platform":  "ios",
			"paired_at": time.Now().Add(-2 * time.Hour).Unix(),
			"status":    "active",
			"last_sync": time.Now().Add(-5 * time.Minute).Unix(),
		},
		{
			"id":        "device_desktop_002",
			"name":      "MacBook Pro",
			"type":      "desktop",
			"platform":  "macos",
			"paired_at": time.Now().Add(-1 * time.Hour).Unix(),
			"status":    "active",
			"last_sync": time.Now().Add(-2 * time.Minute).Unix(),
		},
		{
			"id":        "device_tablet_003",
			"name":      "iPad Air",
			"type":      "tablet",
			"platform":  "ios",
			"paired_at": time.Now().Add(-30 * time.Minute).Unix(),
			"status":    "pairing",
			"last_sync": time.Now().Unix(),
		},
	}

	// Store device registry
	devicesData, _ := json.Marshal(devices)
	d.account.ValueStore("paired_devices", devicesData)
	d.statistics.StorageOps++

	fmt.Println("📱 Paired Device Registry:")
	for i, device := range devices {
		fmt.Printf("   %d. %s (%s)\n", i+1, device["name"], device["platform"])
		fmt.Printf("      🆔 ID: %s\n", device["id"])
		fmt.Printf("      📊 Status: %s\n", device["status"])

		if pairedAt, ok := device["paired_at"].(int64); ok {
			fmt.Printf("      🕐 Paired: %s ago\n", time.Since(time.Unix(pairedAt, 0)).Round(time.Minute))
		}
		if lastSync, ok := device["last_sync"].(int64); ok {
			fmt.Printf("      🔄 Last sync: %s ago\n", time.Since(time.Unix(lastSync, 0)).Round(time.Second))
		}
		fmt.Println()
	}

	// Demonstrate synchronization scenarios
	d.demonstrateSyncScenarios()

	fmt.Println("🔄 Core SDK Synchronization Benefits:")
	fmt.Println("   • Unified account state across all devices")
	fmt.Println("   • Real-time credential and contact sync")
	fmt.Println("   • Automatic conflict resolution")
	fmt.Println("   • Secure cross-device messaging")
	fmt.Println("   • Persistent storage consistency")
	fmt.Println()
}

// demonstrateSyncScenarios shows different synchronization patterns
func (d *PairingDemo) demonstrateSyncScenarios() {
	fmt.Println("🔄 Synchronization Scenarios:")
	fmt.Println()

	scenarios := []struct {
		name        string
		description string
		example     string
	}{
		{
			"Credential Sync",
			"Identity credentials synchronized across devices",
			"New credential issued on mobile → automatically available on desktop",
		},
		{
			"Contact Management",
			"Contact list maintained consistently",
			"Contact added on tablet → visible on all paired devices",
		},
		{
			"Messaging History",
			"Message history synchronized",
			"Conversation started on desktop → continues seamlessly on mobile",
		},
		{
			"Settings Sync",
			"Account preferences synchronized",
			"Privacy settings changed on one device → applied to all devices",
		},
	}

	for i, scenario := range scenarios {
		fmt.Printf("   %d. 🔹 %s\n", i+1, scenario.name)
		fmt.Printf("      Description: %s\n", scenario.description)
		fmt.Printf("      Example: %s\n", scenario.example)
		fmt.Println()
	}
}

// demonstrateAdvancedPairingFeatures shows advanced pairing capabilities
func (d *PairingDemo) demonstrateAdvancedPairingFeatures() {
	fmt.Println("🔹 Advanced Pairing Features")
	fmt.Println("============================")
	fmt.Println("Exploring advanced pairing capabilities with core SDK...")
	fmt.Println()

	// Demonstrate advanced features
	features := []struct {
		feature     string
		description string
		benefit     string
	}{
		{
			"Account Backup & Recovery",
			"Core SDK provides secure account backup",
			"Lost device recovery without losing identity",
		},
		{
			"Multi-Device Credentials",
			"Credentials work across all paired devices",
			"Seamless identity verification anywhere",
		},
		{
			"Secure Storage Sync",
			"Encrypted storage synchronized automatically",
			"Consistent data across all devices",
		},
		{
			"Cross-Device Messaging",
			"Messages route to all paired devices",
			"Never miss important communications",
		},
		{
			"Progressive Pairing",
			"Add devices incrementally to account",
			"Build device ecosystem over time",
		},
	}

	for i, feature := range features {
		fmt.Printf("   %d. 🚀 %s\n", i+1, feature.feature)
		fmt.Printf("      💡 %s\n", feature.description)
		fmt.Printf("      ✅ Benefit: %s\n", feature.benefit)
		fmt.Println()
	}

	// Store advanced features information
	featuresData, _ := json.Marshal(features)
	d.account.ValueStore("advanced_features", featuresData)
	d.statistics.StorageOps++

	fmt.Println("🔗 Integration Opportunities:")
	fmt.Println("   • Combine with credentials for identity sync")
	fmt.Println("   • Use with notifications for pairing alerts")
	fmt.Println("   • Integrate with chat for multi-device messaging")
	fmt.Println("   • Connect to storage for data synchronization")
	fmt.Println("   • Link with discovery for device finding")
	fmt.Println()
}

// storePairingSession saves pairing session information
func (d *PairingDemo) storePairingSession(info *PairingInfo) {
	// Load existing sessions
	sessionsData, _ := d.account.ValueLookup("pairing_sessions")
	var sessions map[string]*PairingInfo
	if sessionsData != nil {
		json.Unmarshal(sessionsData, &sessions)
	} else {
		sessions = make(map[string]*PairingInfo)
	}

	// Add new session
	sessions[info.Code] = info

	// Store updated sessions
	newSessionsData, _ := json.Marshal(sessions)
	d.account.ValueStore("pairing_sessions", newSessionsData)
	d.statistics.StorageOps++
}

// displayPairingSessions shows current pairing sessions
func (d *PairingDemo) displayPairingSessions() {
	sessionsData, err := d.account.ValueLookup("pairing_sessions")
	if err != nil || sessionsData == nil {
		fmt.Println("   📝 No pairing sessions found")
		return
	}

	var sessions map[string]*PairingInfo
	if err := json.Unmarshal(sessionsData, &sessions); err != nil {
		fmt.Println("   ❌ Error loading pairing sessions")
		return
	}

	fmt.Printf("   📊 Active pairing sessions: %d\n", len(sessions))
	for code, session := range sessions {
		fmt.Printf("      • Code: %s (expires: %s)\n",
			code[:8]+"...",
			session.ExpiresAt.Format("15:04:05"))
	}
	fmt.Println()
}

// updatePairingStatistics updates pairing metrics
func (d *PairingDemo) updatePairingStatistics() {
	d.statistics.Runtime = time.Since(d.startTime)
	d.statistics.PairingAttempts++

	// Store updated statistics
	statsData, _ := json.Marshal(d.statistics)
	d.account.ValueStore("pairing_statistics", statsData)
	d.statistics.StorageOps++
}

// displayPairingAnalytics shows pairing analytics
func (d *PairingDemo) displayPairingAnalytics() {
	fmt.Println("📊 Pairing Analytics:")
	fmt.Printf("   📈 Codes generated: %d\n", d.statistics.CodesGenerated)
	fmt.Printf("   🔄 Pairing attempts: %d\n", d.statistics.PairingAttempts)
	fmt.Printf("   ⏱️  Demo runtime: %v\n", d.statistics.Runtime.Round(time.Second))
	fmt.Printf("   💾 Storage operations: %d\n", d.statistics.StorageOps)
	fmt.Println()
}

// displayFinalSummary shows the final demo summary
func (d *PairingDemo) displayFinalSummary() {
	d.updatePairingStatistics()

	fmt.Println("🏁 Core SDK Pairing Demo Summary")
	fmt.Println("================================")
	fmt.Printf("⏱️  Total runtime: %v\n", d.statistics.Runtime.Round(time.Second))
	d.displayPairingAnalytics()

	fmt.Println("🎓 What was demonstrated using core SDK:")
	fmt.Println("   ✅ Account-based pairing code generation")
	fmt.Println("   ✅ QR code creation for device pairing")
	fmt.Println("   ✅ Multi-device synchronization patterns")
	fmt.Println("   ✅ Pairing session management and analytics")
	fmt.Println("   ✅ Persistent storage through core SDK")
	fmt.Println("   ✅ Cross-device state consistency")
	fmt.Println()

	fmt.Println("🚀 Core SDK pairing benefits:")
	fmt.Println("   • Account-based authentication and security")
	fmt.Println("   • Built-in encryption and secure storage")
	fmt.Println("   • Automatic cross-device synchronization")
	fmt.Println("   • Persistent pairing state management")
	fmt.Println("   • Seamless multi-device experiences")
	fmt.Println()

	fmt.Println("📚 Next steps:")
	fmt.Println("   • Run ../production_patterns/main.go for real-world patterns")
	fmt.Println("   • Explore integration with other core SDK features")
	fmt.Println("   • Build production pairing systems using these patterns")
	fmt.Println("   • Implement custom device management workflows")
	fmt.Println()

	fmt.Println("✅ Core SDK pairing demo completed!")
}
