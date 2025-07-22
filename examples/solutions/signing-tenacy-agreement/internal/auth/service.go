// Package auth provides a production-ready authentication service using Self SDK.
//
// This package implements a complete authentication system that can be plugged
// into production applications while maintaining educational clarity.
package auth

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log/slog"
	"sync"
	"time"

	"github.com/joinself/self-go-sdk/account"
	"github.com/joinself/self-go-sdk/credential"
	"github.com/joinself/self-go-sdk/event"
	"github.com/joinself/self-go-sdk/keypair/signing"
	"github.com/joinself/self-go-sdk/message"
	"github.com/joinself/self-go-sdk/object"
)

// AuthService manages Self authentication flows with unified request tracking
type AuthService struct {
	account      *account.Account
	config       *Config
	authRequests map[string]*AuthRequest // Single source of truth - maps request ID to auth request with embedded state
	sessions     map[string]*Session     // Active user sessions
	signRequests map[string]*SignRequest // Signing requests tracking
	mutex        sync.RWMutex
	logger       *slog.Logger
}

// Config holds essential authentication service configuration
type Config struct {
	StoragePath string // Where to store account data
	StorageKey  []byte // Encryption key for storage (REQUIRED)
}

// Session represents an authenticated user session
type Session struct {
	ID           string
	UserDID      *signing.PublicKey
	CreatedAt    time.Time
	ExpiresAt    time.Time
	Claims       map[string]interface{}
	ConnectionID string
}

// AuthRequest represents a pending authentication request with embedded state
type AuthRequest struct {
	ID             string
	QRCode         string
	CreatedAt      time.Time
	ExpiresAt      time.Time
	RequiredClaims []string
	CompleteChan   chan *AuthResult
	ContentID      string // 🔑 Content ID from the discovery request for direct mapping

	// Embedded state that replaces separate tracking maps
	UserDID             *signing.PublicKey // Set when user connects
	ConnectionID        string             // Set when connection established
	CredentialRequestID string             // Set when credential request sent
	Status              AuthRequestStatus  // Tracks request progress
	ConnectedAt         time.Time          // When user connected
}

// SignRequest represents a pending signing request
type SignRequest struct {
	ID           string
	UserDID      *signing.PublicKey
	AgreementPDF *object.Object
	CreatedAt    time.Time
	ExpiresAt    time.Time
	CompleteChan chan *SignResult
	Status       SignRequestStatus
}

// SignRequestStatus represents the state of a signing request
type SignRequestStatus int

const (
	SignRequestPending SignRequestStatus = iota
	SignRequestSent
	SignRequestCompleted
	SignRequestRejected
	SignRequestExpired
	SignRequestFailed
)

// SignResult contains the result of a signing attempt
type SignResult struct {
	Success bool
	Error   error
	UserDID *signing.PublicKey
	Claims  map[string]interface{}
}

// AuthRequestStatus represents the state of an authentication request
type AuthRequestStatus int

const (
	AuthRequestPending             AuthRequestStatus = iota // Initial state, waiting for user
	AuthRequestConnected                                    // User connected via QR
	AuthRequestCredentialRequested                          // Liveness check sent
	AuthRequestCompleted                                    // Authentication successful
	AuthRequestExpired                                      // Request timed out
	AuthRequestFailed                                       // Authentication failed
)

// AuthResult contains the result of an authentication attempt
type AuthResult struct {
	Success bool
	Session *Session
	Error   error
	UserDID *signing.PublicKey
	Claims  map[string]interface{}
}

// NewAuthService creates a new authentication service with unified request tracking
func NewAuthService(config *Config, logger *slog.Logger) (*AuthService, error) {
	// Merge provided config with defaults, filling in missing values
	mergedConfig := mergeConfigWithDefaults(config)

	if logger == nil {
		// Create a default logger if none provided
		slogLogger := slog.Default()
		logger = slogLogger
	}

	service := &AuthService{
		config:       mergedConfig,
		authRequests: make(map[string]*AuthRequest),
		sessions:     make(map[string]*Session),
		signRequests: make(map[string]*SignRequest),
		logger:       logger,
	}

	// Create Self account for authentication service with service callbacks
	accountConfig := &account.Config{
		StorageKey:  mergedConfig.StorageKey,
		StoragePath: mergedConfig.StoragePath,
		Environment: account.TargetSandbox, // Use sandbox environment for examples
		LogLevel:    account.LogInfo,       // Use info level logging
		Callbacks: account.Callbacks{
			OnWelcome: service.onWelcome,
			OnMessage: service.onMessage,
		},
	}

	selfAccount, err := account.New(accountConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create Self account: %w", err)
	}

	service.account = selfAccount

	service.logger.Info("Authentication service initialized with unified request tracking", slog.String("service_did", service.GetServiceDID()))

	return service, nil
}

