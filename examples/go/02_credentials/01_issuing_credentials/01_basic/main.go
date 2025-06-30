// Package main demonstrates basic credential issuance using the Self SDK.
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
	fmt.Println("🎓 Basic Credential Issuance Demo")
	fmt.Println()

	issuer, holder := createAccounts()
	defer issuer.Close()
	defer holder.Close()

	displayAccounts(issuer, holder)
	createEmailCredential(issuer, holder)

	fmt.Println("✅ Demo completed!")
}

func createAccounts() (*account.Account, *account.Account) {
	issuer := common.SetupAccount(common.AccountConfig{
		StorageDir: "./basic_issuer",
		Callbacks:  account.Callbacks{},
	})

	holder := common.SetupAccount(common.AccountConfig{
		StorageDir: "./basic_holder",
		Callbacks:  account.Callbacks{},
	})

	return issuer, holder
}

func displayAccounts(issuer, holder *account.Account) {
	common.DisplayAccountInfo(issuer, "Issuer")
	common.DisplayAccountInfo(holder, "Holder")
	fmt.Println()
}

func createEmailCredential(issuer, holder *account.Account) {
	fmt.Println("📧 Creating email credential...")

	issuerAddress, err := issuer.InboxOpen()
	if err != nil {
		log.Fatal("Failed to open issuer inbox:", err)
	}

	holderAddress, err := holder.InboxOpen()
	if err != nil {
		log.Fatal("Failed to open holder inbox:", err)
	}

	claims := map[string]interface{}{
		"emailAddress": "alice@example.com",
		"verified":     true,
		"domain":       "example.com",
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
		log.Printf("Failed to build credential: %v", err)
		return
	}

	emailCredential, err := issuer.CredentialIssue(unsignedCredential)
	if err != nil {
		log.Printf("Failed to issue credential: %v", err)
		return
	}

	fmt.Printf("✅ Email credential created\n")
	fmt.Printf("   📧 Email: alice@example.com\n")
	fmt.Printf("   🔒 Type: %v\n", emailCredential.CredentialType())
	fmt.Printf("   🆔 Subject: %s\n", emailCredential.CredentialSubject().String())
	fmt.Println()
}
