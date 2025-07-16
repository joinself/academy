// Package main demonstrates a complete production-ready Self authentication system.
//
// This application showcases how to build a secure, passwordless authentication
// backend using the Self SDK. It includes:
// - QR code generation for mobile connections
// - Credential verification and session management
// - REST API endpoints for authentication flows
// - Middleware for protecting routes
// - Example protected endpoints
//
// 🎯 Educational Goals:
// - Understand production-ready Self authentication patterns
// - Learn modular architecture for authentication systems
// - See how to integrate Self SDK with HTTP servers
// - Experience real-world authentication flows
//
// 🚀 Production Features:
// - Graceful shutdown and signal handling
// - Comprehensive logging and error handling
// - Session management with timeouts
// - Rate limiting and security middleware
// - Configuration management
package main

import (
	"context"
	"encoding/base64"
	"log/slog"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/joinself/academy/examples/solutions/age-verifier/internal/auth"
	"github.com/joinself/academy/examples/solutions/age-verifier/internal/server"
)

func main() {
	// Set up structured logging with simple slog
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})).With("service", "self-auth-system", "version", "1.0.0")

	// Load authentication config directly from environment
	authConfig := auth.DefaultConfig()
	if storagePath := os.Getenv("SELF_AUTH_STORAGE_PATH"); storagePath != "" {
		authConfig.StoragePath = storagePath
	}
	if storageKeyEncoded := os.Getenv("SELF_AUTH_STORAGE_KEY"); storageKeyEncoded != "" {
		if storageKey, err := base64.StdEncoding.DecodeString(storageKeyEncoded); err == nil {
			authConfig.StorageKey = storageKey
		}
	}

	authService, err := auth.NewAuthService(authConfig, logger.With("component", "auth"))
	if err != nil {
		logger.Error("Failed to create auth service", slog.String("error", err.Error()))
		os.Exit(1)
	}
	defer authService.Close()

	// Load server config directly from environment
	serverConfig := server.DefaultServerConfig()
	if address := os.Getenv("SELF_SERVER_ADDRESS"); address != "" {
		serverConfig.Address = address
	}
	if portStr := os.Getenv("SELF_SERVER_PORT"); portStr != "" {
		if port, err := strconv.Atoi(portStr); err == nil {
			serverConfig.Port = port
		}
	}

	httpServer := server.NewServer(authService, serverConfig, logger.With("component", "server"))

	// Set up graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle shutdown signals
	go handleShutdown(ctx, authService, httpServer, logger.With("component", "main"))

	// Start the server
	logger.Info("Starting server",
		slog.String("address", serverConfig.Address),
		slog.Int("port", serverConfig.Port),
	)
	logger.Info("Ready to accept Self authentication requests")

	if err := httpServer.Start(); err != nil {
		logger.Error("Server error", slog.String("error", err.Error()))
	}
}

// handleShutdown manages graceful shutdown
func handleShutdown(ctx context.Context, authService *auth.AuthService, httpServer *server.Server, logger *slog.Logger) {
	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case sig := <-quit:
		logger.Info("Received shutdown signal", slog.String("signal", sig.String()))
	case <-ctx.Done():
		logger.Info("Context cancelled")
	}

	logger.Info("Initiating graceful shutdown")

	// Create shutdown context with timeout
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Shutdown HTTP server
	if err := httpServer.Stop(shutdownCtx); err != nil {
		logger.Error("HTTP server shutdown error", slog.String("error", err.Error()))
	}

	// Close authentication service
	if err := authService.Close(); err != nil {
		logger.Error("Auth service shutdown error", slog.String("error", err.Error()))
	}

	logger.Info("Shutdown complete")
	os.Exit(0)
}
