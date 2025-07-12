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
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joinself/academy/examples/server/auth-system/internal/auth"
	"github.com/joinself/academy/examples/server/auth-system/internal/logging"
	"github.com/joinself/academy/examples/server/auth-system/internal/server"
)

func main() {
	// Set up structured logging
	logConfig := logging.FromEnv()
	slogLogger := logging.Setup(logConfig)
	logger := logging.New(slogLogger).WithComponent("main")

	// Create authentication service
	logger.Info("Initializing authentication service")
	authConfig := auth.DefaultConfig()
	authConfig.SessionTimeout = 1 * time.Hour      // Extended for demo
	authConfig.QRCodeExpiration = 10 * time.Minute // Extended for demo

	authService, err := auth.NewAuthService(authConfig, logging.New(slogLogger).WithComponent("auth"))
	if err != nil {
		logger.Error("Failed to create auth service", slog.String("error", err.Error()))
		os.Exit(1)
	}
	defer authService.Close()

	// Create HTTP server
	logger.Info("Setting up HTTP server")
	serverConfig := server.DefaultServerConfig()
	serverConfig.Port = 8081
	serverConfig.Address = "localhost"

	httpServer := server.NewServer(authService, serverConfig, logging.New(slogLogger).WithComponent("server"))

	// Set up graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle shutdown signals
	go handleShutdown(ctx, authService, httpServer, logging.New(slogLogger).WithComponent("main"))

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
func handleShutdown(ctx context.Context, authService *auth.AuthService, httpServer *server.Server, logger *logging.Logger) {
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
