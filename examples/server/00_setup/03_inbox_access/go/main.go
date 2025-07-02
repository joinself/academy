// Package main demonstrates accessing a Self account's inbox address.
// See README.md for detailed explanations and concepts.
package main

import (
	"fmt"
	"log"

	"github.com/joinself/academy/examples/go/common"
)

func main() {
	fmt.Println("📧 Inbox Access Example")
	fmt.Println("========================")

	// Load or create account
	fmt.Println("🔧 Setting up account...")
	selfAccount := common.SetupAccount(common.AccountConfig{})
	defer selfAccount.Close()

	// Access inbox to get the address
	fmt.Println("\n📬 Accessing inbox...")
	inboxAddress, err := selfAccount.InboxOpen()
	if err != nil {
		log.Fatalf("❌ Failed to open inbox: %v", err)
	}

	// Display the inbox address
	fmt.Printf("✅ Inbox opened successfully!\n")
	fmt.Printf("📬 Your inbox address: %s\n", inboxAddress.String())
	fmt.Printf("\n💡 This address can be shared with others to receive:\n")
	fmt.Printf("   • Messages\n")
	fmt.Printf("   • Connection requests\n")
	fmt.Printf("   • Credentials\n")

	fmt.Println("\n✅ Inbox access demonstration complete!")
}