// mergeConfigWithDefaults merges a provided config with default values, filling in missing fields
func mergeConfigWithDefaults(config *Config) *Config {
	defaults := DefaultConfig()

	// If no config provided, return defaults
	if config == nil {
		return defaults
	}

	merged := &Config{}

	// StoragePath: use provided or default
	if config.StoragePath != "" {
		merged.StoragePath = config.StoragePath
	} else {
		merged.StoragePath = defaults.StoragePath
	}

	// StorageKey: use provided or default
	if config.StorageKey != nil && len(config.StorageKey) > 0 {
		merged.StorageKey = config.StorageKey
	} else {
		merged.StorageKey = defaults.StorageKey
	}

	return merged
}

// DefaultConfig returns a default configuration suitable for development
func DefaultConfig() *Config {
	return &Config{
		StoragePath: "./auth_service_storage",
		StorageKey:  generateStorageKey("auth_service"),
	}
}

// GetServiceDID returns the DID of the authentication service
func (a *AuthService) GetServiceDID() string {
	inboxAddress, err := a.account.InboxOpen()
	if err != nil {
		a.logger.Error("Failed to get service DID", slog.String("error", err.Error()))
		return ""
	}
	return inboxAddress.String()
}

// GenerateAuthRequest creates a new authentication request with QR code
func (a *AuthService) GenerateAuthRequest(ctx context.Context, requiredClaims []string) (*AuthRequest, error) {
	// Use default liveness claim if none provided
	if requiredClaims == nil || len(requiredClaims) == 0 {
		requiredClaims = []string{"liveness"} // Default to requiring liveness verification
	}

	requestID := generateID("auth")

	qrCode, contentID, err := a.generateConnectionQR()
	if err != nil {
		a.logger.Error("Failed to generate QR code", slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to generate QR code: %w", err)
	}

	authRequest := &AuthRequest{
		ID:             requestID,
		QRCode:         qrCode,
		CreatedAt:      time.Now(),
		ExpiresAt:      time.Now().Add(5 * time.Minute), // Default 5 minute expiration
		RequiredClaims: requiredClaims,
		CompleteChan:   make(chan *AuthResult, 1),
		ContentID:      contentID,          // 🔑 Store the discovery request content ID
		Status:         AuthRequestPending, // Initialize with pending status
	}

	// Store auth request in unified map
	a.mutex.Lock()
	a.authRequests[requestID] = authRequest
	a.mutex.Unlock()

	a.logger.Info("Generated auth request", slog.String("request_id", requestID), slog.String("content_id", contentID))

	// Set up cleanup timer
	go a.cleanupExpiredAuthRequest(requestID)

	return authRequest, nil
}

// WaitForAuth waits for authentication to complete for a given request
func (a *AuthService) WaitForAuth(ctx context.Context, requestID string) (*AuthResult, error) {
	a.mutex.RLock()
	authReq, exists := a.authRequests[requestID]
	a.mutex.RUnlock()

	if !exists {
		return &AuthResult{Success: false, Error: fmt.Errorf("auth request not found")}, nil
	}

	select {
	case result := <-authReq.CompleteChan:
		// Clean up the auth request
		a.mutex.Lock()
		delete(a.authRequests, requestID)
		a.mutex.Unlock()
		return result, nil
	case <-ctx.Done():
		return &AuthResult{Success: false, Error: ctx.Err()}, nil
	case <-time.After(5 * time.Minute): // Default 5 minute timeout
		return &AuthResult{Success: false, Error: fmt.Errorf("authentication timeout")}, nil
	}
}

// RevokeSession invalidates a user session
func (a *AuthService) RevokeSession(sessionID string) error {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	session, exists := a.sessions[sessionID]
	if !exists {
		return fmt.Errorf("session not found")
	}

	delete(a.sessions, sessionID)
	a.logger.Info("Revoked session", slog.String("session_id", sessionID), slog.String("user_did", session.UserDID.String()))

	return nil
}

// Close gracefully shuts down the authentication service
func (a *AuthService) Close() error {
	a.logger.Info("Shutting down authentication service...")

	// Clean up all pending auth requests
	a.mutex.Lock()
	for requestID, authReq := range a.authRequests {
		close(authReq.CompleteChan)
		delete(a.authRequests, requestID)
	}
	a.mutex.Unlock()

	// Close Self account connection
	if err := a.account.Close(); err != nil {
		a.logger.Error("Failed to close Self account", slog.String("error", err.Error()))
		return err
	}

	a.logger.Info("Authentication service shutdown complete")
	return nil
}

// generateConnectionQR creates a QR code for mobile connection
func (a *AuthService) generateConnectionQR() (string, string, error) {
	// Open inbox for receiving mobile connections
	inboxAddress, err := a.account.InboxOpen()
	if err != nil {
		return "", "", fmt.Errorf("failed to open inbox: %w", err)
	}

	// Generate key package for secure communication
	keyPackage, err := a.account.ConnectionNegotiateOutOfBand(
		inboxAddress,
		time.Now().Add(5*time.Minute), // Default 5 minute expiration
	)
	if err != nil {
		return "", "", fmt.Errorf("failed to generate key package: %w", err)
	}

	// Build discovery request for mobile connection
	content, err := message.NewDiscoveryRequest().
		KeyPackage(keyPackage).
		Expires(time.Now().Add(5 * time.Minute)). // Default 5 minute expiration
		Finish()
	if err != nil {
		return "", "", fmt.Errorf("failed to build discovery request: %w", err)
	}

	// Create QR code
	anonymousMsg := event.NewAnonymousMessage(content)
	anonymousMsg.SetFlags(event.MessageFlagTargetSandbox)

	qrCode, err := anonymousMsg.EncodeToQR(event.QREncodingSVG)
	if err != nil {
		return "", "", fmt.Errorf("failed to generate QR code: %w", err)
	}

	return string(qrCode), hex.EncodeToString(content.ID()), nil
}

// Utility functions

func generateStorageKey(seed string) []byte {
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		h := sha256.Sum256([]byte(fmt.Sprintf("self-sdk-%s-%d", seed, time.Now().UnixNano())))
		return h[:]
	}
	return key
}

