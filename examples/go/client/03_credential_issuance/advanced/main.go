// Package main demonstrates comprehensive credential issuance capabilities using the underlying Self SDK directly.
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
	fmt.Println("🎓 Self SDK Credential Issuance Demo")
	fmt.Println("=====================================")
	fmt.Println("📚 This demo showcases comprehensive credential issuance capabilities:")
	fmt.Println("   • Creating various types of verifiable credentials")
	fmt.Println("   • Using the credential builder pattern")
	fmt.Println("   • Attaching evidence and files to credentials")
	fmt.Println("   • Managing complex nested claims")
	fmt.Println("   • Creating verifiable presentations")
	fmt.Println("   • Handling credential request/response workflows")
	fmt.Println()

	// 🏗️ STEP 1: CLIENT SETUP - Initialize issuer and holder clients
	// The issuer creates and signs credentials, while the holder receives and stores them
	issuerAccount, holderAccount := setupClients()
	defer issuerAccount.Close()
	defer holderAccount.Close()

	// 🆔 IDENTITY DISPLAY: Show the unique DIDs for both parties
	// DIDs (Decentralized Identifiers) are cryptographically verifiable identities
	issuerInbox, err := issuerAccount.InboxOpen()
	if err != nil {
		log.Fatal("Failed to open issuer inbox:", err)
	}
	holderInbox, err := holderAccount.InboxOpen()
	if err != nil {
		log.Fatal("Failed to open holder inbox:", err)
	}

	fmt.Printf("🏢 Issuer DID: %s\n", issuerInbox.String())
	fmt.Printf("   This is the credential issuer's unique decentralized identity\n")
	fmt.Printf("👤 Holder DID: %s\n", holderInbox.String())
	fmt.Printf("   This is the credential holder's unique decentralized identity\n")
	fmt.Println()

	// 🔧 STEP 2: HANDLER SETUP - Configure credential request/response handlers
	// These handlers demonstrate how to process incoming credential requests
	setupCredentialHandlers(issuerAccount, holderAccount)

	// 🎯 STEP 3: CREDENTIAL ISSUANCE EXAMPLES
	// Progressive examples from simple to complex credential types
	fmt.Println("📚 CREDENTIAL ISSUANCE EXAMPLES")
	fmt.Println("================================")
	fmt.Println("🎯 The following examples demonstrate progressive complexity:")
	fmt.Println("   Each example builds upon concepts from the previous ones")
	fmt.Println()

	// 📧 EXAMPLE 1: Basic Email Credential - Foundation concepts
	demonstrateBasicCredential(issuerAccount, holderAccount)

	// 👤 EXAMPLE 2: Profile Credential - Multiple claims
	demonstrateProfileCredential(issuerAccount, holderAccount)

	// 🎓 EXAMPLE 3: Custom Credential - Evidence and presentations
	demonstrateCustomCredentialWithEvidence(issuerAccount, holderAccount)

	// 🏢 EXAMPLE 4: Organization Credential - Complex data structures
	demonstrateOrganizationCredential(issuerAccount, holderAccount)

	// 🔗 STEP 4: OPTIONAL DISCOVERY DEMO
	// Discovery workflow is separated to maintain focus on credential issuance
	fmt.Println("\n🔗 DISCOVERY & CONNECTION (Optional)")
	fmt.Println("====================================")
	fmt.Println("📱 The discovery workflow demonstrates peer-to-peer connections")
	fmt.Println("   For credential issuance focus, this section is optional")
	fmt.Println("   Uncomment runDiscoveryDemo() below to enable QR code discovery")
	fmt.Println()

	// Uncomment the line below to run the discovery demo
	// runDiscoveryDemo(issuerAccount, holderAccount)

	// 🎉 STEP 5: EDUCATIONAL SUMMARY
	// Comprehensive summary of demonstrated features and next steps
	printSummary()
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
	fmt.Println("🔧 SETTING UP SELF SDK ACCOUNTS")
	fmt.Println("===============================")
	fmt.Println("🏗️ Initializing issuer and holder accounts...")
	fmt.Println("   The issuer creates and signs credentials")
	fmt.Println("   The holder receives and manages credentials")
	fmt.Println()

	// 🏢 ISSUER ACCOUNT: Creates and signs verifiable credentials
	// The issuer account has the authority to create credentials for subjects
	fmt.Println("🏢 Creating issuer account...")
	issuerAccount, err := account.New(&account.Config{
		StorageKey:  generateStorageKey("issuer"), // Unique key for issuer storage encryption
		StoragePath: issuerStorageDir,             // Dedicated storage directory for issuer
		Environment: account.TargetSandbox,        // Use Sandbox environment for development
		LogLevel:    account.LogWarn,              // Show warning log messages
	})
	if err != nil {
		log.Fatal("❌ Failed to create issuer account:", err)
	}

	// 👤 HOLDER ACCOUNT: Receives and stores verifiable credentials
	// The holder account manages credentials issued by various issuers
	fmt.Println("👤 Creating holder account...")
	holderAccount, err := account.New(&account.Config{
		StorageKey:  generateStorageKey("holder"), // Unique key for holder storage encryption
		StoragePath: holderStorageDir,             // Dedicated storage directory for holder
		Environment: account.TargetSandbox,        // Use Sandbox environment for development
		LogLevel:    account.LogWarn,              // Show warning log messages
	})
	if err != nil {
		log.Fatal("❌ Failed to create holder account:", err)
	}

	fmt.Println("✅ Accounts created successfully")
	fmt.Println("   🔐 Both accounts use encrypted local storage")
	fmt.Println("   🌐 Connected to Self Sandbox environment")
	fmt.Println()

	return issuerAccount, holderAccount
}

