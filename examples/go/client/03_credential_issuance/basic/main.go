// Package main demonstrates basic credential issuance using the underlying Self SDK directly.
//
// This is the BASIC level of credential issuance examples.
// Start here if you're new to credential issuance concepts.
//
// This example shows the basics of:
// - Setting up issuer and holder accounts using the core SDK
// - Creating a simple email credential using credential.NewCredential()
// - Understanding the core SDK credential builder pattern
// - Direct credential signing and issuance with account.CredentialIssue()
//
// 🎯 What you'll learn:
// • How credential issuance works at the SDK level
// • Direct account setup and configuration
// • Core credential creation patterns
// • Direct credential signing and issuance
// • Foundation concepts for all credential types
//
// 📚 Next steps:
// • multi_claim/main.go - Multiple claims in credentials
// • evidence/main.go - Evidence and asset management
// • complex/main.go - Complex nested data structures
// • advanced/main.go - All features combined
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
	fmt.Println("🎓 Basic Credential Issuance Demo (Core SDK)")
	fmt.Println("=============================================")
	fmt.Println()

	// Step 1: Create issuer and holder accounts using core SDK
	issuer, holder := createAccounts()
	defer issuer.Close()
	defer holder.Close()

	// Step 2: Display account information
	displayAccountInfo(issuer, holder)

	// Step 3: Create and issue a credential
	createEmailCredential(issuer, holder)

	fmt.Println("✅ Demo completed!")
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
	fmt.Println("🔧 Setting up accounts...")

	// Create issuer account with core SDK
	issuer, err := account.New(&account.Config{
		StorageKey:  generateStorageKey("basic_issuer"),
		StoragePath: "./basic_issuer_storage",
		Environment: account.TargetSandbox,
		LogLevel:    account.LogWarn,
	})
	if err != nil {
		log.Fatal("Failed to create issuer account:", err)
	}

	// Create holder account with core SDK
	holder, err := account.New(&account.Config{
		StorageKey:  generateStorageKey("basic_holder"),
		StoragePath: "./basic_holder_storage",
		Environment: account.TargetSandbox,
		LogLevel:    account.LogWarn,
	})
	if err != nil {
		log.Fatal("Failed to create holder account:", err)
	}

	fmt.Println("✅ Accounts created successfully")
	return issuer, holder
}

// displayAccountInfo shows the account DIDs
func displayAccountInfo(issuer, holder *account.Account) {
	// Get DIDs from accounts by opening inboxes
	issuerInbox, err := issuer.InboxOpen()
	if err != nil {
		log.Fatal("Failed to open issuer inbox:", err)
	}
	holderInbox, err := holder.InboxOpen()
	if err != nil {
		log.Fatal("Failed to open holder inbox:", err)
	}

	fmt.Printf("🏢 Issuer DID: %s\n", issuerInbox.String())
	fmt.Printf("👤 Holder DID: %s\n", holderInbox.String())
	fmt.Println()
}

// createEmailCredential creates a basic email credential using the core SDK
func createEmailCredential(issuer, holder *account.Account) {
	fmt.Println("📧 Creating email credential...")

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
		"emailAddress":     "john.doe@example.com",
		"verified":         true,
		"verificationDate": time.Now().Format("2006-01-02"),
	}

	// Build credential using core SDK
	credentialBuilder := credential.NewCredential().
		CredentialType(credential.CredentialTypeEmail).
		CredentialSubject(credential.AddressKey(holderAddress)).
		Issuer(credential.AddressKey(issuerAddress)).
		CredentialSubjectClaims(claims).
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

	// Display results
	fmt.Printf("✅ Email credential issued successfully\n")
	fmt.Printf("   📧 Email: john.doe@example.com\n")
	fmt.Printf("   ✔️  Verified: true\n")
	fmt.Printf("   📅 Date: %s\n", time.Now().Format("2006-01-02"))
	fmt.Printf("   🔒 Type: %v\n", emailCredential.CredentialType())
	fmt.Println()
}