func generateID(prefix string) string {
	random := make([]byte, 16)
	rand.Read(random)
	return fmt.Sprintf("%s_%s", prefix, hex.EncodeToString(random))
}

// Callback implementations

func (a *AuthService) onWelcome(acc *account.Account, welcome *event.Welcome) {
	a.logger.Info("New connection from", slog.String("from_address", welcome.FromAddress().String()))

	// Accept the mobile connection
	_, err := acc.ConnectionAccept(welcome.ToAddress(), welcome.Welcome())
	if err != nil {
		a.logger.Error("Failed to accept connection", slog.String("error", err.Error()))
		return
	}

	// 🔑 Connection established - we'll wait for discovery response to correlate with request
	a.logger.Info("Connection accepted, waiting for discovery response from", slog.String("from_address", welcome.FromAddress().String()))
}

func (a *AuthService) onMessage(acc *account.Account, msg *event.Message) {
	userDID := msg.FromAddress()

	// Handle different message types
	switch event.ContentTypeOf(msg) {
	case message.ContentTypeDiscoveryResponse:
		a.handleDiscoveryResponse(acc, msg)
	case message.ContentTypeCredentialPresentationResponse:
		a.handleCredentialResponse(acc, msg)

	default:
		a.logger.Info("Received message",
			slog.String("content_type", fmt.Sprintf("%d", event.ContentTypeOf(msg))),
			slog.String("from_address", userDID.String()))
	}
}

