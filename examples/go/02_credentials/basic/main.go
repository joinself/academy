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
	"fmt"
	"log"
	"time"

	"github.com/joinself/academy/examples/go/common"
	"github.com/joinself/self-go-sdk/account"
	"github.com/joinself/self-go-sdk/credential"
)

func main() {
	fmt.Println("🎓 Basic Credential Issuance Demo (Core SDK)")
	fmt.Println("=============================================")
	fmt.Println()

	// Step 1: Create issuer and holder accounts using centralized setup
	issuer, holder := common.SetupIssuerHolder(
		account.Callbacks{}, // issuer callbacks
		account.Callbacks{}, // holder callbacks
	)
	defer issuer.Close()
	defer holder.Close()

	// Step 2: Display account information
	common.DisplayAccountPair(issuer, holder)

	// Step 3: Create and issue a credential
	createEmailCredential(issuer, holder)

	fmt.Println("✅ Demo completed!")
}

// createEmailCredential creates a basic email credential using the core SDK
func createEmailCredential(issuer, holder *account.Account) {
	fmt.Println("📧 Creating email credential using core SDK...")

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
		"emailAddress": "alice@example.com",
		"verified":     true,
		"domain":       "example.com",
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

	// Display credential information
	fmt.Printf("✅ Email credential created successfully\n")
	fmt.Printf("   📧 Email: alice@example.com\n")
	fmt.Printf("   ✔️  Verified: true\n")
	fmt.Printf("   🏷️  Domain: example.com\n")
	fmt.Printf("   🔒 Type: %v\n", emailCredential.CredentialType())
	fmt.Printf("   🆔 Subject: %s\n", emailCredential.CredentialSubject().String())
	fmt.Printf("   🏢 Issuer: %s\n", emailCredential.Issuer().String())
	fmt.Println()

	fmt.Println("🎓 What happened:")
	fmt.Println("   1. Created issuer and holder accounts using common setup")
	fmt.Println("   2. Built credential with claims using credential.NewCredential()")
	fmt.Println("   3. Signed credential with issuer's cryptographic key")
	fmt.Println("   4. Issued credential through account.CredentialIssue()")
	fmt.Println("   5. Credential is now ready for verification or sharing")
	fmt.Println()

	fmt.Println("📚 Ready for the next level?")
	fmt.Println("   • cd ../multi_claim && go run main.go - Multiple claims in credentials")
	fmt.Println("   • cd ../evidence && go run main.go - Evidence and asset management")
	fmt.Println("   • cd ../complex && go run main.go - Complex nested data structures")
	fmt.Println("   • cd ../advanced && go run main.go - All features combined")
}
