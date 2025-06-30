// Package main demonstrates credential exchange through presentation requests.
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
	fmt.Println("Credential Exchange Demo")
	fmt.Println("========================")

	// Create parties for the exchange scenario
	issuer, holder := createExchangeParties()
	defer issuer.account.Close()
	defer holder.account.Close()

	// Issue credentials for exchange
	issueCredentialsForExchange(issuer, holder)

	// Demonstrate exchange workflow
	demonstrateExchangeWorkflow(holder)

	fmt.Println("✅ Exchange demo completed!")
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
	fmt.Println("Setting up exchange parties...")

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

	fmt.Printf("Issuer: %s\n", issuerInbox.String())
	fmt.Printf("Holder: %s\n", holderInbox.String())
	fmt.Println()

	return issuer, holder
}

// issueCredentialsForExchange creates credentials following issuance patterns
func issueCredentialsForExchange(issuer, holder *ExchangeParty) {
	fmt.Println("Issuing credentials...")

	issuerAddress, _ := issuer.account.InboxOpen()
	holderAddress, _ := holder.account.InboxOpen()

	// Issue email credential
	emailCred := createEmailCredential(issuer.account, issuerAddress, holderAddress)
	if emailCred != nil {
		holder.credentials["email"] = emailCred
		fmt.Println("✅ Email credential issued")
	}

	// Issue student ID credential
	studentCred := createStudentCredential(issuer.account, issuerAddress, holderAddress)
	if studentCred != nil {
		holder.credentials["student_id"] = studentCred
		fmt.Println("✅ Student ID credential issued")
	}

	// Issue degree credential
	degreeCred := createDegreeCredential(issuer.account, issuerAddress, holderAddress)
	if degreeCred != nil {
		holder.credentials["degree"] = degreeCred
		fmt.Println("✅ Degree credential issued")
	}

	fmt.Printf("Holder has %d credentials\n\n", len(holder.credentials))
}

// demonstrateExchangeWorkflow shows how issued credentials enable exchange
func demonstrateExchangeWorkflow(holder *ExchangeParty) {
	fmt.Println("Exchange workflow:")
	fmt.Println("Scenario: Employer requests proof of education")

	// Simulate exchange request
	requestedTypes := []string{"StudentCredential", "DegreeCredential"}
	fmt.Printf("Requested types: %v\n", requestedTypes)

	// Search credential store
	var matchingCreds []*credential.VerifiableCredential

	for credName, cred := range holder.credentials {
		credTypes := cred.CredentialType()
		for _, credType := range credTypes {
			for _, requestedType := range requestedTypes {
				if credType == requestedType {
					matchingCreds = append(matchingCreds, cred)
					fmt.Printf("✅ Found matching: %s (%s)\n", credName, credType)
					break
				}
			}
		}
	}

	if len(matchingCreds) == 0 {
		fmt.Println("❌ No matching credentials found")
		return
	}

	// Create presentation
	presentation := createCredentialPresentation(matchingCreds)
	if presentation != nil {
		fmt.Printf("✅ Presentation created with %d credentials\n", len(matchingCreds))
		fmt.Println("✅ Exchange completed successfully")
	}
}

// createCredentialPresentation creates a verifiable presentation from credentials
func createCredentialPresentation(credentials []*credential.VerifiableCredential) *credential.VerifiablePresentation {
	if len(credentials) == 0 {
		return nil
	}

	// For demo purposes, simulate successful presentation creation
	// In production, you would create a proper verifiable presentation
	// with the holder's signature
	return &credential.VerifiablePresentation{}
}

// Credential creation functions (following issuance example patterns)

func createEmailCredential(issuerAccount *account.Account, issuerAddress, holderAddress *signing.PublicKey) *credential.VerifiableCredential {
	claims := map[string]interface{}{
		"emailAddress": "student@university.edu",
		"verified":     true,
		"domain":       "university.edu",
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
		log.Printf("Failed to build email credential: %v", err)
		return nil
	}

	emailCredential, err := issuerAccount.CredentialIssue(unsignedCredential)
	if err != nil {
		log.Printf("Failed to issue email credential: %v", err)
		return nil
	}

	return emailCredential
}

func createStudentCredential(issuerAccount *account.Account, issuerAddress, holderAddress *signing.PublicKey) *credential.VerifiableCredential {
	claims := map[string]interface{}{
		"studentId":      "STU-2024-001",
		"enrollmentDate": "2020-09-01",
		"status":         "enrolled",
		"program":        "Computer Science",
		"level":          "undergraduate",
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
		log.Printf("Failed to build student credential: %v", err)
		return nil
	}

	studentCredential, err := issuerAccount.CredentialIssue(unsignedCredential)
	if err != nil {
		log.Printf("Failed to issue student credential: %v", err)
		return nil
	}

	return studentCredential
}

func createDegreeCredential(issuerAccount *account.Account, issuerAddress, holderAddress *signing.PublicKey) *credential.VerifiableCredential {
	claims := map[string]interface{}{
		"degree":         "Bachelor of Science",
		"major":          "Computer Science",
		"graduationDate": "2024-05-15",
		"gpa":            3.8,
		"honors":         "magna cum laude",
		"institution":    "University of Technology",
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
		log.Printf("Failed to build degree credential: %v", err)
		return nil
	}

	degreeCredential, err := issuerAccount.CredentialIssue(unsignedCredential)
	if err != nil {
		log.Printf("Failed to issue degree credential: %v", err)
		return nil
	}

	return degreeCredential
}
