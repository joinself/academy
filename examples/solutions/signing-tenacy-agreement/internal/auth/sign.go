package auth

import (
	"encoding/hex"
	"fmt"
	"log/slog"
	"time"

	"github.com/joinself/self-go-sdk/account"
	"github.com/joinself/self-go-sdk/credential"
	"github.com/joinself/self-go-sdk/event"
	"github.com/joinself/self-go-sdk/keypair/signing"
	"github.com/joinself/self-go-sdk/message"
	"github.com/joinself/self-go-sdk/object"
)

// sendSigningRequest sends a credential issuance request to sign the agreement
func (a *AuthService) sendSigningRequest(userDID *signing.PublicKey, agreementPDF *object.Object, requestID string) {
	// time.Sleep(2 * time.Second) // Wait for request to be stored

	a.logger.Info("Sending signing request to user", slog.String("user_did", userDID.String()), slog.String("request_id", requestID))

	// Create claims for the signing credential
	claims := map[string]interface{}{
		"agreementType":    "Tenancy Agreement",
		"agreementId":      fmt.Sprintf("%x", agreementPDF.Id()),
		"documentHash":     fmt.Sprintf("%x", agreementPDF.Hash()),
		"signingDate":      time.Now().Format("2006-01-02"),
		"signingTimestamp": time.Now().Unix(),
		"documentType":     "application/pdf",
		"evidenceId":       fmt.Sprintf("%x", agreementPDF.Id()),
	}

	// Get service address for issuing
	serviceAddress, err := a.account.InboxOpen()
	if err != nil {
		a.logger.Error("Failed to get service address", slog.String("error", err.Error()))
		a.completeSignRequest(requestID, &SignResult{Success: false, Error: err})
		return
	}

	// Create credential for signing
	credentialBuilder := credential.NewCredential().
		CredentialType([]string{"VerifiableCredential", "AgreementCredential"}).
		CredentialSubject(credential.AddressKey(serviceAddress)).
		CredentialSubjectClaims(claims).
		Issuer(credential.AddressKey(serviceAddress)).
		ValidFrom(time.Now()).
		SignWith(serviceAddress, time.Now())

	unsignedCredential, err := credentialBuilder.Finish()
	if err != nil {
		a.logger.Error("Failed to build signing credential", slog.String("error", err.Error()))
		a.completeSignRequest(requestID, &SignResult{Success: false, Error: err})
		return
	}

	// Issue the credential
	signedAgreementCredential, err := a.account.CredentialIssue(unsignedCredential)
	if err != nil {
		a.logger.Error("Failed to issue signing credential", slog.String("error", err.Error()))
		a.completeSignRequest(requestID, &SignResult{Success: false, Error: err})
		return
	}

	// Create presentation with the signed credential
	unsignedAgreementPresentation, err := credential.NewPresentation().
		PresentationType([]string{"VerifiablePresentation", "AgreementPresentation"}).
		Holder(credential.AddressKey(serviceAddress)).
		CredentialAdd(signedAgreementCredential).
		Finish()

	if err != nil {
		a.logger.Error("Failed to create presentation", slog.String("error", err.Error()))
		a.completeSignRequest(requestID, &SignResult{Success: false, Error: err})
		return
	}

	// Sign the presentation
	signedAgreementPresentation, err := a.account.PresentationIssue(unsignedAgreementPresentation)
	if err != nil {
		a.logger.Error("Failed to issue presentation", slog.String("error", err.Error()))
		a.completeSignRequest(requestID, &SignResult{Success: false, Error: err})
		return
	}

	// Create credential verification request with the agreement evidence
	content, err := message.NewCredentialVerificationRequest().
		Type([]string{"VerifiableCredential", "AgreementCredential"}).
		Evidence("agreement", agreementPDF).
		Proof(signedAgreementPresentation).
		Expires(time.Now().Add(time.Hour * 24)).
		Finish()

	if err != nil {
		a.logger.Error("Failed to build verification request", slog.String("error", err.Error()))
		a.completeSignRequest(requestID, &SignResult{Success: false, Error: err})
		return
	}

	// Send the signing request to the user
	err = a.account.MessageSend(userDID, content)
	if err != nil {
		a.logger.Error("Failed to send signing request", slog.String("error", err.Error()), slog.String("user_did", userDID.String()))
		a.completeSignRequest(requestID, &SignResult{Success: false, Error: err})
		return
	}

	// Update request status
	a.mutex.Lock()
	if signReq, exists := a.signRequests[requestID]; exists {
		signReq.Status = SignRequestSent
	}
	a.mutex.Unlock()

	a.logger.Info("Signing request sent successfully", slog.String("request_id", requestID), slog.String("user_did", userDID.String()))
}

// handleCredentialVerificationResponse processes responses from credential verification (signing requests)
func (a *AuthService) handleCredentialVerificationResponse(acc *account.Account, msg *event.Message) {
	userDID := msg.FromAddress()
	a.logger.Info("Received credential verification response from", slog.String("user_did", userDID.String()))

	// Find the signing request for this user
	a.mutex.Lock()
	var matchingSignReq *SignRequest
	var requestID string
	for reqID, signReq := range a.signRequests {
		if signReq.UserDID.String() == userDID.String() && signReq.Status == SignRequestSent {
			matchingSignReq = signReq
			requestID = reqID
			break
		}
	}
	a.mutex.Unlock()

	if matchingSignReq == nil {
		a.logger.Warn("No matching signing request found for user", slog.String("user_did", userDID.String()))
		return
	}

	// Extract claims from the credential verification response
	claims := make(map[string]interface{})

	// Decode the credential verification response
	credentialVerificationResponse, err := message.DecodeCredentialVerificationResponse(msg.Content())
	if err != nil {
		a.logger.Error("Failed to decode credential verification response", slog.String("error", err.Error()))
		a.completeSignRequest(requestID, &SignResult{Success: false, Error: err})
		return
	}

	// Log the response details for debugging
	a.logger.Info("Processing credential verification response",
		slog.String("response_to_id", hex.EncodeToString(credentialVerificationResponse.ResponseTo())))

	// Extract claims from the credential verification response
	var presentations []*credential.VerifiablePresentation
	if presentationsMethod, ok := interface{}(credentialVerificationResponse).(interface {
		Presentations() []*credential.VerifiablePresentation
	}); ok {
		presentations = presentationsMethod.Presentations()
		a.logger.Info("Found presentations in verification response", slog.Int("presentations_count", len(presentations)))
	} else {
		a.logger.Info("No Presentations() method found on verification response, checking for alternative structure")
	}

	for i, presentation := range presentations {
		a.logger.Info("Processing presentation", slog.Int("presentation_index", i))

		// Extract credentials from the presentation
		credentials := presentation.Credentials()
		a.logger.Info("Found credentials in presentation", slog.Int("credentials_count", len(credentials)))

		for j, cred := range credentials {
			a.logger.Info("Processing credential", slog.Int("credential_index", j))

			// Extract claims from the credential
			if credClaims, err := cred.CredentialSubjectClaims(); err == nil {
				a.logger.Info("Extracted credential claims", slog.Any("claims", credClaims))

				// Merge claims into our result
				for key, value := range credClaims {
					claims[key] = value
				}
			} else {
				a.logger.Warn("Failed to extract claims from credential", slog.String("error", err.Error()))
			}
		}
	}

	// Complete the signing request successfully
	a.completeSignRequest(requestID, &SignResult{
		Success: true,
		UserDID: userDID,
		Claims:  claims,
	})

	a.logger.Info("Signing request completed successfully",
		slog.String("request_id", requestID),
		slog.String("user_did", userDID.String()))
}
