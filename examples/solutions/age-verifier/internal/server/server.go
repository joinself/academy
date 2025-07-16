// Package server provides HTTP server functionality for the Self authentication system.
package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/joinself/academy/examples/solutions/age-verifier/internal/auth"
)

// Server wraps the HTTP server with authentication functionality
type Server struct {
	authService *auth.AuthService
	httpServer  *http.Server
	logger      *slog.Logger
}

// Config holds essential server configuration
type Config struct {
	Address string // Server address (default: localhost)
	Port    int    // Server port (default: 8081)
}

// AuthRequest represents an authentication request
type AuthRequest struct {
	RequiredClaims []string `json:"required_claims,omitempty"`
}

// AuthResponse represents an authentication response
type AuthResponse struct {
	Success   bool   `json:"success"`
	RequestID string `json:"request_id"`
	QRCode    string `json:"qr_code"`
	ExpiresAt string `json:"expires_at"`
	Error     string `json:"error,omitempty"`
}

// AuthStatusResponse represents the response for auth status check
type AuthStatusResponse struct {
	Status  string       `json:"status"`
	Session *SessionData `json:"session,omitempty"`
	Error   string       `json:"error,omitempty"`
}

// SessionData represents session information for the frontend
type SessionData struct {
	ID      string                 `json:"id"`
	UserDID string                 `json:"user_did"`
	Claims  map[string]interface{} `json:"claims"`
}

// ErrorResponse represents an API error response
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message,omitempty"`
	Code    int    `json:"code"`
}

// NewServer creates a new HTTP server with authentication endpoints
func NewServer(authService *auth.AuthService, config *Config, logger *slog.Logger) *Server {
	if config == nil {
		config = DefaultServerConfig()
	}
	if logger == nil {
		logger = slog.Default()
	}

	server := &Server{
		authService: authService,
		logger:      logger,
	}

	// Set up routes
	router := server.setupRoutes()

	// Configure HTTP server with sensible defaults
	server.httpServer = &http.Server{
		Addr:         fmt.Sprintf("%s:%d", config.Address, config.Port),
		Handler:      router,
		ReadTimeout:  30 * time.Second, // Default timeout
		WriteTimeout: 30 * time.Second, // Default timeout
	}

	return server
}

// DefaultServerConfig returns a default server configuration
func DefaultServerConfig() *Config {
	return &Config{
		Address: "localhost",
		Port:    8081,
	}
}

// Start starts the HTTP server
func (s *Server) Start() error {
	s.logger.Info("Starting authentication server", slog.String("address", s.httpServer.Addr))
	return s.httpServer.ListenAndServe()
}

// Stop gracefully stops the HTTP server
func (s *Server) Stop(ctx context.Context) error {
	s.logger.Info("Stopping authentication server")
	return s.httpServer.Shutdown(ctx)
}

// setupRoutes configures all HTTP routes
func (s *Server) setupRoutes() *mux.Router {
	router := mux.NewRouter()

	// API routes (updated to match frontend expectations)
	api := router.PathPrefix("/api").Subrouter()

	// Authentication endpoints (simplified path structure)
	auth := api.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/start", s.handleAuthStart).Methods("POST", "OPTIONS")
	auth.HandleFunc("/status/{requestId}", s.handleAuthStatus).Methods("GET", "OPTIONS")

	// Static files (for demo UI)
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./web/")))

	return router
}

// handleAuthStart starts a new age verification request
func (s *Server) handleAuthStart(w http.ResponseWriter, r *http.Request) {
	var req AuthRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response := AuthResponse{
			Success: false,
			Error:   "Invalid request body",
		}
		s.sendJSON(w, response, http.StatusBadRequest)
		return
	}

	// Generate authentication request with dateOfBirth requirement
	authReq, err := s.authService.GenerateAuthRequest(r.Context(), req.RequiredClaims)
	if err != nil {
		s.logger.Error("Failed to generate age verification request", slog.String("error", err.Error()))
		response := AuthResponse{
			Success: false,
			Error:   "Failed to generate age verification request",
		}
		s.sendJSON(w, response, http.StatusInternalServerError)
		return
	}

	response := AuthResponse{
		Success:   true,
		RequestID: authReq.ID,
		QRCode:    authReq.QRCode,
		ExpiresAt: authReq.ExpiresAt.Format(time.RFC3339),
	}

	s.logger.Info("Generated age verification request", slog.String("request_id", authReq.ID))
	s.sendJSON(w, response, http.StatusOK)
}

// handleAuthStatus checks the status of an age verification request
func (s *Server) handleAuthStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	requestID := vars["requestId"]

	if requestID == "" {
		response := AuthStatusResponse{
			Status: "failed",
			Error:  "Request ID is required",
		}
		s.sendJSON(w, response, http.StatusBadRequest)
		return
	}

	// Check current status without waiting (for polling)
	status, result := s.authService.GetRequestStatus(requestID)

	response := AuthStatusResponse{
		Status: status,
	}

	if result != nil {
		if result.Success {
			response.Session = &SessionData{
				ID:      result.Session.ID,
				UserDID: result.UserDID.String(),
				Claims:  result.Claims,
			}
			s.logger.Info("Age verification completed successfully",
				slog.String("request_id", requestID),
				slog.String("user_did", result.UserDID.String()),
				slog.String("session_id", result.Session.ID))
		} else {
			response.Status = "failed"
			if result.Error != nil {
				response.Error = result.Error.Error()
			}
			s.logger.Info("Age verification failed",
				slog.String("request_id", requestID),
				slog.String("error", response.Error))
		}
	}

	s.sendJSON(w, response, http.StatusOK)
}

// Helper methods

func (s *Server) sendJSON(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func (s *Server) sendError(w http.ResponseWriter, message string, statusCode int) {
	errorResp := ErrorResponse{
		Error:   http.StatusText(statusCode),
		Message: message,
		Code:    statusCode,
	}
	s.sendJSON(w, errorResp, statusCode)
}