// handleDiscoveryResponse processes discovery responses from users
func (a *AuthService) handleDiscoveryResponse(acc *account.Account, msg *event.Message) {
	discoveryResponse, err := message.DecodeDiscoveryResponse(msg.Content())
	if err != nil {
		a.logger.Error("Failed to decode discovery response", slog.String("error", err.Error()))
		return
	}

	userDID := msg.FromAddress()
	a.logger.Info("Received discovery response from", slog.String("user_did", userDID.String()))

	// 🔑 Use SDK's native response matching to find the original request
	responseToID := hex.EncodeToString(discoveryResponse.ResponseTo())

	// Find the auth request by content ID
	a.mutex.Lock()
	var matchingAuthReq *AuthRequest
	var matchingRequestID string
	for reqID, authReq := range a.authRequests {
		if authReq.ContentID == responseToID {
			matchingAuthReq = authReq
			matchingRequestID = reqID
			break
		}
	}

	if matchingAuthReq == nil {
		a.mutex.Unlock()
		a.logger.Warn("No matching auth request found for discovery response", slog.String("response_to_id", responseToID))
		return
	}

	// Update auth request with connection info
	connectionID := generateID("conn")
	matchingAuthReq.UserDID = userDID
	matchingAuthReq.ConnectionID = connectionID
	matchingAuthReq.Status = AuthRequestConnected
	matchingAuthReq.ConnectedAt = time.Now()
	a.mutex.Unlock()

	a.logger.Info("Connection established",
		slog.String("user_did", userDID.String()),
		slog.String("connection_id", connectionID),
		slog.String("request_id", matchingRequestID))

	// Request credentials from the user
	go a.requestCredentials(userDID, responseToID, matchingRequestID)
}

// requestCredentials sends a liveness check request to a user
func (a *AuthService) requestCredentials(userDID *signing.PublicKey, contentID string, requestID string) {
	time.Sleep(2 * time.Second) // Wait for connection to stabilize

	a.logger.Info("Requesting liveness check from user", slog.String("user_did", userDID.String()), slog.String("content_id", contentID))

	// Create presentation request for liveness verification only
	content, err := message.NewCredentialPresentationRequest().
		Type([]string{"VerifiablePresentation"}).
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
		Finish()

	if err != nil {
		a.logger.Error("Failed to create liveness request", slog.String("error", err.Error()))
		return
	}

	// 🔑 Store the credential request content ID for response matching
	credentialRequestID := hex.EncodeToString(content.ID())

	// Send the liveness check request
	err = a.account.MessageSend(userDID, content)
	if err != nil {
		a.logger.Error("Failed to send credential request", slog.String("error", err.Error()))
		return
	}

	// Update auth request with credential request info
	a.mutex.Lock()
	if authReq, exists := a.authRequests[requestID]; exists {
		authReq.CredentialRequestID = credentialRequestID
		authReq.Status = AuthRequestCredentialRequested
	}
	a.mutex.Unlock()

	a.logger.Info("Liveness check sent", slog.String("credential_request_id", credentialRequestID))
}

// handleCredentialResponse processes liveness check responses
func (a *AuthService) handleCredentialResponse(acc *account.Account, msg *event.Message) {
	credentialResponse, err := message.DecodeCredentialPresentationResponse(msg.Content())
	if err != nil {
		a.logger.Error("Failed to decode liveness response", slog.String("error", err.Error()))
		return
	}

	userDID := msg.FromAddress()
	a.logger.Info("Received liveness proof from", slog.String("user_did", userDID.String()))

	// 🔑 Use SDK's native response matching to find the original request
	responseToID := hex.EncodeToString(credentialResponse.ResponseTo())

	// Find the auth request by credential request ID
	a.mutex.Lock()
	var matchingAuthReq *AuthRequest
	var originalContentID string
	for _, authReq := range a.authRequests {
		if authReq.CredentialRequestID == responseToID {
			matchingAuthReq = authReq
			originalContentID = authReq.ContentID
			break
		}
	}
	a.mutex.Unlock()

	if matchingAuthReq == nil {
		a.logger.Warn("No matching auth request found for credential response", slog.String("credential_response_id", responseToID))
		return
	}

	// Extract claims from the credential
	claims := make(map[string]interface{})
	presentations := credentialResponse.Presentations()

	if len(presentations) > 0 {
		for _, presentation := range presentations {
			for _, cred := range presentation.Credentials() {
				if credClaims, err := cred.CredentialSubjectClaims(); err == nil {
					for key, value := range credClaims {
						claims[key] = value
					}
				}
			}
		}
	}

	// Create session for authenticated user
	sessionID := generateID("sess")

	session := &Session{
		ID:           sessionID,
		UserDID:      userDID,
		CreatedAt:    time.Now(),
		ExpiresAt:    time.Now().Add(30 * time.Minute), // Default 30 minute session timeout
		Claims:       claims,
		ConnectionID: matchingAuthReq.ConnectionID,
	}

	a.mutex.Lock()
	a.sessions[sessionID] = session
	matchingAuthReq.Status = AuthRequestCompleted
	a.mutex.Unlock()

	a.logger.Info("User authenticated successfully",
		slog.String("user_did", userDID.String()),
		slog.String("session_id", sessionID))

	// Complete the specific auth request using the original content ID
	a.completeSpecificAuthRequest(originalContentID, &AuthResult{
		Success: true,
		Session: session,
		UserDID: userDID,
		Claims:  claims,
	})
}

