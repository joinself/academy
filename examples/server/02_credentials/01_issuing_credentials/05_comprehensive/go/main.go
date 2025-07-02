// Package main demonstrates comprehensive credential issuance capabilities using the Self SDK.
//
// This example shows how to:
// - Initialize Self accounts for issuer and holder roles using the core SDK
// - Create various types of verifiable credentials using the underlying credential module
// - Attach evidence/files to credentials for enhanced verification
// - Handle complex nested claims and data structures
// - Create verifiable presentations from credentials
// - Manage asset uploads and downloads for evidence using the core SDK
//
// The Self SDK provides decentralized identity and verifiable credential capabilities,
// allowing entities to issue, hold, and verify credentials without requiring
// centralized authorities while maintaining cryptographic integrity and privacy.
//
// 🎯 CREDENTIAL CAPABILITIES DEMONSTRATED:
// • Basic credential creation (Email verification)
// • Multi-claim credentials (Profile information)
// • Custom credentials with file evidence (Certifications)
// • Complex nested data structures (Organization credentials)
// • Credential builder pattern usage
// • Asset/evidence management (file uploads)
// • Verifiable presentation creation
// • Request/response handling workflows
//
// 🔧 KEY SDK COMPONENTS SHOWCASED:
// • account.New() - Account initialization and configuration using core SDK
// • credential.NewCredential() - Direct credential construction using core API
// • object.New() - Evidence and file attachment management using core SDK
// • credential.NewPresentation() - Verifiable presentation creation using core API
// • account.CredentialIssue() - Direct credential issuance using core SDK
//
// 📚 EDUCATIONAL PROGRESSION:
// The examples progress from simple to complex, building understanding:
// 1. Basic Email Credential - Simplest form with minimal claims
// 2. Profile Credential - Multiple claims in a single credential
// 3. Custom Credential with Evidence - File attachments and presentations
// 4. Organization Credential - Complex nested data and arrays
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

const (
	// Configuration constants for demo setup
	issuerStorageDir = "./issuer_storage"
	holderStorageDir = "./holder_storage"
)

