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
)

// AuthService manages Self authentication flows with unified request tracking
type AuthService struct {
	account      *account.Account
	config       *Config
	authRequests map[string]*AuthRequest // Single source of truth - maps request ID to auth request with embedded state
	sessions     map[string]*Session     // Active user sessions
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

// GetRequestStatus returns the current status of an auth request without blocking
func (a *AuthService) GetRequestStatus(requestID string) (string, *AuthResult) {
	a.mutex.RLock()
	defer a.mutex.RUnlock()

	authReq, exists := a.authRequests[requestID]
	if !exists {
		return "not_found", nil
	}

	switch authReq.Status {
	case AuthRequestPending:
		return "pending", nil
	case AuthRequestConnected:
		return "connected", nil
	case AuthRequestCredentialRequested:
		return "credential_requested", nil
	case AuthRequestCompleted:
		// Check if there's a result in the completion channel without blocking
		select {
		case result := <-authReq.CompleteChan:
			return "completed", result
		default:
			return "completed", nil
		}
	case AuthRequestExpired:
		return "expired", nil
	case AuthRequestFailed:
		return "failed", &AuthResult{
			Success: false,
			Error:   fmt.Errorf("authentication request failed"),
		}
	default:
		return "pending", nil
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

// requestCredentials sends a dateOfBirth credential request to a user for age verification
func (a *AuthService) requestCredentials(userDID *signing.PublicKey, contentID string, requestID string) {
	time.Sleep(2 * time.Second) // Wait for connection to stabilize

	a.logger.Info("Requesting age verification from user", slog.String("user_did", userDID.String()), slog.String("content_id", contentID))

	// Calculate the date 18 years ago for age verification
	eighteenYearsAgo := time.Now().AddDate(-18, 0, 0).Format("2006-01-02")

	// Create presentation request for age verification (18+ years old)
	// Uses conditional disclosure: credential shared only if user is 18+, otherwise no response
	content, err := message.NewCredentialPresentationRequest().
		Type([]string{"VerifiablePresentation"}).
		Details(
			credential.CredentialTypePassport,
			[]*message.CredentialPresentationDetailParameter{
				message.NewCredentialPresentationDetailParameter(
					message.OperatorLessThan,
					"dateOfBirth",
					eighteenYearsAgo,
				),
			},
		).
		Finish()

	if err != nil {
		a.logger.Error("Failed to create age verification request", slog.String("error", err.Error()))
		return
	}

	// 🔑 Store the credential request content ID for response matching
	credentialRequestID := hex.EncodeToString(content.ID())

	// Send the age verification request
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

	a.logger.Info("Age verification request sent", slog.String("credential_request_id", credentialRequestID))
}

// handleCredentialResponse processes age verification responses
func (a *AuthService) handleCredentialResponse(acc *account.Account, msg *event.Message) {
	credentialResponse, err := message.DecodeCredentialPresentationResponse(msg.Content())
	if err != nil {
		a.logger.Error("Failed to decode age verification response", slog.String("error", err.Error()))
		return
	}

	userDID := msg.FromAddress()
	a.logger.Info("Received age verification from", slog.String("user_did", userDID.String()))

	// 🔑 Use SDK's native response matching to find the original request
	responseToID := hex.EncodeToString(credentialResponse.ResponseTo())

	// Find the auth request by credential request ID
	a.mutex.Lock()
	var matchingAuthReq *AuthRequest
	for _, authReq := range a.authRequests {
		if authReq.CredentialRequestID == responseToID {
			matchingAuthReq = authReq
			break
		}
	}
	a.mutex.Unlock()

	if matchingAuthReq == nil {
		a.logger.Warn("No matching auth request found for credential response", slog.String("credential_response_id", responseToID))
		return
	}

	// Log for debugging and validate that we received credentials with the required claim
	a.logger.Info("Checking for presentation data for debugging...", slog.Int("presentations_count", len(credentialResponse.Presentations())))
	if len(credentialResponse.Presentations()) > 0 {
		a.logger.Info("Checking for credentials in presentation for debugging...", slog.Int("credentials_count", len(credentialResponse.Presentations()[0].Credentials())))
		if len(credentialResponse.Presentations()[0].Credentials()) > 0 {
			if claims, err := credentialResponse.Presentations()[0].Credentials()[0].CredentialSubjectClaims(); err == nil {
				a.logger.Info("Extracted claims for debugging", slog.Any("claims", claims))
				if dob, ok := claims["dateOfBirth"]; ok {
					// --- SUCCESS CASE ---
					a.logger.Info("Retrieved date of birth for debugging", slog.Any("dateOfBirth", dob))
					a.logger.Info("Age verification successful - dateOfBirth claim received", slog.String("user_did", userDID.String()))

					// Create session with user details
					sessionClaims := make(map[string]interface{})
					sessionClaims["ageVerified"] = true
					sessionClaims["dateOfBirth"] = dob

					session := &Session{
						ID:           generateID("sess"),
						UserDID:      userDID,
						CreatedAt:    time.Now(),
						ExpiresAt:    time.Now().Add(24 * time.Hour), // 24-hour session
						Claims:       sessionClaims,
						ConnectionID: matchingAuthReq.ConnectionID,
					}

					// Store the session
					a.mutex.Lock()
					a.sessions[session.ID] = session
					matchingAuthReq.Status = AuthRequestCompleted
					a.mutex.Unlock()

					// Complete authentication successfully
					result := &AuthResult{
						Success: true,
						Session: session,
						UserDID: userDID,
						Claims:  sessionClaims,
					}

					a.logger.Info("Authentication completed successfully",
						slog.String("user_did", userDID.String()),
						slog.String("session_id", session.ID))

					// Send result on the auth request completion channel
					select {
					case matchingAuthReq.CompleteChan <- result:
					default:
						a.logger.Warn("Failed to send auth result - channel may be closed")
					}
					return // End execution for success case
				}
			}
		}
	}

	// --- FAILURE CASE ---
	// If we reach here, it means the dateOfBirth claim was not found.
	failureReason := "age verification failed: dateOfBirth claim not found in response"
	a.logger.Error(failureReason, slog.String("user_did", userDID.String()))

	a.mutex.Lock()
	matchingAuthReq.Status = AuthRequestFailed
	a.mutex.Unlock()

	result := &AuthResult{
		Success: false,
		Error:   fmt.Errorf(failureReason),
		UserDID: userDID,
	}

	// Send failure result
	select {
	case matchingAuthReq.CompleteChan <- result:
	default:
		a.logger.Warn("Failed to send auth failure result - channel may be closed")
	}
}
