// Package common provides shared utilities for Self SDK examples.
//
// This package contains common functionality used across multiple examples
// to reduce code duplication and provide consistent patterns.
package common

import (
	"fmt"
	"log"

	"github.com/joinself/self-go-sdk/account"
)

const defaultStorageKey = "276cb6191a345753adb0897c2c0a89370aebf44ef99e612747bee3cd4e757ffa"

// AccountConfig provides configuration options for account creation
type AccountConfig struct {
	// StorageDir is the base directory for storage (optional, defaults to current dir)
	StorageDir string

	// Callbacks for the account
	Callbacks account.Callbacks
}

// SetupAccount creates a Self SDK account with the provided configuration.
// This centralizes the common account setup pattern used across examples.
func SetupAccount(config AccountConfig) *account.Account {
	fmt.Println("🔧 Setting up Self account...")

	storagePath := "./storage"
	if config.StorageDir != "" {
		storagePath = fmt.Sprintf("%s_storage", config.StorageDir)
	}

	cfg := &account.Config{
		StorageKey:  []byte(defaultStorageKey),
		StoragePath: storagePath,
		Environment: account.TargetSandbox,
		LogLevel:    account.LogWarn,
		Callbacks:   config.Callbacks,
	}

	selfAccount, err := account.New(cfg)
	if err != nil {
		log.Fatalf("Failed to create Self account: %v", err)
	}

	fmt.Println("✅ Self account created successfully")
	return selfAccount
}

// SetupIssuerHolder creates both issuer and holder accounts with optional callback customization.
// This is a common pattern in credential examples.
func SetupIssuerHolder(issuerCallbacks, holderCallbacks account.Callbacks) (*account.Account, *account.Account) {
	fmt.Println("🔧 Setting up issuer and holder accounts...")

	issuer := SetupAccount(AccountConfig{
		StorageDir: "./issuer",
		Callbacks:  issuerCallbacks,
	})

	holder := SetupAccount(AccountConfig{
		StorageDir: "./holder",
		Callbacks:  holderCallbacks,
	})

	return issuer, holder
}

// SetupBasicAccount creates a simple account with minimal configuration.
// Useful for examples that don't need custom callbacks.
func SetupBasicAccount() *account.Account {
	return SetupAccount(AccountConfig{
		Callbacks: account.Callbacks{
			OnConnect: func(acc *account.Account) {
				fmt.Println("🔗 Connected to Self network")
			},
		},
	})
}

// DisplayAccountInfo shows connection information for an account
func DisplayAccountInfo(acc *account.Account, name string) {
	inboxAddress, err := acc.InboxOpen()
	if err != nil {
		log.Fatalf("Failed to open %s inbox: %v", name, err)
	}

	fmt.Printf("📬 %s Address: %s\n", name, inboxAddress.String())
}

// DisplayAccountPair shows information for both issuer and holder accounts
func DisplayAccountPair(issuer, holder *account.Account) {
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
