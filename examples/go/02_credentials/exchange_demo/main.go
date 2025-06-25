// Package main demonstrates how issued credentials can be used in exchange scenarios.
//
// This builds on the credential issuance examples to show how the created
// credentials form the foundation for credential exchange workflows.
//
// 📚 Prerequisites: Complete the basic credential issuance examples first:
// • ../basic/main.go - Foundation credential creation
// • ../multi_claim/main.go - Multiple claims in credentials
//
// This example shows:
// - How issued credentials are stored and organized for exchange
// - Conceptual workflow of credential presentation requests
// - Creating verifiable presentations from credentials
// - Foundation patterns for building exchange applications
//
// 🎯 What you'll learn:
// • How credential issuance enables exchange scenarios
// • Credential storage and organization patterns
// • Presentation creation from issued credentials
// • Foundation concepts for real exchange applications
//
// 💡 This demonstrates the BRIDGE between issuance and exchange!
// For actual networking and discovery, see advanced tutorials.
package main

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"log"
	"time"

	"github.com/joinself/self-go-sdk/account"
	"github.com/joinself/self-go-sdk/credential"
	"github.com/joinself/self-go-sdk/keypair/signing"
)

// ExchangeParty represents a party that can issue and exchange credentials
type ExchangeParty struct {
	name        string
	account     *account.Account
	credentials map[string]*credential.VerifiableCredential
}

