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

	"github.com/joinself/academy/examples/server/auth-system/internal/logging"
	"github.com/joinself/self-go-sdk/account"
	"github.com/joinself/self-go-sdk/credential"
	"github.com/joinself/self-go-sdk/event"
	"github.com/joinself/self-go-sdk/keypair/signing"
	"github.com/joinself/self-go-sdk/message"
)

// AuthService manages Self authentication flows with proper multi-user support
type AuthService struct {
	account              *account.Account
	config               *Config
	connections          map[string]*Connection
	sessions             map[string]*Session
	pendingAuth          map[string]*AuthRequest
	discoveryResponseMap map[string]string       // Maps discovery response ID to original content ID
	credentialRequestMap map[string]string       // Maps credential request ID to original content ID
	userChannels         map[string]*UserChannel // Maps user DID to their communication channel
	mutex                sync.RWMutex
	logger               *logging.Logger
}

// Config holds authentication service configuration
type Config struct {
	StoragePath      string
	StorageKey       []byte
	Environment      account.Target
	LogLevel         account.LogLevel
	SessionTimeout   time.Duration
	QRCodeExpiration time.Duration
	RequiredClaims   []string
}

// Connection represents an established connection with a user
type Connection struct {
	ID            string
	UserDID       *signing.PublicKey
	EstablishedAt time.Time
	LastSeen      time.Time
	Status        ConnectionStatus
	ContentID     string // 🔑 Content ID from the original request this connection belongs to
	ChannelID     string // 🔑 Channel ID for this connection
}

// Session represents an authenticated user session
type Session struct {
	ID           string
	UserDID      *signing.PublicKey
	CreatedAt    time.Time
	ExpiresAt    time.Time
	Claims       map[string]interface{}
	ConnectionID string
	ChannelID    string // 🔑 Channel ID for this session
}

// AuthRequest represents a pending authentication request
type AuthRequest struct {
	ID             string
	QRCode         string
	CreatedAt      time.Time
	ExpiresAt      time.Time
	RequiredClaims []string
	CompleteChan   chan *AuthResult
	ContentID      string // 🔑 Content ID from the discovery request for direct mapping
}

// AuthResult contains the result of an authentication attempt
type AuthResult struct {
	Success bool
	Session *Session
	Error   error
	UserDID *signing.PublicKey
	Claims  map[string]interface{}
}

// UserChannel represents a communication channel with a specific user
type UserChannel struct {
	ID                string
	UserDID           *signing.PublicKey
	CreatedAt         time.Time
	LastActivity      time.Time
	ConnectionID      string
	OriginalContentID string // 🔑 The content ID that established this channel
	Status            ChannelStatus
}

// ConnectionStatus represents the state of a connection
type ConnectionStatus int

const (
	ConnectionPending ConnectionStatus = iota
	ConnectionEstablished
	ConnectionDisconnected
)

// ChannelStatus represents the state of a user channel
type ChannelStatus int

const (
	ChannelActive ChannelStatus = iota
	ChannelIdle
	ChannelClosed
)

