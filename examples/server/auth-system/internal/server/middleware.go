// Package server provides middleware for the Self authentication HTTP server.
package server

import (
	"context"
	"log/slog"
	"net/http"
	"strings"
	"time"
)

// ContextKey defines custom types for context keys to avoid collisions
type ContextKey string

const (
	SessionIDKey ContextKey = "session_id"
	UserDIDKey   ContextKey = "user_did"
	ClaimsKey    ContextKey = "claims"
)

// loggingMiddleware logs HTTP requests
func (s *Server) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Wrap the response writer to capture status code
		lrw := &loggingResponseWriter{ResponseWriter: w, statusCode: http.StatusOK}

		next.ServeHTTP(lrw, r)

		duration := time.Since(start)
		s.logger.Info("HTTP request",
			slog.String("method", r.Method),
			slog.String("path", r.URL.Path),
			slog.Int("status_code", lrw.statusCode),
			slog.Duration("duration", duration),
			slog.String("remote_addr", getClientIP(r)),
		)
	})
}

// corsMiddleware handles CORS headers
func (s *Server) corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*") // Configure appropriately for production
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// authMiddleware validates authentication and sets user context
func (s *Server) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Try to get session from cookie
		session, err := s.sessionStore.Get(r, "self-auth-session")
		if err != nil {
			s.logger.Debug("Invalid session cookie", slog.String("error", err.Error()))
			s.sendError(w, "Invalid session", http.StatusUnauthorized)
			return
		}

		sessionID, ok := session.Values["session_id"].(string)
		if !ok || sessionID == "" {
			s.logger.Debug("No active session found")
			s.sendError(w, "No active session", http.StatusUnauthorized)
			return
		}

		// Validate session with auth service
		authSession, err := s.authService.ValidateSession(sessionID)
		if err != nil {
			s.logger.Debug("Session validation failed",
				slog.String("session_id", sessionID),
				slog.String("error", err.Error()),
			)
			s.sendError(w, "Session expired or invalid", http.StatusUnauthorized)
			return
		}

		// Add session information to request context
		ctx := context.WithValue(r.Context(), SessionIDKey, authSession.ID)
		ctx = context.WithValue(ctx, UserDIDKey, authSession.UserDID.String())
		ctx = context.WithValue(ctx, ClaimsKey, authSession.Claims)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// loggingResponseWriter wraps http.ResponseWriter to capture status code
type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

// getClientIP extracts the client IP from the request
func getClientIP(r *http.Request) string {
	// Check X-Forwarded-For header
	forwarded := r.Header.Get("X-Forwarded-For")
	if forwarded != "" {
		// Take the first IP if multiple are present
		ips := strings.Split(forwarded, ",")
		return strings.TrimSpace(ips[0])
	}

	// Check X-Real-IP header
	realIP := r.Header.Get("X-Real-IP")
	if realIP != "" {
		return realIP
	}

	// Fall back to RemoteAddr
	ip := r.RemoteAddr
	if colon := strings.LastIndex(ip, ":"); colon != -1 {
		ip = ip[:colon]
	}
	return ip
}
