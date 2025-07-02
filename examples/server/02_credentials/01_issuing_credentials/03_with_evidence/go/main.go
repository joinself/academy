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
	fmt.Println("Evidence-Based Credential Issuance Demo")
	fmt.Println("=======================================")

	// Create issuer and holder accounts
	issuer, holder := createAccounts()
	defer issuer.Close()
	defer holder.Close()

	// Display account information
	displayAccountInfo(issuer, holder)

	// Create credentials with evidence
	createCertificationWithEvidence(issuer, holder)

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
		StorageKey:  generateStorageKey("evidence_issuer"),
		StoragePath: "./evidence_issuer_storage",
		Environment: account.TargetSandbox,
		LogLevel:    account.LogWarn,
	}

	issuer, err := account.New(issuerCfg)
	if err != nil {
		log.Fatal("Failed to create issuer account:", err)
	}

	holderCfg := &account.Config{
		StorageKey:  generateStorageKey("evidence_holder"),
		StoragePath: "./evidence_holder_storage",
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

func createCertificationWithEvidence(issuer, holder *account.Account) {
	fmt.Println("Creating certification with evidence...")

	// Create evidence asset
	evidence := createEvidence(issuer)
	if evidence == nil {
		return
	}

	// Create credential with evidence reference
	cred := createCredentialWithEvidence(issuer, holder, evidence)
	if cred == nil {
		return
	}

	// Create verifiable presentation
	createPresentation(issuer, cred)
}

func createEvidence(issuer *account.Account) *object.Object {
	fmt.Println("Creating evidence asset...")

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

	evidence, err := object.New("certificate.pdf", certificateData)
	if err != nil {
		log.Printf("Failed to create evidence object: %v", err)
		return nil
	}

	err = issuer.ObjectUpload(evidence, true)
	if err != nil {
		log.Printf("Failed to upload evidence: %v", err)
		return nil
	}

	fmt.Printf("✅ Evidence uploaded: %x (%d bytes)\n", evidence.Id(), len(certificateData))
	return evidence
}

func createCredentialWithEvidence(issuer, holder *account.Account, evidence *object.Object) *credential.VerifiableCredential {
	fmt.Println("Building certification credential...")

	issuerAddress, err := issuer.InboxOpen()
	if err != nil {
		log.Fatal("Failed to open issuer inbox:", err)
	}

	holderAddress, err := holder.InboxOpen()
	if err != nil {
		log.Fatal("Failed to open holder inbox:", err)
	}

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

	credentialBuilder := credential.NewCredential().
		CredentialType([]string{"VerifiableCredential", "CertificationCredential"}).
		CredentialSubject(credential.AddressKey(holderAddress)).
		CredentialSubjectClaims(claims).
		Issuer(credential.AddressKey(issuerAddress)).
		ValidFrom(time.Now()).
		SignWith(issuerAddress, time.Now())

	unsignedCredential, err := credentialBuilder.Finish()
	if err != nil {
		log.Printf("Failed to build credential: %v", err)
		return nil
	}

	customCredential, err := issuer.CredentialIssue(unsignedCredential)
	if err != nil {
		log.Printf("Failed to issue credential: %v", err)
		return nil
	}

	fmt.Printf("✅ Certification credential issued (Type: %v)\n", customCredential.CredentialType())
	return customCredential
}

func createPresentation(issuer *account.Account, cred *credential.VerifiableCredential) {
	fmt.Println("Creating verifiable presentation...")

	issuerAddress, err := issuer.InboxOpen()
	if err != nil {
		log.Fatal("Failed to open issuer inbox:", err)
	}

	presentationBuilder := credential.NewPresentation().
		PresentationType([]string{"VerifiablePresentation", "CertificationPresentation"}).
		Holder(credential.AddressKey(issuerAddress)).
		CredentialAdd(cred)

	unsignedPresentation, err := presentationBuilder.Finish()
	if err != nil {
		log.Printf("Failed to build presentation: %v", err)
		return
	}

	presentation, err := issuer.PresentationIssue(unsignedPresentation)
	if err != nil {
		log.Printf("Failed to issue presentation: %v", err)
		return
	}

	fmt.Printf("✅ Presentation created (Type: %v, Credentials: %d)\n",
		presentation.PresentationType(), len(presentation.Credentials()))
}
