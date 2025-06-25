// Package main demonstrates credential issuance with evidence using the Self SDK.
//
// This is the EVIDENCE level of credential issuance examples.
// Prerequisites: Complete basic/main.go and multi_claim/main.go first.
//
// This example shows:
// - Creating custom credential types
// - Attaching file evidence to credentials
// - Asset management and upload functionality
// - Creating verifiable presentations
// - Linking evidence to credential claims
//
// 🎯 What you'll learn:
// • How to attach evidence files to credentials
// • Asset management and secure storage
// • Creating verifiable presentations
// • Linking evidence to claims with hashes
// • Custom credential types
//
// 📚 Next steps:
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
	"github.com/joinself/self-go-sdk/object"
)

func main() {
	fmt.Println("🎓 Evidence-Based Credential Issuance Demo (Core SDK)")
	fmt.Println("======================================================")
	fmt.Println("This demo shows credentials with file evidence using the core SDK.")
	fmt.Println()

	// Step 1: Create issuer and holder accounts
	issuer, holder := createAccounts()
	defer issuer.Close()
	defer holder.Close()

	// Display account information
	displayAccountInfo(issuer, holder)

	// Step 2: Create credentials with evidence
	createCertificationWithEvidence(issuer, holder)

	fmt.Println("✅ Evidence demo completed!")
	fmt.Println()
	fmt.Println("📚 Ready for the next level?")
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
		StorageKey:  generateStorageKey("evidence_issuer"),
		StoragePath: "./evidence_issuer_storage",
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
		StorageKey:  generateStorageKey("evidence_holder"),
		StoragePath: "./evidence_holder_storage",
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

// createCertificationWithEvidence creates a certification credential with file evidence
func createCertificationWithEvidence(issuer, holder *account.Account) {
	fmt.Println("🎓 Creating certification credential with evidence...")

	// Step 1: Create evidence asset
	evidence := createEvidence(issuer)
	if evidence == nil {
		return
	}

	// Step 2: Create credential with evidence reference
	cred := createCredentialWithEvidence(issuer, holder, evidence)
	if cred == nil {
		return
	}

	// Step 3: Create verifiable presentation
	createPresentation(issuer, cred)
}

// createEvidence creates and uploads supporting documentation
func createEvidence(issuer *account.Account) *object.Object {
	fmt.Println("📄 Creating evidence asset...")

	// Create mock certificate document
	certificateData := []byte(`Certificate of Completion
Advanced Go Programming Course

Student: John Doe
Course: Advanced Go Programming with Self SDK
Institution: Self SDK Academy
Grade: A+
Credits: 40 hours
Date: ` + time.Now().Format("2006-01-02") + `

This certificate verifies that the above-named student has
successfully completed the Advanced Go Programming course
with distinction.

Instructor: Jane Smith
Director: Dr. Alice Johnson`)

	// Create object using the core SDK
	evidence, err := object.New("certificate.pdf", certificateData)
	if err != nil {
		log.Printf("Failed to create evidence object: %v", err)
		return nil
	}

	// Upload evidence using the issuer's account
	err = issuer.ObjectUpload(evidence, true)
	if err != nil {
		log.Printf("Failed to upload evidence: %v", err)
		return nil
	}

	fmt.Printf("   📄 Evidence created: certificate.pdf\n")
	fmt.Printf("   🔗 Asset ID: %x\n", evidence.Id())
	fmt.Printf("   📏 Size: %d bytes\n", len(certificateData))
	fmt.Printf("   ✅ Evidence uploaded to secure storage\n")
	fmt.Println()

	return evidence
}

// createCredentialWithEvidence creates a custom credential with evidence reference
func createCredentialWithEvidence(issuer, holder *account.Account, evidence *object.Object) *credential.VerifiableCredential {
	fmt.Println("🏗️ Building custom certification credential...")

	// Get inbox addresses for credential creation
	issuerAddress, err := issuer.InboxOpen()
	if err != nil {
		log.Fatal("Failed to open issuer inbox:", err)
	}

	holderAddress, err := holder.InboxOpen()
	if err != nil {
		log.Fatal("Failed to open holder inbox:", err)
	}

	// Create claims with evidence reference
	claims := map[string]interface{}{
		"certificationType": "Professional Development",
		"courseName":        "Advanced Go Programming",
		"completionDate":    time.Now().Format("2006-01-02"),
		"grade":             "A+",
		"institution":       "Self SDK Academy",
		"courseHours":       40,
		"instructor":        "Jane Smith",
		"evidenceId":        fmt.Sprintf("%x", evidence.Id()),
	}

	// Build the credential using the core SDK
	credentialBuilder := credential.NewCredential().
		CredentialType([]string{"VerifiableCredential", "CertificationCredential"}).
		CredentialSubject(credential.AddressKey(holderAddress)).
		CredentialSubjectClaims(claims).
		Issuer(credential.AddressKey(issuerAddress)).
		ValidFrom(time.Now()).
		SignWith(issuerAddress, time.Now())

	// Finish building the unsigned credential
	unsignedCredential, err := credentialBuilder.Finish()
	if err != nil {
		log.Printf("Failed to build credential: %v", err)
		return nil
	}

	// Issue the credential using the issuer's account
	customCredential, err := issuer.CredentialIssue(unsignedCredential)
	if err != nil {
		log.Printf("Failed to issue credential: %v", err)
		return nil
	}

	fmt.Printf("   ✅ Certification credential created successfully\n")
	fmt.Printf("   🎓 Course: Advanced Go Programming\n")
	fmt.Printf("   📅 Completed: %s\n", time.Now().Format("2006-01-02"))
	fmt.Printf("   🏆 Grade: A+\n")
	fmt.Printf("   🏫 Institution: Self SDK Academy\n")
	fmt.Printf("   ⏱️  Duration: 40 hours\n")
	fmt.Printf("   🔒 Type: %v\n", customCredential.CredentialType())
	fmt.Printf("   🔗 Evidence ID: %x\n", evidence.Id())
	fmt.Println()

	return customCredential
}

// createPresentation creates a verifiable presentation from the credential
func createPresentation(issuer *account.Account, cred *credential.VerifiableCredential) {
	fmt.Println("📋 Creating verifiable presentation...")

	// Get issuer address for presentation creation
	issuerAddress, err := issuer.InboxOpen()
	if err != nil {
		log.Fatal("Failed to open issuer inbox:", err)
	}

	// Create presentation using the core SDK
	presentationBuilder := credential.NewPresentation().
		PresentationType([]string{"VerifiablePresentation", "CertificationPresentation"}).
		Holder(credential.AddressKey(issuerAddress)).
		CredentialAdd(cred)

	// Finish building the unsigned presentation
	unsignedPresentation, err := presentationBuilder.Finish()
	if err != nil {
		log.Printf("Failed to build presentation: %v", err)
		return
	}

	// Issue the presentation using the issuer's account
	presentation, err := issuer.PresentationIssue(unsignedPresentation)
	if err != nil {
		log.Printf("Failed to issue presentation: %v", err)
		return
	}

	fmt.Printf("   ✅ Presentation created successfully\n")
	fmt.Printf("   📋 Type: %v\n", presentation.PresentationType())
	fmt.Printf("   📄 Credentials included: %d\n", len(presentation.Credentials()))
	fmt.Println()
	fmt.Println("🎓 What happened:")
	fmt.Println("   1. Created evidence asset (PDF certificate)")
	fmt.Println("   2. Uploaded evidence to secure storage")
	fmt.Println("   3. Created credential with evidence reference")
	fmt.Println("   4. Created verifiable presentation for sharing")
	fmt.Println()
}
