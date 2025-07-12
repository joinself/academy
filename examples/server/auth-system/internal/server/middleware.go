// Package server provides middleware for the Self authentication HTTP server.
package server

import (
	"context"
	"log/slog"
	"net/http"
	"strings"
	"time"
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
		ctx := context.WithValue(r.Context(), "session_id", authSession.ID)
		ctx = context.WithValue(ctx, "user_did", authSession.UserDID.String())
		ctx = context.WithValue(ctx, "claims", authSession.Claims)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// rateLimitMiddleware implements basic rate limiting (optional)
func (s *Server) rateLimitMiddleware(requestsPerMinute int) func(http.Handler) http.Handler {
	// This is a simple implementation - for production, use a proper rate limiting library
	clients := make(map[string][]time.Time)

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			clientIP := getClientIP(r)
			now := time.Now()

			// Clean old entries
			if times, exists := clients[clientIP]; exists {
				var validTimes []time.Time
				for _, t := range times {
					if now.Sub(t) < time.Minute {
						validTimes = append(validTimes, t)
					}
				}
				clients[clientIP] = validTimes
			}

			// Check rate limit
			if len(clients[clientIP]) >= requestsPerMinute {
				s.logger.Warn("Rate limit exceeded",
					slog.String("client_ip", clientIP),
					slog.Int("requests_per_minute", requestsPerMinute),
				)
				s.sendError(w, "Rate limit exceeded", http.StatusTooManyRequests)
				return
			}

			// Add current request
			clients[clientIP] = append(clients[clientIP], now)

			next.ServeHTTP(w, r)
		})
	}
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