// NewAuthService creates a new authentication service with enhanced multi-user support
func NewAuthService(config *Config, logger *logging.Logger) (*AuthService, error) {
	// Merge provided config with defaults, filling in missing values
	mergedConfig := mergeConfigWithDefaults(config)

	if logger == nil {
		// Create a default logger if none provided
		slogLogger := slog.Default()
		logger = logging.New(slogLogger)
	}

	service := &AuthService{
		config:               mergedConfig,
		connections:          make(map[string]*Connection),
		sessions:             make(map[string]*Session),
		pendingAuth:          make(map[string]*AuthRequest),
		discoveryResponseMap: make(map[string]string),
		credentialRequestMap: make(map[string]string),
		userChannels:         make(map[string]*UserChannel),
		logger:               logger,
	}

	// Create Self account for authentication service with service callbacks
	accountConfig := &account.Config{
		StorageKey:  mergedConfig.StorageKey,
		StoragePath: mergedConfig.StoragePath,
		Environment: &mergedConfig.Environment,
		LogLevel:    mergedConfig.LogLevel,
		Callbacks: account.Callbacks{
			OnConnect:    service.onConnect,
			OnDisconnect: service.onDisconnect,
			OnWelcome:    service.onWelcome,
			OnMessage:    service.onMessage,
		},
	}

	selfAccount, err := account.New(accountConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create Self account: %w", err)
	}

	service.account = selfAccount

	service.logger.Info("Authentication service initialized with enhanced multi-user support", slog.String("service_did", service.GetServiceDID()))

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

	// Environment: use provided or default
	if config.Environment != (account.Target{}) {
		merged.Environment = config.Environment
	} else {
		merged.Environment = defaults.Environment
	}

	// LogLevel: use provided or default
	if config.LogLevel != 0 {
		merged.LogLevel = config.LogLevel
	} else {
		merged.LogLevel = defaults.LogLevel
	}

	// SessionTimeout: use provided or default
	if config.SessionTimeout != 0 {
		merged.SessionTimeout = config.SessionTimeout
	} else {
		merged.SessionTimeout = defaults.SessionTimeout
	}

	// QRCodeExpiration: use provided or default
	if config.QRCodeExpiration != 0 {
		merged.QRCodeExpiration = config.QRCodeExpiration
	} else {
		merged.QRCodeExpiration = defaults.QRCodeExpiration
	}

	// RequiredClaims: use provided or default
	if config.RequiredClaims != nil && len(config.RequiredClaims) > 0 {
		merged.RequiredClaims = config.RequiredClaims
	} else {
		merged.RequiredClaims = defaults.RequiredClaims
	}

	return merged
}

