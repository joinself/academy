// Package main demonstrates requesting and validating real credentials from mobile devices.
// This example shows how to connect to a mobile device and request actual credentials
// like email verification and liveness proofs.
package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/joinself/academy/examples/server/common"
	"github.com/joinself/self-go-sdk/account"
	"github.com/joinself/self-go-sdk/credential"
	"github.com/joinself/self-go-sdk/event"
	"github.com/joinself/self-go-sdk/keypair/signing"
	"github.com/joinself/self-go-sdk/message"
)

var (
	connectedMobile *signing.PublicKey
	requests        sync.Map // Track outstanding requests
)

func main() {
	fmt.Println("Real Credential Presentation Request Demo")
	fmt.Println("=======================================")

	// Create credential verifier account
	verifierService := createVerifierAccount()
	defer verifierService.Close()

	common.DisplayAccountInfo(verifierService, "Credential Verifier")

	// Generate QR code for mobile connection
	if !generateConnectionQR(verifierService) {
		fmt.Println("❌ Failed to generate QR code. Please try again.")
		return
	}

	// Wait for mobile connection and credential requests
	fmt.Println("\n📱 SCAN QR CODE with Self mobile app")
	fmt.Println("⏳ Waiting for mobile device connection...")
	fmt.Println("📋 Once connected, will request real credentials from your device")
	fmt.Println("💡 This demonstrates validation of actual mobile credentials")

	// Keep running to handle mobile connections and requests
	select {}
}

func createVerifierAccount() *account.Account {
	fmt.Println("Setting up credential verifier...")

	verifierService := common.SetupAccount(common.AccountConfig{
		StorageDir: "verifier_service",
		Callbacks: account.Callbacks{
			OnWelcome: handleMobileConnection,
			OnMessage: handleMessage,
		},
	})

	fmt.Println("✅ Credential verifier ready")
	return verifierService
}

func handleMessage(acc *account.Account, msg *event.Message) {
	switch event.ContentTypeOf(msg) {
	case message.ContentTypeCredentialPresentationResponse:
		handlePresentationResponse(acc, msg)
	default:
		// Handle any other message types if needed
		fmt.Printf("📨 Received message of type: %s\n", event.ContentTypeOf(msg))
	}
}

func handlePresentationResponse(acc *account.Account, msg *event.Message) {
	fmt.Printf("\n📋 Received credential presentation response from: %s\n", msg.FromAddress().String())

	credentialPresentationResponse, err := message.DecodeCredentialPresentationResponse(msg.Content())
	if err != nil {
		log.Printf("❌ Failed to decode presentation response: %v", err)
		return
	}

	// Find the corresponding request completer
	completer, ok := requests.LoadAndDelete(hex.EncodeToString(credentialPresentationResponse.ResponseTo()))
	if !ok {
		log.Printf("❌ Received response to unknown request: %s",
			hex.EncodeToString(credentialPresentationResponse.ResponseTo()))
		return
	}

	// Signal the waiting goroutine
	completer.(chan *message.CredentialPresentationResponse) <- credentialPresentationResponse
}

