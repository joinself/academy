// Package main demonstrates creating a brand new Self account from scratch.
// See README.md for detailed explanations and concepts.
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joinself/academy/examples/go/common"
	"github.com/joinself/self-go-sdk/account"
	"github.com/joinself/self-go-sdk/event"
)

func main() {
	fmt.Println("🆕 New Account Creation Example")
	fmt.Println("===============================")

	// Check if account already exists
	fmt.Println("🔍 Checking for existing account...")
	if accountExists() {
		fmt.Println("⚠️  Account already exists in ./storage/")
		fmt.Println("💡 Delete ./storage/ to create fresh account, or use '../02_existing_account'")
		return
	}

	// Create new account
	fmt.Println("🔧 Creating new Self account...")
	selfAccount := createNewAccount()
	defer selfAccount.Close()

	// Display account information
	fmt.Println("\n📋 Account Information:")
	displayAccountInfo(selfAccount)
	displayStorageInfo()

	fmt.Println("\n✅ New Self account ready!")
}

func accountExists() bool {
	_, err := os.Stat("./storage")
	return !os.IsNotExist(err)
}

func createNewAccount() *account.Account {
	return common.SetupAccount(common.AccountConfig{
		Callbacks: account.Callbacks{
			OnConnect: func(acc *account.Account) {
				fmt.Println("✅ Connected to Self network")
			},
			OnMessage: func(selfAccount *account.Account, msg *event.Message) {
				// Message handler (unused in this example)
			},
		},
	})
}

func displayAccountInfo(selfAccount *account.Account) {
	inboxAddress, err := selfAccount.InboxOpen()
	if err != nil {
		log.Printf("⚠️ Could not open inbox: %v", err)
		return
	}

	fmt.Printf("🆔 Account DID: %s\n", inboxAddress.String())
	fmt.Printf("🔐 Status: Ready for secure communication\n")

	fmt.Println("\n💡 Your DID can be shared with others for connections")
}

func displayStorageInfo() {
	fmt.Println("\n📁 Storage: ./storage/ (encrypted)")
	fmt.Println("🔄 Use '../02_existing_account' to reload this account")
}
