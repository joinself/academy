// Package main demonstrates complex credential issuance using the Self SDK.
//
// This is the COMPLEX level of credential issuance examples.
// Prerequisites: Complete basic/main.go, multi_claim/main.go, and evidence/main.go first.
//
// This example shows:
// - Complex nested objects in claims
// - Arrays and collections in credentials
// - Hierarchical data organization
// - Real-world organizational data modeling
// - Advanced claim structuring techniques
//
// 🎯 What you'll learn:
// • How to structure complex nested data in credentials
// • Arrays and collections in claims
// • Hierarchical data organization
// • Real-world data modeling patterns
// • Advanced claim structuring
//
// 📚 Next steps:
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
	fmt.Println("🎓 Complex Credential Issuance Demo (Core SDK)")
	fmt.Println("===============================================")
	fmt.Println("This demo shows complex nested data structures in credentials.")
	fmt.Println()

	// Step 1: Create issuer and holder accounts
	issuer, holder := createAccounts()
	defer issuer.Close()
	defer holder.Close()

	// Display account information
	displayAccountInfo(issuer, holder)

	// Step 2: Create credentials with complex data structures
	createOrganizationCredential(issuer, holder)

	fmt.Println("✅ Complex demo completed!")
	fmt.Println()
	fmt.Println("📚 Ready for the next level?")
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
		StorageKey:  generateStorageKey("complex_issuer"),
		StoragePath: "./complex_issuer_storage",
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
		StorageKey:  generateStorageKey("complex_holder"),
		StoragePath: "./complex_holder_storage",
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

// createOrganizationCredential creates an organization credential with complex nested data
func createOrganizationCredential(issuer, holder *account.Account) {
	fmt.Println("🏢 Creating organization credential with complex nested data...")

	// Get inbox addresses for credential creation
	issuerAddress, err := issuer.InboxOpen()
	if err != nil {
		log.Fatal("Failed to open issuer inbox:", err)
	}

	holderAddress, err := holder.InboxOpen()
	if err != nil {
		log.Fatal("Failed to open holder inbox:", err)
	}

	// Create complex claims structure with nested objects and arrays
	claims := map[string]interface{}{
		"organizationName": "TechCorp Inc.", // Company name
		"employeeId":       "EMP-2024-001",  // Employee identifier
		"position": map[string]interface{}{ // Nested position object
			"title":      "Senior Software Engineer", // Job title
			"department": "Engineering",              // Department
			"level":      "L5",                       // Career level
			"startDate":  "2024-01-15",               // Start date
			"manager":    "jane.smith@techcorp.com",  // Manager reference
			"team":       "Backend Infrastructure",   // Team assignment
		},
		"permissions": []string{ // Array of permissions
			"read:repositories",    // Repository access
			"write:code",           // Code modification
			"deploy:staging",       // Staging deployment
			"review:pull-requests", // Code review
			"admin:team-resources", // Team administration
		},
		"contact": map[string]interface{}{ // Contact information
			"email":    "john.doe@techcorp.com",        // Work email
			"phone":    "+1-555-0123",                  // Work phone
			"office":   "Building A, Floor 3, Desk 42", // Office location
			"timezone": "America/New_York",             // Timezone
			"address": map[string]interface{}{ // Nested address
				"street":  "123 Tech Street",
				"city":    "San Francisco",
				"state":   "CA",
				"zipCode": "94105",
				"country": "United States",
			},
		},
		"benefits": map[string]interface{}{ // Benefits package
			"healthInsurance": true, // Health coverage
			"retirement401k":  true, // Retirement plan
			"paidTimeOff":     25,   // PTO days
			"stockOptions":    1000, // Stock options
			"remoteWork":      true, // Remote work eligibility
			"wellness": map[string]interface{}{ // Nested wellness benefits
				"gymMembership":    true,
				"mentalHealth":     true,
				"annualWellness":   "$1000",
				"flexibleSchedule": true,
			},
		},
		"certifications": []map[string]interface{}{ // Array of certifications
			{
				"name":       "AWS Solutions Architect", // Certification name
				"level":      "Professional",            // Certification level
				"issueDate":  "2023-06-15",              // Issue date
				"expiryDate": "2026-06-15",              // Expiry date
				"verified":   true,                      // Verification status
				"provider":   "Amazon Web Services",     // Certification provider
			},
			{
				"name":       "Kubernetes Administrator", // Second certification
				"level":      "Certified",                // Certification level
				"issueDate":  "2023-09-20",               // Issue date
				"expiryDate": "2026-09-20",               // Expiry date
				"verified":   true,                       // Verification status
				"provider":   "Cloud Native Computing Foundation",
			},
		},
		"projects": []map[string]interface{}{ // Array of projects
			{
				"name":         "Payment Gateway Redesign",
				"role":         "Lead Developer",
				"startDate":    "2023-01-01",
				"endDate":      "2023-06-30",
				"status":       "Completed",
				"technologies": []string{"Go", "PostgreSQL", "Redis", "Docker"},
			},
			{
				"name":         "Microservices Migration",
				"role":         "Senior Engineer",
				"startDate":    "2023-07-01",
				"endDate":      "2024-01-31",
				"status":       "Completed",
				"technologies": []string{"Go", "Kubernetes", "gRPC", "Prometheus"},
			},
		},
	}

	// Build the credential using the core SDK
	credentialBuilder := credential.NewCredential().
		CredentialType(credential.CredentialTypeOrganisation).
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
	orgCredential, err := issuer.CredentialIssue(unsignedCredential)
	if err != nil {
		log.Printf("Failed to issue credential: %v", err)
		return
	}

	// Display essential organization information
	fmt.Printf("✅ Organization credential created successfully\n")
	fmt.Printf("   🏢 Company: TechCorp Inc.\n")
	fmt.Printf("   💼 Position: Senior Software Engineer (L5)\n")
	fmt.Printf("   🏬 Department: Engineering - Backend Infrastructure\n")
	fmt.Printf("   🆔 Employee ID: EMP-2024-001\n")
	fmt.Printf("   📧 Email: john.doe@techcorp.com\n")
	fmt.Printf("   🔑 Permissions: 5 access levels\n")
	fmt.Printf("   🏆 Certifications: 2 professional certifications\n")
	fmt.Printf("   🚀 Projects: 2 completed projects\n")
	fmt.Printf("   🔒 Type: %v\n", orgCredential.CredentialType())
	fmt.Println()
	fmt.Println("🎓 What happened:")
	fmt.Println("   1. Created credential with deeply nested data structures")
	fmt.Println("   2. Used arrays for multiple values (permissions, certifications, projects)")
	fmt.Println("   3. Nested objects for hierarchical organization (position, contact, benefits)")
	fmt.Println("   4. Maintained cryptographic integrity for all nested data")
	fmt.Println()
}