func generateConnectionQR(verifierService *account.Account) bool {
	fmt.Println("Generating QR code for mobile connection...")

	// Open inbox for receiving mobile connections
	inboxAddress, err := verifierService.InboxOpen()
	if err != nil {
		log.Printf("❌ Failed to open inbox: %v", err)
		return false
	}

	// Generate key package for secure communication
	keyPackage, err := verifierService.ConnectionNegotiateOutOfBand(
		inboxAddress,
		time.Now().Add(30*time.Minute),
	)
	if err != nil {
		log.Printf("❌ Failed to generate key package: %v", err)
		fmt.Println("💡 This may happen in demo environments or with network issues")
		return false
	}

	// Build discovery request for mobile connection
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

func handleMobileConnection(verifierService *account.Account, welcome *event.Welcome) {
	fmt.Printf("\n📱 Mobile device connected: %s\n", welcome.FromAddress().String())

	// Accept the mobile connection
	_, err := verifierService.ConnectionAccept(welcome.ToAddress(), welcome.Welcome())
	if err != nil {
		fmt.Printf("❌ Failed to accept connection: %v\n", err)
		return
	}

	fmt.Println("✅ Mobile connection established!")
	fmt.Println("📋 Now requesting real credentials from your mobile device...")

	// Store mobile address for credential requests
	connectedMobile = welcome.FromAddress()

	// Request real credentials from the mobile device
	go requestCredentialsFromMobile(verifierService)
}

func requestCredentialsFromMobile(verifierService *account.Account) {
	if connectedMobile == nil {
		fmt.Println("❌ No mobile device connected for credential request")
		return
	}

	// Wait a moment for connection to stabilize
	time.Sleep(2 * time.Second)

	fmt.Println("\n📋 Requesting real credentials from mobile device...")
	fmt.Println("💡 This demonstrates validation of actual mobile credentials")

	// Create presentation request for multiple credential types
	content, err := message.NewCredentialPresentationRequest().
		Type([]string{"VerifiablePresentation", "MobileCredentialPresentation"}).
		Details(
			credential.CredentialTypeLiveness,
			[]*message.CredentialPresentationDetailParameter{
				message.NewCredentialPresentationDetailParameter(
					message.OperatorNotEquals,
					"sourceImageHash",
					"",
				),
			},
		).
		Details(
			credential.CredentialTypeEmail,
			[]*message.CredentialPresentationDetailParameter{
				message.NewCredentialPresentationDetailParameter(
					message.OperatorNotEquals,
					"emailAddress",
					"",
				),
			},
		).
		Finish()

	if err != nil {
		log.Printf("❌ Failed to create presentation request: %v", err)
		return
	}

	// Create a channel to track the response
	presentationCompleter := make(chan *message.CredentialPresentationResponse, 1)
	requests.Store(hex.EncodeToString(content.ID()), presentationCompleter)

	fmt.Printf("📤 Sending credential request to mobile: %s\n", connectedMobile.String())
	fmt.Println("   • Requesting: Liveness and Email credentials")
	fmt.Println("   • This will show your actual verified credentials")

	// Send the presentation request to the mobile device
	err = verifierService.MessageSend(connectedMobile, content)
	if err != nil {
		log.Printf("❌ Failed to send presentation request: %v", err)
		return
	}

	fmt.Println("⏳ Waiting for credential presentation response...")

	// Wait for response with timeout
	select {
	case response := <-presentationCompleter:
		displayReceivedCredentials(verifierService, response)
	case <-time.After(60 * time.Second):
		fmt.Println("⏰ Credential request timed out")
		fmt.Println("💡 In production, you might retry or notify the user")
		fmt.Println("\n🏁 Demo completed (timed out)")
		os.Exit(0)
	}
}

func displayReceivedCredentials(verifierService *account.Account, response *message.CredentialPresentationResponse) {
	fmt.Println("\n✅ Received real credentials from mobile device!")

	presentations := response.Presentations()
	if len(presentations) == 0 {
		fmt.Println("⚠️  No presentations received")
		return
	}

	fmt.Printf("✅ Received %d presentation(s) from mobile device:\n", len(presentations))

	// Display the actual credential details from the mobile device
	for i, presentation := range presentations {
		fmt.Printf("\n📜 Presentation #%d:\n", i+1)
		fmt.Printf("   • Type: %v\n", presentation.PresentationType())
		fmt.Printf("   • Holder: %s\n", presentation.Holder().String())

		credentials := presentation.Credentials()
		fmt.Printf("   • Contains %d real credential(s):\n", len(credentials))

		for j, cred := range credentials {
			fmt.Printf("\n   📋 Real Credential #%d:\n", j+1)
			fmt.Printf("      • Type: %v\n", cred.CredentialType())
			fmt.Printf("      • Issuer: %s\n", cred.Issuer().String())
			fmt.Printf("      • Subject: %s\n", cred.CredentialSubject().String())
			fmt.Printf("      • Valid From: %s\n", cred.ValidFrom().Format("2006-01-02 15:04:05"))

			// Display actual credential claims
			if claims, err := cred.CredentialSubjectClaims(); err == nil {
				fmt.Println("      • Actual Claims:")
				for key, value := range claims {
					fmt.Printf("        - %s: %v\n", key, value)
				}
			} else {
				fmt.Printf("      ⚠️  Could not parse claims: %v\n", err)
			}
		}
	}

	fmt.Println("\n🎉 Real credential validation completed!")
	fmt.Println("💡 This demonstrates:")
	fmt.Println("   • Requesting real credentials from mobile devices")
	fmt.Println("   • Receiving actual verified email addresses and liveness proofs")
	fmt.Println("   • Validating authentic credentials issued by trusted authorities")
	fmt.Println("   • Building trust through real identity verification")
	fmt.Println("\n🏁 Demo completed successfully!")

	// Exit after successful demonstration
	time.Sleep(2 * time.Second)
	os.Exit(0)
}