// setupCredentialHandlers demonstrates account setup for credential operations
// Note: The core SDK uses a different pattern for event handling than the client package
func setupCredentialHandlers(issuerAccount, holderAccount *account.Account) {
	fmt.Println("🔧 ACCOUNT SETUP COMPLETE")
	fmt.Println("=========================")
	fmt.Println("📨 Accounts are ready for credential operations...")
	fmt.Println("   Core SDK uses direct method calls instead of event handlers")
	fmt.Println("   Credential issuance, verification, and presentation are handled synchronously")
	fmt.Println()
}

// demonstrateBasicCredential showcases basic credential creation using the core SDK
// This example demonstrates:
// - Direct credential creation using credential.NewCredential()
// - Account-based credential issuance
// - Credential signing and issuance
// - Foundation concepts for all credential types
func demonstrateBasicCredential(issuerAccount, holderAccount *account.Account) {
	fmt.Println("1️⃣ BASIC EMAIL CREDENTIAL")
	fmt.Println("==========================")
	fmt.Println("📧 Creating a simple email verification credential...")
	fmt.Println("   This demonstrates the foundation of credential issuance")
	fmt.Println("   Key concepts: builder pattern, claims, signing, issuance")
	fmt.Println()

	// Get DIDs from accounts
	issuerInbox, err := issuerAccount.InboxOpen()
	if err != nil {
		log.Fatal("Failed to open issuer inbox:", err)
	}
	holderInbox, err := holderAccount.InboxOpen()
	if err != nil {
		log.Fatal("Failed to open holder inbox:", err)
	}

	// 🏗️ CREDENTIAL BUILDER: Use the core SDK builder pattern for credential creation
	fmt.Println("🏗️ Using core SDK credential builder pattern...")

	// Create credential claims
	claims := map[string]interface{}{
		"emailAddress":     "john.doe@example.com",
		"verified":         true,
		"verificationDate": time.Now().Format("2006-01-02"),
	}

	// Build the credential using the core SDK
	credentialBuilder := credential.NewCredential().
		CredentialType(credential.CredentialTypeEmail).
		CredentialSubject(credential.AddressKey(holderInbox)).
		CredentialSubjectClaims(claims).
		Issuer(credential.AddressKey(issuerInbox)).
		ValidFrom(time.Now()).
		SignWith(issuerInbox, time.Now())

	// Finish building the unsigned credential
	unsignedCredential, err := credentialBuilder.Finish()
	if err != nil {
		log.Printf("❌ Failed to build credential: %v", err)
		return
	}

	// Issue the credential using the issuer's account
	emailCredential, err := issuerAccount.CredentialIssue(unsignedCredential)
	if err != nil {
		log.Printf("❌ Failed to issue credential: %v", err)
		return
	}

	// ✅ SUCCESS REPORTING: Display credential creation results
	fmt.Printf("   ✅ Email credential created successfully\n")
	fmt.Printf("   📧 Email: john.doe@example.com\n")
	fmt.Printf("   ✔️  Verified: true\n")
	fmt.Printf("   📅 Verification Date: %s\n", time.Now().Format("2006-01-02"))
	fmt.Printf("   🔒 Credential Type: %v\n", emailCredential.CredentialType())
	fmt.Printf("   🆔 Subject: %s\n", emailCredential.CredentialSubject().String())
	fmt.Printf("   🏢 Issuer: %s\n", emailCredential.Issuer().String())
	fmt.Println()
	fmt.Println("📚 Key Learning Points:")
	fmt.Println("   • Credentials contain claims about a subject")
	fmt.Println("   • Core SDK provides direct credential construction")
	fmt.Println("   • Cryptographic signatures ensure integrity")
	fmt.Println("   • Timestamps establish validity periods")
	fmt.Println()
}