// completeSpecificAuthRequest completes a specific authentication request by content ID
func (a *AuthService) completeSpecificAuthRequest(contentID string, result *AuthResult) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	// 🔑 Find the specific request by content ID
	var targetRequestID string
	for requestID, authReq := range a.authRequests {
		if authReq.ContentID == contentID {
			targetRequestID = requestID
			break
		}
	}

	if targetRequestID == "" {
		a.logger.Warn("No auth request found for content ID", slog.String("content_id", contentID))
		return
	}

	authReq := a.authRequests[targetRequestID]

	// Send result through the channel
	select {
	case authReq.CompleteChan <- result:
		a.logger.Info("Auth request completed", slog.String("request_id", targetRequestID), slog.Bool("success", result.Success))
	default:
		a.logger.Warn("Failed to send auth result, channel may be closed", slog.String("request_id", targetRequestID))
	}
}

// cleanupExpiredAuthRequest removes expired authentication requests
func (a *AuthService) cleanupExpiredAuthRequest(requestID string) {
	// Wait for expiration time
	a.mutex.RLock()
	authReq, exists := a.authRequests[requestID]
	a.mutex.RUnlock()

	if !exists {
		return
	}

	time.Sleep(time.Until(authReq.ExpiresAt))

	// Check if request is still pending and remove it
	a.mutex.Lock()
	if authReq, exists := a.authRequests[requestID]; exists && authReq.Status == AuthRequestPending {
		authReq.Status = AuthRequestExpired
		close(authReq.CompleteChan)
		delete(a.authRequests, requestID)
		a.logger.Info("Cleaned up expired auth request", slog.String("request_id", requestID))
	}
	a.mutex.Unlock()
}

// ============================================================================
// SIGNING WORKFLOW METHODS
// ============================================================================

