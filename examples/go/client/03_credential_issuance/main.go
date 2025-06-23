// Package main demonstrates simple credential issuance using the Self SDK.
//
// This is a simplified version of the credential issuance tutorial.
// For the complete educational progression, see the individual tutorial files:
//
// 📚 Educational Progression:
// 1. basic/main.go - Foundation concepts (start here)
// 2. multi_claim/main.go - Multiple claims in credentials
// 3. evidence/main.go - Evidence and asset management
// 4. complex/main.go - Complex nested data structures
// 5. advanced/main.go - All features combined
//
// This example shows the basics of:
// - Setting up issuer and holder accounts using the core SDK
// - Creating a simple credential using the underlying SDK
// - Understanding the issuance workflow without the client facade
// - Direct credential builder usage
//
// 🎯 What you'll learn:
// • How credential issuance works at the SDK level
// • Direct account management and configuration
// • Core credential creation patterns
// • Direct credential signing and issuance
package main

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"log"
	"time"

	"github.com/joinself/self-go-sdk/account"
	"github.com/joinself/self-go-sdk/credential"
)

func main() {
	fmt.Println("🎓 Simple Credential Issuance Demo (Core SDK)")
	fmt.Println("==============================================")
	fmt.Println("This demo shows basic credential issuance using the underlying Self SDK.")
	fmt.Println()

	// Step 1: Create issuer and holder accounts
	issuer, holder := createAccounts()
	defer issuer.Close()
	defer holder.Close()

	// Get DIDs from accounts
	issuerInbox, err := issuer.InboxOpen()
	if err != nil {
		log.Fatal("Failed to open issuer inbox:", err)
	}
	holderInbox, err := holder.InboxOpen()
	if err != nil {
		log.Fatal("Failed to open holder inbox:", err)
	}

	fmt.Printf("🏢 Issuer: %s\n", issuerInbox.String())
	fmt.Printf("👤 Holder: %s\n", holderInbox.String())
	fmt.Println()

	// Step 2: Create a simple credential
	createSimpleCredential(issuer, holder)

	fmt.Println("✅ Basic demo completed!")
	fmt.Println()
	fmt.Println("📚 Ready for the next level?")
	fmt.Println("   • cd basic && go run main.go - Foundation concepts (start here)")
	fmt.Println("   • cd multi_claim && go run main.go - Multiple claims in credentials")
	fmt.Println("   • cd evidence && go run main.go - Evidence and asset management")
	fmt.Println("   • cd complex && go run main.go - Complex nested data structures")
	fmt.Println("   • cd advanced && go run main.go - All features combined")
	fmt.Println()
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

// createAccounts sets up the issuer and holder accounts using the core SDK
func createAccounts() (*account.Account, *account.Account) {
	fmt.Println("🔧 Setting up accounts using core SDK...")

	// Create issuer account configuration
	issuerCfg := &account.Config{
		StorageKey:  generateStorageKey("simple_issuer"),
		StoragePath: "./simple_issuer_storage",
		Environment: account.TargetSandbox,
		LogLevel:    account.LogWarn,
	}

	// Create issuer account
	issuer, err := account.New(issuerCfg)
	if err != nil {
		log.Fatal("Failed to create issuer account:", err)
	}

	// Create holder account configuration
	holderCfg := &account.Config{
		StorageKey:  generateStorageKey("simple_holder"),
		StoragePath: "./simple_holder_storage",
		Environment: account.TargetSandbox,
		LogLevel:    account.LogWarn,
	}

	// Create holder account
	holder, err := account.New(holderCfg)
	if err != nil {
		log.Fatal("Failed to create holder account:", err)
	}

	fmt.Println("✅ Accounts created successfully")
	return issuer, holder
}

// createSimpleCredential creates a basic email credential using the core SDK
func createSimpleCredential(issuer, holder *account.Account) {
	fmt.Println("📧 Creating simple email credential using core SDK...")

	// Get inbox addresses for credential creation
	issuerAddress, err := issuer.InboxOpen()
	if err != nil {
		log.Fatal("Failed to open issuer inbox:", err)
	}

	holderAddress, err := holder.InboxOpen()
	if err != nil {
		log.Fatal("Failed to open holder inbox:", err)
	}

	// Create credential claims
	claims := map[string]interface{}{
		"emailAddress": "demo@example.com",
		"verified":     true,
	}

	// Build the credential using the core SDK
	credentialBuilder := credential.NewCredential().
		CredentialType(credential.CredentialTypeEmail).
		CredentialSubject(credential.AddressKey(holderAddress)).
		CredentialSubjectClaims(claims).
		Issuer(credential.AddressKey(issuerAddress)).
		ValidFrom(time.Now()).
		SignWith(issuerAddress, time.Now())

	// Finish building the unsigned credential
	unsignedCredential, err := credentialBuilder.Finish()
	if err != nil {
		log.Printf("Failed to build credential: %v", err)
		return
	}

	// Issue the credential using the issuer's account
	emailCredential, err := issuer.CredentialIssue(unsignedCredential)
	if err != nil {
		log.Printf("Failed to issue credential: %v", err)
		return
	}

	fmt.Printf("✅ Email credential created successfully\n")
	fmt.Printf("   📧 Email: demo@example.com\n")
	fmt.Printf("   ✔️  Verified: true\n")
	fmt.Printf("   🔒 Type: %v\n", emailCredential.CredentialType())
	fmt.Printf("   🆔 Subject: %s\n", emailCredential.CredentialSubject().String())
	fmt.Println()
	fmt.Println("🎓 What happened:")
	fmt.Println("   1. Created issuer and holder accounts using core SDK")
	fmt.Println("   2. Built credential with claims using credential.NewCredential()")
	fmt.Println("   3. Signed credential with issuer's cryptographic key")
	fmt.Println("   4. Issued credential through account.CredentialIssue()")
	fmt.Println("   5. Credential is now ready for sharing")
	fmt.Println()
}
