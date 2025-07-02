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
	fmt.Println("Multi-Claim Credential Issuance Demo")
	fmt.Println("====================================")

	// Create issuer and holder accounts
	issuer, holder := createAccounts()
	defer issuer.Close()
	defer holder.Close()

	// Display account information
	displayAccountInfo(issuer, holder)

	// Create credentials with multiple claims
	createProfileCredential(issuer, holder)
	createEducationCredential(issuer, holder)

	fmt.Println("✅ Demo completed successfully!")
}

func generateStorageKey(seed string) []byte {
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		h := sha256.Sum256([]byte(fmt.Sprintf("self-sdk-%s-%d", seed, time.Now().UnixNano())))
		return h[:]
	}
	return key
}

func createAccounts() (*account.Account, *account.Account) {
	issuerCfg := &account.Config{
		StorageKey:  generateStorageKey("multi_issuer"),
		StoragePath: "./multi_issuer_storage",
		Environment: account.TargetSandbox,
		LogLevel:    account.LogWarn,
	}

	issuer, err := account.New(issuerCfg)
	if err != nil {
		log.Fatal("Failed to create issuer account:", err)
	}

	holderCfg := &account.Config{
		StorageKey:  generateStorageKey("multi_holder"),
		StoragePath: "./multi_holder_storage",
		Environment: account.TargetSandbox,
		LogLevel:    account.LogWarn,
	}

	holder, err := account.New(holderCfg)
	if err != nil {
		log.Fatal("Failed to create holder account:", err)
	}

	return issuer, holder
}

func displayAccountInfo(issuer, holder *account.Account) {
	issuerInbox, err := issuer.InboxOpen()
	if err != nil {
		log.Fatal("Failed to open issuer inbox:", err)
	}
	holderInbox, err := holder.InboxOpen()
	if err != nil {
		log.Fatal("Failed to open holder inbox:", err)
	}

	fmt.Printf("Issuer: %s\n", issuerInbox.String())
	fmt.Printf("Holder: %s\n", holderInbox.String())
	fmt.Println()
}

func createProfileCredential(issuer, holder *account.Account) {
	fmt.Println("Creating profile credential...")

	issuerAddress, err := issuer.InboxOpen()
	if err != nil {
		log.Fatal("Failed to open issuer inbox:", err)
	}

	holderAddress, err := holder.InboxOpen()
	if err != nil {
		log.Fatal("Failed to open holder inbox:", err)
	}

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

	credentialBuilder := credential.NewCredential().
		CredentialType(credential.CredentialTypeProfileName).
		CredentialSubject(credential.AddressKey(holderAddress)).
		CredentialSubjectClaims(claims).
		Issuer(credential.AddressKey(issuerAddress)).
		ValidFrom(time.Now()).
		SignWith(issuerAddress, time.Now())

	unsignedCredential, err := credentialBuilder.Finish()
	if err != nil {
		log.Printf("Failed to build credential: %v", err)
		return
	}

	profileCredential, err := issuer.CredentialIssue(unsignedCredential)
	if err != nil {
		log.Printf("Failed to issue credential: %v", err)
		return
	}

	fmt.Printf("✅ Profile credential issued (Type: %v)\n", profileCredential.CredentialType())
}

func createEducationCredential(issuer, holder *account.Account) {
	fmt.Println("Creating education credential...")

	issuerAddress, err := issuer.InboxOpen()
	if err != nil {
		log.Fatal("Failed to open issuer inbox:", err)
	}

	holderAddress, err := holder.InboxOpen()
	if err != nil {
		log.Fatal("Failed to open holder inbox:", err)
	}

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

	credentialBuilder := credential.NewCredential().
		CredentialType([]string{"VerifiableCredential", "EducationCredential"}).
		CredentialSubject(credential.AddressKey(holderAddress)).
		CredentialSubjectClaims(claims).
		Issuer(credential.AddressKey(issuerAddress)).
		ValidFrom(time.Now()).
		SignWith(issuerAddress, time.Now())

	unsignedCredential, err := credentialBuilder.Finish()
	if err != nil {
		log.Printf("Failed to build credential: %v", err)
		return
	}

	educationCredential, err := issuer.CredentialIssue(unsignedCredential)
	if err != nil {
		log.Printf("Failed to issue credential: %v", err)
		return
	}

	fmt.Printf("✅ Education credential issued (Type: %v)\n", educationCredential.CredentialType())
}
