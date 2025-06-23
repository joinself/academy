// Package main demonstrates multi-claim credential issuance using the Self SDK.
//
// This is the MULTI-CLAIM level of credential issuance examples.
// Prerequisites: Complete basic/main.go first.
//
// This example shows:
// - Creating credentials with multiple claims
// - Different data types in claims
// - Organizing related information in one credential
// - Building upon basic credential concepts
//
// 🎯 What you'll learn:
// • How to add multiple claims to a single credential
// • Different data types in claims (strings, booleans, numbers)
// • Organizing related identity information
// • Efficient credential structuring
//
// 📚 Next steps:
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
	fmt.Println("🎓 Multi-Claim Credential Issuance Demo (Core SDK)")
	fmt.Println("===================================================")
	fmt.Println("This demo shows credentials with multiple claims using the core SDK.")
	fmt.Println()

	// Step 1: Create issuer and holder accounts
	issuer, holder := createAccounts()
	defer issuer.Close()
	defer holder.Close()

	// Display account information
	displayAccountInfo(issuer, holder)

	// Step 2: Create credentials with multiple claims
	createProfileCredential(issuer, holder)
	createEducationCredential(issuer, holder)

	fmt.Println("✅ Multi-claim demo completed!")
	fmt.Println()
	fmt.Println("📚 Ready for the next level?")
	fmt.Println("   • cd ../evidence && go run main.go - Evidence and asset management")
	fmt.Println("   • cd ../complex && go run main.go - Complex nested data structures")
	fmt.Println("   • cd ../advanced && go run main.go - All features combined")
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
	// Create issuer account configuration
	issuerCfg := &account.Config{
		StorageKey:  generateStorageKey("multi_issuer"),
		StoragePath: "./multi_issuer_storage",
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
		StorageKey:  generateStorageKey("multi_holder"),
		StoragePath: "./multi_holder_storage",
		Environment: account.TargetSandbox,
		LogLevel:    account.LogWarn,
	}

	// Create holder account
	holder, err := account.New(holderCfg)
	if err != nil {
		log.Fatal("Failed to create holder account:", err)
	}

	return issuer, holder
}

// displayAccountInfo shows the account information
func displayAccountInfo(issuer, holder *account.Account) {
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
}

// createProfileCredential creates a profile credential with multiple claims
func createProfileCredential(issuer, holder *account.Account) {
	fmt.Println("👤 Creating profile credential with multiple claims...")

	// Get inbox addresses for credential creation
	issuerAddress, err := issuer.InboxOpen()
	if err != nil {
		log.Fatal("Failed to open issuer inbox:", err)
	}

	holderAddress, err := holder.InboxOpen()
	if err != nil {
		log.Fatal("Failed to open holder inbox:", err)
	}

	// Create profile claims with multiple data types
	claims := map[string]interface{}{
		"firstName":        "John",
		"lastName":         "Doe",
		"displayName":      "John Doe",
		"profileLevel":     "verified",
		"country":          "United States",
		"age":              30,
		"isActive":         true,
		"registrationDate": time.Now().Format("2006-01-02"),
	}

	// Build the credential using the core SDK
	credentialBuilder := credential.NewCredential().
		CredentialType(credential.CredentialTypeProfileName).
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
	profileCredential, err := issuer.CredentialIssue(unsignedCredential)
	if err != nil {
		log.Printf("Failed to issue credential: %v", err)
		return
	}

	fmt.Printf("✅ Profile credential created successfully\n")
	fmt.Printf("   👤 Name: John Doe\n")
	fmt.Printf("   🌍 Country: United States\n")
	fmt.Printf("   🎂 Age: 30\n")
	fmt.Printf("   ⭐ Level: verified\n")
	fmt.Printf("   ✅ Active: true\n")
	fmt.Printf("   📅 Registration: %s\n", time.Now().Format("2006-01-02"))
	fmt.Printf("   🔒 Type: %v\n", profileCredential.CredentialType())
	fmt.Println()
}

// createEducationCredential creates an education credential with academic claims
func createEducationCredential(issuer, holder *account.Account) {
	fmt.Println("🎓 Creating education credential with academic claims...")

	// Get inbox addresses for credential creation
	issuerAddress, err := issuer.InboxOpen()
	if err != nil {
		log.Fatal("Failed to open issuer inbox:", err)
	}

	holderAddress, err := holder.InboxOpen()
	if err != nil {
		log.Fatal("Failed to open holder inbox:", err)
	}

	// Create education claims with academic information
	claims := map[string]interface{}{
		"institution":      "University of Technology",
		"degree":           "Bachelor of Science",
		"major":            "Computer Science",
		"graduationYear":   2020,
		"gpa":              3.8,
		"honors":           true,
		"creditsCompleted": 120,
		"thesis":           "Machine Learning Applications",
		"graduationDate":   "2020-05-15",
	}

	// Build the credential using the core SDK
	credentialBuilder := credential.NewCredential().
		CredentialType([]string{"VerifiableCredential", "EducationCredential"}).
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
	educationCredential, err := issuer.CredentialIssue(unsignedCredential)
	if err != nil {
		log.Printf("Failed to issue credential: %v", err)
		return
	}

	fmt.Printf("✅ Education credential created successfully\n")
	fmt.Printf("   🏫 Institution: University of Technology\n")
	fmt.Printf("   🎓 Degree: Bachelor of Science in Computer Science\n")
	fmt.Printf("   📅 Graduated: 2020-05-15\n")
	fmt.Printf("   📊 GPA: 3.8\n")
	fmt.Printf("   🏆 Honors: true\n")
	fmt.Printf("   📚 Credits: 120\n")
	fmt.Printf("   📝 Thesis: Machine Learning Applications\n")
	fmt.Printf("   🔒 Type: %v\n", educationCredential.CredentialType())
	fmt.Println()
	fmt.Println("🎓 What happened:")
	fmt.Println("   1. Created two credentials with multiple claims each")
	fmt.Println("   2. Used different data types: strings, numbers, booleans")
	fmt.Println("   3. Grouped related information in single credentials")
	fmt.Println("   4. Each credential maintains cryptographic integrity")
	fmt.Println()
}
