// Package main demonstrates email credential verification through mobile delivery.
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joinself/academy/examples/go/common"
	"github.com/joinself/self-go-sdk/account"
	"github.com/joinself/self-go-sdk/credential"
	"github.com/joinself/self-go-sdk/event"
	"github.com/joinself/self-go-sdk/keypair/signing"
	"github.com/joinself/self-go-sdk/message"
)

var connectedMobile *signing.PublicKey

func main() {
	fmt.Println("Email Credential Verification Demo")
	fmt.Println("=================================")

	// Create email service provider account
	emailService := createEmailServiceAccount()
	defer emailService.Close()

	common.DisplayAccountInfo(emailService, "Email Service Provider")

	// Generate QR code for mobile connection
	if !generateEmailVerificationQR(emailService) {
		fmt.Println("❌ Failed to generate QR code. Please try again.")
		return
	}

	// Wait for mobile connection and credential delivery
	fmt.Println("\n📱 SCAN QR CODE with Self mobile app to verify email")
	fmt.Println("⏳ Waiting for mobile device connection...")
	fmt.Println("🔐 Once connected, email verification credential will be created")

	// Keep running to handle mobile connections
	select {}
}

func createEmailServiceAccount() *account.Account {
	fmt.Println("Setting up email service provider...")

	emailService := common.SetupAccount(common.AccountConfig{
		StorageDir: "email_service",
		Callbacks: account.Callbacks{
			OnWelcome: handleMobileConnection,
			OnMessage: func(acc *account.Account, msg *event.Message) {
				// Handle any additional messaging if needed
			},
		},
	})

	fmt.Println("✅ Email service provider ready")
	return emailService
}

func generateEmailVerificationQR(emailService *account.Account) bool {
	fmt.Println("Generating QR code for mobile email verification...")

	// Open inbox for receiving mobile connections
	inboxAddress, err := emailService.InboxOpen()
	if err != nil {
		log.Printf("❌ Failed to open inbox: %v", err)
		return false
	}

	// Generate key package for secure communication
	keyPackage, err := emailService.ConnectionNegotiateOutOfBand(
		inboxAddress,
		time.Now().Add(30*time.Minute),
	)
	if err != nil {
		log.Printf("❌ Failed to generate key package: %v", err)
		fmt.Println("💡 This may happen in demo environments or with network issues")
		return false
	}

	// Build discovery request for email verification
	content, err := message.NewDiscoveryRequest().
		KeyPackage(keyPackage).
		Expires(time.Now().Add(30 * time.Minute)).
		Finish()
	if err != nil {
		log.Printf("❌ Failed to build discovery request: %v", err)
		return false
	}

	// Create QR code
	anonymousMsg := event.NewAnonymousMessage(content)
	anonymousMsg.SetFlags(event.MessageFlagTargetSandbox)

	qrCode, err := anonymousMsg.EncodeToQR(event.QREncodingUnicode)
	if err != nil {
		log.Printf("❌ Failed to generate QR code: %v", err)
		return false
	}

	// Display QR code for mobile scanning
	fmt.Println("\n" + string(qrCode))
	fmt.Printf("⏱️  Expires: %s\n", time.Now().Add(30*time.Minute).Format("15:04:05"))

	return true
}

func handleMobileConnection(emailService *account.Account, welcome *event.Welcome) {
	fmt.Printf("\n📱 Mobile device connected: %s\n", welcome.FromAddress().String())

	// Accept the mobile connection
	_, err := emailService.ConnectionAccept(welcome.ToAddress(), welcome.Welcome())
	if err != nil {
		fmt.Printf("❌ Failed to accept connection: %v\n", err)
		return
	}

	fmt.Println("✅ Mobile connection established!")

	// Store mobile address for credential delivery
	connectedMobile = welcome.FromAddress()

	// Create and demonstrate email verification credential
	demonstrateEmailCredentialCreation(emailService)
}

func demonstrateEmailCredentialCreation(emailService *account.Account) {
	if connectedMobile == nil {
		fmt.Println("❌ No mobile device connected")
		return
	}

	fmt.Println("\n📧 Creating email verification credential...")

	issuerAddress, _ := emailService.InboxOpen()

	// Create email verification credential
	emailCredential := createEmailVerificationCredential(emailService, issuerAddress, connectedMobile)
	if emailCredential == nil {
		fmt.Println("❌ Failed to create email credential")
		return
	}

	fmt.Println("✅ Email verification credential created successfully!")
	fmt.Println("📱 Credential details:")
	fmt.Println("   • Email: user@example.com")
	fmt.Println("   • Status: verified")
	fmt.Println("   • Domain: example.com")
	fmt.Println("   • Method: email_link_clicked")
	fmt.Println("   • Issuer: Example Email Service Provider")
	fmt.Println()
	fmt.Println("🎉 Email verification workflow completed!")
	fmt.Println("💡 In production, this credential would be:")
	fmt.Println("   • Sent directly to the mobile device")
	fmt.Println("   • Stored in the user's credential wallet")
	fmt.Println("   • Available for proving email ownership")
	fmt.Println("   • Reusable across different services")

	// Exit after successful credential creation
	os.Exit(0)
}

func createEmailVerificationCredential(issuerAccount *account.Account, issuerAddress, holderAddress *signing.PublicKey) *credential.VerifiableCredential {
	// Create comprehensive email verification claims
	claims := map[string]interface{}{
		"emailAddress":       "user@example.com",
		"verified":           true,
		"verificationDate":   time.Now().Format("2006-01-02T15:04:05Z"),
		"domain":             "example.com",
		"verificationMethod": "email_link_clicked",
		"issuerName":         "Example Email Service Provider",
		"verificationLevel":  "standard",
		"ipAddress":          "192.168.1.100", // Where verification occurred
		"userAgent":          "Mozilla/5.0 (Mobile)",
	}

	credentialBuilder := credential.NewCredential().
		CredentialType([]string{"VerifiableCredential", "EmailCredential"}).
		CredentialSubject(credential.AddressKey(holderAddress)).
		CredentialSubjectClaims(claims).
		Issuer(credential.AddressKey(issuerAddress)).
		ValidFrom(time.Now()).
		SignWith(issuerAddress, time.Now())

	unsignedCredential, err := credentialBuilder.Finish()
	if err != nil {
		log.Printf("Failed to build email credential: %v", err)
		return nil
	}

	emailCredential, err := issuerAccount.CredentialIssue(unsignedCredential)
	if err != nil {
		log.Printf("Failed to issue email credential: %v", err)
		return nil
	}

	return emailCredential
}