// demonstrateProfileCredential showcases credentials with multiple claims using core SDK
// This example demonstrates:
// - Adding multiple claims to a single credential
// - Different data types in claims
// - Organizing related information in one credential
// - Building upon basic credential concepts
func demonstrateProfileCredential(issuerAccount, holderAccount *account.Account) {
	fmt.Println("2️⃣ PROFILE CREDENTIAL WITH MULTIPLE CLAIMS")
	fmt.Println("===========================================")
	fmt.Println("👤 Creating a profile credential with multiple claims...")
	fmt.Println("   This demonstrates how to include multiple pieces of information")
	fmt.Println("   in a single credential for related identity attributes")
	fmt.Println()

	// Get DIDs from accounts
	issuerInbox, err := issuerAccount.InboxOpen()
	if err != nil {
		log.Fatal("Failed to open issuer inbox:", err)
	}
	holderInbox, err := holderAccount.InboxOpen()
	if err != nil {
		log.Fatal("Failed to open holder inbox:", err)
	}

	// 🏗️ MULTI-CLAIM BUILDER: Demonstrate adding multiple related claims
	fmt.Println("🏗️ Building credential with multiple claims...")

	// Create credential claims
	claims := map[string]interface{}{
		"firstName":        "John",
		"lastName":         "Doe",
		"displayName":      "John Doe",
		"profileLevel":     "verified",
		"country":          "United States",
		"registrationDate": time.Now().Format("2006-01-02"),
	}

	// Build the credential using the core SDK
	credentialBuilder := credential.NewCredential().
		CredentialType(credential.CredentialTypeProfileName).
		CredentialSubject(credential.AddressKey(holderInbox)).
		CredentialSubjectClaims(claims).
		Issuer(credential.AddressKey(issuerInbox)).
		ValidFrom(time.Now()).
		SignWith(issuerInbox, time.Now())

	// Finish building the unsigned credential
	unsignedCredential, err := credentialBuilder.Finish()
	if err != nil {
		log.Printf("❌ Failed to build credential: %v", err)
		return
	}

	// Issue the credential using the issuer's account
	profileCredential, err := issuerAccount.CredentialIssue(unsignedCredential)
	if err != nil {
		log.Printf("❌ Failed to issue credential: %v", err)
		return
	}

	// ✅ SUCCESS REPORTING: Display comprehensive credential information
	fmt.Printf("   ✅ Profile credential created successfully\n")
	fmt.Printf("   👤 Name: John Doe\n")
	fmt.Printf("   🌍 Country: United States\n")
	fmt.Printf("   ⭐ Profile Level: verified\n")
	fmt.Printf("   📅 Registration: %s\n", time.Now().Format("2006-01-02"))
	fmt.Printf("   🔒 Credential Type: %v\n", profileCredential.CredentialType())
	fmt.Println()
	fmt.Println("📚 Key Learning Points:")
	fmt.Println("   • Multiple related claims can be grouped in one credential")
	fmt.Println("   • Claims can contain different data types (strings, booleans, dates)")
	fmt.Println("   • Grouping related information improves efficiency")
	fmt.Println("   • Each claim is cryptographically protected")
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
	fmt.Println("3️⃣ CUSTOM CREDENTIAL WITH EVIDENCE")
	fmt.Println("===================================")
	fmt.Println("🎓 Creating a certification credential with file evidence...")
	fmt.Println("   This demonstrates advanced features: custom types, evidence, presentations")
	fmt.Println("   Evidence provides additional proof supporting credential claims")
	fmt.Println()

	// 📄 EVIDENCE CREATION: Create and upload supporting documentation using core SDK
	fmt.Println("📄 Creating evidence asset...")
	fmt.Println("   Evidence can be any file type: PDFs, images, documents, etc.")
	certificateData := []byte("This is a mock certificate document for demonstration purposes.\n" +
		"Certificate of Completion\n" +
		"Advanced Go Programming Course\n" +
		"Student: John Doe\n" +
		"Grade: A+\n" +
		"Date: " + time.Now().Format("2006-01-02"))

	// Create encrypted object using core SDK
	evidenceObj, err := object.New("application/pdf", certificateData)
	if err != nil {
		log.Printf("❌ Failed to create evidence object: %v", err)
		return
	}

	// Upload to object store
	err = issuerAccount.ObjectUpload(evidenceObj, false)
	if err != nil {
		log.Printf("❌ Failed to upload evidence: %v", err)
		return
	}

	fmt.Printf("   📄 Evidence created: certificate.pdf\n")
	fmt.Printf("   🔗 Asset ID: %x\n", evidenceObj.Id())
	fmt.Printf("   🔐 Content Hash: %x\n", evidenceObj.Hash())
	fmt.Println("   ✅ Evidence uploaded to secure storage")
	fmt.Println()

	// Get DIDs from accounts
	issuerInbox, err := issuerAccount.InboxOpen()
	if err != nil {
		log.Fatal("Failed to open issuer inbox:", err)
	}
	holderInbox, err := holderAccount.InboxOpen()
	if err != nil {
		log.Fatal("Failed to open holder inbox:", err)
	}

	// 🏗️ CUSTOM CREDENTIAL: Create credential with evidence reference
	fmt.Println("🏗️ Building custom certification credential...")

	// Create credential claims
	claims := map[string]interface{}{
		"certificationType": "Professional Development",
		"courseName":        "Advanced Go Programming",
		"completionDate":    time.Now().Format("2006-01-02"),
		"certificateHash":   fmt.Sprintf("%x", evidenceObj.Hash()),
		"grade":             "A+",
		"institution":       "Self SDK Academy",
		"courseHours":       40,
	}

	// Build the credential using the core SDK
	credentialBuilder := credential.NewCredential().
		CredentialType([]string{"VerifiableCredential", "CertificationCredential"}).
		CredentialSubject(credential.AddressKey(holderInbox)).
		CredentialSubjectClaims(claims).
		Issuer(credential.AddressKey(issuerInbox)).
		ValidFrom(time.Now()).
		SignWith(issuerInbox, time.Now())

	// Finish building the unsigned credential
	unsignedCredential, err := credentialBuilder.Finish()
	if err != nil {
		log.Printf("❌ Failed to build credential: %v", err)
		return
	}

	// Issue the credential using the issuer's account
	customCredential, err := issuerAccount.CredentialIssue(unsignedCredential)
	if err != nil {
		log.Printf("❌ Failed to issue credential: %v", err)
		return
	}

	// 📋 PRESENTATION CREATION: Create verifiable presentation from credential
	fmt.Println("📋 Creating verifiable presentation...")
	fmt.Println("   Presentations package credentials for sharing with verifiers")
	presentation, err := createPresentation(issuerAccount, customCredential)
	if err != nil {
		log.Printf("❌ Failed to create presentation: %v", err)
		return
	}

	// ✅ SUCCESS REPORTING: Display comprehensive results
	fmt.Printf("   ✅ Certification credential created successfully\n")
	fmt.Printf("   🎓 Course: Advanced Go Programming\n")
	fmt.Printf("   📅 Completed: %s\n", time.Now().Format("2006-01-02"))
	fmt.Printf("   🏆 Grade: A+\n")
	fmt.Printf("   🏫 Institution: Self SDK Academy\n")
	fmt.Printf("   ⏱️  Duration: 40 hours\n")
	fmt.Printf("   🔒 Credential Type: %v\n", customCredential.CredentialType())
	fmt.Printf("   📋 Presentation Type: %v\n", presentation.PresentationType())
	fmt.Printf("   🔗 Evidence Hash: %x\n", evidenceObj.Hash())
	fmt.Println()
	fmt.Println("📚 Key Learning Points:")
	fmt.Println("   • Custom credential types support specific use cases")
	fmt.Println("   • Evidence provides additional verification material")
	fmt.Println("   • Core SDK object.New() handles secure file storage")
	fmt.Println("   • Presentations package credentials for sharing")
	fmt.Println("   • Hash references link credentials to evidence")
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
	fmt.Println("4️⃣ ORGANIZATION CREDENTIAL WITH COMPLEX CLAIMS")
	fmt.Println("===============================================")
	fmt.Println("🏢 Creating an organization credential with complex nested data...")
	fmt.Println("   This demonstrates advanced data structures: nested objects, arrays")
	fmt.Println("   Real-world credentials often contain hierarchical information")
	fmt.Println()

	// Get DIDs from accounts
	issuerInbox, err := issuerAccount.InboxOpen()
	if err != nil {
		log.Fatal("Failed to open issuer inbox:", err)
	}
	holderInbox, err := holderAccount.InboxOpen()
	if err != nil {
		log.Fatal("Failed to open holder inbox:", err)
	}

	// 🏗️ COMPLEX CLAIMS: Demonstrate nested objects and arrays
	fmt.Println("🏗️ Building credential with complex nested claims...")

	// Create complex credential claims with nested data
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

	// Build the credential using the core SDK
	credentialBuilder := credential.NewCredential().
		CredentialType([]string{"VerifiableCredential", "EmploymentCredential"}).
		CredentialSubject(credential.AddressKey(holderInbox)).
		CredentialSubjectClaims(claims).
		Issuer(credential.AddressKey(issuerInbox)).
		ValidFrom(time.Now()).
		SignWith(issuerInbox, time.Now())

	// Finish building the unsigned credential
	unsignedCredential, err := credentialBuilder.Finish()
	if err != nil {
		log.Printf("❌ Failed to build credential: %v", err)
		return
	}

	// Issue the credential using the issuer's account
	organizationCredential, err := issuerAccount.CredentialIssue(unsignedCredential)
	if err != nil {
		log.Printf("❌ Failed to issue credential: %v", err)
		return
	}

	// ✅ SUCCESS REPORTING: Display comprehensive organizational credential
	fmt.Printf("   ✅ Organization credential created successfully\n")
	fmt.Printf("   🏢 Organization: Acme Corporation\n")
	fmt.Printf("   👤 Employee: EMP-12345\n")
	fmt.Printf("   💼 Position: Senior Software Engineer\n")
	fmt.Printf("   🏭 Department: Engineering\n")
	fmt.Printf("   📅 Start Date: 2020-01-15\n")
	fmt.Printf("   🌟 Performance: Exceeds Expectations\n")
	fmt.Printf("   🎯 Skills: 5 technical competencies\n")
	fmt.Printf("   🏆 Certifications: 2 professional certifications\n")
	fmt.Printf("   📊 Projects: 2 major project contributions\n")
	fmt.Printf("   🔒 Credential Type: %v\n", organizationCredential.CredentialType())
	fmt.Println()
	fmt.Println("📚 Key Learning Points:")
	fmt.Println("   • Complex nested objects represent real-world data structures")
	fmt.Println("   • Arrays handle collections of related information")
	fmt.Println("   • Hierarchical organization improves data clarity")
	fmt.Println("   • All nested data maintains cryptographic integrity")
	fmt.Println("   • Core SDK supports arbitrary JSON structures in claims")
	fmt.Println()
}