// CreateTenancyAgreement creates a PDF agreement for signing
func (a *AuthService) CreateTenancyAgreement(tenantName, propertyAddress, rentAmount, startDate, endDate string) (*object.Object, error) {
	// Create PDF content for the tenancy agreement
	pdfContent := fmt.Sprintf(`TENANCY AGREEMENT

This agreement is made on %s between:

LANDLORD: Self Property Management Ltd.
TENANT: %s

PROPERTY: %s

TERMS:
- Monthly Rent: $%s
- Start Date: %s
- End Date: %s
- Security Deposit: $%s

CONDITIONS:
1. The tenant agrees to pay rent on time
2. The property must be kept in good condition
3. No pets without written permission
4. 30 days notice required for termination

This agreement is legally binding and constitutes a digital signature
when accepted through the Self platform.

Agreement ID: %s
Created: %s`,
		time.Now().Format("January 2, 2006"),
		tenantName,
		propertyAddress,
		rentAmount,
		startDate,
		endDate,
		rentAmount,
		generateID("agreement"),
		time.Now().Format("2006-01-02 15:04:05"))

	// Create evidence object for the PDF
	agreementPDF, err := object.New("application/pdf", []byte(pdfContent))
	if err != nil {
		a.logger.Error("Failed to create agreement PDF", slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to create agreement PDF: %w", err)
	}

	// Upload the PDF to secure storage
	err = a.account.ObjectUpload(agreementPDF, true)
	if err != nil {
		a.logger.Error("Failed to upload agreement PDF", slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to upload agreement PDF: %w", err)
	}

	a.logger.Info("Tenancy agreement PDF created and uploaded",
		slog.String("agreement_id", fmt.Sprintf("%x", agreementPDF.Id())),
		slog.Int("size_bytes", len(pdfContent)))

	return agreementPDF, nil
}

// GenerateSignRequest creates a new signing request for an authenticated user
func (a *AuthService) GenerateSignRequest(ctx context.Context, userDID *signing.PublicKey, agreementPDF *object.Object) (*SignRequest, error) {
	requestID := generateID("sign")

	signRequest := &SignRequest{
		ID:           requestID,
		UserDID:      userDID,
		AgreementPDF: agreementPDF,
		CreatedAt:    time.Now(),
		ExpiresAt:    time.Now().Add(10 * time.Minute), // 10 minute expiration for signing
		CompleteChan: make(chan *SignResult, 1),
		Status:       SignRequestPending,
	}

	// Store signing request
	a.mutex.Lock()
	a.signRequests[requestID] = signRequest
	a.mutex.Unlock()

	a.logger.Info("Generated signing request", slog.String("request_id", requestID), slog.String("user_did", userDID.String()))

	// Set up cleanup timer
	go a.cleanupExpiredSignRequest(requestID)

	// Send the signing request to the user
	go a.sendSigningRequest(userDID, agreementPDF, requestID)

	return signRequest, nil
}

// WaitForSign waits for a signing request to complete
func (a *AuthService) WaitForSign(ctx context.Context, requestID string) (*SignResult, error) {
	a.mutex.RLock()
	signReq, exists := a.signRequests[requestID]
	a.mutex.RUnlock()

	if !exists {
		return &SignResult{Success: false, Error: fmt.Errorf("signing request not found")}, nil
	}

	select {
	case result := <-signReq.CompleteChan:
		// Clean up the signing request
		a.mutex.Lock()
		delete(a.signRequests, requestID)
		a.mutex.Unlock()
		return result, nil
	case <-ctx.Done():
		return &SignResult{Success: false, Error: ctx.Err()}, nil
	case <-time.After(10 * time.Minute): // 10 minute timeout
		return &SignResult{Success: false, Error: fmt.Errorf("signing timeout")}, nil
	}
}

// sendSigningRequest sends a credential issuance request to sign the agreement
func (a *AuthService) sendSigningRequest(userDID *signing.PublicKey, agreementPDF *object.Object, requestID string) {
	time.Sleep(2 * time.Second) // Wait for request to be stored

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
		CredentialSubject(credential.AddressKey(userDID)).
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

// completeSignRequest completes a signing request with the given result
func (a *AuthService) completeSignRequest(requestID string, result *SignResult) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	signReq, exists := a.signRequests[requestID]
	if !exists {
		a.logger.Warn("No signing request found for completion", slog.String("request_id", requestID))
		return
	}

	// Update status based on result
	if result.Success {
		signReq.Status = SignRequestCompleted
	} else {
		signReq.Status = SignRequestFailed
	}

	// Send result through the channel
	select {
	case signReq.CompleteChan <- result:
		a.logger.Info("Signing request completed", slog.String("request_id", requestID), slog.Bool("success", result.Success))
	default:
		a.logger.Warn("Failed to send signing result, channel may be closed", slog.String("request_id", requestID))
	}
}

// cleanupExpiredSignRequest removes expired signing requests
func (a *AuthService) cleanupExpiredSignRequest(requestID string) {
	// Wait for expiration time
	a.mutex.RLock()
	signReq, exists := a.signRequests[requestID]
	a.mutex.RUnlock()

	if !exists {
		return
	}

	time.Sleep(time.Until(signReq.ExpiresAt))

	// Check if request is still pending and remove it
	a.mutex.Lock()
	if signReq, exists := a.signRequests[requestID]; exists && signReq.Status == SignRequestPending {
		signReq.Status = SignRequestExpired
		close(signReq.CompleteChan)
		delete(a.signRequests, requestID)
		a.logger.Info("Cleaned up expired signing request", slog.String("request_id", requestID))
	}
	a.mutex.Unlock()
}

// handleCredentialIssuanceResponse processes responses from credential issuance (signing requests)
func (a *AuthService) handleCredentialIssuanceResponse(acc *account.Account, msg *event.Message) {
	userDID := msg.FromAddress()
	a.logger.Info("Received credential issuance response from", slog.String("user_did", userDID.String()))

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

	// Extract claims from the credential response
	claims := make(map[string]interface{})

	// For credential issuance responses, we need to decode the credential
	// This is a simplified implementation - in a real scenario, you'd decode the actual credential
	claims["agreementType"] = "Tenancy Agreement"
	claims["signingDate"] = time.Now().Format("2006-01-02")
	claims["signingTimestamp"] = time.Now().Unix()
	claims["documentType"] = "application/pdf"
	claims["evidenceId"] = fmt.Sprintf("%x", matchingSignReq.AgreementPDF.Id())

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