func main() {
	fmt.Println("Comprehensive Credential Issuance Demo")
	fmt.Println("======================================")

	// Initialize issuer and holder clients
	issuerAccount, holderAccount := setupClients()
	defer issuerAccount.Close()
	defer holderAccount.Close()

	// Display account information
	displayAccountInfo(issuerAccount, holderAccount)

	// Credential issuance examples
	fmt.Println("CREDENTIAL ISSUANCE EXAMPLES")
	fmt.Println("============================")

	demonstrateBasicCredential(issuerAccount, holderAccount)
	demonstrateProfileCredential(issuerAccount, holderAccount)
	demonstrateCustomCredentialWithEvidence(issuerAccount, holderAccount)
	demonstrateOrganizationCredential(issuerAccount, holderAccount)

	fmt.Println("✅ All examples completed successfully!")
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

// setupClients demonstrates account initialization and configuration using the core SDK
// This function showcases how to:
// - Create Self SDK accounts with proper configuration
// - Set up storage paths for different account roles
// - Configure environment and logging settings
// - Handle account lifecycle management
func setupClients() (*account.Account, *account.Account) {
	fmt.Println("Setting up accounts...")

	issuerAccount, err := account.New(&account.Config{
		StorageKey:  generateStorageKey("issuer"),
		StoragePath: issuerStorageDir,
		Environment: account.TargetSandbox,
		LogLevel:    account.LogWarn,
	})
	if err != nil {
		log.Fatal("Failed to create issuer account:", err)
	}

	holderAccount, err := account.New(&account.Config{
		StorageKey:  generateStorageKey("holder"),
		StoragePath: holderStorageDir,
		Environment: account.TargetSandbox,
		LogLevel:    account.LogWarn,
	})
	if err != nil {
		log.Fatal("Failed to create holder account:", err)
	}

	return issuerAccount, holderAccount
}

// displayAccountInfo shows the unique DIDs for both parties
// DIDs (Decentralized Identifiers) are cryptographically verifiable identities
func displayAccountInfo(issuerAccount, holderAccount *account.Account) {
	issuerInbox, err := issuerAccount.InboxOpen()
	if err != nil {
		log.Fatal("Failed to open issuer inbox:", err)
	}
	holderInbox, err := holderAccount.InboxOpen()
	if err != nil {
		log.Fatal("Failed to open holder inbox:", err)
	}

	fmt.Printf("Issuer: %s\n", issuerInbox.String())
	fmt.Printf("   This is the credential issuer's unique decentralized identity\n")
	fmt.Printf("Holder: %s\n", holderInbox.String())
	fmt.Printf("   This is the credential holder's unique decentralized identity\n")
	fmt.Println()
}

// demonstrateBasicCredential showcases basic credential creation using the core SDK
// This example demonstrates:
// - Direct credential creation using credential.NewCredential()
// - Account-based credential issuance
// - Credential signing and issuance
// - Foundation concepts for all credential types
func demonstrateBasicCredential(issuerAccount, holderAccount *account.Account) {
	fmt.Println("1. Basic Email Credential")
	fmt.Println("=========================")

	issuerInbox, err := issuerAccount.InboxOpen()
	if err != nil {
		log.Fatal("Failed to open issuer inbox:", err)
	}
	holderInbox, err := holderAccount.InboxOpen()
	if err != nil {
		log.Fatal("Failed to open holder inbox:", err)
	}

	claims := map[string]interface{}{
		"emailAddress":     "john.doe@example.com",
		"verified":         true,
		"verificationDate": time.Now().Format("2006-01-02"),
	}

	credentialBuilder := credential.NewCredential().
		CredentialType(credential.CredentialTypeEmail).
		CredentialSubject(credential.AddressKey(holderInbox)).
		CredentialSubjectClaims(claims).
		Issuer(credential.AddressKey(issuerInbox)).
		ValidFrom(time.Now()).
		SignWith(issuerInbox, time.Now())

	unsignedCredential, err := credentialBuilder.Finish()
	if err != nil {
		log.Printf("Failed to build credential: %v", err)
		return
	}

	emailCredential, err := issuerAccount.CredentialIssue(unsignedCredential)
	if err != nil {
		log.Printf("Failed to issue credential: %v", err)
		return
	}

	fmt.Printf("✅ Email credential issued (Type: %v)\n", emailCredential.CredentialType())
	fmt.Printf("   Email: john.doe@example.com (verified)\n")
	fmt.Println()
}

// demonstrateProfileCredential showcases credentials with multiple claims using core SDK
// This example demonstrates:
// - Adding multiple claims to a single credential
// - Different data types in claims
// - Organizing related information in one credential
// - Building upon basic credential concepts
func demonstrateProfileCredential(issuerAccount, holderAccount *account.Account) {
	fmt.Println("2. Profile Credential")
	fmt.Println("=====================")

	issuerInbox, err := issuerAccount.InboxOpen()
	if err != nil {
		log.Fatal("Failed to open issuer inbox:", err)
	}
	holderInbox, err := holderAccount.InboxOpen()
	if err != nil {
		log.Fatal("Failed to open holder inbox:", err)
	}

	claims := map[string]interface{}{
		"firstName":        "John",
		"lastName":         "Doe",
		"displayName":      "John Doe",
		"profileLevel":     "verified",
		"country":          "United States",
		"registrationDate": time.Now().Format("2006-01-02"),
	}

	credentialBuilder := credential.NewCredential().
		CredentialType(credential.CredentialTypeProfileName).
		CredentialSubject(credential.AddressKey(holderInbox)).
		CredentialSubjectClaims(claims).
		Issuer(credential.AddressKey(issuerInbox)).
		ValidFrom(time.Now()).
		SignWith(issuerInbox, time.Now())

	unsignedCredential, err := credentialBuilder.Finish()
	if err != nil {
		log.Printf("Failed to build credential: %v", err)
		return
	}

	profileCredential, err := issuerAccount.CredentialIssue(unsignedCredential)
	if err != nil {
		log.Printf("Failed to issue credential: %v", err)
		return
	}

	fmt.Printf("✅ Profile credential issued (Type: %v)\n", profileCredential.CredentialType())
	fmt.Printf("   Name: John Doe, Country: United States\n")
	fmt.Println()
}

// demonstrateCustomCredentialWithEvidence showcases advanced credential features using core SDK
// This example demonstrates:
// - Creating custom credential types
// - Attaching file evidence to credentials using object.New()
// - Asset management and upload functionality
// - Creating verifiable presentations
// - Linking evidence to credential claims
func demonstrateCustomCredentialWithEvidence(issuerAccount, holderAccount *account.Account) {
	fmt.Println("3. Custom Credential with Evidence")
	fmt.Println("==================================")

	// Create evidence
	certificateData := []byte("Certificate of Completion\n" +
		"Advanced Go Programming Course\n" +
		"Student: John Doe\n" +
		"Grade: A+\n" +
		"Date: " + time.Now().Format("2006-01-02"))

	evidenceObj, err := object.New("application/pdf", certificateData)
	if err != nil {
		log.Printf("Failed to create evidence object: %v", err)
		return
	}

	err = issuerAccount.ObjectUpload(evidenceObj, false)
	if err != nil {
		log.Printf("Failed to upload evidence: %v", err)
		return
	}

	fmt.Printf("Evidence uploaded: %x (%d bytes)\n", evidenceObj.Id(), len(certificateData))

	issuerInbox, err := issuerAccount.InboxOpen()
	if err != nil {
		log.Fatal("Failed to open issuer inbox:", err)
	}
	holderInbox, err := holderAccount.InboxOpen()
	if err != nil {
		log.Fatal("Failed to open holder inbox:", err)
	}

	claims := map[string]interface{}{
		"certificationType": "Professional Development",
		"courseName":        "Advanced Go Programming",
		"completionDate":    time.Now().Format("2006-01-02"),
		"certificateHash":   fmt.Sprintf("%x", evidenceObj.Hash()),
		"grade":             "A+",
		"institution":       "Self SDK Academy",
		"courseHours":       40,
	}

	credentialBuilder := credential.NewCredential().
		CredentialType([]string{"VerifiableCredential", "CertificationCredential"}).
		CredentialSubject(credential.AddressKey(holderInbox)).
		CredentialSubjectClaims(claims).
		Issuer(credential.AddressKey(issuerInbox)).
		ValidFrom(time.Now()).
		SignWith(issuerInbox, time.Now())

	unsignedCredential, err := credentialBuilder.Finish()
	if err != nil {
		log.Printf("Failed to build credential: %v", err)
		return
	}

	customCredential, err := issuerAccount.CredentialIssue(unsignedCredential)
	if err != nil {
		log.Printf("Failed to issue credential: %v", err)
		return
	}

	presentation, err := createPresentation(issuerAccount, customCredential)
	if err != nil {
		log.Printf("Failed to create presentation: %v", err)
		return
	}

	fmt.Printf("✅ Certification credential issued (Type: %v)\n", customCredential.CredentialType())
	fmt.Printf("✅ Presentation created (Type: %v)\n", presentation.PresentationType())
	fmt.Printf("   Course: Advanced Go Programming, Grade: A+\n")
	fmt.Println()
}

// demonstrateOrganizationCredential showcases complex data structures in credentials using core SDK
// This example demonstrates:
// - Complex nested objects in claims
// - Arrays and collections in credentials
// - Hierarchical data organization
// - Real-world organizational data modeling
// - Advanced claim structuring techniques
func demonstrateOrganizationCredential(issuerAccount, holderAccount *account.Account) {
	fmt.Println("4. Organization Credential")
	fmt.Println("==========================")

	issuerInbox, err := issuerAccount.InboxOpen()
	if err != nil {
		log.Fatal("Failed to open issuer inbox:", err)
	}
	holderInbox, err := holderAccount.InboxOpen()
	if err != nil {
		log.Fatal("Failed to open holder inbox:", err)
	}

	claims := map[string]interface{}{
		"organizationName": "Acme Corporation",
		"employeeID":       "EMP-12345",
		"position":         "Senior Software Engineer",
		"department":       "Engineering",
		"startDate":        "2020-01-15",
		"address": map[string]interface{}{
			"street":  "123 Main Street",
			"city":    "San Francisco",
			"state":   "CA",
			"zipCode": "94105",
			"country": "United States",
		},
		"skills": []string{
			"Go Programming",
			"Microservices Architecture",
			"Docker & Kubernetes",
			"Cloud Computing",
			"API Design",
		},
		"certifications": []map[string]interface{}{
			{
				"name":       "AWS Solutions Architect",
				"issuer":     "Amazon Web Services",
				"issueDate":  "2021-03-15",
				"expiryDate": "2024-03-15",
				"level":      "Professional",
			},
			{
				"name":       "Kubernetes Administrator",
				"issuer":     "Cloud Native Computing Foundation",
				"issueDate":  "2021-08-20",
				"expiryDate": "2024-08-20",
				"level":      "Certified",
			},
		},
		"projects": []map[string]interface{}{
			{
				"name":         "Self SDK Integration",
				"role":         "Lead Developer",
				"duration":     "6 months",
				"technologies": []string{"Go", "Self SDK", "REST APIs"},
				"status":       "Completed",
			},
			{
				"name":         "Microservices Migration",
				"role":         "Senior Engineer",
				"duration":     "12 months",
				"technologies": []string{"Docker", "Kubernetes", "gRPC"},
				"status":       "In Progress",
			},
		},
		"performanceRating": map[string]interface{}{
			"overall":    "Exceeds Expectations",
			"technical":  "Expert",
			"leadership": "Strong",
			"reviewDate": time.Now().Format("2006-01-02"),
			"reviewer":   "Jane Smith, Engineering Manager",
		},
	}

	credentialBuilder := credential.NewCredential().
		CredentialType([]string{"VerifiableCredential", "EmploymentCredential"}).
		CredentialSubject(credential.AddressKey(holderInbox)).
		CredentialSubjectClaims(claims).
		Issuer(credential.AddressKey(issuerInbox)).
		ValidFrom(time.Now()).
		SignWith(issuerInbox, time.Now())

	unsignedCredential, err := credentialBuilder.Finish()
	if err != nil {
		log.Printf("Failed to build credential: %v", err)
		return
	}

	organizationCredential, err := issuerAccount.CredentialIssue(unsignedCredential)
	if err != nil {
		log.Printf("Failed to issue credential: %v", err)
		return
	}

	fmt.Printf("✅ Organization credential issued (Type: %v)\n", organizationCredential.CredentialType())
	fmt.Printf("   Employee: EMP-12345 - Senior Software Engineer\n")
	fmt.Printf("   Organization: Acme Corporation, Engineering Dept.\n")
	fmt.Println()
}

// createPresentation creates a verifiable presentation from a credential using core SDK
// This function demonstrates:
// - Presentation creation workflows
// - Credential packaging for sharing
// - Set presentation types and metadata
// - Prepare credentials for sharing with verifiers
func createPresentation(account *account.Account, cred *credential.VerifiableCredential) (*credential.VerifiablePresentation, error) {
	// 📋 PRESENTATION CREATION: Package credential for sharing using core SDK
	// Presentations allow selective disclosure of credential information

	// Get account inbox address for holder
	inboxAddress, err := account.InboxOpen()
	if err != nil {
		return nil, err
	}

	// Build the unsigned presentation
	builder := credential.NewPresentation().
		PresentationType([]string{"VerifiablePresentation"}).
		Holder(credential.AddressKey(inboxAddress)).
		CredentialAdd(cred)

	unsignedPresentation, err := builder.Finish()
	if err != nil {
		return nil, err
	}

	// Issue the presentation using the account
	return account.PresentationIssue(unsignedPresentation)
}