// DefaultConfig returns a default configuration suitable for development
func DefaultConfig() *Config {
	return &Config{
		StoragePath:      "./auth_service_storage",
		StorageKey:       generateStorageKey("auth_service"),
		Environment:      *account.TargetSandbox,
		LogLevel:         account.LogWarn,
		SessionTimeout:   30 * time.Minute,
		QRCodeExpiration: 5 * time.Minute,
		RequiredClaims:   []string{"liveness"}, // Default to requiring liveness verification
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
	if requiredClaims == nil {
		requiredClaims = a.config.RequiredClaims
	}

	// Generate unique request ID
	requestID := generateID("auth_req")

	// Create QR code for mobile connection and capture content ID
	qrCode, contentID, err := a.generateConnectionQR()
	if err != nil {
		return nil, fmt.Errorf("failed to generate QR code: %w", err)
	}

	// Create auth request
	authRequest := &AuthRequest{
		ID:             requestID,
		QRCode:         qrCode,
		CreatedAt:      time.Now(),
		ExpiresAt:      time.Now().Add(a.config.QRCodeExpiration),
		RequiredClaims: requiredClaims,
		CompleteChan:   make(chan *AuthResult, 1),
		ContentID:      contentID, // 🔑 Store the discovery request content ID
	}

	// Store pending request
	a.mutex.Lock()
	a.pendingAuth[requestID] = authRequest
	a.mutex.Unlock()

	a.logger.Info("Generated auth request", slog.String("request_id", requestID), slog.String("content_id", contentID))

	// Set up cleanup timer
	go a.cleanupExpiredAuthRequest(requestID)

	return authRequest, nil
}

// WaitForAuth waits for authentication to complete for a given request
func (a *AuthService) WaitForAuth(ctx context.Context, requestID string) (*AuthResult, error) {
	a.mutex.RLock()
	authReq, exists := a.pendingAuth[requestID]
	a.mutex.RUnlock()

	if !exists {
		return &AuthResult{Success: false, Error: fmt.Errorf("auth request not found")}, nil
	}

	select {
	case result := <-authReq.CompleteChan:
		// Clean up the auth request
		a.mutex.Lock()
		delete(a.pendingAuth, requestID)
		a.mutex.Unlock()
		return result, nil
	case <-ctx.Done():
		return &AuthResult{Success: false, Error: ctx.Err()}, nil
	case <-time.After(a.config.QRCodeExpiration):
		return &AuthResult{Success: false, Error: fmt.Errorf("authentication timeout")}, nil
	}
}

// ValidateSession checks if a session is valid and returns the session data
func (a *AuthService) ValidateSession(sessionID string) (*Session, error) {
	a.mutex.RLock()
	session, exists := a.sessions[sessionID]
	a.mutex.RUnlock()

	if !exists {
		return nil, fmt.Errorf("session not found")
	}

	if time.Now().After(session.ExpiresAt) {
		// Session expired, clean up
		a.mutex.Lock()
		delete(a.sessions, sessionID)
		a.mutex.Unlock()
		return nil, fmt.Errorf("session expired")
	}

	return session, nil
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

// GetUserSessions returns all active sessions for a user
func (a *AuthService) GetUserSessions(userDID *signing.PublicKey) []*Session {
	a.mutex.RLock()
	defer a.mutex.RUnlock()

	var userSessions []*Session
	for _, session := range a.sessions {
		if session.UserDID.String() == userDID.String() {
			userSessions = append(userSessions, session)
		}
	}

	return userSessions
}

// Close shuts down the authentication service
func (a *AuthService) Close() error {
	a.logger.Info("Shutting down authentication service...")

	// Clean up all pending auth requests
	a.mutex.Lock()
	for requestID, authReq := range a.pendingAuth {
		close(authReq.CompleteChan)
		delete(a.pendingAuth, requestID)
	}

	// Clean up all user channels
	for channelID := range a.userChannels {
		delete(a.userChannels, channelID)
	}

	// Clean up all connections
	for connectionID := range a.connections {
		delete(a.connections, connectionID)
	}

	// Clean up response mappings
	a.discoveryResponseMap = make(map[string]string)
	a.credentialRequestMap = make(map[string]string)
	a.mutex.Unlock()

	// Close Self account
	if a.account != nil {
		a.account.Close()
	}

	a.logger.Info("Authentication service shutdown complete")
	return nil
}

// generateConnectionQR creates a QR code for mobile connection
func (a *AuthService) generateConnectionQR() (string, string, error) {
	// Open inbox for receiving connections
	inboxAddress, err := a.account.InboxOpen()
	if err != nil {
		return "", "", fmt.Errorf("failed to open inbox: %w", err)
	}

	// Generate key package for secure communication
	keyPackage, err := a.account.ConnectionNegotiateOutOfBand(
		inboxAddress,
		time.Now().Add(a.config.QRCodeExpiration),
	)
	if err != nil {
		return "", "", fmt.Errorf("failed to generate key package: %w", err)
	}

	// Build discovery request for mobile connection
	content, err := message.NewDiscoveryRequest().
		KeyPackage(keyPackage).
		Expires(time.Now().Add(a.config.QRCodeExpiration)).
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

func (a *AuthService) onConnect(acc *account.Account) {
	a.logger.Info("Auth service connected to Self network")
}

func (a *AuthService) onDisconnect(acc *account.Account, err error) {
	if err != nil {
		a.logger.Error("Auth service disconnected with error", slog.String("error", err.Error()))
	} else {
		a.logger.Info("Auth service disconnected")
	}
}

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

	// Update channel activity
	a.updateChannelActivity(userDID)

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

// handleDiscoveryResponse processes discovery responses and establishes channels
func (a *AuthService) handleDiscoveryResponse(acc *account.Account, msg *event.Message) {
	discoveryResponse, err := message.DecodeDiscoveryResponse(msg.Content())
	if err != nil {
		a.logger.Error("Failed to decode discovery response", slog.String("error", err.Error()))
		return
	}

	userDID := msg.FromAddress()
	responseToID := hex.EncodeToString(discoveryResponse.ResponseTo())

	a.logger.Info("Received discovery response",
		slog.String("user_did", userDID.String()),
		slog.String("response_to", responseToID))

	// 🔑 Find the matching auth request using the response ID
	a.mutex.Lock()
	var matchingRequest *AuthRequest
	var matchingRequestID string

	for requestID, authReq := range a.pendingAuth {
		if authReq.ContentID == responseToID {
			matchingRequest = authReq
			matchingRequestID = requestID
			break
		}
	}
	a.mutex.Unlock()

	if matchingRequest == nil {
		a.logger.Warn("No matching auth request found for discovery response",
			slog.String("response_to", responseToID),
			slog.String("user_did", userDID.String()))
		return
	}

	// 🔑 Establish user channel and connection
	channelID := generateID("channel")
	connectionID := generateID("conn")

	// Create user channel
	channel := &UserChannel{
		ID:                channelID,
		UserDID:           userDID,
		CreatedAt:         time.Now(),
		LastActivity:      time.Now(),
		ConnectionID:      connectionID,
		OriginalContentID: responseToID,
		Status:            ChannelActive,
	}

	// Create connection record
	connection := &Connection{
		ID:            connectionID,
		UserDID:       userDID,
		EstablishedAt: time.Now(),
		LastSeen:      time.Now(),
		Status:        ConnectionEstablished,
		ContentID:     responseToID,
		ChannelID:     channelID,
	}

	a.mutex.Lock()
	a.userChannels[userDID.String()] = channel
	a.connections[connectionID] = connection
	a.mutex.Unlock()

	a.logger.Info("Channel and connection established",
		slog.String("user_did", userDID.String()),
		slog.String("channel_id", channelID),
		slog.String("connection_id", connectionID),
		slog.String("request_id", matchingRequestID))

	// Now request credentials from the connected user
	go a.requestCredentials(userDID, responseToID)
}

// updateChannelActivity updates the last activity time for a user channel
func (a *AuthService) updateChannelActivity(userDID *signing.PublicKey) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	if channel, exists := a.userChannels[userDID.String()]; exists {
		channel.LastActivity = time.Now()
		a.logger.Debug("Updated channel activity", slog.String("user_did", userDID.String()))
	}
}

// requestCredentials sends a liveness check request to a user
func (a *AuthService) requestCredentials(userDID *signing.PublicKey, contentID string) {
	time.Sleep(2 * time.Second) // Wait for connection to stabilize

	a.logger.Info("Requesting liveness check from user", slog.String("user_did", userDID.String()), slog.String("content_id", contentID))

	// Create presentation request for liveness verification only
	content, err := message.NewCredentialPresentationRequest().
		Type([]string{"VerifiablePresentation"}).
		Details(
			credential.CredentialTypeLiveness,
			[]*message.CredentialPresentationDetailParameter{},
		).
		Finish()

	if err != nil {
		a.logger.Error("Failed to create liveness request", slog.String("error", err.Error()))
		return
	}

	// 🔑 Store the credential request content ID for response matching
	credentialRequestID := hex.EncodeToString(content.ID())
	a.storeCredentialRequestMapping(userDID, credentialRequestID, contentID)

	// Send the liveness check request
	err = a.account.MessageSend(userDID, content)
	if err != nil {
		a.logger.Error("Failed to send liveness request", slog.String("error", err.Error()))
		return
	}

	a.logger.Info("Liveness check sent to", slog.String("user_did", userDID.String()), slog.String("request_id", credentialRequestID))
}

// storeCredentialRequestMapping stores the mapping between credential request ID and original content ID
func (a *AuthService) storeCredentialRequestMapping(userDID *signing.PublicKey, credentialRequestID, originalContentID string) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	a.credentialRequestMap[credentialRequestID] = originalContentID
	a.logger.Info("Stored credential request mapping", slog.String("credential_request_id", credentialRequestID), slog.String("original_content_id", originalContentID))
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

	a.mutex.Lock()
	originalContentID, exists := a.credentialRequestMap[responseToID]
	delete(a.credentialRequestMap, responseToID) // Clean up mapping
	a.mutex.Unlock()

	if !exists {
		a.logger.Warn("No mapping found for credential response ID", slog.String("credential_response_id", responseToID))
		return
	}

	a.logger.Info("Matched credential response", slog.String("credential_response_id", responseToID), slog.String("original_request_content_id", originalContentID))

	// Extract claims from the credentials
	claims := make(map[string]interface{})
	presentations := credentialResponse.Presentations()

	for _, presentation := range presentations {
		for _, cred := range presentation.Credentials() {
			if credClaims, err := cred.CredentialSubjectClaims(); err == nil {
				for key, value := range credClaims {
					claims[key] = value
				}
			}
		}
	}

	// Create session for authenticated user
	sessionID := generateID("sess")
	connectionID := a.findConnectionID(userDID)
	channelID := a.findChannelID(userDID)

	session := &Session{
		ID:           sessionID,
		UserDID:      userDID,
		CreatedAt:    time.Now(),
		ExpiresAt:    time.Now().Add(a.config.SessionTimeout),
		Claims:       claims,
		ConnectionID: connectionID,
		ChannelID:    channelID,
	}

	a.mutex.Lock()
	a.sessions[sessionID] = session
	a.mutex.Unlock()

	a.logger.Info("User authenticated successfully",
		slog.String("user_did", userDID.String()),
		slog.String("session_id", sessionID),
		slog.String("channel_id", channelID))

	// Complete the specific auth request using the original content ID
	a.completeSpecificAuthRequest(originalContentID, &AuthResult{
		Success: true,
		Session: session,
		UserDID: userDID,
		Claims:  claims,
	})
}

// findConnectionID finds the connection ID for a given user DID
func (a *AuthService) findConnectionID(userDID *signing.PublicKey) string {
	a.mutex.RLock()
	defer a.mutex.RUnlock()

	for _, conn := range a.connections {
		if conn.UserDID.String() == userDID.String() {
			return conn.ID
		}
	}
	return ""
}

// findChannelID finds the channel ID for a given user DID
func (a *AuthService) findChannelID(userDID *signing.PublicKey) string {
	a.mutex.RLock()
	defer a.mutex.RUnlock()

	if channel, exists := a.userChannels[userDID.String()]; exists {
		return channel.ID
	}
	return ""
}

// completeSpecificAuthRequest completes a specific authentication request by content ID
func (a *AuthService) completeSpecificAuthRequest(contentID string, result *AuthResult) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	// 🔑 Find the specific request by content ID
	var targetRequestID string
	var targetRequest *AuthRequest

	for requestID, authReq := range a.pendingAuth {
		if authReq.ContentID == contentID {
			targetRequestID = requestID
			targetRequest = authReq
			break
		}
	}

	if targetRequest == nil {
		a.logger.Warn("No pending request found for content ID", slog.String("content_id", contentID))
		return
	}

	// Complete only the specific request
	select {
	case targetRequest.CompleteChan <- result:
		a.logger.Info("Completed auth request", slog.String("request_id", targetRequestID), slog.String("content_id", contentID), slog.String("user_did", result.UserDID.String()))
	default:
		a.logger.Warn("Could not complete auth request (channel full)", slog.String("request_id", targetRequestID))
	}
}

// cleanupExpiredAuthRequest removes expired authentication requests
func (a *AuthService) cleanupExpiredAuthRequest(requestID string) {
	time.Sleep(a.config.QRCodeExpiration)

	a.mutex.Lock()
	defer a.mutex.Unlock()

	if authReq, exists := a.pendingAuth[requestID]; exists {
		if time.Now().After(authReq.ExpiresAt) {
			close(authReq.CompleteChan)
			delete(a.pendingAuth, requestID)
			a.logger.Info("Cleaned up expired auth request", slog.String("request_id", requestID))
		}
	}
}
