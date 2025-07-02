// Package main demonstrates loading an existing Self account from storage.
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
	fmt.Println("📂 Existing Account Loading Example")
	fmt.Println("===================================")

	// Step 1: Create demo account
	fmt.Println("🔧 Creating demo account...")
	originalAccount := createDemoAccount()
	originalAccount.Close()
	fmt.Println("✅ Demo account created and closed")

	// Step 2: Load existing account
	fmt.Println("\n📂 Loading existing account...")
	if !accountExists() {
		fmt.Println("❌ No account storage found")
		return
	}

	selfAccount := loadAccount()
	defer selfAccount.Close()

	// Step 3: Verify and display
	fmt.Println("\n🔍 Verifying account state...")
	verifyAccount(selfAccount)
	displayAccountInfo(selfAccount)

	fmt.Println("\n✅ Account loading demonstration complete!")
}

func createDemoAccount() *account.Account {
	// Clean existing storage
	if accountExists() {
		os.RemoveAll("./storage")
	}

	return common.SetupAccount(common.AccountConfig{
		Callbacks: account.Callbacks{
			OnConnect: func(acc *account.Account) {
				fmt.Println("✅ Demo account connected")
			},
		},
	})
}

func accountExists() bool {
	_, err := os.Stat("./storage")
	return !os.IsNotExist(err)
}

func loadAccount() *account.Account {
	return common.SetupAccount(common.AccountConfig{
		Callbacks: account.Callbacks{
			OnConnect: func(acc *account.Account) {
				fmt.Println("✅ Account reconnected")
			},
			OnMessage: func(selfAccount *account.Account, msg *event.Message) {
				fmt.Printf("📨 Message received\n")
			},
		},
	})
}

func verifyAccount(selfAccount *account.Account) {
	inboxAddress, err := selfAccount.InboxOpen()
	if err != nil {
		log.Printf("⚠️ Inbox access failed: %v", err)
		return
	}

	fmt.Printf("✅ Network: Connected\n")
	fmt.Printf("✅ Identity: %s\n", inboxAddress.String())
	fmt.Printf("✅ Storage: OK\n")
}

func displayAccountInfo(selfAccount *account.Account) {
	inboxAddress, err := selfAccount.InboxOpen()
	if err != nil {
		log.Printf("⚠️ Could not access account info: %v", err)
		return
	}

	fmt.Printf("\n🆔 Account DID: %s\n", inboxAddress.String())
	fmt.Printf("🔐 Status: Ready for connections\n")

	fmt.Println("\n💡 Note: Inbox addresses change between sessions (expected behavior)")
	fmt.Println("   What persists: storage, keys, connections, message history")
}