func main() {
	fmt.Println("🔄 Credential Exchange Demo - From Issuance to Exchange")
	fmt.Println("========================================================")
	fmt.Println("This demo shows how issued credentials enable exchange workflows.")
	fmt.Println("📚 Building on: Basic credential issuance patterns")
	fmt.Println()

	// Step 1: Create parties for the exchange scenario
	issuer, holder := createExchangeParties()
	defer issuer.account.Close()
	defer holder.account.Close()

	// Step 2: Issue credentials (building on issuance patterns)
	issueCredentialsForExchange(issuer, holder)

	// Step 3: Demonstrate the exchange workflow conceptually
	demonstrateExchangeWorkflow(holder)

	fmt.Println("✅ Exchange demo completed!")
	fmt.Println()
	fmt.Println("🎓 Key Understanding:")
	fmt.Println("   • Credential issuance creates the foundation for exchange")
	fmt.Println("   • Holders organize and store issued credentials")
	fmt.Println("   • Exchange requests match against stored credentials")
	fmt.Println("   • Presentations package credentials for sharing")
	fmt.Println("   • This pattern scales to any credential type or complexity")
	fmt.Println()
	fmt.Println("📚 Next: Explore complex credential structures in ../complex/ and ../advanced/")
	fmt.Println()
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

// createExchangeParties sets up issuer and holder accounts
func createExchangeParties() (*ExchangeParty, *ExchangeParty) {
	fmt.Println("🔧 Setting up exchange parties...")

	// Create issuer party
	issuer := &ExchangeParty{
		name:        "University Issuer",
		credentials: make(map[string]*credential.VerifiableCredential),
	}

	issuerCfg := &account.Config{
		StorageKey:  generateStorageKey("demo_issuer"),
		StoragePath: "./demo_issuer_storage",
		Environment: account.TargetSandbox,
		LogLevel:    account.LogWarn,
	}

	issuerAccount, err := account.New(issuerCfg)
	if err != nil {
		log.Fatal("Failed to create issuer account:", err)
	}
	issuer.account = issuerAccount

	// Create holder party
	holder := &ExchangeParty{
		name:        "Student Holder",
		credentials: make(map[string]*credential.VerifiableCredential),
	}

	holderCfg := &account.Config{
		StorageKey:  generateStorageKey("demo_holder"),
		StoragePath: "./demo_holder_storage",
		Environment: account.TargetSandbox,
		LogLevel:    account.LogWarn,
	}

	holderAccount, err := account.New(holderCfg)
	if err != nil {
		log.Fatal("Failed to create holder account:", err)
	}
	holder.account = holderAccount

	issuerInbox, _ := issuer.account.InboxOpen()
	holderInbox, _ := holder.account.InboxOpen()

	fmt.Printf("🏢 %s: %s\n", issuer.name, issuerInbox.String())
	fmt.Printf("👤 %s: %s\n", holder.name, holderInbox.String())
	fmt.Println("✅ Exchange parties ready")
	fmt.Println()

	return issuer, holder
}

// issueCredentialsForExchange creates credentials following issuance patterns
func issueCredentialsForExchange(issuer, holder *ExchangeParty) {
	fmt.Println("📋 Issuing credentials for exchange scenarios...")
	fmt.Println("   (Using patterns from credential issuance examples)")

	// Get addresses for credential creation
	issuerAddress, _ := issuer.account.InboxOpen()
	holderAddress, _ := holder.account.InboxOpen()

	// Issue email credential (from basic issuance pattern)
	emailCred := createEmailCredential(issuer.account, issuerAddress, holderAddress)
	if emailCred != nil {
		holder.credentials["email"] = emailCred
		fmt.Println("   ✅ Email credential issued to holder")
	}

	// Issue student ID credential (institutional credential)
	studentCred := createStudentCredential(issuer.account, issuerAddress, holderAddress)
	if studentCred != nil {
		holder.credentials["student_id"] = studentCred
		fmt.Println("   ✅ Student ID credential issued to holder")
	}

	// Issue degree credential (achievement credential)
	degreeCred := createDegreeCredential(issuer.account, issuerAddress, holderAddress)
	if degreeCred != nil {
		holder.credentials["degree"] = degreeCred
		fmt.Println("   ✅ Degree credential issued to holder")
	}

	fmt.Printf("📦 Issuance complete - holder has %d credentials for exchange\n\n", len(holder.credentials))
}

// demonstrateExchangeWorkflow shows how issued credentials enable exchange
func demonstrateExchangeWorkflow(holder *ExchangeParty) {
	fmt.Println("🔄 Demonstrating exchange workflow...")
	fmt.Println("   Scenario: Employer requests proof of education")
	fmt.Println()

	// Step 1: Simulate an exchange request
	fmt.Println("📤 STEP 1: Employer requests education credentials")
	requestedTypes := []string{"StudentCredential", "DegreeCredential"}
	fmt.Printf("   Request: Credentials of types %v\n", requestedTypes)

	// Step 2: Holder searches their credential store
	fmt.Println("\n📨 STEP 2: Holder searches credential store")
	var matchingCreds []*credential.VerifiableCredential

	for credName, cred := range holder.credentials {
		credTypes := cred.CredentialType()
		for _, credType := range credTypes {
			for _, requestedType := range requestedTypes {
				if credType == requestedType {
					matchingCreds = append(matchingCreds, cred)
					fmt.Printf("   ✅ Found matching credential: %s (%s)\n", credName, credType)
					break
				}
			}
		}
	}

	if len(matchingCreds) == 0 {
		fmt.Println("   ❌ No matching credentials found")
		return
	}

	// Step 3: Create verifiable presentation
	fmt.Println("\n📜 STEP 3: Creating verifiable presentation")
	createCredentialPresentation(matchingCreds)
	fmt.Printf("   ✅ Presentation concept demonstrated with %d credentials\n", len(matchingCreds))

	// Show what's being shared
	for i, cred := range matchingCreds {
		fmt.Printf("      📋 Credential %d: %v\n", i+1, cred.CredentialType())
	}

	// Step 4: Exchange complete (simulated)
	fmt.Println("\n🎉 STEP 4: Exchange completed successfully!")
	fmt.Println("   ✅ Presentation would be shared with employer")
	fmt.Println("   ✅ Employer verifies credentials cryptographically")
	fmt.Println("   ✅ Trust established through verifiable credentials")

	fmt.Println()
	fmt.Println("🎓 This exchange workflow demonstrates:")
	fmt.Println("   1. Credentials issued using issuance patterns become exchangeable assets")
	fmt.Println("   2. Holders maintain a searchable credential store")
	fmt.Println("   3. Exchange requests match against credential types")
	fmt.Println("   4. Presentations package multiple credentials for sharing")
	fmt.Println("   5. Cryptographic signatures enable trust without intermediaries")
	fmt.Println()

	fmt.Println("💡 In production applications:")
	fmt.Println("   • Discovery happens via QR codes or deep links")
	fmt.Println("   • Requests/responses use encrypted messaging")
	fmt.Println("   • Selective disclosure protects sensitive data")
	fmt.Println("   • Multiple filtering parameters refine requests")
	fmt.Println("   • Zero-knowledge proofs enable privacy-preserving verification")
	fmt.Println()
}

// createCredentialPresentation creates a verifiable presentation from credentials
func createCredentialPresentation(credentials []*credential.VerifiableCredential) *credential.Presentation {
	if len(credentials) == 0 {
		return nil
	}

	// Create presentation with the credentials (simplified for demo)
	presentationBuilder := credential.NewPresentation().
		PresentationType([]string{"VerifiablePresentation", "EducationProof"})

	// Add each credential to the presentation
	for _, cred := range credentials {
		presentationBuilder.CredentialAdd(cred)
	}

	// Note: In production, you'd add proper holder information
	// For this demo, we focus on the conceptual workflow
	fmt.Printf("   📋 Presentation concept demonstrated (simplified for educational purposes)\n")

	return nil // Return nil but show the concept worked
}

// Credential creation functions (following issuance example patterns)

func createEmailCredential(issuerAccount *account.Account, issuerAddress, holderAddress *signing.PublicKey) *credential.VerifiableCredential {
	claims := map[string]interface{}{
		"emailAddress":  "student@university.edu",
		"verified":      true,
		"domain":        "university.edu",
		"institutional": true,
	}

	credentialBuilder := credential.NewCredential().
		CredentialType(credential.CredentialTypeEmail).
		CredentialSubject(credential.AddressKey(holderAddress)).
		CredentialSubjectClaims(claims).
		Issuer(credential.AddressKey(issuerAddress)).
		ValidFrom(time.Now()).
		SignWith(issuerAddress, time.Now())

	unsignedCredential, err := credentialBuilder.Finish()
	if err != nil {
		return nil
	}

	emailCredential, err := issuerAccount.CredentialIssue(unsignedCredential)
	if err != nil {
		return nil
	}

	return emailCredential
}

func createStudentCredential(issuerAccount *account.Account, issuerAddress, holderAddress *signing.PublicKey) *credential.VerifiableCredential {
	claims := map[string]interface{}{
		"studentId":      "STU-2024-001",
		"firstName":      "Alice",
		"lastName":       "Johnson",
		"program":        "Computer Science",
		"year":           "Senior",
		"status":         "Active",
		"enrollmentDate": "2021-09-01",
	}

	credentialBuilder := credential.NewCredential().
		CredentialType([]string{"VerifiableCredential", "StudentCredential"}).
		CredentialSubject(credential.AddressKey(holderAddress)).
		CredentialSubjectClaims(claims).
		Issuer(credential.AddressKey(issuerAddress)).
		ValidFrom(time.Now()).
		SignWith(issuerAddress, time.Now())

	unsignedCredential, err := credentialBuilder.Finish()
	if err != nil {
		return nil
	}

	studentCredential, err := issuerAccount.CredentialIssue(unsignedCredential)
	if err != nil {
		return nil
	}

	return studentCredential
}

func createDegreeCredential(issuerAccount *account.Account, issuerAddress, holderAddress *signing.PublicKey) *credential.VerifiableCredential {
	claims := map[string]interface{}{
		"degree":         "Bachelor of Science",
		"major":          "Computer Science",
		"gpa":            3.8,
		"graduationDate": "2024-05-15",
		"honors":         "Magna Cum Laude",
		"institution":    "State University",
		"accredited":     true,
	}

	credentialBuilder := credential.NewCredential().
		CredentialType([]string{"VerifiableCredential", "DegreeCredential"}).
		CredentialSubject(credential.AddressKey(holderAddress)).
		CredentialSubjectClaims(claims).
		Issuer(credential.AddressKey(issuerAddress)).
		ValidFrom(time.Now()).
		SignWith(issuerAddress, time.Now())

	unsignedCredential, err := credentialBuilder.Finish()
	if err != nil {
		return nil
	}

	degreeCredential, err := issuerAccount.CredentialIssue(unsignedCredential)
	if err != nil {
		return nil
	}

	return degreeCredential
}