// runDiscoveryDemo demonstrates peer discovery and connection workflows (optional)
// This function showcases:
// - QR code generation for peer discovery
// - Connection establishment between peers
// - Integration of discovery with credential workflows
// - Error handling and timeout management
func runDiscoveryDemo(issuerAccount, holderAccount *account.Account) {
	fmt.Println("🔗 PEER DISCOVERY DEMONSTRATION")
	fmt.Println("===============================")
	fmt.Println("📱 Setting up peer discovery workflow...")
	fmt.Println("   Discovery enables peers to find and connect with each other")
	fmt.Println("   QR codes provide a user-friendly connection method")
	fmt.Println()

	fmt.Println("⚠️  Discovery demo requires additional setup:")
	fmt.Println("   • QR code display functionality")
	fmt.Println("   • Network connectivity configuration")
	fmt.Println("   • Peer discovery service integration")
	fmt.Println()
	fmt.Println("🔄 For credential issuance focus, discovery is optional")
	fmt.Println("   The core credential operations work without discovery")
	fmt.Println("   Discovery enhances peer-to-peer credential exchange")
	fmt.Println()
}

// printSummary provides comprehensive educational summary
func printSummary() {
	fmt.Println("🎉 CREDENTIAL ISSUANCE DEMO COMPLETE")
	fmt.Println("====================================")
	fmt.Println("🎓 Congratulations! You've successfully completed the credential issuance demo.")
	fmt.Println("   This advanced example demonstrated comprehensive credential capabilities:")
	fmt.Println()
	fmt.Println("✅ FEATURES DEMONSTRATED:")
	fmt.Println("   1️⃣ Basic Email Credential - Foundation concepts with core SDK")
	fmt.Println("   2️⃣ Profile Credential - Multiple claims in single credential")
	fmt.Println("   3️⃣ Custom Credential - Evidence attachments and presentations")
	fmt.Println("   4️⃣ Organization Credential - Complex nested data structures")
	fmt.Println()
	fmt.Println("🔧 KEY SDK COMPONENTS UTILIZED:")
	fmt.Println("   • account.New() - Account initialization and configuration")
	fmt.Println("   • credential.NewCredential() - Direct credential construction")
	fmt.Println("   • object.New() - Evidence and file attachment management")
	fmt.Println("   • credential.NewPresentation() - Verifiable presentation creation")
	fmt.Println("   • account.CredentialIssue() - Direct credential issuance")
	fmt.Println()
	fmt.Println("📚 EDUCATIONAL PROGRESSION COMPLETE:")
	fmt.Println("   • Core SDK concepts and architecture")
	fmt.Println("   • Direct credential building patterns")
	fmt.Println("   • Evidence and asset management")
	fmt.Println("   • Complex data modeling techniques")
	fmt.Println("   • Production-ready development patterns")
	fmt.Println()
	fmt.Println("🚀 NEXT STEPS:")
	fmt.Println("   • Explore credential exchange examples")
	fmt.Println("   • Implement custom credential types for your use case")
	fmt.Println("   • Integrate Self SDK into your applications")
	fmt.Println("   • Build end-to-end identity solutions")
	fmt.Println()
	fmt.Println("📖 Additional Resources:")
	fmt.Println("   • Self SDK Documentation: https://docs.joinself.com")
	fmt.Println("   • Example Applications: ../../../examples/")
	fmt.Println("   • Community Support: https://github.com/joinself/self-go-sdk")
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
